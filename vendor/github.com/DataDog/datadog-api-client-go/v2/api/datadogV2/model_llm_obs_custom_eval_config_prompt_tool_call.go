// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigPromptToolCall A tool call within a prompt message.
type LLMObsCustomEvalConfigPromptToolCall struct {
	// JSON-encoded arguments for the tool call.
	Arguments *string `json:"arguments,omitempty"`
	// Unique identifier of the tool call.
	Id *string `json:"id,omitempty"`
	// Name of the tool being called.
	Name *string `json:"name,omitempty"`
	// Type of the tool call.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCustomEvalConfigPromptToolCall instantiates a new LLMObsCustomEvalConfigPromptToolCall object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCustomEvalConfigPromptToolCall() *LLMObsCustomEvalConfigPromptToolCall {
	this := LLMObsCustomEvalConfigPromptToolCall{}
	return &this
}

// NewLLMObsCustomEvalConfigPromptToolCallWithDefaults instantiates a new LLMObsCustomEvalConfigPromptToolCall object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCustomEvalConfigPromptToolCallWithDefaults() *LLMObsCustomEvalConfigPromptToolCall {
	this := LLMObsCustomEvalConfigPromptToolCall{}
	return &this
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptToolCall) GetArguments() string {
	if o == nil || o.Arguments == nil {
		var ret string
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptToolCall) GetArgumentsOk() (*string, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptToolCall) HasArguments() bool {
	return o != nil && o.Arguments != nil
}

// SetArguments gets a reference to the given string and assigns it to the Arguments field.
func (o *LLMObsCustomEvalConfigPromptToolCall) SetArguments(v string) {
	o.Arguments = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptToolCall) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptToolCall) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptToolCall) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LLMObsCustomEvalConfigPromptToolCall) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptToolCall) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptToolCall) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptToolCall) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LLMObsCustomEvalConfigPromptToolCall) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptToolCall) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptToolCall) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptToolCall) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LLMObsCustomEvalConfigPromptToolCall) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCustomEvalConfigPromptToolCall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
func (o *LLMObsCustomEvalConfigPromptToolCall) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arguments *string `json:"arguments,omitempty"`
		Id        *string `json:"id,omitempty"`
		Name      *string `json:"name,omitempty"`
		Type      *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arguments", "id", "name", "type"})
	} else {
		return err
	}
	o.Arguments = all.Arguments
	o.Id = all.Id
	o.Name = all.Name
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
