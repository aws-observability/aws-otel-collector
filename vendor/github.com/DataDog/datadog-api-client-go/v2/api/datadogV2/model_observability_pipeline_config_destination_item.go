// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigDestinationItem - A destination for the pipeline.
type ObservabilityPipelineConfigDestinationItem struct {
	ObservabilityPipelineDatadogLogsDestination *ObservabilityPipelineDatadogLogsDestination

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineDatadogLogsDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineDatadogLogsDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineDatadogLogsDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineDatadogLogsDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineDatadogLogsDestination: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineConfigDestinationItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineDatadogLogsDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineDatadogLogsDestination)
	if err == nil {
		if obj.ObservabilityPipelineDatadogLogsDestination != nil && obj.ObservabilityPipelineDatadogLogsDestination.UnparsedObject == nil {
			jsonObservabilityPipelineDatadogLogsDestination, _ := datadog.Marshal(obj.ObservabilityPipelineDatadogLogsDestination)
			if string(jsonObservabilityPipelineDatadogLogsDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineDatadogLogsDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineDatadogLogsDestination = nil
		}
	} else {
		obj.ObservabilityPipelineDatadogLogsDestination = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineDatadogLogsDestination = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineConfigDestinationItem) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineDatadogLogsDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineDatadogLogsDestination)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineConfigDestinationItem) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineDatadogLogsDestination != nil {
		return obj.ObservabilityPipelineDatadogLogsDestination
	}

	// all schemas are nil
	return nil
}
