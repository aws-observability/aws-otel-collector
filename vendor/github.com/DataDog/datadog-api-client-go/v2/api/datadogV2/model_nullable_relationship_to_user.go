// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NullableRelationshipToUser Relationship to user.
type NullableRelationshipToUser struct {
	// Relationship to user object.
	Data NullableNullableRelationshipToUserData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNullableRelationshipToUser instantiates a new NullableRelationshipToUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNullableRelationshipToUser(data NullableNullableRelationshipToUserData) *NullableRelationshipToUser {
	this := NullableRelationshipToUser{}
	this.Data = data
	return &this
}

// NewNullableRelationshipToUserWithDefaults instantiates a new NullableRelationshipToUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNullableRelationshipToUserWithDefaults() *NullableRelationshipToUser {
	this := NullableRelationshipToUser{}
	return &this
}

// GetData returns the Data field value.
// If the value is explicit nil, the zero value for NullableRelationshipToUserData will be returned.
func (o *NullableRelationshipToUser) GetData() NullableRelationshipToUserData {
	if o == nil || o.Data.Get() == nil {
		var ret NullableRelationshipToUserData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NullableRelationshipToUser) GetDataOk() (*NullableRelationshipToUserData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// SetData sets field value.
func (o *NullableRelationshipToUser) SetData(v NullableRelationshipToUserData) {
	o.Data.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o NullableRelationshipToUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NullableRelationshipToUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data NullableNullableRelationshipToUserData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Data.IsSet() {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableNullableRelationshipToUser handles when a null is used for NullableRelationshipToUser.
type NullableNullableRelationshipToUser struct {
	value *NullableRelationshipToUser
	isSet bool
}

// Get returns the associated value.
func (v NullableNullableRelationshipToUser) Get() *NullableRelationshipToUser {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableNullableRelationshipToUser) Set(val *NullableRelationshipToUser) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableNullableRelationshipToUser) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableNullableRelationshipToUser) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNullableRelationshipToUser initializes the struct as if Set has been called.
func NewNullableNullableRelationshipToUser(val *NullableRelationshipToUser) *NullableNullableRelationshipToUser {
	return &NullableNullableRelationshipToUser{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableNullableRelationshipToUser) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableNullableRelationshipToUser) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
