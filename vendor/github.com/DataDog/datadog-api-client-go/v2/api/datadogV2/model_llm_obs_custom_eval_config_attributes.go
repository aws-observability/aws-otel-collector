// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigAttributes Attributes of a custom LLM Observability evaluator configuration.
type LLMObsCustomEvalConfigAttributes struct {
	// Category of the evaluator.
	Category *string `json:"category,omitempty"`
	// Timestamp when the evaluator configuration was created.
	CreatedAt time.Time `json:"created_at"`
	// A Datadog user associated with a custom evaluator configuration.
	CreatedBy *LLMObsCustomEvalConfigUser `json:"created_by,omitempty"`
	// Name of the custom evaluator.
	EvalName string `json:"eval_name"`
	// A Datadog user associated with a custom evaluator configuration.
	LastUpdatedBy *LLMObsCustomEvalConfigUser `json:"last_updated_by,omitempty"`
	// LLM judge configuration for a custom evaluator.
	LlmJudgeConfig *LLMObsCustomEvalConfigLLMJudgeConfig `json:"llm_judge_config,omitempty"`
	// LLM provider configuration for a custom evaluator.
	LlmProvider *LLMObsCustomEvalConfigLLMProvider `json:"llm_provider,omitempty"`
	// Target application configuration for a custom evaluator.
	Target *LLMObsCustomEvalConfigTarget `json:"target,omitempty"`
	// Timestamp when the evaluator configuration was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCustomEvalConfigAttributes instantiates a new LLMObsCustomEvalConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCustomEvalConfigAttributes(createdAt time.Time, evalName string, updatedAt time.Time) *LLMObsCustomEvalConfigAttributes {
	this := LLMObsCustomEvalConfigAttributes{}
	this.CreatedAt = createdAt
	this.EvalName = evalName
	this.UpdatedAt = updatedAt
	return &this
}

// NewLLMObsCustomEvalConfigAttributesWithDefaults instantiates a new LLMObsCustomEvalConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCustomEvalConfigAttributesWithDefaults() *LLMObsCustomEvalConfigAttributes {
	this := LLMObsCustomEvalConfigAttributes{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigAttributes) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigAttributes) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigAttributes) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *LLMObsCustomEvalConfigAttributes) SetCategory(v string) {
	o.Category = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsCustomEvalConfigAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsCustomEvalConfigAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigAttributes) GetCreatedBy() LLMObsCustomEvalConfigUser {
	if o == nil || o.CreatedBy == nil {
		var ret LLMObsCustomEvalConfigUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigAttributes) GetCreatedByOk() (*LLMObsCustomEvalConfigUser, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given LLMObsCustomEvalConfigUser and assigns it to the CreatedBy field.
func (o *LLMObsCustomEvalConfigAttributes) SetCreatedBy(v LLMObsCustomEvalConfigUser) {
	o.CreatedBy = &v
}

// GetEvalName returns the EvalName field value.
func (o *LLMObsCustomEvalConfigAttributes) GetEvalName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvalName
}

// GetEvalNameOk returns a tuple with the EvalName field value
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigAttributes) GetEvalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvalName, true
}

// SetEvalName sets field value.
func (o *LLMObsCustomEvalConfigAttributes) SetEvalName(v string) {
	o.EvalName = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigAttributes) GetLastUpdatedBy() LLMObsCustomEvalConfigUser {
	if o == nil || o.LastUpdatedBy == nil {
		var ret LLMObsCustomEvalConfigUser
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigAttributes) GetLastUpdatedByOk() (*LLMObsCustomEvalConfigUser, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigAttributes) HasLastUpdatedBy() bool {
	return o != nil && o.LastUpdatedBy != nil
}

// SetLastUpdatedBy gets a reference to the given LLMObsCustomEvalConfigUser and assigns it to the LastUpdatedBy field.
func (o *LLMObsCustomEvalConfigAttributes) SetLastUpdatedBy(v LLMObsCustomEvalConfigUser) {
	o.LastUpdatedBy = &v
}

// GetLlmJudgeConfig returns the LlmJudgeConfig field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigAttributes) GetLlmJudgeConfig() LLMObsCustomEvalConfigLLMJudgeConfig {
	if o == nil || o.LlmJudgeConfig == nil {
		var ret LLMObsCustomEvalConfigLLMJudgeConfig
		return ret
	}
	return *o.LlmJudgeConfig
}

// GetLlmJudgeConfigOk returns a tuple with the LlmJudgeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigAttributes) GetLlmJudgeConfigOk() (*LLMObsCustomEvalConfigLLMJudgeConfig, bool) {
	if o == nil || o.LlmJudgeConfig == nil {
		return nil, false
	}
	return o.LlmJudgeConfig, true
}

// HasLlmJudgeConfig returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigAttributes) HasLlmJudgeConfig() bool {
	return o != nil && o.LlmJudgeConfig != nil
}

// SetLlmJudgeConfig gets a reference to the given LLMObsCustomEvalConfigLLMJudgeConfig and assigns it to the LlmJudgeConfig field.
func (o *LLMObsCustomEvalConfigAttributes) SetLlmJudgeConfig(v LLMObsCustomEvalConfigLLMJudgeConfig) {
	o.LlmJudgeConfig = &v
}

// GetLlmProvider returns the LlmProvider field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigAttributes) GetLlmProvider() LLMObsCustomEvalConfigLLMProvider {
	if o == nil || o.LlmProvider == nil {
		var ret LLMObsCustomEvalConfigLLMProvider
		return ret
	}
	return *o.LlmProvider
}

// GetLlmProviderOk returns a tuple with the LlmProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigAttributes) GetLlmProviderOk() (*LLMObsCustomEvalConfigLLMProvider, bool) {
	if o == nil || o.LlmProvider == nil {
		return nil, false
	}
	return o.LlmProvider, true
}

// HasLlmProvider returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigAttributes) HasLlmProvider() bool {
	return o != nil && o.LlmProvider != nil
}

// SetLlmProvider gets a reference to the given LLMObsCustomEvalConfigLLMProvider and assigns it to the LlmProvider field.
func (o *LLMObsCustomEvalConfigAttributes) SetLlmProvider(v LLMObsCustomEvalConfigLLMProvider) {
	o.LlmProvider = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigAttributes) GetTarget() LLMObsCustomEvalConfigTarget {
	if o == nil || o.Target == nil {
		var ret LLMObsCustomEvalConfigTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigAttributes) GetTargetOk() (*LLMObsCustomEvalConfigTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigAttributes) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given LLMObsCustomEvalConfigTarget and assigns it to the Target field.
func (o *LLMObsCustomEvalConfigAttributes) SetTarget(v LLMObsCustomEvalConfigTarget) {
	o.Target = &v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *LLMObsCustomEvalConfigAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *LLMObsCustomEvalConfigAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCustomEvalConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	toSerialize["eval_name"] = o.EvalName
	if o.LastUpdatedBy != nil {
		toSerialize["last_updated_by"] = o.LastUpdatedBy
	}
	if o.LlmJudgeConfig != nil {
		toSerialize["llm_judge_config"] = o.LlmJudgeConfig
	}
	if o.LlmProvider != nil {
		toSerialize["llm_provider"] = o.LlmProvider
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCustomEvalConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category       *string                               `json:"category,omitempty"`
		CreatedAt      *time.Time                            `json:"created_at"`
		CreatedBy      *LLMObsCustomEvalConfigUser           `json:"created_by,omitempty"`
		EvalName       *string                               `json:"eval_name"`
		LastUpdatedBy  *LLMObsCustomEvalConfigUser           `json:"last_updated_by,omitempty"`
		LlmJudgeConfig *LLMObsCustomEvalConfigLLMJudgeConfig `json:"llm_judge_config,omitempty"`
		LlmProvider    *LLMObsCustomEvalConfigLLMProvider    `json:"llm_provider,omitempty"`
		Target         *LLMObsCustomEvalConfigTarget         `json:"target,omitempty"`
		UpdatedAt      *time.Time                            `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.EvalName == nil {
		return fmt.Errorf("required field eval_name missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "created_at", "created_by", "eval_name", "last_updated_by", "llm_judge_config", "llm_provider", "target", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Category = all.Category
	o.CreatedAt = *all.CreatedAt
	if all.CreatedBy != nil && all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = all.CreatedBy
	o.EvalName = *all.EvalName
	if all.LastUpdatedBy != nil && all.LastUpdatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastUpdatedBy = all.LastUpdatedBy
	if all.LlmJudgeConfig != nil && all.LlmJudgeConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LlmJudgeConfig = all.LlmJudgeConfig
	if all.LlmProvider != nil && all.LlmProvider.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LlmProvider = all.LlmProvider
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
