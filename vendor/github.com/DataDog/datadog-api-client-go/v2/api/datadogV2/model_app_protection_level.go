// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppProtectionLevel The publication protection level of the app. `approval_required` means changes must go through an approval workflow before being published.
type AppProtectionLevel string

// List of AppProtectionLevel.
const (
	APPPROTECTIONLEVEL_DIRECT_PUBLISH    AppProtectionLevel = "direct_publish"
	APPPROTECTIONLEVEL_APPROVAL_REQUIRED AppProtectionLevel = "approval_required"
)

var allowedAppProtectionLevelEnumValues = []AppProtectionLevel{
	APPPROTECTIONLEVEL_DIRECT_PUBLISH,
	APPPROTECTIONLEVEL_APPROVAL_REQUIRED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppProtectionLevel) GetAllowedValues() []AppProtectionLevel {
	return allowedAppProtectionLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppProtectionLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppProtectionLevel(value)
	return nil
}

// NewAppProtectionLevelFromValue returns a pointer to a valid AppProtectionLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppProtectionLevelFromValue(v string) (*AppProtectionLevel, error) {
	ev := AppProtectionLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppProtectionLevel: valid values are %v", v, allowedAppProtectionLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppProtectionLevel) IsValid() bool {
	for _, existing := range allowedAppProtectionLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppProtectionLevel value.
func (v AppProtectionLevel) Ptr() *AppProtectionLevel {
	return &v
}
