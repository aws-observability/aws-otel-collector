// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3ServiceKind The definition of Entity V3 Service Kind object.
type EntityV3ServiceKind string

// List of EntityV3ServiceKind.
const (
	ENTITYV3SERVICEKIND_SERVICE EntityV3ServiceKind = "service"
)

var allowedEntityV3ServiceKindEnumValues = []EntityV3ServiceKind{
	ENTITYV3SERVICEKIND_SERVICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityV3ServiceKind) GetAllowedValues() []EntityV3ServiceKind {
	return allowedEntityV3ServiceKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityV3ServiceKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityV3ServiceKind(value)
	return nil
}

// NewEntityV3ServiceKindFromValue returns a pointer to a valid EntityV3ServiceKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityV3ServiceKindFromValue(v string) (*EntityV3ServiceKind, error) {
	ev := EntityV3ServiceKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityV3ServiceKind: valid values are %v", v, allowedEntityV3ServiceKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityV3ServiceKind) IsValid() bool {
	for _, existing := range allowedEntityV3ServiceKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityV3ServiceKind value.
func (v EntityV3ServiceKind) Ptr() *EntityV3ServiceKind {
	return &v
}
