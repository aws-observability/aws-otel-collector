package metrics

import (
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/splunk/stef/go/otel/oteltef"
	"github.com/splunk/stef/go/pdata/metrics/sortedbymetric"
)

// OtlpToStefSorted sorts and converts OTLP metrics to STEF format.
type OtlpToStefSorted struct{}

var _ OtlpToStef = (*OtlpToStefSorted)(nil)

// Convert OTLP metrics to STEF format and writes them to the provided writer.
// Metrics are sorted before conversion, typically improving STEF compression ratio.
// Input src metrics are not modified, sorting is done on an internal copy.
// Will not call Flush() on the writer at the end.
func (d *OtlpToStefSorted) Convert(src pmetric.Metrics, writer *oteltef.MetricsWriter) error {
	tree, err := sortedbymetric.OtlpToSortedTree(src)
	if err != nil {
		return err
	}
	return tree.ToStef(writer)
}
