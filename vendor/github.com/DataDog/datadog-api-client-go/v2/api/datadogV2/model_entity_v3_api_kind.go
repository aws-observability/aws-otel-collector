// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3APIKind The definition of Entity V3 API Kind object.
type EntityV3APIKind string

// List of EntityV3APIKind.
const (
	ENTITYV3APIKIND_API EntityV3APIKind = "api"
)

var allowedEntityV3APIKindEnumValues = []EntityV3APIKind{
	ENTITYV3APIKIND_API,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityV3APIKind) GetAllowedValues() []EntityV3APIKind {
	return allowedEntityV3APIKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityV3APIKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityV3APIKind(value)
	return nil
}

// NewEntityV3APIKindFromValue returns a pointer to a valid EntityV3APIKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityV3APIKindFromValue(v string) (*EntityV3APIKind, error) {
	ev := EntityV3APIKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityV3APIKind: valid values are %v", v, allowedEntityV3APIKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityV3APIKind) IsValid() bool {
	for _, existing := range allowedEntityV3APIKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityV3APIKind value.
func (v EntityV3APIKind) Ptr() *EntityV3APIKind {
	return &v
}
