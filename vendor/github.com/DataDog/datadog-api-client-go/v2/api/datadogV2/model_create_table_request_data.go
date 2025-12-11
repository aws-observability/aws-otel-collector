// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTableRequestData The data object containing the table definition.
type CreateTableRequestData struct {
	// Attributes that define the reference table's configuration and properties.
	Attributes *CreateTableRequestDataAttributes `json:"attributes,omitempty"`
	// Reference table resource type.
	Type CreateTableRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCreateTableRequestData instantiates a new CreateTableRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTableRequestData(typeVar CreateTableRequestDataType) *CreateTableRequestData {
	this := CreateTableRequestData{}
	this.Type = typeVar
	return &this
}

// NewCreateTableRequestDataWithDefaults instantiates a new CreateTableRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTableRequestDataWithDefaults() *CreateTableRequestData {
	this := CreateTableRequestData{}
	var typeVar CreateTableRequestDataType = CREATETABLEREQUESTDATATYPE_REFERENCE_TABLE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CreateTableRequestData) GetAttributes() CreateTableRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret CreateTableRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTableRequestData) GetAttributesOk() (*CreateTableRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CreateTableRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CreateTableRequestDataAttributes and assigns it to the Attributes field.
func (o *CreateTableRequestData) SetAttributes(v CreateTableRequestDataAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *CreateTableRequestData) GetType() CreateTableRequestDataType {
	if o == nil {
		var ret CreateTableRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestData) GetTypeOk() (*CreateTableRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateTableRequestData) SetType(v CreateTableRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTableRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateTableRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CreateTableRequestDataAttributes `json:"attributes,omitempty"`
		Type       *CreateTableRequestDataType       `json:"type"`
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
