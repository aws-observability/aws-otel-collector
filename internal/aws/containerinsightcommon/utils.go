package common

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
)

// aggregate fileds
func AggregateFields(fields []map[string]interface{}) map[string]float64 {
	if len(fields) == 0 {
		return nil
	}

	result := make(map[string]float64)
	// Use the first element as the base
	for k, v := range fields[0] {
		result[k] = v.(float64)
	}

	if len(fields) == 1 {
		return result
	}

	for i := 1; i < len(fields); i++ {
		for k, v := range result {
			if fields[i][k] == nil {
				continue
			}
			result[k] = v + fields[i][k].(float64)
		}
	}
	return result
}

func IsNode(mType string) bool {
	return mType == TypeNode || mType == TypeNodeNet || mType == TypeNodeFS || mType == TypeNodeDiskIO
}
func IsInstance(mType string) bool {
	return mType == TypeInstance || mType == TypeInstanceNet || mType == TypeInstanceFS || mType == TypeInstanceDiskIO
}
func IsContainer(mType string) bool {
	return mType == TypeContainer || mType == TypeContainerDiskIO || mType == TypeContainerFS
}
func IsPod(mType string) bool {
	return mType == TypePod || mType == TypePodNet
}

func getPrefixByMetricType(mType string) string {
	prefix := ""
	instancePrefix := "instance_"
	nodePrefix := "node_"
	instanceNetPrefix := "instance_interface_"
	nodeNetPrefix := "node_interface_"
	podPrefix := "pod_"
	podNetPrefix := "pod_interface_"
	containerPrefix := "container_"
	service := "service_"
	cluster := "cluster_"
	namespace := "namespace_"

	switch mType {
	case TypeInstance:
		prefix = instancePrefix
	case TypeInstanceFS:
		prefix = instancePrefix
	case TypeInstanceDiskIO:
		prefix = instancePrefix
	case TypeInstanceNet:
		prefix = instanceNetPrefix
	case TypeNode:
		prefix = nodePrefix
	case TypeNodeFS:
		prefix = nodePrefix
	case TypeNodeDiskIO:
		prefix = nodePrefix
	case TypeNodeNet:
		prefix = nodeNetPrefix
	case TypePod:
		prefix = podPrefix
	case TypePodNet:
		prefix = podNetPrefix
	case TypeContainer:
		prefix = containerPrefix
	case TypeContainerDiskIO:
		prefix = containerPrefix
	case TypeContainerFS:
		prefix = containerPrefix
	case TypeService:
		prefix = service
	case TypeCluster:
		prefix = cluster
	case TypeClusterService:
		prefix = service
	case TypeClusterNamespace:
		prefix = namespace
	case K8sNamespace:
		prefix = namespace
	default:
		log.Printf("E! Unexpected MetricType: %s", mType)
	}
	return prefix
}

func MetricName(mType string, name string) string {
	return getPrefixByMetricType(mType) + name
}

func RemovePrefix(mType string, metricName string) string {
	prefix := getPrefixByMetricType(mType)
	return strings.Replace(metricName, prefix, "", 1)
}

func GetUnitForMetric(metric string) string {
	return metricToUnitMap[metric]
}

func ConvertToOTLPMetrics(fields map[string]interface{}, tags map[string]string, logger *zap.Logger) pdata.Metrics {
	md := pdata.NewMetrics()
	rms := md.ResourceMetrics()
	rms.Resize(1)
	rm := rms.At(0)

	var timestamp pdata.Timestamp
	resource := rm.Resource()
	for tagKey, tagValue := range tags {
		if tagKey == Timestamp {
			timeNs, _ := strconv.ParseUint(tagValue, 10, 64)
			timestamp = pdata.Timestamp(timeNs)
			// convert from nanosecond to millisecond (as emf log use millisecond timestamp)
			tagValue = strconv.FormatUint(timeNs/uint64(time.Millisecond), 10)
		}
		resource.Attributes().UpsertString(tagKey, tagValue)
	}

	ilms := rm.InstrumentationLibraryMetrics()

	metricType := tags[MetricType]
	for key, value := range fields {
		metric := RemovePrefix(metricType, key)
		unit := GetUnitForMetric(metric)
		switch t := value.(type) {
		case int:
			ilms.Append(intGauge(key, unit, int64(t), timestamp))
		case int32:
			ilms.Append(intGauge(key, unit, int64(t), timestamp))
		case int64:
			ilms.Append(intGauge(key, unit, t, timestamp))
		case uint:
			ilms.Append(doubleGauge(key, unit, float64(t), timestamp))
		case uint32:
			ilms.Append(doubleGauge(key, unit, float64(t), timestamp))
		case uint64:
			ilms.Append(doubleGauge(key, unit, float64(t), timestamp))
		case float32:
			ilms.Append(doubleGauge(key, unit, float64(t), timestamp))
		case float64:
			ilms.Append(doubleGauge(key, unit, t, timestamp))
		default:
			valueType := fmt.Sprintf("%T", value)
			logger.Warn("Detected unexpected field", zap.String("key", key), zap.Any("value", value), zap.String("value type", valueType))
		}
	}

	return md
}

func intGauge(metricName string, unit string, value int64, ts pdata.Timestamp) pdata.InstrumentationLibraryMetrics {
	ilm := pdata.NewInstrumentationLibraryMetrics()

	metric := initMetric(ilm, metricName, unit)

	metric.SetDataType(pdata.MetricDataTypeIntGauge)
	intGauge := metric.IntGauge()

	updateIntDataPoint(intGauge.DataPoints(), value, ts)

	return ilm
}

func doubleGauge(metricName string, unit string, value float64, ts pdata.Timestamp) pdata.InstrumentationLibraryMetrics {
	ilm := pdata.NewInstrumentationLibraryMetrics()

	metric := initMetric(ilm, metricName, unit)

	metric.SetDataType(pdata.MetricDataTypeDoubleGauge)
	doubleGauge := metric.DoubleGauge()
	dataPoints := doubleGauge.DataPoints()
	dataPoints.Resize(1)
	dataPoint := dataPoints.At(0)

	dataPoint.SetValue(value)
	dataPoint.SetTimestamp(ts)

	return ilm
}

func updateIntDataPoint(dataPoints pdata.IntDataPointSlice, value int64, ts pdata.Timestamp) {
	dataPoints.Resize(1)
	dataPoint := dataPoints.At(0)
	dataPoint.SetValue(value)
	dataPoint.SetTimestamp(ts)
}

func initMetric(ilm pdata.InstrumentationLibraryMetrics, name, unit string) pdata.Metric {
	ilm.Metrics().Resize(1)
	metric := ilm.Metrics().At(0)
	metric.SetName(name)
	metric.SetUnit(unit)

	return metric
}
