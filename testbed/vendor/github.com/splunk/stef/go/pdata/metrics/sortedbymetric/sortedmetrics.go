package sortedbymetric

import (
	"io"
	"slices"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"modernc.org/b/v2"

	"github.com/splunk/stef/go/pdata/metrics/internal"
	"github.com/splunk/stef/go/pkg"

	"github.com/splunk/stef/go/otel/oteltef"

	"github.com/splunk/stef/go/pdata/internal/otlptools"
)

type SortedTree struct {
	//encoder  anyvalue.Encoder
	otlp2tef otlptools.Otlp2Stef
	byMetric *b.Tree[*oteltef.Metric, *ByMetric]
}

type ByMetric struct {
	//encoder    *anyvalue.Encoder
	otlp2tef   *otlptools.Otlp2Stef
	byResource *b.Tree[*oteltef.Resource, *ByResource]
	allocators *oteltef.Allocators
}

type ByResource struct {
	//encoder  *anyvalue.Encoder
	otlp2tef   *otlptools.Otlp2Stef
	byScope    *b.Tree[*oteltef.Scope, *ByScope]
	allocators *oteltef.Allocators
}

type Points []*oteltef.Point

func (p Points) SortValues() {
	slices.SortFunc(
		p, func(a, b *oteltef.Point) int {
			return pkg.Uint64Compare(a.Timestamp(), b.Timestamp())
		},
	)
}

type ByScope struct {
	byAttrs    *b.Tree[*oteltef.Attributes, *Points]
	allocators *oteltef.Allocators
}

func NewSortedMetrics() *SortedTree {
	return &SortedTree{byMetric: b.TreeNew[*oteltef.Metric, *ByMetric](oteltef.CmpMetric)}
}

func (s *SortedTree) ToStef(writer *oteltef.MetricsWriter) error {
	i := 0
	err := s.Iter(
		func(metric *oteltef.Metric, byMetric *ByMetric) error {
			writer.Record.SetMetric(metric)
			err := byMetric.Iter(
				func(resource *oteltef.Resource, byResource *ByResource) error {
					writer.Record.SetResource(resource)
					err := byResource.Iter(
						func(scope *oteltef.Scope, byScope *ByScope) error {
							writer.Record.SetScope(scope)
							err := byScope.Iter(
								func(attrs *oteltef.Attributes, points *Points) error {
									writer.Record.Attributes().CopyFrom(attrs)
									for _, value := range *points {
										writer.Record.Point().CopyFrom(value)
										if err := writer.Write(); err != nil {
											return err
										}
										i++
									}
									return nil
								},
							)
							return err
						},
					)
					return err
				},
			)
			return err
		},
	)
	return err
}

func (s *SortedTree) ByMetric(
	metric pmetric.Metric, metricType oteltef.MetricType, flags internal.MetricFlags,
	histogramBounds []float64,
) *ByMetric {
	metr := metric2metric(metric, metricType, flags, histogramBounds, &s.otlp2tef)
	elem, exists := s.byMetric.Get(metr)
	if !exists {
		elem = &ByMetric{
			otlp2tef: &s.otlp2tef, byResource: b.TreeNew[*oteltef.Resource, *ByResource](oteltef.CmpResource),
		}
		s.byMetric.Set(metr, elem)
	}
	return elem
}

func metric2metric(
	metric pmetric.Metric, metricType oteltef.MetricType, flags internal.MetricFlags, histogramBounds []float64,
	otlp2tef *otlptools.Otlp2Stef,
) *oteltef.Metric {

	var dst oteltef.Metric
	otlp2tef.MapSorted(metric.Metadata(), dst.Metadata())
	dst.SetName(metric.Name())
	dst.SetDescription(metric.Description())
	dst.SetUnit(metric.Unit())
	dst.SetType(metricType)
	dst.HistogramBounds().CopyFromSlice(histogramBounds)
	dst.SetMonotonic(flags.Monotonic)
	dst.SetAggregationTemporality(flags.Temporality)

	return &dst
}

func (s *SortedTree) SortValues() {
	iter, err := s.byMetric.SeekFirst()
	if err != nil {
		return
	}
	for {
		_, v, err := iter.Next()
		if err == io.EOF {
			break
		}
		v.SortValues()
	}
}

func (s *SortedTree) Iter(f func(metric *oteltef.Metric, byMetric *ByMetric) error) error {
	iter, err := s.byMetric.SeekFirst()
	if err != nil {
		return nil
	}
	for {
		k, v, err := iter.Next()
		if err == io.EOF {
			break
		}
		if err := f(k, v); err != nil {
			return err
		}
	}
	return nil
}

func (m *ByMetric) ByResource(res *oteltef.Resource) *ByResource {
	elem, exists := m.byResource.Get(res)
	if !exists {
		elem = &ByResource{
			otlp2tef: m.otlp2tef, byScope: b.TreeNew[*oteltef.Scope, *ByScope](oteltef.CmpScope),
			allocators: m.allocators,
		}
		m.byResource.Set(res, elem)
	}
	return elem
}

func (m *ByMetric) SortValues() {
	iter, err := m.byResource.SeekFirst()
	if err != nil {
		return
	}
	for {
		_, v, err := iter.Next()
		if err == io.EOF {
			break
		}
		v.SortValues()
	}
}

func (m *ByMetric) Iter(f func(resource *oteltef.Resource, byResource *ByResource) error) error {
	iter, err := m.byResource.SeekFirst()
	if err != nil {
		return nil
	}
	for {
		k, v, err := iter.Next()
		if err == io.EOF {
			break
		}
		if err := f(k, v); err != nil {
			return err
		}
	}
	return nil
}

func (m *ByResource) ByScope(dst *oteltef.Scope) *ByScope {
	elem, exists := m.byScope.Get(dst)
	if !exists {
		elem = &ByScope{
			byAttrs:    b.TreeNew[*oteltef.Attributes, *Points](oteltef.CmpAttributes),
			allocators: m.allocators,
		}
		m.byScope.Set(dst, elem)
	}
	return elem
}

func (m *ByResource) Iter(f func(scope *oteltef.Scope, byScope *ByScope) error) error {
	iter, err := m.byScope.SeekFirst()
	if err != nil {
		return nil
	}
	for {
		k, v, err := iter.Next()
		if err == io.EOF {
			break
		}
		if err := f(k, v); err != nil {
			return err
		}
	}
	return nil
}

func (m *ByResource) SortValues() {
	iter, err := m.byScope.SeekFirst()
	if err != nil {
		return
	}
	for {
		_, v, err := iter.Next()
		if err == io.EOF {
			break
		}
		v.SortValues()
	}
}

func (m *ByScope) ByAttrs(attrs *oteltef.Attributes) *Points {
	elem, exists := m.byAttrs.Get(attrs)
	if !exists {
		elem = new(Points)
		var attrsClone oteltef.Attributes
		attrsClone.CopyFrom(attrs)
		m.byAttrs.Set(&attrsClone, elem)
	}
	return elem
}

func (m *ByScope) SortValues() {
	iter, err := m.byAttrs.SeekFirst()
	if err != nil {
		return
	}
	for {
		_, v, err := iter.Next()
		if err == io.EOF {
			break
		}
		v.SortValues()
	}
}

func (m *ByScope) Iter(f func(attrs *oteltef.Attributes, points *Points) error) error {
	iter, err := m.byAttrs.SeekFirst()
	if err != nil {
		return nil
	}
	for {
		k, v, err := iter.Next()
		if err == io.EOF {
			break
		}
		if err := f(k, v); err != nil {
			return err
		}
	}
	return nil
}
