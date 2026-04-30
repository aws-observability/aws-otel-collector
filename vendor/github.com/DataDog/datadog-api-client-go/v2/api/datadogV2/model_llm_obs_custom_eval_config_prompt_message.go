// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigPromptMessage A message in the prompt template for a custom LLM judge evaluator.
type LLMObsCustomEvalConfigPromptMessage struct {
	// Text content of the message.
	Content *string `json:"content,omitempty"`
	// Multi-part content blocks for the message.
	Contents []LLMObsCustomEvalConfigPromptContent `json:"contents,omitempty"`
	// Role of the message author.
	Role string `json:"role"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCustomEvalConfigPromptMessage instantiates a new LLMObsCustomEvalConfigPromptMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCustomEvalConfigPromptMessage(role string) *LLMObsCustomEvalConfigPromptMessage {
	this := LLMObsCustomEvalConfigPromptMessage{}
	this.Role = role
	return &this
}

// NewLLMObsCustomEvalConfigPromptMessageWithDefaults instantiates a new LLMObsCustomEvalConfigPromptMessage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCustomEvalConfigPromptMessageWithDefaults() *LLMObsCustomEvalConfigPromptMessage {
	this := LLMObsCustomEvalConfigPromptMessage{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptMessage) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptMessage) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptMessage) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *LLMObsCustomEvalConfigPromptMessage) SetContent(v string) {
	o.Content = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigPromptMessage) GetContents() []LLMObsCustomEvalConfigPromptContent {
	if o == nil || o.Contents == nil {
		var ret []LLMObsCustomEvalConfigPromptContent
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptMessage) GetContentsOk() (*[]LLMObsCustomEvalConfigPromptContent, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return &o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigPromptMessage) HasContents() bool {
	return o != nil && o.Contents != nil
}

// SetContents gets a reference to the given []LLMObsCustomEvalConfigPromptContent and assigns it to the Contents field.
func (o *LLMObsCustomEvalConfigPromptMessage) SetContents(v []LLMObsCustomEvalConfigPromptContent) {
	o.Contents = v
}

// GetRole returns the Role field value.
func (o *LLMObsCustomEvalConfigPromptMessage) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigPromptMessage) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *LLMObsCustomEvalConfigPromptMessage) SetRole(v string) {
	o.Role = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCustomEvalConfigPromptMessage) MarshalJSON() ([]byte, error) {
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
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCustomEvalConfigPromptMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content  *string                               `json:"content,omitempty"`
		Contents []LLMObsCustomEvalConfigPromptContent `json:"contents,omitempty"`
		Role     *string                               `json:"role"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "contents", "role"})
	} else {
		return err
	}
	o.Content = all.Content
	o.Contents = all.Contents
	o.Role = *all.Role

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
