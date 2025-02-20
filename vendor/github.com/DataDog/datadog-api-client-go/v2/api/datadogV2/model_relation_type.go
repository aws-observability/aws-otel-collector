// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RelationType Supported relation types.
type RelationType string

// List of RelationType.
const (
	RELATIONTYPE_RELATIONTYPEOWNS          RelationType = "RelationTypeOwns"
	RELATIONTYPE_RELATIONTYPEOWNEDBY       RelationType = "RelationTypeOwnedBy"
	RELATIONTYPE_RELATIONTYPEDEPENDSON     RelationType = "RelationTypeDependsOn"
	RELATIONTYPE_RELATIONTYPEDEPENDENCYOF  RelationType = "RelationTypeDependencyOf"
	RELATIONTYPE_RELATIONTYPEPARTSOF       RelationType = "RelationTypePartsOf"
	RELATIONTYPE_RELATIONTYPEHASPART       RelationType = "RelationTypeHasPart"
	RELATIONTYPE_RELATIONTYPEOTHEROWNS     RelationType = "RelationTypeOtherOwns"
	RELATIONTYPE_RELATIONTYPEOTHEROWNEDBY  RelationType = "RelationTypeOtherOwnedBy"
	RELATIONTYPE_RELATIONTYPEIMPLEMENTEDBY RelationType = "RelationTypeImplementedBy"
	RELATIONTYPE_RELATIONTYPEIMPLEMENTS    RelationType = "RelationTypeImplements"
)

var allowedRelationTypeEnumValues = []RelationType{
	RELATIONTYPE_RELATIONTYPEOWNS,
	RELATIONTYPE_RELATIONTYPEOWNEDBY,
	RELATIONTYPE_RELATIONTYPEDEPENDSON,
	RELATIONTYPE_RELATIONTYPEDEPENDENCYOF,
	RELATIONTYPE_RELATIONTYPEPARTSOF,
	RELATIONTYPE_RELATIONTYPEHASPART,
	RELATIONTYPE_RELATIONTYPEOTHEROWNS,
	RELATIONTYPE_RELATIONTYPEOTHEROWNEDBY,
	RELATIONTYPE_RELATIONTYPEIMPLEMENTEDBY,
	RELATIONTYPE_RELATIONTYPEIMPLEMENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RelationType) GetAllowedValues() []RelationType {
	return allowedRelationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RelationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RelationType(value)
	return nil
}

// NewRelationTypeFromValue returns a pointer to a valid RelationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRelationTypeFromValue(v string) (*RelationType, error) {
	ev := RelationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RelationType: valid values are %v", v, allowedRelationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RelationType) IsValid() bool {
	for _, existing := range allowedRelationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RelationType value.
func (v RelationType) Ptr() *RelationType {
	return &v
}
