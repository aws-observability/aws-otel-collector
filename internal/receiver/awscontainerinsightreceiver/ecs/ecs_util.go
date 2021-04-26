package ecs

import (
	"fmt"
	common "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
	"reflect"
)

func insertMetrics(key string, value interface{}, ilms pdata.InstrumentationLibraryMetricsSlice, metricType string, timestamp pdata.Timestamp, logger *zap.Logger) {

	metric := common.RemovePrefix(metricType, key)
	unit := common.GetUnitForMetric(metric)

	switch t := value.(type) {
	case int:
		ilms.Append(common.IntGauge(key, unit, int64(t), timestamp))
	case int32:
		ilms.Append(common.IntGauge(key, unit, int64(t), timestamp))
	case int64:
		ilms.Append(common.IntGauge(key, unit, t, timestamp))
	case uint:
		ilms.Append(common.DoubleGauge(key, unit, float64(t), timestamp))
	case uint32:
		ilms.Append(common.DoubleGauge(key, unit, float64(t), timestamp))
	case uint64:
		ilms.Append(common.DoubleGauge(key, unit, float64(t), timestamp))
	case float32:
		ilms.Append(common.DoubleGauge(key, unit, float64(t), timestamp))
	case float64:
		ilms.Append(common.DoubleGauge(key, unit, t, timestamp))
	default:
		valueType := fmt.Sprintf("%T", value)
		logger.Warn("Detected unexpected field", zap.String("key", key), zap.Any("value", value), zap.String("value type", valueType))
	}
}

func getMetricValue(metricName string, rms pdata.ResourceMetrics) (bool, interface{}, pdata.Timestamp) {
	ilms := rms.InstrumentationLibraryMetrics()
	for i := 0; i < ilms.Len(); i++ {
		ilm := ilms.At(i)
		metricSlice := ilm.Metrics()
		for j := 0; j < metricSlice.Len(); j++ {
			metric := metricSlice.At(j)
			if metric.Name() == metricName {
				if metric.DataType() == pdata.MetricDataTypeIntGauge {
					dataPoints := metric.IntGauge().DataPoints()
					if dataPoints.Len() > 0 {
						timestamp := dataPoints.At(0).Timestamp()
						val := dataPoints.At(0).Value()
						return true, val, timestamp
					}
				} else if metric.DataType() == pdata.MetricDataTypeDoubleGauge {
					dataPoints := metric.DoubleGauge().DataPoints()
					if dataPoints.Len() > 0 {
						timestamp := dataPoints.At(0).Timestamp()
						val := dataPoints.At(0).Value()
						return true, val, timestamp
					}
				}
			}
		}
	}
	return false, nil, 0
}

func AddAttribute(key string, value interface{}, attributes pdata.AttributeMap) {
	attributes.Upsert(key, fromVal(value))
}
func fromVal(v interface{}) pdata.AttributeValue {
	switch val := v.(type) {
	case string:
		return pdata.NewAttributeValueString(val)
	case int:
		return pdata.NewAttributeValueInt(int64(val))
	case float64:
		return pdata.NewAttributeValueDouble(val)
	case map[string]interface{}:
		return fromMap(val)
	case map[string]string:
		return fromStringMap(val)
	case []interface{}:
		return fromArray(val)
	case []string:
		return fromStringArray(val)
	}

	panic("data type is not supported in fromVal()" + reflect.TypeOf(v).String())
}

func fromMap(v map[string]interface{}) pdata.AttributeValue {
	av := pdata.NewAttributeValueMap()
	m := av.MapVal()
	for k, v := range v {
		m.Insert(k, fromVal(v))
	}
	m.Sort()
	return av
}

func fromStringMap(v map[string]string) pdata.AttributeValue {
	av := pdata.NewAttributeValueMap()
	m := av.MapVal()
	for k, v := range v {
		m.Insert(k, fromVal(v))
	}
	m.Sort()
	return av
}
func fromArray(v []interface{}) pdata.AttributeValue {
	av := pdata.NewAttributeValueArray()
	arr := av.ArrayVal()
	for _, v := range v {
		arr.Append(fromVal(v))
	}
	return av
}

func fromStringArray(v []string) pdata.AttributeValue {
	av := pdata.NewAttributeValueArray()
	arr := av.ArrayVal()
	for _, v := range v {
		arr.Append(fromVal(v))
	}
	return av
}
