package internal

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/splunk/stef/go/otel/oteltef"
	"github.com/splunk/stef/go/pdata/internal/otlptools"
)

type BaseSTEFToOTLP struct {
}

func (s *BaseSTEFToOTLP) convertNumberPoint(src *oteltef.Point, dst pmetric.NumberDataPoint) error {
	dst.SetStartTimestamp(pcommon.Timestamp(src.StartTimestamp()))
	dst.SetTimestamp(pcommon.Timestamp(src.Timestamp()))

	switch src.Value().Type() {
	case oteltef.PointValueTypeInt64:
		dst.SetIntValue(src.Value().Int64())
	case oteltef.PointValueTypeFloat64:
		dst.SetDoubleValue(src.Value().Float64())
	case oteltef.PointValueTypeNone:
		dst.SetFlags(pmetric.DefaultDataPointFlags.WithNoRecordedValue(true))
		return nil
	default:
		panic("unexpected type")
	}

	for i := range src.Exemplars().Len() {
		exemplar := src.Exemplars().At(i)
		if err := s.ConvertExemplar(exemplar, dst.Exemplars().AppendEmpty()); err != nil {
			return err
		}
	}
	return nil
}

func (c *BaseSTEFToOTLP) ConvertExemplar(src *oteltef.Exemplar, dst pmetric.Exemplar) error {
	dst.SetTimestamp(pcommon.Timestamp(src.Timestamp()))
	switch src.Value().Type() {
	case oteltef.ExemplarValueTypeInt64:
		dst.SetIntValue(src.Value().Int64())
	case oteltef.ExemplarValueTypeFloat64:
		dst.SetDoubleValue(src.Value().Float64())
	case oteltef.ExemplarValueTypeNone:
	default:
		panic("unknown exemplar value type")
	}
	dst.SetTraceID(pcommon.TraceID([]byte(src.TraceID())))
	dst.SetSpanID(pcommon.SpanID([]byte(src.SpanID())))

	otlpAttrs := pcommon.NewMap()
	err := otlptools.TefToOtlpMap(src.FilteredAttributes(), otlpAttrs)
	if err != nil {
		return err
	}
	otlpAttrs.MoveTo(dst.FilteredAttributes())
	return nil
}

func (s *BaseSTEFToOTLP) AppendOTLPPoint(
	srcMetric *oteltef.Metric, srcAttrs *oteltef.Attributes, srcPoint *oteltef.Point, dstMetric pmetric.Metric,
) error {
	otlpAttrs := pcommon.NewMap()
	err := otlptools.TefToOtlpMap(srcAttrs, otlpAttrs)
	if err != nil {
		return err
	}

	switch dstMetric.Type() {
	case pmetric.MetricTypeGauge:
		point := dstMetric.Gauge().DataPoints().AppendEmpty()
		if err := s.convertNumberPoint(srcPoint, point); err != nil {
			return err
		}
		otlpAttrs.MoveTo(point.Attributes())
	case pmetric.MetricTypeSum:
		point := dstMetric.Sum().DataPoints().AppendEmpty()
		if err := s.convertNumberPoint(srcPoint, point); err != nil {
			return err
		}
		otlpAttrs.MoveTo(point.Attributes())
		dstMetric.Sum().SetIsMonotonic(srcMetric.Monotonic())
		dstMetric.Sum().SetAggregationTemporality(aggregationTemporalityToOtlp(srcMetric.AggregationTemporality()))
	case pmetric.MetricTypeHistogram:
		point := dstMetric.Histogram().DataPoints().AppendEmpty()
		if err := s.convertHistogramPoint(srcMetric, srcPoint, point); err != nil {
			return err
		}
		otlpAttrs.MoveTo(point.Attributes())
		dstMetric.Histogram().SetAggregationTemporality(aggregationTemporalityToOtlp(srcMetric.AggregationTemporality()))
	case pmetric.MetricTypeExponentialHistogram:
		point := dstMetric.ExponentialHistogram().DataPoints().AppendEmpty()
		if err := s.convertExpHistogramPoint(srcPoint, point); err != nil {
			return err
		}
		otlpAttrs.MoveTo(point.Attributes())
		dstMetric.ExponentialHistogram().SetAggregationTemporality(aggregationTemporalityToOtlp(srcMetric.AggregationTemporality()))
	case pmetric.MetricTypeSummary:
		point := dstMetric.Summary().DataPoints().AppendEmpty()
		if err := s.convertSumaryPoint(srcPoint, point); err != nil {
			return err
		}
		otlpAttrs.MoveTo(point.Attributes())
	default:
		panic("not implemented")
	}

	return nil
}

func aggregationTemporalityToOtlp(temp oteltef.AggregationTemporality) pmetric.AggregationTemporality {
	switch temp {
	case oteltef.AggregationTemporalityDelta:
		return pmetric.AggregationTemporalityDelta
	case oteltef.AggregationTemporalityCumulative:
		return pmetric.AggregationTemporalityCumulative
	case oteltef.AggregationTemporalityUnspecified:
		return pmetric.AggregationTemporalityUnspecified
	default:
		panic("unexpected aggregation temporality")
	}
}

func (c *BaseSTEFToOTLP) convertHistogramPoint(
	srcMetric *oteltef.Metric, srcPoint *oteltef.Point, dstPoint pmetric.HistogramDataPoint,
) error {
	dstPoint.SetStartTimestamp(pcommon.Timestamp(srcPoint.StartTimestamp()))
	dstPoint.SetTimestamp(pcommon.Timestamp(srcPoint.Timestamp()))

	if srcPoint.Value().Type() == oteltef.PointValueTypeNone {
		dstPoint.SetFlags(pmetric.DefaultDataPointFlags.WithNoRecordedValue(true))
		return nil
	}

	dstPoint.SetCount(uint64(srcPoint.Value().Histogram().Count()))

	dstPoint.BucketCounts().EnsureCapacity(srcPoint.Value().Histogram().BucketCounts().Len())
	for i := 0; i < srcPoint.Value().Histogram().BucketCounts().Len(); i++ {
		dstPoint.BucketCounts().Append(srcPoint.Value().Histogram().BucketCounts().At(i))
	}

	dstPoint.ExplicitBounds().EnsureCapacity(srcMetric.HistogramBounds().Len())
	for i := range srcMetric.HistogramBounds().Len() {
		dstPoint.ExplicitBounds().Append(srcMetric.HistogramBounds().At(i))
	}

	if srcPoint.Value().Histogram().HasSum() {
		dstPoint.SetSum(srcPoint.Value().Histogram().Sum())
	}

	if srcPoint.Value().Histogram().HasMin() {
		dstPoint.SetMin(srcPoint.Value().Histogram().Min())
	}

	if srcPoint.Value().Histogram().HasMax() {
		dstPoint.SetMax(srcPoint.Value().Histogram().Max())
	}

	for i := range srcPoint.Exemplars().Len() {
		exemplar := srcPoint.Exemplars().At(i)
		if err := c.ConvertExemplar(exemplar, dstPoint.Exemplars().AppendEmpty()); err != nil {
			return err
		}
	}
	return nil
}

func (c *BaseSTEFToOTLP) convertExpHistogramPoint(
	srcPoint *oteltef.Point, dstPoint pmetric.ExponentialHistogramDataPoint,
) error {
	dstPoint.SetStartTimestamp(pcommon.Timestamp(srcPoint.StartTimestamp()))
	dstPoint.SetTimestamp(pcommon.Timestamp(srcPoint.Timestamp()))
	dstPoint.SetCount(srcPoint.Value().ExpHistogram().Count())

	if srcPoint.Value().Type() == oteltef.PointValueTypeNone {
		dstPoint.SetFlags(pmetric.DefaultDataPointFlags.WithNoRecordedValue(true))
		return nil
	}

	expBucketsFromStef(dstPoint.Positive(), srcPoint.Value().ExpHistogram().PositiveBuckets())
	expBucketsFromStef(dstPoint.Negative(), srcPoint.Value().ExpHistogram().NegativeBuckets())

	if srcPoint.Value().ExpHistogram().HasSum() {
		dstPoint.SetSum(srcPoint.Value().ExpHistogram().Sum())
	}

	if srcPoint.Value().ExpHistogram().HasMin() {
		dstPoint.SetMin(srcPoint.Value().ExpHistogram().Min())
	}

	if srcPoint.Value().ExpHistogram().HasMax() {
		dstPoint.SetMax(srcPoint.Value().ExpHistogram().Max())
	}

	dstPoint.SetScale(int32(srcPoint.Value().ExpHistogram().Scale()))
	dstPoint.SetZeroCount(srcPoint.Value().ExpHistogram().ZeroCount())
	dstPoint.SetZeroThreshold(srcPoint.Value().ExpHistogram().ZeroThreshold())

	for i := range srcPoint.Exemplars().Len() {
		exemplar := srcPoint.Exemplars().At(i)
		if err := c.ConvertExemplar(exemplar, dstPoint.Exemplars().AppendEmpty()); err != nil {
			return err
		}
	}
	return nil
}

func expBucketsFromStef(
	dst pmetric.ExponentialHistogramDataPointBuckets, src *oteltef.ExpHistogramBuckets,
) {
	dst.SetOffset(int32(src.Offset()))
	dst.BucketCounts().EnsureCapacity(src.BucketCounts().Len())
	for i := 0; i < src.BucketCounts().Len(); i++ {
		dst.BucketCounts().Append(src.BucketCounts().At(i))
	}
}

func (c *BaseSTEFToOTLP) convertSumaryPoint(
	srcPoint *oteltef.Point, dstPoint pmetric.SummaryDataPoint,
) error {
	dstPoint.SetStartTimestamp(pcommon.Timestamp(srcPoint.StartTimestamp()))
	dstPoint.SetTimestamp(pcommon.Timestamp(srcPoint.Timestamp()))

	if srcPoint.Value().Type() == oteltef.PointValueTypeNone {
		dstPoint.SetFlags(pmetric.DefaultDataPointFlags.WithNoRecordedValue(true))
		return nil
	}

	dstPoint.SetCount(srcPoint.Value().Summary().Count())
	dstPoint.SetSum(srcPoint.Value().Summary().Sum())

	quantilesFromStef(dstPoint.QuantileValues(), srcPoint.Value().Summary().QuantileValues())

	return nil
}

func quantilesFromStef(dst pmetric.SummaryDataPointValueAtQuantileSlice, src *oteltef.QuantileValueArray) {
	dst.EnsureCapacity(src.Len())
	for i := 0; i < src.Len(); i++ {
		dstQ := dst.AppendEmpty()
		dstQ.SetQuantile(src.At(i).Quantile())
		dstQ.SetValue(src.At(i).Value())
	}
}
