package metrics

import (
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type PDataSorter struct {
	leftAttrs, rightAttrs []attrAccessible
}

func (s *PDataSorter) SortMetrics(data pmetric.Metrics, combineSlices bool) {
	resourceMetrics := data.ResourceMetrics()

	resourceMetrics.Sort(
		func(a, b pmetric.ResourceMetrics) bool {
			return s.cmpResourceMetrics(a, b) < 0
		},
	)

	if combineSlices {
		for i := 0; i < resourceMetrics.Len()-1; {
			if s.cmpResourceMetrics(resourceMetrics.At(i), resourceMetrics.At(i+1)) == 0 {
				resourceMetrics.At(i + 1).ScopeMetrics().MoveAndAppendTo(resourceMetrics.At(i).ScopeMetrics())
				j := 0
				resourceMetrics.RemoveIf(
					func(metrics pmetric.ResourceMetrics) bool {
						j++
						return j == i+2
					},
				)
			} else {
				i++
			}
		}
	}

	for i := 0; i < resourceMetrics.Len(); i++ {
		s.sortScopeMetrics(resourceMetrics.At(i).ScopeMetrics(), combineSlices)
	}
}

func (s *PDataSorter) sortScopeMetrics(scopeMetrics pmetric.ScopeMetricsSlice, combineSlices bool) {
	scopeMetrics.Sort(
		func(a, b pmetric.ScopeMetrics) bool {
			return s.cmpScopeMetrics(a, b) < 0
		},
	)

	if combineSlices {
		for i := 0; i < scopeMetrics.Len()-1; {
			if s.cmpScopeMetrics(scopeMetrics.At(i), scopeMetrics.At(i+1)) == 0 {
				scopeMetrics.At(i + 1).Metrics().MoveAndAppendTo(scopeMetrics.At(i).Metrics())
				j := 0
				scopeMetrics.RemoveIf(
					func(s pmetric.ScopeMetrics) bool {
						j++
						return j == i+2
					},
				)
			} else {
				i++
			}
		}
	}

	for i := 0; i < scopeMetrics.Len(); i++ {
		s.sortMetricSlice(scopeMetrics.At(i).Metrics(), combineSlices)
	}
}

func (s *PDataSorter) sortMetricSlice(metrics pmetric.MetricSlice, combineSlices bool) {
	metrics.Sort(
		func(a, b pmetric.Metric) bool {
			return s.cmpMetric(a, b) < 0
		},
	)

	if combineSlices {
		for i := 0; i < metrics.Len()-1; {
			if s.cmpMetric(metrics.At(i), metrics.At(i+1)) == 0 {
				appendMetricData(metrics.At(i), metrics.At(i+1))
				j := 0
				metrics.RemoveIf(
					func(s pmetric.Metric) bool {
						j++
						return j == i+2
					},
				)
			} else {
				i++
			}
		}
	}

	for i := 0; i < metrics.Len(); i++ {
		s.sortMetricData(metrics.At(i))
	}
}

func (s *PDataSorter) sortMetricData(metric pmetric.Metric) {
	switch metric.Type() {
	case pmetric.MetricTypeGauge:
		sortNumberDatapointSlice(metric.Gauge().DataPoints())
		metric.Gauge().DataPoints().Sort(
			func(a, b pmetric.NumberDataPoint) bool {
				return s.cmpNumberDataPoint(a, b) < 0
			},
		)
	case pmetric.MetricTypeSum:
		sortNumberDatapointSlice(metric.Sum().DataPoints())
		metric.Sum().DataPoints().Sort(
			func(a, b pmetric.NumberDataPoint) bool {
				return s.cmpNumberDataPoint(a, b) < 0
			},
		)
	case pmetric.MetricTypeHistogram:
		metric.Histogram().DataPoints().Sort(
			func(a, b pmetric.HistogramDataPoint) bool {
				return s.cmpHistogramDataPoint(a, b) < 0
			},
		)
	default:
		panic("not implemented")
	}

}

func sortNumberDatapointSlice(points pmetric.NumberDataPointSlice) {
	points.RemoveIf(
		func(point pmetric.NumberDataPoint) bool {
			return point.ValueType() == pmetric.NumberDataPointValueTypeEmpty
		},
	)
}

func (s *PDataSorter) cmpHistogramDataPoint(left pmetric.HistogramDataPoint, right pmetric.HistogramDataPoint) int {
	c := s.cmpAttrs(left.Attributes(), right.Attributes())
	if c != 0 {
		return c
	}

	if left.Timestamp() < right.Timestamp() {
		return -1
	}
	if left.Timestamp() > right.Timestamp() {
		return 1
	}
	return 0
}

func (s *PDataSorter) cmpNumberDataPoint(left pmetric.NumberDataPoint, right pmetric.NumberDataPoint) int {
	c := s.cmpAttrs(left.Attributes(), right.Attributes())
	if c != 0 {
		return c
	}

	if left.Timestamp() < right.Timestamp() {
		return -1
	}
	if left.Timestamp() > right.Timestamp() {
		return 1
	}
	return 0
}

func appendMetricData(left, right pmetric.Metric) {
	switch left.Type() {
	case pmetric.MetricTypeGauge:
		right.Gauge().DataPoints().MoveAndAppendTo(left.Gauge().DataPoints())
	case pmetric.MetricTypeSum:
		right.Sum().DataPoints().MoveAndAppendTo(left.Sum().DataPoints())
	case pmetric.MetricTypeHistogram:
		right.Histogram().DataPoints().MoveAndAppendTo(left.Histogram().DataPoints())
	default:
		panic("not implemented")
	}
}

type attrAccessible struct {
	Key   string
	Value pcommon.Value
}

func map2attrs(src pcommon.Map, dst *[]attrAccessible) {
	*dst = (*dst)[:0]
	src.Range(
		func(k string, v pcommon.Value) bool {
			*dst = append(*dst, attrAccessible{Key: k, Value: v})
			return true
		},
	)
}

func (s *PDataSorter) cmpResourceMetrics(left, right pmetric.ResourceMetrics) int {
	c := strings.Compare(left.SchemaUrl(), right.SchemaUrl())
	if c != 0 {
		return c
	}
	return s.cmpAttrs(left.Resource().Attributes(), right.Resource().Attributes())
}

func (s *PDataSorter) cmpScopeMetrics(left, right pmetric.ScopeMetrics) int {
	c := strings.Compare(left.Scope().Name(), right.Scope().Name())
	if c != 0 {
		return c
	}
	c = strings.Compare(left.Scope().Version(), right.Scope().Version())
	if c != 0 {
		return c
	}
	c = strings.Compare(left.SchemaUrl(), right.SchemaUrl())
	if c != 0 {
		return c
	}

	return s.cmpAttrs(left.Scope().Attributes(), right.Scope().Attributes())
}

func (s *PDataSorter) cmpMetric(left, right pmetric.Metric) int {
	c := strings.Compare(left.Name(), right.Name())
	if c != 0 {
		return c
	}
	c = strings.Compare(left.Unit(), right.Unit())
	if c != 0 {
		return c
	}
	c = strings.Compare(left.Description(), right.Description())
	if c != 0 {
		return c
	}
	c = s.cmpAttrs(left.Metadata(), right.Metadata())
	if c != 0 {
		return c
	}
	return metricTypeIndex(left) - metricTypeIndex(right)
}

func metricTypeIndex(metric pmetric.Metric) int {
	switch metric.Type() {
	case pmetric.MetricTypeGauge:
		return 0
	case pmetric.MetricTypeSum:
		return 1
	case pmetric.MetricTypeHistogram:
		return 2
	case pmetric.MetricTypeSummary:
		return 3
	case pmetric.MetricTypeExponentialHistogram:
		return 4
	default:
		panic("unknown metric type")
	}
}

func (s *PDataSorter) cmpAttrs(a, b pcommon.Map) int {
	map2attrs(a, &s.leftAttrs)
	map2attrs(b, &s.rightAttrs)

	lenDiff := len(s.leftAttrs) - len(s.rightAttrs)
	l := min(len(s.leftAttrs), len(s.rightAttrs))
	for i := 0; i < l; i++ {
		c := strings.Compare(s.leftAttrs[i].Key, s.rightAttrs[i].Key)
		if c != 0 {
			return c
		}
	}
	if lenDiff != 0 {
		return lenDiff
	}
	for i := 0; i < l; i++ {
		c := cmpVal(s.leftAttrs[i].Value, s.rightAttrs[i].Value)
		if c != 0 {
			return c
		}
	}
	return lenDiff
}

func cmpVal(left, right pcommon.Value) int {
	c := int(left.Type()) - int(right.Type())
	if c != 0 {
		return c
	}

	switch left.Type() {
	case pcommon.ValueTypeStr:
		return strings.Compare(left.Str(), right.Str())

	case pcommon.ValueTypeInt:
		return cmpInt64(left.Int(), right.Int())

	case pcommon.ValueTypeBool:
		return cmpBool(left.Bool(), right.Bool())

	case pcommon.ValueTypeSlice:
		left := left.Slice()
		right := right.Slice()
		if left.Len() != right.Len() {
			return left.Len() - right.Len()
		}
		for i := 0; i < left.Len(); i++ {
			c := cmpVal(left.At(i), right.At(i))
			if c != 0 {
				return c
			}
		}
		return 0

	case pcommon.ValueTypeEmpty:
		return 0

	default:
		panic("comparison not implemented")
	}
}

func cmpBool(v1, v2 bool) int {
	if v1 == v2 {
		return 0
	}
	if v1 {
		return 1
	}
	return -1
}

func cmpInt64(v1, v2 int64) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
}
