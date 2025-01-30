// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppsSortField The field and direction to sort apps by
type AppsSortField string

// List of AppsSortField.
const (
	APPSSORTFIELD_NAME            AppsSortField = "name"
	APPSSORTFIELD_CREATED_AT      AppsSortField = "created_at"
	APPSSORTFIELD_UPDATED_AT      AppsSortField = "updated_at"
	APPSSORTFIELD_USER_NAME       AppsSortField = "user_name"
	APPSSORTFIELD_NAME_DESC       AppsSortField = "-name"
	APPSSORTFIELD_CREATED_AT_DESC AppsSortField = "-created_at"
	APPSSORTFIELD_UPDATED_AT_DESC AppsSortField = "-updated_at"
	APPSSORTFIELD_USER_NAME_DESC  AppsSortField = "-user_name"
)

var allowedAppsSortFieldEnumValues = []AppsSortField{
	APPSSORTFIELD_NAME,
	APPSSORTFIELD_CREATED_AT,
	APPSSORTFIELD_UPDATED_AT,
	APPSSORTFIELD_USER_NAME,
	APPSSORTFIELD_NAME_DESC,
	APPSSORTFIELD_CREATED_AT_DESC,
	APPSSORTFIELD_UPDATED_AT_DESC,
	APPSSORTFIELD_USER_NAME_DESC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppsSortField) GetAllowedValues() []AppsSortField {
	return allowedAppsSortFieldEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppsSortField) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppsSortField(value)
	return nil
}

// NewAppsSortFieldFromValue returns a pointer to a valid AppsSortField
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppsSortFieldFromValue(v string) (*AppsSortField, error) {
	ev := AppsSortField(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppsSortField: valid values are %v", v, allowedAppsSortFieldEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppsSortField) IsValid() bool {
	for _, existing := range allowedAppsSortFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppsSortField value.
func (v AppsSortField) Ptr() *AppsSortField {
	return &v
}
