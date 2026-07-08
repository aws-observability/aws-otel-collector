// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipOwnerType The owner type for an ownership inference.
type OwnershipOwnerType string

// List of OwnershipOwnerType.
const (
	OWNERSHIPOWNERTYPE_USER    OwnershipOwnerType = "user"
	OWNERSHIPOWNERTYPE_TEAM    OwnershipOwnerType = "team"
	OWNERSHIPOWNERTYPE_SERVICE OwnershipOwnerType = "service"
	OWNERSHIPOWNERTYPE_UNKNOWN OwnershipOwnerType = "unknown"
)

var allowedOwnershipOwnerTypeEnumValues = []OwnershipOwnerType{
	OWNERSHIPOWNERTYPE_USER,
	OWNERSHIPOWNERTYPE_TEAM,
	OWNERSHIPOWNERTYPE_SERVICE,
	OWNERSHIPOWNERTYPE_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OwnershipOwnerType) GetAllowedValues() []OwnershipOwnerType {
	return allowedOwnershipOwnerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OwnershipOwnerType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OwnershipOwnerType(value)
	return nil
}

// NewOwnershipOwnerTypeFromValue returns a pointer to a valid OwnershipOwnerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOwnershipOwnerTypeFromValue(v string) (*OwnershipOwnerType, error) {
	ev := OwnershipOwnerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OwnershipOwnerType: valid values are %v", v, allowedOwnershipOwnerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OwnershipOwnerType) IsValid() bool {
	for _, existing := range allowedOwnershipOwnerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OwnershipOwnerType value.
func (v OwnershipOwnerType) Ptr() *OwnershipOwnerType {
	return &v
}
