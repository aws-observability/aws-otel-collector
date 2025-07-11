// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberFormatUnitCanonical Canonical unit.
type NumberFormatUnitCanonical struct {
	// The name of the unit per item.
	PerUnitName *string `json:"per_unit_name,omitempty"`
	// The type of unit scale.
	Type *NumberFormatUnitScaleType `json:"type,omitempty"`
	// The name of the unit.
	UnitName *string `json:"unit_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNumberFormatUnitCanonical instantiates a new NumberFormatUnitCanonical object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNumberFormatUnitCanonical() *NumberFormatUnitCanonical {
	this := NumberFormatUnitCanonical{}
	return &this
}

// NewNumberFormatUnitCanonicalWithDefaults instantiates a new NumberFormatUnitCanonical object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNumberFormatUnitCanonicalWithDefaults() *NumberFormatUnitCanonical {
	this := NumberFormatUnitCanonical{}
	return &this
}

// GetPerUnitName returns the PerUnitName field value if set, zero value otherwise.
func (o *NumberFormatUnitCanonical) GetPerUnitName() string {
	if o == nil || o.PerUnitName == nil {
		var ret string
		return ret
	}
	return *o.PerUnitName
}

// GetPerUnitNameOk returns a tuple with the PerUnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormatUnitCanonical) GetPerUnitNameOk() (*string, bool) {
	if o == nil || o.PerUnitName == nil {
		return nil, false
	}
	return o.PerUnitName, true
}

// HasPerUnitName returns a boolean if a field has been set.
func (o *NumberFormatUnitCanonical) HasPerUnitName() bool {
	return o != nil && o.PerUnitName != nil
}

// SetPerUnitName gets a reference to the given string and assigns it to the PerUnitName field.
func (o *NumberFormatUnitCanonical) SetPerUnitName(v string) {
	o.PerUnitName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NumberFormatUnitCanonical) GetType() NumberFormatUnitScaleType {
	if o == nil || o.Type == nil {
		var ret NumberFormatUnitScaleType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormatUnitCanonical) GetTypeOk() (*NumberFormatUnitScaleType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NumberFormatUnitCanonical) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given NumberFormatUnitScaleType and assigns it to the Type field.
func (o *NumberFormatUnitCanonical) SetType(v NumberFormatUnitScaleType) {
	o.Type = &v
}

// GetUnitName returns the UnitName field value if set, zero value otherwise.
func (o *NumberFormatUnitCanonical) GetUnitName() string {
	if o == nil || o.UnitName == nil {
		var ret string
		return ret
	}
	return *o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormatUnitCanonical) GetUnitNameOk() (*string, bool) {
	if o == nil || o.UnitName == nil {
		return nil, false
	}
	return o.UnitName, true
}

// HasUnitName returns a boolean if a field has been set.
func (o *NumberFormatUnitCanonical) HasUnitName() bool {
	return o != nil && o.UnitName != nil
}

// SetUnitName gets a reference to the given string and assigns it to the UnitName field.
func (o *NumberFormatUnitCanonical) SetUnitName(v string) {
	o.UnitName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NumberFormatUnitCanonical) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PerUnitName != nil {
		toSerialize["per_unit_name"] = o.PerUnitName
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
func (o *NumberFormatUnitCanonical) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PerUnitName *string                    `json:"per_unit_name,omitempty"`
		Type        *NumberFormatUnitScaleType `json:"type,omitempty"`
		UnitName    *string                    `json:"unit_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"per_unit_name", "type", "unit_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.PerUnitName = all.PerUnitName
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
