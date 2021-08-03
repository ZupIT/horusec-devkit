package tracer

import (
	"errors"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
	jaegerPrometheus "github.com/uber/jaeger-lib/metrics/prometheus"
	"io"
	"time"
)

type (
	Jaeger struct {
		Name     string
		LogError bool
		LogInfo  bool
	}
)

func (j *Jaeger) Config() (io.Closer, error) {

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

	jLogger := &stdLogger{
		LogError: j.LogError,
		LogInfo:  j.LogInfo,
	}
	jMetricsFactory := jaegerPrometheus.New()
	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()

	closer, err := cfg.InitGlobalTracer(
		j.Name,
		config.Injector(opentracing.HTTPHeaders, zipkinPropagator),
		config.Extractor(opentracing.HTTPHeaders, zipkinPropagator),
		config.ZipkinSharedRPCSpan(true),
		config.Logger(jLogger),
		config.Metrics(jMetricsFactory),
	)
	if err != nil {
		return nil, err
	}
	return closer, nil
}

type stdLogger struct {
	LogError bool
	LogInfo  bool
}

func (l *stdLogger) Error(msg string) {
	if l.LogError {
		logger.LogError(msg, errors.New(msg))
	}
}

func (l *stdLogger) Infof(msg string, args ...interface{}) {
	if l.LogInfo {
		logger.LogInfo(msg, args)
	}
}
