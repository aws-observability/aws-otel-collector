// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppKeyRegistrationDataType The definition of `AppKeyRegistrationDataType` object.
type AppKeyRegistrationDataType string

// List of AppKeyRegistrationDataType.
const (
	APPKEYREGISTRATIONDATATYPE_APP_KEY_REGISTRATION AppKeyRegistrationDataType = "app_key_registration"
)

var allowedAppKeyRegistrationDataTypeEnumValues = []AppKeyRegistrationDataType{
	APPKEYREGISTRATIONDATATYPE_APP_KEY_REGISTRATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppKeyRegistrationDataType) GetAllowedValues() []AppKeyRegistrationDataType {
	return allowedAppKeyRegistrationDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppKeyRegistrationDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppKeyRegistrationDataType(value)
	return nil
}

// NewAppKeyRegistrationDataTypeFromValue returns a pointer to a valid AppKeyRegistrationDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppKeyRegistrationDataTypeFromValue(v string) (*AppKeyRegistrationDataType, error) {
	ev := AppKeyRegistrationDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppKeyRegistrationDataType: valid values are %v", v, allowedAppKeyRegistrationDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppKeyRegistrationDataType) IsValid() bool {
	for _, existing := range allowedAppKeyRegistrationDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppKeyRegistrationDataType value.
func (v AppKeyRegistrationDataType) Ptr() *AppKeyRegistrationDataType {
	return &v
}
