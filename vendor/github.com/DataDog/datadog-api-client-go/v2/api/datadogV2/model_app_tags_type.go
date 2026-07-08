// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppTagsType The tags resource type.
type AppTagsType string

// List of AppTagsType.
const (
	APPTAGSTYPE_TAGS AppTagsType = "tags"
)

var allowedAppTagsTypeEnumValues = []AppTagsType{
	APPTAGSTYPE_TAGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppTagsType) GetAllowedValues() []AppTagsType {
	return allowedAppTagsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppTagsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppTagsType(value)
	return nil
}

// NewAppTagsTypeFromValue returns a pointer to a valid AppTagsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppTagsTypeFromValue(v string) (*AppTagsType, error) {
	ev := AppTagsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppTagsType: valid values are %v", v, allowedAppTagsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppTagsType) IsValid() bool {
	for _, existing := range allowedAppTagsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppTagsType value.
func (v AppTagsType) Ptr() *AppTagsType {
	return &v
}
