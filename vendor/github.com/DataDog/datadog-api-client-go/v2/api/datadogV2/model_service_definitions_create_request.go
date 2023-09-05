// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// ServiceDefinitionsCreateRequest - Create service definitions request.
type ServiceDefinitionsCreateRequest struct {
	ServiceDefinitionV2Dot1 *ServiceDefinitionV2Dot1
	ServiceDefinitionV2     *ServiceDefinitionV2
	ServiceDefinitionRaw    *string

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ServiceDefinitionV2Dot1AsServiceDefinitionsCreateRequest is a convenience function that returns ServiceDefinitionV2Dot1 wrapped in ServiceDefinitionsCreateRequest.
func ServiceDefinitionV2Dot1AsServiceDefinitionsCreateRequest(v *ServiceDefinitionV2Dot1) ServiceDefinitionsCreateRequest {
	return ServiceDefinitionsCreateRequest{ServiceDefinitionV2Dot1: v}
}

// ServiceDefinitionV2AsServiceDefinitionsCreateRequest is a convenience function that returns ServiceDefinitionV2 wrapped in ServiceDefinitionsCreateRequest.
func ServiceDefinitionV2AsServiceDefinitionsCreateRequest(v *ServiceDefinitionV2) ServiceDefinitionsCreateRequest {
	return ServiceDefinitionsCreateRequest{ServiceDefinitionV2: v}
}

// ServiceDefinitionRawAsServiceDefinitionsCreateRequest is a convenience function that returns string wrapped in ServiceDefinitionsCreateRequest.
func ServiceDefinitionRawAsServiceDefinitionsCreateRequest(v *string) ServiceDefinitionsCreateRequest {
	return ServiceDefinitionsCreateRequest{ServiceDefinitionRaw: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ServiceDefinitionsCreateRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ServiceDefinitionV2Dot1
	err = json.Unmarshal(data, &obj.ServiceDefinitionV2Dot1)
	if err == nil {
		if obj.ServiceDefinitionV2Dot1 != nil && obj.ServiceDefinitionV2Dot1.UnparsedObject == nil {
			jsonServiceDefinitionV2Dot1, _ := json.Marshal(obj.ServiceDefinitionV2Dot1)
			if string(jsonServiceDefinitionV2Dot1) == "{}" { // empty struct
				obj.ServiceDefinitionV2Dot1 = nil
			} else {
				match++
			}
		} else {
			obj.ServiceDefinitionV2Dot1 = nil
		}
	} else {
		obj.ServiceDefinitionV2Dot1 = nil
	}

	// try to unmarshal data into ServiceDefinitionV2
	err = json.Unmarshal(data, &obj.ServiceDefinitionV2)
	if err == nil {
		if obj.ServiceDefinitionV2 != nil && obj.ServiceDefinitionV2.UnparsedObject == nil {
			jsonServiceDefinitionV2, _ := json.Marshal(obj.ServiceDefinitionV2)
			if string(jsonServiceDefinitionV2) == "{}" { // empty struct
				obj.ServiceDefinitionV2 = nil
			} else {
				match++
			}
		} else {
			obj.ServiceDefinitionV2 = nil
		}
	} else {
		obj.ServiceDefinitionV2 = nil
	}

	// try to unmarshal data into ServiceDefinitionRaw
	err = json.Unmarshal(data, &obj.ServiceDefinitionRaw)
	if err == nil {
		if obj.ServiceDefinitionRaw != nil {
			jsonServiceDefinitionRaw, _ := json.Marshal(obj.ServiceDefinitionRaw)
			if string(jsonServiceDefinitionRaw) == "{}" { // empty struct
				obj.ServiceDefinitionRaw = nil
			} else {
				match++
			}
		} else {
			obj.ServiceDefinitionRaw = nil
		}
	} else {
		obj.ServiceDefinitionRaw = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ServiceDefinitionV2Dot1 = nil
		obj.ServiceDefinitionV2 = nil
		obj.ServiceDefinitionRaw = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ServiceDefinitionsCreateRequest) MarshalJSON() ([]byte, error) {
	if obj.ServiceDefinitionV2Dot1 != nil {
		return json.Marshal(&obj.ServiceDefinitionV2Dot1)
	}

	if obj.ServiceDefinitionV2 != nil {
		return json.Marshal(&obj.ServiceDefinitionV2)
	}

	if obj.ServiceDefinitionRaw != nil {
		return json.Marshal(&obj.ServiceDefinitionRaw)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ServiceDefinitionsCreateRequest) GetActualInstance() interface{} {
	if obj.ServiceDefinitionV2Dot1 != nil {
		return obj.ServiceDefinitionV2Dot1
	}

	if obj.ServiceDefinitionV2 != nil {
		return obj.ServiceDefinitionV2
	}

	if obj.ServiceDefinitionRaw != nil {
		return obj.ServiceDefinitionRaw
	}

	// all schemas are nil
	return nil
}
