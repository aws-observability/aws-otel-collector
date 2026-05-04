// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateNotificationChannelConfig - Defines the configuration for creating an On-Call notification channel
type CreateNotificationChannelConfig struct {
	CreatePhoneNotificationChannelConfig *CreatePhoneNotificationChannelConfig
	CreateEmailNotificationChannelConfig *CreateEmailNotificationChannelConfig

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CreatePhoneNotificationChannelConfigAsCreateNotificationChannelConfig is a convenience function that returns CreatePhoneNotificationChannelConfig wrapped in CreateNotificationChannelConfig.
func CreatePhoneNotificationChannelConfigAsCreateNotificationChannelConfig(v *CreatePhoneNotificationChannelConfig) CreateNotificationChannelConfig {
	return CreateNotificationChannelConfig{CreatePhoneNotificationChannelConfig: v}
}

// CreateEmailNotificationChannelConfigAsCreateNotificationChannelConfig is a convenience function that returns CreateEmailNotificationChannelConfig wrapped in CreateNotificationChannelConfig.
func CreateEmailNotificationChannelConfigAsCreateNotificationChannelConfig(v *CreateEmailNotificationChannelConfig) CreateNotificationChannelConfig {
	return CreateNotificationChannelConfig{CreateEmailNotificationChannelConfig: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CreateNotificationChannelConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreatePhoneNotificationChannelConfig
	err = datadog.Unmarshal(data, &obj.CreatePhoneNotificationChannelConfig)
	if err == nil {
		if obj.CreatePhoneNotificationChannelConfig != nil && obj.CreatePhoneNotificationChannelConfig.UnparsedObject == nil {
			jsonCreatePhoneNotificationChannelConfig, _ := datadog.Marshal(obj.CreatePhoneNotificationChannelConfig)
			if string(jsonCreatePhoneNotificationChannelConfig) == "{}" { // empty struct
				obj.CreatePhoneNotificationChannelConfig = nil
			} else {
				match++
			}
		} else {
			obj.CreatePhoneNotificationChannelConfig = nil
		}
	} else {
		obj.CreatePhoneNotificationChannelConfig = nil
	}

	// try to unmarshal data into CreateEmailNotificationChannelConfig
	err = datadog.Unmarshal(data, &obj.CreateEmailNotificationChannelConfig)
	if err == nil {
		if obj.CreateEmailNotificationChannelConfig != nil && obj.CreateEmailNotificationChannelConfig.UnparsedObject == nil {
			jsonCreateEmailNotificationChannelConfig, _ := datadog.Marshal(obj.CreateEmailNotificationChannelConfig)
			if string(jsonCreateEmailNotificationChannelConfig) == "{}" { // empty struct
				obj.CreateEmailNotificationChannelConfig = nil
			} else {
				match++
			}
		} else {
			obj.CreateEmailNotificationChannelConfig = nil
		}
	} else {
		obj.CreateEmailNotificationChannelConfig = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CreatePhoneNotificationChannelConfig = nil
		obj.CreateEmailNotificationChannelConfig = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CreateNotificationChannelConfig) MarshalJSON() ([]byte, error) {
	if obj.CreatePhoneNotificationChannelConfig != nil {
		return datadog.Marshal(&obj.CreatePhoneNotificationChannelConfig)
	}

	if obj.CreateEmailNotificationChannelConfig != nil {
		return datadog.Marshal(&obj.CreateEmailNotificationChannelConfig)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CreateNotificationChannelConfig) GetActualInstance() interface{} {
	if obj.CreatePhoneNotificationChannelConfig != nil {
		return obj.CreatePhoneNotificationChannelConfig
	}

	if obj.CreateEmailNotificationChannelConfig != nil {
		return obj.CreateEmailNotificationChannelConfig
	}

	// all schemas are nil
	return nil
}
