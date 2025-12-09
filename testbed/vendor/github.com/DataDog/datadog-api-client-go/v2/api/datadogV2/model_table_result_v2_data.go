// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableResultV2Data The data object containing the reference table configuration and state.
type TableResultV2Data struct {
	// Attributes that define the reference table's configuration and properties.
	Attributes *TableResultV2DataAttributes `json:"attributes,omitempty"`
	// Unique identifier for the reference table.
	Id *string `json:"id,omitempty"`
	// Reference table resource type.
	Type TableResultV2DataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableResultV2Data instantiates a new TableResultV2Data object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableResultV2Data(typeVar TableResultV2DataType) *TableResultV2Data {
	this := TableResultV2Data{}
	this.Type = typeVar
	return &this
}

// NewTableResultV2DataWithDefaults instantiates a new TableResultV2Data object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableResultV2DataWithDefaults() *TableResultV2Data {
	this := TableResultV2Data{}
	var typeVar TableResultV2DataType = TABLERESULTV2DATATYPE_REFERENCE_TABLE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TableResultV2Data) GetAttributes() TableResultV2DataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TableResultV2DataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2Data) GetAttributesOk() (*TableResultV2DataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TableResultV2Data) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given TableResultV2DataAttributes and assigns it to the Attributes field.
func (o *TableResultV2Data) SetAttributes(v TableResultV2DataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TableResultV2Data) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2Data) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TableResultV2Data) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TableResultV2Data) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *TableResultV2Data) GetType() TableResultV2DataType {
	if o == nil {
		var ret TableResultV2DataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TableResultV2Data) GetTypeOk() (*TableResultV2DataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TableResultV2Data) SetType(v TableResultV2DataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableResultV2Data) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableResultV2Data) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TableResultV2DataAttributes `json:"attributes,omitempty"`
		Id         *string                      `json:"id,omitempty"`
		Type       *TableResultV2DataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
