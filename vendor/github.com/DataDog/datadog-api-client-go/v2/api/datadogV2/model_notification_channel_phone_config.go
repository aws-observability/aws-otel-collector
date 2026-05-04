// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationChannelPhoneConfig Phone notification channel configuration
type NotificationChannelPhoneConfig struct {
	// The formatted international version of Number (e.g. +33 7 1 23 45 67).
	FormattedNumber string `json:"formatted_number"`
	// The E-164 formatted phone number (e.g. +3371234567)
	Number string `json:"number"`
	// The ISO 3166-1 alpha-2 two-letter country code.
	Region string `json:"region"`
	// If present, the date the user subscribed this number to SMS messages
	SmsSubscribedAt datadog.NullableTime `json:"sms_subscribed_at,omitempty"`
	// Indicates that the notification channel is a phone
	Type NotificationChannelPhoneConfigType `json:"type"`
	// Indicates whether this phone has been verified by the user in Datadog On-Call
	Verified bool `json:"verified"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotificationChannelPhoneConfig instantiates a new NotificationChannelPhoneConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotificationChannelPhoneConfig(formattedNumber string, number string, region string, typeVar NotificationChannelPhoneConfigType, verified bool) *NotificationChannelPhoneConfig {
	this := NotificationChannelPhoneConfig{}
	this.FormattedNumber = formattedNumber
	this.Number = number
	this.Region = region
	this.Type = typeVar
	this.Verified = verified
	return &this
}

// NewNotificationChannelPhoneConfigWithDefaults instantiates a new NotificationChannelPhoneConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotificationChannelPhoneConfigWithDefaults() *NotificationChannelPhoneConfig {
	this := NotificationChannelPhoneConfig{}
	var typeVar NotificationChannelPhoneConfigType = NOTIFICATIONCHANNELPHONECONFIGTYPE_PHONE
	this.Type = typeVar
	return &this
}

// GetFormattedNumber returns the FormattedNumber field value.
func (o *NotificationChannelPhoneConfig) GetFormattedNumber() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FormattedNumber
}

// GetFormattedNumberOk returns a tuple with the FormattedNumber field value
// and a boolean to check if the value has been set.
func (o *NotificationChannelPhoneConfig) GetFormattedNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormattedNumber, true
}

// SetFormattedNumber sets field value.
func (o *NotificationChannelPhoneConfig) SetFormattedNumber(v string) {
	o.FormattedNumber = v
}

// GetNumber returns the Number field value.
func (o *NotificationChannelPhoneConfig) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *NotificationChannelPhoneConfig) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value.
func (o *NotificationChannelPhoneConfig) SetNumber(v string) {
	o.Number = v
}

// GetRegion returns the Region field value.
func (o *NotificationChannelPhoneConfig) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *NotificationChannelPhoneConfig) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *NotificationChannelPhoneConfig) SetRegion(v string) {
	o.Region = v
}

// GetSmsSubscribedAt returns the SmsSubscribedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationChannelPhoneConfig) GetSmsSubscribedAt() time.Time {
	if o == nil || o.SmsSubscribedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SmsSubscribedAt.Get()
}

// GetSmsSubscribedAtOk returns a tuple with the SmsSubscribedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NotificationChannelPhoneConfig) GetSmsSubscribedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmsSubscribedAt.Get(), o.SmsSubscribedAt.IsSet()
}

// HasSmsSubscribedAt returns a boolean if a field has been set.
func (o *NotificationChannelPhoneConfig) HasSmsSubscribedAt() bool {
	return o != nil && o.SmsSubscribedAt.IsSet()
}

// SetSmsSubscribedAt gets a reference to the given datadog.NullableTime and assigns it to the SmsSubscribedAt field.
func (o *NotificationChannelPhoneConfig) SetSmsSubscribedAt(v time.Time) {
	o.SmsSubscribedAt.Set(&v)
}

// SetSmsSubscribedAtNil sets the value for SmsSubscribedAt to be an explicit nil.
func (o *NotificationChannelPhoneConfig) SetSmsSubscribedAtNil() {
	o.SmsSubscribedAt.Set(nil)
}

// UnsetSmsSubscribedAt ensures that no value is present for SmsSubscribedAt, not even an explicit nil.
func (o *NotificationChannelPhoneConfig) UnsetSmsSubscribedAt() {
	o.SmsSubscribedAt.Unset()
}

// GetType returns the Type field value.
func (o *NotificationChannelPhoneConfig) GetType() NotificationChannelPhoneConfigType {
	if o == nil {
		var ret NotificationChannelPhoneConfigType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationChannelPhoneConfig) GetTypeOk() (*NotificationChannelPhoneConfigType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NotificationChannelPhoneConfig) SetType(v NotificationChannelPhoneConfigType) {
	o.Type = v
}

// GetVerified returns the Verified field value.
func (o *NotificationChannelPhoneConfig) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *NotificationChannelPhoneConfig) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value.
func (o *NotificationChannelPhoneConfig) SetVerified(v bool) {
	o.Verified = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotificationChannelPhoneConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["formatted_number"] = o.FormattedNumber
	toSerialize["number"] = o.Number
	toSerialize["region"] = o.Region
	if o.SmsSubscribedAt.IsSet() {
		toSerialize["sms_subscribed_at"] = o.SmsSubscribedAt.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["verified"] = o.Verified

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotificationChannelPhoneConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FormattedNumber *string                             `json:"formatted_number"`
		Number          *string                             `json:"number"`
		Region          *string                             `json:"region"`
		SmsSubscribedAt datadog.NullableTime                `json:"sms_subscribed_at,omitempty"`
		Type            *NotificationChannelPhoneConfigType `json:"type"`
		Verified        *bool                               `json:"verified"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FormattedNumber == nil {
		return fmt.Errorf("required field formatted_number missing")
	}
	if all.Number == nil {
		return fmt.Errorf("required field number missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Verified == nil {
		return fmt.Errorf("required field verified missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"formatted_number", "number", "region", "sms_subscribed_at", "type", "verified"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FormattedNumber = *all.FormattedNumber
	o.Number = *all.Number
	o.Region = *all.Region
	o.SmsSubscribedAt = all.SmsSubscribedAt
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Verified = *all.Verified

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
