// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimelineCellMarkdownCreateAttributes Timeline cell data for Markdown timeline cells for a create request.
type IncidentTimelineCellMarkdownCreateAttributes struct {
	// Type of the Markdown timeline cell.
	CellType IncidentTimelineCellMarkdownContentType `json:"cell_type"`
	// The Markdown timeline cell contents.
	Content IncidentTimelineCellMarkdownCreateAttributesContent `json:"content"`
	// A flag indicating whether the timeline cell is important and should be highlighted.
	Important *bool `json:"important,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTimelineCellMarkdownCreateAttributes instantiates a new IncidentTimelineCellMarkdownCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTimelineCellMarkdownCreateAttributes(cellType IncidentTimelineCellMarkdownContentType, content IncidentTimelineCellMarkdownCreateAttributesContent) *IncidentTimelineCellMarkdownCreateAttributes {
	this := IncidentTimelineCellMarkdownCreateAttributes{}
	this.CellType = cellType
	this.Content = content
	var important bool = false
	this.Important = &important
	return &this
}

// NewIncidentTimelineCellMarkdownCreateAttributesWithDefaults instantiates a new IncidentTimelineCellMarkdownCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTimelineCellMarkdownCreateAttributesWithDefaults() *IncidentTimelineCellMarkdownCreateAttributes {
	this := IncidentTimelineCellMarkdownCreateAttributes{}
	var cellType IncidentTimelineCellMarkdownContentType = INCIDENTTIMELINECELLMARKDOWNCONTENTTYPE_MARKDOWN
	this.CellType = cellType
	var important bool = false
	this.Important = &important
	return &this
}

// GetCellType returns the CellType field value.
func (o *IncidentTimelineCellMarkdownCreateAttributes) GetCellType() IncidentTimelineCellMarkdownContentType {
	if o == nil {
		var ret IncidentTimelineCellMarkdownContentType
		return ret
	}
	return o.CellType
}

// GetCellTypeOk returns a tuple with the CellType field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineCellMarkdownCreateAttributes) GetCellTypeOk() (*IncidentTimelineCellMarkdownContentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CellType, true
}

// SetCellType sets field value.
func (o *IncidentTimelineCellMarkdownCreateAttributes) SetCellType(v IncidentTimelineCellMarkdownContentType) {
	o.CellType = v
}

// GetContent returns the Content field value.
func (o *IncidentTimelineCellMarkdownCreateAttributes) GetContent() IncidentTimelineCellMarkdownCreateAttributesContent {
	if o == nil {
		var ret IncidentTimelineCellMarkdownCreateAttributesContent
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineCellMarkdownCreateAttributes) GetContentOk() (*IncidentTimelineCellMarkdownCreateAttributesContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *IncidentTimelineCellMarkdownCreateAttributes) SetContent(v IncidentTimelineCellMarkdownCreateAttributesContent) {
	o.Content = v
}

// GetImportant returns the Important field value if set, zero value otherwise.
func (o *IncidentTimelineCellMarkdownCreateAttributes) GetImportant() bool {
	if o == nil || o.Important == nil {
		var ret bool
		return ret
	}
	return *o.Important
}

// GetImportantOk returns a tuple with the Important field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTimelineCellMarkdownCreateAttributes) GetImportantOk() (*bool, bool) {
	if o == nil || o.Important == nil {
		return nil, false
	}
	return o.Important, true
}

// HasImportant returns a boolean if a field has been set.
func (o *IncidentTimelineCellMarkdownCreateAttributes) HasImportant() bool {
	return o != nil && o.Important != nil
}

// SetImportant gets a reference to the given bool and assigns it to the Important field.
func (o *IncidentTimelineCellMarkdownCreateAttributes) SetImportant(v bool) {
	o.Important = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTimelineCellMarkdownCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cell_type"] = o.CellType
	toSerialize["content"] = o.Content
	if o.Important != nil {
		toSerialize["important"] = o.Important
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTimelineCellMarkdownCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CellType  *IncidentTimelineCellMarkdownContentType             `json:"cell_type"`
		Content   *IncidentTimelineCellMarkdownCreateAttributesContent `json:"content"`
		Important *bool                                                `json:"important,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CellType == nil {
		return fmt.Errorf("required field cell_type missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cell_type", "content", "important"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.CellType.IsValid() {
		hasInvalidField = true
	} else {
		o.CellType = *all.CellType
	}
	if all.Content.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Content = *all.Content
	o.Important = all.Important

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
