// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigPromptContentValue Value of a prompt message content block.
type LLMObsCustomEvalConfigPromptContentValue struct {
	// Text content of the message block.
	Text *string `json:"text,omitempty"`
	// A tool call within a prompt message.
	ToolCall *LLMObsCustomEvalConfigPromptToolCall `json:"tool_call,omitempty"`
	// A tool call result within a prompt message.
	ToolCallResult *LLMObsCustomEvalConfigPromptToolResult `json:"tool_call_result,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCustomEvalConfigPromptContentValue instantiates a new LLMObsCustomEvalConfigPromptContentValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCustomEvalConfigPromptContentValue() *LLMObsCustomEvalConfigPromptContentValue {
	this := LLMObsCustomEvalConfigPromptContentValue{}
	return &this
}

// NewLLMObsCustomEvalConfigPromptContentValueWithDefaults instantiates a new LLMObsCustomEvalConfigPromptContentValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCustomEvalConfigPromptContentValueWithDefaults() *LLMObsCustomEvalConfigPromptContentValue {
	this := LLMObsCustomEvalConfigPromptContentValue{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptContentValue) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptContentValue) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptContentValue) HasText() bool {
	return o != nil && o.Text != nil
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *LLMObsCustomEvalConfigPromptContentValue) SetText(v string) {
	o.Text = &v
}

// GetToolCall returns the ToolCall field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptContentValue) GetToolCall() LLMObsCustomEvalConfigPromptToolCall {
	if o == nil || o.ToolCall == nil {
		var ret LLMObsCustomEvalConfigPromptToolCall
		return ret
	}
	return *o.ToolCall
}

// GetToolCallOk returns a tuple with the ToolCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptContentValue) GetToolCallOk() (*LLMObsCustomEvalConfigPromptToolCall, bool) {
	if o == nil || o.ToolCall == nil {
		return nil, false
	}
	return o.ToolCall, true
}

// HasToolCall returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptContentValue) HasToolCall() bool {
	return o != nil && o.ToolCall != nil
}

// SetToolCall gets a reference to the given LLMObsCustomEvalConfigPromptToolCall and assigns it to the ToolCall field.
func (o *LLMObsCustomEvalConfigPromptContentValue) SetToolCall(v LLMObsCustomEvalConfigPromptToolCall) {
	o.ToolCall = &v
}

// GetToolCallResult returns the ToolCallResult field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptContentValue) GetToolCallResult() LLMObsCustomEvalConfigPromptToolResult {
	if o == nil || o.ToolCallResult == nil {
		var ret LLMObsCustomEvalConfigPromptToolResult
		return ret
	}
	return *o.ToolCallResult
}

// GetToolCallResultOk returns a tuple with the ToolCallResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptContentValue) GetToolCallResultOk() (*LLMObsCustomEvalConfigPromptToolResult, bool) {
	if o == nil || o.ToolCallResult == nil {
		return nil, false
	}
	return o.ToolCallResult, true
}

// HasToolCallResult returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptContentValue) HasToolCallResult() bool {
	return o != nil && o.ToolCallResult != nil
}

// SetToolCallResult gets a reference to the given LLMObsCustomEvalConfigPromptToolResult and assigns it to the ToolCallResult field.
func (o *LLMObsCustomEvalConfigPromptContentValue) SetToolCallResult(v LLMObsCustomEvalConfigPromptToolResult) {
	o.ToolCallResult = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCustomEvalConfigPromptContentValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.ToolCall != nil {
		toSerialize["tool_call"] = o.ToolCall
	}
	if o.ToolCallResult != nil {
		toSerialize["tool_call_result"] = o.ToolCallResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCustomEvalConfigPromptContentValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Text           *string                                 `json:"text,omitempty"`
		ToolCall       *LLMObsCustomEvalConfigPromptToolCall   `json:"tool_call,omitempty"`
		ToolCallResult *LLMObsCustomEvalConfigPromptToolResult `json:"tool_call_result,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"text", "tool_call", "tool_call_result"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Text = all.Text
	if all.ToolCall != nil && all.ToolCall.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ToolCall = all.ToolCall
	if all.ToolCallResult != nil && all.ToolCallResult.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ToolCallResult = all.ToolCallResult

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
