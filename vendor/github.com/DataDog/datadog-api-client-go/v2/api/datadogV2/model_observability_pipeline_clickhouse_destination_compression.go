// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineClickhouseDestinationCompression - Compression setting for outbound HTTP requests to ClickHouse.
// Can be specified as a shorthand string (`"gzip"` or `"none"`) or as an object
// with an `algorithm` field and an optional `level` (gzip only, 1–9).
type ObservabilityPipelineClickhouseDestinationCompression struct {
	ObservabilityPipelineClickhouseDestinationCompressionAlgorithm *ObservabilityPipelineClickhouseDestinationCompressionAlgorithm
	ObservabilityPipelineClickhouseDestinationCompressionObject    *ObservabilityPipelineClickhouseDestinationCompressionObject

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineClickhouseDestinationCompressionAlgorithmAsObservabilityPipelineClickhouseDestinationCompression is a convenience function that returns ObservabilityPipelineClickhouseDestinationCompressionAlgorithm wrapped in ObservabilityPipelineClickhouseDestinationCompression.
func ObservabilityPipelineClickhouseDestinationCompressionAlgorithmAsObservabilityPipelineClickhouseDestinationCompression(v *ObservabilityPipelineClickhouseDestinationCompressionAlgorithm) ObservabilityPipelineClickhouseDestinationCompression {
	return ObservabilityPipelineClickhouseDestinationCompression{ObservabilityPipelineClickhouseDestinationCompressionAlgorithm: v}
}

// ObservabilityPipelineClickhouseDestinationCompressionObjectAsObservabilityPipelineClickhouseDestinationCompression is a convenience function that returns ObservabilityPipelineClickhouseDestinationCompressionObject wrapped in ObservabilityPipelineClickhouseDestinationCompression.
func ObservabilityPipelineClickhouseDestinationCompressionObjectAsObservabilityPipelineClickhouseDestinationCompression(v *ObservabilityPipelineClickhouseDestinationCompressionObject) ObservabilityPipelineClickhouseDestinationCompression {
	return ObservabilityPipelineClickhouseDestinationCompression{ObservabilityPipelineClickhouseDestinationCompressionObject: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineClickhouseDestinationCompression) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineClickhouseDestinationCompressionAlgorithm
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineClickhouseDestinationCompressionAlgorithm)
	if err == nil {
		if obj.ObservabilityPipelineClickhouseDestinationCompressionAlgorithm != nil {
			jsonObservabilityPipelineClickhouseDestinationCompressionAlgorithm, _ := datadog.Marshal(obj.ObservabilityPipelineClickhouseDestinationCompressionAlgorithm)
			if string(jsonObservabilityPipelineClickhouseDestinationCompressionAlgorithm) == "{}" && string(data) != "{}" { // empty struct
				obj.ObservabilityPipelineClickhouseDestinationCompressionAlgorithm = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineClickhouseDestinationCompressionAlgorithm = nil
		}
	} else {
		obj.ObservabilityPipelineClickhouseDestinationCompressionAlgorithm = nil
	}

	// try to unmarshal data into ObservabilityPipelineClickhouseDestinationCompressionObject
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineClickhouseDestinationCompressionObject)
	if err == nil {
		if obj.ObservabilityPipelineClickhouseDestinationCompressionObject != nil && obj.ObservabilityPipelineClickhouseDestinationCompressionObject.UnparsedObject == nil {
			jsonObservabilityPipelineClickhouseDestinationCompressionObject, _ := datadog.Marshal(obj.ObservabilityPipelineClickhouseDestinationCompressionObject)
			if string(jsonObservabilityPipelineClickhouseDestinationCompressionObject) == "{}" { // empty struct
				obj.ObservabilityPipelineClickhouseDestinationCompressionObject = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineClickhouseDestinationCompressionObject = nil
		}
	} else {
		obj.ObservabilityPipelineClickhouseDestinationCompressionObject = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineClickhouseDestinationCompressionAlgorithm = nil
		obj.ObservabilityPipelineClickhouseDestinationCompressionObject = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineClickhouseDestinationCompression) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineClickhouseDestinationCompressionAlgorithm != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineClickhouseDestinationCompressionAlgorithm)
	}

	if obj.ObservabilityPipelineClickhouseDestinationCompressionObject != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineClickhouseDestinationCompressionObject)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineClickhouseDestinationCompression) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineClickhouseDestinationCompressionAlgorithm != nil {
		return obj.ObservabilityPipelineClickhouseDestinationCompressionAlgorithm
	}

	if obj.ObservabilityPipelineClickhouseDestinationCompressionObject != nil {
		return obj.ObservabilityPipelineClickhouseDestinationCompressionObject
	}

	// all schemas are nil
	return nil
}
