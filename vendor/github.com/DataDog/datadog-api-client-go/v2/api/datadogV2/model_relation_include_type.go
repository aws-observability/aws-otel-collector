// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RelationIncludeType Supported include types for relations.
type RelationIncludeType string

// List of RelationIncludeType.
const (
	RELATIONINCLUDETYPE_ENTITY RelationIncludeType = "entity"
	RELATIONINCLUDETYPE_SCHEMA RelationIncludeType = "schema"
)

var allowedRelationIncludeTypeEnumValues = []RelationIncludeType{
	RELATIONINCLUDETYPE_ENTITY,
	RELATIONINCLUDETYPE_SCHEMA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RelationIncludeType) GetAllowedValues() []RelationIncludeType {
	return allowedRelationIncludeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RelationIncludeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RelationIncludeType(value)
	return nil
}

// NewRelationIncludeTypeFromValue returns a pointer to a valid RelationIncludeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRelationIncludeTypeFromValue(v string) (*RelationIncludeType, error) {
	ev := RelationIncludeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RelationIncludeType: valid values are %v", v, allowedRelationIncludeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RelationIncludeType) IsValid() bool {
	for _, existing := range allowedRelationIncludeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RelationIncludeType value.
func (v RelationIncludeType) Ptr() *RelationIncludeType {
	return &v
}
