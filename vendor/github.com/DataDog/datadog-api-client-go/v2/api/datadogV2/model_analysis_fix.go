// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisFix A fix suggestion for a rule violation, consisting of one or more edit operations.
type AnalysisFix struct {
	// A human-readable description of what the fix does.
	Description string `json:"description"`
	// The list of edit operations that constitute the fix.
	Edits []AnalysisEdit `json:"edits"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnalysisFix instantiates a new AnalysisFix object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisFix(description string, edits []AnalysisEdit) *AnalysisFix {
	this := AnalysisFix{}
	this.Description = description
	this.Edits = edits
	return &this
}

// NewAnalysisFixWithDefaults instantiates a new AnalysisFix object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisFixWithDefaults() *AnalysisFix {
	this := AnalysisFix{}
	return &this
}

// GetDescription returns the Description field value.
func (o *AnalysisFix) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AnalysisFix) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *AnalysisFix) SetDescription(v string) {
	o.Description = v
}

// GetEdits returns the Edits field value.
func (o *AnalysisFix) GetEdits() []AnalysisEdit {
	if o == nil {
		var ret []AnalysisEdit
		return ret
	}
	return o.Edits
}

// GetEditsOk returns a tuple with the Edits field value
// and a boolean to check if the value has been set.
func (o *AnalysisFix) GetEditsOk() (*[]AnalysisEdit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edits, true
}

// SetEdits sets field value.
func (o *AnalysisFix) SetEdits(v []AnalysisEdit) {
	o.Edits = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisFix) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["edits"] = o.Edits

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisFix) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string         `json:"description"`
		Edits       *[]AnalysisEdit `json:"edits"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Edits == nil {
		return fmt.Errorf("required field edits missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "edits"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.Edits = *all.Edits

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
