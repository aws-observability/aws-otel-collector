// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NullableUserRelationshipData Relationship to user object.
type NullableUserRelationshipData struct {
	// A unique identifier that represents the user.
	Id string `json:"id"`
	// User resource type.
	Type UserResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNullableUserRelationshipData instantiates a new NullableUserRelationshipData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNullableUserRelationshipData(id string, typeVar UserResourceType) *NullableUserRelationshipData {
	this := NullableUserRelationshipData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewNullableUserRelationshipDataWithDefaults instantiates a new NullableUserRelationshipData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNullableUserRelationshipDataWithDefaults() *NullableUserRelationshipData {
	this := NullableUserRelationshipData{}
	var typeVar UserResourceType = USERRESOURCETYPE_USER
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *NullableUserRelationshipData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NullableUserRelationshipData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *NullableUserRelationshipData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *NullableUserRelationshipData) GetType() UserResourceType {
	if o == nil {
		var ret UserResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NullableUserRelationshipData) GetTypeOk() (*UserResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NullableUserRelationshipData) SetType(v UserResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NullableUserRelationshipData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NullableUserRelationshipData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string           `json:"id"`
		Type *UserResourceType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
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

// NullableNullableUserRelationshipData handles when a null is used for NullableUserRelationshipData.
type NullableNullableUserRelationshipData struct {
	value *NullableUserRelationshipData
	isSet bool
}

// Get returns the associated value.
func (v NullableNullableUserRelationshipData) Get() *NullableUserRelationshipData {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableNullableUserRelationshipData) Set(val *NullableUserRelationshipData) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableNullableUserRelationshipData) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableNullableUserRelationshipData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNullableUserRelationshipData initializes the struct as if Set has been called.
func NewNullableNullableUserRelationshipData(val *NullableUserRelationshipData) *NullableNullableUserRelationshipData {
	return &NullableNullableUserRelationshipData{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableNullableUserRelationshipData) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableNullableUserRelationshipData) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
