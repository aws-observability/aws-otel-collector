// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cadvisor

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
// 	"go.opentelemetry.io/collector/consumer/pdata"
// 	"go.uber.org/zap"
// )

// func convertToOTLPMetrics(fields map[string]interface{}, tags map[string]string, logger *zap.Logger) pdata.ResourceMetricsSlice {
// 	rms := pdata.NewResourceMetricsSlice()
// 	// logger.Info("convertToOTLPMetrics", zap.Any("cadvisorMetrics", m.GetFields()), zap.String("timestamp", m.GetAllTags()[common.Timestamp]))
// 	rms.Resize(1)
// 	rm := rms.At(0)

// 	var timestamp pdata.TimestampUnixNano
// 	resource := rm.Resource()
// 	for tagKey, tagValue := range tags {
// 		if tagKey == common.Timestamp {
// 			time, _ := strconv.ParseUint(tagValue, 10, 64)
// 			timestamp = pdata.TimestampUnixNano(time)
// 			continue
// 		}
// 		resource.Attributes().UpsertString(tagKey, tagValue)
// 	}

// 	ilms := rm.InstrumentationLibraryMetrics()

// 	metricType := tags[common.MetricType]
// 	for key, value := range fields {
// 		metric := common.RemovePrefix(metricType, key)
// 		unit := common.GetUnitForMetric(metric)
// 		// val_type := fmt.Sprintf("%T", value)
// 		// logger.Info("DEBUG", zap.Any("metric: ", key), zap.Any("unit: ", unit), zap.Any("value", value), zap.Any("type", val_type))
// 		switch t := value.(type) {
// 		case int:
// 			ilms.Append(intGauge(key, unit, int64(t), timestamp))
// 		case int32:
// 			ilms.Append(intGauge(key, unit, int64(t), timestamp))
// 		case int64:
// 			ilms.Append(intGauge(key, unit, t, timestamp))
// 		case uint:
// 			ilms.Append(doubleGauge(key, unit, float64(t), timestamp))
// 		case uint32:
// 			ilms.Append(doubleGauge(key, unit, float64(t), timestamp))
// 		case uint64:
// 			ilms.Append(doubleGauge(key, unit, float64(t), timestamp))
// 		case float32:
// 			ilms.Append(doubleGauge(key, unit, float64(t), timestamp))
// 		case float64:
// 			ilms.Append(doubleGauge(key, unit, t, timestamp))
// 		default:
// 			valueType := fmt.Sprintf("%T", value)
// 			logger.Warn("Detected unexpected field", zap.String("key", key), zap.Any("value", value), zap.String("value type", valueType))
// 		}
// 	}

// 	return rms
// }

// func intGauge(metricName string, unit string, value int64, ts pdata.TimestampUnixNano) pdata.InstrumentationLibraryMetrics {
// 	ilm := pdata.NewInstrumentationLibraryMetrics()

// 	metric := initMetric(ilm, metricName, unit)

// 	metric.SetDataType(pdata.MetricDataTypeIntGauge)
// 	intGauge := metric.IntGauge()

// 	updateIntDataPoint(intGauge.DataPoints(), value, ts)

// 	return ilm
// }

// func intSum(metricName string, unit string, value int64, ts pdata.TimestampUnixNano) pdata.InstrumentationLibraryMetrics {
// 	ilm := pdata.NewInstrumentationLibraryMetrics()

// 	metric := initMetric(ilm, metricName, unit)

// 	metric.SetDataType(pdata.MetricDataTypeIntSum)
// 	intSum := metric.IntSum()
// 	intSum.SetAggregationTemporality(pdata.AggregationTemporalityCumulative)

// 	updateIntDataPoint(intSum.DataPoints(), value, ts)

// 	return ilm
// }

// func doubleGauge(metricName string, unit string, value float64, ts pdata.TimestampUnixNano) pdata.InstrumentationLibraryMetrics {
// 	ilm := pdata.NewInstrumentationLibraryMetrics()

// 	metric := initMetric(ilm, metricName, unit)

// 	metric.SetDataType(pdata.MetricDataTypeDoubleGauge)
// 	doubleGauge := metric.DoubleGauge()
// 	dataPoints := doubleGauge.DataPoints()
// 	dataPoints.Resize(1)
// 	dataPoint := dataPoints.At(0)

// 	dataPoint.SetValue(value)
// 	dataPoint.SetTimestamp(ts)

// 	return ilm
// }

// func updateIntDataPoint(dataPoints pdata.IntDataPointSlice, value int64, ts pdata.TimestampUnixNano) {
// 	dataPoints.Resize(1)
// 	dataPoint := dataPoints.At(0)
// 	dataPoint.SetValue(value)
// 	dataPoint.SetTimestamp(ts)
// }

// func initMetric(ilm pdata.InstrumentationLibraryMetrics, name, unit string) pdata.Metric {
// 	ilm.Metrics().Resize(1)
// 	metric := ilm.Metrics().At(0)
// 	metric.SetName(name)
// 	metric.SetUnit(unit)

// 	return metric
// }
