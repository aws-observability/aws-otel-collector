// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3SystemKind The definition of Entity V3 System Kind object.
type EntityV3SystemKind string

// List of EntityV3SystemKind.
const (
	ENTITYV3SYSTEMKIND_SYSTEM EntityV3SystemKind = "system"
)

var allowedEntityV3SystemKindEnumValues = []EntityV3SystemKind{
	ENTITYV3SYSTEMKIND_SYSTEM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityV3SystemKind) GetAllowedValues() []EntityV3SystemKind {
	return allowedEntityV3SystemKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityV3SystemKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityV3SystemKind(value)
	return nil
}

// NewEntityV3SystemKindFromValue returns a pointer to a valid EntityV3SystemKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityV3SystemKindFromValue(v string) (*EntityV3SystemKind, error) {
	ev := EntityV3SystemKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityV3SystemKind: valid values are %v", v, allowedEntityV3SystemKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityV3SystemKind) IsValid() bool {
	for _, existing := range allowedEntityV3SystemKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityV3SystemKind value.
func (v EntityV3SystemKind) Ptr() *EntityV3SystemKind {
	return &v
}
