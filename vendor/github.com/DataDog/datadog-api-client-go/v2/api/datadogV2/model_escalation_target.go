// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationTarget - Represents an escalation target, which can be a team, user, or schedule.
type EscalationTarget struct {
	TeamTarget     *TeamTarget
	UserTarget     *UserTarget
	ScheduleTarget *ScheduleTarget

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TeamTargetAsEscalationTarget is a convenience function that returns TeamTarget wrapped in EscalationTarget.
func TeamTargetAsEscalationTarget(v *TeamTarget) EscalationTarget {
	return EscalationTarget{TeamTarget: v}
}

// UserTargetAsEscalationTarget is a convenience function that returns UserTarget wrapped in EscalationTarget.
func UserTargetAsEscalationTarget(v *UserTarget) EscalationTarget {
	return EscalationTarget{UserTarget: v}
}

// ScheduleTargetAsEscalationTarget is a convenience function that returns ScheduleTarget wrapped in EscalationTarget.
func ScheduleTargetAsEscalationTarget(v *ScheduleTarget) EscalationTarget {
	return EscalationTarget{ScheduleTarget: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *EscalationTarget) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TeamTarget
	err = datadog.Unmarshal(data, &obj.TeamTarget)
	if err == nil {
		if obj.TeamTarget != nil && obj.TeamTarget.UnparsedObject == nil {
			jsonTeamTarget, _ := datadog.Marshal(obj.TeamTarget)
			if string(jsonTeamTarget) == "{}" { // empty struct
				obj.TeamTarget = nil
			} else {
				match++
			}
		} else {
			obj.TeamTarget = nil
		}
	} else {
		obj.TeamTarget = nil
	}

	// try to unmarshal data into UserTarget
	err = datadog.Unmarshal(data, &obj.UserTarget)
	if err == nil {
		if obj.UserTarget != nil && obj.UserTarget.UnparsedObject == nil {
			jsonUserTarget, _ := datadog.Marshal(obj.UserTarget)
			if string(jsonUserTarget) == "{}" { // empty struct
				obj.UserTarget = nil
			} else {
				match++
			}
		} else {
			obj.UserTarget = nil
		}
	} else {
		obj.UserTarget = nil
	}

	// try to unmarshal data into ScheduleTarget
	err = datadog.Unmarshal(data, &obj.ScheduleTarget)
	if err == nil {
		if obj.ScheduleTarget != nil && obj.ScheduleTarget.UnparsedObject == nil {
			jsonScheduleTarget, _ := datadog.Marshal(obj.ScheduleTarget)
			if string(jsonScheduleTarget) == "{}" { // empty struct
				obj.ScheduleTarget = nil
			} else {
				match++
			}
		} else {
			obj.ScheduleTarget = nil
		}
	} else {
		obj.ScheduleTarget = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TeamTarget = nil
		obj.UserTarget = nil
		obj.ScheduleTarget = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj EscalationTarget) MarshalJSON() ([]byte, error) {
	if obj.TeamTarget != nil {
		return datadog.Marshal(&obj.TeamTarget)
	}

	if obj.UserTarget != nil {
		return datadog.Marshal(&obj.UserTarget)
	}

	if obj.ScheduleTarget != nil {
		return datadog.Marshal(&obj.ScheduleTarget)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *EscalationTarget) GetActualInstance() interface{} {
	if obj.TeamTarget != nil {
		return obj.TeamTarget
	}

	if obj.UserTarget != nil {
		return obj.UserTarget
	}

	if obj.ScheduleTarget != nil {
		return obj.ScheduleTarget
	}

	// all schemas are nil
	return nil
}
