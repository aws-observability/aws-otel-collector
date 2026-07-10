// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSpanSearchOptions Additional options for a span search request.
type LLMObsSpanSearchOptions struct {
	// Whether to include attachment data in the response. Defaults to `true`.
	IncludeAttachments *bool `json:"include_attachments,omitempty"`
	// Offset in seconds applied to both `from` and `to` timestamps.
	TimeOffset *int64 `json:"time_offset,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsSpanSearchOptions instantiates a new LLMObsSpanSearchOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsSpanSearchOptions() *LLMObsSpanSearchOptions {
	this := LLMObsSpanSearchOptions{}
	return &this
}

// NewLLMObsSpanSearchOptionsWithDefaults instantiates a new LLMObsSpanSearchOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsSpanSearchOptionsWithDefaults() *LLMObsSpanSearchOptions {
	this := LLMObsSpanSearchOptions{}
	return &this
}

// GetIncludeAttachments returns the IncludeAttachments field value if set, zero value otherwise.
func (o *LLMObsSpanSearchOptions) GetIncludeAttachments() bool {
	if o == nil || o.IncludeAttachments == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAttachments
}

// GetIncludeAttachmentsOk returns a tuple with the IncludeAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanSearchOptions) GetIncludeAttachmentsOk() (*bool, bool) {
	if o == nil || o.IncludeAttachments == nil {
		return nil, false
	}
	return o.IncludeAttachments, true
}

// HasIncludeAttachments returns a boolean if a field has been set.
func (o *LLMObsSpanSearchOptions) HasIncludeAttachments() bool {
	return o != nil && o.IncludeAttachments != nil
}

// SetIncludeAttachments gets a reference to the given bool and assigns it to the IncludeAttachments field.
func (o *LLMObsSpanSearchOptions) SetIncludeAttachments(v bool) {
	o.IncludeAttachments = &v
}

// GetTimeOffset returns the TimeOffset field value if set, zero value otherwise.
func (o *LLMObsSpanSearchOptions) GetTimeOffset() int64 {
	if o == nil || o.TimeOffset == nil {
		var ret int64
		return ret
	}
	return *o.TimeOffset
}

// GetTimeOffsetOk returns a tuple with the TimeOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanSearchOptions) GetTimeOffsetOk() (*int64, bool) {
	if o == nil || o.TimeOffset == nil {
		return nil, false
	}
	return o.TimeOffset, true
}

// HasTimeOffset returns a boolean if a field has been set.
func (o *LLMObsSpanSearchOptions) HasTimeOffset() bool {
	return o != nil && o.TimeOffset != nil
}

// SetTimeOffset gets a reference to the given int64 and assigns it to the TimeOffset field.
func (o *LLMObsSpanSearchOptions) SetTimeOffset(v int64) {
	o.TimeOffset = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsSpanSearchOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncludeAttachments != nil {
		toSerialize["include_attachments"] = o.IncludeAttachments
	}
	if o.TimeOffset != nil {
		toSerialize["time_offset"] = o.TimeOffset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsSpanSearchOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeAttachments *bool  `json:"include_attachments,omitempty"`
		TimeOffset         *int64 `json:"time_offset,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_attachments", "time_offset"})
	} else {
		return err
	}
	o.IncludeAttachments = all.IncludeAttachments
	o.TimeOffset = all.TimeOffset

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
