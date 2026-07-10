// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipInferenceType The type of the ownership inference resource. The value should always be `ownership_inference`.
type OwnershipInferenceType string

// List of OwnershipInferenceType.
const (
	OWNERSHIPINFERENCETYPE_OWNERSHIP_INFERENCE OwnershipInferenceType = "ownership_inference"
)

var allowedOwnershipInferenceTypeEnumValues = []OwnershipInferenceType{
	OWNERSHIPINFERENCETYPE_OWNERSHIP_INFERENCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OwnershipInferenceType) GetAllowedValues() []OwnershipInferenceType {
	return allowedOwnershipInferenceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OwnershipInferenceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OwnershipInferenceType(value)
	return nil
}

// NewOwnershipInferenceTypeFromValue returns a pointer to a valid OwnershipInferenceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOwnershipInferenceTypeFromValue(v string) (*OwnershipInferenceType, error) {
	ev := OwnershipInferenceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OwnershipInferenceType: valid values are %v", v, allowedOwnershipInferenceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OwnershipInferenceType) IsValid() bool {
	for _, existing := range allowedOwnershipInferenceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OwnershipInferenceType value.
func (v OwnershipInferenceType) Ptr() *OwnershipInferenceType {
	return &v
}
