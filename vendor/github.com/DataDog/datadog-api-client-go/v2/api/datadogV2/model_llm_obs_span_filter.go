// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSpanFilter Filter criteria for an LLM Observability span search.
type LLMObsSpanFilter struct {
	// Start of the time range. Accepts ISO 8601 or relative format (e.g., `now-15m`). Defaults to `now-15m`.
	From *string `json:"from,omitempty"`
	// Filter by ML application name.
	MlApp *string `json:"ml_app,omitempty"`
	// Search query using LLM Observability query syntax. Supports attribute filters using the field:value syntax (e.g. session_id, trace_id, ml_app, meta.span.kind). When provided, structured field filters (`span_id`, `trace_id`, etc.) are ignored.
	Query *string `json:"query,omitempty"`
	// Filter by exact span ID.
	SpanId *string `json:"span_id,omitempty"`
	// Filter by span kind (e.g., llm, agent, tool, task, workflow).
	SpanKind *string `json:"span_kind,omitempty"`
	// Filter by span name.
	SpanName *string `json:"span_name,omitempty"`
	// Filter by tag key-value pairs.
	Tags map[string]string `json:"tags,omitempty"`
	// End of the time range. Accepts ISO 8601 or relative format (e.g., `now`). Defaults to `now`.
	To *string `json:"to,omitempty"`
	// Filter by exact trace ID.
	TraceId *string `json:"trace_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsSpanFilter instantiates a new LLMObsSpanFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsSpanFilter() *LLMObsSpanFilter {
	this := LLMObsSpanFilter{}
	return &this
}

// NewLLMObsSpanFilterWithDefaults instantiates a new LLMObsSpanFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsSpanFilterWithDefaults() *LLMObsSpanFilter {
	this := LLMObsSpanFilter{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *LLMObsSpanFilter) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanFilter) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *LLMObsSpanFilter) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *LLMObsSpanFilter) SetFrom(v string) {
	o.From = &v
}

// GetMlApp returns the MlApp field value if set, zero value otherwise.
func (o *LLMObsSpanFilter) GetMlApp() string {
	if o == nil || o.MlApp == nil {
		var ret string
		return ret
	}
	return *o.MlApp
}

// GetMlAppOk returns a tuple with the MlApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanFilter) GetMlAppOk() (*string, bool) {
	if o == nil || o.MlApp == nil {
		return nil, false
	}
	return o.MlApp, true
}

// HasMlApp returns a boolean if a field has been set.
func (o *LLMObsSpanFilter) HasMlApp() bool {
	return o != nil && o.MlApp != nil
}

// SetMlApp gets a reference to the given string and assigns it to the MlApp field.
func (o *LLMObsSpanFilter) SetMlApp(v string) {
	o.MlApp = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *LLMObsSpanFilter) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanFilter) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *LLMObsSpanFilter) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *LLMObsSpanFilter) SetQuery(v string) {
	o.Query = &v
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *LLMObsSpanFilter) GetSpanId() string {
	if o == nil || o.SpanId == nil {
		var ret string
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanFilter) GetSpanIdOk() (*string, bool) {
	if o == nil || o.SpanId == nil {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *LLMObsSpanFilter) HasSpanId() bool {
	return o != nil && o.SpanId != nil
}

// SetSpanId gets a reference to the given string and assigns it to the SpanId field.
func (o *LLMObsSpanFilter) SetSpanId(v string) {
	o.SpanId = &v
}

// GetSpanKind returns the SpanKind field value if set, zero value otherwise.
func (o *LLMObsSpanFilter) GetSpanKind() string {
	if o == nil || o.SpanKind == nil {
		var ret string
		return ret
	}
	return *o.SpanKind
}

// GetSpanKindOk returns a tuple with the SpanKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanFilter) GetSpanKindOk() (*string, bool) {
	if o == nil || o.SpanKind == nil {
		return nil, false
	}
	return o.SpanKind, true
}

// HasSpanKind returns a boolean if a field has been set.
func (o *LLMObsSpanFilter) HasSpanKind() bool {
	return o != nil && o.SpanKind != nil
}

// SetSpanKind gets a reference to the given string and assigns it to the SpanKind field.
func (o *LLMObsSpanFilter) SetSpanKind(v string) {
	o.SpanKind = &v
}

// GetSpanName returns the SpanName field value if set, zero value otherwise.
func (o *LLMObsSpanFilter) GetSpanName() string {
	if o == nil || o.SpanName == nil {
		var ret string
		return ret
	}
	return *o.SpanName
}

// GetSpanNameOk returns a tuple with the SpanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanFilter) GetSpanNameOk() (*string, bool) {
	if o == nil || o.SpanName == nil {
		return nil, false
	}
	return o.SpanName, true
}

// HasSpanName returns a boolean if a field has been set.
func (o *LLMObsSpanFilter) HasSpanName() bool {
	return o != nil && o.SpanName != nil
}

// SetSpanName gets a reference to the given string and assigns it to the SpanName field.
func (o *LLMObsSpanFilter) SetSpanName(v string) {
	o.SpanName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LLMObsSpanFilter) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanFilter) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LLMObsSpanFilter) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *LLMObsSpanFilter) SetTags(v map[string]string) {
	o.Tags = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *LLMObsSpanFilter) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanFilter) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *LLMObsSpanFilter) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *LLMObsSpanFilter) SetTo(v string) {
	o.To = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *LLMObsSpanFilter) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanFilter) GetTraceIdOk() (*string, bool) {
	if o == nil || o.TraceId == nil {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *LLMObsSpanFilter) HasTraceId() bool {
	return o != nil && o.TraceId != nil
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *LLMObsSpanFilter) SetTraceId(v string) {
	o.TraceId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsSpanFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.MlApp != nil {
		toSerialize["ml_app"] = o.MlApp
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.SpanId != nil {
		toSerialize["span_id"] = o.SpanId
	}
	if o.SpanKind != nil {
		toSerialize["span_kind"] = o.SpanKind
	}
	if o.SpanName != nil {
		toSerialize["span_name"] = o.SpanName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.To != nil {
		toSerialize["to"] = o.To
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
func (o *LLMObsSpanFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From     *string           `json:"from,omitempty"`
		MlApp    *string           `json:"ml_app,omitempty"`
		Query    *string           `json:"query,omitempty"`
		SpanId   *string           `json:"span_id,omitempty"`
		SpanKind *string           `json:"span_kind,omitempty"`
		SpanName *string           `json:"span_name,omitempty"`
		Tags     map[string]string `json:"tags,omitempty"`
		To       *string           `json:"to,omitempty"`
		TraceId  *string           `json:"trace_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "ml_app", "query", "span_id", "span_kind", "span_name", "tags", "to", "trace_id"})
	} else {
		return err
	}
	o.From = all.From
	o.MlApp = all.MlApp
	o.Query = all.Query
	o.SpanId = all.SpanId
	o.SpanKind = all.SpanKind
	o.SpanName = all.SpanName
	o.Tags = all.Tags
	o.To = all.To
	o.TraceId = all.TraceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
