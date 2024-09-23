// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateOpenAPIResponseData Data envelope for `CreateOpenAPIResponse`.
type CreateOpenAPIResponseData struct {
	// Attributes for `CreateOpenAPI`.
	Attributes *CreateOpenAPIResponseAttributes `json:"attributes,omitempty"`
	// API identifier.
	Id *uuid.UUID `json:"id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateOpenAPIResponseData instantiates a new CreateOpenAPIResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateOpenAPIResponseData() *CreateOpenAPIResponseData {
	this := CreateOpenAPIResponseData{}
	return &this
}

// NewCreateOpenAPIResponseDataWithDefaults instantiates a new CreateOpenAPIResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateOpenAPIResponseDataWithDefaults() *CreateOpenAPIResponseData {
	this := CreateOpenAPIResponseData{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CreateOpenAPIResponseData) GetAttributes() CreateOpenAPIResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret CreateOpenAPIResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOpenAPIResponseData) GetAttributesOk() (*CreateOpenAPIResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CreateOpenAPIResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CreateOpenAPIResponseAttributes and assigns it to the Attributes field.
func (o *CreateOpenAPIResponseData) SetAttributes(v CreateOpenAPIResponseAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateOpenAPIResponseData) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOpenAPIResponseData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateOpenAPIResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *CreateOpenAPIResponseData) SetId(v uuid.UUID) {
	o.Id = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateOpenAPIResponseData) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateOpenAPIResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CreateOpenAPIResponseAttributes `json:"attributes,omitempty"`
		Id         *uuid.UUID                       `json:"id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
