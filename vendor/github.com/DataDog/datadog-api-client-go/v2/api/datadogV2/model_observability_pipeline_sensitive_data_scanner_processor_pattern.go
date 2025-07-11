// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorPattern - Pattern detection configuration for identifying sensitive data using either a custom regex or a library reference.
type ObservabilityPipelineSensitiveDataScannerProcessorPattern struct {
	ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern  *ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern
	ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternAsObservabilityPipelineSensitiveDataScannerProcessorPattern is a convenience function that returns ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern wrapped in ObservabilityPipelineSensitiveDataScannerProcessorPattern.
func ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternAsObservabilityPipelineSensitiveDataScannerProcessorPattern(v *ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern) ObservabilityPipelineSensitiveDataScannerProcessorPattern {
	return ObservabilityPipelineSensitiveDataScannerProcessorPattern{ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern: v}
}

// ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternAsObservabilityPipelineSensitiveDataScannerProcessorPattern is a convenience function that returns ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern wrapped in ObservabilityPipelineSensitiveDataScannerProcessorPattern.
func ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternAsObservabilityPipelineSensitiveDataScannerProcessorPattern(v *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern) ObservabilityPipelineSensitiveDataScannerProcessorPattern {
	return ObservabilityPipelineSensitiveDataScannerProcessorPattern{ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineSensitiveDataScannerProcessorPattern) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern)
	if err == nil {
		if obj.ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern != nil && obj.ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern.UnparsedObject == nil {
			jsonObservabilityPipelineSensitiveDataScannerProcessorCustomPattern, _ := datadog.Marshal(obj.ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern)
			if string(jsonObservabilityPipelineSensitiveDataScannerProcessorCustomPattern) == "{}" { // empty struct
				obj.ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern = nil
		}
	} else {
		obj.ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern = nil
	}

	// try to unmarshal data into ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern)
	if err == nil {
		if obj.ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern != nil && obj.ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern.UnparsedObject == nil {
			jsonObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern, _ := datadog.Marshal(obj.ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern)
			if string(jsonObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern) == "{}" { // empty struct
				obj.ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern = nil
		}
	} else {
		obj.ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern = nil
		obj.ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineSensitiveDataScannerProcessorPattern) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern)
	}

	if obj.ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineSensitiveDataScannerProcessorPattern) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern != nil {
		return obj.ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern
	}

	if obj.ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern != nil {
		return obj.ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern
	}

	// all schemas are nil
	return nil
}
