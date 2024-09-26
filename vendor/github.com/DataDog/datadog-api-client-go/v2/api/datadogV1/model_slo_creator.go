// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOCreator The creator of the SLO
type SLOCreator struct {
	// Email of the creator.
	Email *string `json:"email,omitempty"`
	// User ID of the creator.
	Id *int64 `json:"id,omitempty"`
	// Name of the creator.
	Name datadog.NullableString `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSLOCreator instantiates a new SLOCreator object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOCreator() *SLOCreator {
	this := SLOCreator{}
	return &this
}

// NewSLOCreatorWithDefaults instantiates a new SLOCreator object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOCreatorWithDefaults() *SLOCreator {
	this := SLOCreator{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SLOCreator) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCreator) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SLOCreator) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SLOCreator) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SLOCreator) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCreator) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SLOCreator) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SLOCreator) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SLOCreator) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SLOCreator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SLOCreator) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *SLOCreator) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *SLOCreator) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *SLOCreator) UnsetName() {
	o.Name.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOCreator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOCreator) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Email *string                `json:"email,omitempty"`
		Id    *int64                 `json:"id,omitempty"`
		Name  datadog.NullableString `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"email", "id", "name"})
	} else {
		return err
	}
	o.Email = all.Email
	o.Id = all.Id
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableSLOCreator handles when a null is used for SLOCreator.
type NullableSLOCreator struct {
	value *SLOCreator
	isSet bool
}

// Get returns the associated value.
func (v NullableSLOCreator) Get() *SLOCreator {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableSLOCreator) Set(val *SLOCreator) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableSLOCreator) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableSLOCreator) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSLOCreator initializes the struct as if Set has been called.
func NewNullableSLOCreator(val *SLOCreator) *NullableSLOCreator {
	return &NullableSLOCreator{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableSLOCreator) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableSLOCreator) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
