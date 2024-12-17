// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3DatastoreKind The definition of Entity V3 Datastore Kind object.
type EntityV3DatastoreKind string

// List of EntityV3DatastoreKind.
const (
	ENTITYV3DATASTOREKIND_DATASTORE EntityV3DatastoreKind = "datastore"
)

var allowedEntityV3DatastoreKindEnumValues = []EntityV3DatastoreKind{
	ENTITYV3DATASTOREKIND_DATASTORE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityV3DatastoreKind) GetAllowedValues() []EntityV3DatastoreKind {
	return allowedEntityV3DatastoreKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityV3DatastoreKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityV3DatastoreKind(value)
	return nil
}

// NewEntityV3DatastoreKindFromValue returns a pointer to a valid EntityV3DatastoreKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityV3DatastoreKindFromValue(v string) (*EntityV3DatastoreKind, error) {
	ev := EntityV3DatastoreKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityV3DatastoreKind: valid values are %v", v, allowedEntityV3DatastoreKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityV3DatastoreKind) IsValid() bool {
	for _, existing := range allowedEntityV3DatastoreKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityV3DatastoreKind value.
func (v EntityV3DatastoreKind) Ptr() *EntityV3DatastoreKind {
	return &v
}
