package metrics

import (
	"go-metrics-demo/pkg/metrics/custom"
)

func RecordDatabaseQueryDuration(label *custom.DatabaseQueryDurationLabel, value float64) {
	custom.DatabaseQueryDuration.
		WithLabelValues(label.Database, label.Table, label.Operation, label.Sql, label.Exception, label.Function).
		Observe(value)
}

func RecordExternalCallDuration(label *custom.ExternalCallDurationLabel, value float64) {
	custom.ExternalCallDuration.
		WithLabelValues(label.Method, label.Host, label.Action, label.Version, label.Outcome, label.Status, label.Exception).
		Observe(value)
}

func RecordBusinessFailureCountInc(label *custom.BusinessFailureCountLabel) {
	custom.BusinessFailureCount.
		WithLabelValues(label.Business, label.Severity, label.ErrorType, label.Source).
		Inc()
}

func RecordBusinessFailureCountAdd(label *custom.BusinessFailureCountLabel, value float64) {
	custom.BusinessFailureCount.
		WithLabelValues(label.Business, label.Severity, label.ErrorType, label.Source).
		Add(value)
}
