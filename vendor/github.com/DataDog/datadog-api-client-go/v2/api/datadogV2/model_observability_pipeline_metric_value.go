// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMetricValue - Specifies how the value of the generated metric is computed.
type ObservabilityPipelineMetricValue struct {
	ObservabilityPipelineGeneratedMetricIncrementByOne   *ObservabilityPipelineGeneratedMetricIncrementByOne
	ObservabilityPipelineGeneratedMetricIncrementByField *ObservabilityPipelineGeneratedMetricIncrementByField

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineGeneratedMetricIncrementByOneAsObservabilityPipelineMetricValue is a convenience function that returns ObservabilityPipelineGeneratedMetricIncrementByOne wrapped in ObservabilityPipelineMetricValue.
func ObservabilityPipelineGeneratedMetricIncrementByOneAsObservabilityPipelineMetricValue(v *ObservabilityPipelineGeneratedMetricIncrementByOne) ObservabilityPipelineMetricValue {
	return ObservabilityPipelineMetricValue{ObservabilityPipelineGeneratedMetricIncrementByOne: v}
}

// ObservabilityPipelineGeneratedMetricIncrementByFieldAsObservabilityPipelineMetricValue is a convenience function that returns ObservabilityPipelineGeneratedMetricIncrementByField wrapped in ObservabilityPipelineMetricValue.
func ObservabilityPipelineGeneratedMetricIncrementByFieldAsObservabilityPipelineMetricValue(v *ObservabilityPipelineGeneratedMetricIncrementByField) ObservabilityPipelineMetricValue {
	return ObservabilityPipelineMetricValue{ObservabilityPipelineGeneratedMetricIncrementByField: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineMetricValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineGeneratedMetricIncrementByOne
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineGeneratedMetricIncrementByOne)
	if err == nil {
		if obj.ObservabilityPipelineGeneratedMetricIncrementByOne != nil && obj.ObservabilityPipelineGeneratedMetricIncrementByOne.UnparsedObject == nil {
			jsonObservabilityPipelineGeneratedMetricIncrementByOne, _ := datadog.Marshal(obj.ObservabilityPipelineGeneratedMetricIncrementByOne)
			if string(jsonObservabilityPipelineGeneratedMetricIncrementByOne) == "{}" { // empty struct
				obj.ObservabilityPipelineGeneratedMetricIncrementByOne = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineGeneratedMetricIncrementByOne = nil
		}
	} else {
		obj.ObservabilityPipelineGeneratedMetricIncrementByOne = nil
	}

	// try to unmarshal data into ObservabilityPipelineGeneratedMetricIncrementByField
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineGeneratedMetricIncrementByField)
	if err == nil {
		if obj.ObservabilityPipelineGeneratedMetricIncrementByField != nil && obj.ObservabilityPipelineGeneratedMetricIncrementByField.UnparsedObject == nil {
			jsonObservabilityPipelineGeneratedMetricIncrementByField, _ := datadog.Marshal(obj.ObservabilityPipelineGeneratedMetricIncrementByField)
			if string(jsonObservabilityPipelineGeneratedMetricIncrementByField) == "{}" { // empty struct
				obj.ObservabilityPipelineGeneratedMetricIncrementByField = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineGeneratedMetricIncrementByField = nil
		}
	} else {
		obj.ObservabilityPipelineGeneratedMetricIncrementByField = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineGeneratedMetricIncrementByOne = nil
		obj.ObservabilityPipelineGeneratedMetricIncrementByField = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineMetricValue) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineGeneratedMetricIncrementByOne != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineGeneratedMetricIncrementByOne)
	}

	if obj.ObservabilityPipelineGeneratedMetricIncrementByField != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineGeneratedMetricIncrementByField)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineMetricValue) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineGeneratedMetricIncrementByOne != nil {
		return obj.ObservabilityPipelineGeneratedMetricIncrementByOne
	}

	if obj.ObservabilityPipelineGeneratedMetricIncrementByField != nil {
		return obj.ObservabilityPipelineGeneratedMetricIncrementByField
	}

	// all schemas are nil
	return nil
}
