// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOcsfMapperProcessorMappingMapping - Defines a single mapping rule for transforming logs into the OCSF schema.
type ObservabilityPipelineOcsfMapperProcessorMappingMapping struct {
	ObservabilityPipelineOcsfMappingLibrary *ObservabilityPipelineOcsfMappingLibrary

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineOcsfMappingLibraryAsObservabilityPipelineOcsfMapperProcessorMappingMapping is a convenience function that returns ObservabilityPipelineOcsfMappingLibrary wrapped in ObservabilityPipelineOcsfMapperProcessorMappingMapping.
func ObservabilityPipelineOcsfMappingLibraryAsObservabilityPipelineOcsfMapperProcessorMappingMapping(v *ObservabilityPipelineOcsfMappingLibrary) ObservabilityPipelineOcsfMapperProcessorMappingMapping {
	return ObservabilityPipelineOcsfMapperProcessorMappingMapping{ObservabilityPipelineOcsfMappingLibrary: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineOcsfMapperProcessorMappingMapping) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineOcsfMappingLibrary
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineOcsfMappingLibrary)
	if err == nil {
		if obj.ObservabilityPipelineOcsfMappingLibrary != nil {
			jsonObservabilityPipelineOcsfMappingLibrary, _ := datadog.Marshal(obj.ObservabilityPipelineOcsfMappingLibrary)
			if string(jsonObservabilityPipelineOcsfMappingLibrary) == "{}" && string(data) != "{}" { // empty struct
				obj.ObservabilityPipelineOcsfMappingLibrary = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineOcsfMappingLibrary = nil
		}
	} else {
		obj.ObservabilityPipelineOcsfMappingLibrary = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineOcsfMappingLibrary = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineOcsfMapperProcessorMappingMapping) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineOcsfMappingLibrary != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineOcsfMappingLibrary)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineOcsfMapperProcessorMappingMapping) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineOcsfMappingLibrary != nil {
		return obj.ObservabilityPipelineOcsfMappingLibrary
	}

	// all schemas are nil
	return nil
}
