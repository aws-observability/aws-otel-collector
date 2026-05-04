// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetNewFixedSpan Used for fixed span times, such as 'March 1 to March 7'.
type WidgetNewFixedSpan struct {
	// Start time in milliseconds since epoch.
	From int64 `json:"from"`
	// Whether to hide incomplete cost data in the widget.
	HideIncompleteCostData *bool `json:"hide_incomplete_cost_data,omitempty"`
	// End time in milliseconds since epoch.
	To int64 `json:"to"`
	// Type "fixed" denotes a fixed span.
	Type WidgetNewFixedSpanType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetNewFixedSpan instantiates a new WidgetNewFixedSpan object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetNewFixedSpan(from int64, to int64, typeVar WidgetNewFixedSpanType) *WidgetNewFixedSpan {
	this := WidgetNewFixedSpan{}
	this.From = from
	this.To = to
	this.Type = typeVar
	return &this
}

// NewWidgetNewFixedSpanWithDefaults instantiates a new WidgetNewFixedSpan object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetNewFixedSpanWithDefaults() *WidgetNewFixedSpan {
	this := WidgetNewFixedSpan{}
	return &this
}

// GetFrom returns the From field value.
func (o *WidgetNewFixedSpan) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *WidgetNewFixedSpan) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *WidgetNewFixedSpan) SetFrom(v int64) {
	o.From = v
}

// GetHideIncompleteCostData returns the HideIncompleteCostData field value if set, zero value otherwise.
func (o *WidgetNewFixedSpan) GetHideIncompleteCostData() bool {
	if o == nil || o.HideIncompleteCostData == nil {
		var ret bool
		return ret
	}
	return *o.HideIncompleteCostData
}

// GetHideIncompleteCostDataOk returns a tuple with the HideIncompleteCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetNewFixedSpan) GetHideIncompleteCostDataOk() (*bool, bool) {
	if o == nil || o.HideIncompleteCostData == nil {
		return nil, false
	}
	return o.HideIncompleteCostData, true
}

// HasHideIncompleteCostData returns a boolean if a field has been set.
func (o *WidgetNewFixedSpan) HasHideIncompleteCostData() bool {
	return o != nil && o.HideIncompleteCostData != nil
}

// SetHideIncompleteCostData gets a reference to the given bool and assigns it to the HideIncompleteCostData field.
func (o *WidgetNewFixedSpan) SetHideIncompleteCostData(v bool) {
	o.HideIncompleteCostData = &v
}

// GetTo returns the To field value.
func (o *WidgetNewFixedSpan) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *WidgetNewFixedSpan) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *WidgetNewFixedSpan) SetTo(v int64) {
	o.To = v
}

// GetType returns the Type field value.
func (o *WidgetNewFixedSpan) GetType() WidgetNewFixedSpanType {
	if o == nil {
		var ret WidgetNewFixedSpanType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WidgetNewFixedSpan) GetTypeOk() (*WidgetNewFixedSpanType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WidgetNewFixedSpan) SetType(v WidgetNewFixedSpanType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetNewFixedSpan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from"] = o.From
	if o.HideIncompleteCostData != nil {
		toSerialize["hide_incomplete_cost_data"] = o.HideIncompleteCostData
	}
	toSerialize["to"] = o.To
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetNewFixedSpan) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From                   *int64                  `json:"from"`
		HideIncompleteCostData *bool                   `json:"hide_incomplete_cost_data,omitempty"`
		To                     *int64                  `json:"to"`
		Type                   *WidgetNewFixedSpanType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "hide_incomplete_cost_data", "to", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.From = *all.From
	o.HideIncompleteCostData = all.HideIncompleteCostData
	o.To = *all.To
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
