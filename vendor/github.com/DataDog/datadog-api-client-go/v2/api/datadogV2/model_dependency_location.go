// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DependencyLocation Static library vulnerability location.
type DependencyLocation struct {
	// Location column end.
	ColumnEnd int64 `json:"column_end"`
	// Location column start.
	ColumnStart int64 `json:"column_start"`
	// Location file name.
	FileName string `json:"file_name"`
	// Location line end.
	LineEnd int64 `json:"line_end"`
	// Location line start.
	LineStart int64 `json:"line_start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDependencyLocation instantiates a new DependencyLocation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDependencyLocation(columnEnd int64, columnStart int64, fileName string, lineEnd int64, lineStart int64) *DependencyLocation {
	this := DependencyLocation{}
	this.ColumnEnd = columnEnd
	this.ColumnStart = columnStart
	this.FileName = fileName
	this.LineEnd = lineEnd
	this.LineStart = lineStart
	return &this
}

// NewDependencyLocationWithDefaults instantiates a new DependencyLocation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDependencyLocationWithDefaults() *DependencyLocation {
	this := DependencyLocation{}
	return &this
}

// GetColumnEnd returns the ColumnEnd field value.
func (o *DependencyLocation) GetColumnEnd() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ColumnEnd
}

// GetColumnEndOk returns a tuple with the ColumnEnd field value
// and a boolean to check if the value has been set.
func (o *DependencyLocation) GetColumnEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnEnd, true
}

// SetColumnEnd sets field value.
func (o *DependencyLocation) SetColumnEnd(v int64) {
	o.ColumnEnd = v
}

// GetColumnStart returns the ColumnStart field value.
func (o *DependencyLocation) GetColumnStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ColumnStart
}

// GetColumnStartOk returns a tuple with the ColumnStart field value
// and a boolean to check if the value has been set.
func (o *DependencyLocation) GetColumnStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnStart, true
}

// SetColumnStart sets field value.
func (o *DependencyLocation) SetColumnStart(v int64) {
	o.ColumnStart = v
}

// GetFileName returns the FileName field value.
func (o *DependencyLocation) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *DependencyLocation) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value.
func (o *DependencyLocation) SetFileName(v string) {
	o.FileName = v
}

// GetLineEnd returns the LineEnd field value.
func (o *DependencyLocation) GetLineEnd() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LineEnd
}

// GetLineEndOk returns a tuple with the LineEnd field value
// and a boolean to check if the value has been set.
func (o *DependencyLocation) GetLineEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LineEnd, true
}

// SetLineEnd sets field value.
func (o *DependencyLocation) SetLineEnd(v int64) {
	o.LineEnd = v
}

// GetLineStart returns the LineStart field value.
func (o *DependencyLocation) GetLineStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LineStart
}

// GetLineStartOk returns a tuple with the LineStart field value
// and a boolean to check if the value has been set.
func (o *DependencyLocation) GetLineStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LineStart, true
}

// SetLineStart sets field value.
func (o *DependencyLocation) SetLineStart(v int64) {
	o.LineStart = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DependencyLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column_end"] = o.ColumnEnd
	toSerialize["column_start"] = o.ColumnStart
	toSerialize["file_name"] = o.FileName
	toSerialize["line_end"] = o.LineEnd
	toSerialize["line_start"] = o.LineStart

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DependencyLocation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnEnd   *int64  `json:"column_end"`
		ColumnStart *int64  `json:"column_start"`
		FileName    *string `json:"file_name"`
		LineEnd     *int64  `json:"line_end"`
		LineStart   *int64  `json:"line_start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ColumnEnd == nil {
		return fmt.Errorf("required field column_end missing")
	}
	if all.ColumnStart == nil {
		return fmt.Errorf("required field column_start missing")
	}
	if all.FileName == nil {
		return fmt.Errorf("required field file_name missing")
	}
	if all.LineEnd == nil {
		return fmt.Errorf("required field line_end missing")
	}
	if all.LineStart == nil {
		return fmt.Errorf("required field line_start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_end", "column_start", "file_name", "line_end", "line_start"})
	} else {
		return err
	}
	o.ColumnEnd = *all.ColumnEnd
	o.ColumnStart = *all.ColumnStart
	o.FileName = *all.FileName
	o.LineEnd = *all.LineEnd
	o.LineStart = *all.LineStart

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
