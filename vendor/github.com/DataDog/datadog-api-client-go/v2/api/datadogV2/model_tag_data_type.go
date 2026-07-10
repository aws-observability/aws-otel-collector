// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagDataType The resource type for a tag.
type TagDataType string

// List of TagDataType.
const (
	TAGDATATYPE_TAG TagDataType = "tag"
)

var allowedTagDataTypeEnumValues = []TagDataType{
	TAGDATATYPE_TAG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagDataType) GetAllowedValues() []TagDataType {
	return allowedTagDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagDataType(value)
	return nil
}

// NewTagDataTypeFromValue returns a pointer to a valid TagDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagDataTypeFromValue(v string) (*TagDataType, error) {
	ev := TagDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagDataType: valid values are %v", v, allowedTagDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagDataType) IsValid() bool {
	for _, existing := range allowedTagDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagDataType value.
func (v TagDataType) Ptr() *TagDataType {
	return &v
}
