// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetNumberFormat Number format options for the widget.
type WidgetNumberFormat struct {
	// Number format unit.
	Unit *NumberFormatUnit `json:"unit,omitempty"`
	// The definition of `NumberFormatUnitScale` object.
	UnitScale NullableNumberFormatUnitScale `json:"unit_scale,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetNumberFormat instantiates a new WidgetNumberFormat object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetNumberFormat() *WidgetNumberFormat {
	this := WidgetNumberFormat{}
	return &this
}

// NewWidgetNumberFormatWithDefaults instantiates a new WidgetNumberFormat object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetNumberFormatWithDefaults() *WidgetNumberFormat {
	this := WidgetNumberFormat{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *WidgetNumberFormat) GetUnit() NumberFormatUnit {
	if o == nil || o.Unit == nil {
		var ret NumberFormatUnit
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetNumberFormat) GetUnitOk() (*NumberFormatUnit, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *WidgetNumberFormat) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given NumberFormatUnit and assigns it to the Unit field.
func (o *WidgetNumberFormat) SetUnit(v NumberFormatUnit) {
	o.Unit = &v
}

// GetUnitScale returns the UnitScale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetNumberFormat) GetUnitScale() NumberFormatUnitScale {
	if o == nil || o.UnitScale.Get() == nil {
		var ret NumberFormatUnitScale
		return ret
	}
	return *o.UnitScale.Get()
}

// GetUnitScaleOk returns a tuple with the UnitScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *WidgetNumberFormat) GetUnitScaleOk() (*NumberFormatUnitScale, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitScale.Get(), o.UnitScale.IsSet()
}

// HasUnitScale returns a boolean if a field has been set.
func (o *WidgetNumberFormat) HasUnitScale() bool {
	return o != nil && o.UnitScale.IsSet()
}

// SetUnitScale gets a reference to the given NullableNumberFormatUnitScale and assigns it to the UnitScale field.
func (o *WidgetNumberFormat) SetUnitScale(v NumberFormatUnitScale) {
	o.UnitScale.Set(&v)
}

// SetUnitScaleNil sets the value for UnitScale to be an explicit nil.
func (o *WidgetNumberFormat) SetUnitScaleNil() {
	o.UnitScale.Set(nil)
}

// UnsetUnitScale ensures that no value is present for UnitScale, not even an explicit nil.
func (o *WidgetNumberFormat) UnsetUnitScale() {
	o.UnitScale.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetNumberFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.UnitScale.IsSet() {
		toSerialize["unit_scale"] = o.UnitScale.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetNumberFormat) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Unit      *NumberFormatUnit             `json:"unit,omitempty"`
		UnitScale NullableNumberFormatUnitScale `json:"unit_scale,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"unit", "unit_scale"})
	} else {
		return err
	}
	o.Unit = all.Unit
	o.UnitScale = all.UnitScale

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
