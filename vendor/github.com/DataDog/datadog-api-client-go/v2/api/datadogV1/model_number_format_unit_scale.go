// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberFormatUnitScale The definition of `NumberFormatUnitScale` object.
type NumberFormatUnitScale struct {
	// The type of unit scale.
	Type *NumberFormatUnitScaleType `json:"type,omitempty"`
	// The name of the unit.
	UnitName *string `json:"unit_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNumberFormatUnitScale instantiates a new NumberFormatUnitScale object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNumberFormatUnitScale() *NumberFormatUnitScale {
	this := NumberFormatUnitScale{}
	return &this
}

// NewNumberFormatUnitScaleWithDefaults instantiates a new NumberFormatUnitScale object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNumberFormatUnitScaleWithDefaults() *NumberFormatUnitScale {
	this := NumberFormatUnitScale{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NumberFormatUnitScale) GetType() NumberFormatUnitScaleType {
	if o == nil || o.Type == nil {
		var ret NumberFormatUnitScaleType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormatUnitScale) GetTypeOk() (*NumberFormatUnitScaleType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NumberFormatUnitScale) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given NumberFormatUnitScaleType and assigns it to the Type field.
func (o *NumberFormatUnitScale) SetType(v NumberFormatUnitScaleType) {
	o.Type = &v
}

// GetUnitName returns the UnitName field value if set, zero value otherwise.
func (o *NumberFormatUnitScale) GetUnitName() string {
	if o == nil || o.UnitName == nil {
		var ret string
		return ret
	}
	return *o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormatUnitScale) GetUnitNameOk() (*string, bool) {
	if o == nil || o.UnitName == nil {
		return nil, false
	}
	return o.UnitName, true
}

// HasUnitName returns a boolean if a field has been set.
func (o *NumberFormatUnitScale) HasUnitName() bool {
	return o != nil && o.UnitName != nil
}

// SetUnitName gets a reference to the given string and assigns it to the UnitName field.
func (o *NumberFormatUnitScale) SetUnitName(v string) {
	o.UnitName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NumberFormatUnitScale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UnitName != nil {
		toSerialize["unit_name"] = o.UnitName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NumberFormatUnitScale) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type     *NumberFormatUnitScaleType `json:"type,omitempty"`
		UnitName *string                    `json:"unit_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type", "unit_name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.UnitName = all.UnitName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableNumberFormatUnitScale handles when a null is used for NumberFormatUnitScale.
type NullableNumberFormatUnitScale struct {
	value *NumberFormatUnitScale
	isSet bool
}

// Get returns the associated value.
func (v NullableNumberFormatUnitScale) Get() *NumberFormatUnitScale {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableNumberFormatUnitScale) Set(val *NumberFormatUnitScale) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableNumberFormatUnitScale) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableNumberFormatUnitScale) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNumberFormatUnitScale initializes the struct as if Set has been called.
func NewNullableNumberFormatUnitScale(val *NumberFormatUnitScale) *NullableNumberFormatUnitScale {
	return &NullableNumberFormatUnitScale{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableNumberFormatUnitScale) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableNumberFormatUnitScale) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
