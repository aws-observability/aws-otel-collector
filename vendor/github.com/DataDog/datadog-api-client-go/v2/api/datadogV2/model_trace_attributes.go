// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TraceAttributes The attributes of a trace returned by the Get trace by ID endpoint.
type TraceAttributes struct {
	// Indicates whether the trace was truncated because its size exceeded the maximum response payload.
	IsTruncated bool `json:"is_truncated"`
	// The list of spans that compose the trace.
	Spans []APMTraceSpan `json:"spans"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTraceAttributes instantiates a new TraceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTraceAttributes(isTruncated bool, spans []APMTraceSpan) *TraceAttributes {
	this := TraceAttributes{}
	this.IsTruncated = isTruncated
	this.Spans = spans
	return &this
}

// NewTraceAttributesWithDefaults instantiates a new TraceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTraceAttributesWithDefaults() *TraceAttributes {
	this := TraceAttributes{}
	return &this
}

// GetIsTruncated returns the IsTruncated field value.
func (o *TraceAttributes) GetIsTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsTruncated
}

// GetIsTruncatedOk returns a tuple with the IsTruncated field value
// and a boolean to check if the value has been set.
func (o *TraceAttributes) GetIsTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTruncated, true
}

// SetIsTruncated sets field value.
func (o *TraceAttributes) SetIsTruncated(v bool) {
	o.IsTruncated = v
}

// GetSpans returns the Spans field value.
func (o *TraceAttributes) GetSpans() []APMTraceSpan {
	if o == nil {
		var ret []APMTraceSpan
		return ret
	}
	return o.Spans
}

// GetSpansOk returns a tuple with the Spans field value
// and a boolean to check if the value has been set.
func (o *TraceAttributes) GetSpansOk() (*[]APMTraceSpan, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spans, true
}

// SetSpans sets field value.
func (o *TraceAttributes) SetSpans(v []APMTraceSpan) {
	o.Spans = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TraceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["is_truncated"] = o.IsTruncated
	toSerialize["spans"] = o.Spans

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TraceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsTruncated *bool           `json:"is_truncated"`
		Spans       *[]APMTraceSpan `json:"spans"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsTruncated == nil {
		return fmt.Errorf("required field is_truncated missing")
	}
	if all.Spans == nil {
		return fmt.Errorf("required field spans missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_truncated", "spans"})
	} else {
		return err
	}
	o.IsTruncated = *all.IsTruncated
	o.Spans = *all.Spans

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
