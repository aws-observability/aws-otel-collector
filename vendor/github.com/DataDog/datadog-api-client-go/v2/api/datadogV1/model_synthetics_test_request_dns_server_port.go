// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestRequestDNSServerPort - DNS server port to use for DNS tests.
type SyntheticsTestRequestDNSServerPort struct {
	SyntheticsTestRequestNumericalDNSServerPort *int64
	SyntheticsTestRequestVariableDNSServerPort  *string

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SyntheticsTestRequestNumericalDNSServerPortAsSyntheticsTestRequestDNSServerPort is a convenience function that returns int64 wrapped in SyntheticsTestRequestDNSServerPort.
func SyntheticsTestRequestNumericalDNSServerPortAsSyntheticsTestRequestDNSServerPort(v *int64) SyntheticsTestRequestDNSServerPort {
	return SyntheticsTestRequestDNSServerPort{SyntheticsTestRequestNumericalDNSServerPort: v}
}

// SyntheticsTestRequestVariableDNSServerPortAsSyntheticsTestRequestDNSServerPort is a convenience function that returns string wrapped in SyntheticsTestRequestDNSServerPort.
func SyntheticsTestRequestVariableDNSServerPortAsSyntheticsTestRequestDNSServerPort(v *string) SyntheticsTestRequestDNSServerPort {
	return SyntheticsTestRequestDNSServerPort{SyntheticsTestRequestVariableDNSServerPort: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SyntheticsTestRequestDNSServerPort) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SyntheticsTestRequestNumericalDNSServerPort
	err = datadog.Unmarshal(data, &obj.SyntheticsTestRequestNumericalDNSServerPort)
	if err == nil {
		if obj.SyntheticsTestRequestNumericalDNSServerPort != nil {
			jsonSyntheticsTestRequestNumericalDNSServerPort, _ := datadog.Marshal(obj.SyntheticsTestRequestNumericalDNSServerPort)
			if string(jsonSyntheticsTestRequestNumericalDNSServerPort) == "{}" { // empty struct
				obj.SyntheticsTestRequestNumericalDNSServerPort = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsTestRequestNumericalDNSServerPort = nil
		}
	} else {
		obj.SyntheticsTestRequestNumericalDNSServerPort = nil
	}

	// try to unmarshal data into SyntheticsTestRequestVariableDNSServerPort
	err = datadog.Unmarshal(data, &obj.SyntheticsTestRequestVariableDNSServerPort)
	if err == nil {
		if obj.SyntheticsTestRequestVariableDNSServerPort != nil {
			jsonSyntheticsTestRequestVariableDNSServerPort, _ := datadog.Marshal(obj.SyntheticsTestRequestVariableDNSServerPort)
			if string(jsonSyntheticsTestRequestVariableDNSServerPort) == "{}" { // empty struct
				obj.SyntheticsTestRequestVariableDNSServerPort = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsTestRequestVariableDNSServerPort = nil
		}
	} else {
		obj.SyntheticsTestRequestVariableDNSServerPort = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SyntheticsTestRequestNumericalDNSServerPort = nil
		obj.SyntheticsTestRequestVariableDNSServerPort = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SyntheticsTestRequestDNSServerPort) MarshalJSON() ([]byte, error) {
	if obj.SyntheticsTestRequestNumericalDNSServerPort != nil {
		return datadog.Marshal(&obj.SyntheticsTestRequestNumericalDNSServerPort)
	}

	if obj.SyntheticsTestRequestVariableDNSServerPort != nil {
		return datadog.Marshal(&obj.SyntheticsTestRequestVariableDNSServerPort)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SyntheticsTestRequestDNSServerPort) GetActualInstance() interface{} {
	if obj.SyntheticsTestRequestNumericalDNSServerPort != nil {
		return obj.SyntheticsTestRequestNumericalDNSServerPort
	}

	if obj.SyntheticsTestRequestVariableDNSServerPort != nil {
		return obj.SyntheticsTestRequestVariableDNSServerPort
	}

	// all schemas are nil
	return nil
}
