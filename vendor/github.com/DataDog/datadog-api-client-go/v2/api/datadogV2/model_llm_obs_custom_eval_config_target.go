// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigTarget Target application configuration for a custom evaluator.
type LLMObsCustomEvalConfigTarget struct {
	// Name of the ML application this evaluator targets.
	ApplicationName string `json:"application_name"`
	// Whether the evaluator is active for the target application.
	Enabled bool `json:"enabled"`
	// Scope at which to evaluate spans.
	EvalScope *LLMObsCustomEvalConfigEvalScope `json:"eval_scope,omitempty"`
	// Filter expression to select which spans to evaluate.
	Filter datadog.NullableString `json:"filter,omitempty"`
	// When true, only root spans are evaluated.
	RootSpansOnly datadog.NullableBool `json:"root_spans_only,omitempty"`
	// Percentage of traces to evaluate. Must be greater than 0 and at most 100.
	SamplingPercentage datadog.NullableFloat64 `json:"sampling_percentage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCustomEvalConfigTarget instantiates a new LLMObsCustomEvalConfigTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCustomEvalConfigTarget(applicationName string, enabled bool) *LLMObsCustomEvalConfigTarget {
	this := LLMObsCustomEvalConfigTarget{}
	this.ApplicationName = applicationName
	this.Enabled = enabled
	return &this
}

// NewLLMObsCustomEvalConfigTargetWithDefaults instantiates a new LLMObsCustomEvalConfigTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCustomEvalConfigTargetWithDefaults() *LLMObsCustomEvalConfigTarget {
	this := LLMObsCustomEvalConfigTarget{}
	return &this
}

// GetApplicationName returns the ApplicationName field value.
func (o *LLMObsCustomEvalConfigTarget) GetApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigTarget) GetApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationName, true
}

// SetApplicationName sets field value.
func (o *LLMObsCustomEvalConfigTarget) SetApplicationName(v string) {
	o.ApplicationName = v
}

// GetEnabled returns the Enabled field value.
func (o *LLMObsCustomEvalConfigTarget) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigTarget) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *LLMObsCustomEvalConfigTarget) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEvalScope returns the EvalScope field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigTarget) GetEvalScope() LLMObsCustomEvalConfigEvalScope {
	if o == nil || o.EvalScope == nil {
		var ret LLMObsCustomEvalConfigEvalScope
		return ret
	}
	return *o.EvalScope
}

// GetEvalScopeOk returns a tuple with the EvalScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigTarget) GetEvalScopeOk() (*LLMObsCustomEvalConfigEvalScope, bool) {
	if o == nil || o.EvalScope == nil {
		return nil, false
	}
	return o.EvalScope, true
}

// HasEvalScope returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigTarget) HasEvalScope() bool {
	return o != nil && o.EvalScope != nil
}

// SetEvalScope gets a reference to the given LLMObsCustomEvalConfigEvalScope and assigns it to the EvalScope field.
func (o *LLMObsCustomEvalConfigTarget) SetEvalScope(v LLMObsCustomEvalConfigEvalScope) {
	o.EvalScope = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigTarget) GetFilter() string {
	if o == nil || o.Filter.Get() == nil {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigTarget) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigTarget) HasFilter() bool {
	return o != nil && o.Filter.IsSet()
}

// SetFilter gets a reference to the given datadog.NullableString and assigns it to the Filter field.
func (o *LLMObsCustomEvalConfigTarget) SetFilter(v string) {
	o.Filter.Set(&v)
}

// SetFilterNil sets the value for Filter to be an explicit nil.
func (o *LLMObsCustomEvalConfigTarget) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil.
func (o *LLMObsCustomEvalConfigTarget) UnsetFilter() {
	o.Filter.Unset()
}

// GetRootSpansOnly returns the RootSpansOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigTarget) GetRootSpansOnly() bool {
	if o == nil || o.RootSpansOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RootSpansOnly.Get()
}

// GetRootSpansOnlyOk returns a tuple with the RootSpansOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigTarget) GetRootSpansOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootSpansOnly.Get(), o.RootSpansOnly.IsSet()
}

// HasRootSpansOnly returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigTarget) HasRootSpansOnly() bool {
	return o != nil && o.RootSpansOnly.IsSet()
}

// SetRootSpansOnly gets a reference to the given datadog.NullableBool and assigns it to the RootSpansOnly field.
func (o *LLMObsCustomEvalConfigTarget) SetRootSpansOnly(v bool) {
	o.RootSpansOnly.Set(&v)
}

// SetRootSpansOnlyNil sets the value for RootSpansOnly to be an explicit nil.
func (o *LLMObsCustomEvalConfigTarget) SetRootSpansOnlyNil() {
	o.RootSpansOnly.Set(nil)
}

// UnsetRootSpansOnly ensures that no value is present for RootSpansOnly, not even an explicit nil.
func (o *LLMObsCustomEvalConfigTarget) UnsetRootSpansOnly() {
	o.RootSpansOnly.Unset()
}

// GetSamplingPercentage returns the SamplingPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigTarget) GetSamplingPercentage() float64 {
	if o == nil || o.SamplingPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SamplingPercentage.Get()
}

// GetSamplingPercentageOk returns a tuple with the SamplingPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigTarget) GetSamplingPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SamplingPercentage.Get(), o.SamplingPercentage.IsSet()
}

// HasSamplingPercentage returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigTarget) HasSamplingPercentage() bool {
	return o != nil && o.SamplingPercentage.IsSet()
}

// SetSamplingPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the SamplingPercentage field.
func (o *LLMObsCustomEvalConfigTarget) SetSamplingPercentage(v float64) {
	o.SamplingPercentage.Set(&v)
}

// SetSamplingPercentageNil sets the value for SamplingPercentage to be an explicit nil.
func (o *LLMObsCustomEvalConfigTarget) SetSamplingPercentageNil() {
	o.SamplingPercentage.Set(nil)
}

// UnsetSamplingPercentage ensures that no value is present for SamplingPercentage, not even an explicit nil.
func (o *LLMObsCustomEvalConfigTarget) UnsetSamplingPercentage() {
	o.SamplingPercentage.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCustomEvalConfigTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_name"] = o.ApplicationName
	toSerialize["enabled"] = o.Enabled
	if o.EvalScope != nil {
		toSerialize["eval_scope"] = o.EvalScope
	}
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	if o.RootSpansOnly.IsSet() {
		toSerialize["root_spans_only"] = o.RootSpansOnly.Get()
	}
	if o.SamplingPercentage.IsSet() {
		toSerialize["sampling_percentage"] = o.SamplingPercentage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCustomEvalConfigTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationName    *string                          `json:"application_name"`
		Enabled            *bool                            `json:"enabled"`
		EvalScope          *LLMObsCustomEvalConfigEvalScope `json:"eval_scope,omitempty"`
		Filter             datadog.NullableString           `json:"filter,omitempty"`
		RootSpansOnly      datadog.NullableBool             `json:"root_spans_only,omitempty"`
		SamplingPercentage datadog.NullableFloat64          `json:"sampling_percentage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationName == nil {
		return fmt.Errorf("required field application_name missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_name", "enabled", "eval_scope", "filter", "root_spans_only", "sampling_percentage"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationName = *all.ApplicationName
	o.Enabled = *all.Enabled
	if all.EvalScope != nil && !all.EvalScope.IsValid() {
		hasInvalidField = true
	} else {
		o.EvalScope = all.EvalScope
	}
	o.Filter = all.Filter
	o.RootSpansOnly = all.RootSpansOnly
	o.SamplingPercentage = all.SamplingPercentage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
