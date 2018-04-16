// package instrumentation provides utilities for instrumenting Go code.
package instrumentation

import (
	"go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"
)


func NewJaegerExporter(opt jaeger.Options, sampler trace.Sampler) (*jaeger.Exporter, error) {
	exporter, err := jaeger.NewExporter(opt)
	if err != nil {
		return nil, err
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{ DefaultSampler: sampler })

	return exporter, nil
}

func LocalJaegerExporter(service string) (*jaeger.Exporter, error) {
	localOpt := jaeger.Options{
		Endpoint: 	   "http://localhost:14268",
		AgentEndpoint: "localhost:6831",
		ServiceName:   service,
	}
	exporter, err := NewJaegerExporter(localOpt, trace.AlwaysSample())
	if err != nil {
		return nil, err
	}

	return exporter, nil
}
