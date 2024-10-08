package metrics

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
)

type Config struct {
	Name string
}

func Middleware(config Config) (middleware.Middleware, error) {
	//meter := otel.Meter(config.Name)
	//var err error
	//_metricRequests, err := metrics.DefaultRequestsCounter(meter, metrics.DefaultServerRequestsCounterName)
	//if err != nil {
	//	return nil, err
	//}
	//
	//_metricSeconds, err := metrics.DefaultSecondsHistogram(meter, metrics.DefaultServerSecondsHistogramName)
	//if err != nil {
	//	return nil, err
	//}
	// TODO: add metrics middleware
	return metrics.Server(
	//metrics.WithSeconds(_metricSeconds),
	//metrics.WithRequests(_metricRequests),
	), nil
}
