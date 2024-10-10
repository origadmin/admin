package config

type Metrics struct {
	Enabled bool
	Name    string
}

type Traces struct {
	Enabled bool
	Name    string
}

type Logger struct {
	Enabled bool
	Name    string
}

type Cors struct {
	Enabled         bool
	AllowAllOrigins bool
	// AllowOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// Default value is []
	AllowOrigins []string
	// AllowMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET, POST, PUT, PATCH, DELETE, HEAD, and OPTIONS)
	AllowMethods []string
	// AllowHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
	AllowHeaders []string
	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool
	// ExposeHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposeHeaders []string
	// MaxAge indicates how long (with second-precision) the results of a preflight request
	// can be cached
	MaxAge int
	// Allows to add origins like http://some-domain/*, https://api.* or http://some.*.subdomain.com
	AllowWildcard bool
	// Allows usage of popular browser extensions schemas
	AllowBrowserExtensions bool
	// Allows usage of WebSocket protocol
	AllowWebSockets bool
	// Allows usage of file:// schema (dangerous!) use it only when you 100% sure it's needed
	AllowFiles bool
}

type Middleware struct {
	Cors    Cors
	Metrics Metrics
	Traces  Traces
	Logger  Logger
}
