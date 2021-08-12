// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tracer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/go-chi/chi/middleware"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Tracer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isHealthRoute(r) || isSwaggerRoute(r) {
			next.ServeHTTP(w, r)
			return
		}
		span, traceCtx := createSpanAndContext(r, opentracing.GlobalTracer())
		defer span.Finish()
		defer setSpanIfPanic(span, w)
		status := setHTTPSpans(traceCtx, w, r, span, next)
		setSpanErrorIfStatus(status, span)
	})
}

func createSpanAndContext(r *http.Request, tr opentracing.Tracer) (opentracing.Span, context.Context) {
	operationName, serverSpanCtx := getContextFromHeader(r, tr)
	span, traceCtx := opentracing.StartSpanFromContext(r.Context(),
		operationName, ext.RPCServerOption(serverSpanCtx))
	return span, traceCtx
}

func setSpanIfPanic(span opentracing.Span, w http.ResponseWriter) {
	if err := recover(); err != nil {
		ext.HTTPStatusCode.Set(span, http.StatusInternalServerError)
		SetSpanError(span, errors.New(fmt.Sprint(err)))
		span.SetTag("error.type", "panic")
		span.LogKV(
			"event", "error",
			"error.kind", "panic",
			"message", err,
			"stack", string(debug.Stack()),
		)
		span.Finish()
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func setHTTPSpans(traceCtx context.Context, w http.ResponseWriter, r *http.Request,
	span opentracing.Span, next http.Handler) int {
	ext.HTTPMethod.Set(span, r.Method)
	ext.HTTPUrl.Set(span, r.URL.Path)

	resourceName := r.URL.Path
	resourceName = r.Method + " " + resourceName
	span.SetTag("resource.name", resourceName)

	ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
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
	if status >= 500 && status < 600 {
		SetSpanError(span, errors.New("status code "+strconv.Itoa(status)))
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

type ContextValue string

const ActiveOperationName ContextValue = "operationName"

func AddOperationName(ctx context.Context, operationName string) context.Context {
	return context.WithValue(ctx, ActiveOperationName, operationName)
}

func GetOperationName(ctx context.Context) string {
	val := ctx.Value(ActiveOperationName)
	if sp, ok := val.(string); ok {
		return sp
	}
	return ""
}
