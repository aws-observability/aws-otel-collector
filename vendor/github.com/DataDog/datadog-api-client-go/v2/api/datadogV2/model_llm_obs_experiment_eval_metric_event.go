// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentEvalMetricEvent An evaluation metric event associated with an experiment span.
type LLMObsExperimentEvalMetricEvent struct {
	// Assessment result for an LLM Observability experiment metric.
	Assessment *LLMObsMetricAssessment `json:"assessment,omitempty"`
	// Boolean value. Present when `metric_type` is `boolean`.
	BooleanValue datadog.NullableBool `json:"boolean_value,omitempty"`
	// Categorical value. Present when `metric_type` is `categorical`.
	CategoricalValue datadog.NullableString `json:"categorical_value,omitempty"`
	// Source type of the evaluation.
	EvalSourceType *string `json:"eval_source_type,omitempty"`
	// Unique identifier of the evaluation metric event.
	Id *string `json:"id,omitempty"`
	// JSON value. Present when `metric_type` is `json`.
	JsonValue map[string]interface{} `json:"json_value,omitempty"`
	// Label or name for the metric.
	Label *string `json:"label,omitempty"`
	// Arbitrary metadata associated with the metric.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Source of the metric. Either `custom` (user-submitted) or `summary` (experiment-level aggregate).
	MetricSource *string `json:"metric_source,omitempty"`
	// Type of metric recorded for an LLM Observability experiment.
	MetricType *LLMObsMetricScoreType `json:"metric_type,omitempty"`
	// Human-readable reasoning for the metric value.
	Reasoning datadog.NullableString `json:"reasoning,omitempty"`
	// Numeric score. Present when `metric_type` is `score`.
	ScoreValue datadog.NullableFloat64 `json:"score_value,omitempty"`
	// Span ID this metric is associated with.
	SpanId *string `json:"span_id,omitempty"`
	// Tags associated with the metric.
	Tags []string `json:"tags,omitempty"`
	// Timestamp when the metric was recorded, in milliseconds since Unix epoch.
	TimestampMs *int64 `json:"timestamp_ms,omitempty"`
	// Trace ID linking this metric to a span.
	TraceId *string `json:"trace_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentEvalMetricEvent instantiates a new LLMObsExperimentEvalMetricEvent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentEvalMetricEvent() *LLMObsExperimentEvalMetricEvent {
	this := LLMObsExperimentEvalMetricEvent{}
	return &this
}

// NewLLMObsExperimentEvalMetricEventWithDefaults instantiates a new LLMObsExperimentEvalMetricEvent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentEvalMetricEventWithDefaults() *LLMObsExperimentEvalMetricEvent {
	this := LLMObsExperimentEvalMetricEvent{}
	return &this
}

// GetAssessment returns the Assessment field value if set, zero value otherwise.
func (o *LLMObsExperimentEvalMetricEvent) GetAssessment() LLMObsMetricAssessment {
	if o == nil || o.Assessment == nil {
		var ret LLMObsMetricAssessment
		return ret
	}
	return *o.Assessment
}

// GetAssessmentOk returns a tuple with the Assessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEvalMetricEvent) GetAssessmentOk() (*LLMObsMetricAssessment, bool) {
	if o == nil || o.Assessment == nil {
		return nil, false
	}
	return o.Assessment, true
}

// HasAssessment returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasAssessment() bool {
	return o != nil && o.Assessment != nil
}

// SetAssessment gets a reference to the given LLMObsMetricAssessment and assigns it to the Assessment field.
func (o *LLMObsExperimentEvalMetricEvent) SetAssessment(v LLMObsMetricAssessment) {
	o.Assessment = &v
}

// GetBooleanValue returns the BooleanValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentEvalMetricEvent) GetBooleanValue() bool {
	if o == nil || o.BooleanValue.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BooleanValue.Get()
}

// GetBooleanValueOk returns a tuple with the BooleanValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentEvalMetricEvent) GetBooleanValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BooleanValue.Get(), o.BooleanValue.IsSet()
}

// HasBooleanValue returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasBooleanValue() bool {
	return o != nil && o.BooleanValue.IsSet()
}

// SetBooleanValue gets a reference to the given datadog.NullableBool and assigns it to the BooleanValue field.
func (o *LLMObsExperimentEvalMetricEvent) SetBooleanValue(v bool) {
	o.BooleanValue.Set(&v)
}

// SetBooleanValueNil sets the value for BooleanValue to be an explicit nil.
func (o *LLMObsExperimentEvalMetricEvent) SetBooleanValueNil() {
	o.BooleanValue.Set(nil)
}

// UnsetBooleanValue ensures that no value is present for BooleanValue, not even an explicit nil.
func (o *LLMObsExperimentEvalMetricEvent) UnsetBooleanValue() {
	o.BooleanValue.Unset()
}

// GetCategoricalValue returns the CategoricalValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentEvalMetricEvent) GetCategoricalValue() string {
	if o == nil || o.CategoricalValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.CategoricalValue.Get()
}

// GetCategoricalValueOk returns a tuple with the CategoricalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentEvalMetricEvent) GetCategoricalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CategoricalValue.Get(), o.CategoricalValue.IsSet()
}

// HasCategoricalValue returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasCategoricalValue() bool {
	return o != nil && o.CategoricalValue.IsSet()
}

// SetCategoricalValue gets a reference to the given datadog.NullableString and assigns it to the CategoricalValue field.
func (o *LLMObsExperimentEvalMetricEvent) SetCategoricalValue(v string) {
	o.CategoricalValue.Set(&v)
}

// SetCategoricalValueNil sets the value for CategoricalValue to be an explicit nil.
func (o *LLMObsExperimentEvalMetricEvent) SetCategoricalValueNil() {
	o.CategoricalValue.Set(nil)
}

// UnsetCategoricalValue ensures that no value is present for CategoricalValue, not even an explicit nil.
func (o *LLMObsExperimentEvalMetricEvent) UnsetCategoricalValue() {
	o.CategoricalValue.Unset()
}

// GetEvalSourceType returns the EvalSourceType field value if set, zero value otherwise.
func (o *LLMObsExperimentEvalMetricEvent) GetEvalSourceType() string {
	if o == nil || o.EvalSourceType == nil {
		var ret string
		return ret
	}
	return *o.EvalSourceType
}

// GetEvalSourceTypeOk returns a tuple with the EvalSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEvalMetricEvent) GetEvalSourceTypeOk() (*string, bool) {
	if o == nil || o.EvalSourceType == nil {
		return nil, false
	}
	return o.EvalSourceType, true
}

// HasEvalSourceType returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasEvalSourceType() bool {
	return o != nil && o.EvalSourceType != nil
}

// SetEvalSourceType gets a reference to the given string and assigns it to the EvalSourceType field.
func (o *LLMObsExperimentEvalMetricEvent) SetEvalSourceType(v string) {
	o.EvalSourceType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LLMObsExperimentEvalMetricEvent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEvalMetricEvent) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LLMObsExperimentEvalMetricEvent) SetId(v string) {
	o.Id = &v
}

// GetJsonValue returns the JsonValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentEvalMetricEvent) GetJsonValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.JsonValue
}

// GetJsonValueOk returns a tuple with the JsonValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentEvalMetricEvent) GetJsonValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.JsonValue == nil {
		return nil, false
	}
	return &o.JsonValue, true
}

// HasJsonValue returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasJsonValue() bool {
	return o != nil && o.JsonValue != nil
}

// SetJsonValue gets a reference to the given map[string]interface{} and assigns it to the JsonValue field.
func (o *LLMObsExperimentEvalMetricEvent) SetJsonValue(v map[string]interface{}) {
	o.JsonValue = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *LLMObsExperimentEvalMetricEvent) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEvalMetricEvent) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *LLMObsExperimentEvalMetricEvent) SetLabel(v string) {
	o.Label = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentEvalMetricEvent) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentEvalMetricEvent) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LLMObsExperimentEvalMetricEvent) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMetricSource returns the MetricSource field value if set, zero value otherwise.
func (o *LLMObsExperimentEvalMetricEvent) GetMetricSource() string {
	if o == nil || o.MetricSource == nil {
		var ret string
		return ret
	}
	return *o.MetricSource
}

// GetMetricSourceOk returns a tuple with the MetricSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEvalMetricEvent) GetMetricSourceOk() (*string, bool) {
	if o == nil || o.MetricSource == nil {
		return nil, false
	}
	return o.MetricSource, true
}

// HasMetricSource returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasMetricSource() bool {
	return o != nil && o.MetricSource != nil
}

// SetMetricSource gets a reference to the given string and assigns it to the MetricSource field.
func (o *LLMObsExperimentEvalMetricEvent) SetMetricSource(v string) {
	o.MetricSource = &v
}

// GetMetricType returns the MetricType field value if set, zero value otherwise.
func (o *LLMObsExperimentEvalMetricEvent) GetMetricType() LLMObsMetricScoreType {
	if o == nil || o.MetricType == nil {
		var ret LLMObsMetricScoreType
		return ret
	}
	return *o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEvalMetricEvent) GetMetricTypeOk() (*LLMObsMetricScoreType, bool) {
	if o == nil || o.MetricType == nil {
		return nil, false
	}
	return o.MetricType, true
}

// HasMetricType returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasMetricType() bool {
	return o != nil && o.MetricType != nil
}

// SetMetricType gets a reference to the given LLMObsMetricScoreType and assigns it to the MetricType field.
func (o *LLMObsExperimentEvalMetricEvent) SetMetricType(v LLMObsMetricScoreType) {
	o.MetricType = &v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentEvalMetricEvent) GetReasoning() string {
	if o == nil || o.Reasoning.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reasoning.Get()
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentEvalMetricEvent) GetReasoningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reasoning.Get(), o.Reasoning.IsSet()
}

// HasReasoning returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasReasoning() bool {
	return o != nil && o.Reasoning.IsSet()
}

// SetReasoning gets a reference to the given datadog.NullableString and assigns it to the Reasoning field.
func (o *LLMObsExperimentEvalMetricEvent) SetReasoning(v string) {
	o.Reasoning.Set(&v)
}

// SetReasoningNil sets the value for Reasoning to be an explicit nil.
func (o *LLMObsExperimentEvalMetricEvent) SetReasoningNil() {
	o.Reasoning.Set(nil)
}

// UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil.
func (o *LLMObsExperimentEvalMetricEvent) UnsetReasoning() {
	o.Reasoning.Unset()
}

// GetScoreValue returns the ScoreValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentEvalMetricEvent) GetScoreValue() float64 {
	if o == nil || o.ScoreValue.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ScoreValue.Get()
}

// GetScoreValueOk returns a tuple with the ScoreValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentEvalMetricEvent) GetScoreValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScoreValue.Get(), o.ScoreValue.IsSet()
}

// HasScoreValue returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasScoreValue() bool {
	return o != nil && o.ScoreValue.IsSet()
}

// SetScoreValue gets a reference to the given datadog.NullableFloat64 and assigns it to the ScoreValue field.
func (o *LLMObsExperimentEvalMetricEvent) SetScoreValue(v float64) {
	o.ScoreValue.Set(&v)
}

// SetScoreValueNil sets the value for ScoreValue to be an explicit nil.
func (o *LLMObsExperimentEvalMetricEvent) SetScoreValueNil() {
	o.ScoreValue.Set(nil)
}

// UnsetScoreValue ensures that no value is present for ScoreValue, not even an explicit nil.
func (o *LLMObsExperimentEvalMetricEvent) UnsetScoreValue() {
	o.ScoreValue.Unset()
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *LLMObsExperimentEvalMetricEvent) GetSpanId() string {
	if o == nil || o.SpanId == nil {
		var ret string
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEvalMetricEvent) GetSpanIdOk() (*string, bool) {
	if o == nil || o.SpanId == nil {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasSpanId() bool {
	return o != nil && o.SpanId != nil
}

// SetSpanId gets a reference to the given string and assigns it to the SpanId field.
func (o *LLMObsExperimentEvalMetricEvent) SetSpanId(v string) {
	o.SpanId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LLMObsExperimentEvalMetricEvent) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEvalMetricEvent) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LLMObsExperimentEvalMetricEvent) SetTags(v []string) {
	o.Tags = v
}

// GetTimestampMs returns the TimestampMs field value if set, zero value otherwise.
func (o *LLMObsExperimentEvalMetricEvent) GetTimestampMs() int64 {
	if o == nil || o.TimestampMs == nil {
		var ret int64
		return ret
	}
	return *o.TimestampMs
}

// GetTimestampMsOk returns a tuple with the TimestampMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEvalMetricEvent) GetTimestampMsOk() (*int64, bool) {
	if o == nil || o.TimestampMs == nil {
		return nil, false
	}
	return o.TimestampMs, true
}

// HasTimestampMs returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasTimestampMs() bool {
	return o != nil && o.TimestampMs != nil
}

// SetTimestampMs gets a reference to the given int64 and assigns it to the TimestampMs field.
func (o *LLMObsExperimentEvalMetricEvent) SetTimestampMs(v int64) {
	o.TimestampMs = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *LLMObsExperimentEvalMetricEvent) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEvalMetricEvent) GetTraceIdOk() (*string, bool) {
	if o == nil || o.TraceId == nil {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *LLMObsExperimentEvalMetricEvent) HasTraceId() bool {
	return o != nil && o.TraceId != nil
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *LLMObsExperimentEvalMetricEvent) SetTraceId(v string) {
	o.TraceId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentEvalMetricEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assessment != nil {
		toSerialize["assessment"] = o.Assessment
	}
	if o.BooleanValue.IsSet() {
		toSerialize["boolean_value"] = o.BooleanValue.Get()
	}
	if o.CategoricalValue.IsSet() {
		toSerialize["categorical_value"] = o.CategoricalValue.Get()
	}
	if o.EvalSourceType != nil {
		toSerialize["eval_source_type"] = o.EvalSourceType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.JsonValue != nil {
		toSerialize["json_value"] = o.JsonValue
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.MetricSource != nil {
		toSerialize["metric_source"] = o.MetricSource
	}
	if o.MetricType != nil {
		toSerialize["metric_type"] = o.MetricType
	}
	if o.Reasoning.IsSet() {
		toSerialize["reasoning"] = o.Reasoning.Get()
	}
	if o.ScoreValue.IsSet() {
		toSerialize["score_value"] = o.ScoreValue.Get()
	}
	if o.SpanId != nil {
		toSerialize["span_id"] = o.SpanId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TimestampMs != nil {
		toSerialize["timestamp_ms"] = o.TimestampMs
	}
	if o.TraceId != nil {
		toSerialize["trace_id"] = o.TraceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentEvalMetricEvent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assessment       *LLMObsMetricAssessment `json:"assessment,omitempty"`
		BooleanValue     datadog.NullableBool    `json:"boolean_value,omitempty"`
		CategoricalValue datadog.NullableString  `json:"categorical_value,omitempty"`
		EvalSourceType   *string                 `json:"eval_source_type,omitempty"`
		Id               *string                 `json:"id,omitempty"`
		JsonValue        map[string]interface{}  `json:"json_value,omitempty"`
		Label            *string                 `json:"label,omitempty"`
		Metadata         map[string]interface{}  `json:"metadata,omitempty"`
		MetricSource     *string                 `json:"metric_source,omitempty"`
		MetricType       *LLMObsMetricScoreType  `json:"metric_type,omitempty"`
		Reasoning        datadog.NullableString  `json:"reasoning,omitempty"`
		ScoreValue       datadog.NullableFloat64 `json:"score_value,omitempty"`
		SpanId           *string                 `json:"span_id,omitempty"`
		Tags             []string                `json:"tags,omitempty"`
		TimestampMs      *int64                  `json:"timestamp_ms,omitempty"`
		TraceId          *string                 `json:"trace_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assessment", "boolean_value", "categorical_value", "eval_source_type", "id", "json_value", "label", "metadata", "metric_source", "metric_type", "reasoning", "score_value", "span_id", "tags", "timestamp_ms", "trace_id"})
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
	o.EvalSourceType = all.EvalSourceType
	o.Id = all.Id
	o.JsonValue = all.JsonValue
	o.Label = all.Label
	o.Metadata = all.Metadata
	o.MetricSource = all.MetricSource
	if all.MetricType != nil && !all.MetricType.IsValid() {
		hasInvalidField = true
	} else {
		o.MetricType = all.MetricType
	}
	o.Reasoning = all.Reasoning
	o.ScoreValue = all.ScoreValue
	o.SpanId = all.SpanId
	o.Tags = all.Tags
	o.TimestampMs = all.TimestampMs
	o.TraceId = all.TraceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
