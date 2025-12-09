package metrics

import (
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/splunk/stef/go/otel/oteltef"
	"github.com/splunk/stef/go/pdata/internal/otlptools"
	"github.com/splunk/stef/go/pdata/metrics/internal"
)

// OtlpToStefUnsorted converts OTLP metrics to STEF format, without sorting.
type OtlpToStefUnsorted struct {
	base internal.BaseOtlpToStef
}

var _ OtlpToStef = (*OtlpToStefUnsorted)(nil)

// Convert OTLP metrics to STEF format and writes them to the provided writer.
// Will not call Flush() on the writer at the end.
func (d *OtlpToStefUnsorted) Convert(src pmetric.Metrics, writer *oteltef.MetricsWriter) error {
	otlp2stef := &otlptools.Otlp2Stef{}
	for i := 0; i < src.ResourceMetrics().Len(); i++ {
		rmm := src.ResourceMetrics().At(i)
		otlp2stef.ResourceUnsorted(writer.Record.Resource(), rmm.Resource(), rmm.SchemaUrl())

		for j := 0; j < rmm.ScopeMetrics().Len(); j++ {
			smm := rmm.ScopeMetrics().At(j)
			otlp2stef.ScopeUnsorted(writer.Record.Scope(), smm.Scope(), smm.SchemaUrl())
			for k := 0; k < smm.Metrics().Len(); k++ {
				m := smm.Metrics().At(k)
				metric2metric(m, writer.Record.Metric(), otlp2stef)
				var err error
				switch m.Type() {
				case pmetric.MetricTypeGauge:
					err = d.writeNumeric(writer, m.Gauge().DataPoints())
				case pmetric.MetricTypeSum:
					writer.Record.Metric().SetAggregationTemporality(internal.AggregationTemporalityToStef(m.Sum().AggregationTemporality()))
					writer.Record.Metric().SetMonotonic(m.Sum().IsMonotonic())
					err = d.writeNumeric(writer, m.Sum().DataPoints())
				case pmetric.MetricTypeHistogram:
					writer.Record.Metric().SetAggregationTemporality(internal.AggregationTemporalityToStef(m.Histogram().AggregationTemporality()))
					err = d.writeHistogram(writer, m.Histogram().DataPoints())
				case pmetric.MetricTypeExponentialHistogram:
					writer.Record.Metric().SetAggregationTemporality(internal.AggregationTemporalityToStef(m.ExponentialHistogram().AggregationTemporality()))
					err = d.writeExpHistogram(writer, m.ExponentialHistogram().DataPoints())
				case pmetric.MetricTypeSummary:
					err = d.writeSummary(writer, m.Summary().DataPoints())
				default:
					panic("Unsupported metric type")
				}
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func metricType(typ pmetric.MetricType) oteltef.MetricType {
	switch typ {
	case pmetric.MetricTypeGauge:
		return oteltef.MetricTypeGauge
	case pmetric.MetricTypeSum:
		return oteltef.MetricTypeSum
	case pmetric.MetricTypeHistogram:
		return oteltef.MetricTypeHistogram
	case pmetric.MetricTypeExponentialHistogram:
		return oteltef.MetricTypeExpHistogram
	case pmetric.MetricTypeSummary:
		return oteltef.MetricTypeSummary
	default:
		panic("Unsupported metric value")
	}
	return 0
}

func metric2metric(
	src pmetric.Metric, // histogramBounds []float64,
	dst *oteltef.Metric,
	otlp2stef *otlptools.Otlp2Stef,
) {
	otlp2stef.MapUnsorted(src.Metadata(), dst.Metadata())
	dst.SetName(src.Name())
	dst.SetDescription(src.Description())
	dst.SetUnit(src.Unit())
	dst.SetType(metricType(src.Type()))
}

func (d *OtlpToStefUnsorted) writeNumeric(writer *oteltef.MetricsWriter, src pmetric.NumberDataPointSlice) error {
	for i := 0; i < src.Len(); i++ {
		srcPoint := src.At(i)
		dstPoint := writer.Record.Point()

		d.base.ConvertNumDatapoint(dstPoint, srcPoint)
		d.base.Otlp2tef.MapUnsorted(srcPoint.Attributes(), writer.Record.Attributes())
		d.base.ConvertExemplars(dstPoint.Exemplars(), srcPoint.Exemplars())

		err := writer.Write()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *OtlpToStefUnsorted) writeHistogram(writer *oteltef.MetricsWriter, src pmetric.HistogramDataPointSlice) error {
	for i := 0; i < src.Len(); i++ {
		src := src.At(i)
		dst := writer.Record.Point()

		err := d.base.ConvertHistogram(dst, src)
		if err != nil {
			return err
		}

		d.base.Otlp2tef.MapUnsorted(src.Attributes(), writer.Record.Attributes())
		writer.Record.Metric().HistogramBounds().CopyFromSlice(src.ExplicitBounds().AsRaw())
		d.base.ConvertExemplars(dst.Exemplars(), src.Exemplars())

		err = writer.Write()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *OtlpToStefUnsorted) writeExpHistogram(
	writer *oteltef.MetricsWriter, src pmetric.ExponentialHistogramDataPointSlice,
) error {
	for i := 0; i < src.Len(); i++ {
		src := src.At(i)
		dst := writer.Record.Point()

		err := d.base.ConvertExpHistogram(dst, src)
		if err != nil {
			return err
		}

		d.base.Otlp2tef.MapUnsorted(src.Attributes(), writer.Record.Attributes())
		d.base.ConvertExemplars(dst.Exemplars(), src.Exemplars())

		err = writer.Write()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *OtlpToStefUnsorted) writeSummary(writer *oteltef.MetricsWriter, src pmetric.SummaryDataPointSlice) error {
	for i := 0; i < src.Len(); i++ {
		srcPoint := src.At(i)
		dstPoint := writer.Record.Point()

		err := d.base.ConvertSummary(dstPoint, srcPoint)
		if err != nil {
			return err
		}

		d.base.Otlp2tef.MapUnsorted(srcPoint.Attributes(), writer.Record.Attributes())

		err = writer.Write()
		if err != nil {
			return err
		}
	}
	return nil
}
