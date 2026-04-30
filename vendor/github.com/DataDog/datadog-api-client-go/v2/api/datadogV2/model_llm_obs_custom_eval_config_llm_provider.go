// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigLLMProvider LLM provider configuration for a custom evaluator.
type LLMObsCustomEvalConfigLLMProvider struct {
	// AWS Bedrock-specific options for LLM provider configuration.
	Bedrock *LLMObsCustomEvalConfigBedrockOptions `json:"bedrock,omitempty"`
	// Integration account identifier.
	IntegrationAccountId *string `json:"integration_account_id,omitempty"`
	// Name of the LLM integration provider.
	IntegrationProvider *LLMObsCustomEvalConfigIntegrationProvider `json:"integration_provider,omitempty"`
	// Name of the LLM model.
	ModelName *string `json:"model_name,omitempty"`
	// Google Vertex AI-specific options for LLM provider configuration.
	VertexAi *LLMObsCustomEvalConfigVertexAIOptions `json:"vertex_ai,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCustomEvalConfigLLMProvider instantiates a new LLMObsCustomEvalConfigLLMProvider object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCustomEvalConfigLLMProvider() *LLMObsCustomEvalConfigLLMProvider {
	this := LLMObsCustomEvalConfigLLMProvider{}
	return &this
}

// NewLLMObsCustomEvalConfigLLMProviderWithDefaults instantiates a new LLMObsCustomEvalConfigLLMProvider object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCustomEvalConfigLLMProviderWithDefaults() *LLMObsCustomEvalConfigLLMProvider {
	this := LLMObsCustomEvalConfigLLMProvider{}
	return &this
}

// GetBedrock returns the Bedrock field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigLLMProvider) GetBedrock() LLMObsCustomEvalConfigBedrockOptions {
	if o == nil || o.Bedrock == nil {
		var ret LLMObsCustomEvalConfigBedrockOptions
		return ret
	}
	return *o.Bedrock
}

// GetBedrockOk returns a tuple with the Bedrock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigLLMProvider) GetBedrockOk() (*LLMObsCustomEvalConfigBedrockOptions, bool) {
	if o == nil || o.Bedrock == nil {
		return nil, false
	}
	return o.Bedrock, true
}

// HasBedrock returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMProvider) HasBedrock() bool {
	return o != nil && o.Bedrock != nil
}

// SetBedrock gets a reference to the given LLMObsCustomEvalConfigBedrockOptions and assigns it to the Bedrock field.
func (o *LLMObsCustomEvalConfigLLMProvider) SetBedrock(v LLMObsCustomEvalConfigBedrockOptions) {
	o.Bedrock = &v
}

// GetIntegrationAccountId returns the IntegrationAccountId field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigLLMProvider) GetIntegrationAccountId() string {
	if o == nil || o.IntegrationAccountId == nil {
		var ret string
		return ret
	}
	return *o.IntegrationAccountId
}

// GetIntegrationAccountIdOk returns a tuple with the IntegrationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigLLMProvider) GetIntegrationAccountIdOk() (*string, bool) {
	if o == nil || o.IntegrationAccountId == nil {
		return nil, false
	}
	return o.IntegrationAccountId, true
}

// HasIntegrationAccountId returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMProvider) HasIntegrationAccountId() bool {
	return o != nil && o.IntegrationAccountId != nil
}

// SetIntegrationAccountId gets a reference to the given string and assigns it to the IntegrationAccountId field.
func (o *LLMObsCustomEvalConfigLLMProvider) SetIntegrationAccountId(v string) {
	o.IntegrationAccountId = &v
}

// GetIntegrationProvider returns the IntegrationProvider field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigLLMProvider) GetIntegrationProvider() LLMObsCustomEvalConfigIntegrationProvider {
	if o == nil || o.IntegrationProvider == nil {
		var ret LLMObsCustomEvalConfigIntegrationProvider
		return ret
	}
	return *o.IntegrationProvider
}

// GetIntegrationProviderOk returns a tuple with the IntegrationProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigLLMProvider) GetIntegrationProviderOk() (*LLMObsCustomEvalConfigIntegrationProvider, bool) {
	if o == nil || o.IntegrationProvider == nil {
		return nil, false
	}
	return o.IntegrationProvider, true
}

// HasIntegrationProvider returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMProvider) HasIntegrationProvider() bool {
	return o != nil && o.IntegrationProvider != nil
}

// SetIntegrationProvider gets a reference to the given LLMObsCustomEvalConfigIntegrationProvider and assigns it to the IntegrationProvider field.
func (o *LLMObsCustomEvalConfigLLMProvider) SetIntegrationProvider(v LLMObsCustomEvalConfigIntegrationProvider) {
	o.IntegrationProvider = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigLLMProvider) GetModelName() string {
	if o == nil || o.ModelName == nil {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigLLMProvider) GetModelNameOk() (*string, bool) {
	if o == nil || o.ModelName == nil {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMProvider) HasModelName() bool {
	return o != nil && o.ModelName != nil
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *LLMObsCustomEvalConfigLLMProvider) SetModelName(v string) {
	o.ModelName = &v
}

// GetVertexAi returns the VertexAi field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigLLMProvider) GetVertexAi() LLMObsCustomEvalConfigVertexAIOptions {
	if o == nil || o.VertexAi == nil {
		var ret LLMObsCustomEvalConfigVertexAIOptions
		return ret
	}
	return *o.VertexAi
}

// GetVertexAiOk returns a tuple with the VertexAi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigLLMProvider) GetVertexAiOk() (*LLMObsCustomEvalConfigVertexAIOptions, bool) {
	if o == nil || o.VertexAi == nil {
		return nil, false
	}
	return o.VertexAi, true
}

// HasVertexAi returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMProvider) HasVertexAi() bool {
	return o != nil && o.VertexAi != nil
}

// SetVertexAi gets a reference to the given LLMObsCustomEvalConfigVertexAIOptions and assigns it to the VertexAi field.
func (o *LLMObsCustomEvalConfigLLMProvider) SetVertexAi(v LLMObsCustomEvalConfigVertexAIOptions) {
	o.VertexAi = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCustomEvalConfigLLMProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Bedrock != nil {
		toSerialize["bedrock"] = o.Bedrock
	}
	if o.IntegrationAccountId != nil {
		toSerialize["integration_account_id"] = o.IntegrationAccountId
	}
	if o.IntegrationProvider != nil {
		toSerialize["integration_provider"] = o.IntegrationProvider
	}
	if o.ModelName != nil {
		toSerialize["model_name"] = o.ModelName
	}
	if o.VertexAi != nil {
		toSerialize["vertex_ai"] = o.VertexAi
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCustomEvalConfigLLMProvider) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bedrock              *LLMObsCustomEvalConfigBedrockOptions      `json:"bedrock,omitempty"`
		IntegrationAccountId *string                                    `json:"integration_account_id,omitempty"`
		IntegrationProvider  *LLMObsCustomEvalConfigIntegrationProvider `json:"integration_provider,omitempty"`
		ModelName            *string                                    `json:"model_name,omitempty"`
		VertexAi             *LLMObsCustomEvalConfigVertexAIOptions     `json:"vertex_ai,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bedrock", "integration_account_id", "integration_provider", "model_name", "vertex_ai"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Bedrock != nil && all.Bedrock.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Bedrock = all.Bedrock
	o.IntegrationAccountId = all.IntegrationAccountId
	if all.IntegrationProvider != nil && !all.IntegrationProvider.IsValid() {
		hasInvalidField = true
	} else {
		o.IntegrationProvider = all.IntegrationProvider
	}
	o.ModelName = all.ModelName
	if all.VertexAi != nil && all.VertexAi.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.VertexAi = all.VertexAi

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
