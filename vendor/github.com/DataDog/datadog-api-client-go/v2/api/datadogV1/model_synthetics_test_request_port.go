// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestRequestPort - Port to use when performing the test.
type SyntheticsTestRequestPort struct {
	SyntheticsTestRequestNumericalPort *int64
	SyntheticsTestRequestVariablePort  *string

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SyntheticsTestRequestNumericalPortAsSyntheticsTestRequestPort is a convenience function that returns int64 wrapped in SyntheticsTestRequestPort.
func SyntheticsTestRequestNumericalPortAsSyntheticsTestRequestPort(v *int64) SyntheticsTestRequestPort {
	return SyntheticsTestRequestPort{SyntheticsTestRequestNumericalPort: v}
}

// SyntheticsTestRequestVariablePortAsSyntheticsTestRequestPort is a convenience function that returns string wrapped in SyntheticsTestRequestPort.
func SyntheticsTestRequestVariablePortAsSyntheticsTestRequestPort(v *string) SyntheticsTestRequestPort {
	return SyntheticsTestRequestPort{SyntheticsTestRequestVariablePort: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SyntheticsTestRequestPort) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SyntheticsTestRequestNumericalPort
	err = datadog.Unmarshal(data, &obj.SyntheticsTestRequestNumericalPort)
	if err == nil {
		if obj.SyntheticsTestRequestNumericalPort != nil {
			jsonSyntheticsTestRequestNumericalPort, _ := datadog.Marshal(obj.SyntheticsTestRequestNumericalPort)
			if string(jsonSyntheticsTestRequestNumericalPort) == "{}" { // empty struct
				obj.SyntheticsTestRequestNumericalPort = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsTestRequestNumericalPort = nil
		}
	} else {
		obj.SyntheticsTestRequestNumericalPort = nil
	}

	// try to unmarshal data into SyntheticsTestRequestVariablePort
	err = datadog.Unmarshal(data, &obj.SyntheticsTestRequestVariablePort)
	if err == nil {
		if obj.SyntheticsTestRequestVariablePort != nil {
			jsonSyntheticsTestRequestVariablePort, _ := datadog.Marshal(obj.SyntheticsTestRequestVariablePort)
			if string(jsonSyntheticsTestRequestVariablePort) == "{}" { // empty struct
				obj.SyntheticsTestRequestVariablePort = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsTestRequestVariablePort = nil
		}
	} else {
		obj.SyntheticsTestRequestVariablePort = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SyntheticsTestRequestNumericalPort = nil
		obj.SyntheticsTestRequestVariablePort = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SyntheticsTestRequestPort) MarshalJSON() ([]byte, error) {
	if obj.SyntheticsTestRequestNumericalPort != nil {
		return datadog.Marshal(&obj.SyntheticsTestRequestNumericalPort)
	}

	if obj.SyntheticsTestRequestVariablePort != nil {
		return datadog.Marshal(&obj.SyntheticsTestRequestVariablePort)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SyntheticsTestRequestPort) GetActualInstance() interface{} {
	if obj.SyntheticsTestRequestNumericalPort != nil {
		return obj.SyntheticsTestRequestNumericalPort
	}

	if obj.SyntheticsTestRequestVariablePort != nil {
		return obj.SyntheticsTestRequestVariablePort
	}

	// all schemas are nil
	return nil
}
