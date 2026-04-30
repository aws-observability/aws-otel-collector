// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallNotificationRuleChannelSettings - Defines the configuration for a channel associated with a notification rule
type OnCallNotificationRuleChannelSettings struct {
	OnCallPhoneNotificationRuleSettings *OnCallPhoneNotificationRuleSettings

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OnCallPhoneNotificationRuleSettingsAsOnCallNotificationRuleChannelSettings is a convenience function that returns OnCallPhoneNotificationRuleSettings wrapped in OnCallNotificationRuleChannelSettings.
func OnCallPhoneNotificationRuleSettingsAsOnCallNotificationRuleChannelSettings(v *OnCallPhoneNotificationRuleSettings) OnCallNotificationRuleChannelSettings {
	return OnCallNotificationRuleChannelSettings{OnCallPhoneNotificationRuleSettings: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OnCallNotificationRuleChannelSettings) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OnCallPhoneNotificationRuleSettings
	err = datadog.Unmarshal(data, &obj.OnCallPhoneNotificationRuleSettings)
	if err == nil {
		if obj.OnCallPhoneNotificationRuleSettings != nil && obj.OnCallPhoneNotificationRuleSettings.UnparsedObject == nil {
			jsonOnCallPhoneNotificationRuleSettings, _ := datadog.Marshal(obj.OnCallPhoneNotificationRuleSettings)
			if string(jsonOnCallPhoneNotificationRuleSettings) == "{}" { // empty struct
				obj.OnCallPhoneNotificationRuleSettings = nil
			} else {
				match++
			}
		} else {
			obj.OnCallPhoneNotificationRuleSettings = nil
		}
	} else {
		obj.OnCallPhoneNotificationRuleSettings = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OnCallPhoneNotificationRuleSettings = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OnCallNotificationRuleChannelSettings) MarshalJSON() ([]byte, error) {
	if obj.OnCallPhoneNotificationRuleSettings != nil {
		return datadog.Marshal(&obj.OnCallPhoneNotificationRuleSettings)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OnCallNotificationRuleChannelSettings) GetActualInstance() interface{} {
	if obj.OnCallPhoneNotificationRuleSettings != nil {
		return obj.OnCallPhoneNotificationRuleSettings
	}

	// all schemas are nil
	return nil
}
