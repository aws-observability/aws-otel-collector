// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentMetric A metric associated with an LLM Observability experiment span.
type LLMObsExperimentMetric struct {
	// Assessment result for an LLM Observability experiment metric.
	Assessment *LLMObsMetricAssessment `json:"assessment,omitempty"`
	// Boolean value. Used when `metric_type` is `boolean`.
	BooleanValue *bool `json:"boolean_value,omitempty"`
	// Categorical value. Used when `metric_type` is `categorical`.
	CategoricalValue *string `json:"categorical_value,omitempty"`
	// Error details for an experiment metric evaluation.
	Error *LLMObsExperimentMetricError `json:"error,omitempty"`
	// JSON value. Used when `metric_type` is `json`.
	JsonValue map[string]interface{} `json:"json_value,omitempty"`
	// Label or name for the metric.
	Label string `json:"label"`
	// Arbitrary metadata associated with the metric.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Type of metric recorded for an LLM Observability experiment.
	MetricType LLMObsMetricScoreType `json:"metric_type"`
	// Human-readable reasoning for the metric value.
	Reasoning *string `json:"reasoning,omitempty"`
	// Numeric score value. Used when `metric_type` is `score`.
	ScoreValue *float64 `json:"score_value,omitempty"`
	// The ID of the span this metric measures.
	SpanId string `json:"span_id"`
	// List of tags associated with the metric.
	Tags []string `json:"tags,omitempty"`
	// Timestamp when the metric was recorded, in milliseconds since Unix epoch.
	TimestampMs int64 `json:"timestamp_ms"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentMetric instantiates a new LLMObsExperimentMetric object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentMetric(label string, metricType LLMObsMetricScoreType, spanId string, timestampMs int64) *LLMObsExperimentMetric {
	this := LLMObsExperimentMetric{}
	this.Label = label
	this.MetricType = metricType
	this.SpanId = spanId
	this.TimestampMs = timestampMs
	return &this
}

// NewLLMObsExperimentMetricWithDefaults instantiates a new LLMObsExperimentMetric object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentMetricWithDefaults() *LLMObsExperimentMetric {
	this := LLMObsExperimentMetric{}
	return &this
}

// GetAssessment returns the Assessment field value if set, zero value otherwise.
func (o *LLMObsExperimentMetric) GetAssessment() LLMObsMetricAssessment {
	if o == nil || o.Assessment == nil {
		var ret LLMObsMetricAssessment
		return ret
	}
	return *o.Assessment
}

// GetAssessmentOk returns a tuple with the Assessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetAssessmentOk() (*LLMObsMetricAssessment, bool) {
	if o == nil || o.Assessment == nil {
		return nil, false
	}
	return o.Assessment, true
}

// HasAssessment returns a boolean if a field has been set.
func (o *LLMObsExperimentMetric) HasAssessment() bool {
	return o != nil && o.Assessment != nil
}

// SetAssessment gets a reference to the given LLMObsMetricAssessment and assigns it to the Assessment field.
func (o *LLMObsExperimentMetric) SetAssessment(v LLMObsMetricAssessment) {
	o.Assessment = &v
}

// GetBooleanValue returns the BooleanValue field value if set, zero value otherwise.
func (o *LLMObsExperimentMetric) GetBooleanValue() bool {
	if o == nil || o.BooleanValue == nil {
		var ret bool
		return ret
	}
	return *o.BooleanValue
}

// GetBooleanValueOk returns a tuple with the BooleanValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetBooleanValueOk() (*bool, bool) {
	if o == nil || o.BooleanValue == nil {
		return nil, false
	}
	return o.BooleanValue, true
}

// HasBooleanValue returns a boolean if a field has been set.
func (o *LLMObsExperimentMetric) HasBooleanValue() bool {
	return o != nil && o.BooleanValue != nil
}

// SetBooleanValue gets a reference to the given bool and assigns it to the BooleanValue field.
func (o *LLMObsExperimentMetric) SetBooleanValue(v bool) {
	o.BooleanValue = &v
}

// GetCategoricalValue returns the CategoricalValue field value if set, zero value otherwise.
func (o *LLMObsExperimentMetric) GetCategoricalValue() string {
	if o == nil || o.CategoricalValue == nil {
		var ret string
		return ret
	}
	return *o.CategoricalValue
}

// GetCategoricalValueOk returns a tuple with the CategoricalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetCategoricalValueOk() (*string, bool) {
	if o == nil || o.CategoricalValue == nil {
		return nil, false
	}
	return o.CategoricalValue, true
}

// HasCategoricalValue returns a boolean if a field has been set.
func (o *LLMObsExperimentMetric) HasCategoricalValue() bool {
	return o != nil && o.CategoricalValue != nil
}

// SetCategoricalValue gets a reference to the given string and assigns it to the CategoricalValue field.
func (o *LLMObsExperimentMetric) SetCategoricalValue(v string) {
	o.CategoricalValue = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *LLMObsExperimentMetric) GetError() LLMObsExperimentMetricError {
	if o == nil || o.Error == nil {
		var ret LLMObsExperimentMetricError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetErrorOk() (*LLMObsExperimentMetricError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *LLMObsExperimentMetric) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given LLMObsExperimentMetricError and assigns it to the Error field.
func (o *LLMObsExperimentMetric) SetError(v LLMObsExperimentMetricError) {
	o.Error = &v
}

// GetJsonValue returns the JsonValue field value if set, zero value otherwise.
func (o *LLMObsExperimentMetric) GetJsonValue() map[string]interface{} {
	if o == nil || o.JsonValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.JsonValue
}

// GetJsonValueOk returns a tuple with the JsonValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetJsonValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.JsonValue == nil {
		return nil, false
	}
	return &o.JsonValue, true
}

// HasJsonValue returns a boolean if a field has been set.
func (o *LLMObsExperimentMetric) HasJsonValue() bool {
	return o != nil && o.JsonValue != nil
}

// SetJsonValue gets a reference to the given map[string]interface{} and assigns it to the JsonValue field.
func (o *LLMObsExperimentMetric) SetJsonValue(v map[string]interface{}) {
	o.JsonValue = v
}

// GetLabel returns the Label field value.
func (o *LLMObsExperimentMetric) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *LLMObsExperimentMetric) SetLabel(v string) {
	o.Label = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LLMObsExperimentMetric) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LLMObsExperimentMetric) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LLMObsExperimentMetric) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMetricType returns the MetricType field value.
func (o *LLMObsExperimentMetric) GetMetricType() LLMObsMetricScoreType {
	if o == nil {
		var ret LLMObsMetricScoreType
		return ret
	}
	return o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetMetricTypeOk() (*LLMObsMetricScoreType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricType, true
}

// SetMetricType sets field value.
func (o *LLMObsExperimentMetric) SetMetricType(v LLMObsMetricScoreType) {
	o.MetricType = v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *LLMObsExperimentMetric) GetReasoning() string {
	if o == nil || o.Reasoning == nil {
		var ret string
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetReasoningOk() (*string, bool) {
	if o == nil || o.Reasoning == nil {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *LLMObsExperimentMetric) HasReasoning() bool {
	return o != nil && o.Reasoning != nil
}

// SetReasoning gets a reference to the given string and assigns it to the Reasoning field.
func (o *LLMObsExperimentMetric) SetReasoning(v string) {
	o.Reasoning = &v
}

// GetScoreValue returns the ScoreValue field value if set, zero value otherwise.
func (o *LLMObsExperimentMetric) GetScoreValue() float64 {
	if o == nil || o.ScoreValue == nil {
		var ret float64
		return ret
	}
	return *o.ScoreValue
}

// GetScoreValueOk returns a tuple with the ScoreValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetScoreValueOk() (*float64, bool) {
	if o == nil || o.ScoreValue == nil {
		return nil, false
	}
	return o.ScoreValue, true
}

// HasScoreValue returns a boolean if a field has been set.
func (o *LLMObsExperimentMetric) HasScoreValue() bool {
	return o != nil && o.ScoreValue != nil
}

// SetScoreValue gets a reference to the given float64 and assigns it to the ScoreValue field.
func (o *LLMObsExperimentMetric) SetScoreValue(v float64) {
	o.ScoreValue = &v
}

// GetSpanId returns the SpanId field value.
func (o *LLMObsExperimentMetric) GetSpanId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetSpanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanId, true
}

// SetSpanId sets field value.
func (o *LLMObsExperimentMetric) SetSpanId(v string) {
	o.SpanId = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LLMObsExperimentMetric) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LLMObsExperimentMetric) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LLMObsExperimentMetric) SetTags(v []string) {
	o.Tags = v
}

// GetTimestampMs returns the TimestampMs field value.
func (o *LLMObsExperimentMetric) GetTimestampMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TimestampMs
}

// GetTimestampMsOk returns a tuple with the TimestampMs field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentMetric) GetTimestampMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampMs, true
}

// SetTimestampMs sets field value.
func (o *LLMObsExperimentMetric) SetTimestampMs(v int64) {
	o.TimestampMs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assessment != nil {
		toSerialize["assessment"] = o.Assessment
	}
	if o.BooleanValue != nil {
		toSerialize["boolean_value"] = o.BooleanValue
	}
	if o.CategoricalValue != nil {
		toSerialize["categorical_value"] = o.CategoricalValue
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.JsonValue != nil {
		toSerialize["json_value"] = o.JsonValue
	}
	toSerialize["label"] = o.Label
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["metric_type"] = o.MetricType
	if o.Reasoning != nil {
		toSerialize["reasoning"] = o.Reasoning
	}
	if o.ScoreValue != nil {
		toSerialize["score_value"] = o.ScoreValue
	}
	toSerialize["span_id"] = o.SpanId
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["timestamp_ms"] = o.TimestampMs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentMetric) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assessment       *LLMObsMetricAssessment      `json:"assessment,omitempty"`
		BooleanValue     *bool                        `json:"boolean_value,omitempty"`
		CategoricalValue *string                      `json:"categorical_value,omitempty"`
		Error            *LLMObsExperimentMetricError `json:"error,omitempty"`
		JsonValue        map[string]interface{}       `json:"json_value,omitempty"`
		Label            *string                      `json:"label"`
		Metadata         map[string]interface{}       `json:"metadata,omitempty"`
		MetricType       *LLMObsMetricScoreType       `json:"metric_type"`
		Reasoning        *string                      `json:"reasoning,omitempty"`
		ScoreValue       *float64                     `json:"score_value,omitempty"`
		SpanId           *string                      `json:"span_id"`
		Tags             []string                     `json:"tags,omitempty"`
		TimestampMs      *int64                       `json:"timestamp_ms"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.MetricType == nil {
		return fmt.Errorf("required field metric_type missing")
	}
	if all.SpanId == nil {
		return fmt.Errorf("required field span_id missing")
	}
	if all.TimestampMs == nil {
		return fmt.Errorf("required field timestamp_ms missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assessment", "boolean_value", "categorical_value", "error", "json_value", "label", "metadata", "metric_type", "reasoning", "score_value", "span_id", "tags", "timestamp_ms"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Assessment != nil && !all.Assessment.IsValid() {
		hasInvalidField = true
	} else {
		o.Assessment = all.Assessment
	}
	o.BooleanValue = all.BooleanValue
	o.CategoricalValue = all.CategoricalValue
	if all.Error != nil && all.Error.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Error = all.Error
	o.JsonValue = all.JsonValue
	o.Label = *all.Label
	o.Metadata = all.Metadata
	if !all.MetricType.IsValid() {
		hasInvalidField = true
	} else {
		o.MetricType = *all.MetricType
	}
	o.Reasoning = all.Reasoning
	o.ScoreValue = all.ScoreValue
	o.SpanId = *all.SpanId
	o.Tags = all.Tags
	o.TimestampMs = *all.TimestampMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
