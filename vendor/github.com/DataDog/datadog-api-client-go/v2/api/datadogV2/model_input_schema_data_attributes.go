// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InputSchemaDataAttributes The definition of `InputSchemaDataAttributes` object.
type InputSchemaDataAttributes struct {
	// The `attributes` `parameters`.
	Parameters []InputSchemaDataAttributesParametersItems `json:"parameters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInputSchemaDataAttributes instantiates a new InputSchemaDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInputSchemaDataAttributes() *InputSchemaDataAttributes {
	this := InputSchemaDataAttributes{}
	return &this
}

// NewInputSchemaDataAttributesWithDefaults instantiates a new InputSchemaDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInputSchemaDataAttributesWithDefaults() *InputSchemaDataAttributes {
	this := InputSchemaDataAttributes{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *InputSchemaDataAttributes) GetParameters() []InputSchemaDataAttributesParametersItems {
	if o == nil || o.Parameters == nil {
		var ret []InputSchemaDataAttributesParametersItems
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSchemaDataAttributes) GetParametersOk() (*[]InputSchemaDataAttributesParametersItems, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *InputSchemaDataAttributes) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given []InputSchemaDataAttributesParametersItems and assigns it to the Parameters field.
func (o *InputSchemaDataAttributes) SetParameters(v []InputSchemaDataAttributesParametersItems) {
	o.Parameters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InputSchemaDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InputSchemaDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Parameters []InputSchemaDataAttributesParametersItems `json:"parameters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"parameters"})
	} else {
		return err
	}
	o.Parameters = all.Parameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
