package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-metrics-demo/pkg/logger"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var (
	metricRequestTotal    = "request_total"
	metricRequestBody     = "request_body_size_bytes"
	metricResponseBody    = "response_body_size_bytes"
	metricRequestDuration = "request_duration_seconds"
)

var (
	defaultMetricPath         = "/metrics"
	defaultReqDurationBuckets = []float64{0.01, 0.05, 0.08, 0.1, 0.15, 0.2, 0.3}
)

// MonitorOptions defines options to create a Monitor
type MonitorOptions struct {
	MetricsPath        string
	Namespace          string
	SubSystemName      string
	ReqDurationBuckets []float64
	SkipPaths          []string
}

// Monitor defines Prometheus exporter for gin server.
type Monitor struct {
	metricsPath   string
	ns            string
	subsystemName string
	skipPaths     []string

	reqTotalCnt        *prometheus.CounterVec
	reqDuration        prometheus.HistogramVec
	reqDurationBuckets []float64
	reqBodySize        prometheus.Summary
	respBodySize       prometheus.Summary
}

// NewMonitor creates a new Monitor
func NewMonitor(o MonitorOptions) *Monitor {
	m := &Monitor{
		ns:                 o.Namespace,
		subsystemName:      o.SubSystemName,
		metricsPath:        defaultMetricPath,
		reqDurationBuckets: defaultReqDurationBuckets,
	}

	if o.MetricsPath != "" {
		m.metricsPath = o.MetricsPath
	}

	if o.ReqDurationBuckets != nil {
		m.reqDurationBuckets = o.ReqDurationBuckets
	}

	m.skipPaths = []string{m.metricsPath}
	if len(o.SkipPaths) > 0 {
		m.skipPaths = append(m.skipPaths, o.SkipPaths...)
	}

	m.init()
	return m
}

// UsedBy adds GinMonitor to Gin engine as a middleware
func (m *Monitor) UsedBy(e *gin.Engine) {
	e.Use(func(c *gin.Context) {
		for _, s := range m.skipPaths {
			if s == c.Request.URL.Path {
				c.Next()
				// skip collect metrics
				return
			}
		}
		start := time.Now()
		// deferred func collects mterics even panic happens in the following handelrs
		defer func() {
			m.collect(c, start)
		}()
		c.Next()
	})
	e.GET(m.metricsPath, func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})
}

func (m *Monitor) collect(c *gin.Context, startTime time.Time) {
	code := strconv.Itoa(c.Writer.Status())
	method := c.Request.Method
	path := c.Request.URL.Path
	m.reqTotalCnt.WithLabelValues(path, code, method).Inc()

	elapsed := time.Since(startTime).Seconds()
	m.reqDuration.WithLabelValues(path, code, method).Observe(elapsed)

	var reqSize, respSize float64
	if (c.Request.ContentLength) > 0 {
		reqSize = float64(c.Request.ContentLength)
	}
	m.reqBodySize.Observe(reqSize)

	if c.Writer.Size() > 0 {
		respSize = float64(c.Writer.Size())
	}
	m.respBodySize.Observe(respSize)
}

func (m *Monitor) init() {
	m.reqTotalCnt = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   m.ns,
		Subsystem:   m.subsystemName,
		Name:        metricRequestTotal,
		Help:        "How many HTTP requests processed, partitioned by status code and HTTP method.",
		ConstLabels: nil,
	}, []string{"path", "code", "method"})

	m.reqDuration = *prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   m.ns,
		Subsystem:   m.subsystemName,
		Name:        metricRequestDuration,
		Help:        "The HTTP request latencies in seconds.",
		ConstLabels: map[string]string{},
		Buckets:     []float64{},
	}, []string{"path", "code", "method"})

	m.reqBodySize = prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: m.ns,
		Subsystem: m.subsystemName,
		Name:      metricRequestBody,
		Help:      "The HTTP request sizes in bytes.",
	})

	m.respBodySize = prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: m.ns,
		Subsystem: m.subsystemName,
		Name:      metricResponseBody,
		Help:      "The HTTP response sizes in bytes.",
	})

	collectors := []prometheus.Collector{m.reqTotalCnt, m.reqDuration, m.reqBodySize, m.respBodySize}
	for _, x := range collectors {
		if err := prometheus.Register(x); err != nil {
			logger.Log.Warn("register collector failed", zap.Error(err))
		}
	}
}
