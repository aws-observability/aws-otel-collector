// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationChannelConfig - Defines the configuration for an On-Call notification channel
type NotificationChannelConfig struct {
	NotificationChannelPhoneConfig *NotificationChannelPhoneConfig
	NotificationChannelEmailConfig *NotificationChannelEmailConfig
	NotificationChannelPushConfig  *NotificationChannelPushConfig

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NotificationChannelPhoneConfigAsNotificationChannelConfig is a convenience function that returns NotificationChannelPhoneConfig wrapped in NotificationChannelConfig.
func NotificationChannelPhoneConfigAsNotificationChannelConfig(v *NotificationChannelPhoneConfig) NotificationChannelConfig {
	return NotificationChannelConfig{NotificationChannelPhoneConfig: v}
}

// NotificationChannelEmailConfigAsNotificationChannelConfig is a convenience function that returns NotificationChannelEmailConfig wrapped in NotificationChannelConfig.
func NotificationChannelEmailConfigAsNotificationChannelConfig(v *NotificationChannelEmailConfig) NotificationChannelConfig {
	return NotificationChannelConfig{NotificationChannelEmailConfig: v}
}

// NotificationChannelPushConfigAsNotificationChannelConfig is a convenience function that returns NotificationChannelPushConfig wrapped in NotificationChannelConfig.
func NotificationChannelPushConfigAsNotificationChannelConfig(v *NotificationChannelPushConfig) NotificationChannelConfig {
	return NotificationChannelConfig{NotificationChannelPushConfig: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *NotificationChannelConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotificationChannelPhoneConfig
	err = datadog.Unmarshal(data, &obj.NotificationChannelPhoneConfig)
	if err == nil {
		if obj.NotificationChannelPhoneConfig != nil && obj.NotificationChannelPhoneConfig.UnparsedObject == nil {
			jsonNotificationChannelPhoneConfig, _ := datadog.Marshal(obj.NotificationChannelPhoneConfig)
			if string(jsonNotificationChannelPhoneConfig) == "{}" { // empty struct
				obj.NotificationChannelPhoneConfig = nil
			} else {
				match++
			}
		} else {
			obj.NotificationChannelPhoneConfig = nil
		}
	} else {
		obj.NotificationChannelPhoneConfig = nil
	}

	// try to unmarshal data into NotificationChannelEmailConfig
	err = datadog.Unmarshal(data, &obj.NotificationChannelEmailConfig)
	if err == nil {
		if obj.NotificationChannelEmailConfig != nil && obj.NotificationChannelEmailConfig.UnparsedObject == nil {
			jsonNotificationChannelEmailConfig, _ := datadog.Marshal(obj.NotificationChannelEmailConfig)
			if string(jsonNotificationChannelEmailConfig) == "{}" { // empty struct
				obj.NotificationChannelEmailConfig = nil
			} else {
				match++
			}
		} else {
			obj.NotificationChannelEmailConfig = nil
		}
	} else {
		obj.NotificationChannelEmailConfig = nil
	}

	// try to unmarshal data into NotificationChannelPushConfig
	err = datadog.Unmarshal(data, &obj.NotificationChannelPushConfig)
	if err == nil {
		if obj.NotificationChannelPushConfig != nil && obj.NotificationChannelPushConfig.UnparsedObject == nil {
			jsonNotificationChannelPushConfig, _ := datadog.Marshal(obj.NotificationChannelPushConfig)
			if string(jsonNotificationChannelPushConfig) == "{}" { // empty struct
				obj.NotificationChannelPushConfig = nil
			} else {
				match++
			}
		} else {
			obj.NotificationChannelPushConfig = nil
		}
	} else {
		obj.NotificationChannelPushConfig = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.NotificationChannelPhoneConfig = nil
		obj.NotificationChannelEmailConfig = nil
		obj.NotificationChannelPushConfig = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj NotificationChannelConfig) MarshalJSON() ([]byte, error) {
	if obj.NotificationChannelPhoneConfig != nil {
		return datadog.Marshal(&obj.NotificationChannelPhoneConfig)
	}

	if obj.NotificationChannelEmailConfig != nil {
		return datadog.Marshal(&obj.NotificationChannelEmailConfig)
	}

	if obj.NotificationChannelPushConfig != nil {
		return datadog.Marshal(&obj.NotificationChannelPushConfig)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *NotificationChannelConfig) GetActualInstance() interface{} {
	if obj.NotificationChannelPhoneConfig != nil {
		return obj.NotificationChannelPhoneConfig
	}

	if obj.NotificationChannelEmailConfig != nil {
		return obj.NotificationChannelEmailConfig
	}

	if obj.NotificationChannelPushConfig != nil {
		return obj.NotificationChannelPushConfig
	}

	// all schemas are nil
	return nil
}
