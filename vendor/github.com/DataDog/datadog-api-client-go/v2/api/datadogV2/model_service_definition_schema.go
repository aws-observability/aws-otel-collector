// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionSchema - Service definition schema.
type ServiceDefinitionSchema struct {
	ServiceDefinitionV1     *ServiceDefinitionV1
	ServiceDefinitionV2     *ServiceDefinitionV2
	ServiceDefinitionV2Dot1 *ServiceDefinitionV2Dot1
	ServiceDefinitionV2Dot2 *ServiceDefinitionV2Dot2

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ServiceDefinitionV1AsServiceDefinitionSchema is a convenience function that returns ServiceDefinitionV1 wrapped in ServiceDefinitionSchema.
func ServiceDefinitionV1AsServiceDefinitionSchema(v *ServiceDefinitionV1) ServiceDefinitionSchema {
	return ServiceDefinitionSchema{ServiceDefinitionV1: v}
}

// ServiceDefinitionV2AsServiceDefinitionSchema is a convenience function that returns ServiceDefinitionV2 wrapped in ServiceDefinitionSchema.
func ServiceDefinitionV2AsServiceDefinitionSchema(v *ServiceDefinitionV2) ServiceDefinitionSchema {
	return ServiceDefinitionSchema{ServiceDefinitionV2: v}
}

// ServiceDefinitionV2Dot1AsServiceDefinitionSchema is a convenience function that returns ServiceDefinitionV2Dot1 wrapped in ServiceDefinitionSchema.
func ServiceDefinitionV2Dot1AsServiceDefinitionSchema(v *ServiceDefinitionV2Dot1) ServiceDefinitionSchema {
	return ServiceDefinitionSchema{ServiceDefinitionV2Dot1: v}
}

// ServiceDefinitionV2Dot2AsServiceDefinitionSchema is a convenience function that returns ServiceDefinitionV2Dot2 wrapped in ServiceDefinitionSchema.
func ServiceDefinitionV2Dot2AsServiceDefinitionSchema(v *ServiceDefinitionV2Dot2) ServiceDefinitionSchema {
	return ServiceDefinitionSchema{ServiceDefinitionV2Dot2: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ServiceDefinitionSchema) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ServiceDefinitionV1
	err = datadog.Unmarshal(data, &obj.ServiceDefinitionV1)
	if err == nil {
		if obj.ServiceDefinitionV1 != nil && obj.ServiceDefinitionV1.UnparsedObject == nil {
			jsonServiceDefinitionV1, _ := datadog.Marshal(obj.ServiceDefinitionV1)
			if string(jsonServiceDefinitionV1) == "{}" { // empty struct
				obj.ServiceDefinitionV1 = nil
			} else {
				match++
			}
		} else {
			obj.ServiceDefinitionV1 = nil
		}
	} else {
		obj.ServiceDefinitionV1 = nil
	}

	// try to unmarshal data into ServiceDefinitionV2
	err = datadog.Unmarshal(data, &obj.ServiceDefinitionV2)
	if err == nil {
		if obj.ServiceDefinitionV2 != nil && obj.ServiceDefinitionV2.UnparsedObject == nil {
			jsonServiceDefinitionV2, _ := datadog.Marshal(obj.ServiceDefinitionV2)
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

	// try to unmarshal data into ServiceDefinitionV2Dot1
	err = datadog.Unmarshal(data, &obj.ServiceDefinitionV2Dot1)
	if err == nil {
		if obj.ServiceDefinitionV2Dot1 != nil && obj.ServiceDefinitionV2Dot1.UnparsedObject == nil {
			jsonServiceDefinitionV2Dot1, _ := datadog.Marshal(obj.ServiceDefinitionV2Dot1)
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

	// try to unmarshal data into ServiceDefinitionV2Dot2
	err = datadog.Unmarshal(data, &obj.ServiceDefinitionV2Dot2)
	if err == nil {
		if obj.ServiceDefinitionV2Dot2 != nil && obj.ServiceDefinitionV2Dot2.UnparsedObject == nil {
			jsonServiceDefinitionV2Dot2, _ := datadog.Marshal(obj.ServiceDefinitionV2Dot2)
			if string(jsonServiceDefinitionV2Dot2) == "{}" { // empty struct
				obj.ServiceDefinitionV2Dot2 = nil
			} else {
				match++
			}
		} else {
			obj.ServiceDefinitionV2Dot2 = nil
		}
	} else {
		obj.ServiceDefinitionV2Dot2 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ServiceDefinitionV1 = nil
		obj.ServiceDefinitionV2 = nil
		obj.ServiceDefinitionV2Dot1 = nil
		obj.ServiceDefinitionV2Dot2 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ServiceDefinitionSchema) MarshalJSON() ([]byte, error) {
	if obj.ServiceDefinitionV1 != nil {
		return datadog.Marshal(&obj.ServiceDefinitionV1)
	}

	if obj.ServiceDefinitionV2 != nil {
		return datadog.Marshal(&obj.ServiceDefinitionV2)
	}

	if obj.ServiceDefinitionV2Dot1 != nil {
		return datadog.Marshal(&obj.ServiceDefinitionV2Dot1)
	}

	if obj.ServiceDefinitionV2Dot2 != nil {
		return datadog.Marshal(&obj.ServiceDefinitionV2Dot2)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ServiceDefinitionSchema) GetActualInstance() interface{} {
	if obj.ServiceDefinitionV1 != nil {
		return obj.ServiceDefinitionV1
	}

	if obj.ServiceDefinitionV2 != nil {
		return obj.ServiceDefinitionV2
	}

	if obj.ServiceDefinitionV2Dot1 != nil {
		return obj.ServiceDefinitionV2Dot1
	}

	if obj.ServiceDefinitionV2Dot2 != nil {
		return obj.ServiceDefinitionV2Dot2
	}

	// all schemas are nil
	return nil
}
