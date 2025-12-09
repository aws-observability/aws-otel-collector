// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RulesValidateQueryResponseDataAttributes The definition of `RulesValidateQueryResponseDataAttributes` object.
type RulesValidateQueryResponseDataAttributes struct {
	// The `attributes` `Canonical`.
	Canonical string `json:"Canonical"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRulesValidateQueryResponseDataAttributes instantiates a new RulesValidateQueryResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRulesValidateQueryResponseDataAttributes(canonical string) *RulesValidateQueryResponseDataAttributes {
	this := RulesValidateQueryResponseDataAttributes{}
	this.Canonical = canonical
	return &this
}

// NewRulesValidateQueryResponseDataAttributesWithDefaults instantiates a new RulesValidateQueryResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRulesValidateQueryResponseDataAttributesWithDefaults() *RulesValidateQueryResponseDataAttributes {
	this := RulesValidateQueryResponseDataAttributes{}
	return &this
}

// GetCanonical returns the Canonical field value.
func (o *RulesValidateQueryResponseDataAttributes) GetCanonical() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *RulesValidateQueryResponseDataAttributes) GetCanonicalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value.
func (o *RulesValidateQueryResponseDataAttributes) SetCanonical(v string) {
	o.Canonical = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RulesValidateQueryResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["Canonical"] = o.Canonical

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RulesValidateQueryResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Canonical *string `json:"Canonical"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Canonical == nil {
		return fmt.Errorf("required field Canonical missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"Canonical"})
	} else {
		return err
	}
	o.Canonical = *all.Canonical

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
