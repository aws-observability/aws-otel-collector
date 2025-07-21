// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamOnCallRespondersIncluded - Represents an union of related resources included in the response, such as users and escalation steps.
type TeamOnCallRespondersIncluded struct {
	User       *User
	Escalation *Escalation

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// UserAsTeamOnCallRespondersIncluded is a convenience function that returns User wrapped in TeamOnCallRespondersIncluded.
func UserAsTeamOnCallRespondersIncluded(v *User) TeamOnCallRespondersIncluded {
	return TeamOnCallRespondersIncluded{User: v}
}

// EscalationAsTeamOnCallRespondersIncluded is a convenience function that returns Escalation wrapped in TeamOnCallRespondersIncluded.
func EscalationAsTeamOnCallRespondersIncluded(v *Escalation) TeamOnCallRespondersIncluded {
	return TeamOnCallRespondersIncluded{Escalation: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TeamOnCallRespondersIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into User
	err = datadog.Unmarshal(data, &obj.User)
	if err == nil {
		if obj.User != nil && obj.User.UnparsedObject == nil {
			jsonUser, _ := datadog.Marshal(obj.User)
			if string(jsonUser) == "{}" && string(data) != "{}" { // empty struct
				obj.User = nil
			} else {
				match++
			}
		} else {
			obj.User = nil
		}
	} else {
		obj.User = nil
	}

	// try to unmarshal data into Escalation
	err = datadog.Unmarshal(data, &obj.Escalation)
	if err == nil {
		if obj.Escalation != nil && obj.Escalation.UnparsedObject == nil {
			jsonEscalation, _ := datadog.Marshal(obj.Escalation)
			if string(jsonEscalation) == "{}" { // empty struct
				obj.Escalation = nil
			} else {
				match++
			}
		} else {
			obj.Escalation = nil
		}
	} else {
		obj.Escalation = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.User = nil
		obj.Escalation = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TeamOnCallRespondersIncluded) MarshalJSON() ([]byte, error) {
	if obj.User != nil {
		return datadog.Marshal(&obj.User)
	}

	if obj.Escalation != nil {
		return datadog.Marshal(&obj.Escalation)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TeamOnCallRespondersIncluded) GetActualInstance() interface{} {
	if obj.User != nil {
		return obj.User
	}

	if obj.Escalation != nil {
		return obj.Escalation
	}

	// all schemas are nil
	return nil
}
