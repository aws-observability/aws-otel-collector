package stores

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	. "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"github.com/aws-observability/aws-otel-collector/internal/aws/k8scommon/k8sutil"
	"go.opentelemetry.io/collector/consumer/pdata"
	corev1 "k8s.io/api/core/v1"
)

func createPodKeyFromMetaData(pod *corev1.Pod) string {
	namespace := pod.Namespace
	podName := pod.Name
	return k8sutil.CreatePodKey(namespace, podName)
}

func GetStringValFromAttributes(key string, attributes pdata.AttributeMap) string {
	var strVal string
	if val, ok := attributes.Get(key); ok {
		strVal = val.StringVal()
	}
	return strVal
}

func AddAttribute(key string, value interface{}, attributes pdata.AttributeMap) {
	attributes.Upsert(key, fromVal(value))
}

func RemoveAttribute(key string, attributes pdata.AttributeMap) {
	attributes.Delete(key)
}

func createPodKeyFromMetric(attributes pdata.AttributeMap) string {
	namespace := GetStringValFromAttributes(K8sNamespace, attributes)
	podName := GetStringValFromAttributes(K8sPodNameKey, attributes)
	return k8sutil.CreatePodKey(namespace, podName)
}

func createContainerKeyFromMetric(attributes pdata.AttributeMap) string {
	namespace := GetStringValFromAttributes(K8sNamespace, attributes)
	podName := GetStringValFromAttributes(K8sPodNameKey, attributes)
	containerName := GetStringValFromAttributes(ContainerNamekey, attributes)
	return k8sutil.CreateContainerKey(namespace, podName, containerName)
}

const (
	// kubeAllowedStringAlphaNums holds the characters allowed in replicaset names from as parent deployment
	// https://github.com/kubernetes/apimachinery/blob/master/pkg/util/rand/rand.go#L83
	kubeAllowedStringAlphaNums = "bcdfghjklmnpqrstvwxz2456789"
	cronJobAllowedString       = "0123456789"
)

// get the deployment name by stripping the last dash following some rules
// return empty if it is not following the rule
func parseDeploymentFromReplicaSet(name string) string {
	lastDash := strings.LastIndexAny(name, "-")
	if lastDash == -1 {
		// No dash
		return ""
	}
	suffix := name[lastDash+1:]
	if len(suffix) < 3 {
		// Invalid suffix if it is less than 3
		return ""
	}

	if !stringInRuneset(suffix, kubeAllowedStringAlphaNums) {
		// Invalid suffix
		return ""
	}

	return name[:lastDash]
}

// get the cronJob name by stripping the last dash following some rules
// return empty if it is not following the rule
func parseCronJobFromJob(name string) string {
	lastDash := strings.LastIndexAny(name, "-")
	if lastDash == -1 {
		// No dash
		return ""
	}
	suffix := name[lastDash+1:]
	if len(suffix) != 10 {
		// Invalid suffix if it is not 10 rune
		return ""
	}

	if !stringInRuneset(suffix, cronJobAllowedString) {
		// Invalid suffix
		return ""
	}

	return name[:lastDash]
}

func stringInRuneset(name, subset string) bool {
	for _, r := range name {
		if !strings.ContainsRune(subset, r) {
			// Found an unexpected rune in suffix
			return false
		}
	}
	return true
}

//used to create an attribute out of a json
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
	fmt.Printf("value = %v", v)
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

func fromJSONMap(jsonStr string) pdata.AttributeValue {
	var src map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &src)
	if err != nil {
		panic("Invalid input jsonStr:" + jsonStr)
	}
	return fromMap(src)
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

func fromJSONArray(jsonStr string) pdata.AttributeValue {
	var src []interface{}
	err := json.Unmarshal([]byte(jsonStr), &src)
	if err != nil {
		panic("Invalid input jsonStr:" + jsonStr)
	}
	return fromArray(src)
}

func getMetricValue(metricName string, rms pdata.ResourceMetrics) interface{} {
	ilms := rms.InstrumentationLibraryMetrics()
	for i := 0; i < ilms.Len(); i++ {
		ilm := ilms.At(i)
		metricSlice := ilm.Metrics()
		for j := 0; j < metricSlice.Len(); j++ {
			metric := metricSlice.At(j)
			if metric.Name() == metricName {
				if metric.DataType() == pdata.MetricDataTypeIntGauge {
					dataPoints := metric.IntGauge().DataPoints()
					//assume there is only one datapoint
					return dataPoints.At(0).Value()
				} else if metric.DataType() == pdata.MetricDataTypeDoubleGauge {
					dataPoints := metric.DoubleGauge().DataPoints()
					//assume there is only one datapoint
					return dataPoints.At(0).Value()
				}
			}
		}
	}
	return nil
}

func containsMetric(metricName string, rms pdata.ResourceMetrics) bool {
	return getMetricValue(metricName, rms) != nil
}

func addMetric(metricName string, metricValue interface{}, unit string, rms pdata.ResourceMetrics) bool {
	ilms := rms.InstrumentationLibraryMetrics()
	var timestamp pdata.Timestamp
	var newMetric pdata.InstrumentationLibraryMetrics

	//find the timestamp for current metrics
	for i := 0; i < ilms.Len(); i++ {
		ilm := ilms.At(i)
		metricSlice := ilm.Metrics()
		if metricSlice.Len() > 0 {
			metric := metricSlice.At(0)
			//only handle the two cases
			if metric.DataType() == pdata.MetricDataTypeIntGauge {
				dataPoints := metric.IntGauge().DataPoints()
				if dataPoints.Len() > 0 {
					timestamp = dataPoints.At(0).Timestamp()
					break
				}
			} else if metric.DataType() == pdata.MetricDataTypeDoubleGauge {
				dataPoints := metric.DoubleGauge().DataPoints()
				if dataPoints.Len() > 0 {
					timestamp = dataPoints.At(0).Timestamp()
					break
				}
			}
		}
	}

	if timestamp == 0 {
		return false
	}

	switch t := metricValue.(type) {
	case int:
		newMetric = intGauge(metricName, unit, int64(t), timestamp)
	case int32:
		newMetric = intGauge(metricName, unit, int64(t), timestamp)
	case int64:
		newMetric = intGauge(metricName, unit, t, timestamp)
	case uint:
		newMetric = doubleGauge(metricName, unit, float64(t), timestamp)
	case uint32:
		newMetric = doubleGauge(metricName, unit, float64(t), timestamp)
	case uint64:
		newMetric = doubleGauge(metricName, unit, float64(t), timestamp)
	case float32:
		newMetric = doubleGauge(metricName, unit, float64(t), timestamp)
	case float64:
		newMetric = doubleGauge(metricName, unit, t, timestamp)
	default:
		// valueType := fmt.Sprintf("%T", value)
		// logger.Warn("Detected unexpected field", zap.String("key", metricName), zap.Any("value", metricValue), zap.String("value type", valueType))
		return false
	}

	ilms.Append(newMetric)
	return true
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
