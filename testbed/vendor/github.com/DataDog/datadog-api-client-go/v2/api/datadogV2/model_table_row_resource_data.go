// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableRowResourceData The data object containing the row column names and values.
type TableRowResourceData struct {
	// Column values for this row in the reference table.
	Attributes *TableRowResourceDataAttributes `json:"attributes,omitempty"`
	// Row identifier, corresponding to the primary key value.
	Id *string `json:"id,omitempty"`
	// Row resource type.
	Type TableRowResourceDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableRowResourceData instantiates a new TableRowResourceData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableRowResourceData(typeVar TableRowResourceDataType) *TableRowResourceData {
	this := TableRowResourceData{}
	this.Type = typeVar
	return &this
}

// NewTableRowResourceDataWithDefaults instantiates a new TableRowResourceData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableRowResourceDataWithDefaults() *TableRowResourceData {
	this := TableRowResourceData{}
	var typeVar TableRowResourceDataType = TABLEROWRESOURCEDATATYPE_ROW
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TableRowResourceData) GetAttributes() TableRowResourceDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TableRowResourceDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableRowResourceData) GetAttributesOk() (*TableRowResourceDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TableRowResourceData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given TableRowResourceDataAttributes and assigns it to the Attributes field.
func (o *TableRowResourceData) SetAttributes(v TableRowResourceDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TableRowResourceData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableRowResourceData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TableRowResourceData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TableRowResourceData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *TableRowResourceData) GetType() TableRowResourceDataType {
	if o == nil {
		var ret TableRowResourceDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TableRowResourceData) GetTypeOk() (*TableRowResourceDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TableRowResourceData) SetType(v TableRowResourceDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableRowResourceData) MarshalJSON() ([]byte, error) {
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
func (o *TableRowResourceData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TableRowResourceDataAttributes `json:"attributes,omitempty"`
		Id         *string                         `json:"id,omitempty"`
		Type       *TableRowResourceDataType       `json:"type"`
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
