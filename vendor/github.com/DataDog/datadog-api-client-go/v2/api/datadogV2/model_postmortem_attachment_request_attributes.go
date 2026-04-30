// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PostmortemAttachmentRequestAttributes Postmortem attachment attributes
type PostmortemAttachmentRequestAttributes struct {
	// The cells of the postmortem
	Cells []PostmortemCell `json:"cells,omitempty"`
	// The content of the postmortem
	Content *string `json:"content,omitempty"`
	// The ID of the postmortem template
	PostmortemTemplateId *string `json:"postmortem_template_id,omitempty"`
	// The title of the postmortem
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostmortemAttachmentRequestAttributes instantiates a new PostmortemAttachmentRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostmortemAttachmentRequestAttributes() *PostmortemAttachmentRequestAttributes {
	this := PostmortemAttachmentRequestAttributes{}
	return &this
}

// NewPostmortemAttachmentRequestAttributesWithDefaults instantiates a new PostmortemAttachmentRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostmortemAttachmentRequestAttributesWithDefaults() *PostmortemAttachmentRequestAttributes {
	this := PostmortemAttachmentRequestAttributes{}
	return &this
}

// GetCells returns the Cells field value if set, zero value otherwise.
func (o *PostmortemAttachmentRequestAttributes) GetCells() []PostmortemCell {
	if o == nil || o.Cells == nil {
		var ret []PostmortemCell
		return ret
	}
	return o.Cells
}

// GetCellsOk returns a tuple with the Cells field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemAttachmentRequestAttributes) GetCellsOk() (*[]PostmortemCell, bool) {
	if o == nil || o.Cells == nil {
		return nil, false
	}
	return &o.Cells, true
}

// HasCells returns a boolean if a field has been set.
func (o *PostmortemAttachmentRequestAttributes) HasCells() bool {
	return o != nil && o.Cells != nil
}

// SetCells gets a reference to the given []PostmortemCell and assigns it to the Cells field.
func (o *PostmortemAttachmentRequestAttributes) SetCells(v []PostmortemCell) {
	o.Cells = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *PostmortemAttachmentRequestAttributes) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemAttachmentRequestAttributes) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *PostmortemAttachmentRequestAttributes) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *PostmortemAttachmentRequestAttributes) SetContent(v string) {
	o.Content = &v
}

// GetPostmortemTemplateId returns the PostmortemTemplateId field value if set, zero value otherwise.
func (o *PostmortemAttachmentRequestAttributes) GetPostmortemTemplateId() string {
	if o == nil || o.PostmortemTemplateId == nil {
		var ret string
		return ret
	}
	return *o.PostmortemTemplateId
}

// GetPostmortemTemplateIdOk returns a tuple with the PostmortemTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemAttachmentRequestAttributes) GetPostmortemTemplateIdOk() (*string, bool) {
	if o == nil || o.PostmortemTemplateId == nil {
		return nil, false
	}
	return o.PostmortemTemplateId, true
}

// HasPostmortemTemplateId returns a boolean if a field has been set.
func (o *PostmortemAttachmentRequestAttributes) HasPostmortemTemplateId() bool {
	return o != nil && o.PostmortemTemplateId != nil
}

// SetPostmortemTemplateId gets a reference to the given string and assigns it to the PostmortemTemplateId field.
func (o *PostmortemAttachmentRequestAttributes) SetPostmortemTemplateId(v string) {
	o.PostmortemTemplateId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PostmortemAttachmentRequestAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemAttachmentRequestAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PostmortemAttachmentRequestAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PostmortemAttachmentRequestAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostmortemAttachmentRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cells != nil {
		toSerialize["cells"] = o.Cells
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.PostmortemTemplateId != nil {
		toSerialize["postmortem_template_id"] = o.PostmortemTemplateId
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostmortemAttachmentRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cells                []PostmortemCell `json:"cells,omitempty"`
		Content              *string          `json:"content,omitempty"`
		PostmortemTemplateId *string          `json:"postmortem_template_id,omitempty"`
		Title                *string          `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cells", "content", "postmortem_template_id", "title"})
	} else {
		return err
	}
	o.Cells = all.Cells
	o.Content = all.Content
	o.PostmortemTemplateId = all.PostmortemTemplateId
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
