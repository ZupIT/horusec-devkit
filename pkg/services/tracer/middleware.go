package tracer

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/urfave/negroni"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Tracer(tr opentracing.Tracer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if isHealthRoute(r) || isSwaggerRoute(r) {
				return
			}
			span, traceCtx := createSpanAndContext(r, tr)
			defer span.Finish()
			defer setSpanIfPanic(span)
			status := setHTTPSpans(traceCtx, w, r, span, next)
			setSpanErrorIfStatus(status, span)
		})
	}
}

func createSpanAndContext(r *http.Request, tr opentracing.Tracer) (opentracing.Span, context.Context) {
	operationName, serverSpanCtx := getContextFromHeader(r, tr)
	span, traceCtx := opentracing.StartSpanFromContextWithTracer(r.Context(),
		tr, operationName, ext.RPCServerOption(serverSpanCtx))
	return span, traceCtx
}

func setSpanIfPanic(span opentracing.Span) {
	if err := recover(); err != nil {
		ext.HTTPStatusCode.Set(span, http.StatusInternalServerError)
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
}

func setHTTPSpans(traceCtx context.Context, w http.ResponseWriter, r *http.Request,
	span opentracing.Span, next http.Handler) int {
	ext.HTTPMethod.Set(span, r.Method)
	ext.HTTPUrl.Set(span, r.URL.Path)

	resourceName := r.URL.Path
	resourceName = r.Method + " " + resourceName
	span.SetTag("resource.name", resourceName)

	ww := negroni.NewResponseWriter(w)
	next.ServeHTTP(ww, r.WithContext(traceCtx))

	status := ww.Status()
	ext.HTTPStatusCode.Set(span, uint16(status))
	return status
}

func getContextFromHeader(r *http.Request, tr opentracing.Tracer) (string, opentracing.SpanContext) {
	operationName := GetOperationName(r.Context())
	if operationName == "" {
		operationName = fmt.Sprintf("%s %s", r.Method, r.URL.Path)
	}
	serverSpanCtx, _ := tr.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	return operationName, serverSpanCtx
}

func setSpanErrorIfStatus(status int, span opentracing.Span) {
	if status >= 400 && status < 600 {
		ext.Error.Set(span, true)
		span.SetTag("error.type", fmt.Sprintf("%d: %s", status, http.StatusText(status)))
		span.LogKV(
			"event", "error",
			"message", fmt.Sprintf("%d: %s", status, http.StatusText(status)),
		)
	}
}

func isSwaggerRoute(r *http.Request) bool {
	return strings.Contains(r.URL.Path, "swagger")
}

func isHealthRoute(r *http.Request) bool {
	return strings.Contains(r.URL.Path, "health")
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
