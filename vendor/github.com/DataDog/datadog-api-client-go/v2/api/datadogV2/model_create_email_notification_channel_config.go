// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateEmailNotificationChannelConfig Configuration to create an e-mail notification channel
type CreateEmailNotificationChannelConfig struct {
	// The e-mail address to be notified
	Address string `json:"address"`
	// Preferred content formats for notifications.
	Formats []NotificationChannelEmailFormatType `json:"formats"`
	// Indicates that the notification channel is an e-mail address
	Type NotificationChannelEmailConfigType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateEmailNotificationChannelConfig instantiates a new CreateEmailNotificationChannelConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateEmailNotificationChannelConfig(address string, formats []NotificationChannelEmailFormatType, typeVar NotificationChannelEmailConfigType) *CreateEmailNotificationChannelConfig {
	this := CreateEmailNotificationChannelConfig{}
	this.Address = address
	this.Formats = formats
	this.Type = typeVar
	return &this
}

// NewCreateEmailNotificationChannelConfigWithDefaults instantiates a new CreateEmailNotificationChannelConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateEmailNotificationChannelConfigWithDefaults() *CreateEmailNotificationChannelConfig {
	this := CreateEmailNotificationChannelConfig{}
	var typeVar NotificationChannelEmailConfigType = NOTIFICATIONCHANNELEMAILCONFIGTYPE_EMAIL
	this.Type = typeVar
	return &this
}

// GetAddress returns the Address field value.
func (o *CreateEmailNotificationChannelConfig) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateEmailNotificationChannelConfig) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value.
func (o *CreateEmailNotificationChannelConfig) SetAddress(v string) {
	o.Address = v
}

// GetFormats returns the Formats field value.
func (o *CreateEmailNotificationChannelConfig) GetFormats() []NotificationChannelEmailFormatType {
	if o == nil {
		var ret []NotificationChannelEmailFormatType
		return ret
	}
	return o.Formats
}

// GetFormatsOk returns a tuple with the Formats field value
// and a boolean to check if the value has been set.
func (o *CreateEmailNotificationChannelConfig) GetFormatsOk() (*[]NotificationChannelEmailFormatType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formats, true
}

// SetFormats sets field value.
func (o *CreateEmailNotificationChannelConfig) SetFormats(v []NotificationChannelEmailFormatType) {
	o.Formats = v
}

// GetType returns the Type field value.
func (o *CreateEmailNotificationChannelConfig) GetType() NotificationChannelEmailConfigType {
	if o == nil {
		var ret NotificationChannelEmailConfigType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateEmailNotificationChannelConfig) GetTypeOk() (*NotificationChannelEmailConfigType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateEmailNotificationChannelConfig) SetType(v NotificationChannelEmailConfigType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateEmailNotificationChannelConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["address"] = o.Address
	toSerialize["formats"] = o.Formats
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateEmailNotificationChannelConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Address *string                               `json:"address"`
		Formats *[]NotificationChannelEmailFormatType `json:"formats"`
		Type    *NotificationChannelEmailConfigType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Address == nil {
		return fmt.Errorf("required field address missing")
	}
	if all.Formats == nil {
		return fmt.Errorf("required field formats missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"address", "formats", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Address = *all.Address
	o.Formats = *all.Formats
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
