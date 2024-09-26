// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAttributes JSON object containing all span attributes and their associated values.
type SpansAttributes struct {
	// JSON object of attributes from your span.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// JSON object of custom spans data.
	Custom map[string]interface{} `json:"custom,omitempty"`
	// End timestamp of your span.
	EndTimestamp *time.Time `json:"end_timestamp,omitempty"`
	// Name of the environment from where the spans are being sent.
	Env *string `json:"env,omitempty"`
	// Name of the machine from where the spans are being sent.
	Host *string `json:"host,omitempty"`
	// The reason why the span was ingested.
	IngestionReason *string `json:"ingestion_reason,omitempty"`
	// Id of the span that's parent of this span.
	ParentId *string `json:"parent_id,omitempty"`
	// Unique identifier of the resource.
	ResourceHash *string `json:"resource_hash,omitempty"`
	// The name of the resource.
	ResourceName *string `json:"resource_name,omitempty"`
	// The reason why the span was indexed.
	RetainedBy *string `json:"retained_by,omitempty"`
	// The name of the application or service generating the span events.
	// It is used to switch from APM to Logs, so make sure you define the same
	// value when you use both products.
	Service *string `json:"service,omitempty"`
	// Whether or not the span was collected as a stand-alone span. Always associated to "single_span" ingestion_reason if true.
	SingleSpan *bool `json:"single_span,omitempty"`
	// Id of the span.
	SpanId *string `json:"span_id,omitempty"`
	// Start timestamp of your span.
	StartTimestamp *time.Time `json:"start_timestamp,omitempty"`
	// Array of tags associated with your span.
	Tags []string `json:"tags,omitempty"`
	// Id of the trace to which the span belongs.
	TraceId *string `json:"trace_id,omitempty"`
	// The type of the span.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpansAttributes instantiates a new SpansAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansAttributes() *SpansAttributes {
	this := SpansAttributes{}
	return &this
}

// NewSpansAttributesWithDefaults instantiates a new SpansAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansAttributesWithDefaults() *SpansAttributes {
	this := SpansAttributes{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SpansAttributes) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SpansAttributes) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *SpansAttributes) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *SpansAttributes) GetCustom() map[string]interface{} {
	if o == nil || o.Custom == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetCustomOk() (*map[string]interface{}, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return &o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *SpansAttributes) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given map[string]interface{} and assigns it to the Custom field.
func (o *SpansAttributes) SetCustom(v map[string]interface{}) {
	o.Custom = v
}

// GetEndTimestamp returns the EndTimestamp field value if set, zero value otherwise.
func (o *SpansAttributes) GetEndTimestamp() time.Time {
	if o == nil || o.EndTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTimestamp
}

// GetEndTimestampOk returns a tuple with the EndTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetEndTimestampOk() (*time.Time, bool) {
	if o == nil || o.EndTimestamp == nil {
		return nil, false
	}
	return o.EndTimestamp, true
}

// HasEndTimestamp returns a boolean if a field has been set.
func (o *SpansAttributes) HasEndTimestamp() bool {
	return o != nil && o.EndTimestamp != nil
}

// SetEndTimestamp gets a reference to the given time.Time and assigns it to the EndTimestamp field.
func (o *SpansAttributes) SetEndTimestamp(v time.Time) {
	o.EndTimestamp = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *SpansAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *SpansAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *SpansAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *SpansAttributes) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *SpansAttributes) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *SpansAttributes) SetHost(v string) {
	o.Host = &v
}

// GetIngestionReason returns the IngestionReason field value if set, zero value otherwise.
func (o *SpansAttributes) GetIngestionReason() string {
	if o == nil || o.IngestionReason == nil {
		var ret string
		return ret
	}
	return *o.IngestionReason
}

// GetIngestionReasonOk returns a tuple with the IngestionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetIngestionReasonOk() (*string, bool) {
	if o == nil || o.IngestionReason == nil {
		return nil, false
	}
	return o.IngestionReason, true
}

// HasIngestionReason returns a boolean if a field has been set.
func (o *SpansAttributes) HasIngestionReason() bool {
	return o != nil && o.IngestionReason != nil
}

// SetIngestionReason gets a reference to the given string and assigns it to the IngestionReason field.
func (o *SpansAttributes) SetIngestionReason(v string) {
	o.IngestionReason = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *SpansAttributes) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *SpansAttributes) HasParentId() bool {
	return o != nil && o.ParentId != nil
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *SpansAttributes) SetParentId(v string) {
	o.ParentId = &v
}

// GetResourceHash returns the ResourceHash field value if set, zero value otherwise.
func (o *SpansAttributes) GetResourceHash() string {
	if o == nil || o.ResourceHash == nil {
		var ret string
		return ret
	}
	return *o.ResourceHash
}

// GetResourceHashOk returns a tuple with the ResourceHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetResourceHashOk() (*string, bool) {
	if o == nil || o.ResourceHash == nil {
		return nil, false
	}
	return o.ResourceHash, true
}

// HasResourceHash returns a boolean if a field has been set.
func (o *SpansAttributes) HasResourceHash() bool {
	return o != nil && o.ResourceHash != nil
}

// SetResourceHash gets a reference to the given string and assigns it to the ResourceHash field.
func (o *SpansAttributes) SetResourceHash(v string) {
	o.ResourceHash = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *SpansAttributes) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *SpansAttributes) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *SpansAttributes) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetRetainedBy returns the RetainedBy field value if set, zero value otherwise.
func (o *SpansAttributes) GetRetainedBy() string {
	if o == nil || o.RetainedBy == nil {
		var ret string
		return ret
	}
	return *o.RetainedBy
}

// GetRetainedByOk returns a tuple with the RetainedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetRetainedByOk() (*string, bool) {
	if o == nil || o.RetainedBy == nil {
		return nil, false
	}
	return o.RetainedBy, true
}

// HasRetainedBy returns a boolean if a field has been set.
func (o *SpansAttributes) HasRetainedBy() bool {
	return o != nil && o.RetainedBy != nil
}

// SetRetainedBy gets a reference to the given string and assigns it to the RetainedBy field.
func (o *SpansAttributes) SetRetainedBy(v string) {
	o.RetainedBy = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SpansAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SpansAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SpansAttributes) SetService(v string) {
	o.Service = &v
}

// GetSingleSpan returns the SingleSpan field value if set, zero value otherwise.
func (o *SpansAttributes) GetSingleSpan() bool {
	if o == nil || o.SingleSpan == nil {
		var ret bool
		return ret
	}
	return *o.SingleSpan
}

// GetSingleSpanOk returns a tuple with the SingleSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetSingleSpanOk() (*bool, bool) {
	if o == nil || o.SingleSpan == nil {
		return nil, false
	}
	return o.SingleSpan, true
}

// HasSingleSpan returns a boolean if a field has been set.
func (o *SpansAttributes) HasSingleSpan() bool {
	return o != nil && o.SingleSpan != nil
}

// SetSingleSpan gets a reference to the given bool and assigns it to the SingleSpan field.
func (o *SpansAttributes) SetSingleSpan(v bool) {
	o.SingleSpan = &v
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *SpansAttributes) GetSpanId() string {
	if o == nil || o.SpanId == nil {
		var ret string
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetSpanIdOk() (*string, bool) {
	if o == nil || o.SpanId == nil {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *SpansAttributes) HasSpanId() bool {
	return o != nil && o.SpanId != nil
}

// SetSpanId gets a reference to the given string and assigns it to the SpanId field.
func (o *SpansAttributes) SetSpanId(v string) {
	o.SpanId = &v
}

// GetStartTimestamp returns the StartTimestamp field value if set, zero value otherwise.
func (o *SpansAttributes) GetStartTimestamp() time.Time {
	if o == nil || o.StartTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetStartTimestampOk() (*time.Time, bool) {
	if o == nil || o.StartTimestamp == nil {
		return nil, false
	}
	return o.StartTimestamp, true
}

// HasStartTimestamp returns a boolean if a field has been set.
func (o *SpansAttributes) HasStartTimestamp() bool {
	return o != nil && o.StartTimestamp != nil
}

// SetStartTimestamp gets a reference to the given time.Time and assigns it to the StartTimestamp field.
func (o *SpansAttributes) SetStartTimestamp(v time.Time) {
	o.StartTimestamp = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SpansAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SpansAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SpansAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *SpansAttributes) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetTraceIdOk() (*string, bool) {
	if o == nil || o.TraceId == nil {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *SpansAttributes) HasTraceId() bool {
	return o != nil && o.TraceId != nil
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *SpansAttributes) SetTraceId(v string) {
	o.TraceId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SpansAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SpansAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SpansAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	if o.EndTimestamp != nil {
		if o.EndTimestamp.Nanosecond() == 0 {
			toSerialize["end_timestamp"] = o.EndTimestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["end_timestamp"] = o.EndTimestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.IngestionReason != nil {
		toSerialize["ingestion_reason"] = o.IngestionReason
	}
	if o.ParentId != nil {
		toSerialize["parent_id"] = o.ParentId
	}
	if o.ResourceHash != nil {
		toSerialize["resource_hash"] = o.ResourceHash
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.RetainedBy != nil {
		toSerialize["retained_by"] = o.RetainedBy
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.SingleSpan != nil {
		toSerialize["single_span"] = o.SingleSpan
	}
	if o.SpanId != nil {
		toSerialize["span_id"] = o.SpanId
	}
	if o.StartTimestamp != nil {
		if o.StartTimestamp.Nanosecond() == 0 {
			toSerialize["start_timestamp"] = o.StartTimestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["start_timestamp"] = o.StartTimestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TraceId != nil {
		toSerialize["trace_id"] = o.TraceId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes      map[string]interface{} `json:"attributes,omitempty"`
		Custom          map[string]interface{} `json:"custom,omitempty"`
		EndTimestamp    *time.Time             `json:"end_timestamp,omitempty"`
		Env             *string                `json:"env,omitempty"`
		Host            *string                `json:"host,omitempty"`
		IngestionReason *string                `json:"ingestion_reason,omitempty"`
		ParentId        *string                `json:"parent_id,omitempty"`
		ResourceHash    *string                `json:"resource_hash,omitempty"`
		ResourceName    *string                `json:"resource_name,omitempty"`
		RetainedBy      *string                `json:"retained_by,omitempty"`
		Service         *string                `json:"service,omitempty"`
		SingleSpan      *bool                  `json:"single_span,omitempty"`
		SpanId          *string                `json:"span_id,omitempty"`
		StartTimestamp  *time.Time             `json:"start_timestamp,omitempty"`
		Tags            []string               `json:"tags,omitempty"`
		TraceId         *string                `json:"trace_id,omitempty"`
		Type            *string                `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "custom", "end_timestamp", "env", "host", "ingestion_reason", "parent_id", "resource_hash", "resource_name", "retained_by", "service", "single_span", "span_id", "start_timestamp", "tags", "trace_id", "type"})
	} else {
		return err
	}
	o.Attributes = all.Attributes
	o.Custom = all.Custom
	o.EndTimestamp = all.EndTimestamp
	o.Env = all.Env
	o.Host = all.Host
	o.IngestionReason = all.IngestionReason
	o.ParentId = all.ParentId
	o.ResourceHash = all.ResourceHash
	o.ResourceName = all.ResourceName
	o.RetainedBy = all.RetainedBy
	o.Service = all.Service
	o.SingleSpan = all.SingleSpan
	o.SpanId = all.SpanId
	o.StartTimestamp = all.StartTimestamp
	o.Tags = all.Tags
	o.TraceId = all.TraceId
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
