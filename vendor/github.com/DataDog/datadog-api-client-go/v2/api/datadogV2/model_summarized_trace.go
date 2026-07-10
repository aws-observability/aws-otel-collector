// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SummarizedTrace A summarized, hierarchical view of a trace.
type SummarizedTrace struct {
	// A node in the pruned trace tree.
	Root SummarizedSpan `json:"root"`
	// The full 128-bit trace ID, encoded as a 32-character hexadecimal string.
	TraceId string `json:"traceId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSummarizedTrace instantiates a new SummarizedTrace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSummarizedTrace(root SummarizedSpan, traceId string) *SummarizedTrace {
	this := SummarizedTrace{}
	this.Root = root
	this.TraceId = traceId
	return &this
}

// NewSummarizedTraceWithDefaults instantiates a new SummarizedTrace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSummarizedTraceWithDefaults() *SummarizedTrace {
	this := SummarizedTrace{}
	return &this
}

// GetRoot returns the Root field value.
func (o *SummarizedTrace) GetRoot() SummarizedSpan {
	if o == nil {
		var ret SummarizedSpan
		return ret
	}
	return o.Root
}

// GetRootOk returns a tuple with the Root field value
// and a boolean to check if the value has been set.
func (o *SummarizedTrace) GetRootOk() (*SummarizedSpan, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Root, true
}

// SetRoot sets field value.
func (o *SummarizedTrace) SetRoot(v SummarizedSpan) {
	o.Root = v
}

// GetTraceId returns the TraceId field value.
func (o *SummarizedTrace) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *SummarizedTrace) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value.
func (o *SummarizedTrace) SetTraceId(v string) {
	o.TraceId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SummarizedTrace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["root"] = o.Root
	toSerialize["traceId"] = o.TraceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SummarizedTrace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Root    *SummarizedSpan `json:"root"`
		TraceId *string         `json:"traceId"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Root == nil {
		return fmt.Errorf("required field root missing")
	}
	if all.TraceId == nil {
		return fmt.Errorf("required field traceId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"root", "traceId"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Root.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Root = *all.Root
	o.TraceId = *all.TraceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
