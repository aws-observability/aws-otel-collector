// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigPromptToolResult A tool call result within a prompt message.
type LLMObsCustomEvalConfigPromptToolResult struct {
	// Name of the tool that produced this result.
	Name *string `json:"name,omitempty"`
	// The result returned by the tool.
	Result *string `json:"result,omitempty"`
	// Identifier of the tool call this result corresponds to.
	ToolId *string `json:"tool_id,omitempty"`
	// Type of the tool result.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCustomEvalConfigPromptToolResult instantiates a new LLMObsCustomEvalConfigPromptToolResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCustomEvalConfigPromptToolResult() *LLMObsCustomEvalConfigPromptToolResult {
	this := LLMObsCustomEvalConfigPromptToolResult{}
	return &this
}

// NewLLMObsCustomEvalConfigPromptToolResultWithDefaults instantiates a new LLMObsCustomEvalConfigPromptToolResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCustomEvalConfigPromptToolResultWithDefaults() *LLMObsCustomEvalConfigPromptToolResult {
	this := LLMObsCustomEvalConfigPromptToolResult{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptToolResult) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptToolResult) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptToolResult) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LLMObsCustomEvalConfigPromptToolResult) SetName(v string) {
	o.Name = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptToolResult) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptToolResult) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptToolResult) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *LLMObsCustomEvalConfigPromptToolResult) SetResult(v string) {
	o.Result = &v
}

// GetToolId returns the ToolId field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptToolResult) GetToolId() string {
	if o == nil || o.ToolId == nil {
		var ret string
		return ret
	}
	return *o.ToolId
}

// GetToolIdOk returns a tuple with the ToolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptToolResult) GetToolIdOk() (*string, bool) {
	if o == nil || o.ToolId == nil {
		return nil, false
	}
	return o.ToolId, true
}

// HasToolId returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptToolResult) HasToolId() bool {
	return o != nil && o.ToolId != nil
}

// SetToolId gets a reference to the given string and assigns it to the ToolId field.
func (o *LLMObsCustomEvalConfigPromptToolResult) SetToolId(v string) {
	o.ToolId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptToolResult) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptToolResult) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptToolResult) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LLMObsCustomEvalConfigPromptToolResult) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCustomEvalConfigPromptToolResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.ToolId != nil {
		toSerialize["tool_id"] = o.ToolId
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
func (o *LLMObsCustomEvalConfigPromptToolResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string `json:"name,omitempty"`
		Result *string `json:"result,omitempty"`
		ToolId *string `json:"tool_id,omitempty"`
		Type   *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "result", "tool_id", "type"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Result = all.Result
	o.ToolId = all.ToolId
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
