// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchTableRequestData The data object containing the partial table definition updates.
type PatchTableRequestData struct {
	// Attributes that define the updates to the reference table's configuration and properties.
	Attributes *PatchTableRequestDataAttributes `json:"attributes,omitempty"`
	// Reference table resource type.
	Type PatchTableRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewPatchTableRequestData instantiates a new PatchTableRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchTableRequestData(typeVar PatchTableRequestDataType) *PatchTableRequestData {
	this := PatchTableRequestData{}
	this.Type = typeVar
	return &this
}

// NewPatchTableRequestDataWithDefaults instantiates a new PatchTableRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchTableRequestDataWithDefaults() *PatchTableRequestData {
	this := PatchTableRequestData{}
	var typeVar PatchTableRequestDataType = PATCHTABLEREQUESTDATATYPE_REFERENCE_TABLE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PatchTableRequestData) GetAttributes() PatchTableRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret PatchTableRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestData) GetAttributesOk() (*PatchTableRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PatchTableRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given PatchTableRequestDataAttributes and assigns it to the Attributes field.
func (o *PatchTableRequestData) SetAttributes(v PatchTableRequestDataAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *PatchTableRequestData) GetType() PatchTableRequestDataType {
	if o == nil {
		var ret PatchTableRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PatchTableRequestData) GetTypeOk() (*PatchTableRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PatchTableRequestData) SetType(v PatchTableRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchTableRequestData) MarshalJSON() ([]byte, error) {
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
func (o *PatchTableRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *PatchTableRequestDataAttributes `json:"attributes,omitempty"`
		Type       *PatchTableRequestDataType       `json:"type"`
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
