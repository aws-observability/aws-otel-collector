// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloStatusDataAttributes The attributes of the SLO status.
type SloStatusDataAttributes struct {
	// The percentage of error budget remaining.
	ErrorBudgetRemaining float64 `json:"error_budget_remaining"`
	// The raw error budget remaining for the SLO.
	RawErrorBudgetRemaining RawErrorBudgetRemaining `json:"raw_error_budget_remaining"`
	// The current Service Level Indicator (SLI) value as a percentage.
	Sli float64 `json:"sli"`
	// The precision of the time span in seconds.
	SpanPrecision int64 `json:"span_precision"`
	// The current state of the SLO (for example, `breached`, `warning`, `ok`).
	State string `json:"state"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSloStatusDataAttributes instantiates a new SloStatusDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSloStatusDataAttributes(errorBudgetRemaining float64, rawErrorBudgetRemaining RawErrorBudgetRemaining, sli float64, spanPrecision int64, state string) *SloStatusDataAttributes {
	this := SloStatusDataAttributes{}
	this.ErrorBudgetRemaining = errorBudgetRemaining
	this.RawErrorBudgetRemaining = rawErrorBudgetRemaining
	this.Sli = sli
	this.SpanPrecision = spanPrecision
	this.State = state
	return &this
}

// NewSloStatusDataAttributesWithDefaults instantiates a new SloStatusDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSloStatusDataAttributesWithDefaults() *SloStatusDataAttributes {
	this := SloStatusDataAttributes{}
	return &this
}

// GetErrorBudgetRemaining returns the ErrorBudgetRemaining field value.
func (o *SloStatusDataAttributes) GetErrorBudgetRemaining() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ErrorBudgetRemaining
}

// GetErrorBudgetRemainingOk returns a tuple with the ErrorBudgetRemaining field value
// and a boolean to check if the value has been set.
func (o *SloStatusDataAttributes) GetErrorBudgetRemainingOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorBudgetRemaining, true
}

// SetErrorBudgetRemaining sets field value.
func (o *SloStatusDataAttributes) SetErrorBudgetRemaining(v float64) {
	o.ErrorBudgetRemaining = v
}

// GetRawErrorBudgetRemaining returns the RawErrorBudgetRemaining field value.
func (o *SloStatusDataAttributes) GetRawErrorBudgetRemaining() RawErrorBudgetRemaining {
	if o == nil {
		var ret RawErrorBudgetRemaining
		return ret
	}
	return o.RawErrorBudgetRemaining
}

// GetRawErrorBudgetRemainingOk returns a tuple with the RawErrorBudgetRemaining field value
// and a boolean to check if the value has been set.
func (o *SloStatusDataAttributes) GetRawErrorBudgetRemainingOk() (*RawErrorBudgetRemaining, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawErrorBudgetRemaining, true
}

// SetRawErrorBudgetRemaining sets field value.
func (o *SloStatusDataAttributes) SetRawErrorBudgetRemaining(v RawErrorBudgetRemaining) {
	o.RawErrorBudgetRemaining = v
}

// GetSli returns the Sli field value.
func (o *SloStatusDataAttributes) GetSli() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Sli
}

// GetSliOk returns a tuple with the Sli field value
// and a boolean to check if the value has been set.
func (o *SloStatusDataAttributes) GetSliOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sli, true
}

// SetSli sets field value.
func (o *SloStatusDataAttributes) SetSli(v float64) {
	o.Sli = v
}

// GetSpanPrecision returns the SpanPrecision field value.
func (o *SloStatusDataAttributes) GetSpanPrecision() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SpanPrecision
}

// GetSpanPrecisionOk returns a tuple with the SpanPrecision field value
// and a boolean to check if the value has been set.
func (o *SloStatusDataAttributes) GetSpanPrecisionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanPrecision, true
}

// SetSpanPrecision sets field value.
func (o *SloStatusDataAttributes) SetSpanPrecision(v int64) {
	o.SpanPrecision = v
}

// GetState returns the State field value.
func (o *SloStatusDataAttributes) GetState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SloStatusDataAttributes) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *SloStatusDataAttributes) SetState(v string) {
	o.State = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SloStatusDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["error_budget_remaining"] = o.ErrorBudgetRemaining
	toSerialize["raw_error_budget_remaining"] = o.RawErrorBudgetRemaining
	toSerialize["sli"] = o.Sli
	toSerialize["span_precision"] = o.SpanPrecision
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SloStatusDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ErrorBudgetRemaining    *float64                 `json:"error_budget_remaining"`
		RawErrorBudgetRemaining *RawErrorBudgetRemaining `json:"raw_error_budget_remaining"`
		Sli                     *float64                 `json:"sli"`
		SpanPrecision           *int64                   `json:"span_precision"`
		State                   *string                  `json:"state"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ErrorBudgetRemaining == nil {
		return fmt.Errorf("required field error_budget_remaining missing")
	}
	if all.RawErrorBudgetRemaining == nil {
		return fmt.Errorf("required field raw_error_budget_remaining missing")
	}
	if all.Sli == nil {
		return fmt.Errorf("required field sli missing")
	}
	if all.SpanPrecision == nil {
		return fmt.Errorf("required field span_precision missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error_budget_remaining", "raw_error_budget_remaining", "sli", "span_precision", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ErrorBudgetRemaining = *all.ErrorBudgetRemaining
	if all.RawErrorBudgetRemaining.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RawErrorBudgetRemaining = *all.RawErrorBudgetRemaining
	o.Sli = *all.Sli
	o.SpanPrecision = *all.SpanPrecision
	o.State = *all.State

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
