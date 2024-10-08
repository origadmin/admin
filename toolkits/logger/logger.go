package logger

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
)

type Config struct {
	Name string
}

func Middleware(config Config, logger log.Logger) (middleware.Middleware, error) {
	if logger == nil {
		// todo: init logger from config
		logger = log.NewStdLogger(os.Stdout)
	}
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
	return logging.Server(logger), nil
}

func NewLogger(config Config) log.Logger {
	return log.NewStdLogger(os.Stdout)
}
