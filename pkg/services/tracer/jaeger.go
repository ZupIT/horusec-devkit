package tracer

import (
	"errors"
	"io"
	"time"

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
