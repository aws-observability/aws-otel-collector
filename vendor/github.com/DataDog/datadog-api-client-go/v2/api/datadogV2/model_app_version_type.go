// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppVersionType The app-version resource type.
type AppVersionType string

// List of AppVersionType.
const (
	APPVERSIONTYPE_APPVERSIONS AppVersionType = "appVersions"
)

var allowedAppVersionTypeEnumValues = []AppVersionType{
	APPVERSIONTYPE_APPVERSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppVersionType) GetAllowedValues() []AppVersionType {
	return allowedAppVersionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppVersionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppVersionType(value)
	return nil
}

// NewAppVersionTypeFromValue returns a pointer to a valid AppVersionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppVersionTypeFromValue(v string) (*AppVersionType, error) {
	ev := AppVersionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppVersionType: valid values are %v", v, allowedAppVersionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppVersionType) IsValid() bool {
	for _, existing := range allowedAppVersionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppVersionType value.
func (v AppVersionType) Ptr() *AppVersionType {
	return &v
}
