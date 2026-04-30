// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigUpdateAttributes Attributes for creating or updating a custom LLM Observability evaluator configuration.
type LLMObsCustomEvalConfigUpdateAttributes struct {
	// Category of the evaluator.
	Category *string `json:"category,omitempty"`
	// Name of the custom evaluator. If provided, must match the eval_name path parameter.
	EvalName *string `json:"eval_name,omitempty"`
	// LLM judge configuration for a custom evaluator.
	LlmJudgeConfig *LLMObsCustomEvalConfigLLMJudgeConfig `json:"llm_judge_config,omitempty"`
	// LLM provider configuration for a custom evaluator.
	LlmProvider *LLMObsCustomEvalConfigLLMProvider `json:"llm_provider,omitempty"`
	// Target application configuration for a custom evaluator.
	Target LLMObsCustomEvalConfigTarget `json:"target"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCustomEvalConfigUpdateAttributes instantiates a new LLMObsCustomEvalConfigUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCustomEvalConfigUpdateAttributes(target LLMObsCustomEvalConfigTarget) *LLMObsCustomEvalConfigUpdateAttributes {
	this := LLMObsCustomEvalConfigUpdateAttributes{}
	this.Target = target
	return &this
}

// NewLLMObsCustomEvalConfigUpdateAttributesWithDefaults instantiates a new LLMObsCustomEvalConfigUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCustomEvalConfigUpdateAttributesWithDefaults() *LLMObsCustomEvalConfigUpdateAttributes {
	this := LLMObsCustomEvalConfigUpdateAttributes{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigUpdateAttributes) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigUpdateAttributes) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigUpdateAttributes) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *LLMObsCustomEvalConfigUpdateAttributes) SetCategory(v string) {
	o.Category = &v
}

// GetEvalName returns the EvalName field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigUpdateAttributes) GetEvalName() string {
	if o == nil || o.EvalName == nil {
		var ret string
		return ret
	}
	return *o.EvalName
}

// GetEvalNameOk returns a tuple with the EvalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigUpdateAttributes) GetEvalNameOk() (*string, bool) {
	if o == nil || o.EvalName == nil {
		return nil, false
	}
	return o.EvalName, true
}

// HasEvalName returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigUpdateAttributes) HasEvalName() bool {
	return o != nil && o.EvalName != nil
}

// SetEvalName gets a reference to the given string and assigns it to the EvalName field.
func (o *LLMObsCustomEvalConfigUpdateAttributes) SetEvalName(v string) {
	o.EvalName = &v
}

// GetLlmJudgeConfig returns the LlmJudgeConfig field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigUpdateAttributes) GetLlmJudgeConfig() LLMObsCustomEvalConfigLLMJudgeConfig {
	if o == nil || o.LlmJudgeConfig == nil {
		var ret LLMObsCustomEvalConfigLLMJudgeConfig
		return ret
	}
	return *o.LlmJudgeConfig
}

// GetLlmJudgeConfigOk returns a tuple with the LlmJudgeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigUpdateAttributes) GetLlmJudgeConfigOk() (*LLMObsCustomEvalConfigLLMJudgeConfig, bool) {
	if o == nil || o.LlmJudgeConfig == nil {
		return nil, false
	}
	return o.LlmJudgeConfig, true
}

// HasLlmJudgeConfig returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigUpdateAttributes) HasLlmJudgeConfig() bool {
	return o != nil && o.LlmJudgeConfig != nil
}

// SetLlmJudgeConfig gets a reference to the given LLMObsCustomEvalConfigLLMJudgeConfig and assigns it to the LlmJudgeConfig field.
func (o *LLMObsCustomEvalConfigUpdateAttributes) SetLlmJudgeConfig(v LLMObsCustomEvalConfigLLMJudgeConfig) {
	o.LlmJudgeConfig = &v
}

// GetLlmProvider returns the LlmProvider field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigUpdateAttributes) GetLlmProvider() LLMObsCustomEvalConfigLLMProvider {
	if o == nil || o.LlmProvider == nil {
		var ret LLMObsCustomEvalConfigLLMProvider
		return ret
	}
	return *o.LlmProvider
}

// GetLlmProviderOk returns a tuple with the LlmProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigUpdateAttributes) GetLlmProviderOk() (*LLMObsCustomEvalConfigLLMProvider, bool) {
	if o == nil || o.LlmProvider == nil {
		return nil, false
	}
	return o.LlmProvider, true
}

// HasLlmProvider returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigUpdateAttributes) HasLlmProvider() bool {
	return o != nil && o.LlmProvider != nil
}

// SetLlmProvider gets a reference to the given LLMObsCustomEvalConfigLLMProvider and assigns it to the LlmProvider field.
func (o *LLMObsCustomEvalConfigUpdateAttributes) SetLlmProvider(v LLMObsCustomEvalConfigLLMProvider) {
	o.LlmProvider = &v
}

// GetTarget returns the Target field value.
func (o *LLMObsCustomEvalConfigUpdateAttributes) GetTarget() LLMObsCustomEvalConfigTarget {
	if o == nil {
		var ret LLMObsCustomEvalConfigTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigUpdateAttributes) GetTargetOk() (*LLMObsCustomEvalConfigTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *LLMObsCustomEvalConfigUpdateAttributes) SetTarget(v LLMObsCustomEvalConfigTarget) {
	o.Target = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCustomEvalConfigUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.EvalName != nil {
		toSerialize["eval_name"] = o.EvalName
	}
	if o.LlmJudgeConfig != nil {
		toSerialize["llm_judge_config"] = o.LlmJudgeConfig
	}
	if o.LlmProvider != nil {
		toSerialize["llm_provider"] = o.LlmProvider
	}
	toSerialize["target"] = o.Target

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCustomEvalConfigUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category       *string                               `json:"category,omitempty"`
		EvalName       *string                               `json:"eval_name,omitempty"`
		LlmJudgeConfig *LLMObsCustomEvalConfigLLMJudgeConfig `json:"llm_judge_config,omitempty"`
		LlmProvider    *LLMObsCustomEvalConfigLLMProvider    `json:"llm_provider,omitempty"`
		Target         *LLMObsCustomEvalConfigTarget         `json:"target"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "eval_name", "llm_judge_config", "llm_provider", "target"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Category = all.Category
	o.EvalName = all.EvalName
	if all.LlmJudgeConfig != nil && all.LlmJudgeConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LlmJudgeConfig = all.LlmJudgeConfig
	if all.LlmProvider != nil && all.LlmProvider.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LlmProvider = all.LlmProvider
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = *all.Target

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
