package tracer

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Tracer(tr opentracing.Tracer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			operationName := GetOperationName(r.Context())
			if operationName == "" {
				operationName = fmt.Sprintf("%s %s", r.Method, r.URL.Path)
			}
			// Pass request through tracer
			serverSpanCtx, _ := tr.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
			span, traceCtx := opentracing.StartSpanFromContextWithTracer(r.Context(), tr, operationName, ext.RPCServerOption(serverSpanCtx))
			defer span.Finish()

			defer func() {
				if err := recover(); err != nil {
					ext.HTTPStatusCode.Set(span, uint16(500))
					ext.Error.Set(span, true)
					span.SetTag("error.type", "panic")
					span.LogKV(
						"event", "error",
						"error.kind", "panic",
						"message", err,
						"stack", string(debug.Stack()),
					)
					span.Finish()

					panic(err)
				}
			}()

			ext.SpanKind.Set(span, ext.SpanKindRPCServerEnum)
			ext.HTTPMethod.Set(span, r.Method)
			ext.HTTPUrl.Set(span, r.URL.Path)

			resourceName := r.URL.Path
			resourceName = r.Method + " " + resourceName
			span.SetTag("resource.name", resourceName)

			// pass the span through the request context and serve the request to the next middleware
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r.WithContext(traceCtx))

			// set the status code
			status := ww.Status()
			ext.HTTPStatusCode.Set(span, uint16(status))

			if status >= 500 && status < 600 {
				// mark 5xx server error
				ext.Error.Set(span, true)
				span.SetTag("error.type", fmt.Sprintf("%d: %s", status, http.StatusText(status)))
				span.LogKV(
					"event", "error",
					"message", fmt.Sprintf("%d: %s", status, http.StatusText(status)),
				)
			}
			//TODO: to be decided
			//if status >= 400 && status < 500 {
			//	// mark 4xx server error
			//	ext.Error.Set(span, true)
			//	span.SetTag("error.type", fmt.Sprintf("%d: %s", status, http.StatusText(status)))
			//	span.LogKV(
			//		"event", "error",
			//		"message", fmt.Sprintf("%d: %s", status, http.StatusText(status)),
			//	)
			//}
		})
	}
}

type contextKey struct{}

var ActiveOperationName = contextKey{}
var ActiveOperationNameString = "operationName"

func AddOperationName(ctx context.Context, operationName string) context.Context {
	return context.WithValue(ctx, ActiveOperationName, operationName)
}

func GetOperationName(ctx context.Context) string {
	val := ctx.Value(ActiveOperationName)
	if sp, ok := val.(string); ok {
		return sp
	}
	val = ctx.Value(ActiveOperationNameString)
	if sp, ok := val.(string); ok {
		return sp
	}
	return ""
}
