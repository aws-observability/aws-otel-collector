// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SummarizedSpan A node in the pruned trace tree.
type SummarizedSpan struct {
	// The child spans of this node in the pruned tree.
	Children []SummarizedSpan `json:"children"`
	// The duration of the span, in seconds.
	DurationSeconds float64 `json:"durationSeconds"`
	// The end time of the span, in RFC3339 format.
	EndTime time.Time `json:"endTime"`
	// Error flag for a span. `1` when the span is in error, `0` otherwise.
	Error APMSpanErrorFlag `json:"error"`
	// The number of child spans that were pruned from this node when summarizing the trace.
	HiddenChildSpansCount int32 `json:"hidden_child_spans_count"`
	// String-valued tags attached to the span.
	Meta map[string]string `json:"meta"`
	// Numeric metrics attached to the span.
	Metrics map[string]float64 `json:"metrics"`
	// The operation name of the span.
	Name string `json:"name"`
	// The ID of the parent span, or `0` when the span is the trace root.
	ParentId int64 `json:"parentID"`
	// The resource that the span describes.
	Resource string `json:"resource"`
	// The name of the service that emitted the span.
	Service string `json:"service"`
	// The span ID, as an unsigned 64-bit integer.
	SpanId int64 `json:"spanID"`
	// The OpenTelemetry span kind, for example `INTERNAL`, `SERVER`, `CLIENT`,
	// `PRODUCER`, or `CONSUMER`.
	SpanKind string `json:"span_kind"`
	// The start time of the span, in RFC3339 format.
	StartTime time.Time `json:"startTime"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSummarizedSpan instantiates a new SummarizedSpan object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSummarizedSpan(children []SummarizedSpan, durationSeconds float64, endTime time.Time, error APMSpanErrorFlag, hiddenChildSpansCount int32, meta map[string]string, metrics map[string]float64, name string, parentId int64, resource string, service string, spanId int64, spanKind string, startTime time.Time) *SummarizedSpan {
	this := SummarizedSpan{}
	this.Children = children
	this.DurationSeconds = durationSeconds
	this.EndTime = endTime
	this.Error = error
	this.HiddenChildSpansCount = hiddenChildSpansCount
	this.Meta = meta
	this.Metrics = metrics
	this.Name = name
	this.ParentId = parentId
	this.Resource = resource
	this.Service = service
	this.SpanId = spanId
	this.SpanKind = spanKind
	this.StartTime = startTime
	return &this
}

// NewSummarizedSpanWithDefaults instantiates a new SummarizedSpan object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSummarizedSpanWithDefaults() *SummarizedSpan {
	this := SummarizedSpan{}
	return &this
}

// GetChildren returns the Children field value.
func (o *SummarizedSpan) GetChildren() []SummarizedSpan {
	if o == nil {
		var ret []SummarizedSpan
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetChildrenOk() (*[]SummarizedSpan, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Children, true
}

// SetChildren sets field value.
func (o *SummarizedSpan) SetChildren(v []SummarizedSpan) {
	o.Children = v
}

// GetDurationSeconds returns the DurationSeconds field value.
func (o *SummarizedSpan) GetDurationSeconds() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetDurationSecondsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationSeconds, true
}

// SetDurationSeconds sets field value.
func (o *SummarizedSpan) SetDurationSeconds(v float64) {
	o.DurationSeconds = v
}

// GetEndTime returns the EndTime field value.
func (o *SummarizedSpan) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value.
func (o *SummarizedSpan) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetError returns the Error field value.
func (o *SummarizedSpan) GetError() APMSpanErrorFlag {
	if o == nil {
		var ret APMSpanErrorFlag
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetErrorOk() (*APMSpanErrorFlag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value.
func (o *SummarizedSpan) SetError(v APMSpanErrorFlag) {
	o.Error = v
}

// GetHiddenChildSpansCount returns the HiddenChildSpansCount field value.
func (o *SummarizedSpan) GetHiddenChildSpansCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.HiddenChildSpansCount
}

// GetHiddenChildSpansCountOk returns a tuple with the HiddenChildSpansCount field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetHiddenChildSpansCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HiddenChildSpansCount, true
}

// SetHiddenChildSpansCount sets field value.
func (o *SummarizedSpan) SetHiddenChildSpansCount(v int32) {
	o.HiddenChildSpansCount = v
}

// GetMeta returns the Meta field value.
func (o *SummarizedSpan) GetMeta() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetMetaOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *SummarizedSpan) SetMeta(v map[string]string) {
	o.Meta = v
}

// GetMetrics returns the Metrics field value.
func (o *SummarizedSpan) GetMetrics() map[string]float64 {
	if o == nil {
		var ret map[string]float64
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetMetricsOk() (*map[string]float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value.
func (o *SummarizedSpan) SetMetrics(v map[string]float64) {
	o.Metrics = v
}

// GetName returns the Name field value.
func (o *SummarizedSpan) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SummarizedSpan) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value.
func (o *SummarizedSpan) GetParentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetParentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value.
func (o *SummarizedSpan) SetParentId(v int64) {
	o.ParentId = v
}

// GetResource returns the Resource field value.
func (o *SummarizedSpan) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value.
func (o *SummarizedSpan) SetResource(v string) {
	o.Resource = v
}

// GetService returns the Service field value.
func (o *SummarizedSpan) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *SummarizedSpan) SetService(v string) {
	o.Service = v
}

// GetSpanId returns the SpanId field value.
func (o *SummarizedSpan) GetSpanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetSpanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanId, true
}

// SetSpanId sets field value.
func (o *SummarizedSpan) SetSpanId(v int64) {
	o.SpanId = v
}

// GetSpanKind returns the SpanKind field value.
func (o *SummarizedSpan) GetSpanKind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpanKind
}

// GetSpanKindOk returns a tuple with the SpanKind field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetSpanKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanKind, true
}

// SetSpanKind sets field value.
func (o *SummarizedSpan) SetSpanKind(v string) {
	o.SpanKind = v
}

// GetStartTime returns the StartTime field value.
func (o *SummarizedSpan) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *SummarizedSpan) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value.
func (o *SummarizedSpan) SetStartTime(v time.Time) {
	o.StartTime = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SummarizedSpan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["children"] = o.Children
	toSerialize["durationSeconds"] = o.DurationSeconds
	if o.EndTime.Nanosecond() == 0 {
		toSerialize["endTime"] = o.EndTime.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["endTime"] = o.EndTime.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["error"] = o.Error
	toSerialize["hidden_child_spans_count"] = o.HiddenChildSpansCount
	toSerialize["meta"] = o.Meta
	toSerialize["metrics"] = o.Metrics
	toSerialize["name"] = o.Name
	toSerialize["parentID"] = o.ParentId
	toSerialize["resource"] = o.Resource
	toSerialize["service"] = o.Service
	toSerialize["spanID"] = o.SpanId
	toSerialize["span_kind"] = o.SpanKind
	if o.StartTime.Nanosecond() == 0 {
		toSerialize["startTime"] = o.StartTime.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["startTime"] = o.StartTime.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SummarizedSpan) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Children              *[]SummarizedSpan   `json:"children"`
		DurationSeconds       *float64            `json:"durationSeconds"`
		EndTime               *time.Time          `json:"endTime"`
		Error                 *APMSpanErrorFlag   `json:"error"`
		HiddenChildSpansCount *int32              `json:"hidden_child_spans_count"`
		Meta                  *map[string]string  `json:"meta"`
		Metrics               *map[string]float64 `json:"metrics"`
		Name                  *string             `json:"name"`
		ParentId              *int64              `json:"parentID"`
		Resource              *string             `json:"resource"`
		Service               *string             `json:"service"`
		SpanId                *int64              `json:"spanID"`
		SpanKind              *string             `json:"span_kind"`
		StartTime             *time.Time          `json:"startTime"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Children == nil {
		return fmt.Errorf("required field children missing")
	}
	if all.DurationSeconds == nil {
		return fmt.Errorf("required field durationSeconds missing")
	}
	if all.EndTime == nil {
		return fmt.Errorf("required field endTime missing")
	}
	if all.Error == nil {
		return fmt.Errorf("required field error missing")
	}
	if all.HiddenChildSpansCount == nil {
		return fmt.Errorf("required field hidden_child_spans_count missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	if all.Metrics == nil {
		return fmt.Errorf("required field metrics missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ParentId == nil {
		return fmt.Errorf("required field parentID missing")
	}
	if all.Resource == nil {
		return fmt.Errorf("required field resource missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.SpanId == nil {
		return fmt.Errorf("required field spanID missing")
	}
	if all.SpanKind == nil {
		return fmt.Errorf("required field span_kind missing")
	}
	if all.StartTime == nil {
		return fmt.Errorf("required field startTime missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"children", "durationSeconds", "endTime", "error", "hidden_child_spans_count", "meta", "metrics", "name", "parentID", "resource", "service", "spanID", "span_kind", "startTime"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Children = *all.Children
	o.DurationSeconds = *all.DurationSeconds
	o.EndTime = *all.EndTime
	if !all.Error.IsValid() {
		hasInvalidField = true
	} else {
		o.Error = *all.Error
	}
	o.HiddenChildSpansCount = *all.HiddenChildSpansCount
	o.Meta = *all.Meta
	o.Metrics = *all.Metrics
	o.Name = *all.Name
	o.ParentId = *all.ParentId
	o.Resource = *all.Resource
	o.Service = *all.Service
	o.SpanId = *all.SpanId
	o.SpanKind = *all.SpanKind
	o.StartTime = *all.StartTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
