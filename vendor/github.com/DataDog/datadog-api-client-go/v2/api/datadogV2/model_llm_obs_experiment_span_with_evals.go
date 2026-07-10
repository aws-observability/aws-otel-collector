// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentSpanWithEvals An experiment span enriched with its associated evaluation metrics.
type LLMObsExperimentSpanWithEvals struct {
	// ID of the dataset record this span evaluated.
	DatasetRecordId datadog.NullableString `json:"dataset_record_id,omitempty"`
	// Duration of the span in nanoseconds.
	Duration *float64 `json:"duration,omitempty"`
	// Evaluation metrics associated with this span.
	EvalMetrics []LLMObsExperimentEvalMetricEvent `json:"eval_metrics,omitempty"`
	// Unique identifier of the span.
	Id *string `json:"id,omitempty"`
	// Metadata associated with an experiment span.
	Meta *LLMObsExperimentSpanMeta `json:"meta,omitempty"`
	// Numeric metrics attached to the span.
	Metrics map[string]float64 `json:"metrics,omitempty"`
	// Name of the span.
	Name *string `json:"name,omitempty"`
	// Parent span ID, if any.
	ParentId *string `json:"parent_id,omitempty"`
	// Span ID.
	SpanId *string `json:"span_id,omitempty"`
	// Start time in nanoseconds since Unix epoch.
	StartNs *int64 `json:"start_ns,omitempty"`
	// Status of the span.
	Status *LLMObsExperimentSpanStatus `json:"status,omitempty"`
	// Tags associated with the span.
	Tags []string `json:"tags,omitempty"`
	// Trace ID.
	TraceId *string `json:"trace_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentSpanWithEvals instantiates a new LLMObsExperimentSpanWithEvals object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentSpanWithEvals() *LLMObsExperimentSpanWithEvals {
	this := LLMObsExperimentSpanWithEvals{}
	return &this
}

// NewLLMObsExperimentSpanWithEvalsWithDefaults instantiates a new LLMObsExperimentSpanWithEvals object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentSpanWithEvalsWithDefaults() *LLMObsExperimentSpanWithEvals {
	this := LLMObsExperimentSpanWithEvals{}
	return &this
}

// GetDatasetRecordId returns the DatasetRecordId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentSpanWithEvals) GetDatasetRecordId() string {
	if o == nil || o.DatasetRecordId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatasetRecordId.Get()
}

// GetDatasetRecordIdOk returns a tuple with the DatasetRecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentSpanWithEvals) GetDatasetRecordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatasetRecordId.Get(), o.DatasetRecordId.IsSet()
}

// HasDatasetRecordId returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasDatasetRecordId() bool {
	return o != nil && o.DatasetRecordId.IsSet()
}

// SetDatasetRecordId gets a reference to the given datadog.NullableString and assigns it to the DatasetRecordId field.
func (o *LLMObsExperimentSpanWithEvals) SetDatasetRecordId(v string) {
	o.DatasetRecordId.Set(&v)
}

// SetDatasetRecordIdNil sets the value for DatasetRecordId to be an explicit nil.
func (o *LLMObsExperimentSpanWithEvals) SetDatasetRecordIdNil() {
	o.DatasetRecordId.Set(nil)
}

// UnsetDatasetRecordId ensures that no value is present for DatasetRecordId, not even an explicit nil.
func (o *LLMObsExperimentSpanWithEvals) UnsetDatasetRecordId() {
	o.DatasetRecordId.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanWithEvals) GetDuration() float64 {
	if o == nil || o.Duration == nil {
		var ret float64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanWithEvals) GetDurationOk() (*float64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given float64 and assigns it to the Duration field.
func (o *LLMObsExperimentSpanWithEvals) SetDuration(v float64) {
	o.Duration = &v
}

// GetEvalMetrics returns the EvalMetrics field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanWithEvals) GetEvalMetrics() []LLMObsExperimentEvalMetricEvent {
	if o == nil || o.EvalMetrics == nil {
		var ret []LLMObsExperimentEvalMetricEvent
		return ret
	}
	return o.EvalMetrics
}

// GetEvalMetricsOk returns a tuple with the EvalMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanWithEvals) GetEvalMetricsOk() (*[]LLMObsExperimentEvalMetricEvent, bool) {
	if o == nil || o.EvalMetrics == nil {
		return nil, false
	}
	return &o.EvalMetrics, true
}

// HasEvalMetrics returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasEvalMetrics() bool {
	return o != nil && o.EvalMetrics != nil
}

// SetEvalMetrics gets a reference to the given []LLMObsExperimentEvalMetricEvent and assigns it to the EvalMetrics field.
func (o *LLMObsExperimentSpanWithEvals) SetEvalMetrics(v []LLMObsExperimentEvalMetricEvent) {
	o.EvalMetrics = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanWithEvals) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanWithEvals) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LLMObsExperimentSpanWithEvals) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanWithEvals) GetMeta() LLMObsExperimentSpanMeta {
	if o == nil || o.Meta == nil {
		var ret LLMObsExperimentSpanMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanWithEvals) GetMetaOk() (*LLMObsExperimentSpanMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given LLMObsExperimentSpanMeta and assigns it to the Meta field.
func (o *LLMObsExperimentSpanWithEvals) SetMeta(v LLMObsExperimentSpanMeta) {
	o.Meta = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanWithEvals) GetMetrics() map[string]float64 {
	if o == nil || o.Metrics == nil {
		var ret map[string]float64
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanWithEvals) GetMetricsOk() (*map[string]float64, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given map[string]float64 and assigns it to the Metrics field.
func (o *LLMObsExperimentSpanWithEvals) SetMetrics(v map[string]float64) {
	o.Metrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanWithEvals) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanWithEvals) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LLMObsExperimentSpanWithEvals) SetName(v string) {
	o.Name = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanWithEvals) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanWithEvals) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasParentId() bool {
	return o != nil && o.ParentId != nil
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *LLMObsExperimentSpanWithEvals) SetParentId(v string) {
	o.ParentId = &v
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanWithEvals) GetSpanId() string {
	if o == nil || o.SpanId == nil {
		var ret string
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanWithEvals) GetSpanIdOk() (*string, bool) {
	if o == nil || o.SpanId == nil {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasSpanId() bool {
	return o != nil && o.SpanId != nil
}

// SetSpanId gets a reference to the given string and assigns it to the SpanId field.
func (o *LLMObsExperimentSpanWithEvals) SetSpanId(v string) {
	o.SpanId = &v
}

// GetStartNs returns the StartNs field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanWithEvals) GetStartNs() int64 {
	if o == nil || o.StartNs == nil {
		var ret int64
		return ret
	}
	return *o.StartNs
}

// GetStartNsOk returns a tuple with the StartNs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanWithEvals) GetStartNsOk() (*int64, bool) {
	if o == nil || o.StartNs == nil {
		return nil, false
	}
	return o.StartNs, true
}

// HasStartNs returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasStartNs() bool {
	return o != nil && o.StartNs != nil
}

// SetStartNs gets a reference to the given int64 and assigns it to the StartNs field.
func (o *LLMObsExperimentSpanWithEvals) SetStartNs(v int64) {
	o.StartNs = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanWithEvals) GetStatus() LLMObsExperimentSpanStatus {
	if o == nil || o.Status == nil {
		var ret LLMObsExperimentSpanStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanWithEvals) GetStatusOk() (*LLMObsExperimentSpanStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given LLMObsExperimentSpanStatus and assigns it to the Status field.
func (o *LLMObsExperimentSpanWithEvals) SetStatus(v LLMObsExperimentSpanStatus) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanWithEvals) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanWithEvals) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LLMObsExperimentSpanWithEvals) SetTags(v []string) {
	o.Tags = v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanWithEvals) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanWithEvals) GetTraceIdOk() (*string, bool) {
	if o == nil || o.TraceId == nil {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanWithEvals) HasTraceId() bool {
	return o != nil && o.TraceId != nil
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *LLMObsExperimentSpanWithEvals) SetTraceId(v string) {
	o.TraceId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentSpanWithEvals) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DatasetRecordId.IsSet() {
		toSerialize["dataset_record_id"] = o.DatasetRecordId.Get()
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.EvalMetrics != nil {
		toSerialize["eval_metrics"] = o.EvalMetrics
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ParentId != nil {
		toSerialize["parent_id"] = o.ParentId
	}
	if o.SpanId != nil {
		toSerialize["span_id"] = o.SpanId
	}
	if o.StartNs != nil {
		toSerialize["start_ns"] = o.StartNs
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
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
func (o *LLMObsExperimentSpanWithEvals) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasetRecordId datadog.NullableString            `json:"dataset_record_id,omitempty"`
		Duration        *float64                          `json:"duration,omitempty"`
		EvalMetrics     []LLMObsExperimentEvalMetricEvent `json:"eval_metrics,omitempty"`
		Id              *string                           `json:"id,omitempty"`
		Meta            *LLMObsExperimentSpanMeta         `json:"meta,omitempty"`
		Metrics         map[string]float64                `json:"metrics,omitempty"`
		Name            *string                           `json:"name,omitempty"`
		ParentId        *string                           `json:"parent_id,omitempty"`
		SpanId          *string                           `json:"span_id,omitempty"`
		StartNs         *int64                            `json:"start_ns,omitempty"`
		Status          *LLMObsExperimentSpanStatus       `json:"status,omitempty"`
		Tags            []string                          `json:"tags,omitempty"`
		TraceId         *string                           `json:"trace_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dataset_record_id", "duration", "eval_metrics", "id", "meta", "metrics", "name", "parent_id", "span_id", "start_ns", "status", "tags", "trace_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DatasetRecordId = all.DatasetRecordId
	o.Duration = all.Duration
	o.EvalMetrics = all.EvalMetrics
	o.Id = all.Id
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	o.Metrics = all.Metrics
	o.Name = all.Name
	o.ParentId = all.ParentId
	o.SpanId = all.SpanId
	o.StartNs = all.StartNs
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Tags = all.Tags
	o.TraceId = all.TraceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
