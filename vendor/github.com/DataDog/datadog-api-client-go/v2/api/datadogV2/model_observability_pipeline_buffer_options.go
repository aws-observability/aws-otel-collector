// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineBufferOptions - Configuration for buffer settings on destination components.
type ObservabilityPipelineBufferOptions struct {
	ObservabilityPipelineDiskBufferOptions       *ObservabilityPipelineDiskBufferOptions
	ObservabilityPipelineMemoryBufferOptions     *ObservabilityPipelineMemoryBufferOptions
	ObservabilityPipelineMemoryBufferSizeOptions *ObservabilityPipelineMemoryBufferSizeOptions

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineDiskBufferOptionsAsObservabilityPipelineBufferOptions is a convenience function that returns ObservabilityPipelineDiskBufferOptions wrapped in ObservabilityPipelineBufferOptions.
func ObservabilityPipelineDiskBufferOptionsAsObservabilityPipelineBufferOptions(v *ObservabilityPipelineDiskBufferOptions) ObservabilityPipelineBufferOptions {
	return ObservabilityPipelineBufferOptions{ObservabilityPipelineDiskBufferOptions: v}
}

// ObservabilityPipelineMemoryBufferOptionsAsObservabilityPipelineBufferOptions is a convenience function that returns ObservabilityPipelineMemoryBufferOptions wrapped in ObservabilityPipelineBufferOptions.
func ObservabilityPipelineMemoryBufferOptionsAsObservabilityPipelineBufferOptions(v *ObservabilityPipelineMemoryBufferOptions) ObservabilityPipelineBufferOptions {
	return ObservabilityPipelineBufferOptions{ObservabilityPipelineMemoryBufferOptions: v}
}

// ObservabilityPipelineMemoryBufferSizeOptionsAsObservabilityPipelineBufferOptions is a convenience function that returns ObservabilityPipelineMemoryBufferSizeOptions wrapped in ObservabilityPipelineBufferOptions.
func ObservabilityPipelineMemoryBufferSizeOptionsAsObservabilityPipelineBufferOptions(v *ObservabilityPipelineMemoryBufferSizeOptions) ObservabilityPipelineBufferOptions {
	return ObservabilityPipelineBufferOptions{ObservabilityPipelineMemoryBufferSizeOptions: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineBufferOptions) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineDiskBufferOptions
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineDiskBufferOptions)
	if err == nil {
		if obj.ObservabilityPipelineDiskBufferOptions != nil && obj.ObservabilityPipelineDiskBufferOptions.UnparsedObject == nil {
			jsonObservabilityPipelineDiskBufferOptions, _ := datadog.Marshal(obj.ObservabilityPipelineDiskBufferOptions)
			if string(jsonObservabilityPipelineDiskBufferOptions) == "{}" { // empty struct
				obj.ObservabilityPipelineDiskBufferOptions = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineDiskBufferOptions = nil
		}
	} else {
		obj.ObservabilityPipelineDiskBufferOptions = nil
	}

	// try to unmarshal data into ObservabilityPipelineMemoryBufferOptions
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineMemoryBufferOptions)
	if err == nil {
		if obj.ObservabilityPipelineMemoryBufferOptions != nil && obj.ObservabilityPipelineMemoryBufferOptions.UnparsedObject == nil {
			jsonObservabilityPipelineMemoryBufferOptions, _ := datadog.Marshal(obj.ObservabilityPipelineMemoryBufferOptions)
			if string(jsonObservabilityPipelineMemoryBufferOptions) == "{}" { // empty struct
				obj.ObservabilityPipelineMemoryBufferOptions = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineMemoryBufferOptions = nil
		}
	} else {
		obj.ObservabilityPipelineMemoryBufferOptions = nil
	}

	// try to unmarshal data into ObservabilityPipelineMemoryBufferSizeOptions
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineMemoryBufferSizeOptions)
	if err == nil {
		if obj.ObservabilityPipelineMemoryBufferSizeOptions != nil && obj.ObservabilityPipelineMemoryBufferSizeOptions.UnparsedObject == nil {
			jsonObservabilityPipelineMemoryBufferSizeOptions, _ := datadog.Marshal(obj.ObservabilityPipelineMemoryBufferSizeOptions)
			if string(jsonObservabilityPipelineMemoryBufferSizeOptions) == "{}" { // empty struct
				obj.ObservabilityPipelineMemoryBufferSizeOptions = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineMemoryBufferSizeOptions = nil
		}
	} else {
		obj.ObservabilityPipelineMemoryBufferSizeOptions = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineDiskBufferOptions = nil
		obj.ObservabilityPipelineMemoryBufferOptions = nil
		obj.ObservabilityPipelineMemoryBufferSizeOptions = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineBufferOptions) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineDiskBufferOptions != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineDiskBufferOptions)
	}

	if obj.ObservabilityPipelineMemoryBufferOptions != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineMemoryBufferOptions)
	}

	if obj.ObservabilityPipelineMemoryBufferSizeOptions != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineMemoryBufferSizeOptions)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineBufferOptions) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineDiskBufferOptions != nil {
		return obj.ObservabilityPipelineDiskBufferOptions
	}

	if obj.ObservabilityPipelineMemoryBufferOptions != nil {
		return obj.ObservabilityPipelineMemoryBufferOptions
	}

	if obj.ObservabilityPipelineMemoryBufferSizeOptions != nil {
		return obj.ObservabilityPipelineMemoryBufferSizeOptions
	}

	// all schemas are nil
	return nil
}
