package metrics

import (
	"github.com/splunk/stef/go/otel/oteltef"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

// OtlpToStef defines a converter from OTLP metrics to STEF format.
type OtlpToStef interface {
	// Convert OTLP metrics to STEF format and write them to the provided writer.
	// Will not call Flush() on the writer at the end.
	Convert(src pmetric.Metrics, writer *oteltef.MetricsWriter) error
}

func NewOtlpToStef(sorted bool) OtlpToStef {
	if sorted {
		return &OtlpToStefSorted{}
	}
	return &OtlpToStefUnsorted{}
}
