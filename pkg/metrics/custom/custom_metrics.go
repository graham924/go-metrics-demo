package custom

import "github.com/prometheus/client_golang/prometheus"

var (
	metricDatabaseQueryDuration = "database_query_duration"
	metricExternalCallDuration  = "external_call_duration"
	metricBusinessFailureCount  = "business_failure_count"
)

var (
	DatabaseQueryDuration *prometheus.HistogramVec
	ExternalCallDuration  *prometheus.HistogramVec
	BusinessFailureCount  *prometheus.CounterVec
)

func init() {
	DatabaseQueryDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        metricDatabaseQueryDuration,
		Help:        "The HTTP request latencies in seconds.",
		ConstLabels: map[string]string{},
		Buckets:     []float64{},
	}, []string{"database", "table", "operation", "sql", "exception", "function"})

	ExternalCallDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        metricExternalCallDuration,
		Help:        "The HTTP request latencies in seconds.",
		ConstLabels: map[string]string{},
		Buckets:     []float64{},
	}, []string{"method", "host", "action", "version", "outcome", "status", "exception"})

	BusinessFailureCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: metricBusinessFailureCount,
		Help: "The HTTP request latencies in seconds.",
	}, []string{"business", "severity", "errorType", "source"})

	prometheus.MustRegister(DatabaseQueryDuration)
	prometheus.MustRegister(ExternalCallDuration)
	prometheus.MustRegister(BusinessFailureCount)
}

type DatabaseQueryDurationLabel struct {
	Database  string
	Table     string
	Operation string
	Sql       string
	Exception string
	Function  string
}

type ExternalCallDurationLabel struct {
	Method    string
	Host      string
	Action    string
	Version   string
	Outcome   string
	Status    string
	Exception string
}

type BusinessFailureCountLabel struct {
	Business  string
	Severity  string
	ErrorType string
	Source    string
}
