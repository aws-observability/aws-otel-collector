package metrics

import (
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/splunk/stef/go/otel/oteltef"
)

// StefToOtlp defines a converter from STEF format to OTLP metrics.
type StefToOtlp interface {
	// Convert by reading from STEF reader using specified ReadOptions and return
	// converted metrics.
	//
	// If untilEOF is true, will read until EOF is reached, otherwise will
	// read at least one record and then continues reading records from the
	// reader until the end of the current frame.
	Convert(reader *oteltef.MetricsReader, untilEOF bool) (pmetric.Metrics, error)
}
