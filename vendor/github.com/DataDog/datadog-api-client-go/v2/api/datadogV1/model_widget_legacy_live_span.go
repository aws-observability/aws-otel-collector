// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetLegacyLiveSpan Wrapper for live span
type WidgetLegacyLiveSpan struct {
	// Whether to hide incomplete cost data in the widget.
	HideIncompleteCostData *bool `json:"hide_incomplete_cost_data,omitempty"`
	// The available timeframes depend on the widget you are using.
	LiveSpan *WidgetLiveSpan `json:"live_span,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWidgetLegacyLiveSpan instantiates a new WidgetLegacyLiveSpan object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetLegacyLiveSpan() *WidgetLegacyLiveSpan {
	this := WidgetLegacyLiveSpan{}
	return &this
}

// NewWidgetLegacyLiveSpanWithDefaults instantiates a new WidgetLegacyLiveSpan object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetLegacyLiveSpanWithDefaults() *WidgetLegacyLiveSpan {
	this := WidgetLegacyLiveSpan{}
	return &this
}

// GetHideIncompleteCostData returns the HideIncompleteCostData field value if set, zero value otherwise.
func (o *WidgetLegacyLiveSpan) GetHideIncompleteCostData() bool {
	if o == nil || o.HideIncompleteCostData == nil {
		var ret bool
		return ret
	}
	return *o.HideIncompleteCostData
}

// GetHideIncompleteCostDataOk returns a tuple with the HideIncompleteCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetLegacyLiveSpan) GetHideIncompleteCostDataOk() (*bool, bool) {
	if o == nil || o.HideIncompleteCostData == nil {
		return nil, false
	}
	return o.HideIncompleteCostData, true
}

// HasHideIncompleteCostData returns a boolean if a field has been set.
func (o *WidgetLegacyLiveSpan) HasHideIncompleteCostData() bool {
	return o != nil && o.HideIncompleteCostData != nil
}

// SetHideIncompleteCostData gets a reference to the given bool and assigns it to the HideIncompleteCostData field.
func (o *WidgetLegacyLiveSpan) SetHideIncompleteCostData(v bool) {
	o.HideIncompleteCostData = &v
}

// GetLiveSpan returns the LiveSpan field value if set, zero value otherwise.
func (o *WidgetLegacyLiveSpan) GetLiveSpan() WidgetLiveSpan {
	if o == nil || o.LiveSpan == nil {
		var ret WidgetLiveSpan
		return ret
	}
	return *o.LiveSpan
}

// GetLiveSpanOk returns a tuple with the LiveSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetLegacyLiveSpan) GetLiveSpanOk() (*WidgetLiveSpan, bool) {
	if o == nil || o.LiveSpan == nil {
		return nil, false
	}
	return o.LiveSpan, true
}

// HasLiveSpan returns a boolean if a field has been set.
func (o *WidgetLegacyLiveSpan) HasLiveSpan() bool {
	return o != nil && o.LiveSpan != nil
}

// SetLiveSpan gets a reference to the given WidgetLiveSpan and assigns it to the LiveSpan field.
func (o *WidgetLegacyLiveSpan) SetLiveSpan(v WidgetLiveSpan) {
	o.LiveSpan = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetLegacyLiveSpan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HideIncompleteCostData != nil {
		toSerialize["hide_incomplete_cost_data"] = o.HideIncompleteCostData
	}
	if o.LiveSpan != nil {
		toSerialize["live_span"] = o.LiveSpan
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetLegacyLiveSpan) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HideIncompleteCostData *bool           `json:"hide_incomplete_cost_data,omitempty"`
		LiveSpan               *WidgetLiveSpan `json:"live_span,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.HideIncompleteCostData = all.HideIncompleteCostData
	if all.LiveSpan != nil && !all.LiveSpan.IsValid() {
		hasInvalidField = true
	} else {
		o.LiveSpan = all.LiveSpan
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
