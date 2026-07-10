// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsInferenceToolCall A tool call made during LLM inference.
type LLMObsInferenceToolCall struct {
	// The arguments passed to the tool.
	Arguments map[string]interface{} `json:"arguments,omitempty"`
	// The name of the tool being called.
	Name *string `json:"name,omitempty"`
	// Unique identifier for the tool call.
	ToolId *string `json:"tool_id,omitempty"`
	// The type of tool call.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsInferenceToolCall instantiates a new LLMObsInferenceToolCall object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsInferenceToolCall() *LLMObsInferenceToolCall {
	this := LLMObsInferenceToolCall{}
	return &this
}

// NewLLMObsInferenceToolCallWithDefaults instantiates a new LLMObsInferenceToolCall object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsInferenceToolCallWithDefaults() *LLMObsInferenceToolCall {
	this := LLMObsInferenceToolCall{}
	return &this
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *LLMObsInferenceToolCall) GetArguments() map[string]interface{} {
	if o == nil || o.Arguments == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceToolCall) GetArgumentsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *LLMObsInferenceToolCall) HasArguments() bool {
	return o != nil && o.Arguments != nil
}

// SetArguments gets a reference to the given map[string]interface{} and assigns it to the Arguments field.
func (o *LLMObsInferenceToolCall) SetArguments(v map[string]interface{}) {
	o.Arguments = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LLMObsInferenceToolCall) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceToolCall) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LLMObsInferenceToolCall) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LLMObsInferenceToolCall) SetName(v string) {
	o.Name = &v
}

// GetToolId returns the ToolId field value if set, zero value otherwise.
func (o *LLMObsInferenceToolCall) GetToolId() string {
	if o == nil || o.ToolId == nil {
		var ret string
		return ret
	}
	return *o.ToolId
}

// GetToolIdOk returns a tuple with the ToolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceToolCall) GetToolIdOk() (*string, bool) {
	if o == nil || o.ToolId == nil {
		return nil, false
	}
	return o.ToolId, true
}

// HasToolId returns a boolean if a field has been set.
func (o *LLMObsInferenceToolCall) HasToolId() bool {
	return o != nil && o.ToolId != nil
}

// SetToolId gets a reference to the given string and assigns it to the ToolId field.
func (o *LLMObsInferenceToolCall) SetToolId(v string) {
	o.ToolId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LLMObsInferenceToolCall) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceToolCall) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LLMObsInferenceToolCall) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LLMObsInferenceToolCall) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsInferenceToolCall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
func (o *LLMObsInferenceToolCall) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arguments map[string]interface{} `json:"arguments,omitempty"`
		Name      *string                `json:"name,omitempty"`
		ToolId    *string                `json:"tool_id,omitempty"`
		Type      *string                `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arguments", "name", "tool_id", "type"})
	} else {
		return err
	}
	o.Arguments = all.Arguments
	o.Name = all.Name
	o.ToolId = all.ToolId
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
