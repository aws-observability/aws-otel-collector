// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleDataIncludedItem - Any additional resources related to this schedule, such as teams and layers.
type ScheduleDataIncludedItem struct {
	TeamReference  *TeamReference
	Layer          *Layer
	ScheduleMember *ScheduleMember
	ScheduleUser   *ScheduleUser

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TeamReferenceAsScheduleDataIncludedItem is a convenience function that returns TeamReference wrapped in ScheduleDataIncludedItem.
func TeamReferenceAsScheduleDataIncludedItem(v *TeamReference) ScheduleDataIncludedItem {
	return ScheduleDataIncludedItem{TeamReference: v}
}

// LayerAsScheduleDataIncludedItem is a convenience function that returns Layer wrapped in ScheduleDataIncludedItem.
func LayerAsScheduleDataIncludedItem(v *Layer) ScheduleDataIncludedItem {
	return ScheduleDataIncludedItem{Layer: v}
}

// ScheduleMemberAsScheduleDataIncludedItem is a convenience function that returns ScheduleMember wrapped in ScheduleDataIncludedItem.
func ScheduleMemberAsScheduleDataIncludedItem(v *ScheduleMember) ScheduleDataIncludedItem {
	return ScheduleDataIncludedItem{ScheduleMember: v}
}

// ScheduleUserAsScheduleDataIncludedItem is a convenience function that returns ScheduleUser wrapped in ScheduleDataIncludedItem.
func ScheduleUserAsScheduleDataIncludedItem(v *ScheduleUser) ScheduleDataIncludedItem {
	return ScheduleDataIncludedItem{ScheduleUser: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ScheduleDataIncludedItem) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into Layer
	err = datadog.Unmarshal(data, &obj.Layer)
	if err == nil {
		if obj.Layer != nil && obj.Layer.UnparsedObject == nil {
			jsonLayer, _ := datadog.Marshal(obj.Layer)
			if string(jsonLayer) == "{}" { // empty struct
				obj.Layer = nil
			} else {
				match++
			}
		} else {
			obj.Layer = nil
		}
	} else {
		obj.Layer = nil
	}

	// try to unmarshal data into ScheduleMember
	err = datadog.Unmarshal(data, &obj.ScheduleMember)
	if err == nil {
		if obj.ScheduleMember != nil && obj.ScheduleMember.UnparsedObject == nil {
			jsonScheduleMember, _ := datadog.Marshal(obj.ScheduleMember)
			if string(jsonScheduleMember) == "{}" { // empty struct
				obj.ScheduleMember = nil
			} else {
				match++
			}
		} else {
			obj.ScheduleMember = nil
		}
	} else {
		obj.ScheduleMember = nil
	}

	// try to unmarshal data into ScheduleUser
	err = datadog.Unmarshal(data, &obj.ScheduleUser)
	if err == nil {
		if obj.ScheduleUser != nil && obj.ScheduleUser.UnparsedObject == nil {
			jsonScheduleUser, _ := datadog.Marshal(obj.ScheduleUser)
			if string(jsonScheduleUser) == "{}" { // empty struct
				obj.ScheduleUser = nil
			} else {
				match++
			}
		} else {
			obj.ScheduleUser = nil
		}
	} else {
		obj.ScheduleUser = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TeamReference = nil
		obj.Layer = nil
		obj.ScheduleMember = nil
		obj.ScheduleUser = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ScheduleDataIncludedItem) MarshalJSON() ([]byte, error) {
	if obj.TeamReference != nil {
		return datadog.Marshal(&obj.TeamReference)
	}

	if obj.Layer != nil {
		return datadog.Marshal(&obj.Layer)
	}

	if obj.ScheduleMember != nil {
		return datadog.Marshal(&obj.ScheduleMember)
	}

	if obj.ScheduleUser != nil {
		return datadog.Marshal(&obj.ScheduleUser)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ScheduleDataIncludedItem) GetActualInstance() interface{} {
	if obj.TeamReference != nil {
		return obj.TeamReference
	}

	if obj.Layer != nil {
		return obj.Layer
	}

	if obj.ScheduleMember != nil {
		return obj.ScheduleMember
	}

	if obj.ScheduleUser != nil {
		return obj.ScheduleUser
	}

	// all schemas are nil
	return nil
}
