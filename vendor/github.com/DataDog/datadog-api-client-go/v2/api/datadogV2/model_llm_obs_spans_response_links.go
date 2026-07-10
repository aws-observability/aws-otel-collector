// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSpansResponseLinks Pagination links accompanying the spans response.
type LLMObsSpansResponseLinks struct {
	// URL to retrieve the next page of results.
	Next *string `json:"next,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsSpansResponseLinks instantiates a new LLMObsSpansResponseLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsSpansResponseLinks() *LLMObsSpansResponseLinks {
	this := LLMObsSpansResponseLinks{}
	return &this
}

// NewLLMObsSpansResponseLinksWithDefaults instantiates a new LLMObsSpansResponseLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsSpansResponseLinksWithDefaults() *LLMObsSpansResponseLinks {
	this := LLMObsSpansResponseLinks{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *LLMObsSpansResponseLinks) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpansResponseLinks) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *LLMObsSpansResponseLinks) HasNext() bool {
	return o != nil && o.Next != nil
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *LLMObsSpansResponseLinks) SetNext(v string) {
	o.Next = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsSpansResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsSpansResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Next *string `json:"next,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next"})
	} else {
		return err
	}
	o.Next = all.Next

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
