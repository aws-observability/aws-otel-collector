// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentSpanMeta Metadata associated with an experiment span.
type LLMObsExperimentSpanMeta struct {
	// Error details for an experiment span.
	Error *LLMObsExperimentSpanError `json:"error,omitempty"`
	// Expected output for the span, used for evaluation.
	ExpectedOutput map[string]interface{} `json:"expected_output,omitempty"`
	// Represents any valid JSON value.
	Input NullableAnyValue `json:"input,omitempty"`
	// Represents any valid JSON value.
	Output NullableAnyValue `json:"output,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentSpanMeta instantiates a new LLMObsExperimentSpanMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentSpanMeta() *LLMObsExperimentSpanMeta {
	this := LLMObsExperimentSpanMeta{}
	return &this
}

// NewLLMObsExperimentSpanMetaWithDefaults instantiates a new LLMObsExperimentSpanMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentSpanMetaWithDefaults() *LLMObsExperimentSpanMeta {
	this := LLMObsExperimentSpanMeta{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanMeta) GetError() LLMObsExperimentSpanError {
	if o == nil || o.Error == nil {
		var ret LLMObsExperimentSpanError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanMeta) GetErrorOk() (*LLMObsExperimentSpanError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanMeta) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given LLMObsExperimentSpanError and assigns it to the Error field.
func (o *LLMObsExperimentSpanMeta) SetError(v LLMObsExperimentSpanError) {
	o.Error = &v
}

// GetExpectedOutput returns the ExpectedOutput field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanMeta) GetExpectedOutput() map[string]interface{} {
	if o == nil || o.ExpectedOutput == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExpectedOutput
}

// GetExpectedOutputOk returns a tuple with the ExpectedOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanMeta) GetExpectedOutputOk() (*map[string]interface{}, bool) {
	if o == nil || o.ExpectedOutput == nil {
		return nil, false
	}
	return &o.ExpectedOutput, true
}

// HasExpectedOutput returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanMeta) HasExpectedOutput() bool {
	return o != nil && o.ExpectedOutput != nil
}

// SetExpectedOutput gets a reference to the given map[string]interface{} and assigns it to the ExpectedOutput field.
func (o *LLMObsExperimentSpanMeta) SetExpectedOutput(v map[string]interface{}) {
	o.ExpectedOutput = v
}

// GetInput returns the Input field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentSpanMeta) GetInput() AnyValue {
	if o == nil || o.Input.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.Input.Get()
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentSpanMeta) GetInputOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Input.Get(), o.Input.IsSet()
}

// HasInput returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanMeta) HasInput() bool {
	return o != nil && o.Input.IsSet()
}

// SetInput gets a reference to the given NullableAnyValue and assigns it to the Input field.
func (o *LLMObsExperimentSpanMeta) SetInput(v AnyValue) {
	o.Input.Set(&v)
}

// SetInputNil sets the value for Input to be an explicit nil.
func (o *LLMObsExperimentSpanMeta) SetInputNil() {
	o.Input.Set(nil)
}

// UnsetInput ensures that no value is present for Input, not even an explicit nil.
func (o *LLMObsExperimentSpanMeta) UnsetInput() {
	o.Input.Unset()
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentSpanMeta) GetOutput() AnyValue {
	if o == nil || o.Output.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.Output.Get()
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentSpanMeta) GetOutputOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Output.Get(), o.Output.IsSet()
}

// HasOutput returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanMeta) HasOutput() bool {
	return o != nil && o.Output.IsSet()
}

// SetOutput gets a reference to the given NullableAnyValue and assigns it to the Output field.
func (o *LLMObsExperimentSpanMeta) SetOutput(v AnyValue) {
	o.Output.Set(&v)
}

// SetOutputNil sets the value for Output to be an explicit nil.
func (o *LLMObsExperimentSpanMeta) SetOutputNil() {
	o.Output.Set(nil)
}

// UnsetOutput ensures that no value is present for Output, not even an explicit nil.
func (o *LLMObsExperimentSpanMeta) UnsetOutput() {
	o.Output.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentSpanMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ExpectedOutput != nil {
		toSerialize["expected_output"] = o.ExpectedOutput
	}
	if o.Input.IsSet() {
		toSerialize["input"] = o.Input.Get()
	}
	if o.Output.IsSet() {
		toSerialize["output"] = o.Output.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentSpanMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error          *LLMObsExperimentSpanError `json:"error,omitempty"`
		ExpectedOutput map[string]interface{}     `json:"expected_output,omitempty"`
		Input          NullableAnyValue           `json:"input,omitempty"`
		Output         NullableAnyValue           `json:"output,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error", "expected_output", "input", "output"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Error != nil && all.Error.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Error = all.Error
	o.ExpectedOutput = all.ExpectedOutput
	o.Input = all.Input
	o.Output = all.Output

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
