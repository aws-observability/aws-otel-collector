// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipInferencesType The type of the ownership inferences collection resource. The value should always be `ownership_inferences`.
type OwnershipInferencesType string

// List of OwnershipInferencesType.
const (
	OWNERSHIPINFERENCESTYPE_OWNERSHIP_INFERENCES OwnershipInferencesType = "ownership_inferences"
)

var allowedOwnershipInferencesTypeEnumValues = []OwnershipInferencesType{
	OWNERSHIPINFERENCESTYPE_OWNERSHIP_INFERENCES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OwnershipInferencesType) GetAllowedValues() []OwnershipInferencesType {
	return allowedOwnershipInferencesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OwnershipInferencesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OwnershipInferencesType(value)
	return nil
}

// NewOwnershipInferencesTypeFromValue returns a pointer to a valid OwnershipInferencesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOwnershipInferencesTypeFromValue(v string) (*OwnershipInferencesType, error) {
	ev := OwnershipInferencesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OwnershipInferencesType: valid values are %v", v, allowedOwnershipInferencesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OwnershipInferencesType) IsValid() bool {
	for _, existing := range allowedOwnershipInferencesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OwnershipInferencesType value.
func (v OwnershipInferencesType) Ptr() *OwnershipInferencesType {
	return &v
}
