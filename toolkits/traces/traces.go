package traces

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
)

type Config struct {
	Name string
}

func Middleware(config Config) (middleware.Middleware, error) {

	//// WithPropagator with tracer propagator.
	//func WithPropagator(propagator propagation.TextMapPropagator) Option {
	//	return func(opts *options) {
	//	opts.propagator = propagator
	//}
	//}
	//
	//// WithTracerProvider with tracer provider.
	//// By default, it uses the global provider that is set by otel.SetTracerProvider(provider).
	//func WithTracerProvider(provider trace.TracerProvider) Option {
	//	return func(opts *options) {
	//	opts.tracerProvider = provider
	//}
	//}
	//
	//// WithTracerName with tracer name
	//func WithTracerName(tracerName string) Option {
	//	return func(opts *options) {
	//	opts.tracerName = tracerName
	//}
	//}
	return tracing.Server(
	//metrics.WithSeconds(_metricSeconds),
	//metrics.WithRequests(_metricRequests),
	), nil
}
