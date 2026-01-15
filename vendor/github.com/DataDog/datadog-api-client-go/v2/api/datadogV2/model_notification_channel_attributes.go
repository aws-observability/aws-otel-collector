// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationChannelAttributes Attributes for an on-call notification channel.
type NotificationChannelAttributes struct {
	// Whether the notification channel is currently active.
	Active *bool `json:"active,omitempty"`
	// Defines the configuration for an On-Call notification channel
	Config *NotificationChannelConfig `json:"config,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotificationChannelAttributes instantiates a new NotificationChannelAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotificationChannelAttributes() *NotificationChannelAttributes {
	this := NotificationChannelAttributes{}
	return &this
}

// NewNotificationChannelAttributesWithDefaults instantiates a new NotificationChannelAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotificationChannelAttributesWithDefaults() *NotificationChannelAttributes {
	this := NotificationChannelAttributes{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *NotificationChannelAttributes) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationChannelAttributes) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *NotificationChannelAttributes) HasActive() bool {
	return o != nil && o.Active != nil
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *NotificationChannelAttributes) SetActive(v bool) {
	o.Active = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *NotificationChannelAttributes) GetConfig() NotificationChannelConfig {
	if o == nil || o.Config == nil {
		var ret NotificationChannelConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationChannelAttributes) GetConfigOk() (*NotificationChannelConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *NotificationChannelAttributes) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given NotificationChannelConfig and assigns it to the Config field.
func (o *NotificationChannelAttributes) SetConfig(v NotificationChannelConfig) {
	o.Config = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotificationChannelAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotificationChannelAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Active *bool                      `json:"active,omitempty"`
		Config *NotificationChannelConfig `json:"config,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"active", "config"})
	} else {
		return err
	}
	o.Active = all.Active
	o.Config = all.Config

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
