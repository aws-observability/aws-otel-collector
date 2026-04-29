// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectUpdate Project update.
type ProjectUpdate struct {
	// Project update attributes.
	Attributes *ProjectUpdateAttributes `json:"attributes,omitempty"`
	// Project resource type.
	Type ProjectResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProjectUpdate instantiates a new ProjectUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProjectUpdate(typeVar ProjectResourceType) *ProjectUpdate {
	this := ProjectUpdate{}
	this.Type = typeVar
	return &this
}

// NewProjectUpdateWithDefaults instantiates a new ProjectUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectUpdateWithDefaults() *ProjectUpdate {
	this := ProjectUpdate{}
	var typeVar ProjectResourceType = PROJECTRESOURCETYPE_PROJECT
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProjectUpdate) GetAttributes() ProjectUpdateAttributes {
	if o == nil || o.Attributes == nil {
		var ret ProjectUpdateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdate) GetAttributesOk() (*ProjectUpdateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProjectUpdate) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ProjectUpdateAttributes and assigns it to the Attributes field.
func (o *ProjectUpdate) SetAttributes(v ProjectUpdateAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *ProjectUpdate) GetType() ProjectResourceType {
	if o == nil {
		var ret ProjectResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProjectUpdate) GetTypeOk() (*ProjectResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProjectUpdate) SetType(v ProjectResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProjectUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProjectUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ProjectUpdateAttributes `json:"attributes,omitempty"`
		Type       *ProjectResourceType     `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
