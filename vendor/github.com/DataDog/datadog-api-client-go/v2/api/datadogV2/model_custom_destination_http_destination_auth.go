// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationHttpDestinationAuth - Authentication method of the HTTP requests.
type CustomDestinationHttpDestinationAuth struct {
	CustomDestinationHttpDestinationAuthBasic        *CustomDestinationHttpDestinationAuthBasic
	CustomDestinationHttpDestinationAuthCustomHeader *CustomDestinationHttpDestinationAuthCustomHeader

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CustomDestinationHttpDestinationAuthBasicAsCustomDestinationHttpDestinationAuth is a convenience function that returns CustomDestinationHttpDestinationAuthBasic wrapped in CustomDestinationHttpDestinationAuth.
func CustomDestinationHttpDestinationAuthBasicAsCustomDestinationHttpDestinationAuth(v *CustomDestinationHttpDestinationAuthBasic) CustomDestinationHttpDestinationAuth {
	return CustomDestinationHttpDestinationAuth{CustomDestinationHttpDestinationAuthBasic: v}
}

// CustomDestinationHttpDestinationAuthCustomHeaderAsCustomDestinationHttpDestinationAuth is a convenience function that returns CustomDestinationHttpDestinationAuthCustomHeader wrapped in CustomDestinationHttpDestinationAuth.
func CustomDestinationHttpDestinationAuthCustomHeaderAsCustomDestinationHttpDestinationAuth(v *CustomDestinationHttpDestinationAuthCustomHeader) CustomDestinationHttpDestinationAuth {
	return CustomDestinationHttpDestinationAuth{CustomDestinationHttpDestinationAuthCustomHeader: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CustomDestinationHttpDestinationAuth) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomDestinationHttpDestinationAuthBasic
	err = datadog.Unmarshal(data, &obj.CustomDestinationHttpDestinationAuthBasic)
	if err == nil {
		if obj.CustomDestinationHttpDestinationAuthBasic != nil && obj.CustomDestinationHttpDestinationAuthBasic.UnparsedObject == nil {
			jsonCustomDestinationHttpDestinationAuthBasic, _ := datadog.Marshal(obj.CustomDestinationHttpDestinationAuthBasic)
			if string(jsonCustomDestinationHttpDestinationAuthBasic) == "{}" { // empty struct
				obj.CustomDestinationHttpDestinationAuthBasic = nil
			} else {
				match++
			}
		} else {
			obj.CustomDestinationHttpDestinationAuthBasic = nil
		}
	} else {
		obj.CustomDestinationHttpDestinationAuthBasic = nil
	}

	// try to unmarshal data into CustomDestinationHttpDestinationAuthCustomHeader
	err = datadog.Unmarshal(data, &obj.CustomDestinationHttpDestinationAuthCustomHeader)
	if err == nil {
		if obj.CustomDestinationHttpDestinationAuthCustomHeader != nil && obj.CustomDestinationHttpDestinationAuthCustomHeader.UnparsedObject == nil {
			jsonCustomDestinationHttpDestinationAuthCustomHeader, _ := datadog.Marshal(obj.CustomDestinationHttpDestinationAuthCustomHeader)
			if string(jsonCustomDestinationHttpDestinationAuthCustomHeader) == "{}" { // empty struct
				obj.CustomDestinationHttpDestinationAuthCustomHeader = nil
			} else {
				match++
			}
		} else {
			obj.CustomDestinationHttpDestinationAuthCustomHeader = nil
		}
	} else {
		obj.CustomDestinationHttpDestinationAuthCustomHeader = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CustomDestinationHttpDestinationAuthBasic = nil
		obj.CustomDestinationHttpDestinationAuthCustomHeader = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CustomDestinationHttpDestinationAuth) MarshalJSON() ([]byte, error) {
	if obj.CustomDestinationHttpDestinationAuthBasic != nil {
		return datadog.Marshal(&obj.CustomDestinationHttpDestinationAuthBasic)
	}

	if obj.CustomDestinationHttpDestinationAuthCustomHeader != nil {
		return datadog.Marshal(&obj.CustomDestinationHttpDestinationAuthCustomHeader)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CustomDestinationHttpDestinationAuth) GetActualInstance() interface{} {
	if obj.CustomDestinationHttpDestinationAuthBasic != nil {
		return obj.CustomDestinationHttpDestinationAuthBasic
	}

	if obj.CustomDestinationHttpDestinationAuthCustomHeader != nil {
		return obj.CustomDestinationHttpDestinationAuthCustomHeader
	}

	// all schemas are nil
	return nil
}
