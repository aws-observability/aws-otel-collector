// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PostmortemCellType The postmortem cell resource type.
type PostmortemCellType string

// List of PostmortemCellType.
const (
	POSTMORTEMCELLTYPE_MARKDOWN PostmortemCellType = "markdown"
)

var allowedPostmortemCellTypeEnumValues = []PostmortemCellType{
	POSTMORTEMCELLTYPE_MARKDOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PostmortemCellType) GetAllowedValues() []PostmortemCellType {
	return allowedPostmortemCellTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PostmortemCellType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PostmortemCellType(value)
	return nil
}

// NewPostmortemCellTypeFromValue returns a pointer to a valid PostmortemCellType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPostmortemCellTypeFromValue(v string) (*PostmortemCellType, error) {
	ev := PostmortemCellType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PostmortemCellType: valid values are %v", v, allowedPostmortemCellTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PostmortemCellType) IsValid() bool {
	for _, existing := range allowedPostmortemCellTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostmortemCellType value.
func (v PostmortemCellType) Ptr() *PostmortemCellType {
	return &v
}
