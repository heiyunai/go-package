package ctrace

import (
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type Config struct {
	TraceUrl   string
	ServerName string
}

func InitTrace(conf *Config) (*trace.TracerProvider, error) {
	// 创建 Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.TraceUrl)))
	if err != nil {
		return nil, err
	}
	return trace.NewTracerProvider(
		// 将基于父span的采样率设置为100%
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(1.0))),
		// 始终确保在生产中批量处理
		trace.WithBatcher(exp),
		// 在资源中记录有关此应用程序的信息
		trace.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(conf.ServerName),
		)),
	), nil
}
