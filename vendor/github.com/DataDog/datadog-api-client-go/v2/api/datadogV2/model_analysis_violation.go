// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisViolation A rule violation found in the analyzed source code.
type AnalysisViolation struct {
	// The category of the violation.
	Category string `json:"category"`
	// A position in source code, identified by line and column numbers.
	End AnalysisPosition `json:"end"`
	// The list of suggested fixes for this violation.
	Fixes []AnalysisFix `json:"fixes"`
	// A human-readable description of the violation.
	Message string `json:"message"`
	// The severity level of the violation.
	Severity string `json:"severity"`
	// A position in source code, identified by line and column numbers.
	Start AnalysisPosition `json:"start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnalysisViolation instantiates a new AnalysisViolation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisViolation(category string, end AnalysisPosition, fixes []AnalysisFix, message string, severity string, start AnalysisPosition) *AnalysisViolation {
	this := AnalysisViolation{}
	this.Category = category
	this.End = end
	this.Fixes = fixes
	this.Message = message
	this.Severity = severity
	this.Start = start
	return &this
}

// NewAnalysisViolationWithDefaults instantiates a new AnalysisViolation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisViolationWithDefaults() *AnalysisViolation {
	this := AnalysisViolation{}
	return &this
}

// GetCategory returns the Category field value.
func (o *AnalysisViolation) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *AnalysisViolation) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *AnalysisViolation) SetCategory(v string) {
	o.Category = v
}

// GetEnd returns the End field value.
func (o *AnalysisViolation) GetEnd() AnalysisPosition {
	if o == nil {
		var ret AnalysisPosition
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *AnalysisViolation) GetEndOk() (*AnalysisPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *AnalysisViolation) SetEnd(v AnalysisPosition) {
	o.End = v
}

// GetFixes returns the Fixes field value.
func (o *AnalysisViolation) GetFixes() []AnalysisFix {
	if o == nil {
		var ret []AnalysisFix
		return ret
	}
	return o.Fixes
}

// GetFixesOk returns a tuple with the Fixes field value
// and a boolean to check if the value has been set.
func (o *AnalysisViolation) GetFixesOk() (*[]AnalysisFix, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fixes, true
}

// SetFixes sets field value.
func (o *AnalysisViolation) SetFixes(v []AnalysisFix) {
	o.Fixes = v
}

// GetMessage returns the Message field value.
func (o *AnalysisViolation) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AnalysisViolation) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *AnalysisViolation) SetMessage(v string) {
	o.Message = v
}

// GetSeverity returns the Severity field value.
func (o *AnalysisViolation) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *AnalysisViolation) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *AnalysisViolation) SetSeverity(v string) {
	o.Severity = v
}

// GetStart returns the Start field value.
func (o *AnalysisViolation) GetStart() AnalysisPosition {
	if o == nil {
		var ret AnalysisPosition
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *AnalysisViolation) GetStartOk() (*AnalysisPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *AnalysisViolation) SetStart(v AnalysisPosition) {
	o.Start = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisViolation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["end"] = o.End
	toSerialize["fixes"] = o.Fixes
	toSerialize["message"] = o.Message
	toSerialize["severity"] = o.Severity
	toSerialize["start"] = o.Start

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisViolation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category *string           `json:"category"`
		End      *AnalysisPosition `json:"end"`
		Fixes    *[]AnalysisFix    `json:"fixes"`
		Message  *string           `json:"message"`
		Severity *string           `json:"severity"`
		Start    *AnalysisPosition `json:"start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.End == nil {
		return fmt.Errorf("required field end missing")
	}
	if all.Fixes == nil {
		return fmt.Errorf("required field fixes missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "end", "fixes", "message", "severity", "start"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Category = *all.Category
	if all.End.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.End = *all.End
	o.Fixes = *all.Fixes
	o.Message = *all.Message
	o.Severity = *all.Severity
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
