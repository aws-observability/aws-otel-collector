// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseHttpDestinationAuth - Authentication method of the HTTP requests.
type CustomDestinationResponseHttpDestinationAuth struct {
	CustomDestinationResponseHttpDestinationAuthBasic        *CustomDestinationResponseHttpDestinationAuthBasic
	CustomDestinationResponseHttpDestinationAuthCustomHeader *CustomDestinationResponseHttpDestinationAuthCustomHeader

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CustomDestinationResponseHttpDestinationAuthBasicAsCustomDestinationResponseHttpDestinationAuth is a convenience function that returns CustomDestinationResponseHttpDestinationAuthBasic wrapped in CustomDestinationResponseHttpDestinationAuth.
func CustomDestinationResponseHttpDestinationAuthBasicAsCustomDestinationResponseHttpDestinationAuth(v *CustomDestinationResponseHttpDestinationAuthBasic) CustomDestinationResponseHttpDestinationAuth {
	return CustomDestinationResponseHttpDestinationAuth{CustomDestinationResponseHttpDestinationAuthBasic: v}
}

// CustomDestinationResponseHttpDestinationAuthCustomHeaderAsCustomDestinationResponseHttpDestinationAuth is a convenience function that returns CustomDestinationResponseHttpDestinationAuthCustomHeader wrapped in CustomDestinationResponseHttpDestinationAuth.
func CustomDestinationResponseHttpDestinationAuthCustomHeaderAsCustomDestinationResponseHttpDestinationAuth(v *CustomDestinationResponseHttpDestinationAuthCustomHeader) CustomDestinationResponseHttpDestinationAuth {
	return CustomDestinationResponseHttpDestinationAuth{CustomDestinationResponseHttpDestinationAuthCustomHeader: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CustomDestinationResponseHttpDestinationAuth) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomDestinationResponseHttpDestinationAuthBasic
	err = datadog.Unmarshal(data, &obj.CustomDestinationResponseHttpDestinationAuthBasic)
	if err == nil {
		if obj.CustomDestinationResponseHttpDestinationAuthBasic != nil && obj.CustomDestinationResponseHttpDestinationAuthBasic.UnparsedObject == nil {
			jsonCustomDestinationResponseHttpDestinationAuthBasic, _ := datadog.Marshal(obj.CustomDestinationResponseHttpDestinationAuthBasic)
			if string(jsonCustomDestinationResponseHttpDestinationAuthBasic) == "{}" { // empty struct
				obj.CustomDestinationResponseHttpDestinationAuthBasic = nil
			} else {
				match++
			}
		} else {
			obj.CustomDestinationResponseHttpDestinationAuthBasic = nil
		}
	} else {
		obj.CustomDestinationResponseHttpDestinationAuthBasic = nil
	}

	// try to unmarshal data into CustomDestinationResponseHttpDestinationAuthCustomHeader
	err = datadog.Unmarshal(data, &obj.CustomDestinationResponseHttpDestinationAuthCustomHeader)
	if err == nil {
		if obj.CustomDestinationResponseHttpDestinationAuthCustomHeader != nil && obj.CustomDestinationResponseHttpDestinationAuthCustomHeader.UnparsedObject == nil {
			jsonCustomDestinationResponseHttpDestinationAuthCustomHeader, _ := datadog.Marshal(obj.CustomDestinationResponseHttpDestinationAuthCustomHeader)
			if string(jsonCustomDestinationResponseHttpDestinationAuthCustomHeader) == "{}" { // empty struct
				obj.CustomDestinationResponseHttpDestinationAuthCustomHeader = nil
			} else {
				match++
			}
		} else {
			obj.CustomDestinationResponseHttpDestinationAuthCustomHeader = nil
		}
	} else {
		obj.CustomDestinationResponseHttpDestinationAuthCustomHeader = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CustomDestinationResponseHttpDestinationAuthBasic = nil
		obj.CustomDestinationResponseHttpDestinationAuthCustomHeader = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CustomDestinationResponseHttpDestinationAuth) MarshalJSON() ([]byte, error) {
	if obj.CustomDestinationResponseHttpDestinationAuthBasic != nil {
		return datadog.Marshal(&obj.CustomDestinationResponseHttpDestinationAuthBasic)
	}

	if obj.CustomDestinationResponseHttpDestinationAuthCustomHeader != nil {
		return datadog.Marshal(&obj.CustomDestinationResponseHttpDestinationAuthCustomHeader)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CustomDestinationResponseHttpDestinationAuth) GetActualInstance() interface{} {
	if obj.CustomDestinationResponseHttpDestinationAuthBasic != nil {
		return obj.CustomDestinationResponseHttpDestinationAuthBasic
	}

	if obj.CustomDestinationResponseHttpDestinationAuthCustomHeader != nil {
		return obj.CustomDestinationResponseHttpDestinationAuthCustomHeader
	}

	// all schemas are nil
	return nil
}
