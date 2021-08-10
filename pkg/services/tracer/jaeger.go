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
	"errors"
	"io"
	"time"

	"github.com/opentracing/opentracing-go/ext"

	"github.com/uber/jaeger-client-go"

	jaegerPrometheus "github.com/uber/jaeger-lib/metrics/prometheus"

	"github.com/ZupIT/horusec-devkit/pkg/services/tracer/enums"

	"github.com/ZupIT/horusec-devkit/pkg/utils/env"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
)

type (
	Jaeger struct {
		Name     string
		LogError bool
		LogInfo  bool
	}
)

func NewJaeger() (*Jaeger, error) {
	j := &Jaeger{
		Name:     env.GetEnvOrDefault(enums.HorusecJaegerName, ""),
		LogError: env.GetEnvOrDefaultBool(enums.HorusecJaegerLogError, true),
		LogInfo:  env.GetEnvOrDefaultBool(enums.HorusecJaegerLogInfo, false),
	}
	if j.Name == "" {
		return nil, errors.New(enums.ErrorEmptyJaegerName)
	}
	return j, nil
}

//nolint:funlen // need to have more than 15 lines
func (j *Jaeger) Config(setPrometheus bool) (io.Closer, error) {
	defcfg := config.Configuration{
		ServiceName: j.Name,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	cfg, err := defcfg.FromEnv()

	if err != nil {
		return nil, err
	}

	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	var options []config.Option
	prom := jaegerPrometheus.New()
	if setPrometheus {
		options = append(options, config.Metrics(prom))
	}
	options = append(options, config.Injector(opentracing.HTTPHeaders, zipkinPropagator),
		config.Extractor(opentracing.HTTPHeaders, zipkinPropagator),
		config.ZipkinSharedRPCSpan(true),
		config.Logger(jaeger.StdLogger))
	return cfg.InitGlobalTracer(
		j.Name,
		options...,
	)
}

func SetSpanError(span opentracing.Span, err error) {
	span.SetTag("error.message", err.Error())
	ext.LogError(span, err)
}
