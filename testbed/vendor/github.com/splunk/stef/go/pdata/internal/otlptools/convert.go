package otlptools

import (
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/splunk/stef/go/otel/oteltef"
)

func ResourceToOtlp(resource *oteltef.Resource, out pmetric.ResourceMetrics) error {
	if resource == nil {
		return nil
	}
	out.SetSchemaUrl(resource.SchemaURL())
	var err error
	err = TefToOtlpMap(resource.Attributes(), out.Resource().Attributes())
	if err != nil {
		return err
	}
	out.Resource().SetDroppedAttributesCount(uint32(resource.DroppedAttributesCount()))
	return nil
}

func ScopeToOtlp(scope *oteltef.Scope, out pmetric.ScopeMetrics) error {
	if scope == nil {
		return nil
	}
	out.SetSchemaUrl(scope.SchemaURL())
	out.Scope().SetName(scope.Name())
	out.Scope().SetVersion(scope.Version())
	var err error
	err = TefToOtlpMap(scope.Attributes(), out.Scope().Attributes())
	if err != nil {
		return err
	}
	out.Scope().SetDroppedAttributesCount(uint32(scope.DroppedAttributesCount()))
	return nil
}

func MetricToOtlp(metric *oteltef.Metric, out pmetric.Metric) error {
	if metric == nil {
		return nil
	}
	out.SetName(metric.Name())
	out.SetUnit(metric.Unit())
	out.SetDescription(metric.Description())

	var err error
	err = TefToOtlpMap(metric.Metadata(), out.Metadata())
	if err != nil {
		return err
	}

	switch oteltef.MetricType(metric.Type()) {
	case oteltef.MetricTypeGauge:
		out.SetEmptyGauge()
	case oteltef.MetricTypeSum:
		out.SetEmptySum()
	case oteltef.MetricTypeHistogram:
		out.SetEmptyHistogram()
	case oteltef.MetricTypeExpHistogram:
		out.SetEmptyExponentialHistogram()
	case oteltef.MetricTypeSummary:
		out.SetEmptySummary()
	default:
		panic("not implemented")
	}

	return nil
}
