// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSpanIO Input or output content of an LLM Observability span.
type LLMObsSpanIO struct {
	// List of messages in the input or output.
	Messages []LLMObsSpanMessage `json:"messages,omitempty"`
	// Plain-text value of the input or output.
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsSpanIO instantiates a new LLMObsSpanIO object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsSpanIO() *LLMObsSpanIO {
	this := LLMObsSpanIO{}
	return &this
}

// NewLLMObsSpanIOWithDefaults instantiates a new LLMObsSpanIO object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsSpanIOWithDefaults() *LLMObsSpanIO {
	this := LLMObsSpanIO{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *LLMObsSpanIO) GetMessages() []LLMObsSpanMessage {
	if o == nil || o.Messages == nil {
		var ret []LLMObsSpanMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanIO) GetMessagesOk() (*[]LLMObsSpanMessage, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return &o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *LLMObsSpanIO) HasMessages() bool {
	return o != nil && o.Messages != nil
}

// SetMessages gets a reference to the given []LLMObsSpanMessage and assigns it to the Messages field.
func (o *LLMObsSpanIO) SetMessages(v []LLMObsSpanMessage) {
	o.Messages = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LLMObsSpanIO) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanIO) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LLMObsSpanIO) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *LLMObsSpanIO) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsSpanIO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsSpanIO) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Messages []LLMObsSpanMessage `json:"messages,omitempty"`
		Value    *string             `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"messages", "value"})
	} else {
		return err
	}
	o.Messages = all.Messages
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
