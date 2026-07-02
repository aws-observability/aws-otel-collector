// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsIntegrationInferenceResponse The result of an LLM inference request, including input parameters and the model response.
type LLMObsIntegrationInferenceResponse struct {
	// Anthropic-specific metadata for an inference request.
	AnthropicMetadata *LLMObsAnthropicMetadata `json:"anthropic_metadata,omitempty"`
	// Azure OpenAI-specific metadata for an integration account or inference request.
	AzureOpenaiMetadata *LLMObsAzureOpenAIMetadata `json:"azure_openai_metadata,omitempty"`
	// Amazon Bedrock-specific metadata for an inference request.
	BedrockMetadata *LLMObsBedrockMetadata `json:"bedrock_metadata,omitempty"`
	// Error details returned when an inference provider returns an error.
	ErrorResponse *LLMObsInferenceErrorResponse `json:"error_response,omitempty"`
	// Frequency penalty that was applied.
	FrequencyPenalty datadog.NullableFloat64 `json:"frequency_penalty,omitempty"`
	// JSON schema that was applied for structured output.
	JsonSchema datadog.NullableString `json:"json_schema,omitempty"`
	// Maximum number of completion tokens that were configured.
	MaxCompletionTokens datadog.NullableInt64 `json:"max_completion_tokens,omitempty"`
	// Maximum number of tokens that were configured.
	MaxTokens datadog.NullableInt64 `json:"max_tokens,omitempty"`
	// List of messages in an inference conversation.
	Messages []LLMObsInferenceMessage `json:"messages"`
	// The model identifier used for inference.
	ModelId string `json:"model_id"`
	// OpenAI-specific metadata for an inference request.
	OpenaiMetadata *LLMObsOpenAIMetadata `json:"openai_metadata,omitempty"`
	// Presence penalty that was applied.
	PresencePenalty datadog.NullableFloat64 `json:"presence_penalty,omitempty"`
	// The output of a completed LLM inference call.
	Response LLMObsInferenceRunResult `json:"response"`
	// Sampling temperature that was used.
	Temperature datadog.NullableFloat64 `json:"temperature,omitempty"`
	// List of tools available to the model.
	Tools []LLMObsInferenceTool `json:"tools,omitempty"`
	// Top-K sampling parameter that was used.
	TopK datadog.NullableInt64 `json:"top_k,omitempty"`
	// Nucleus sampling parameter that was used.
	TopP datadog.NullableFloat64 `json:"top_p,omitempty"`
	// Vertex AI-specific metadata for an integration account or inference request.
	VertexAiMetadata *LLMObsVertexAIMetadata `json:"vertex_ai_metadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsIntegrationInferenceResponse instantiates a new LLMObsIntegrationInferenceResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsIntegrationInferenceResponse(messages []LLMObsInferenceMessage, modelId string, response LLMObsInferenceRunResult) *LLMObsIntegrationInferenceResponse {
	this := LLMObsIntegrationInferenceResponse{}
	this.Messages = messages
	this.ModelId = modelId
	this.Response = response
	return &this
}

// NewLLMObsIntegrationInferenceResponseWithDefaults instantiates a new LLMObsIntegrationInferenceResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsIntegrationInferenceResponseWithDefaults() *LLMObsIntegrationInferenceResponse {
	this := LLMObsIntegrationInferenceResponse{}
	return &this
}

// GetAnthropicMetadata returns the AnthropicMetadata field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceResponse) GetAnthropicMetadata() LLMObsAnthropicMetadata {
	if o == nil || o.AnthropicMetadata == nil {
		var ret LLMObsAnthropicMetadata
		return ret
	}
	return *o.AnthropicMetadata
}

// GetAnthropicMetadataOk returns a tuple with the AnthropicMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceResponse) GetAnthropicMetadataOk() (*LLMObsAnthropicMetadata, bool) {
	if o == nil || o.AnthropicMetadata == nil {
		return nil, false
	}
	return o.AnthropicMetadata, true
}

// HasAnthropicMetadata returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasAnthropicMetadata() bool {
	return o != nil && o.AnthropicMetadata != nil
}

// SetAnthropicMetadata gets a reference to the given LLMObsAnthropicMetadata and assigns it to the AnthropicMetadata field.
func (o *LLMObsIntegrationInferenceResponse) SetAnthropicMetadata(v LLMObsAnthropicMetadata) {
	o.AnthropicMetadata = &v
}

// GetAzureOpenaiMetadata returns the AzureOpenaiMetadata field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceResponse) GetAzureOpenaiMetadata() LLMObsAzureOpenAIMetadata {
	if o == nil || o.AzureOpenaiMetadata == nil {
		var ret LLMObsAzureOpenAIMetadata
		return ret
	}
	return *o.AzureOpenaiMetadata
}

// GetAzureOpenaiMetadataOk returns a tuple with the AzureOpenaiMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceResponse) GetAzureOpenaiMetadataOk() (*LLMObsAzureOpenAIMetadata, bool) {
	if o == nil || o.AzureOpenaiMetadata == nil {
		return nil, false
	}
	return o.AzureOpenaiMetadata, true
}

// HasAzureOpenaiMetadata returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasAzureOpenaiMetadata() bool {
	return o != nil && o.AzureOpenaiMetadata != nil
}

// SetAzureOpenaiMetadata gets a reference to the given LLMObsAzureOpenAIMetadata and assigns it to the AzureOpenaiMetadata field.
func (o *LLMObsIntegrationInferenceResponse) SetAzureOpenaiMetadata(v LLMObsAzureOpenAIMetadata) {
	o.AzureOpenaiMetadata = &v
}

// GetBedrockMetadata returns the BedrockMetadata field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceResponse) GetBedrockMetadata() LLMObsBedrockMetadata {
	if o == nil || o.BedrockMetadata == nil {
		var ret LLMObsBedrockMetadata
		return ret
	}
	return *o.BedrockMetadata
}

// GetBedrockMetadataOk returns a tuple with the BedrockMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceResponse) GetBedrockMetadataOk() (*LLMObsBedrockMetadata, bool) {
	if o == nil || o.BedrockMetadata == nil {
		return nil, false
	}
	return o.BedrockMetadata, true
}

// HasBedrockMetadata returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasBedrockMetadata() bool {
	return o != nil && o.BedrockMetadata != nil
}

// SetBedrockMetadata gets a reference to the given LLMObsBedrockMetadata and assigns it to the BedrockMetadata field.
func (o *LLMObsIntegrationInferenceResponse) SetBedrockMetadata(v LLMObsBedrockMetadata) {
	o.BedrockMetadata = &v
}

// GetErrorResponse returns the ErrorResponse field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceResponse) GetErrorResponse() LLMObsInferenceErrorResponse {
	if o == nil || o.ErrorResponse == nil {
		var ret LLMObsInferenceErrorResponse
		return ret
	}
	return *o.ErrorResponse
}

// GetErrorResponseOk returns a tuple with the ErrorResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceResponse) GetErrorResponseOk() (*LLMObsInferenceErrorResponse, bool) {
	if o == nil || o.ErrorResponse == nil {
		return nil, false
	}
	return o.ErrorResponse, true
}

// HasErrorResponse returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasErrorResponse() bool {
	return o != nil && o.ErrorResponse != nil
}

// SetErrorResponse gets a reference to the given LLMObsInferenceErrorResponse and assigns it to the ErrorResponse field.
func (o *LLMObsIntegrationInferenceResponse) SetErrorResponse(v LLMObsInferenceErrorResponse) {
	o.ErrorResponse = &v
}

// GetFrequencyPenalty returns the FrequencyPenalty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceResponse) GetFrequencyPenalty() float64 {
	if o == nil || o.FrequencyPenalty.Get() == nil {
		var ret float64
		return ret
	}
	return *o.FrequencyPenalty.Get()
}

// GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceResponse) GetFrequencyPenaltyOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrequencyPenalty.Get(), o.FrequencyPenalty.IsSet()
}

// HasFrequencyPenalty returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasFrequencyPenalty() bool {
	return o != nil && o.FrequencyPenalty.IsSet()
}

// SetFrequencyPenalty gets a reference to the given datadog.NullableFloat64 and assigns it to the FrequencyPenalty field.
func (o *LLMObsIntegrationInferenceResponse) SetFrequencyPenalty(v float64) {
	o.FrequencyPenalty.Set(&v)
}

// SetFrequencyPenaltyNil sets the value for FrequencyPenalty to be an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) SetFrequencyPenaltyNil() {
	o.FrequencyPenalty.Set(nil)
}

// UnsetFrequencyPenalty ensures that no value is present for FrequencyPenalty, not even an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) UnsetFrequencyPenalty() {
	o.FrequencyPenalty.Unset()
}

// GetJsonSchema returns the JsonSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceResponse) GetJsonSchema() string {
	if o == nil || o.JsonSchema.Get() == nil {
		var ret string
		return ret
	}
	return *o.JsonSchema.Get()
}

// GetJsonSchemaOk returns a tuple with the JsonSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceResponse) GetJsonSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JsonSchema.Get(), o.JsonSchema.IsSet()
}

// HasJsonSchema returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasJsonSchema() bool {
	return o != nil && o.JsonSchema.IsSet()
}

// SetJsonSchema gets a reference to the given datadog.NullableString and assigns it to the JsonSchema field.
func (o *LLMObsIntegrationInferenceResponse) SetJsonSchema(v string) {
	o.JsonSchema.Set(&v)
}

// SetJsonSchemaNil sets the value for JsonSchema to be an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) SetJsonSchemaNil() {
	o.JsonSchema.Set(nil)
}

// UnsetJsonSchema ensures that no value is present for JsonSchema, not even an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) UnsetJsonSchema() {
	o.JsonSchema.Unset()
}

// GetMaxCompletionTokens returns the MaxCompletionTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceResponse) GetMaxCompletionTokens() int64 {
	if o == nil || o.MaxCompletionTokens.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MaxCompletionTokens.Get()
}

// GetMaxCompletionTokensOk returns a tuple with the MaxCompletionTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceResponse) GetMaxCompletionTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCompletionTokens.Get(), o.MaxCompletionTokens.IsSet()
}

// HasMaxCompletionTokens returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasMaxCompletionTokens() bool {
	return o != nil && o.MaxCompletionTokens.IsSet()
}

// SetMaxCompletionTokens gets a reference to the given datadog.NullableInt64 and assigns it to the MaxCompletionTokens field.
func (o *LLMObsIntegrationInferenceResponse) SetMaxCompletionTokens(v int64) {
	o.MaxCompletionTokens.Set(&v)
}

// SetMaxCompletionTokensNil sets the value for MaxCompletionTokens to be an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) SetMaxCompletionTokensNil() {
	o.MaxCompletionTokens.Set(nil)
}

// UnsetMaxCompletionTokens ensures that no value is present for MaxCompletionTokens, not even an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) UnsetMaxCompletionTokens() {
	o.MaxCompletionTokens.Unset()
}

// GetMaxTokens returns the MaxTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceResponse) GetMaxTokens() int64 {
	if o == nil || o.MaxTokens.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MaxTokens.Get()
}

// GetMaxTokensOk returns a tuple with the MaxTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceResponse) GetMaxTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTokens.Get(), o.MaxTokens.IsSet()
}

// HasMaxTokens returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasMaxTokens() bool {
	return o != nil && o.MaxTokens.IsSet()
}

// SetMaxTokens gets a reference to the given datadog.NullableInt64 and assigns it to the MaxTokens field.
func (o *LLMObsIntegrationInferenceResponse) SetMaxTokens(v int64) {
	o.MaxTokens.Set(&v)
}

// SetMaxTokensNil sets the value for MaxTokens to be an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) SetMaxTokensNil() {
	o.MaxTokens.Set(nil)
}

// UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) UnsetMaxTokens() {
	o.MaxTokens.Unset()
}

// GetMessages returns the Messages field value.
func (o *LLMObsIntegrationInferenceResponse) GetMessages() []LLMObsInferenceMessage {
	if o == nil {
		var ret []LLMObsInferenceMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceResponse) GetMessagesOk() (*[]LLMObsInferenceMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Messages, true
}

// SetMessages sets field value.
func (o *LLMObsIntegrationInferenceResponse) SetMessages(v []LLMObsInferenceMessage) {
	o.Messages = v
}

// GetModelId returns the ModelId field value.
func (o *LLMObsIntegrationInferenceResponse) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceResponse) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value.
func (o *LLMObsIntegrationInferenceResponse) SetModelId(v string) {
	o.ModelId = v
}

// GetOpenaiMetadata returns the OpenaiMetadata field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceResponse) GetOpenaiMetadata() LLMObsOpenAIMetadata {
	if o == nil || o.OpenaiMetadata == nil {
		var ret LLMObsOpenAIMetadata
		return ret
	}
	return *o.OpenaiMetadata
}

// GetOpenaiMetadataOk returns a tuple with the OpenaiMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceResponse) GetOpenaiMetadataOk() (*LLMObsOpenAIMetadata, bool) {
	if o == nil || o.OpenaiMetadata == nil {
		return nil, false
	}
	return o.OpenaiMetadata, true
}

// HasOpenaiMetadata returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasOpenaiMetadata() bool {
	return o != nil && o.OpenaiMetadata != nil
}

// SetOpenaiMetadata gets a reference to the given LLMObsOpenAIMetadata and assigns it to the OpenaiMetadata field.
func (o *LLMObsIntegrationInferenceResponse) SetOpenaiMetadata(v LLMObsOpenAIMetadata) {
	o.OpenaiMetadata = &v
}

// GetPresencePenalty returns the PresencePenalty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceResponse) GetPresencePenalty() float64 {
	if o == nil || o.PresencePenalty.Get() == nil {
		var ret float64
		return ret
	}
	return *o.PresencePenalty.Get()
}

// GetPresencePenaltyOk returns a tuple with the PresencePenalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceResponse) GetPresencePenaltyOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PresencePenalty.Get(), o.PresencePenalty.IsSet()
}

// HasPresencePenalty returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasPresencePenalty() bool {
	return o != nil && o.PresencePenalty.IsSet()
}

// SetPresencePenalty gets a reference to the given datadog.NullableFloat64 and assigns it to the PresencePenalty field.
func (o *LLMObsIntegrationInferenceResponse) SetPresencePenalty(v float64) {
	o.PresencePenalty.Set(&v)
}

// SetPresencePenaltyNil sets the value for PresencePenalty to be an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) SetPresencePenaltyNil() {
	o.PresencePenalty.Set(nil)
}

// UnsetPresencePenalty ensures that no value is present for PresencePenalty, not even an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) UnsetPresencePenalty() {
	o.PresencePenalty.Unset()
}

// GetResponse returns the Response field value.
func (o *LLMObsIntegrationInferenceResponse) GetResponse() LLMObsInferenceRunResult {
	if o == nil {
		var ret LLMObsInferenceRunResult
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceResponse) GetResponseOk() (*LLMObsInferenceRunResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value.
func (o *LLMObsIntegrationInferenceResponse) SetResponse(v LLMObsInferenceRunResult) {
	o.Response = v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceResponse) GetTemperature() float64 {
	if o == nil || o.Temperature.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Temperature.Get()
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceResponse) GetTemperatureOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Temperature.Get(), o.Temperature.IsSet()
}

// HasTemperature returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasTemperature() bool {
	return o != nil && o.Temperature.IsSet()
}

// SetTemperature gets a reference to the given datadog.NullableFloat64 and assigns it to the Temperature field.
func (o *LLMObsIntegrationInferenceResponse) SetTemperature(v float64) {
	o.Temperature.Set(&v)
}

// SetTemperatureNil sets the value for Temperature to be an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) SetTemperatureNil() {
	o.Temperature.Set(nil)
}

// UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) UnsetTemperature() {
	o.Temperature.Unset()
}

// GetTools returns the Tools field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceResponse) GetTools() []LLMObsInferenceTool {
	if o == nil || o.Tools == nil {
		var ret []LLMObsInferenceTool
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceResponse) GetToolsOk() (*[]LLMObsInferenceTool, bool) {
	if o == nil || o.Tools == nil {
		return nil, false
	}
	return &o.Tools, true
}

// HasTools returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasTools() bool {
	return o != nil && o.Tools != nil
}

// SetTools gets a reference to the given []LLMObsInferenceTool and assigns it to the Tools field.
func (o *LLMObsIntegrationInferenceResponse) SetTools(v []LLMObsInferenceTool) {
	o.Tools = v
}

// GetTopK returns the TopK field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceResponse) GetTopK() int64 {
	if o == nil || o.TopK.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TopK.Get()
}

// GetTopKOk returns a tuple with the TopK field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceResponse) GetTopKOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopK.Get(), o.TopK.IsSet()
}

// HasTopK returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasTopK() bool {
	return o != nil && o.TopK.IsSet()
}

// SetTopK gets a reference to the given datadog.NullableInt64 and assigns it to the TopK field.
func (o *LLMObsIntegrationInferenceResponse) SetTopK(v int64) {
	o.TopK.Set(&v)
}

// SetTopKNil sets the value for TopK to be an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) SetTopKNil() {
	o.TopK.Set(nil)
}

// UnsetTopK ensures that no value is present for TopK, not even an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) UnsetTopK() {
	o.TopK.Unset()
}

// GetTopP returns the TopP field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsIntegrationInferenceResponse) GetTopP() float64 {
	if o == nil || o.TopP.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TopP.Get()
}

// GetTopPOk returns a tuple with the TopP field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsIntegrationInferenceResponse) GetTopPOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopP.Get(), o.TopP.IsSet()
}

// HasTopP returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasTopP() bool {
	return o != nil && o.TopP.IsSet()
}

// SetTopP gets a reference to the given datadog.NullableFloat64 and assigns it to the TopP field.
func (o *LLMObsIntegrationInferenceResponse) SetTopP(v float64) {
	o.TopP.Set(&v)
}

// SetTopPNil sets the value for TopP to be an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) SetTopPNil() {
	o.TopP.Set(nil)
}

// UnsetTopP ensures that no value is present for TopP, not even an explicit nil.
func (o *LLMObsIntegrationInferenceResponse) UnsetTopP() {
	o.TopP.Unset()
}

// GetVertexAiMetadata returns the VertexAiMetadata field value if set, zero value otherwise.
func (o *LLMObsIntegrationInferenceResponse) GetVertexAiMetadata() LLMObsVertexAIMetadata {
	if o == nil || o.VertexAiMetadata == nil {
		var ret LLMObsVertexAIMetadata
		return ret
	}
	return *o.VertexAiMetadata
}

// GetVertexAiMetadataOk returns a tuple with the VertexAiMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationInferenceResponse) GetVertexAiMetadataOk() (*LLMObsVertexAIMetadata, bool) {
	if o == nil || o.VertexAiMetadata == nil {
		return nil, false
	}
	return o.VertexAiMetadata, true
}

// HasVertexAiMetadata returns a boolean if a field has been set.
func (o *LLMObsIntegrationInferenceResponse) HasVertexAiMetadata() bool {
	return o != nil && o.VertexAiMetadata != nil
}

// SetVertexAiMetadata gets a reference to the given LLMObsVertexAIMetadata and assigns it to the VertexAiMetadata field.
func (o *LLMObsIntegrationInferenceResponse) SetVertexAiMetadata(v LLMObsVertexAIMetadata) {
	o.VertexAiMetadata = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsIntegrationInferenceResponse) MarshalJSON() ([]byte, error) {
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
	if o.ErrorResponse != nil {
		toSerialize["error_response"] = o.ErrorResponse
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
	toSerialize["response"] = o.Response
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
func (o *LLMObsIntegrationInferenceResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnthropicMetadata   *LLMObsAnthropicMetadata      `json:"anthropic_metadata,omitempty"`
		AzureOpenaiMetadata *LLMObsAzureOpenAIMetadata    `json:"azure_openai_metadata,omitempty"`
		BedrockMetadata     *LLMObsBedrockMetadata        `json:"bedrock_metadata,omitempty"`
		ErrorResponse       *LLMObsInferenceErrorResponse `json:"error_response,omitempty"`
		FrequencyPenalty    datadog.NullableFloat64       `json:"frequency_penalty,omitempty"`
		JsonSchema          datadog.NullableString        `json:"json_schema,omitempty"`
		MaxCompletionTokens datadog.NullableInt64         `json:"max_completion_tokens,omitempty"`
		MaxTokens           datadog.NullableInt64         `json:"max_tokens,omitempty"`
		Messages            *[]LLMObsInferenceMessage     `json:"messages"`
		ModelId             *string                       `json:"model_id"`
		OpenaiMetadata      *LLMObsOpenAIMetadata         `json:"openai_metadata,omitempty"`
		PresencePenalty     datadog.NullableFloat64       `json:"presence_penalty,omitempty"`
		Response            *LLMObsInferenceRunResult     `json:"response"`
		Temperature         datadog.NullableFloat64       `json:"temperature,omitempty"`
		Tools               []LLMObsInferenceTool         `json:"tools,omitempty"`
		TopK                datadog.NullableInt64         `json:"top_k,omitempty"`
		TopP                datadog.NullableFloat64       `json:"top_p,omitempty"`
		VertexAiMetadata    *LLMObsVertexAIMetadata       `json:"vertex_ai_metadata,omitempty"`
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
	if all.Response == nil {
		return fmt.Errorf("required field response missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"anthropic_metadata", "azure_openai_metadata", "bedrock_metadata", "error_response", "frequency_penalty", "json_schema", "max_completion_tokens", "max_tokens", "messages", "model_id", "openai_metadata", "presence_penalty", "response", "temperature", "tools", "top_k", "top_p", "vertex_ai_metadata"})
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
	if all.ErrorResponse != nil && all.ErrorResponse.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ErrorResponse = all.ErrorResponse
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
	if all.Response.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Response = *all.Response
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
