// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipHistoryType The type of the ownership history resource. The value should always be `ownership_history`.
type OwnershipHistoryType string

// List of OwnershipHistoryType.
const (
	OWNERSHIPHISTORYTYPE_OWNERSHIP_HISTORY OwnershipHistoryType = "ownership_history"
)

var allowedOwnershipHistoryTypeEnumValues = []OwnershipHistoryType{
	OWNERSHIPHISTORYTYPE_OWNERSHIP_HISTORY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OwnershipHistoryType) GetAllowedValues() []OwnershipHistoryType {
	return allowedOwnershipHistoryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OwnershipHistoryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OwnershipHistoryType(value)
	return nil
}

// NewOwnershipHistoryTypeFromValue returns a pointer to a valid OwnershipHistoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOwnershipHistoryTypeFromValue(v string) (*OwnershipHistoryType, error) {
	ev := OwnershipHistoryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OwnershipHistoryType: valid values are %v", v, allowedOwnershipHistoryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OwnershipHistoryType) IsValid() bool {
	for _, existing := range allowedOwnershipHistoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OwnershipHistoryType value.
func (v OwnershipHistoryType) Ptr() *OwnershipHistoryType {
	return &v
}
