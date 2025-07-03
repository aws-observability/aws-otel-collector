// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigSourceItem - A data source for the pipeline.
type ObservabilityPipelineConfigSourceItem struct {
	ObservabilityPipelineKafkaSource        *ObservabilityPipelineKafkaSource
	ObservabilityPipelineDatadogAgentSource *ObservabilityPipelineDatadogAgentSource

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineKafkaSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineKafkaSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineKafkaSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineKafkaSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineKafkaSource: v}
}

// ObservabilityPipelineDatadogAgentSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineDatadogAgentSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineDatadogAgentSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineDatadogAgentSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineDatadogAgentSource: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineConfigSourceItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineKafkaSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineKafkaSource)
	if err == nil {
		if obj.ObservabilityPipelineKafkaSource != nil && obj.ObservabilityPipelineKafkaSource.UnparsedObject == nil {
			jsonObservabilityPipelineKafkaSource, _ := datadog.Marshal(obj.ObservabilityPipelineKafkaSource)
			if string(jsonObservabilityPipelineKafkaSource) == "{}" { // empty struct
				obj.ObservabilityPipelineKafkaSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineKafkaSource = nil
		}
	} else {
		obj.ObservabilityPipelineKafkaSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineDatadogAgentSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineDatadogAgentSource)
	if err == nil {
		if obj.ObservabilityPipelineDatadogAgentSource != nil && obj.ObservabilityPipelineDatadogAgentSource.UnparsedObject == nil {
			jsonObservabilityPipelineDatadogAgentSource, _ := datadog.Marshal(obj.ObservabilityPipelineDatadogAgentSource)
			if string(jsonObservabilityPipelineDatadogAgentSource) == "{}" { // empty struct
				obj.ObservabilityPipelineDatadogAgentSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineDatadogAgentSource = nil
		}
	} else {
		obj.ObservabilityPipelineDatadogAgentSource = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineKafkaSource = nil
		obj.ObservabilityPipelineDatadogAgentSource = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineConfigSourceItem) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineKafkaSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineKafkaSource)
	}

	if obj.ObservabilityPipelineDatadogAgentSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineDatadogAgentSource)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineConfigSourceItem) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineKafkaSource != nil {
		return obj.ObservabilityPipelineKafkaSource
	}

	if obj.ObservabilityPipelineDatadogAgentSource != nil {
		return obj.ObservabilityPipelineDatadogAgentSource
	}

	// all schemas are nil
	return nil
}
