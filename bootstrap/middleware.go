package bootstrap

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/validate"

	"origadmin/backend/internal/config"
	"origadmin/backend/toolkits/logger"
	"origadmin/backend/toolkits/metrics"
	"origadmin/backend/toolkits/traces"
)

func LoadMiddlewares(name string, conf config.Middleware) ([]middleware.Middleware, error) {
	var middlewares []middleware.Middleware
	middlewares = append(middlewares, validate.Validator())
	if conf.Logger.Enabled {
		m, err := logger.Middleware(logger.Config{
			Name: conf.Logger.Name,
		}, nil)
		if err != nil {
			return nil, err
		}
		middlewares = append(middlewares, m)
	}
	if conf.Traces.Enabled {
		m, err := traces.Middleware(traces.Config{
			Name: conf.Traces.Name,
		})
		if err != nil {
			return nil, err
		}
		middlewares = append(middlewares, m)
	}

	if conf.Metrics.Enabled {
		m, err := metrics.Middleware(metrics.Config{
			Name: conf.Traces.Name,
			Side: metrics.SideServer,
		})
		if err != nil {
			return nil, err
		}
		middlewares = append(middlewares, m)
	}
	return middlewares, nil
}
