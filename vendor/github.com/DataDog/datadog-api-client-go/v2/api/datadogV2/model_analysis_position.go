// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisPosition A position in source code, identified by line and column numbers.
type AnalysisPosition struct {
	// The column number in the source file (1-based).
	Col int64 `json:"col"`
	// The line number in the source file (1-based).
	Line int64 `json:"line"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnalysisPosition instantiates a new AnalysisPosition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisPosition(col int64, line int64) *AnalysisPosition {
	this := AnalysisPosition{}
	this.Col = col
	this.Line = line
	return &this
}

// NewAnalysisPositionWithDefaults instantiates a new AnalysisPosition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisPositionWithDefaults() *AnalysisPosition {
	this := AnalysisPosition{}
	return &this
}

// GetCol returns the Col field value.
func (o *AnalysisPosition) GetCol() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Col
}

// GetColOk returns a tuple with the Col field value
// and a boolean to check if the value has been set.
func (o *AnalysisPosition) GetColOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Col, true
}

// SetCol sets field value.
func (o *AnalysisPosition) SetCol(v int64) {
	o.Col = v
}

// GetLine returns the Line field value.
func (o *AnalysisPosition) GetLine() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Line
}

// GetLineOk returns a tuple with the Line field value
// and a boolean to check if the value has been set.
func (o *AnalysisPosition) GetLineOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line, true
}

// SetLine sets field value.
func (o *AnalysisPosition) SetLine(v int64) {
	o.Line = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisPosition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["col"] = o.Col
	toSerialize["line"] = o.Line

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisPosition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Col  *int64 `json:"col"`
		Line *int64 `json:"line"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Col == nil {
		return fmt.Errorf("required field col missing")
	}
	if all.Line == nil {
		return fmt.Errorf("required field line missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"col", "line"})
	} else {
		return err
	}
	o.Col = *all.Col
	o.Line = *all.Line

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
