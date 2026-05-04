// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigInferenceParams LLM inference parameters for a custom evaluator.
type LLMObsCustomEvalConfigInferenceParams struct {
	// Frequency penalty to reduce repetition.
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`
	// Maximum number of tokens to generate.
	MaxTokens *int64 `json:"max_tokens,omitempty"`
	// Presence penalty to reduce repetition.
	PresencePenalty *float64 `json:"presence_penalty,omitempty"`
	// Sampling temperature for the LLM.
	Temperature *float64 `json:"temperature,omitempty"`
	// Top-k sampling parameter.
	TopK *int64 `json:"top_k,omitempty"`
	// Top-p (nucleus) sampling parameter.
	TopP *float64 `json:"top_p,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCustomEvalConfigInferenceParams instantiates a new LLMObsCustomEvalConfigInferenceParams object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCustomEvalConfigInferenceParams() *LLMObsCustomEvalConfigInferenceParams {
	this := LLMObsCustomEvalConfigInferenceParams{}
	return &this
}

// NewLLMObsCustomEvalConfigInferenceParamsWithDefaults instantiates a new LLMObsCustomEvalConfigInferenceParams object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCustomEvalConfigInferenceParamsWithDefaults() *LLMObsCustomEvalConfigInferenceParams {
	this := LLMObsCustomEvalConfigInferenceParams{}
	return &this
}

// GetFrequencyPenalty returns the FrequencyPenalty field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigInferenceParams) GetFrequencyPenalty() float64 {
	if o == nil || o.FrequencyPenalty == nil {
		var ret float64
		return ret
	}
	return *o.FrequencyPenalty
}

// GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigInferenceParams) GetFrequencyPenaltyOk() (*float64, bool) {
	if o == nil || o.FrequencyPenalty == nil {
		return nil, false
	}
	return o.FrequencyPenalty, true
}

// HasFrequencyPenalty returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigInferenceParams) HasFrequencyPenalty() bool {
	return o != nil && o.FrequencyPenalty != nil
}

// SetFrequencyPenalty gets a reference to the given float64 and assigns it to the FrequencyPenalty field.
func (o *LLMObsCustomEvalConfigInferenceParams) SetFrequencyPenalty(v float64) {
	o.FrequencyPenalty = &v
}

// GetMaxTokens returns the MaxTokens field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigInferenceParams) GetMaxTokens() int64 {
	if o == nil || o.MaxTokens == nil {
		var ret int64
		return ret
	}
	return *o.MaxTokens
}

// GetMaxTokensOk returns a tuple with the MaxTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigInferenceParams) GetMaxTokensOk() (*int64, bool) {
	if o == nil || o.MaxTokens == nil {
		return nil, false
	}
	return o.MaxTokens, true
}

// HasMaxTokens returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigInferenceParams) HasMaxTokens() bool {
	return o != nil && o.MaxTokens != nil
}

// SetMaxTokens gets a reference to the given int64 and assigns it to the MaxTokens field.
func (o *LLMObsCustomEvalConfigInferenceParams) SetMaxTokens(v int64) {
	o.MaxTokens = &v
}

// GetPresencePenalty returns the PresencePenalty field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigInferenceParams) GetPresencePenalty() float64 {
	if o == nil || o.PresencePenalty == nil {
		var ret float64
		return ret
	}
	return *o.PresencePenalty
}

// GetPresencePenaltyOk returns a tuple with the PresencePenalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigInferenceParams) GetPresencePenaltyOk() (*float64, bool) {
	if o == nil || o.PresencePenalty == nil {
		return nil, false
	}
	return o.PresencePenalty, true
}

// HasPresencePenalty returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigInferenceParams) HasPresencePenalty() bool {
	return o != nil && o.PresencePenalty != nil
}

// SetPresencePenalty gets a reference to the given float64 and assigns it to the PresencePenalty field.
func (o *LLMObsCustomEvalConfigInferenceParams) SetPresencePenalty(v float64) {
	o.PresencePenalty = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigInferenceParams) GetTemperature() float64 {
	if o == nil || o.Temperature == nil {
		var ret float64
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigInferenceParams) GetTemperatureOk() (*float64, bool) {
	if o == nil || o.Temperature == nil {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigInferenceParams) HasTemperature() bool {
	return o != nil && o.Temperature != nil
}

// SetTemperature gets a reference to the given float64 and assigns it to the Temperature field.
func (o *LLMObsCustomEvalConfigInferenceParams) SetTemperature(v float64) {
	o.Temperature = &v
}

// GetTopK returns the TopK field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigInferenceParams) GetTopK() int64 {
	if o == nil || o.TopK == nil {
		var ret int64
		return ret
	}
	return *o.TopK
}

// GetTopKOk returns a tuple with the TopK field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigInferenceParams) GetTopKOk() (*int64, bool) {
	if o == nil || o.TopK == nil {
		return nil, false
	}
	return o.TopK, true
}

// HasTopK returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigInferenceParams) HasTopK() bool {
	return o != nil && o.TopK != nil
}

// SetTopK gets a reference to the given int64 and assigns it to the TopK field.
func (o *LLMObsCustomEvalConfigInferenceParams) SetTopK(v int64) {
	o.TopK = &v
}

// GetTopP returns the TopP field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigInferenceParams) GetTopP() float64 {
	if o == nil || o.TopP == nil {
		var ret float64
		return ret
	}
	return *o.TopP
}

// GetTopPOk returns a tuple with the TopP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigInferenceParams) GetTopPOk() (*float64, bool) {
	if o == nil || o.TopP == nil {
		return nil, false
	}
	return o.TopP, true
}

// HasTopP returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigInferenceParams) HasTopP() bool {
	return o != nil && o.TopP != nil
}

// SetTopP gets a reference to the given float64 and assigns it to the TopP field.
func (o *LLMObsCustomEvalConfigInferenceParams) SetTopP(v float64) {
	o.TopP = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCustomEvalConfigInferenceParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FrequencyPenalty != nil {
		toSerialize["frequency_penalty"] = o.FrequencyPenalty
	}
	if o.MaxTokens != nil {
		toSerialize["max_tokens"] = o.MaxTokens
	}
	if o.PresencePenalty != nil {
		toSerialize["presence_penalty"] = o.PresencePenalty
	}
	if o.Temperature != nil {
		toSerialize["temperature"] = o.Temperature
	}
	if o.TopK != nil {
		toSerialize["top_k"] = o.TopK
	}
	if o.TopP != nil {
		toSerialize["top_p"] = o.TopP
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCustomEvalConfigInferenceParams) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`
		MaxTokens        *int64   `json:"max_tokens,omitempty"`
		PresencePenalty  *float64 `json:"presence_penalty,omitempty"`
		Temperature      *float64 `json:"temperature,omitempty"`
		TopK             *int64   `json:"top_k,omitempty"`
		TopP             *float64 `json:"top_p,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"frequency_penalty", "max_tokens", "presence_penalty", "temperature", "top_k", "top_p"})
	} else {
		return err
	}
	o.FrequencyPenalty = all.FrequencyPenalty
	o.MaxTokens = all.MaxTokens
	o.PresencePenalty = all.PresencePenalty
	o.Temperature = all.Temperature
	o.TopK = all.TopK
	o.TopP = all.TopP

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
