// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricsAndMetricTagConfigurations - Object for a metrics and metric tag configurations.
type MetricsAndMetricTagConfigurations struct {
	Metric                 *Metric
	MetricTagConfiguration *MetricTagConfiguration

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MetricAsMetricsAndMetricTagConfigurations is a convenience function that returns Metric wrapped in MetricsAndMetricTagConfigurations.
func MetricAsMetricsAndMetricTagConfigurations(v *Metric) MetricsAndMetricTagConfigurations {
	return MetricsAndMetricTagConfigurations{Metric: v}
}

// MetricTagConfigurationAsMetricsAndMetricTagConfigurations is a convenience function that returns MetricTagConfiguration wrapped in MetricsAndMetricTagConfigurations.
func MetricTagConfigurationAsMetricsAndMetricTagConfigurations(v *MetricTagConfiguration) MetricsAndMetricTagConfigurations {
	return MetricsAndMetricTagConfigurations{MetricTagConfiguration: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *MetricsAndMetricTagConfigurations) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Metric
	err = datadog.Unmarshal(data, &obj.Metric)
	if err == nil {
		if obj.Metric != nil && obj.Metric.UnparsedObject == nil {
			jsonMetric, _ := datadog.Marshal(obj.Metric)
			if string(jsonMetric) == "{}" && string(data) != "{}" { // empty struct
				obj.Metric = nil
			} else {
				match++
			}
		} else {
			obj.Metric = nil
		}
	} else {
		obj.Metric = nil
	}

	// try to unmarshal data into MetricTagConfiguration
	err = datadog.Unmarshal(data, &obj.MetricTagConfiguration)
	if err == nil {
		if obj.MetricTagConfiguration != nil && obj.MetricTagConfiguration.UnparsedObject == nil {
			jsonMetricTagConfiguration, _ := datadog.Marshal(obj.MetricTagConfiguration)
			if string(jsonMetricTagConfiguration) == "{}" && string(data) != "{}" { // empty struct
				obj.MetricTagConfiguration = nil
			} else {
				match++
			}
		} else {
			obj.MetricTagConfiguration = nil
		}
	} else {
		obj.MetricTagConfiguration = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.Metric = nil
		obj.MetricTagConfiguration = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj MetricsAndMetricTagConfigurations) MarshalJSON() ([]byte, error) {
	if obj.Metric != nil {
		return datadog.Marshal(&obj.Metric)
	}

	if obj.MetricTagConfiguration != nil {
		return datadog.Marshal(&obj.MetricTagConfiguration)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *MetricsAndMetricTagConfigurations) GetActualInstance() interface{} {
	if obj.Metric != nil {
		return obj.Metric
	}

	if obj.MetricTagConfiguration != nil {
		return obj.MetricTagConfiguration
	}

	// all schemas are nil
	return nil
}
