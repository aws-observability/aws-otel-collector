// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnthropicMetadata Anthropic-specific metadata for an inference request.
type LLMObsAnthropicMetadata struct {
	// The effort level for Anthropic inference.
	Effort NullableLLMObsAnthropicEffort `json:"effort,omitempty"`
	// Configuration for Anthropic extended thinking feature.
	Thinking *LLMObsAnthropicThinkingConfig `json:"thinking,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnthropicMetadata instantiates a new LLMObsAnthropicMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnthropicMetadata() *LLMObsAnthropicMetadata {
	this := LLMObsAnthropicMetadata{}
	return &this
}

// NewLLMObsAnthropicMetadataWithDefaults instantiates a new LLMObsAnthropicMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnthropicMetadataWithDefaults() *LLMObsAnthropicMetadata {
	this := LLMObsAnthropicMetadata{}
	return &this
}

// GetEffort returns the Effort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsAnthropicMetadata) GetEffort() LLMObsAnthropicEffort {
	if o == nil || o.Effort.Get() == nil {
		var ret LLMObsAnthropicEffort
		return ret
	}
	return *o.Effort.Get()
}

// GetEffortOk returns a tuple with the Effort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsAnthropicMetadata) GetEffortOk() (*LLMObsAnthropicEffort, bool) {
	if o == nil {
		return nil, false
	}
	return o.Effort.Get(), o.Effort.IsSet()
}

// HasEffort returns a boolean if a field has been set.
func (o *LLMObsAnthropicMetadata) HasEffort() bool {
	return o != nil && o.Effort.IsSet()
}

// SetEffort gets a reference to the given NullableLLMObsAnthropicEffort and assigns it to the Effort field.
func (o *LLMObsAnthropicMetadata) SetEffort(v LLMObsAnthropicEffort) {
	o.Effort.Set(&v)
}

// SetEffortNil sets the value for Effort to be an explicit nil.
func (o *LLMObsAnthropicMetadata) SetEffortNil() {
	o.Effort.Set(nil)
}

// UnsetEffort ensures that no value is present for Effort, not even an explicit nil.
func (o *LLMObsAnthropicMetadata) UnsetEffort() {
	o.Effort.Unset()
}

// GetThinking returns the Thinking field value if set, zero value otherwise.
func (o *LLMObsAnthropicMetadata) GetThinking() LLMObsAnthropicThinkingConfig {
	if o == nil || o.Thinking == nil {
		var ret LLMObsAnthropicThinkingConfig
		return ret
	}
	return *o.Thinking
}

// GetThinkingOk returns a tuple with the Thinking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnthropicMetadata) GetThinkingOk() (*LLMObsAnthropicThinkingConfig, bool) {
	if o == nil || o.Thinking == nil {
		return nil, false
	}
	return o.Thinking, true
}

// HasThinking returns a boolean if a field has been set.
func (o *LLMObsAnthropicMetadata) HasThinking() bool {
	return o != nil && o.Thinking != nil
}

// SetThinking gets a reference to the given LLMObsAnthropicThinkingConfig and assigns it to the Thinking field.
func (o *LLMObsAnthropicMetadata) SetThinking(v LLMObsAnthropicThinkingConfig) {
	o.Thinking = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnthropicMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Effort.IsSet() {
		toSerialize["effort"] = o.Effort.Get()
	}
	if o.Thinking != nil {
		toSerialize["thinking"] = o.Thinking
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnthropicMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Effort   NullableLLMObsAnthropicEffort  `json:"effort,omitempty"`
		Thinking *LLMObsAnthropicThinkingConfig `json:"thinking,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"effort", "thinking"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Effort.Get() != nil && !all.Effort.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.Effort = all.Effort
	}
	if all.Thinking != nil && all.Thinking.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Thinking = all.Thinking

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
