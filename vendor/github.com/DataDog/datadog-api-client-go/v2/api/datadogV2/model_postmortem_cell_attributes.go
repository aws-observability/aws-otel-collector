// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PostmortemCellAttributes Attributes of a postmortem cell
type PostmortemCellAttributes struct {
	// Definition of a postmortem cell
	Definition *PostmortemCellDefinition `json:"definition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostmortemCellAttributes instantiates a new PostmortemCellAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostmortemCellAttributes() *PostmortemCellAttributes {
	this := PostmortemCellAttributes{}
	return &this
}

// NewPostmortemCellAttributesWithDefaults instantiates a new PostmortemCellAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostmortemCellAttributesWithDefaults() *PostmortemCellAttributes {
	this := PostmortemCellAttributes{}
	return &this
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *PostmortemCellAttributes) GetDefinition() PostmortemCellDefinition {
	if o == nil || o.Definition == nil {
		var ret PostmortemCellDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemCellAttributes) GetDefinitionOk() (*PostmortemCellDefinition, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *PostmortemCellAttributes) HasDefinition() bool {
	return o != nil && o.Definition != nil
}

// SetDefinition gets a reference to the given PostmortemCellDefinition and assigns it to the Definition field.
func (o *PostmortemCellAttributes) SetDefinition(v PostmortemCellDefinition) {
	o.Definition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostmortemCellAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostmortemCellAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Definition *PostmortemCellDefinition `json:"definition,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"definition"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Definition != nil && all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = all.Definition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
