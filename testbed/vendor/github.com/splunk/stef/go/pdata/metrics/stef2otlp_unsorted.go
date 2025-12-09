package metrics

import (
	"errors"
	"io"

	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/splunk/stef/go/otel/oteltef"
	"github.com/splunk/stef/go/pdata/internal/otlptools"
	"github.com/splunk/stef/go/pdata/metrics/internal"
	"github.com/splunk/stef/go/pkg"
)

// StefToOtlpUnsorted reads and converts STEF records to OTLP metrics.
type StefToOtlpUnsorted struct {
	base internal.BaseSTEFToOTLP
}

var _ StefToOtlp = (*StefToOtlpUnsorted)(nil)

// Convert by reading from STEF reader using specified ReadOptions and return
// converted metrics.
//
// If untilEOF is true, will read until EOF is reached, otherwise will
// read at least one record and then continues reading records from the
// reader until the end of the current frame.
//
// If untilEOF is false only the very first record reading may block on I/O.
// Once the first record is read, the subsequent reads will not block on I/O.
// This is useful for converting records from a Reader that uses an underlying
// GrpcReader that can block on reading from network. This ensures that we won't
// read a record and keep it in memory while blocked on further reads.
func (c *StefToOtlpUnsorted) Convert(reader *oteltef.MetricsReader, untilEOF bool) (pmetric.Metrics, error) {
	var resourceMetrics pmetric.ResourceMetrics
	var scopeMetrics pmetric.ScopeMetrics
	var metric pmetric.Metric

	metrics := pmetric.NewMetrics()

	// Read the first record. This call may block on I/O.
	err := reader.Read(pkg.ReadOptions{})
	if err != nil {
		return metrics, err
	}
	record := &reader.Record

	modified := true
	for {
		if modified || record.IsResourceModified() {
			modified = true
			resourceMetrics = metrics.ResourceMetrics().AppendEmpty()
			if err := otlptools.ResourceToOtlp(record.Resource(), resourceMetrics); err != nil {
				return metrics, err
			}
		}

		if modified || record.IsScopeModified() {
			modified = true
			scopeMetrics = resourceMetrics.ScopeMetrics().AppendEmpty()
			if err := otlptools.ScopeToOtlp(record.Scope(), scopeMetrics); err != nil {
				return metrics, err
			}
		}

		if modified || record.IsMetricModified() {
			metric = scopeMetrics.Metrics().AppendEmpty()
			if err := otlptools.MetricToOtlp(record.Metric(), metric); err != nil {
				return metrics, err
			}
		}

		err = c.base.AppendOTLPPoint(record.Metric(), record.Attributes(), record.Point(), metric)
		if err != nil {
			return metrics, err
		}

		modified = false

		if untilEOF {
			err = reader.Read(pkg.ReadOptions{})
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return metrics, err
			}
		} else {
			// Read more records. This will not block on I/O.
			err = reader.Read(pkg.ReadOptions{TillEndOfFrame: true})
			if err != nil {
				if errors.Is(err, pkg.ErrEndOfFrame) {
					break
				}
				return metrics, err
			}
		}
	}

	return metrics, nil
}
