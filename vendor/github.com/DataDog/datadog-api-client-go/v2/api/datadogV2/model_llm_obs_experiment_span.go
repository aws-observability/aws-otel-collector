// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentSpan A span associated with an LLM Observability experiment.
type LLMObsExperimentSpan struct {
	// Dataset ID associated with this span.
	DatasetId string `json:"dataset_id"`
	// Duration of the span in nanoseconds.
	Duration int64 `json:"duration"`
	// Metadata associated with an experiment span.
	Meta *LLMObsExperimentSpanMeta `json:"meta,omitempty"`
	// Name of the span.
	Name string `json:"name"`
	// Project ID associated with this span.
	ProjectId string `json:"project_id"`
	// Unique identifier of the span.
	SpanId string `json:"span_id"`
	// Start time of the span in nanoseconds since Unix epoch.
	StartNs int64 `json:"start_ns"`
	// Status of the span.
	Status LLMObsExperimentSpanStatus `json:"status"`
	// List of tags associated with the span.
	Tags []string `json:"tags,omitempty"`
	// Trace ID for the span.
	TraceId string `json:"trace_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentSpan instantiates a new LLMObsExperimentSpan object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentSpan(datasetId string, duration int64, name string, projectId string, spanId string, startNs int64, status LLMObsExperimentSpanStatus, traceId string) *LLMObsExperimentSpan {
	this := LLMObsExperimentSpan{}
	this.DatasetId = datasetId
	this.Duration = duration
	this.Name = name
	this.ProjectId = projectId
	this.SpanId = spanId
	this.StartNs = startNs
	this.Status = status
	this.TraceId = traceId
	return &this
}

// NewLLMObsExperimentSpanWithDefaults instantiates a new LLMObsExperimentSpan object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentSpanWithDefaults() *LLMObsExperimentSpan {
	this := LLMObsExperimentSpan{}
	return &this
}

// GetDatasetId returns the DatasetId field value.
func (o *LLMObsExperimentSpan) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpan) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value.
func (o *LLMObsExperimentSpan) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetDuration returns the Duration field value.
func (o *LLMObsExperimentSpan) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpan) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value.
func (o *LLMObsExperimentSpan) SetDuration(v int64) {
	o.Duration = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LLMObsExperimentSpan) GetMeta() LLMObsExperimentSpanMeta {
	if o == nil || o.Meta == nil {
		var ret LLMObsExperimentSpanMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpan) GetMetaOk() (*LLMObsExperimentSpanMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LLMObsExperimentSpan) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given LLMObsExperimentSpanMeta and assigns it to the Meta field.
func (o *LLMObsExperimentSpan) SetMeta(v LLMObsExperimentSpanMeta) {
	o.Meta = &v
}

// GetName returns the Name field value.
func (o *LLMObsExperimentSpan) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpan) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsExperimentSpan) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value.
func (o *LLMObsExperimentSpan) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpan) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *LLMObsExperimentSpan) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSpanId returns the SpanId field value.
func (o *LLMObsExperimentSpan) GetSpanId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpan) GetSpanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanId, true
}

// SetSpanId sets field value.
func (o *LLMObsExperimentSpan) SetSpanId(v string) {
	o.SpanId = v
}

// GetStartNs returns the StartNs field value.
func (o *LLMObsExperimentSpan) GetStartNs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StartNs
}

// GetStartNsOk returns a tuple with the StartNs field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpan) GetStartNsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartNs, true
}

// SetStartNs sets field value.
func (o *LLMObsExperimentSpan) SetStartNs(v int64) {
	o.StartNs = v
}

// GetStatus returns the Status field value.
func (o *LLMObsExperimentSpan) GetStatus() LLMObsExperimentSpanStatus {
	if o == nil {
		var ret LLMObsExperimentSpanStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpan) GetStatusOk() (*LLMObsExperimentSpanStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *LLMObsExperimentSpan) SetStatus(v LLMObsExperimentSpanStatus) {
	o.Status = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LLMObsExperimentSpan) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpan) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LLMObsExperimentSpan) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LLMObsExperimentSpan) SetTags(v []string) {
	o.Tags = v
}

// GetTraceId returns the TraceId field value.
func (o *LLMObsExperimentSpan) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpan) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value.
func (o *LLMObsExperimentSpan) SetTraceId(v string) {
	o.TraceId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentSpan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dataset_id"] = o.DatasetId
	toSerialize["duration"] = o.Duration
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["name"] = o.Name
	toSerialize["project_id"] = o.ProjectId
	toSerialize["span_id"] = o.SpanId
	toSerialize["start_ns"] = o.StartNs
	toSerialize["status"] = o.Status
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["trace_id"] = o.TraceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentSpan) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasetId *string                     `json:"dataset_id"`
		Duration  *int64                      `json:"duration"`
		Meta      *LLMObsExperimentSpanMeta   `json:"meta,omitempty"`
		Name      *string                     `json:"name"`
		ProjectId *string                     `json:"project_id"`
		SpanId    *string                     `json:"span_id"`
		StartNs   *int64                      `json:"start_ns"`
		Status    *LLMObsExperimentSpanStatus `json:"status"`
		Tags      []string                    `json:"tags,omitempty"`
		TraceId   *string                     `json:"trace_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatasetId == nil {
		return fmt.Errorf("required field dataset_id missing")
	}
	if all.Duration == nil {
		return fmt.Errorf("required field duration missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.SpanId == nil {
		return fmt.Errorf("required field span_id missing")
	}
	if all.StartNs == nil {
		return fmt.Errorf("required field start_ns missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.TraceId == nil {
		return fmt.Errorf("required field trace_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dataset_id", "duration", "meta", "name", "project_id", "span_id", "start_ns", "status", "tags", "trace_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DatasetId = *all.DatasetId
	o.Duration = *all.Duration
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	o.Name = *all.Name
	o.ProjectId = *all.ProjectId
	o.SpanId = *all.SpanId
	o.StartNs = *all.StartNs
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Tags = all.Tags
	o.TraceId = *all.TraceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
