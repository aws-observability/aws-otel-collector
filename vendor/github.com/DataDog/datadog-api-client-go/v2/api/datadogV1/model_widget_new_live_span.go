// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetNewLiveSpan Used for arbitrary live span times, such as 17 minutes or 6 hours.
type WidgetNewLiveSpan struct {
	// Whether to hide incomplete cost data in the widget.
	HideIncompleteCostData *bool `json:"hide_incomplete_cost_data,omitempty"`
	// Type "live" denotes a live span in the new format.
	Type WidgetNewLiveSpanType `json:"type"`
	// Unit of the time span.
	Unit WidgetLiveSpanUnit `json:"unit"`
	// Value of the time span.
	Value int64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetNewLiveSpan instantiates a new WidgetNewLiveSpan object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetNewLiveSpan(typeVar WidgetNewLiveSpanType, unit WidgetLiveSpanUnit, value int64) *WidgetNewLiveSpan {
	this := WidgetNewLiveSpan{}
	this.Type = typeVar
	this.Unit = unit
	this.Value = value
	return &this
}

// NewWidgetNewLiveSpanWithDefaults instantiates a new WidgetNewLiveSpan object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetNewLiveSpanWithDefaults() *WidgetNewLiveSpan {
	this := WidgetNewLiveSpan{}
	return &this
}

// GetHideIncompleteCostData returns the HideIncompleteCostData field value if set, zero value otherwise.
func (o *WidgetNewLiveSpan) GetHideIncompleteCostData() bool {
	if o == nil || o.HideIncompleteCostData == nil {
		var ret bool
		return ret
	}
	return *o.HideIncompleteCostData
}

// GetHideIncompleteCostDataOk returns a tuple with the HideIncompleteCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetNewLiveSpan) GetHideIncompleteCostDataOk() (*bool, bool) {
	if o == nil || o.HideIncompleteCostData == nil {
		return nil, false
	}
	return o.HideIncompleteCostData, true
}

// HasHideIncompleteCostData returns a boolean if a field has been set.
func (o *WidgetNewLiveSpan) HasHideIncompleteCostData() bool {
	return o != nil && o.HideIncompleteCostData != nil
}

// SetHideIncompleteCostData gets a reference to the given bool and assigns it to the HideIncompleteCostData field.
func (o *WidgetNewLiveSpan) SetHideIncompleteCostData(v bool) {
	o.HideIncompleteCostData = &v
}

// GetType returns the Type field value.
func (o *WidgetNewLiveSpan) GetType() WidgetNewLiveSpanType {
	if o == nil {
		var ret WidgetNewLiveSpanType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WidgetNewLiveSpan) GetTypeOk() (*WidgetNewLiveSpanType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WidgetNewLiveSpan) SetType(v WidgetNewLiveSpanType) {
	o.Type = v
}

// GetUnit returns the Unit field value.
func (o *WidgetNewLiveSpan) GetUnit() WidgetLiveSpanUnit {
	if o == nil {
		var ret WidgetLiveSpanUnit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *WidgetNewLiveSpan) GetUnitOk() (*WidgetLiveSpanUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *WidgetNewLiveSpan) SetUnit(v WidgetLiveSpanUnit) {
	o.Unit = v
}

// GetValue returns the Value field value.
func (o *WidgetNewLiveSpan) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *WidgetNewLiveSpan) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *WidgetNewLiveSpan) SetValue(v int64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetNewLiveSpan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HideIncompleteCostData != nil {
		toSerialize["hide_incomplete_cost_data"] = o.HideIncompleteCostData
	}
	toSerialize["type"] = o.Type
	toSerialize["unit"] = o.Unit
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetNewLiveSpan) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HideIncompleteCostData *bool                  `json:"hide_incomplete_cost_data,omitempty"`
		Type                   *WidgetNewLiveSpanType `json:"type"`
		Unit                   *WidgetLiveSpanUnit    `json:"unit"`
		Value                  *int64                 `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Unit == nil {
		return fmt.Errorf("required field unit missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hide_incomplete_cost_data", "type", "unit", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.HideIncompleteCostData = all.HideIncompleteCostData
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if !all.Unit.IsValid() {
		hasInvalidField = true
	} else {
		o.Unit = *all.Unit
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
