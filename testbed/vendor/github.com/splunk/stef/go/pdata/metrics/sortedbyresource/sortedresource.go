package sortedbyresource

import (
	"io"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"modernc.org/b/v2"

	"github.com/splunk/stef/go/otel/oteltef"

	"github.com/splunk/stef/go/pdata/internal/otlptools"
	"github.com/splunk/stef/go/pdata/metrics/internal"
)

type SortedTree struct {
	internal.BaseSTEFToOTLP
	byResource *b.Tree[*oteltef.Resource, *ByResource]
	allocators oteltef.Allocators
}

type ByResource struct {
	byScope    *b.Tree[*oteltef.Scope, *ByScope]
	allocators *oteltef.Allocators
}

type ByScope struct {
	byMetrics  *b.Tree[*oteltef.Metric, *ByMetric]
	allocators *oteltef.Allocators
}

type Points []*oteltef.Point

type ByMetric struct {
	byAttrs    *b.Tree[*oteltef.Attributes, *Points]
	allocators *oteltef.Allocators
}

func NewSortedByResource() *SortedTree {
	return &SortedTree{byResource: b.TreeNew[*oteltef.Resource, *ByResource](oteltef.CmpResource)}
}

func (s *SortedTree) ToOtlp() (pmetric.Metrics, error) {
	ret := pmetric.NewMetrics()
	rms := ret.ResourceMetrics()

	err := s.Iter(
		func(resource *oteltef.Resource, byResource *ByResource) error {
			rm := rms.AppendEmpty()
			err := otlptools.ResourceToOtlp(resource, rm)
			if err != nil {
				return err
			}
			err = byResource.Iter(
				func(scope *oteltef.Scope, byScope *ByScope) error {
					sc := rm.ScopeMetrics().AppendEmpty()
					err := otlptools.ScopeToOtlp(scope, sc)
					if err != nil {
						return err
					}
					err = byScope.Iter(
						func(metric *oteltef.Metric, byMetric *ByMetric) error {
							metr := sc.Metrics().AppendEmpty()
							err := otlptools.MetricToOtlp(metric, metr)
							if err != nil {
								return err
							}
							err = byMetric.Iter(
								func(attrs *oteltef.Attributes, points *Points) error {
									for _, point := range *points {
										err := s.AppendOTLPPoint(metric, attrs, point, metr)
										if err != nil {
											return err
										}
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
	return ret, err
}

func (s *SortedTree) ByResource(resource *oteltef.Resource) *ByResource {
	elem, exists := s.byResource.Get(resource)
	if !exists {
		elem = &ByResource{byScope: b.TreeNew[*oteltef.Scope, *ByScope](oteltef.CmpScope), allocators: &s.allocators}
		s.byResource.Set(resource.Clone(&s.allocators), elem)
	}
	return elem
}

func (m *ByResource) ByScope(scope *oteltef.Scope) *ByScope {
	elem, exists := m.byScope.Get(scope)
	if !exists {
		elem = &ByScope{byMetrics: b.TreeNew[*oteltef.Metric, *ByMetric](oteltef.CmpMetric), allocators: m.allocators}
		m.byScope.Set(scope.Clone(m.allocators), elem)
	}
	return elem
}

func (m *ByScope) ByMetric(
	metric *oteltef.Metric,
) *ByMetric {
	elem, exists := m.byMetrics.Get(metric)
	if !exists {
		elem = &ByMetric{
			byAttrs:    b.TreeNew[*oteltef.Attributes, *Points](oteltef.CmpAttributes),
			allocators: m.allocators,
		}
		m.byMetrics.Set(metric.Clone(m.allocators), elem)
	}
	return elem
}

func (m *ByMetric) ByAttrs(attrs *oteltef.Attributes) *Points {
	elem, exists := m.byAttrs.Get(attrs)
	if !exists {
		var clone oteltef.Attributes
		clone.CopyFrom(attrs)
		elem = &Points{}
		m.byAttrs.Set(&clone, elem)
	}
	return elem
}

func (s *SortedTree) Iter(f func(resource *oteltef.Resource, byResource *ByResource) error) error {
	iter, err := s.byResource.SeekFirst()
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

func (m *ByScope) Iter(f func(metric *oteltef.Metric, byMetric *ByMetric) error) error {
	iter, err := m.byMetrics.SeekFirst()
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

func (m *ByMetric) Iter(f func(attrs *oteltef.Attributes, points *Points) error) error {
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
