// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UCConfigPairDataAttributes The definition of `UCConfigPairDataAttributes` object.
type UCConfigPairDataAttributes struct {
	// The `attributes` `configs`.
	Configs []UCConfigPairDataAttributesConfigsItems `json:"configs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUCConfigPairDataAttributes instantiates a new UCConfigPairDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUCConfigPairDataAttributes() *UCConfigPairDataAttributes {
	this := UCConfigPairDataAttributes{}
	return &this
}

// NewUCConfigPairDataAttributesWithDefaults instantiates a new UCConfigPairDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUCConfigPairDataAttributesWithDefaults() *UCConfigPairDataAttributes {
	this := UCConfigPairDataAttributes{}
	return &this
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributes) GetConfigs() []UCConfigPairDataAttributesConfigsItems {
	if o == nil || o.Configs == nil {
		var ret []UCConfigPairDataAttributesConfigsItems
		return ret
	}
	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributes) GetConfigsOk() (*[]UCConfigPairDataAttributesConfigsItems, bool) {
	if o == nil || o.Configs == nil {
		return nil, false
	}
	return &o.Configs, true
}

// HasConfigs returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributes) HasConfigs() bool {
	return o != nil && o.Configs != nil
}

// SetConfigs gets a reference to the given []UCConfigPairDataAttributesConfigsItems and assigns it to the Configs field.
func (o *UCConfigPairDataAttributes) SetConfigs(v []UCConfigPairDataAttributesConfigsItems) {
	o.Configs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UCConfigPairDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Configs != nil {
		toSerialize["configs"] = o.Configs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UCConfigPairDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Configs []UCConfigPairDataAttributesConfigsItems `json:"configs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configs"})
	} else {
		return err
	}
	o.Configs = all.Configs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
