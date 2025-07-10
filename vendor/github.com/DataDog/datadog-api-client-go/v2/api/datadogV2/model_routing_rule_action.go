// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoutingRuleAction - Defines an action that is executed when a routing rule matches certain criteria.
type RoutingRuleAction struct {
	SendSlackMessageAction *SendSlackMessageAction
	SendTeamsMessageAction *SendTeamsMessageAction

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SendSlackMessageActionAsRoutingRuleAction is a convenience function that returns SendSlackMessageAction wrapped in RoutingRuleAction.
func SendSlackMessageActionAsRoutingRuleAction(v *SendSlackMessageAction) RoutingRuleAction {
	return RoutingRuleAction{SendSlackMessageAction: v}
}

// SendTeamsMessageActionAsRoutingRuleAction is a convenience function that returns SendTeamsMessageAction wrapped in RoutingRuleAction.
func SendTeamsMessageActionAsRoutingRuleAction(v *SendTeamsMessageAction) RoutingRuleAction {
	return RoutingRuleAction{SendTeamsMessageAction: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *RoutingRuleAction) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SendSlackMessageAction
	err = datadog.Unmarshal(data, &obj.SendSlackMessageAction)
	if err == nil {
		if obj.SendSlackMessageAction != nil && obj.SendSlackMessageAction.UnparsedObject == nil {
			jsonSendSlackMessageAction, _ := datadog.Marshal(obj.SendSlackMessageAction)
			if string(jsonSendSlackMessageAction) == "{}" { // empty struct
				obj.SendSlackMessageAction = nil
			} else {
				match++
			}
		} else {
			obj.SendSlackMessageAction = nil
		}
	} else {
		obj.SendSlackMessageAction = nil
	}

	// try to unmarshal data into SendTeamsMessageAction
	err = datadog.Unmarshal(data, &obj.SendTeamsMessageAction)
	if err == nil {
		if obj.SendTeamsMessageAction != nil && obj.SendTeamsMessageAction.UnparsedObject == nil {
			jsonSendTeamsMessageAction, _ := datadog.Marshal(obj.SendTeamsMessageAction)
			if string(jsonSendTeamsMessageAction) == "{}" { // empty struct
				obj.SendTeamsMessageAction = nil
			} else {
				match++
			}
		} else {
			obj.SendTeamsMessageAction = nil
		}
	} else {
		obj.SendTeamsMessageAction = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SendSlackMessageAction = nil
		obj.SendTeamsMessageAction = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj RoutingRuleAction) MarshalJSON() ([]byte, error) {
	if obj.SendSlackMessageAction != nil {
		return datadog.Marshal(&obj.SendSlackMessageAction)
	}

	if obj.SendTeamsMessageAction != nil {
		return datadog.Marshal(&obj.SendTeamsMessageAction)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *RoutingRuleAction) GetActualInstance() interface{} {
	if obj.SendSlackMessageAction != nil {
		return obj.SendSlackMessageAction
	}

	if obj.SendTeamsMessageAction != nil {
		return obj.SendTeamsMessageAction
	}

	// all schemas are nil
	return nil
}
