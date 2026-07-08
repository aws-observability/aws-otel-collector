// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppFavoriteType The favorite resource type.
type AppFavoriteType string

// List of AppFavoriteType.
const (
	APPFAVORITETYPE_FAVORITES AppFavoriteType = "favorites"
)

var allowedAppFavoriteTypeEnumValues = []AppFavoriteType{
	APPFAVORITETYPE_FAVORITES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppFavoriteType) GetAllowedValues() []AppFavoriteType {
	return allowedAppFavoriteTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppFavoriteType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppFavoriteType(value)
	return nil
}

// NewAppFavoriteTypeFromValue returns a pointer to a valid AppFavoriteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppFavoriteTypeFromValue(v string) (*AppFavoriteType, error) {
	ev := AppFavoriteType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppFavoriteType: valid values are %v", v, allowedAppFavoriteTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppFavoriteType) IsValid() bool {
	for _, existing := range allowedAppFavoriteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppFavoriteType value.
func (v AppFavoriteType) Ptr() *AppFavoriteType {
	return &v
}
