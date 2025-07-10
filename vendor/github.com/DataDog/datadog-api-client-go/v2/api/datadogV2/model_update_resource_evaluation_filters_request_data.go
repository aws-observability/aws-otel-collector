// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateResourceEvaluationFiltersRequestData The definition of `UpdateResourceFilterRequestData` object.
type UpdateResourceEvaluationFiltersRequestData struct {
	// Attributes of a resource filter.
	Attributes ResourceFilterAttributes `json:"attributes"`
	// The `UpdateResourceEvaluationFiltersRequestData` `id`.
	Id *string `json:"id,omitempty"`
	// Constant string to identify the request type.
	Type ResourceFilterRequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateResourceEvaluationFiltersRequestData instantiates a new UpdateResourceEvaluationFiltersRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateResourceEvaluationFiltersRequestData(attributes ResourceFilterAttributes, typeVar ResourceFilterRequestType) *UpdateResourceEvaluationFiltersRequestData {
	this := UpdateResourceEvaluationFiltersRequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewUpdateResourceEvaluationFiltersRequestDataWithDefaults instantiates a new UpdateResourceEvaluationFiltersRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateResourceEvaluationFiltersRequestDataWithDefaults() *UpdateResourceEvaluationFiltersRequestData {
	this := UpdateResourceEvaluationFiltersRequestData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *UpdateResourceEvaluationFiltersRequestData) GetAttributes() ResourceFilterAttributes {
	if o == nil {
		var ret ResourceFilterAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceEvaluationFiltersRequestData) GetAttributesOk() (*ResourceFilterAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *UpdateResourceEvaluationFiltersRequestData) SetAttributes(v ResourceFilterAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateResourceEvaluationFiltersRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateResourceEvaluationFiltersRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateResourceEvaluationFiltersRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateResourceEvaluationFiltersRequestData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *UpdateResourceEvaluationFiltersRequestData) GetType() ResourceFilterRequestType {
	if o == nil {
		var ret ResourceFilterRequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceEvaluationFiltersRequestData) GetTypeOk() (*ResourceFilterRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UpdateResourceEvaluationFiltersRequestData) SetType(v ResourceFilterRequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateResourceEvaluationFiltersRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
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
func (o *UpdateResourceEvaluationFiltersRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ResourceFilterAttributes  `json:"attributes"`
		Id         *string                    `json:"id,omitempty"`
		Type       *ResourceFilterRequestType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
