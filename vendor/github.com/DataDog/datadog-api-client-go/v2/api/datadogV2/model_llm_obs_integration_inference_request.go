// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsIntegrationInferenceRequest Parameters for an LLM inference request.
type LLMObsIntegrationInferenceRequest struct {
	// Anthropic-specific metadata for an inference request.
	AnthropicMetadata *LLMObsAnthropicMetadata `json:"anthropic_metadata,omitempty"`
	// Azure OpenAI-specific metadata for an integration account or inference request.
	AzureOpenaiMetadata *LLMObsAzureOpenAIMetadata `json:"azure_openai_metadata,omitempty"`
	// Amazon Bedrock-specific metadata for an inference request.
	BedrockMetadata *LLMObsBedrockMetadata `json:"bedrock_metadata,omitempty"`
	// Penalty for token frequency to reduce repetition.
	FrequencyPenalty datadog.NullableFloat64 `json:"frequency_penalty,omitempty"`
	// JSON schema for structured output, if supported by the model.
	JsonSchema datadog.NullableString `json:"json_schema,omitempty"`
	// Maximum number of completion tokens to generate (alternative to max_tokens for some providers).
	MaxCompletionTokens datadog.NullableInt64 `json:"max_completion_tokens,omitempty"`
	// Maximum number of tokens to generate.
	MaxTokens datadog.NullableInt64 `json:"max_tokens,omitempty"`
	// List of messages in an inference conversation.
	Messages []LLMObsInferenceMessage `json:"messages"`
	// The model identifier to use for inference.
	ModelId string `json:"model_id"`
	// OpenAI-specific metadata for an inference request.
	OpenaiMetadata *LLMObsOpenAIMetadata `json:"openai_metadata,omitempty"`
	// Penalty for token presence to encourage topic diversity.
	PresencePenalty datadog.NullableFloat64 `json:"presence_penalty,omitempty"`
	// Sampling temperature between 0 and 2. Higher values produce more random output.
	Temperature datadog.NullableFloat64 `json:"temperature,omitempty"`
	// List of tools available to the model.
	Tools []LLMObsInferenceTool `json:"tools,omitempty"`
	// Top-K sampling parameter.
	TopK datadog.NullableInt64 `json:"top_k,omitempty"`
	// Nucleus sampling probability mass.
	TopP datadog.NullableFloat64 `json:"top_p,omitempty"`
	// Vertex AI-specific metadata for an integration account or inference request.
	VertexAiMetadata *LLMObsVertexAIMetadata `json:"vertex_ai_metadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsIntegrationInferenceRequest instantiates a new LLMObsIntegrationInferenceRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsIntegrationInferenceRequest(messages []LLMObsInferenceMessage, modelId string) *LLMObsIntegrationInferenceRequest {
	this := LLMObsIntegrationInferenceRequest{}
	this.Messages = messages
	this.ModelId = modelId
	return &this
}

// NewLLMObsIntegrationInferenceRequestWithDefaults instantiates a new LLMObsIntegrationInferenceRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsIntegrationInferenceRequestWithDefaults() *LLMObsIntegrationInferenceRequest {
	this := LLMObsIntegrationInferenceRequest{}
	return &this
}

// GetAnthropicMetadata returns the AnthropicMetadata field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceRequest) GetAnthropicMetadata() LLMObsAnthropicMetadata {
	if o == nil || o.AnthropicMetadata == nil {
		var ret LLMObsAnthropicMetadata
		return ret
	}
	return *o.AnthropicMetadata
}

// GetAnthropicMetadataOk returns a tuple with the AnthropicMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceRequest) GetAnthropicMetadataOk() (*LLMObsAnthropicMetadata, bool) {
	if o == nil || o.AnthropicMetadata == nil {
		return nil, false
	}
	return o.AnthropicMetadata, true
}

// HasAnthropicMetadata returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasAnthropicMetadata() bool {
	return o != nil && o.AnthropicMetadata != nil
}

// SetAnthropicMetadata gets a reference to the given LLMObsAnthropicMetadata and assigns it to the AnthropicMetadata field.
func (o *LLMObsIntegrationInferenceRequest) SetAnthropicMetadata(v LLMObsAnthropicMetadata) {
	o.AnthropicMetadata = &v
}

// GetAzureOpenaiMetadata returns the AzureOpenaiMetadata field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceRequest) GetAzureOpenaiMetadata() LLMObsAzureOpenAIMetadata {
	if o == nil || o.AzureOpenaiMetadata == nil {
		var ret LLMObsAzureOpenAIMetadata
		return ret
	}
	return *o.AzureOpenaiMetadata
}

// GetAzureOpenaiMetadataOk returns a tuple with the AzureOpenaiMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceRequest) GetAzureOpenaiMetadataOk() (*LLMObsAzureOpenAIMetadata, bool) {
	if o == nil || o.AzureOpenaiMetadata == nil {
		return nil, false
	}
	return o.AzureOpenaiMetadata, true
}

// HasAzureOpenaiMetadata returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasAzureOpenaiMetadata() bool {
	return o != nil && o.AzureOpenaiMetadata != nil
}

// SetAzureOpenaiMetadata gets a reference to the given LLMObsAzureOpenAIMetadata and assigns it to the AzureOpenaiMetadata field.
func (o *LLMObsIntegrationInferenceRequest) SetAzureOpenaiMetadata(v LLMObsAzureOpenAIMetadata) {
	o.AzureOpenaiMetadata = &v
}

// GetBedrockMetadata returns the BedrockMetadata field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceRequest) GetBedrockMetadata() LLMObsBedrockMetadata {
	if o == nil || o.BedrockMetadata == nil {
		var ret LLMObsBedrockMetadata
		return ret
	}
	return *o.BedrockMetadata
}

// GetBedrockMetadataOk returns a tuple with the BedrockMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceRequest) GetBedrockMetadataOk() (*LLMObsBedrockMetadata, bool) {
	if o == nil || o.BedrockMetadata == nil {
		return nil, false
	}
	return o.BedrockMetadata, true
}

// HasBedrockMetadata returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasBedrockMetadata() bool {
	return o != nil && o.BedrockMetadata != nil
}

// SetBedrockMetadata gets a reference to the given LLMObsBedrockMetadata and assigns it to the BedrockMetadata field.
func (o *LLMObsIntegrationInferenceRequest) SetBedrockMetadata(v LLMObsBedrockMetadata) {
	o.BedrockMetadata = &v
}

// GetFrequencyPenalty returns the FrequencyPenalty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceRequest) GetFrequencyPenalty() float64 {
	if o == nil || o.FrequencyPenalty.Get() == nil {
		var ret float64
		return ret
	}
	return *o.FrequencyPenalty.Get()
}

// GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceRequest) GetFrequencyPenaltyOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrequencyPenalty.Get(), o.FrequencyPenalty.IsSet()
}

// HasFrequencyPenalty returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasFrequencyPenalty() bool {
	return o != nil && o.FrequencyPenalty.IsSet()
}

// SetFrequencyPenalty gets a reference to the given datadog.NullableFloat64 and assigns it to the FrequencyPenalty field.
func (o *LLMObsIntegrationInferenceRequest) SetFrequencyPenalty(v float64) {
	o.FrequencyPenalty.Set(&v)
}

// SetFrequencyPenaltyNil sets the value for FrequencyPenalty to be an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) SetFrequencyPenaltyNil() {
	o.FrequencyPenalty.Set(nil)
}

// UnsetFrequencyPenalty ensures that no value is present for FrequencyPenalty, not even an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) UnsetFrequencyPenalty() {
	o.FrequencyPenalty.Unset()
}

// GetJsonSchema returns the JsonSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceRequest) GetJsonSchema() string {
	if o == nil || o.JsonSchema.Get() == nil {
		var ret string
		return ret
	}
	return *o.JsonSchema.Get()
}

// GetJsonSchemaOk returns a tuple with the JsonSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceRequest) GetJsonSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JsonSchema.Get(), o.JsonSchema.IsSet()
}

// HasJsonSchema returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasJsonSchema() bool {
	return o != nil && o.JsonSchema.IsSet()
}

// SetJsonSchema gets a reference to the given datadog.NullableString and assigns it to the JsonSchema field.
func (o *LLMObsIntegrationInferenceRequest) SetJsonSchema(v string) {
	o.JsonSchema.Set(&v)
}

// SetJsonSchemaNil sets the value for JsonSchema to be an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) SetJsonSchemaNil() {
	o.JsonSchema.Set(nil)
}

// UnsetJsonSchema ensures that no value is present for JsonSchema, not even an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) UnsetJsonSchema() {
	o.JsonSchema.Unset()
}

// GetMaxCompletionTokens returns the MaxCompletionTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceRequest) GetMaxCompletionTokens() int64 {
	if o == nil || o.MaxCompletionTokens.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MaxCompletionTokens.Get()
}

// GetMaxCompletionTokensOk returns a tuple with the MaxCompletionTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceRequest) GetMaxCompletionTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCompletionTokens.Get(), o.MaxCompletionTokens.IsSet()
}

// HasMaxCompletionTokens returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasMaxCompletionTokens() bool {
	return o != nil && o.MaxCompletionTokens.IsSet()
}

// SetMaxCompletionTokens gets a reference to the given datadog.NullableInt64 and assigns it to the MaxCompletionTokens field.
func (o *LLMObsIntegrationInferenceRequest) SetMaxCompletionTokens(v int64) {
	o.MaxCompletionTokens.Set(&v)
}

// SetMaxCompletionTokensNil sets the value for MaxCompletionTokens to be an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) SetMaxCompletionTokensNil() {
	o.MaxCompletionTokens.Set(nil)
}

// UnsetMaxCompletionTokens ensures that no value is present for MaxCompletionTokens, not even an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) UnsetMaxCompletionTokens() {
	o.MaxCompletionTokens.Unset()
}

// GetMaxTokens returns the MaxTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceRequest) GetMaxTokens() int64 {
	if o == nil || o.MaxTokens.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MaxTokens.Get()
}

// GetMaxTokensOk returns a tuple with the MaxTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceRequest) GetMaxTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTokens.Get(), o.MaxTokens.IsSet()
}

// HasMaxTokens returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasMaxTokens() bool {
	return o != nil && o.MaxTokens.IsSet()
}

// SetMaxTokens gets a reference to the given datadog.NullableInt64 and assigns it to the MaxTokens field.
func (o *LLMObsIntegrationInferenceRequest) SetMaxTokens(v int64) {
	o.MaxTokens.Set(&v)
}

// SetMaxTokensNil sets the value for MaxTokens to be an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) SetMaxTokensNil() {
	o.MaxTokens.Set(nil)
}

// UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) UnsetMaxTokens() {
	o.MaxTokens.Unset()
}

// GetMessages returns the Messages field value.
func (o *LLMObsIntegrationInferenceRequest) GetMessages() []LLMObsInferenceMessage {
	if o == nil {
		var ret []LLMObsInferenceMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceRequest) GetMessagesOk() (*[]LLMObsInferenceMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Messages, true
}

// SetMessages sets field value.
func (o *LLMObsIntegrationInferenceRequest) SetMessages(v []LLMObsInferenceMessage) {
	o.Messages = v
}

// GetModelId returns the ModelId field value.
func (o *LLMObsIntegrationInferenceRequest) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceRequest) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value.
func (o *LLMObsIntegrationInferenceRequest) SetModelId(v string) {
	o.ModelId = v
}

// GetOpenaiMetadata returns the OpenaiMetadata field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceRequest) GetOpenaiMetadata() LLMObsOpenAIMetadata {
	if o == nil || o.OpenaiMetadata == nil {
		var ret LLMObsOpenAIMetadata
		return ret
	}
	return *o.OpenaiMetadata
}

// GetOpenaiMetadataOk returns a tuple with the OpenaiMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceRequest) GetOpenaiMetadataOk() (*LLMObsOpenAIMetadata, bool) {
	if o == nil || o.OpenaiMetadata == nil {
		return nil, false
	}
	return o.OpenaiMetadata, true
}

// HasOpenaiMetadata returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasOpenaiMetadata() bool {
	return o != nil && o.OpenaiMetadata != nil
}

// SetOpenaiMetadata gets a reference to the given LLMObsOpenAIMetadata and assigns it to the OpenaiMetadata field.
func (o *LLMObsIntegrationInferenceRequest) SetOpenaiMetadata(v LLMObsOpenAIMetadata) {
	o.OpenaiMetadata = &v
}

// GetPresencePenalty returns the PresencePenalty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceRequest) GetPresencePenalty() float64 {
	if o == nil || o.PresencePenalty.Get() == nil {
		var ret float64
		return ret
	}
	return *o.PresencePenalty.Get()
}

// GetPresencePenaltyOk returns a tuple with the PresencePenalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceRequest) GetPresencePenaltyOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PresencePenalty.Get(), o.PresencePenalty.IsSet()
}

// HasPresencePenalty returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasPresencePenalty() bool {
	return o != nil && o.PresencePenalty.IsSet()
}

// SetPresencePenalty gets a reference to the given datadog.NullableFloat64 and assigns it to the PresencePenalty field.
func (o *LLMObsIntegrationInferenceRequest) SetPresencePenalty(v float64) {
	o.PresencePenalty.Set(&v)
}

// SetPresencePenaltyNil sets the value for PresencePenalty to be an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) SetPresencePenaltyNil() {
	o.PresencePenalty.Set(nil)
}

// UnsetPresencePenalty ensures that no value is present for PresencePenalty, not even an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) UnsetPresencePenalty() {
	o.PresencePenalty.Unset()
}

// GetTemperature returns the Temperature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceRequest) GetTemperature() float64 {
	if o == nil || o.Temperature.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Temperature.Get()
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceRequest) GetTemperatureOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Temperature.Get(), o.Temperature.IsSet()
}

// HasTemperature returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasTemperature() bool {
	return o != nil && o.Temperature.IsSet()
}

// SetTemperature gets a reference to the given datadog.NullableFloat64 and assigns it to the Temperature field.
func (o *LLMObsIntegrationInferenceRequest) SetTemperature(v float64) {
	o.Temperature.Set(&v)
}

// SetTemperatureNil sets the value for Temperature to be an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) SetTemperatureNil() {
	o.Temperature.Set(nil)
}

// UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) UnsetTemperature() {
	o.Temperature.Unset()
}

// GetTools returns the Tools field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceRequest) GetTools() []LLMObsInferenceTool {
	if o == nil || o.Tools == nil {
		var ret []LLMObsInferenceTool
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceRequest) GetToolsOk() (*[]LLMObsInferenceTool, bool) {
	if o == nil || o.Tools == nil {
		return nil, false
	}
	return &o.Tools, true
}

// HasTools returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasTools() bool {
	return o != nil && o.Tools != nil
}

// SetTools gets a reference to the given []LLMObsInferenceTool and assigns it to the Tools field.
func (o *LLMObsIntegrationInferenceRequest) SetTools(v []LLMObsInferenceTool) {
	o.Tools = v
}

// GetTopK returns the TopK field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceRequest) GetTopK() int64 {
	if o == nil || o.TopK.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TopK.Get()
}

// GetTopKOk returns a tuple with the TopK field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceRequest) GetTopKOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopK.Get(), o.TopK.IsSet()
}

// HasTopK returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasTopK() bool {
	return o != nil && o.TopK.IsSet()
}

// SetTopK gets a reference to the given datadog.NullableInt64 and assigns it to the TopK field.
func (o *LLMObsIntegrationInferenceRequest) SetTopK(v int64) {
	o.TopK.Set(&v)
}

// SetTopKNil sets the value for TopK to be an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) SetTopKNil() {
	o.TopK.Set(nil)
}

// UnsetTopK ensures that no value is present for TopK, not even an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) UnsetTopK() {
	o.TopK.Unset()
}

// GetTopP returns the TopP field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceRequest) GetTopP() float64 {
	if o == nil || o.TopP.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TopP.Get()
}

// GetTopPOk returns a tuple with the TopP field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceRequest) GetTopPOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopP.Get(), o.TopP.IsSet()
}

// HasTopP returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasTopP() bool {
	return o != nil && o.TopP.IsSet()
}

// SetTopP gets a reference to the given datadog.NullableFloat64 and assigns it to the TopP field.
func (o *LLMObsIntegrationInferenceRequest) SetTopP(v float64) {
	o.TopP.Set(&v)
}

// SetTopPNil sets the value for TopP to be an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) SetTopPNil() {
	o.TopP.Set(nil)
}

// UnsetTopP ensures that no value is present for TopP, not even an explicit nil.
func (o *LLMObsIntegrationInferenceRequest) UnsetTopP() {
	o.TopP.Unset()
}

// GetVertexAiMetadata returns the VertexAiMetadata field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceRequest) GetVertexAiMetadata() LLMObsVertexAIMetadata {
	if o == nil || o.VertexAiMetadata == nil {
		var ret LLMObsVertexAIMetadata
		return ret
	}
	return *o.VertexAiMetadata
}

// GetVertexAiMetadataOk returns a tuple with the VertexAiMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceRequest) GetVertexAiMetadataOk() (*LLMObsVertexAIMetadata, bool) {
	if o == nil || o.VertexAiMetadata == nil {
		return nil, false
	}
	return o.VertexAiMetadata, true
}

// HasVertexAiMetadata returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceRequest) HasVertexAiMetadata() bool {
	return o != nil && o.VertexAiMetadata != nil
}

// SetVertexAiMetadata gets a reference to the given LLMObsVertexAIMetadata and assigns it to the VertexAiMetadata field.
func (o *LLMObsIntegrationInferenceRequest) SetVertexAiMetadata(v LLMObsVertexAIMetadata) {
	o.VertexAiMetadata = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsIntegrationInferenceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnthropicMetadata != nil {
		toSerialize["anthropic_metadata"] = o.AnthropicMetadata
	}
	if o.AzureOpenaiMetadata != nil {
		toSerialize["azure_openai_metadata"] = o.AzureOpenaiMetadata
	}
	if o.BedrockMetadata != nil {
		toSerialize["bedrock_metadata"] = o.BedrockMetadata
	}
	if o.FrequencyPenalty.IsSet() {
		toSerialize["frequency_penalty"] = o.FrequencyPenalty.Get()
	}
	if o.JsonSchema.IsSet() {
		toSerialize["json_schema"] = o.JsonSchema.Get()
	}
	if o.MaxCompletionTokens.IsSet() {
		toSerialize["max_completion_tokens"] = o.MaxCompletionTokens.Get()
	}
	if o.MaxTokens.IsSet() {
		toSerialize["max_tokens"] = o.MaxTokens.Get()
	}
	toSerialize["messages"] = o.Messages
	toSerialize["model_id"] = o.ModelId
	if o.OpenaiMetadata != nil {
		toSerialize["openai_metadata"] = o.OpenaiMetadata
	}
	if o.PresencePenalty.IsSet() {
		toSerialize["presence_penalty"] = o.PresencePenalty.Get()
	}
	if o.Temperature.IsSet() {
		toSerialize["temperature"] = o.Temperature.Get()
	}
	if o.Tools != nil {
		toSerialize["tools"] = o.Tools
	}
	if o.TopK.IsSet() {
		toSerialize["top_k"] = o.TopK.Get()
	}
	if o.TopP.IsSet() {
		toSerialize["top_p"] = o.TopP.Get()
	}
	if o.VertexAiMetadata != nil {
		toSerialize["vertex_ai_metadata"] = o.VertexAiMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsIntegrationInferenceRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnthropicMetadata   *LLMObsAnthropicMetadata   `json:"anthropic_metadata,omitempty"`
		AzureOpenaiMetadata *LLMObsAzureOpenAIMetadata `json:"azure_openai_metadata,omitempty"`
		BedrockMetadata     *LLMObsBedrockMetadata     `json:"bedrock_metadata,omitempty"`
		FrequencyPenalty    datadog.NullableFloat64    `json:"frequency_penalty,omitempty"`
		JsonSchema          datadog.NullableString     `json:"json_schema,omitempty"`
		MaxCompletionTokens datadog.NullableInt64      `json:"max_completion_tokens,omitempty"`
		MaxTokens           datadog.NullableInt64      `json:"max_tokens,omitempty"`
		Messages            *[]LLMObsInferenceMessage  `json:"messages"`
		ModelId             *string                    `json:"model_id"`
		OpenaiMetadata      *LLMObsOpenAIMetadata      `json:"openai_metadata,omitempty"`
		PresencePenalty     datadog.NullableFloat64    `json:"presence_penalty,omitempty"`
		Temperature         datadog.NullableFloat64    `json:"temperature,omitempty"`
		Tools               []LLMObsInferenceTool      `json:"tools,omitempty"`
		TopK                datadog.NullableInt64      `json:"top_k,omitempty"`
		TopP                datadog.NullableFloat64    `json:"top_p,omitempty"`
		VertexAiMetadata    *LLMObsVertexAIMetadata    `json:"vertex_ai_metadata,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Messages == nil {
		return fmt.Errorf("required field messages missing")
	}
	if all.ModelId == nil {
		return fmt.Errorf("required field model_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"anthropic_metadata", "azure_openai_metadata", "bedrock_metadata", "frequency_penalty", "json_schema", "max_completion_tokens", "max_tokens", "messages", "model_id", "openai_metadata", "presence_penalty", "temperature", "tools", "top_k", "top_p", "vertex_ai_metadata"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AnthropicMetadata != nil && all.AnthropicMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AnthropicMetadata = all.AnthropicMetadata
	if all.AzureOpenaiMetadata != nil && all.AzureOpenaiMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AzureOpenaiMetadata = all.AzureOpenaiMetadata
	if all.BedrockMetadata != nil && all.BedrockMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BedrockMetadata = all.BedrockMetadata
	o.FrequencyPenalty = all.FrequencyPenalty
	o.JsonSchema = all.JsonSchema
	o.MaxCompletionTokens = all.MaxCompletionTokens
	o.MaxTokens = all.MaxTokens
	o.Messages = *all.Messages
	o.ModelId = *all.ModelId
	if all.OpenaiMetadata != nil && all.OpenaiMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpenaiMetadata = all.OpenaiMetadata
	o.PresencePenalty = all.PresencePenalty
	o.Temperature = all.Temperature
	o.Tools = all.Tools
	o.TopK = all.TopK
	o.TopP = all.TopP
	if all.VertexAiMetadata != nil && all.VertexAiMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.VertexAiMetadata = all.VertexAiMetadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
