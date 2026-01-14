// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationChannelPushConfig Push notification channel configuration
type NotificationChannelPushConfig struct {
	// The name of the application used to receive push notifications
	ApplicationName string `json:"application_name"`
	// The name of the mobile device being used
	DeviceName string `json:"device_name"`
	// Indicates that the notification channel is a mobile device for push notifications
	Type NotificationChannelPushConfigType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotificationChannelPushConfig instantiates a new NotificationChannelPushConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotificationChannelPushConfig(applicationName string, deviceName string, typeVar NotificationChannelPushConfigType) *NotificationChannelPushConfig {
	this := NotificationChannelPushConfig{}
	this.ApplicationName = applicationName
	this.DeviceName = deviceName
	this.Type = typeVar
	return &this
}

// NewNotificationChannelPushConfigWithDefaults instantiates a new NotificationChannelPushConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotificationChannelPushConfigWithDefaults() *NotificationChannelPushConfig {
	this := NotificationChannelPushConfig{}
	var typeVar NotificationChannelPushConfigType = NOTIFICATIONCHANNELPUSHCONFIGTYPE_PUSH
	this.Type = typeVar
	return &this
}

// GetApplicationName returns the ApplicationName field value.
func (o *NotificationChannelPushConfig) GetApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value
// and a boolean to check if the value has been set.
func (o *NotificationChannelPushConfig) GetApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationName, true
}

// SetApplicationName sets field value.
func (o *NotificationChannelPushConfig) SetApplicationName(v string) {
	o.ApplicationName = v
}

// GetDeviceName returns the DeviceName field value.
func (o *NotificationChannelPushConfig) GetDeviceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value
// and a boolean to check if the value has been set.
func (o *NotificationChannelPushConfig) GetDeviceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceName, true
}

// SetDeviceName sets field value.
func (o *NotificationChannelPushConfig) SetDeviceName(v string) {
	o.DeviceName = v
}

// GetType returns the Type field value.
func (o *NotificationChannelPushConfig) GetType() NotificationChannelPushConfigType {
	if o == nil {
		var ret NotificationChannelPushConfigType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationChannelPushConfig) GetTypeOk() (*NotificationChannelPushConfigType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NotificationChannelPushConfig) SetType(v NotificationChannelPushConfigType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotificationChannelPushConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_name"] = o.ApplicationName
	toSerialize["device_name"] = o.DeviceName
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotificationChannelPushConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationName *string                            `json:"application_name"`
		DeviceName      *string                            `json:"device_name"`
		Type            *NotificationChannelPushConfigType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationName == nil {
		return fmt.Errorf("required field application_name missing")
	}
	if all.DeviceName == nil {
		return fmt.Errorf("required field device_name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_name", "device_name", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationName = *all.ApplicationName
	o.DeviceName = *all.DeviceName
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
