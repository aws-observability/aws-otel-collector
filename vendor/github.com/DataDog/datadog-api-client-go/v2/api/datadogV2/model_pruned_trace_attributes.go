// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PrunedTraceAttributes The attributes of a pruned trace returned by the Get pruned trace by ID endpoint.
type PrunedTraceAttributes struct {
	// Indicates whether the underlying trace was truncated because its size
	// exceeded the maximum that can be retrieved from storage.
	IsTruncated bool `json:"is_truncated"`
	// The size, in bytes, of the original (non-pruned) trace before summarization.
	SizeBytes int32 `json:"size_bytes"`
	// A summarized, hierarchical view of a trace.
	SummarizedTrace SummarizedTrace `json:"summarized_trace"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPrunedTraceAttributes instantiates a new PrunedTraceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPrunedTraceAttributes(isTruncated bool, sizeBytes int32, summarizedTrace SummarizedTrace) *PrunedTraceAttributes {
	this := PrunedTraceAttributes{}
	this.IsTruncated = isTruncated
	this.SizeBytes = sizeBytes
	this.SummarizedTrace = summarizedTrace
	return &this
}

// NewPrunedTraceAttributesWithDefaults instantiates a new PrunedTraceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPrunedTraceAttributesWithDefaults() *PrunedTraceAttributes {
	this := PrunedTraceAttributes{}
	return &this
}

// GetIsTruncated returns the IsTruncated field value.
func (o *PrunedTraceAttributes) GetIsTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsTruncated
}

// GetIsTruncatedOk returns a tuple with the IsTruncated field value
// and a boolean to check if the value has been set.
func (o *PrunedTraceAttributes) GetIsTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTruncated, true
}

// SetIsTruncated sets field value.
func (o *PrunedTraceAttributes) SetIsTruncated(v bool) {
	o.IsTruncated = v
}

// GetSizeBytes returns the SizeBytes field value.
func (o *PrunedTraceAttributes) GetSizeBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value
// and a boolean to check if the value has been set.
func (o *PrunedTraceAttributes) GetSizeBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeBytes, true
}

// SetSizeBytes sets field value.
func (o *PrunedTraceAttributes) SetSizeBytes(v int32) {
	o.SizeBytes = v
}

// GetSummarizedTrace returns the SummarizedTrace field value.
func (o *PrunedTraceAttributes) GetSummarizedTrace() SummarizedTrace {
	if o == nil {
		var ret SummarizedTrace
		return ret
	}
	return o.SummarizedTrace
}

// GetSummarizedTraceOk returns a tuple with the SummarizedTrace field value
// and a boolean to check if the value has been set.
func (o *PrunedTraceAttributes) GetSummarizedTraceOk() (*SummarizedTrace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SummarizedTrace, true
}

// SetSummarizedTrace sets field value.
func (o *PrunedTraceAttributes) SetSummarizedTrace(v SummarizedTrace) {
	o.SummarizedTrace = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PrunedTraceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["is_truncated"] = o.IsTruncated
	toSerialize["size_bytes"] = o.SizeBytes
	toSerialize["summarized_trace"] = o.SummarizedTrace

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PrunedTraceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsTruncated     *bool            `json:"is_truncated"`
		SizeBytes       *int32           `json:"size_bytes"`
		SummarizedTrace *SummarizedTrace `json:"summarized_trace"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsTruncated == nil {
		return fmt.Errorf("required field is_truncated missing")
	}
	if all.SizeBytes == nil {
		return fmt.Errorf("required field size_bytes missing")
	}
	if all.SummarizedTrace == nil {
		return fmt.Errorf("required field summarized_trace missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_truncated", "size_bytes", "summarized_trace"})
	} else {
		return err
	}

	hasInvalidField := false
	o.IsTruncated = *all.IsTruncated
	o.SizeBytes = *all.SizeBytes
	if all.SummarizedTrace.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SummarizedTrace = *all.SummarizedTrace

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
