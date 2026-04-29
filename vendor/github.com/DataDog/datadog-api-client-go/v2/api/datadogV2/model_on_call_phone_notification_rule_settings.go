// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPhoneNotificationRuleSettings Configuration for using a phone notification channel in a notification rule
type OnCallPhoneNotificationRuleSettings struct {
	// Specifies the method in which a phone is used in a notification rule
	Method OnCallPhoneNotificationRuleMethod `json:"method"`
	// Indicates that the notification channel is a phone
	Type NotificationChannelPhoneConfigType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPhoneNotificationRuleSettings instantiates a new OnCallPhoneNotificationRuleSettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPhoneNotificationRuleSettings(method OnCallPhoneNotificationRuleMethod, typeVar NotificationChannelPhoneConfigType) *OnCallPhoneNotificationRuleSettings {
	this := OnCallPhoneNotificationRuleSettings{}
	this.Method = method
	this.Type = typeVar
	return &this
}

// NewOnCallPhoneNotificationRuleSettingsWithDefaults instantiates a new OnCallPhoneNotificationRuleSettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPhoneNotificationRuleSettingsWithDefaults() *OnCallPhoneNotificationRuleSettings {
	this := OnCallPhoneNotificationRuleSettings{}
	var typeVar NotificationChannelPhoneConfigType = NOTIFICATIONCHANNELPHONECONFIGTYPE_PHONE
	this.Type = typeVar
	return &this
}

// GetMethod returns the Method field value.
func (o *OnCallPhoneNotificationRuleSettings) GetMethod() OnCallPhoneNotificationRuleMethod {
	if o == nil {
		var ret OnCallPhoneNotificationRuleMethod
		return ret
	}
	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *OnCallPhoneNotificationRuleSettings) GetMethodOk() (*OnCallPhoneNotificationRuleMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value.
func (o *OnCallPhoneNotificationRuleSettings) SetMethod(v OnCallPhoneNotificationRuleMethod) {
	o.Method = v
}

// GetType returns the Type field value.
func (o *OnCallPhoneNotificationRuleSettings) GetType() NotificationChannelPhoneConfigType {
	if o == nil {
		var ret NotificationChannelPhoneConfigType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnCallPhoneNotificationRuleSettings) GetTypeOk() (*NotificationChannelPhoneConfigType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OnCallPhoneNotificationRuleSettings) SetType(v NotificationChannelPhoneConfigType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPhoneNotificationRuleSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["method"] = o.Method
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPhoneNotificationRuleSettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Method *OnCallPhoneNotificationRuleMethod  `json:"method"`
		Type   *NotificationChannelPhoneConfigType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Method == nil {
		return fmt.Errorf("required field method missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"method", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Method.IsValid() {
		hasInvalidField = true
	} else {
		o.Method = *all.Method
	}
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
