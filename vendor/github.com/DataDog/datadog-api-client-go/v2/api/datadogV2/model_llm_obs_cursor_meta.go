// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCursorMeta Pagination cursor metadata.
type LLMObsCursorMeta struct {
	// Cursor for the next page of results.
	After datadog.NullableString `json:"after,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCursorMeta instantiates a new LLMObsCursorMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCursorMeta() *LLMObsCursorMeta {
	this := LLMObsCursorMeta{}
	return &this
}

// NewLLMObsCursorMetaWithDefaults instantiates a new LLMObsCursorMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCursorMetaWithDefaults() *LLMObsCursorMeta {
	this := LLMObsCursorMeta{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCursorMeta) GetAfter() string {
	if o == nil || o.After.Get() == nil {
		var ret string
		return ret
	}
	return *o.After.Get()
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCursorMeta) GetAfterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.After.Get(), o.After.IsSet()
}

// HasAfter returns a boolean if a field has been set.
func (o *LLMObsCursorMeta) HasAfter() bool {
	return o != nil && o.After.IsSet()
}

// SetAfter gets a reference to the given datadog.NullableString and assigns it to the After field.
func (o *LLMObsCursorMeta) SetAfter(v string) {
	o.After.Set(&v)
}

// SetAfterNil sets the value for After to be an explicit nil.
func (o *LLMObsCursorMeta) SetAfterNil() {
	o.After.Set(nil)
}

// UnsetAfter ensures that no value is present for After, not even an explicit nil.
func (o *LLMObsCursorMeta) UnsetAfter() {
	o.After.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCursorMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.After.IsSet() {
		toSerialize["after"] = o.After.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCursorMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		After datadog.NullableString `json:"after,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"after"})
	} else {
		return err
	}
	o.After = all.After

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
