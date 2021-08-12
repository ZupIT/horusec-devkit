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

	"github.com/opentracing/opentracing-go/ext"

	"github.com/uber/jaeger-client-go"

	jaegerPrometheus "github.com/uber/jaeger-lib/metrics/prometheus"

	"github.com/ZupIT/horusec-devkit/pkg/services/tracer/enums"

	"github.com/ZupIT/horusec-devkit/pkg/utils/env"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

type (
	Jaeger struct {
		Name string
	}
)

func NewJaeger(serviceName string) (*Jaeger, error) {
	j := &Jaeger{
		Name: env.GetEnvOrDefault(enums.JaegerServiceName, serviceName),
	}
	if j.Name == "" {
		return nil, errors.New(enums.ErrorEmptyJaegerServiceName)
	}
	return j, nil
}

func (j *Jaeger) Config(setPrometheus bool) (io.Closer, error) {
	defcfg := j.getDefaultConfig()
	cfg, err := defcfg.FromEnv()

	if err != nil {
		return nil, err
	}
	var options []config.Option
	if setPrometheus {
		options = append(options, config.Metrics(jaegerPrometheus.New()))
	}
	options = append(options,
		config.ZipkinSharedRPCSpan(true),
		config.Logger(jaeger.StdLogger))
	return cfg.InitGlobalTracer(j.Name, options...)
}

func (j *Jaeger) getDefaultConfig() config.Configuration {
	defcfg := config.Configuration{
		ServiceName: j.Name,
	}
	return defcfg
}

func SetSpanError(span opentracing.Span, err error) {
	span.SetTag("error.message", err.Error())
	ext.LogError(span, err)
}
func SpanError(span opentracing.Span, err error) error {
	if err != nil {
		span.SetTag("error.message", err.Error())
		ext.LogError(span, err)
	}
	return err
}
