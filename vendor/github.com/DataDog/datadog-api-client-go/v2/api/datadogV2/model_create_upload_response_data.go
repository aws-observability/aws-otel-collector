// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateUploadResponseData Upload ID and attributes of the created upload.
type CreateUploadResponseData struct {
	// Pre-signed URLs for uploading parts of the file.
	Attributes *CreateUploadResponseDataAttributes `json:"attributes,omitempty"`
	// Unique identifier for this upload. Use this ID when creating the reference table.
	Id *string `json:"id,omitempty"`
	// Upload resource type.
	Type CreateUploadResponseDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCreateUploadResponseData instantiates a new CreateUploadResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateUploadResponseData(typeVar CreateUploadResponseDataType) *CreateUploadResponseData {
	this := CreateUploadResponseData{}
	this.Type = typeVar
	return &this
}

// NewCreateUploadResponseDataWithDefaults instantiates a new CreateUploadResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateUploadResponseDataWithDefaults() *CreateUploadResponseData {
	this := CreateUploadResponseData{}
	var typeVar CreateUploadResponseDataType = CREATEUPLOADRESPONSEDATATYPE_UPLOAD
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CreateUploadResponseData) GetAttributes() CreateUploadResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret CreateUploadResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUploadResponseData) GetAttributesOk() (*CreateUploadResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CreateUploadResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CreateUploadResponseDataAttributes and assigns it to the Attributes field.
func (o *CreateUploadResponseData) SetAttributes(v CreateUploadResponseDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateUploadResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUploadResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateUploadResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateUploadResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *CreateUploadResponseData) GetType() CreateUploadResponseDataType {
	if o == nil {
		var ret CreateUploadResponseDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateUploadResponseData) GetTypeOk() (*CreateUploadResponseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateUploadResponseData) SetType(v CreateUploadResponseDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateUploadResponseData) MarshalJSON() ([]byte, error) {
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
func (o *CreateUploadResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CreateUploadResponseDataAttributes `json:"attributes,omitempty"`
		Id         *string                             `json:"id,omitempty"`
		Type       *CreateUploadResponseDataType       `json:"type"`
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
