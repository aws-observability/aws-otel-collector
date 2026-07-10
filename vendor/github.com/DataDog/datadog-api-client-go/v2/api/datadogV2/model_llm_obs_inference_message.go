// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsInferenceMessage A single message in an LLM inference conversation.
type LLMObsInferenceMessage struct {
	// Plain text content of the message.
	Content *string `json:"content,omitempty"`
	// List of structured content blocks in a message.
	Contents []LLMObsInferenceContent `json:"contents,omitempty"`
	// Unique identifier for the message.
	Id *string `json:"id,omitempty"`
	// The role of the message author.
	Role *string `json:"role,omitempty"`
	// List of tool calls in a message.
	ToolCalls []LLMObsInferenceToolCall `json:"tool_calls,omitempty"`
	// List of tool results in a message.
	ToolResults []LLMObsInferenceToolResult `json:"tool_results,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsInferenceMessage instantiates a new LLMObsInferenceMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsInferenceMessage() *LLMObsInferenceMessage {
	this := LLMObsInferenceMessage{}
	return &this
}

// NewLLMObsInferenceMessageWithDefaults instantiates a new LLMObsInferenceMessage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsInferenceMessageWithDefaults() *LLMObsInferenceMessage {
	this := LLMObsInferenceMessage{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *LLMObsInferenceMessage) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceMessage) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *LLMObsInferenceMessage) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *LLMObsInferenceMessage) SetContent(v string) {
	o.Content = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *LLMObsInferenceMessage) GetContents() []LLMObsInferenceContent {
	if o == nil || o.Contents == nil {
		var ret []LLMObsInferenceContent
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceMessage) GetContentsOk() (*[]LLMObsInferenceContent, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return &o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *LLMObsInferenceMessage) HasContents() bool {
	return o != nil && o.Contents != nil
}

// SetContents gets a reference to the given []LLMObsInferenceContent and assigns it to the Contents field.
func (o *LLMObsInferenceMessage) SetContents(v []LLMObsInferenceContent) {
	o.Contents = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LLMObsInferenceMessage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceMessage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LLMObsInferenceMessage) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LLMObsInferenceMessage) SetId(v string) {
	o.Id = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *LLMObsInferenceMessage) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceMessage) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *LLMObsInferenceMessage) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *LLMObsInferenceMessage) SetRole(v string) {
	o.Role = &v
}

// GetToolCalls returns the ToolCalls field value if set, zero value otherwise.
func (o *LLMObsInferenceMessage) GetToolCalls() []LLMObsInferenceToolCall {
	if o == nil || o.ToolCalls == nil {
		var ret []LLMObsInferenceToolCall
		return ret
	}
	return o.ToolCalls
}

// GetToolCallsOk returns a tuple with the ToolCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceMessage) GetToolCallsOk() (*[]LLMObsInferenceToolCall, bool) {
	if o == nil || o.ToolCalls == nil {
		return nil, false
	}
	return &o.ToolCalls, true
}

// HasToolCalls returns a boolean if a field has been set.
func (o *LLMObsInferenceMessage) HasToolCalls() bool {
	return o != nil && o.ToolCalls != nil
}

// SetToolCalls gets a reference to the given []LLMObsInferenceToolCall and assigns it to the ToolCalls field.
func (o *LLMObsInferenceMessage) SetToolCalls(v []LLMObsInferenceToolCall) {
	o.ToolCalls = v
}

// GetToolResults returns the ToolResults field value if set, zero value otherwise.
func (o *LLMObsInferenceMessage) GetToolResults() []LLMObsInferenceToolResult {
	if o == nil || o.ToolResults == nil {
		var ret []LLMObsInferenceToolResult
		return ret
	}
	return o.ToolResults
}

// GetToolResultsOk returns a tuple with the ToolResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceMessage) GetToolResultsOk() (*[]LLMObsInferenceToolResult, bool) {
	if o == nil || o.ToolResults == nil {
		return nil, false
	}
	return &o.ToolResults, true
}

// HasToolResults returns a boolean if a field has been set.
func (o *LLMObsInferenceMessage) HasToolResults() bool {
	return o != nil && o.ToolResults != nil
}

// SetToolResults gets a reference to the given []LLMObsInferenceToolResult and assigns it to the ToolResults field.
func (o *LLMObsInferenceMessage) SetToolResults(v []LLMObsInferenceToolResult) {
	o.ToolResults = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsInferenceMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.ToolCalls != nil {
		toSerialize["tool_calls"] = o.ToolCalls
	}
	if o.ToolResults != nil {
		toSerialize["tool_results"] = o.ToolResults
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsInferenceMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content     *string                     `json:"content,omitempty"`
		Contents    []LLMObsInferenceContent    `json:"contents,omitempty"`
		Id          *string                     `json:"id,omitempty"`
		Role        *string                     `json:"role,omitempty"`
		ToolCalls   []LLMObsInferenceToolCall   `json:"tool_calls,omitempty"`
		ToolResults []LLMObsInferenceToolResult `json:"tool_results,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "contents", "id", "role", "tool_calls", "tool_results"})
	} else {
		return err
	}
	o.Content = all.Content
	o.Contents = all.Contents
	o.Id = all.Id
	o.Role = all.Role
	o.ToolCalls = all.ToolCalls
	o.ToolResults = all.ToolResults

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
