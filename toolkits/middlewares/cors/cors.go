package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/durationpb"
)

const (
	corsOptionMethod              string = "OPTIONS"
	corsAllowOriginHeader         string = "Access-Control-Allow-Origin"
	corsExposeHeadersHeader       string = "Access-Control-Expose-Headers"
	corsMaxAgeHeader              string = "Access-Control-Max-Age"
	corsAllowMethodsHeader        string = "Access-Control-Allow-Methods"
	corsAllowHeadersHeader        string = "Access-Control-Allow-Headers"
	corsAllowCredentialsHeader    string = "Access-Control-Allow-Credentials"
	corsAllowPrivateNetworkHeader string = "Access-Control-Allow-Private-Network"
	corsRequestMethodHeader       string = "Access-Control-Request-Method"
	corsRequestHeadersHeader      string = "Access-Control-Request-Headers"
	corsRequestPrivateNetwork     string = "Access-Control-Request-Private-Network"
	corsOriginHeader              string = "Origin"
	corsVaryHeader                string = "Vary"
	corsMatchAll                  string = "*"
)

var DefaultCorsConfig = CorsConfig{
	Enabled:                true,
	AllowCredentials:       false,
	AllowOrigins:           []string{"*"},
	AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type"},
	ExposeHeaders:          nil,
	MaxAge:                 durationpb.New(86400),
	AllowWildcard:          true,
	AllowBrowserExtensions: false,
	AllowWebSockets:        true,
	AllowPrivateNetwork:    false,
	AllowFiles:             true,
}

func WithCors(cfg CorsConfig) gin.HandlerFunc {
	if !cfg.Enabled {
		return func(ctx *gin.Context) {
			ctx.Next()
		}
	}

	allOrigins := true
	if len(cfg.AllowOrigins) > 0 && cfg.AllowOrigins[0] != corsMatchAll {
		allOrigins = false
	}

	return cors.New(cors.Config{
		AllowAllOrigins:        allOrigins,
		AllowOrigins:           cfg.AllowOrigins,
		AllowMethods:           cfg.AllowMethods,
		AllowPrivateNetwork:    cfg.AllowPrivateNetwork,
		AllowHeaders:           cfg.AllowHeaders,
		AllowCredentials:       cfg.AllowCredentials,
		ExposeHeaders:          cfg.ExposeHeaders,
		MaxAge:                 cfg.MaxAge.AsDuration(),
		AllowWildcard:          cfg.AllowWildcard,
		AllowBrowserExtensions: cfg.AllowBrowserExtensions,
		AllowWebSockets:        cfg.AllowWebSockets,
		AllowFiles:             cfg.AllowFiles,
	})
}
