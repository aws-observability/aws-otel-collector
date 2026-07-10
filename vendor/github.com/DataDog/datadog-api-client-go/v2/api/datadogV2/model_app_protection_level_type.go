// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppProtectionLevelType The protection-level resource type.
type AppProtectionLevelType string

// List of AppProtectionLevelType.
const (
	APPPROTECTIONLEVELTYPE_PROTECTIONLEVEL AppProtectionLevelType = "protectionLevel"
)

var allowedAppProtectionLevelTypeEnumValues = []AppProtectionLevelType{
	APPPROTECTIONLEVELTYPE_PROTECTIONLEVEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppProtectionLevelType) GetAllowedValues() []AppProtectionLevelType {
	return allowedAppProtectionLevelTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppProtectionLevelType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppProtectionLevelType(value)
	return nil
}

// NewAppProtectionLevelTypeFromValue returns a pointer to a valid AppProtectionLevelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppProtectionLevelTypeFromValue(v string) (*AppProtectionLevelType, error) {
	ev := AppProtectionLevelType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppProtectionLevelType: valid values are %v", v, allowedAppProtectionLevelTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppProtectionLevelType) IsValid() bool {
	for _, existing := range allowedAppProtectionLevelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppProtectionLevelType value.
func (v AppProtectionLevelType) Ptr() *AppProtectionLevelType {
	return &v
}
