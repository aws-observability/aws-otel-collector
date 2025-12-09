// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationRuleIncludedItems - Objects related to a notification rule.
type IncidentNotificationRuleIncludedItems struct {
	User                               *User
	IncidentTypeObject                 *IncidentTypeObject
	IncidentNotificationTemplateObject *IncidentNotificationTemplateObject

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// UserAsIncidentNotificationRuleIncludedItems is a convenience function that returns User wrapped in IncidentNotificationRuleIncludedItems.
func UserAsIncidentNotificationRuleIncludedItems(v *User) IncidentNotificationRuleIncludedItems {
	return IncidentNotificationRuleIncludedItems{User: v}
}

// IncidentTypeObjectAsIncidentNotificationRuleIncludedItems is a convenience function that returns IncidentTypeObject wrapped in IncidentNotificationRuleIncludedItems.
func IncidentTypeObjectAsIncidentNotificationRuleIncludedItems(v *IncidentTypeObject) IncidentNotificationRuleIncludedItems {
	return IncidentNotificationRuleIncludedItems{IncidentTypeObject: v}
}

// IncidentNotificationTemplateObjectAsIncidentNotificationRuleIncludedItems is a convenience function that returns IncidentNotificationTemplateObject wrapped in IncidentNotificationRuleIncludedItems.
func IncidentNotificationTemplateObjectAsIncidentNotificationRuleIncludedItems(v *IncidentNotificationTemplateObject) IncidentNotificationRuleIncludedItems {
	return IncidentNotificationRuleIncludedItems{IncidentNotificationTemplateObject: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IncidentNotificationRuleIncludedItems) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into IncidentTypeObject
	err = datadog.Unmarshal(data, &obj.IncidentTypeObject)
	if err == nil {
		if obj.IncidentTypeObject != nil && obj.IncidentTypeObject.UnparsedObject == nil {
			jsonIncidentTypeObject, _ := datadog.Marshal(obj.IncidentTypeObject)
			if string(jsonIncidentTypeObject) == "{}" { // empty struct
				obj.IncidentTypeObject = nil
			} else {
				match++
			}
		} else {
			obj.IncidentTypeObject = nil
		}
	} else {
		obj.IncidentTypeObject = nil
	}

	// try to unmarshal data into IncidentNotificationTemplateObject
	err = datadog.Unmarshal(data, &obj.IncidentNotificationTemplateObject)
	if err == nil {
		if obj.IncidentNotificationTemplateObject != nil && obj.IncidentNotificationTemplateObject.UnparsedObject == nil {
			jsonIncidentNotificationTemplateObject, _ := datadog.Marshal(obj.IncidentNotificationTemplateObject)
			if string(jsonIncidentNotificationTemplateObject) == "{}" { // empty struct
				obj.IncidentNotificationTemplateObject = nil
			} else {
				match++
			}
		} else {
			obj.IncidentNotificationTemplateObject = nil
		}
	} else {
		obj.IncidentNotificationTemplateObject = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.User = nil
		obj.IncidentTypeObject = nil
		obj.IncidentNotificationTemplateObject = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IncidentNotificationRuleIncludedItems) MarshalJSON() ([]byte, error) {
	if obj.User != nil {
		return datadog.Marshal(&obj.User)
	}

	if obj.IncidentTypeObject != nil {
		return datadog.Marshal(&obj.IncidentTypeObject)
	}

	if obj.IncidentNotificationTemplateObject != nil {
		return datadog.Marshal(&obj.IncidentNotificationTemplateObject)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IncidentNotificationRuleIncludedItems) GetActualInstance() interface{} {
	if obj.User != nil {
		return obj.User
	}

	if obj.IncidentTypeObject != nil {
		return obj.IncidentTypeObject
	}

	if obj.IncidentNotificationTemplateObject != nil {
		return obj.IncidentNotificationTemplateObject
	}

	// all schemas are nil
	return nil
}
