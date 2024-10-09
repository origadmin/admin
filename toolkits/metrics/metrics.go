package metrics

import (
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/origadmin/toolkits/errors"
	toolmetrics "github.com/origadmin/toolkits/metrics"
	gins "github.com/origadmin/toolkits/runtime/kratos/transport/gins"
	"go.opentelemetry.io/otel"
)

const (
	SideClient = "client"
	SideServer = "server"
)

const (
	NameUptime                 = "uptime"
	NameRequestTotal           = "requests_total"
	NameRequestDurationSeconds = "request_duration_seconds"
	NameRequestsInFlight       = "requests_in_flight"
	NameRequestsSlowTotal      = "requests_slow_total"
	NameCounterSendBytes       = "counter_send_bytes"
	NameCounterRecvBytes       = "counter_recv_bytes"
	NameHistogramSeconds       = "histogram_seconds"
	NameHistogramLatency       = "histogram_latency"
	NameSummaryLatency         = "summary_latency"
	NameGaugeState             = "gauge_state"
	NameCounterException       = "counter_exception"
	NameCounterEvent           = "counter_event"
	NameCounterSiteEvent       = "counter_site_event"
)

type Config struct {
	Name    string
	Side    string
	Metrics []string
}

var (
// namespace = "gin"
//
// labels = []string{"status", "path", "method"}
//
// uptime = prometheus.NewCounterVec(
//
//	prometheus.CounterOpts{
//		Namespace: namespace,
//		Name:      "uptime",
//		Help:      "HTTP service uptime, updated every minute",
//	}, nil,
//
// )
//
// reqCount = prometheus.NewCounterVec(
//
//	prometheus.CounterOpts{
//		Namespace: namespace,
//		Name:      "http_request_count_total",
//		Help:      "Total number of HTTP requests made.",
//	}, labels,
//
// )
//
// reqDuration = prometheus.NewHistogramVec(
//
//	prometheus.HistogramOpts{
//		Namespace: namespace,
//		Name:      "http_request_duration_seconds",
//		Help:      "HTTP request latencies in seconds.",
//	}, labels,
//
// )
//
// reqSizeBytes = prometheus.NewSummaryVec(
//
//	prometheus.SummaryOpts{
//		Namespace: namespace,
//		Name:      "http_request_size_bytes",
//		Help:      "HTTP request sizes in bytes.",
//	}, labels,
//
// )
//
// respSizeBytes = prometheus.NewSummaryVec(
//
//	prometheus.SummaryOpts{
//		Namespace: namespace,
//		Name:      "http_response_size_bytes",
//		Help:      "HTTP response sizes in bytes.",
//	}, labels,
//
// )
)

func Middleware(config Config) (middleware.Middleware, error) {
	var (
		m   middleware.Middleware
		err error
	)
	switch config.Side {
	case SideServer:
		m, err = ServerMiddleware(config)
	case SideClient:
		m, err = ClientMiddleware(config)
	default:
		return nil, errors.New("unknown metrics side")
	}
	if err != nil {
		return nil, err
	}
	return m, nil
}

func ServerMiddleware(config Config) (middleware.Middleware, error) {
	meter := otel.Meter(config.Name)
	opts := make([]metrics.Option, 0, len(config.Metrics))
	if slices.Contains(config.Metrics, "requests") {
		metricRequests, err := metrics.DefaultRequestsCounter(meter, metrics.DefaultServerRequestsCounterName)
		if err != nil {
			return nil, err
		}
		opts = append(opts, metrics.WithRequests(metricRequests))
	}
	if slices.Contains(config.Metrics, "seconds") {
		metricSeconds, err := metrics.DefaultSecondsHistogram(meter, metrics.DefaultServerSecondsHistogramName)
		if err != nil {
			return nil, err
		}
		opts = append(opts, metrics.WithSeconds(metricSeconds))
	}
	return metrics.Server(opts...), nil
}

func ClientMiddleware(config Config) (middleware.Middleware, error) {
	meter := otel.Meter(config.Name)
	opts := make([]metrics.Option, 0, len(config.Metrics))
	if slices.Contains(config.Metrics, "requests") {
		metricRequests, err := metrics.DefaultRequestsCounter(meter, metrics.DefaultClientRequestsCounterName)
		if err != nil {
			return nil, err
		}
		opts = append(opts, metrics.WithRequests(metricRequests))
	}
	if slices.Contains(config.Metrics, "seconds") {
		metricSeconds, err := metrics.DefaultSecondsHistogram(meter, metrics.DefaultClientSecondsHistogramName)
		if err != nil {
			return nil, err
		}
		opts = append(opts, metrics.WithSeconds(metricSeconds))
	}
	return metrics.Client(opts...), nil
}

func WithMetrics(metrics toolmetrics.Metrics) (gins.HandlerFunc, error) {
	return func(ctx *gin.Context) {
		if !metrics.Enabled() {
			ctx.Next()
			return
		}
		start := time.Now()
		recv := int64(0)
		if ctx.Request.ContentLength > 0 {
			recv = ctx.Request.ContentLength
		}
		ctx.Next()
		code := ctx.Writer.Status()
		send := int64(ctx.Writer.Size())
		metrics.Observe(ctx.Request.Context(), toolmetrics.Report{
			Endpoint: ctx.FullPath(),
			Method:   ctx.Request.Method,
			Code:     http.StatusText(code),
			RecvSize: recv,
			SendSize: send,
			Latency:  int64(time.Since(start).Seconds()),
			Succeed:  code < 400,
		})
	}, nil
}
