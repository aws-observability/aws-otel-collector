// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnthropicThinkingConfig Configuration for Anthropic extended thinking feature.
type LLMObsAnthropicThinkingConfig struct {
	// Maximum token budget for extended thinking. Required when type is `enabled`.
	BudgetTokens datadog.NullableInt64 `json:"budget_tokens,omitempty"`
	// The thinking mode for Anthropic extended thinking.
	Type LLMObsAnthropicThinkingType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnthropicThinkingConfig instantiates a new LLMObsAnthropicThinkingConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnthropicThinkingConfig(typeVar LLMObsAnthropicThinkingType) *LLMObsAnthropicThinkingConfig {
	this := LLMObsAnthropicThinkingConfig{}
	this.Type = typeVar
	return &this
}

// NewLLMObsAnthropicThinkingConfigWithDefaults instantiates a new LLMObsAnthropicThinkingConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnthropicThinkingConfigWithDefaults() *LLMObsAnthropicThinkingConfig {
	this := LLMObsAnthropicThinkingConfig{}
	return &this
}

// GetBudgetTokens returns the BudgetTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsAnthropicThinkingConfig) GetBudgetTokens() int64 {
	if o == nil || o.BudgetTokens.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BudgetTokens.Get()
}

// GetBudgetTokensOk returns a tuple with the BudgetTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsAnthropicThinkingConfig) GetBudgetTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BudgetTokens.Get(), o.BudgetTokens.IsSet()
}

// HasBudgetTokens returns a boolean if a field has been set.
func (o *LLMObsAnthropicThinkingConfig) HasBudgetTokens() bool {
	return o != nil && o.BudgetTokens.IsSet()
}

// SetBudgetTokens gets a reference to the given datadog.NullableInt64 and assigns it to the BudgetTokens field.
func (o *LLMObsAnthropicThinkingConfig) SetBudgetTokens(v int64) {
	o.BudgetTokens.Set(&v)
}

// SetBudgetTokensNil sets the value for BudgetTokens to be an explicit nil.
func (o *LLMObsAnthropicThinkingConfig) SetBudgetTokensNil() {
	o.BudgetTokens.Set(nil)
}

// UnsetBudgetTokens ensures that no value is present for BudgetTokens, not even an explicit nil.
func (o *LLMObsAnthropicThinkingConfig) UnsetBudgetTokens() {
	o.BudgetTokens.Unset()
}

// GetType returns the Type field value.
func (o *LLMObsAnthropicThinkingConfig) GetType() LLMObsAnthropicThinkingType {
	if o == nil {
		var ret LLMObsAnthropicThinkingType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnthropicThinkingConfig) GetTypeOk() (*LLMObsAnthropicThinkingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsAnthropicThinkingConfig) SetType(v LLMObsAnthropicThinkingType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnthropicThinkingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BudgetTokens.IsSet() {
		toSerialize["budget_tokens"] = o.BudgetTokens.Get()
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnthropicThinkingConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BudgetTokens datadog.NullableInt64        `json:"budget_tokens,omitempty"`
		Type         *LLMObsAnthropicThinkingType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"budget_tokens", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BudgetTokens = all.BudgetTokens
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
