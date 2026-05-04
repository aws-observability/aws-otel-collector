// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigAssessmentCriteria Criteria used to assess the pass/fail result of a custom evaluator.
type LLMObsCustomEvalConfigAssessmentCriteria struct {
	// Maximum numeric threshold for a passing result.
	MaxThreshold datadog.NullableFloat64 `json:"max_threshold,omitempty"`
	// Minimum numeric threshold for a passing result.
	MinThreshold datadog.NullableFloat64 `json:"min_threshold,omitempty"`
	// Specific output values considered as a passing result.
	PassValues datadog.NullableList[string] `json:"pass_values,omitempty"`
	// When true, a boolean output of true is treated as passing.
	PassWhen datadog.NullableBool `json:"pass_when,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCustomEvalConfigAssessmentCriteria instantiates a new LLMObsCustomEvalConfigAssessmentCriteria object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCustomEvalConfigAssessmentCriteria() *LLMObsCustomEvalConfigAssessmentCriteria {
	this := LLMObsCustomEvalConfigAssessmentCriteria{}
	return &this
}

// NewLLMObsCustomEvalConfigAssessmentCriteriaWithDefaults instantiates a new LLMObsCustomEvalConfigAssessmentCriteria object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCustomEvalConfigAssessmentCriteriaWithDefaults() *LLMObsCustomEvalConfigAssessmentCriteria {
	this := LLMObsCustomEvalConfigAssessmentCriteria{}
	return &this
}

// GetMaxThreshold returns the MaxThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigAssessmentCriteria) GetMaxThreshold() float64 {
	if o == nil || o.MaxThreshold.Get() == nil {
		var ret float64
		return ret
	}
	return *o.MaxThreshold.Get()
}

// GetMaxThresholdOk returns a tuple with the MaxThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) GetMaxThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxThreshold.Get(), o.MaxThreshold.IsSet()
}

// HasMaxThreshold returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) HasMaxThreshold() bool {
	return o != nil && o.MaxThreshold.IsSet()
}

// SetMaxThreshold gets a reference to the given datadog.NullableFloat64 and assigns it to the MaxThreshold field.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) SetMaxThreshold(v float64) {
	o.MaxThreshold.Set(&v)
}

// SetMaxThresholdNil sets the value for MaxThreshold to be an explicit nil.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) SetMaxThresholdNil() {
	o.MaxThreshold.Set(nil)
}

// UnsetMaxThreshold ensures that no value is present for MaxThreshold, not even an explicit nil.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) UnsetMaxThreshold() {
	o.MaxThreshold.Unset()
}

// GetMinThreshold returns the MinThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigAssessmentCriteria) GetMinThreshold() float64 {
	if o == nil || o.MinThreshold.Get() == nil {
		var ret float64
		return ret
	}
	return *o.MinThreshold.Get()
}

// GetMinThresholdOk returns a tuple with the MinThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) GetMinThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinThreshold.Get(), o.MinThreshold.IsSet()
}

// HasMinThreshold returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) HasMinThreshold() bool {
	return o != nil && o.MinThreshold.IsSet()
}

// SetMinThreshold gets a reference to the given datadog.NullableFloat64 and assigns it to the MinThreshold field.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) SetMinThreshold(v float64) {
	o.MinThreshold.Set(&v)
}

// SetMinThresholdNil sets the value for MinThreshold to be an explicit nil.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) SetMinThresholdNil() {
	o.MinThreshold.Set(nil)
}

// UnsetMinThreshold ensures that no value is present for MinThreshold, not even an explicit nil.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) UnsetMinThreshold() {
	o.MinThreshold.Unset()
}

// GetPassValues returns the PassValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigAssessmentCriteria) GetPassValues() []string {
	if o == nil || o.PassValues.Get() == nil {
		var ret []string
		return ret
	}
	return *o.PassValues.Get()
}

// GetPassValuesOk returns a tuple with the PassValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) GetPassValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PassValues.Get(), o.PassValues.IsSet()
}

// HasPassValues returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) HasPassValues() bool {
	return o != nil && o.PassValues.IsSet()
}

// SetPassValues gets a reference to the given datadog.NullableList[string] and assigns it to the PassValues field.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) SetPassValues(v []string) {
	o.PassValues.Set(&v)
}

// SetPassValuesNil sets the value for PassValues to be an explicit nil.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) SetPassValuesNil() {
	o.PassValues.Set(nil)
}

// UnsetPassValues ensures that no value is present for PassValues, not even an explicit nil.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) UnsetPassValues() {
	o.PassValues.Unset()
}

// GetPassWhen returns the PassWhen field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigAssessmentCriteria) GetPassWhen() bool {
	if o == nil || o.PassWhen.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PassWhen.Get()
}

// GetPassWhenOk returns a tuple with the PassWhen field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) GetPassWhenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PassWhen.Get(), o.PassWhen.IsSet()
}

// HasPassWhen returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) HasPassWhen() bool {
	return o != nil && o.PassWhen.IsSet()
}

// SetPassWhen gets a reference to the given datadog.NullableBool and assigns it to the PassWhen field.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) SetPassWhen(v bool) {
	o.PassWhen.Set(&v)
}

// SetPassWhenNil sets the value for PassWhen to be an explicit nil.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) SetPassWhenNil() {
	o.PassWhen.Set(nil)
}

// UnsetPassWhen ensures that no value is present for PassWhen, not even an explicit nil.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) UnsetPassWhen() {
	o.PassWhen.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCustomEvalConfigAssessmentCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MaxThreshold.IsSet() {
		toSerialize["max_threshold"] = o.MaxThreshold.Get()
	}
	if o.MinThreshold.IsSet() {
		toSerialize["min_threshold"] = o.MinThreshold.Get()
	}
	if o.PassValues.IsSet() {
		toSerialize["pass_values"] = o.PassValues.Get()
	}
	if o.PassWhen.IsSet() {
		toSerialize["pass_when"] = o.PassWhen.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCustomEvalConfigAssessmentCriteria) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MaxThreshold datadog.NullableFloat64      `json:"max_threshold,omitempty"`
		MinThreshold datadog.NullableFloat64      `json:"min_threshold,omitempty"`
		PassValues   datadog.NullableList[string] `json:"pass_values,omitempty"`
		PassWhen     datadog.NullableBool         `json:"pass_when,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max_threshold", "min_threshold", "pass_values", "pass_when"})
	} else {
		return err
	}
	o.MaxThreshold = all.MaxThreshold
	o.MinThreshold = all.MinThreshold
	o.PassValues = all.PassValues
	o.PassWhen = all.PassWhen

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
