// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMappingResponseDataAttributes
type GetMappingResponseDataAttributes struct {
	//
	Attributes []GetMappingResponseDataAttributesAttributesItems `json:"attributes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetMappingResponseDataAttributes instantiates a new GetMappingResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetMappingResponseDataAttributes() *GetMappingResponseDataAttributes {
	this := GetMappingResponseDataAttributes{}
	return &this
}

// NewGetMappingResponseDataAttributesWithDefaults instantiates a new GetMappingResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetMappingResponseDataAttributesWithDefaults() *GetMappingResponseDataAttributes {
	this := GetMappingResponseDataAttributes{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GetMappingResponseDataAttributes) GetAttributes() []GetMappingResponseDataAttributesAttributesItems {
	if o == nil || o.Attributes == nil {
		var ret []GetMappingResponseDataAttributesAttributesItems
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMappingResponseDataAttributes) GetAttributesOk() (*[]GetMappingResponseDataAttributesAttributesItems, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GetMappingResponseDataAttributes) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given []GetMappingResponseDataAttributesAttributesItems and assigns it to the Attributes field.
func (o *GetMappingResponseDataAttributes) SetAttributes(v []GetMappingResponseDataAttributesAttributesItems) {
	o.Attributes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetMappingResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetMappingResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes []GetMappingResponseDataAttributesAttributesItems `json:"attributes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes"})
	} else {
		return err
	}
	o.Attributes = all.Attributes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
