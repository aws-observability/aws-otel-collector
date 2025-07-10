// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyIncluded - Represents included related resources when retrieving an escalation policy, such as teams, steps, or targets.
type EscalationPolicyIncluded struct {
	TeamReference        *TeamReference
	EscalationPolicyStep *EscalationPolicyStep
	EscalationPolicyUser *EscalationPolicyUser
	ScheduleData         *ScheduleData

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TeamReferenceAsEscalationPolicyIncluded is a convenience function that returns TeamReference wrapped in EscalationPolicyIncluded.
func TeamReferenceAsEscalationPolicyIncluded(v *TeamReference) EscalationPolicyIncluded {
	return EscalationPolicyIncluded{TeamReference: v}
}

// EscalationPolicyStepAsEscalationPolicyIncluded is a convenience function that returns EscalationPolicyStep wrapped in EscalationPolicyIncluded.
func EscalationPolicyStepAsEscalationPolicyIncluded(v *EscalationPolicyStep) EscalationPolicyIncluded {
	return EscalationPolicyIncluded{EscalationPolicyStep: v}
}

// EscalationPolicyUserAsEscalationPolicyIncluded is a convenience function that returns EscalationPolicyUser wrapped in EscalationPolicyIncluded.
func EscalationPolicyUserAsEscalationPolicyIncluded(v *EscalationPolicyUser) EscalationPolicyIncluded {
	return EscalationPolicyIncluded{EscalationPolicyUser: v}
}

// ScheduleDataAsEscalationPolicyIncluded is a convenience function that returns ScheduleData wrapped in EscalationPolicyIncluded.
func ScheduleDataAsEscalationPolicyIncluded(v *ScheduleData) EscalationPolicyIncluded {
	return EscalationPolicyIncluded{ScheduleData: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *EscalationPolicyIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TeamReference
	err = datadog.Unmarshal(data, &obj.TeamReference)
	if err == nil {
		if obj.TeamReference != nil && obj.TeamReference.UnparsedObject == nil {
			jsonTeamReference, _ := datadog.Marshal(obj.TeamReference)
			if string(jsonTeamReference) == "{}" { // empty struct
				obj.TeamReference = nil
			} else {
				match++
			}
		} else {
			obj.TeamReference = nil
		}
	} else {
		obj.TeamReference = nil
	}

	// try to unmarshal data into EscalationPolicyStep
	err = datadog.Unmarshal(data, &obj.EscalationPolicyStep)
	if err == nil {
		if obj.EscalationPolicyStep != nil && obj.EscalationPolicyStep.UnparsedObject == nil {
			jsonEscalationPolicyStep, _ := datadog.Marshal(obj.EscalationPolicyStep)
			if string(jsonEscalationPolicyStep) == "{}" { // empty struct
				obj.EscalationPolicyStep = nil
			} else {
				match++
			}
		} else {
			obj.EscalationPolicyStep = nil
		}
	} else {
		obj.EscalationPolicyStep = nil
	}

	// try to unmarshal data into EscalationPolicyUser
	err = datadog.Unmarshal(data, &obj.EscalationPolicyUser)
	if err == nil {
		if obj.EscalationPolicyUser != nil && obj.EscalationPolicyUser.UnparsedObject == nil {
			jsonEscalationPolicyUser, _ := datadog.Marshal(obj.EscalationPolicyUser)
			if string(jsonEscalationPolicyUser) == "{}" { // empty struct
				obj.EscalationPolicyUser = nil
			} else {
				match++
			}
		} else {
			obj.EscalationPolicyUser = nil
		}
	} else {
		obj.EscalationPolicyUser = nil
	}

	// try to unmarshal data into ScheduleData
	err = datadog.Unmarshal(data, &obj.ScheduleData)
	if err == nil {
		if obj.ScheduleData != nil && obj.ScheduleData.UnparsedObject == nil {
			jsonScheduleData, _ := datadog.Marshal(obj.ScheduleData)
			if string(jsonScheduleData) == "{}" { // empty struct
				obj.ScheduleData = nil
			} else {
				match++
			}
		} else {
			obj.ScheduleData = nil
		}
	} else {
		obj.ScheduleData = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TeamReference = nil
		obj.EscalationPolicyStep = nil
		obj.EscalationPolicyUser = nil
		obj.ScheduleData = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj EscalationPolicyIncluded) MarshalJSON() ([]byte, error) {
	if obj.TeamReference != nil {
		return datadog.Marshal(&obj.TeamReference)
	}

	if obj.EscalationPolicyStep != nil {
		return datadog.Marshal(&obj.EscalationPolicyStep)
	}

	if obj.EscalationPolicyUser != nil {
		return datadog.Marshal(&obj.EscalationPolicyUser)
	}

	if obj.ScheduleData != nil {
		return datadog.Marshal(&obj.ScheduleData)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *EscalationPolicyIncluded) GetActualInstance() interface{} {
	if obj.TeamReference != nil {
		return obj.TeamReference
	}

	if obj.EscalationPolicyStep != nil {
		return obj.EscalationPolicyStep
	}

	if obj.EscalationPolicyUser != nil {
		return obj.EscalationPolicyUser
	}

	if obj.ScheduleData != nil {
		return obj.ScheduleData
	}

	// all schemas are nil
	return nil
}
