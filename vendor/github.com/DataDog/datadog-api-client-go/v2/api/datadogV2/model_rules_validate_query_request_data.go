// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RulesValidateQueryRequestData The definition of `RulesValidateQueryRequestData` object.
type RulesValidateQueryRequestData struct {
	// The definition of `RulesValidateQueryRequestDataAttributes` object.
	Attributes *RulesValidateQueryRequestDataAttributes `json:"attributes,omitempty"`
	// The `RulesValidateQueryRequestData` `id`.
	Id *string `json:"id,omitempty"`
	// Validate query resource type.
	Type RulesValidateQueryRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRulesValidateQueryRequestData instantiates a new RulesValidateQueryRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRulesValidateQueryRequestData(typeVar RulesValidateQueryRequestDataType) *RulesValidateQueryRequestData {
	this := RulesValidateQueryRequestData{}
	this.Type = typeVar
	return &this
}

// NewRulesValidateQueryRequestDataWithDefaults instantiates a new RulesValidateQueryRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRulesValidateQueryRequestDataWithDefaults() *RulesValidateQueryRequestData {
	this := RulesValidateQueryRequestData{}
	var typeVar RulesValidateQueryRequestDataType = RULESVALIDATEQUERYREQUESTDATATYPE_VALIDATE_QUERY
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RulesValidateQueryRequestData) GetAttributes() RulesValidateQueryRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret RulesValidateQueryRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesValidateQueryRequestData) GetAttributesOk() (*RulesValidateQueryRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RulesValidateQueryRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given RulesValidateQueryRequestDataAttributes and assigns it to the Attributes field.
func (o *RulesValidateQueryRequestData) SetAttributes(v RulesValidateQueryRequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RulesValidateQueryRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesValidateQueryRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RulesValidateQueryRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RulesValidateQueryRequestData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *RulesValidateQueryRequestData) GetType() RulesValidateQueryRequestDataType {
	if o == nil {
		var ret RulesValidateQueryRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RulesValidateQueryRequestData) GetTypeOk() (*RulesValidateQueryRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RulesValidateQueryRequestData) SetType(v RulesValidateQueryRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RulesValidateQueryRequestData) MarshalJSON() ([]byte, error) {
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
func (o *RulesValidateQueryRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *RulesValidateQueryRequestDataAttributes `json:"attributes,omitempty"`
		Id         *string                                  `json:"id,omitempty"`
		Type       *RulesValidateQueryRequestDataType       `json:"type"`
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
