// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RulesValidateQueryResponseData The definition of `RulesValidateQueryResponseData` object.
type RulesValidateQueryResponseData struct {
	// The definition of `RulesValidateQueryResponseDataAttributes` object.
	Attributes *RulesValidateQueryResponseDataAttributes `json:"attributes,omitempty"`
	// The `RulesValidateQueryResponseData` `id`.
	Id *string `json:"id,omitempty"`
	// Validate response resource type.
	Type RulesValidateQueryResponseDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRulesValidateQueryResponseData instantiates a new RulesValidateQueryResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRulesValidateQueryResponseData(typeVar RulesValidateQueryResponseDataType) *RulesValidateQueryResponseData {
	this := RulesValidateQueryResponseData{}
	this.Type = typeVar
	return &this
}

// NewRulesValidateQueryResponseDataWithDefaults instantiates a new RulesValidateQueryResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRulesValidateQueryResponseDataWithDefaults() *RulesValidateQueryResponseData {
	this := RulesValidateQueryResponseData{}
	var typeVar RulesValidateQueryResponseDataType = RULESVALIDATEQUERYRESPONSEDATATYPE_VALIDATE_RESPONSE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RulesValidateQueryResponseData) GetAttributes() RulesValidateQueryResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret RulesValidateQueryResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesValidateQueryResponseData) GetAttributesOk() (*RulesValidateQueryResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RulesValidateQueryResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given RulesValidateQueryResponseDataAttributes and assigns it to the Attributes field.
func (o *RulesValidateQueryResponseData) SetAttributes(v RulesValidateQueryResponseDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RulesValidateQueryResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesValidateQueryResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RulesValidateQueryResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RulesValidateQueryResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *RulesValidateQueryResponseData) GetType() RulesValidateQueryResponseDataType {
	if o == nil {
		var ret RulesValidateQueryResponseDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RulesValidateQueryResponseData) GetTypeOk() (*RulesValidateQueryResponseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RulesValidateQueryResponseData) SetType(v RulesValidateQueryResponseDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RulesValidateQueryResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RulesValidateQueryResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *RulesValidateQueryResponseDataAttributes `json:"attributes,omitempty"`
		Id         *string                                   `json:"id,omitempty"`
		Type       *RulesValidateQueryResponseDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
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
