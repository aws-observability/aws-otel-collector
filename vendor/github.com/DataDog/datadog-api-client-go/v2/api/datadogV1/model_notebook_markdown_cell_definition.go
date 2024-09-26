// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookMarkdownCellDefinition Text in a notebook is formatted with [Markdown](https://daringfireball.net/projects/markdown/), which enables the use of headings, subheadings, links, images, lists, and code blocks.
type NotebookMarkdownCellDefinition struct {
	// The markdown content.
	Text string `json:"text"`
	// Type of the markdown cell.
	Type NotebookMarkdownCellDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotebookMarkdownCellDefinition instantiates a new NotebookMarkdownCellDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookMarkdownCellDefinition(text string, typeVar NotebookMarkdownCellDefinitionType) *NotebookMarkdownCellDefinition {
	this := NotebookMarkdownCellDefinition{}
	this.Text = text
	this.Type = typeVar
	return &this
}

// NewNotebookMarkdownCellDefinitionWithDefaults instantiates a new NotebookMarkdownCellDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookMarkdownCellDefinitionWithDefaults() *NotebookMarkdownCellDefinition {
	this := NotebookMarkdownCellDefinition{}
	var typeVar NotebookMarkdownCellDefinitionType = NOTEBOOKMARKDOWNCELLDEFINITIONTYPE_MARKDOWN
	this.Type = typeVar
	return &this
}

// GetText returns the Text field value.
func (o *NotebookMarkdownCellDefinition) GetText() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *NotebookMarkdownCellDefinition) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value.
func (o *NotebookMarkdownCellDefinition) SetText(v string) {
	o.Text = v
}

// GetType returns the Type field value.
func (o *NotebookMarkdownCellDefinition) GetType() NotebookMarkdownCellDefinitionType {
	if o == nil {
		var ret NotebookMarkdownCellDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotebookMarkdownCellDefinition) GetTypeOk() (*NotebookMarkdownCellDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NotebookMarkdownCellDefinition) SetType(v NotebookMarkdownCellDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookMarkdownCellDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["text"] = o.Text
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookMarkdownCellDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Text *string                             `json:"text"`
		Type *NotebookMarkdownCellDefinitionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Text == nil {
		return fmt.Errorf("required field text missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"text", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Text = *all.Text
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
