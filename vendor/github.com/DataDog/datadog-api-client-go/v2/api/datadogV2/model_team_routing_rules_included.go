// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamRoutingRulesIncluded - Represents additional included resources for team routing rules, such as associated routing rules.
type TeamRoutingRulesIncluded struct {
	RoutingRule *RoutingRule

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// RoutingRuleAsTeamRoutingRulesIncluded is a convenience function that returns RoutingRule wrapped in TeamRoutingRulesIncluded.
func RoutingRuleAsTeamRoutingRulesIncluded(v *RoutingRule) TeamRoutingRulesIncluded {
	return TeamRoutingRulesIncluded{RoutingRule: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TeamRoutingRulesIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RoutingRule
	err = datadog.Unmarshal(data, &obj.RoutingRule)
	if err == nil {
		if obj.RoutingRule != nil && obj.RoutingRule.UnparsedObject == nil {
			jsonRoutingRule, _ := datadog.Marshal(obj.RoutingRule)
			if string(jsonRoutingRule) == "{}" { // empty struct
				obj.RoutingRule = nil
			} else {
				match++
			}
		} else {
			obj.RoutingRule = nil
		}
	} else {
		obj.RoutingRule = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.RoutingRule = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TeamRoutingRulesIncluded) MarshalJSON() ([]byte, error) {
	if obj.RoutingRule != nil {
		return datadog.Marshal(&obj.RoutingRule)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TeamRoutingRulesIncluded) GetActualInstance() interface{} {
	if obj.RoutingRule != nil {
		return obj.RoutingRule
	}

	// all schemas are nil
	return nil
}
