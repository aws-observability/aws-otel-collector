// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMApplicationUpdateType RUM application update type.
type RUMApplicationUpdateType string

// List of RUMApplicationUpdateType.
const (
	RUMAPPLICATIONUPDATETYPE_RUM_APPLICATION_UPDATE RUMApplicationUpdateType = "rum_application_update"
)

var allowedRUMApplicationUpdateTypeEnumValues = []RUMApplicationUpdateType{
	RUMAPPLICATIONUPDATETYPE_RUM_APPLICATION_UPDATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMApplicationUpdateType) GetAllowedValues() []RUMApplicationUpdateType {
	return allowedRUMApplicationUpdateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMApplicationUpdateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMApplicationUpdateType(value)
	return nil
}

// NewRUMApplicationUpdateTypeFromValue returns a pointer to a valid RUMApplicationUpdateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMApplicationUpdateTypeFromValue(v string) (*RUMApplicationUpdateType, error) {
	ev := RUMApplicationUpdateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMApplicationUpdateType: valid values are %v", v, allowedRUMApplicationUpdateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMApplicationUpdateType) IsValid() bool {
	for _, existing := range allowedRUMApplicationUpdateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMApplicationUpdateType value.
func (v RUMApplicationUpdateType) Ptr() *RUMApplicationUpdateType {
	return &v
}
