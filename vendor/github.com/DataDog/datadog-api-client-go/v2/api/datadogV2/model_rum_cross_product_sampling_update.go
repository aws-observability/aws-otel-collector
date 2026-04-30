// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumCrossProductSamplingUpdate The configuration for cross-product retention filters. All fields are optional for partial updates.
type RumCrossProductSamplingUpdate struct {
	// Whether the cross-product retention filter for APM traces is enabled.
	TraceEnabled *bool `json:"trace_enabled,omitempty"`
	// The sample rate for the APM cross-product retention filter, between 0 and 100.
	TraceSampleRate *float64 `json:"trace_sample_rate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumCrossProductSamplingUpdate instantiates a new RumCrossProductSamplingUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumCrossProductSamplingUpdate() *RumCrossProductSamplingUpdate {
	this := RumCrossProductSamplingUpdate{}
	return &this
}

// NewRumCrossProductSamplingUpdateWithDefaults instantiates a new RumCrossProductSamplingUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumCrossProductSamplingUpdateWithDefaults() *RumCrossProductSamplingUpdate {
	this := RumCrossProductSamplingUpdate{}
	return &this
}

// GetTraceEnabled returns the TraceEnabled field value if set, zero value otherwise.
func (o *RumCrossProductSamplingUpdate) GetTraceEnabled() bool {
	if o == nil || o.TraceEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TraceEnabled
}

// GetTraceEnabledOk returns a tuple with the TraceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumCrossProductSamplingUpdate) GetTraceEnabledOk() (*bool, bool) {
	if o == nil || o.TraceEnabled == nil {
		return nil, false
	}
	return o.TraceEnabled, true
}

// HasTraceEnabled returns a boolean if a field has been set.
func (o *RumCrossProductSamplingUpdate) HasTraceEnabled() bool {
	return o != nil && o.TraceEnabled != nil
}

// SetTraceEnabled gets a reference to the given bool and assigns it to the TraceEnabled field.
func (o *RumCrossProductSamplingUpdate) SetTraceEnabled(v bool) {
	o.TraceEnabled = &v
}

// GetTraceSampleRate returns the TraceSampleRate field value if set, zero value otherwise.
func (o *RumCrossProductSamplingUpdate) GetTraceSampleRate() float64 {
	if o == nil || o.TraceSampleRate == nil {
		var ret float64
		return ret
	}
	return *o.TraceSampleRate
}

// GetTraceSampleRateOk returns a tuple with the TraceSampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumCrossProductSamplingUpdate) GetTraceSampleRateOk() (*float64, bool) {
	if o == nil || o.TraceSampleRate == nil {
		return nil, false
	}
	return o.TraceSampleRate, true
}

// HasTraceSampleRate returns a boolean if a field has been set.
func (o *RumCrossProductSamplingUpdate) HasTraceSampleRate() bool {
	return o != nil && o.TraceSampleRate != nil
}

// SetTraceSampleRate gets a reference to the given float64 and assigns it to the TraceSampleRate field.
func (o *RumCrossProductSamplingUpdate) SetTraceSampleRate(v float64) {
	o.TraceSampleRate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumCrossProductSamplingUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TraceEnabled != nil {
		toSerialize["trace_enabled"] = o.TraceEnabled
	}
	if o.TraceSampleRate != nil {
		toSerialize["trace_sample_rate"] = o.TraceSampleRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumCrossProductSamplingUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TraceEnabled    *bool    `json:"trace_enabled,omitempty"`
		TraceSampleRate *float64 `json:"trace_sample_rate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"trace_enabled", "trace_sample_rate"})
	} else {
		return err
	}
	o.TraceEnabled = all.TraceEnabled
	o.TraceSampleRate = all.TraceSampleRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
