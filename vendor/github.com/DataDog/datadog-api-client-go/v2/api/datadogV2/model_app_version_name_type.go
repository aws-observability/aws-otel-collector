// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppVersionNameType The version-name resource type.
type AppVersionNameType string

// List of AppVersionNameType.
const (
	APPVERSIONNAMETYPE_VERSIONNAMES AppVersionNameType = "versionNames"
)

var allowedAppVersionNameTypeEnumValues = []AppVersionNameType{
	APPVERSIONNAMETYPE_VERSIONNAMES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppVersionNameType) GetAllowedValues() []AppVersionNameType {
	return allowedAppVersionNameTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppVersionNameType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppVersionNameType(value)
	return nil
}

// NewAppVersionNameTypeFromValue returns a pointer to a valid AppVersionNameType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppVersionNameTypeFromValue(v string) (*AppVersionNameType, error) {
	ev := AppVersionNameType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppVersionNameType: valid values are %v", v, allowedAppVersionNameTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppVersionNameType) IsValid() bool {
	for _, existing := range allowedAppVersionNameTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppVersionNameType value.
func (v AppVersionNameType) Ptr() *AppVersionNameType {
	return &v
}
