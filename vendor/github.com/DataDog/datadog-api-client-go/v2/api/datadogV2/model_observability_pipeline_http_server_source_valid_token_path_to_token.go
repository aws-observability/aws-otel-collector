// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpServerSourceValidTokenPathToToken - Specifies where the worker extracts the token from in the incoming HTTP request.
// This can be either a built-in location (`path` or `address`) or an HTTP header object.
type ObservabilityPipelineHttpServerSourceValidTokenPathToToken struct {
	ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation
	ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader   *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocationAsObservabilityPipelineHttpServerSourceValidTokenPathToToken is a convenience function that returns ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation wrapped in ObservabilityPipelineHttpServerSourceValidTokenPathToToken.
func ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocationAsObservabilityPipelineHttpServerSourceValidTokenPathToToken(v *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation) ObservabilityPipelineHttpServerSourceValidTokenPathToToken {
	return ObservabilityPipelineHttpServerSourceValidTokenPathToToken{ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation: v}
}

// ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeaderAsObservabilityPipelineHttpServerSourceValidTokenPathToToken is a convenience function that returns ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader wrapped in ObservabilityPipelineHttpServerSourceValidTokenPathToToken.
func ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeaderAsObservabilityPipelineHttpServerSourceValidTokenPathToToken(v *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader) ObservabilityPipelineHttpServerSourceValidTokenPathToToken {
	return ObservabilityPipelineHttpServerSourceValidTokenPathToToken{ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineHttpServerSourceValidTokenPathToToken) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation)
	if err == nil {
		if obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation != nil {
			jsonObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation, _ := datadog.Marshal(obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation)
			if string(jsonObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation) == "{}" && string(data) != "{}" { // empty struct
				obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation = nil
		}
	} else {
		obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation = nil
	}

	// try to unmarshal data into ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader)
	if err == nil {
		if obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader != nil && obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader.UnparsedObject == nil {
			jsonObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader, _ := datadog.Marshal(obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader)
			if string(jsonObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader) == "{}" { // empty struct
				obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader = nil
		}
	} else {
		obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation = nil
		obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineHttpServerSourceValidTokenPathToToken) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation)
	}

	if obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineHttpServerSourceValidTokenPathToToken) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation != nil {
		return obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation
	}

	if obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader != nil {
		return obj.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader
	}

	// all schemas are nil
	return nil
}
