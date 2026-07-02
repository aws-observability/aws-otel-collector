// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsInternalReasoning The model's internal reasoning or thinking output, if available.
type LLMObsInternalReasoning struct {
	// Number of tokens used for internal reasoning.
	ReasoningTokens datadog.NullableInt64 `json:"reasoning_tokens,omitempty"`
	// The reasoning text produced by the model.
	Text string `json:"text"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsInternalReasoning instantiates a new LLMObsInternalReasoning object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsInternalReasoning(text string) *LLMObsInternalReasoning {
	this := LLMObsInternalReasoning{}
	this.Text = text
	return &this
}

// NewLLMObsInternalReasoningWithDefaults instantiates a new LLMObsInternalReasoning object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsInternalReasoningWithDefaults() *LLMObsInternalReasoning {
	this := LLMObsInternalReasoning{}
	return &this
}

// GetReasoningTokens returns the ReasoningTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsInternalReasoning) GetReasoningTokens() int64 {
	if o == nil || o.ReasoningTokens.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ReasoningTokens.Get()
}

// GetReasoningTokensOk returns a tuple with the ReasoningTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsInternalReasoning) GetReasoningTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReasoningTokens.Get(), o.ReasoningTokens.IsSet()
}

// HasReasoningTokens returns a boolean if a field has been set.
func (o *LLMObsInternalReasoning) HasReasoningTokens() bool {
	return o != nil && o.ReasoningTokens.IsSet()
}

// SetReasoningTokens gets a reference to the given datadog.NullableInt64 and assigns it to the ReasoningTokens field.
func (o *LLMObsInternalReasoning) SetReasoningTokens(v int64) {
	o.ReasoningTokens.Set(&v)
}

// SetReasoningTokensNil sets the value for ReasoningTokens to be an explicit nil.
func (o *LLMObsInternalReasoning) SetReasoningTokensNil() {
	o.ReasoningTokens.Set(nil)
}

// UnsetReasoningTokens ensures that no value is present for ReasoningTokens, not even an explicit nil.
func (o *LLMObsInternalReasoning) UnsetReasoningTokens() {
	o.ReasoningTokens.Unset()
}

// GetText returns the Text field value.
func (o *LLMObsInternalReasoning) GetText() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *LLMObsInternalReasoning) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value.
func (o *LLMObsInternalReasoning) SetText(v string) {
	o.Text = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsInternalReasoning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ReasoningTokens.IsSet() {
		toSerialize["reasoning_tokens"] = o.ReasoningTokens.Get()
	}
	toSerialize["text"] = o.Text

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsInternalReasoning) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReasoningTokens datadog.NullableInt64 `json:"reasoning_tokens,omitempty"`
		Text            *string               `json:"text"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Text == nil {
		return fmt.Errorf("required field text missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"reasoning_tokens", "text"})
	} else {
		return err
	}
	o.ReasoningTokens = all.ReasoningTokens
	o.Text = *all.Text

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
