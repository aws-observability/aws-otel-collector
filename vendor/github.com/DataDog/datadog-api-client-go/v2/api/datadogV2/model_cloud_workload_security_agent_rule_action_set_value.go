// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleActionSetValue - The value of the set action
type CloudWorkloadSecurityAgentRuleActionSetValue struct {
	String *string
	Int32  *int32
	Bool   *bool

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsCloudWorkloadSecurityAgentRuleActionSetValue is a convenience function that returns string wrapped in CloudWorkloadSecurityAgentRuleActionSetValue.
func StringAsCloudWorkloadSecurityAgentRuleActionSetValue(v *string) CloudWorkloadSecurityAgentRuleActionSetValue {
	return CloudWorkloadSecurityAgentRuleActionSetValue{String: v}
}

// Int32AsCloudWorkloadSecurityAgentRuleActionSetValue is a convenience function that returns int32 wrapped in CloudWorkloadSecurityAgentRuleActionSetValue.
func Int32AsCloudWorkloadSecurityAgentRuleActionSetValue(v *int32) CloudWorkloadSecurityAgentRuleActionSetValue {
	return CloudWorkloadSecurityAgentRuleActionSetValue{Int32: v}
}

// BoolAsCloudWorkloadSecurityAgentRuleActionSetValue is a convenience function that returns bool wrapped in CloudWorkloadSecurityAgentRuleActionSetValue.
func BoolAsCloudWorkloadSecurityAgentRuleActionSetValue(v *bool) CloudWorkloadSecurityAgentRuleActionSetValue {
	return CloudWorkloadSecurityAgentRuleActionSetValue{Bool: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CloudWorkloadSecurityAgentRuleActionSetValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = datadog.Unmarshal(data, &obj.String)
	if err == nil {
		if obj.String != nil {
			jsonString, _ := datadog.Marshal(obj.String)
			if string(jsonString) == "{}" { // empty struct
				obj.String = nil
			} else {
				match++
			}
		} else {
			obj.String = nil
		}
	} else {
		obj.String = nil
	}

	// try to unmarshal data into Int32
	err = datadog.Unmarshal(data, &obj.Int32)
	if err == nil {
		if obj.Int32 != nil {
			jsonInt32, _ := datadog.Marshal(obj.Int32)
			if string(jsonInt32) == "{}" { // empty struct
				obj.Int32 = nil
			} else {
				match++
			}
		} else {
			obj.Int32 = nil
		}
	} else {
		obj.Int32 = nil
	}

	// try to unmarshal data into Bool
	err = datadog.Unmarshal(data, &obj.Bool)
	if err == nil {
		if obj.Bool != nil {
			jsonBool, _ := datadog.Marshal(obj.Bool)
			if string(jsonBool) == "{}" { // empty struct
				obj.Bool = nil
			} else {
				match++
			}
		} else {
			obj.Bool = nil
		}
	} else {
		obj.Bool = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.Int32 = nil
		obj.Bool = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CloudWorkloadSecurityAgentRuleActionSetValue) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.Int32 != nil {
		return datadog.Marshal(&obj.Int32)
	}

	if obj.Bool != nil {
		return datadog.Marshal(&obj.Bool)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CloudWorkloadSecurityAgentRuleActionSetValue) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	// all schemas are nil
	return nil
}
