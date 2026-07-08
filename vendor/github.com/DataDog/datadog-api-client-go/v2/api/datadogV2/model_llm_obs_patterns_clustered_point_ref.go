// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsClusteredPointRef A clustered point attached inline to a topic. The metric fields are populated
// only when the request includes `include_metrics=true`.
type LLMObsPatternsClusteredPointRef struct {
	// Duration of the source span in nanoseconds. Included only when metrics are requested.
	Duration *float64 `json:"duration,omitempty"`
	// Estimated total cost of the source span. Included only when metrics are requested.
	EstimatedTotalCost *float64 `json:"estimated_total_cost,omitempty"`
	// Evaluation results for the source span keyed by evaluation name. Included
	// only when metrics are requested.
	Evaluation map[string]interface{} `json:"evaluation,omitempty"`
	// Number of input tokens of the source span. Included only when metrics are requested.
	InputTokens *float64 `json:"input_tokens,omitempty"`
	// Number of output tokens of the source span. Included only when metrics are requested.
	OutputTokens *float64 `json:"output_tokens,omitempty"`
	// Identifier of the source span.
	SpanId string `json:"span_id"`
	// Status of the source span. Included only when metrics are requested.
	Status *string `json:"status,omitempty"`
	// Total number of tokens of the source span. Included only when metrics are requested.
	TotalTokens *float64 `json:"total_tokens,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsClusteredPointRef instantiates a new LLMObsPatternsClusteredPointRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsClusteredPointRef(spanId string) *LLMObsPatternsClusteredPointRef {
	this := LLMObsPatternsClusteredPointRef{}
	this.SpanId = spanId
	return &this
}

// NewLLMObsPatternsClusteredPointRefWithDefaults instantiates a new LLMObsPatternsClusteredPointRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsClusteredPointRefWithDefaults() *LLMObsPatternsClusteredPointRef {
	this := LLMObsPatternsClusteredPointRef{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *LLMObsPatternsClusteredPointRef) GetDuration() float64 {
	if o == nil || o.Duration == nil {
		var ret float64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPointRef) GetDurationOk() (*float64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *LLMObsPatternsClusteredPointRef) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given float64 and assigns it to the Duration field.
func (o *LLMObsPatternsClusteredPointRef) SetDuration(v float64) {
	o.Duration = &v
}

// GetEstimatedTotalCost returns the EstimatedTotalCost field value if set, zero value otherwise.
func (o *LLMObsPatternsClusteredPointRef) GetEstimatedTotalCost() float64 {
	if o == nil || o.EstimatedTotalCost == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedTotalCost
}

// GetEstimatedTotalCostOk returns a tuple with the EstimatedTotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPointRef) GetEstimatedTotalCostOk() (*float64, bool) {
	if o == nil || o.EstimatedTotalCost == nil {
		return nil, false
	}
	return o.EstimatedTotalCost, true
}

// HasEstimatedTotalCost returns a boolean if a field has been set.
func (o *LLMObsPatternsClusteredPointRef) HasEstimatedTotalCost() bool {
	return o != nil && o.EstimatedTotalCost != nil
}

// SetEstimatedTotalCost gets a reference to the given float64 and assigns it to the EstimatedTotalCost field.
func (o *LLMObsPatternsClusteredPointRef) SetEstimatedTotalCost(v float64) {
	o.EstimatedTotalCost = &v
}

// GetEvaluation returns the Evaluation field value if set, zero value otherwise.
func (o *LLMObsPatternsClusteredPointRef) GetEvaluation() map[string]interface{} {
	if o == nil || o.Evaluation == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Evaluation
}

// GetEvaluationOk returns a tuple with the Evaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPointRef) GetEvaluationOk() (*map[string]interface{}, bool) {
	if o == nil || o.Evaluation == nil {
		return nil, false
	}
	return &o.Evaluation, true
}

// HasEvaluation returns a boolean if a field has been set.
func (o *LLMObsPatternsClusteredPointRef) HasEvaluation() bool {
	return o != nil && o.Evaluation != nil
}

// SetEvaluation gets a reference to the given map[string]interface{} and assigns it to the Evaluation field.
func (o *LLMObsPatternsClusteredPointRef) SetEvaluation(v map[string]interface{}) {
	o.Evaluation = v
}

// GetInputTokens returns the InputTokens field value if set, zero value otherwise.
func (o *LLMObsPatternsClusteredPointRef) GetInputTokens() float64 {
	if o == nil || o.InputTokens == nil {
		var ret float64
		return ret
	}
	return *o.InputTokens
}

// GetInputTokensOk returns a tuple with the InputTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPointRef) GetInputTokensOk() (*float64, bool) {
	if o == nil || o.InputTokens == nil {
		return nil, false
	}
	return o.InputTokens, true
}

// HasInputTokens returns a boolean if a field has been set.
func (o *LLMObsPatternsClusteredPointRef) HasInputTokens() bool {
	return o != nil && o.InputTokens != nil
}

// SetInputTokens gets a reference to the given float64 and assigns it to the InputTokens field.
func (o *LLMObsPatternsClusteredPointRef) SetInputTokens(v float64) {
	o.InputTokens = &v
}

// GetOutputTokens returns the OutputTokens field value if set, zero value otherwise.
func (o *LLMObsPatternsClusteredPointRef) GetOutputTokens() float64 {
	if o == nil || o.OutputTokens == nil {
		var ret float64
		return ret
	}
	return *o.OutputTokens
}

// GetOutputTokensOk returns a tuple with the OutputTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPointRef) GetOutputTokensOk() (*float64, bool) {
	if o == nil || o.OutputTokens == nil {
		return nil, false
	}
	return o.OutputTokens, true
}

// HasOutputTokens returns a boolean if a field has been set.
func (o *LLMObsPatternsClusteredPointRef) HasOutputTokens() bool {
	return o != nil && o.OutputTokens != nil
}

// SetOutputTokens gets a reference to the given float64 and assigns it to the OutputTokens field.
func (o *LLMObsPatternsClusteredPointRef) SetOutputTokens(v float64) {
	o.OutputTokens = &v
}

// GetSpanId returns the SpanId field value.
func (o *LLMObsPatternsClusteredPointRef) GetSpanId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPointRef) GetSpanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanId, true
}

// SetSpanId sets field value.
func (o *LLMObsPatternsClusteredPointRef) SetSpanId(v string) {
	o.SpanId = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LLMObsPatternsClusteredPointRef) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPointRef) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LLMObsPatternsClusteredPointRef) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LLMObsPatternsClusteredPointRef) SetStatus(v string) {
	o.Status = &v
}

// GetTotalTokens returns the TotalTokens field value if set, zero value otherwise.
func (o *LLMObsPatternsClusteredPointRef) GetTotalTokens() float64 {
	if o == nil || o.TotalTokens == nil {
		var ret float64
		return ret
	}
	return *o.TotalTokens
}

// GetTotalTokensOk returns a tuple with the TotalTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPointRef) GetTotalTokensOk() (*float64, bool) {
	if o == nil || o.TotalTokens == nil {
		return nil, false
	}
	return o.TotalTokens, true
}

// HasTotalTokens returns a boolean if a field has been set.
func (o *LLMObsPatternsClusteredPointRef) HasTotalTokens() bool {
	return o != nil && o.TotalTokens != nil
}

// SetTotalTokens gets a reference to the given float64 and assigns it to the TotalTokens field.
func (o *LLMObsPatternsClusteredPointRef) SetTotalTokens(v float64) {
	o.TotalTokens = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsClusteredPointRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.EstimatedTotalCost != nil {
		toSerialize["estimated_total_cost"] = o.EstimatedTotalCost
	}
	if o.Evaluation != nil {
		toSerialize["evaluation"] = o.Evaluation
	}
	if o.InputTokens != nil {
		toSerialize["input_tokens"] = o.InputTokens
	}
	if o.OutputTokens != nil {
		toSerialize["output_tokens"] = o.OutputTokens
	}
	toSerialize["span_id"] = o.SpanId
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TotalTokens != nil {
		toSerialize["total_tokens"] = o.TotalTokens
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsClusteredPointRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration           *float64               `json:"duration,omitempty"`
		EstimatedTotalCost *float64               `json:"estimated_total_cost,omitempty"`
		Evaluation         map[string]interface{} `json:"evaluation,omitempty"`
		InputTokens        *float64               `json:"input_tokens,omitempty"`
		OutputTokens       *float64               `json:"output_tokens,omitempty"`
		SpanId             *string                `json:"span_id"`
		Status             *string                `json:"status,omitempty"`
		TotalTokens        *float64               `json:"total_tokens,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SpanId == nil {
		return fmt.Errorf("required field span_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration", "estimated_total_cost", "evaluation", "input_tokens", "output_tokens", "span_id", "status", "total_tokens"})
	} else {
		return err
	}
	o.Duration = all.Duration
	o.EstimatedTotalCost = all.EstimatedTotalCost
	o.Evaluation = all.Evaluation
	o.InputTokens = all.InputTokens
	o.OutputTokens = all.OutputTokens
	o.SpanId = *all.SpanId
	o.Status = all.Status
	o.TotalTokens = all.TotalTokens

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
