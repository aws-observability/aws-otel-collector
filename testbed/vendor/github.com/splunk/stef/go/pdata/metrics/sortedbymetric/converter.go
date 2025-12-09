package sortedbymetric

import (
	"fmt"

	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/splunk/stef/go/otel/oteltef"

	"github.com/splunk/stef/go/pdata/metrics/internal"
)

type converter struct {
	internal.BaseOtlpToStef
}

func OtlpToSortedTree(data pmetric.Metrics) (*SortedTree, error) {
	rms := data.ResourceMetrics()
	sm := NewSortedMetrics()
	c := converter{}

	for i := 0; i < rms.Len(); i++ {
		rm := rms.At(i)
		resource := oteltef.NewResource()
		c.Otlp2tef.ResourceSorted(resource, rm.Resource(), rm.SchemaUrl())
		resource.Freeze()

		for j := 0; j < rm.ScopeMetrics().Len(); j++ {
			sms := rm.ScopeMetrics().At(j)
			scope := oteltef.NewScope()
			c.Otlp2tef.ScopeSorted(scope, sms.Scope(), sms.SchemaUrl())
			scope.Freeze()

			for k := 0; k < sms.Metrics().Len(); k++ {
				metric := sms.Metrics().At(k)
				switch metric.Type() {
				case pmetric.MetricTypeGauge:
					c.covertNumberDataPoints(
						sm, resource, scope, metric, metric.Gauge().DataPoints(), internal.MetricFlags{},
					)
				case pmetric.MetricTypeSum:
					c.covertNumberDataPoints(
						sm, resource, scope, metric, metric.Sum().DataPoints(),
						calcMetricFlags(metric.Sum().IsMonotonic(), metric.Sum().AggregationTemporality()),
					)
				case pmetric.MetricTypeHistogram:
					err := c.covertHistogramDataPoints(sm, resource, scope, metric, metric.Histogram())
					if err != nil {
						return nil, err
					}
				case pmetric.MetricTypeExponentialHistogram:
					err := c.covertExponentialHistogramDataPoints(
						sm, resource, scope, metric, metric.ExponentialHistogram(),
					)
					if err != nil {
						return nil, err
					}
				case pmetric.MetricTypeSummary:
					err := c.covertSummaryDataPoints(sm, resource, scope, metric, metric.Summary())
					if err != nil {
						return nil, err
					}
				default:
					panic(fmt.Sprintf("Unsupported metric type: %v (metric name=%s)", metric.Type(), metric.Name()))
				}
			}
		}
	}

	sm.SortValues()

	return sm, nil
}

func calcMetricFlags(monotonic bool, temporality pmetric.AggregationTemporality) internal.MetricFlags {
	return internal.MetricFlags{
		Monotonic:   monotonic,
		Temporality: internal.AggregationTemporalityToStef(temporality),
	}
}

func (c *converter) covertNumberDataPoints(
	sm *SortedTree,
	rm *oteltef.Resource,
	sms *oteltef.Scope,
	metric pmetric.Metric,
	srcPoints pmetric.NumberDataPointSlice,
	flags internal.MetricFlags,
) {
	var metricType oteltef.MetricType
	var byMetric *ByMetric
	var byScope *ByScope

	dstPointSlice := make([]oteltef.Point, srcPoints.Len())

	for l := 0; l < srcPoints.Len(); l++ {
		srcPoint := srcPoints.At(l)

		if srcPoint.ValueType() == pmetric.NumberDataPointValueTypeEmpty {
			continue
		}

		mt := calcNumericMetricType(metric)
		if mt != metricType || byMetric == nil {
			metricType = mt
			byMetric = sm.ByMetric(metric, metricType, flags, nil)
			byResource := byMetric.ByResource(rm)
			byScope = byResource.ByScope(sms)
		}

		c.Otlp2tef.MapSorted(srcPoint.Attributes(), &c.TempAttrs)
		dstPoints := byScope.ByAttrs(&c.TempAttrs)

		dstPoint := &dstPointSlice[l]
		dstPoint.Init()

		*dstPoints = append(*dstPoints, dstPoint)
		dstPoint.SetTimestamp(uint64(srcPoint.Timestamp()))
		dstPoint.SetStartTimestamp(uint64(srcPoint.StartTimestamp()))
		c.ConvertExemplars(dstPoint.Exemplars(), srcPoint.Exemplars())

		c.ConvertNumDatapoint(dstPoint, srcPoint)
	}
}

func calcNumericMetricType(metric pmetric.Metric) oteltef.MetricType {
	switch metric.Type() {
	case pmetric.MetricTypeGauge:
		return oteltef.MetricTypeGauge
	case pmetric.MetricTypeSum:
		return oteltef.MetricTypeSum
	default:
		panic("Unsupported metric type")
	}
	return 0
}

func (c *converter) covertHistogramDataPoints(
	sm *SortedTree,
	rm *oteltef.Resource,
	sms *oteltef.Scope,
	metric pmetric.Metric,
	hist pmetric.Histogram,
) error {
	var byMetric *ByMetric
	var byScope *ByScope
	flags := calcMetricFlags(false, hist.AggregationTemporality())
	srcPoints := hist.DataPoints()

	for l := 0; l < srcPoints.Len(); l++ {
		srcPoint := srcPoints.At(l)

		byMetric = sm.ByMetric(metric, oteltef.MetricTypeHistogram, flags, srcPoint.ExplicitBounds().AsRaw())
		byResource := byMetric.ByResource(rm)
		byScope = byResource.ByScope(sms)
		c.Otlp2tef.MapSorted(srcPoint.Attributes(), &c.TempAttrs)

		c.Otlp2tef.MapSorted(srcPoint.Attributes(), &c.TempAttrs)
		dstPoints := byScope.ByAttrs(&c.TempAttrs)
		dstPoint := oteltef.NewPoint()
		*dstPoints = append(*dstPoints, dstPoint)

		c.ConvertExemplars(dstPoint.Exemplars(), srcPoint.Exemplars())

		err := c.ConvertHistogram(dstPoint, srcPoint)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *converter) covertExponentialHistogramDataPoints(
	sm *SortedTree,
	rm *oteltef.Resource,
	sms *oteltef.Scope,
	metric pmetric.Metric,
	hist pmetric.ExponentialHistogram,
) error {
	var byMetric *ByMetric
	var byScope *ByScope
	flags := calcMetricFlags(false, hist.AggregationTemporality())
	srcPoints := hist.DataPoints()

	for l := 0; l < srcPoints.Len(); l++ {
		srcPoint := srcPoints.At(l)

		byMetric = sm.ByMetric(metric, oteltef.MetricTypeExpHistogram, flags, nil)
		byResource := byMetric.ByResource(rm)
		byScope = byResource.ByScope(sms)
		c.Otlp2tef.MapSorted(srcPoint.Attributes(), &c.TempAttrs)

		c.Otlp2tef.MapSorted(srcPoint.Attributes(), &c.TempAttrs)
		dstPoints := byScope.ByAttrs(&c.TempAttrs)
		dstPoint := oteltef.NewPoint()
		*dstPoints = append(*dstPoints, dstPoint)

		c.ConvertExemplars(dstPoint.Exemplars(), srcPoint.Exemplars())

		err := c.ConvertExpHistogram(dstPoint, srcPoint)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *converter) covertSummaryDataPoints(
	sm *SortedTree,
	rm *oteltef.Resource,
	sms *oteltef.Scope,
	metric pmetric.Metric,
	summary pmetric.Summary,
) error {
	var byMetric *ByMetric
	var byScope *ByScope
	flags := internal.MetricFlags{} // No monotonic/temporality for summary
	srcPoints := summary.DataPoints()

	for l := 0; l < srcPoints.Len(); l++ {
		srcPoint := srcPoints.At(l)

		byMetric = sm.ByMetric(metric, oteltef.MetricTypeSummary, flags, nil)
		byResource := byMetric.ByResource(rm)
		byScope = byResource.ByScope(sms)
		c.Otlp2tef.MapSorted(srcPoint.Attributes(), &c.TempAttrs)
		dstPoints := byScope.ByAttrs(&c.TempAttrs)
		dstPoint := oteltef.NewPoint()
		*dstPoints = append(*dstPoints, dstPoint)

		dstPoint.SetTimestamp(uint64(srcPoint.Timestamp()))
		dstPoint.SetStartTimestamp(uint64(srcPoint.StartTimestamp()))
		// No exemplars for summary

		err := c.ConvertSummary(dstPoint, srcPoint)
		if err != nil {
			return err
		}
	}
	return nil
}
