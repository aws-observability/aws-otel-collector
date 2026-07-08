// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSpanEvaluationMetric An evaluation metric associated with an LLM Observability span.
type LLMObsSpanEvaluationMetric struct {
	// Assessment result (e.g., pass or fail).
	Assessment *string `json:"assessment,omitempty"`
	// Type of the evaluation metric (e.g., score, categorical, boolean).
	EvalMetricType *string `json:"eval_metric_type,omitempty"`
	// Human-readable reasoning for the evaluation result.
	Reasoning *string `json:"reasoning,omitempty"`
	// Status of the evaluation execution.
	Status *string `json:"status,omitempty"`
	// Tags associated with the evaluation metric.
	Tags []string `json:"tags,omitempty"`
	// Value of the evaluation result.
	Value interface{} `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsSpanEvaluationMetric instantiates a new LLMObsSpanEvaluationMetric object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsSpanEvaluationMetric() *LLMObsSpanEvaluationMetric {
	this := LLMObsSpanEvaluationMetric{}
	return &this
}

// NewLLMObsSpanEvaluationMetricWithDefaults instantiates a new LLMObsSpanEvaluationMetric object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsSpanEvaluationMetricWithDefaults() *LLMObsSpanEvaluationMetric {
	this := LLMObsSpanEvaluationMetric{}
	return &this
}

// GetAssessment returns the Assessment field value if set, zero value otherwise.
func (o *LLMObsSpanEvaluationMetric) GetAssessment() string {
	if o == nil || o.Assessment == nil {
		var ret string
		return ret
	}
	return *o.Assessment
}

// GetAssessmentOk returns a tuple with the Assessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanEvaluationMetric) GetAssessmentOk() (*string, bool) {
	if o == nil || o.Assessment == nil {
		return nil, false
	}
	return o.Assessment, true
}

// HasAssessment returns a boolean if a field has been set.
func (o *LLMObsSpanEvaluationMetric) HasAssessment() bool {
	return o != nil && o.Assessment != nil
}

// SetAssessment gets a reference to the given string and assigns it to the Assessment field.
func (o *LLMObsSpanEvaluationMetric) SetAssessment(v string) {
	o.Assessment = &v
}

// GetEvalMetricType returns the EvalMetricType field value if set, zero value otherwise.
func (o *LLMObsSpanEvaluationMetric) GetEvalMetricType() string {
	if o == nil || o.EvalMetricType == nil {
		var ret string
		return ret
	}
	return *o.EvalMetricType
}

// GetEvalMetricTypeOk returns a tuple with the EvalMetricType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanEvaluationMetric) GetEvalMetricTypeOk() (*string, bool) {
	if o == nil || o.EvalMetricType == nil {
		return nil, false
	}
	return o.EvalMetricType, true
}

// HasEvalMetricType returns a boolean if a field has been set.
func (o *LLMObsSpanEvaluationMetric) HasEvalMetricType() bool {
	return o != nil && o.EvalMetricType != nil
}

// SetEvalMetricType gets a reference to the given string and assigns it to the EvalMetricType field.
func (o *LLMObsSpanEvaluationMetric) SetEvalMetricType(v string) {
	o.EvalMetricType = &v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *LLMObsSpanEvaluationMetric) GetReasoning() string {
	if o == nil || o.Reasoning == nil {
		var ret string
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanEvaluationMetric) GetReasoningOk() (*string, bool) {
	if o == nil || o.Reasoning == nil {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *LLMObsSpanEvaluationMetric) HasReasoning() bool {
	return o != nil && o.Reasoning != nil
}

// SetReasoning gets a reference to the given string and assigns it to the Reasoning field.
func (o *LLMObsSpanEvaluationMetric) SetReasoning(v string) {
	o.Reasoning = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LLMObsSpanEvaluationMetric) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanEvaluationMetric) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LLMObsSpanEvaluationMetric) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LLMObsSpanEvaluationMetric) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LLMObsSpanEvaluationMetric) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanEvaluationMetric) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LLMObsSpanEvaluationMetric) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LLMObsSpanEvaluationMetric) SetTags(v []string) {
	o.Tags = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LLMObsSpanEvaluationMetric) GetValue() interface{} {
	if o == nil || o.Value == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanEvaluationMetric) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LLMObsSpanEvaluationMetric) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *LLMObsSpanEvaluationMetric) SetValue(v interface{}) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsSpanEvaluationMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assessment != nil {
		toSerialize["assessment"] = o.Assessment
	}
	if o.EvalMetricType != nil {
		toSerialize["eval_metric_type"] = o.EvalMetricType
	}
	if o.Reasoning != nil {
		toSerialize["reasoning"] = o.Reasoning
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsSpanEvaluationMetric) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assessment     *string     `json:"assessment,omitempty"`
		EvalMetricType *string     `json:"eval_metric_type,omitempty"`
		Reasoning      *string     `json:"reasoning,omitempty"`
		Status         *string     `json:"status,omitempty"`
		Tags           []string    `json:"tags,omitempty"`
		Value          interface{} `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assessment", "eval_metric_type", "reasoning", "status", "tags", "value"})
	} else {
		return err
	}
	o.Assessment = all.Assessment
	o.EvalMetricType = all.EvalMetricType
	o.Reasoning = all.Reasoning
	o.Status = all.Status
	o.Tags = all.Tags
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
