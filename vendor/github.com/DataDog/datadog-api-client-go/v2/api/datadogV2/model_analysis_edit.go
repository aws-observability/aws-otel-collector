// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisEdit A single edit operation within a fix suggestion for a rule violation.
type AnalysisEdit struct {
	// The content to insert or replace at the specified position, if applicable.
	Content datadog.NullableString `json:"content"`
	// The type of code edit to apply when fixing a violation.
	EditType AnalysisEditType `json:"edit_type"`
	// A position in source code, identified by line and column numbers.
	End AnalysisPosition `json:"end"`
	// A position in source code, identified by line and column numbers.
	Start AnalysisPosition `json:"start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnalysisEdit instantiates a new AnalysisEdit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisEdit(content datadog.NullableString, editType AnalysisEditType, end AnalysisPosition, start AnalysisPosition) *AnalysisEdit {
	this := AnalysisEdit{}
	this.Content = content
	this.EditType = editType
	this.End = end
	this.Start = start
	return &this
}

// NewAnalysisEditWithDefaults instantiates a new AnalysisEdit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisEditWithDefaults() *AnalysisEdit {
	this := AnalysisEdit{}
	var editType AnalysisEditType = ANALYSISEDITTYPE_ADD
	this.EditType = editType
	return &this
}

// GetContent returns the Content field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *AnalysisEdit) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AnalysisEdit) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// SetContent sets field value.
func (o *AnalysisEdit) SetContent(v string) {
	o.Content.Set(&v)
}

// GetEditType returns the EditType field value.
func (o *AnalysisEdit) GetEditType() AnalysisEditType {
	if o == nil {
		var ret AnalysisEditType
		return ret
	}
	return o.EditType
}

// GetEditTypeOk returns a tuple with the EditType field value
// and a boolean to check if the value has been set.
func (o *AnalysisEdit) GetEditTypeOk() (*AnalysisEditType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditType, true
}

// SetEditType sets field value.
func (o *AnalysisEdit) SetEditType(v AnalysisEditType) {
	o.EditType = v
}

// GetEnd returns the End field value.
func (o *AnalysisEdit) GetEnd() AnalysisPosition {
	if o == nil {
		var ret AnalysisPosition
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *AnalysisEdit) GetEndOk() (*AnalysisPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *AnalysisEdit) SetEnd(v AnalysisPosition) {
	o.End = v
}

// GetStart returns the Start field value.
func (o *AnalysisEdit) GetStart() AnalysisPosition {
	if o == nil {
		var ret AnalysisPosition
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *AnalysisEdit) GetStartOk() (*AnalysisPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *AnalysisEdit) SetStart(v AnalysisPosition) {
	o.Start = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisEdit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["content"] = o.Content.Get()
	toSerialize["edit_type"] = o.EditType
	toSerialize["end"] = o.End
	toSerialize["start"] = o.Start

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisEdit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content  datadog.NullableString `json:"content"`
		EditType *AnalysisEditType      `json:"edit_type"`
		End      *AnalysisPosition      `json:"end"`
		Start    *AnalysisPosition      `json:"start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Content.IsSet() {
		return fmt.Errorf("required field content missing")
	}
	if all.EditType == nil {
		return fmt.Errorf("required field edit_type missing")
	}
	if all.End == nil {
		return fmt.Errorf("required field end missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "edit_type", "end", "start"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Content = all.Content
	if !all.EditType.IsValid() {
		hasInvalidField = true
	} else {
		o.EditType = *all.EditType
	}
	if all.End.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.End = *all.End
	if all.Start.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Start = *all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
