// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateUploadRequestData Request data for creating an upload for a file to be ingested into a reference table.
type CreateUploadRequestData struct {
	// Upload configuration specifying how data is uploaded by the user, and properties of the table to associate the upload with.
	Attributes *CreateUploadRequestDataAttributes `json:"attributes,omitempty"`
	// Upload resource type.
	Type CreateUploadRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCreateUploadRequestData instantiates a new CreateUploadRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateUploadRequestData(typeVar CreateUploadRequestDataType) *CreateUploadRequestData {
	this := CreateUploadRequestData{}
	this.Type = typeVar
	return &this
}

// NewCreateUploadRequestDataWithDefaults instantiates a new CreateUploadRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateUploadRequestDataWithDefaults() *CreateUploadRequestData {
	this := CreateUploadRequestData{}
	var typeVar CreateUploadRequestDataType = CREATEUPLOADREQUESTDATATYPE_UPLOAD
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CreateUploadRequestData) GetAttributes() CreateUploadRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret CreateUploadRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUploadRequestData) GetAttributesOk() (*CreateUploadRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CreateUploadRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CreateUploadRequestDataAttributes and assigns it to the Attributes field.
func (o *CreateUploadRequestData) SetAttributes(v CreateUploadRequestDataAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *CreateUploadRequestData) GetType() CreateUploadRequestDataType {
	if o == nil {
		var ret CreateUploadRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateUploadRequestData) GetTypeOk() (*CreateUploadRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateUploadRequestData) SetType(v CreateUploadRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateUploadRequestData) MarshalJSON() ([]byte, error) {
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
func (o *CreateUploadRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CreateUploadRequestDataAttributes `json:"attributes,omitempty"`
		Type       *CreateUploadRequestDataType       `json:"type"`
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
