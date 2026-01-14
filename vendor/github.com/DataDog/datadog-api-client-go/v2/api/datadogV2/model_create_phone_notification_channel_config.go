// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreatePhoneNotificationChannelConfig Configuration to create a phone notification channel
type CreatePhoneNotificationChannelConfig struct {
	// The E-164 formatted phone number (e.g. +3371234567)
	Number string `json:"number"`
	// Indicates that the notification channel is a phone
	Type NotificationChannelPhoneConfigType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreatePhoneNotificationChannelConfig instantiates a new CreatePhoneNotificationChannelConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreatePhoneNotificationChannelConfig(number string, typeVar NotificationChannelPhoneConfigType) *CreatePhoneNotificationChannelConfig {
	this := CreatePhoneNotificationChannelConfig{}
	this.Number = number
	this.Type = typeVar
	return &this
}

// NewCreatePhoneNotificationChannelConfigWithDefaults instantiates a new CreatePhoneNotificationChannelConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreatePhoneNotificationChannelConfigWithDefaults() *CreatePhoneNotificationChannelConfig {
	this := CreatePhoneNotificationChannelConfig{}
	var typeVar NotificationChannelPhoneConfigType = NOTIFICATIONCHANNELPHONECONFIGTYPE_PHONE
	this.Type = typeVar
	return &this
}

// GetNumber returns the Number field value.
func (o *CreatePhoneNotificationChannelConfig) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *CreatePhoneNotificationChannelConfig) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value.
func (o *CreatePhoneNotificationChannelConfig) SetNumber(v string) {
	o.Number = v
}

// GetType returns the Type field value.
func (o *CreatePhoneNotificationChannelConfig) GetType() NotificationChannelPhoneConfigType {
	if o == nil {
		var ret NotificationChannelPhoneConfigType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreatePhoneNotificationChannelConfig) GetTypeOk() (*NotificationChannelPhoneConfigType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreatePhoneNotificationChannelConfig) SetType(v NotificationChannelPhoneConfigType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreatePhoneNotificationChannelConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["number"] = o.Number
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreatePhoneNotificationChannelConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Number *string                             `json:"number"`
		Type   *NotificationChannelPhoneConfigType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Number == nil {
		return fmt.Errorf("required field number missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"number", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Number = *all.Number
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
