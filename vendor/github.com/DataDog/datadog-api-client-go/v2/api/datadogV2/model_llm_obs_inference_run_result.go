// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsInferenceRunResult The output of a completed LLM inference call.
type LLMObsInferenceRunResult struct {
	// An optional assessment of the inference output quality.
	Assessment datadog.NullableString `json:"assessment"`
	// The text content of the model response.
	Content string `json:"content"`
	// The reason the model stopped generating tokens.
	FinishReason string `json:"finish_reason"`
	// List of generated code snippets for the inference configuration.
	InferenceCodes []LLMObsInferenceCode `json:"inference_codes"`
	// Number of input tokens consumed.
	InputTokens int64 `json:"input_tokens"`
	// The model's internal reasoning or thinking output, if available.
	InternalReasoning *LLMObsInternalReasoning `json:"internal_reasoning,omitempty"`
	// Request latency in milliseconds.
	Latency int64 `json:"latency"`
	// Number of output tokens generated.
	OutputTokens int64 `json:"output_tokens"`
	// List of tools available to the model.
	Tools []LLMObsInferenceTool `json:"tools"`
	// Total tokens used (input plus output).
	TotalTokens int64 `json:"total_tokens"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsInferenceRunResult instantiates a new LLMObsInferenceRunResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsInferenceRunResult(assessment datadog.NullableString, content string, finishReason string, inferenceCodes []LLMObsInferenceCode, inputTokens int64, latency int64, outputTokens int64, tools []LLMObsInferenceTool, totalTokens int64) *LLMObsInferenceRunResult {
	this := LLMObsInferenceRunResult{}
	this.Assessment = assessment
	this.Content = content
	this.FinishReason = finishReason
	this.InferenceCodes = inferenceCodes
	this.InputTokens = inputTokens
	this.Latency = latency
	this.OutputTokens = outputTokens
	this.Tools = tools
	this.TotalTokens = totalTokens
	return &this
}

// NewLLMObsInferenceRunResultWithDefaults instantiates a new LLMObsInferenceRunResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsInferenceRunResultWithDefaults() *LLMObsInferenceRunResult {
	this := LLMObsInferenceRunResult{}
	return &this
}

// GetAssessment returns the Assessment field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *LLMObsInferenceRunResult) GetAssessment() string {
	if o == nil || o.Assessment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Assessment.Get()
}

// GetAssessmentOk returns a tuple with the Assessment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsInferenceRunResult) GetAssessmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assessment.Get(), o.Assessment.IsSet()
}

// SetAssessment sets field value.
func (o *LLMObsInferenceRunResult) SetAssessment(v string) {
	o.Assessment.Set(&v)
}

// GetContent returns the Content field value.
func (o *LLMObsInferenceRunResult) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceRunResult) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *LLMObsInferenceRunResult) SetContent(v string) {
	o.Content = v
}

// GetFinishReason returns the FinishReason field value.
func (o *LLMObsInferenceRunResult) GetFinishReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FinishReason
}

// GetFinishReasonOk returns a tuple with the FinishReason field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceRunResult) GetFinishReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishReason, true
}

// SetFinishReason sets field value.
func (o *LLMObsInferenceRunResult) SetFinishReason(v string) {
	o.FinishReason = v
}

// GetInferenceCodes returns the InferenceCodes field value.
func (o *LLMObsInferenceRunResult) GetInferenceCodes() []LLMObsInferenceCode {
	if o == nil {
		var ret []LLMObsInferenceCode
		return ret
	}
	return o.InferenceCodes
}

// GetInferenceCodesOk returns a tuple with the InferenceCodes field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceRunResult) GetInferenceCodesOk() (*[]LLMObsInferenceCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InferenceCodes, true
}

// SetInferenceCodes sets field value.
func (o *LLMObsInferenceRunResult) SetInferenceCodes(v []LLMObsInferenceCode) {
	o.InferenceCodes = v
}

// GetInputTokens returns the InputTokens field value.
func (o *LLMObsInferenceRunResult) GetInputTokens() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.InputTokens
}

// GetInputTokensOk returns a tuple with the InputTokens field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceRunResult) GetInputTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputTokens, true
}

// SetInputTokens sets field value.
func (o *LLMObsInferenceRunResult) SetInputTokens(v int64) {
	o.InputTokens = v
}

// GetInternalReasoning returns the InternalReasoning field value if set, zero value otherwise.
func (o *LLMObsInferenceRunResult) GetInternalReasoning() LLMObsInternalReasoning {
	if o == nil || o.InternalReasoning == nil {
		var ret LLMObsInternalReasoning
		return ret
	}
	return *o.InternalReasoning
}

// GetInternalReasoningOk returns a tuple with the InternalReasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceRunResult) GetInternalReasoningOk() (*LLMObsInternalReasoning, bool) {
	if o == nil || o.InternalReasoning == nil {
		return nil, false
	}
	return o.InternalReasoning, true
}

// HasInternalReasoning returns a boolean if a field has been set.
func (o *LLMObsInferenceRunResult) HasInternalReasoning() bool {
	return o != nil && o.InternalReasoning != nil
}

// SetInternalReasoning gets a reference to the given LLMObsInternalReasoning and assigns it to the InternalReasoning field.
func (o *LLMObsInferenceRunResult) SetInternalReasoning(v LLMObsInternalReasoning) {
	o.InternalReasoning = &v
}

// GetLatency returns the Latency field value.
func (o *LLMObsInferenceRunResult) GetLatency() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceRunResult) GetLatencyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latency, true
}

// SetLatency sets field value.
func (o *LLMObsInferenceRunResult) SetLatency(v int64) {
	o.Latency = v
}

// GetOutputTokens returns the OutputTokens field value.
func (o *LLMObsInferenceRunResult) GetOutputTokens() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OutputTokens
}

// GetOutputTokensOk returns a tuple with the OutputTokens field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceRunResult) GetOutputTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputTokens, true
}

// SetOutputTokens sets field value.
func (o *LLMObsInferenceRunResult) SetOutputTokens(v int64) {
	o.OutputTokens = v
}

// GetTools returns the Tools field value.
func (o *LLMObsInferenceRunResult) GetTools() []LLMObsInferenceTool {
	if o == nil {
		var ret []LLMObsInferenceTool
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceRunResult) GetToolsOk() (*[]LLMObsInferenceTool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tools, true
}

// SetTools sets field value.
func (o *LLMObsInferenceRunResult) SetTools(v []LLMObsInferenceTool) {
	o.Tools = v
}

// GetTotalTokens returns the TotalTokens field value.
func (o *LLMObsInferenceRunResult) GetTotalTokens() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalTokens
}

// GetTotalTokensOk returns a tuple with the TotalTokens field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceRunResult) GetTotalTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTokens, true
}

// SetTotalTokens sets field value.
func (o *LLMObsInferenceRunResult) SetTotalTokens(v int64) {
	o.TotalTokens = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsInferenceRunResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["assessment"] = o.Assessment.Get()
	toSerialize["content"] = o.Content
	toSerialize["finish_reason"] = o.FinishReason
	toSerialize["inference_codes"] = o.InferenceCodes
	toSerialize["input_tokens"] = o.InputTokens
	if o.InternalReasoning != nil {
		toSerialize["internal_reasoning"] = o.InternalReasoning
	}
	toSerialize["latency"] = o.Latency
	toSerialize["output_tokens"] = o.OutputTokens
	toSerialize["tools"] = o.Tools
	toSerialize["total_tokens"] = o.TotalTokens

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsInferenceRunResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assessment        datadog.NullableString   `json:"assessment"`
		Content           *string                  `json:"content"`
		FinishReason      *string                  `json:"finish_reason"`
		InferenceCodes    *[]LLMObsInferenceCode   `json:"inference_codes"`
		InputTokens       *int64                   `json:"input_tokens"`
		InternalReasoning *LLMObsInternalReasoning `json:"internal_reasoning,omitempty"`
		Latency           *int64                   `json:"latency"`
		OutputTokens      *int64                   `json:"output_tokens"`
		Tools             *[]LLMObsInferenceTool   `json:"tools"`
		TotalTokens       *int64                   `json:"total_tokens"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Assessment.IsSet() {
		return fmt.Errorf("required field assessment missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	if all.FinishReason == nil {
		return fmt.Errorf("required field finish_reason missing")
	}
	if all.InferenceCodes == nil {
		return fmt.Errorf("required field inference_codes missing")
	}
	if all.InputTokens == nil {
		return fmt.Errorf("required field input_tokens missing")
	}
	if all.Latency == nil {
		return fmt.Errorf("required field latency missing")
	}
	if all.OutputTokens == nil {
		return fmt.Errorf("required field output_tokens missing")
	}
	if all.Tools == nil {
		return fmt.Errorf("required field tools missing")
	}
	if all.TotalTokens == nil {
		return fmt.Errorf("required field total_tokens missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assessment", "content", "finish_reason", "inference_codes", "input_tokens", "internal_reasoning", "latency", "output_tokens", "tools", "total_tokens"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Assessment = all.Assessment
	o.Content = *all.Content
	o.FinishReason = *all.FinishReason
	o.InferenceCodes = *all.InferenceCodes
	o.InputTokens = *all.InputTokens
	if all.InternalReasoning != nil && all.InternalReasoning.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.InternalReasoning = all.InternalReasoning
	o.Latency = *all.Latency
	o.OutputTokens = *all.OutputTokens
	o.Tools = *all.Tools
	o.TotalTokens = *all.TotalTokens

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
