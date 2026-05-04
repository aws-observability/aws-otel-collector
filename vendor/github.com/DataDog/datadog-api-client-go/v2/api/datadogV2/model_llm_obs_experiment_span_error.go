// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentSpanError Error details for an experiment span.
type LLMObsExperimentSpanError struct {
	// Error message.
	Message *string `json:"message,omitempty"`
	// Stack trace of the error.
	Stack *string `json:"stack,omitempty"`
	// The error type or exception class name.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentSpanError instantiates a new LLMObsExperimentSpanError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentSpanError() *LLMObsExperimentSpanError {
	this := LLMObsExperimentSpanError{}
	return &this
}

// NewLLMObsExperimentSpanErrorWithDefaults instantiates a new LLMObsExperimentSpanError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentSpanErrorWithDefaults() *LLMObsExperimentSpanError {
	this := LLMObsExperimentSpanError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanError) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LLMObsExperimentSpanError) SetMessage(v string) {
	o.Message = &v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanError) GetStack() string {
	if o == nil || o.Stack == nil {
		var ret string
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanError) GetStackOk() (*string, bool) {
	if o == nil || o.Stack == nil {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanError) HasStack() bool {
	return o != nil && o.Stack != nil
}

// SetStack gets a reference to the given string and assigns it to the Stack field.
func (o *LLMObsExperimentSpanError) SetStack(v string) {
	o.Stack = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LLMObsExperimentSpanError) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentSpanError) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LLMObsExperimentSpanError) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LLMObsExperimentSpanError) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentSpanError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Stack != nil {
		toSerialize["stack"] = o.Stack
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
func (o *LLMObsExperimentSpanError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message *string `json:"message,omitempty"`
		Stack   *string `json:"stack,omitempty"`
		Type    *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message", "stack", "type"})
	} else {
		return err
	}
	o.Message = all.Message
	o.Stack = all.Stack
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
