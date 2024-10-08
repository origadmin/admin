package bootstrap

import (
	"github.com/go-kratos/kratos/v2/middleware"

	"origadmin/backend/internal/config"
	"origadmin/backend/toolkits/logger"
	"origadmin/backend/toolkits/metrics"
	"origadmin/backend/toolkits/traces"
)

func LoadMiddlewares(name string, conf config.Middleware) ([]middleware.Middleware, error) {
	var mids []middleware.Middleware
	if conf.Logger.Enabled {
		m, err := logger.Middleware(logger.Config{
			Name: conf.Logger.Name,
		}, nil)
		if err != nil {
			return nil, err
		}
		//tracing.Server(),
		//logging.Server(logger),
		//validate.Validator(),
		mids = append(mids, m)
	}
	if conf.Traces.Enabled {
		m, err := traces.Middleware(traces.Config{
			Name: conf.Traces.Name,
		})
		if err != nil {
			return nil, err
		}
		//tracing.Server(),
		//logging.Server(logger),
		//validate.Validator(),
		mids = append(mids, m)
	}
	if conf.Metrics.Enabled {
		m, err := metrics.Middleware(metrics.Config{
			Name: conf.Traces.Name,
		})
		if err != nil {
			return nil, err
		}
		mids = append(mids, m)
	}
	return mids, nil
}
