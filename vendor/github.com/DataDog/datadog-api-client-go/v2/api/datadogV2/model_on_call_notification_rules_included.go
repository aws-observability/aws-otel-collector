// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallNotificationRulesIncluded - Represents additional included resources for a on-call notification rules
type OnCallNotificationRulesIncluded struct {
	NotificationChannelData *NotificationChannelData

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NotificationChannelDataAsOnCallNotificationRulesIncluded is a convenience function that returns NotificationChannelData wrapped in OnCallNotificationRulesIncluded.
func NotificationChannelDataAsOnCallNotificationRulesIncluded(v *NotificationChannelData) OnCallNotificationRulesIncluded {
	return OnCallNotificationRulesIncluded{NotificationChannelData: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OnCallNotificationRulesIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotificationChannelData
	err = datadog.Unmarshal(data, &obj.NotificationChannelData)
	if err == nil {
		if obj.NotificationChannelData != nil && obj.NotificationChannelData.UnparsedObject == nil {
			jsonNotificationChannelData, _ := datadog.Marshal(obj.NotificationChannelData)
			if string(jsonNotificationChannelData) == "{}" { // empty struct
				obj.NotificationChannelData = nil
			} else {
				match++
			}
		} else {
			obj.NotificationChannelData = nil
		}
	} else {
		obj.NotificationChannelData = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.NotificationChannelData = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OnCallNotificationRulesIncluded) MarshalJSON() ([]byte, error) {
	if obj.NotificationChannelData != nil {
		return datadog.Marshal(&obj.NotificationChannelData)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OnCallNotificationRulesIncluded) GetActualInstance() interface{} {
	if obj.NotificationChannelData != nil {
		return obj.NotificationChannelData
	}

	// all schemas are nil
	return nil
}
