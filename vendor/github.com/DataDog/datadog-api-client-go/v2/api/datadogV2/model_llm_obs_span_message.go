// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSpanMessage A single message in a span input or output.
type LLMObsSpanMessage struct {
	// Text content of the message.
	Content *string `json:"content,omitempty"`
	// Unique identifier of the message.
	Id *string `json:"id,omitempty"`
	// Role of the message sender (e.g., user, assistant, system).
	Role *string `json:"role,omitempty"`
	// Tool calls made in this message.
	ToolCalls []LLMObsSpanToolCall `json:"tool_calls,omitempty"`
	// Tool results returned in this message.
	ToolResults []LLMObsSpanToolResult `json:"tool_results,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsSpanMessage instantiates a new LLMObsSpanMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsSpanMessage() *LLMObsSpanMessage {
	this := LLMObsSpanMessage{}
	return &this
}

// NewLLMObsSpanMessageWithDefaults instantiates a new LLMObsSpanMessage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsSpanMessageWithDefaults() *LLMObsSpanMessage {
	this := LLMObsSpanMessage{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *LLMObsSpanMessage) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanMessage) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *LLMObsSpanMessage) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *LLMObsSpanMessage) SetContent(v string) {
	o.Content = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LLMObsSpanMessage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanMessage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LLMObsSpanMessage) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LLMObsSpanMessage) SetId(v string) {
	o.Id = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *LLMObsSpanMessage) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanMessage) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *LLMObsSpanMessage) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *LLMObsSpanMessage) SetRole(v string) {
	o.Role = &v
}

// GetToolCalls returns the ToolCalls field value if set, zero value otherwise.
func (o *LLMObsSpanMessage) GetToolCalls() []LLMObsSpanToolCall {
	if o == nil || o.ToolCalls == nil {
		var ret []LLMObsSpanToolCall
		return ret
	}
	return o.ToolCalls
}

// GetToolCallsOk returns a tuple with the ToolCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanMessage) GetToolCallsOk() (*[]LLMObsSpanToolCall, bool) {
	if o == nil || o.ToolCalls == nil {
		return nil, false
	}
	return &o.ToolCalls, true
}

// HasToolCalls returns a boolean if a field has been set.
func (o *LLMObsSpanMessage) HasToolCalls() bool {
	return o != nil && o.ToolCalls != nil
}

// SetToolCalls gets a reference to the given []LLMObsSpanToolCall and assigns it to the ToolCalls field.
func (o *LLMObsSpanMessage) SetToolCalls(v []LLMObsSpanToolCall) {
	o.ToolCalls = v
}

// GetToolResults returns the ToolResults field value if set, zero value otherwise.
func (o *LLMObsSpanMessage) GetToolResults() []LLMObsSpanToolResult {
	if o == nil || o.ToolResults == nil {
		var ret []LLMObsSpanToolResult
		return ret
	}
	return o.ToolResults
}

// GetToolResultsOk returns a tuple with the ToolResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanMessage) GetToolResultsOk() (*[]LLMObsSpanToolResult, bool) {
	if o == nil || o.ToolResults == nil {
		return nil, false
	}
	return &o.ToolResults, true
}

// HasToolResults returns a boolean if a field has been set.
func (o *LLMObsSpanMessage) HasToolResults() bool {
	return o != nil && o.ToolResults != nil
}

// SetToolResults gets a reference to the given []LLMObsSpanToolResult and assigns it to the ToolResults field.
func (o *LLMObsSpanMessage) SetToolResults(v []LLMObsSpanToolResult) {
	o.ToolResults = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsSpanMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
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
func (o *LLMObsSpanMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content     *string                `json:"content,omitempty"`
		Id          *string                `json:"id,omitempty"`
		Role        *string                `json:"role,omitempty"`
		ToolCalls   []LLMObsSpanToolCall   `json:"tool_calls,omitempty"`
		ToolResults []LLMObsSpanToolResult `json:"tool_results,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "id", "role", "tool_calls", "tool_results"})
	} else {
		return err
	}
	o.Content = all.Content
	o.Id = all.Id
	o.Role = all.Role
	o.ToolCalls = all.ToolCalls
	o.ToolResults = all.ToolResults

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
