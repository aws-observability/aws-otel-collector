// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateNotificationChannelAttributes Attributes for creating an on-call notification channel.
type CreateNotificationChannelAttributes struct {
	// Defines the configuration for creating an On-Call notification channel
	Config *CreateNotificationChannelConfig `json:"config,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateNotificationChannelAttributes instantiates a new CreateNotificationChannelAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateNotificationChannelAttributes() *CreateNotificationChannelAttributes {
	this := CreateNotificationChannelAttributes{}
	return &this
}

// NewCreateNotificationChannelAttributesWithDefaults instantiates a new CreateNotificationChannelAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateNotificationChannelAttributesWithDefaults() *CreateNotificationChannelAttributes {
	this := CreateNotificationChannelAttributes{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CreateNotificationChannelAttributes) GetConfig() CreateNotificationChannelConfig {
	if o == nil || o.Config == nil {
		var ret CreateNotificationChannelConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationChannelAttributes) GetConfigOk() (*CreateNotificationChannelConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CreateNotificationChannelAttributes) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given CreateNotificationChannelConfig and assigns it to the Config field.
func (o *CreateNotificationChannelAttributes) SetConfig(v CreateNotificationChannelConfig) {
	o.Config = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateNotificationChannelAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *CreateNotificationChannelAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config *CreateNotificationChannelConfig `json:"config,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config"})
	} else {
		return err
	}
	o.Config = all.Config

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
