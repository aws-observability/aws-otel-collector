// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APMTraceSpan A single APM span returned as part of a trace.
type APMTraceSpan struct {
	// The duration of the span, in nanoseconds.
	Duration int64 `json:"duration"`
	// The end time of the span, in Unix nanoseconds.
	EndTime int64 `json:"endTime"`
	// Error flag for a span. `1` when the span is in error, `0` otherwise.
	Error APMSpanErrorFlag `json:"error"`
	// String-valued tags attached to the span. Tag keys starting with `_` are
	// filtered out of the response.
	Meta map[string]string `json:"meta"`
	// Numeric metrics attached to the span. Metric keys starting with `_` are
	// filtered out of the response.
	Metrics map[string]float64 `json:"metrics"`
	// The operation name of the span.
	Name string `json:"name"`
	// The ID of the parent span, or `0` when the span is a trace root.
	ParentId int64 `json:"parentID"`
	// The resource that the span describes.
	Resource string `json:"resource"`
	// A hash of the resource field.
	ResourceHash *string `json:"resourceHash,omitempty"`
	// Whether access to the span is restricted by the organization's data access policies.
	Restricted *bool `json:"restricted,omitempty"`
	// The time spent in the span itself, excluding time spent in child spans, in nanoseconds.
	SelfTime *float64 `json:"self_time,omitempty"`
	// The name of the service that emitted the span.
	Service string `json:"service"`
	// The span ID, as an unsigned 64-bit integer.
	SpanId int64 `json:"spanID"`
	// The start time of the span, in Unix nanoseconds.
	StartTime int64 `json:"startTime"`
	// The lower 64 bits of the trace ID, as an unsigned 64-bit integer.
	TraceId int64 `json:"traceID"`
	// The full 128-bit trace ID, encoded as a 32-character hexadecimal string.
	TraceIdFull string `json:"traceIDFull"`
	// The type of the span (for example, `web`, `db`, or `rpc`).
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAPMTraceSpan instantiates a new APMTraceSpan object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAPMTraceSpan(duration int64, endTime int64, error APMSpanErrorFlag, meta map[string]string, metrics map[string]float64, name string, parentId int64, resource string, service string, spanId int64, startTime int64, traceId int64, traceIdFull string, typeVar string) *APMTraceSpan {
	this := APMTraceSpan{}
	this.Duration = duration
	this.EndTime = endTime
	this.Error = error
	this.Meta = meta
	this.Metrics = metrics
	this.Name = name
	this.ParentId = parentId
	this.Resource = resource
	this.Service = service
	this.SpanId = spanId
	this.StartTime = startTime
	this.TraceId = traceId
	this.TraceIdFull = traceIdFull
	this.Type = typeVar
	return &this
}

// NewAPMTraceSpanWithDefaults instantiates a new APMTraceSpan object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAPMTraceSpanWithDefaults() *APMTraceSpan {
	this := APMTraceSpan{}
	return &this
}

// GetDuration returns the Duration field value.
func (o *APMTraceSpan) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value.
func (o *APMTraceSpan) SetDuration(v int64) {
	o.Duration = v
}

// GetEndTime returns the EndTime field value.
func (o *APMTraceSpan) GetEndTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value.
func (o *APMTraceSpan) SetEndTime(v int64) {
	o.EndTime = v
}

// GetError returns the Error field value.
func (o *APMTraceSpan) GetError() APMSpanErrorFlag {
	if o == nil {
		var ret APMSpanErrorFlag
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetErrorOk() (*APMSpanErrorFlag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value.
func (o *APMTraceSpan) SetError(v APMSpanErrorFlag) {
	o.Error = v
}

// GetMeta returns the Meta field value.
func (o *APMTraceSpan) GetMeta() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetMetaOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *APMTraceSpan) SetMeta(v map[string]string) {
	o.Meta = v
}

// GetMetrics returns the Metrics field value.
func (o *APMTraceSpan) GetMetrics() map[string]float64 {
	if o == nil {
		var ret map[string]float64
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetMetricsOk() (*map[string]float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value.
func (o *APMTraceSpan) SetMetrics(v map[string]float64) {
	o.Metrics = v
}

// GetName returns the Name field value.
func (o *APMTraceSpan) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *APMTraceSpan) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value.
func (o *APMTraceSpan) GetParentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetParentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value.
func (o *APMTraceSpan) SetParentId(v int64) {
	o.ParentId = v
}

// GetResource returns the Resource field value.
func (o *APMTraceSpan) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value.
func (o *APMTraceSpan) SetResource(v string) {
	o.Resource = v
}

// GetResourceHash returns the ResourceHash field value if set, zero value otherwise.
func (o *APMTraceSpan) GetResourceHash() string {
	if o == nil || o.ResourceHash == nil {
		var ret string
		return ret
	}
	return *o.ResourceHash
}

// GetResourceHashOk returns a tuple with the ResourceHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetResourceHashOk() (*string, bool) {
	if o == nil || o.ResourceHash == nil {
		return nil, false
	}
	return o.ResourceHash, true
}

// HasResourceHash returns a boolean if a field has been set.
func (o *APMTraceSpan) HasResourceHash() bool {
	return o != nil && o.ResourceHash != nil
}

// SetResourceHash gets a reference to the given string and assigns it to the ResourceHash field.
func (o *APMTraceSpan) SetResourceHash(v string) {
	o.ResourceHash = &v
}

// GetRestricted returns the Restricted field value if set, zero value otherwise.
func (o *APMTraceSpan) GetRestricted() bool {
	if o == nil || o.Restricted == nil {
		var ret bool
		return ret
	}
	return *o.Restricted
}

// GetRestrictedOk returns a tuple with the Restricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetRestrictedOk() (*bool, bool) {
	if o == nil || o.Restricted == nil {
		return nil, false
	}
	return o.Restricted, true
}

// HasRestricted returns a boolean if a field has been set.
func (o *APMTraceSpan) HasRestricted() bool {
	return o != nil && o.Restricted != nil
}

// SetRestricted gets a reference to the given bool and assigns it to the Restricted field.
func (o *APMTraceSpan) SetRestricted(v bool) {
	o.Restricted = &v
}

// GetSelfTime returns the SelfTime field value if set, zero value otherwise.
func (o *APMTraceSpan) GetSelfTime() float64 {
	if o == nil || o.SelfTime == nil {
		var ret float64
		return ret
	}
	return *o.SelfTime
}

// GetSelfTimeOk returns a tuple with the SelfTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetSelfTimeOk() (*float64, bool) {
	if o == nil || o.SelfTime == nil {
		return nil, false
	}
	return o.SelfTime, true
}

// HasSelfTime returns a boolean if a field has been set.
func (o *APMTraceSpan) HasSelfTime() bool {
	return o != nil && o.SelfTime != nil
}

// SetSelfTime gets a reference to the given float64 and assigns it to the SelfTime field.
func (o *APMTraceSpan) SetSelfTime(v float64) {
	o.SelfTime = &v
}

// GetService returns the Service field value.
func (o *APMTraceSpan) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *APMTraceSpan) SetService(v string) {
	o.Service = v
}

// GetSpanId returns the SpanId field value.
func (o *APMTraceSpan) GetSpanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetSpanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanId, true
}

// SetSpanId sets field value.
func (o *APMTraceSpan) SetSpanId(v int64) {
	o.SpanId = v
}

// GetStartTime returns the StartTime field value.
func (o *APMTraceSpan) GetStartTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value.
func (o *APMTraceSpan) SetStartTime(v int64) {
	o.StartTime = v
}

// GetTraceId returns the TraceId field value.
func (o *APMTraceSpan) GetTraceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetTraceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value.
func (o *APMTraceSpan) SetTraceId(v int64) {
	o.TraceId = v
}

// GetTraceIdFull returns the TraceIdFull field value.
func (o *APMTraceSpan) GetTraceIdFull() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TraceIdFull
}

// GetTraceIdFullOk returns a tuple with the TraceIdFull field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetTraceIdFullOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceIdFull, true
}

// SetTraceIdFull sets field value.
func (o *APMTraceSpan) SetTraceIdFull(v string) {
	o.TraceIdFull = v
}

// GetType returns the Type field value.
func (o *APMTraceSpan) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *APMTraceSpan) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *APMTraceSpan) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o APMTraceSpan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["duration"] = o.Duration
	toSerialize["endTime"] = o.EndTime
	toSerialize["error"] = o.Error
	toSerialize["meta"] = o.Meta
	toSerialize["metrics"] = o.Metrics
	toSerialize["name"] = o.Name
	toSerialize["parentID"] = o.ParentId
	toSerialize["resource"] = o.Resource
	if o.ResourceHash != nil {
		toSerialize["resourceHash"] = o.ResourceHash
	}
	if o.Restricted != nil {
		toSerialize["restricted"] = o.Restricted
	}
	if o.SelfTime != nil {
		toSerialize["self_time"] = o.SelfTime
	}
	toSerialize["service"] = o.Service
	toSerialize["spanID"] = o.SpanId
	toSerialize["startTime"] = o.StartTime
	toSerialize["traceID"] = o.TraceId
	toSerialize["traceIDFull"] = o.TraceIdFull
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *APMTraceSpan) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration     *int64              `json:"duration"`
		EndTime      *int64              `json:"endTime"`
		Error        *APMSpanErrorFlag   `json:"error"`
		Meta         *map[string]string  `json:"meta"`
		Metrics      *map[string]float64 `json:"metrics"`
		Name         *string             `json:"name"`
		ParentId     *int64              `json:"parentID"`
		Resource     *string             `json:"resource"`
		ResourceHash *string             `json:"resourceHash,omitempty"`
		Restricted   *bool               `json:"restricted,omitempty"`
		SelfTime     *float64            `json:"self_time,omitempty"`
		Service      *string             `json:"service"`
		SpanId       *int64              `json:"spanID"`
		StartTime    *int64              `json:"startTime"`
		TraceId      *int64              `json:"traceID"`
		TraceIdFull  *string             `json:"traceIDFull"`
		Type         *string             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Duration == nil {
		return fmt.Errorf("required field duration missing")
	}
	if all.EndTime == nil {
		return fmt.Errorf("required field endTime missing")
	}
	if all.Error == nil {
		return fmt.Errorf("required field error missing")
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
	if all.StartTime == nil {
		return fmt.Errorf("required field startTime missing")
	}
	if all.TraceId == nil {
		return fmt.Errorf("required field traceID missing")
	}
	if all.TraceIdFull == nil {
		return fmt.Errorf("required field traceIDFull missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration", "endTime", "error", "meta", "metrics", "name", "parentID", "resource", "resourceHash", "restricted", "self_time", "service", "spanID", "startTime", "traceID", "traceIDFull", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Duration = *all.Duration
	o.EndTime = *all.EndTime
	if !all.Error.IsValid() {
		hasInvalidField = true
	} else {
		o.Error = *all.Error
	}
	o.Meta = *all.Meta
	o.Metrics = *all.Metrics
	o.Name = *all.Name
	o.ParentId = *all.ParentId
	o.Resource = *all.Resource
	o.ResourceHash = all.ResourceHash
	o.Restricted = all.Restricted
	o.SelfTime = all.SelfTime
	o.Service = *all.Service
	o.SpanId = *all.SpanId
	o.StartTime = *all.StartTime
	o.TraceId = *all.TraceId
	o.TraceIdFull = *all.TraceIdFull
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
