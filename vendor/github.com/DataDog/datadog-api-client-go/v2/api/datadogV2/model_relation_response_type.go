// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RelationResponseType Relation type.
type RelationResponseType string

// List of RelationResponseType.
const (
	RELATIONRESPONSETYPE_RELATION RelationResponseType = "relation"
)

var allowedRelationResponseTypeEnumValues = []RelationResponseType{
	RELATIONRESPONSETYPE_RELATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RelationResponseType) GetAllowedValues() []RelationResponseType {
	return allowedRelationResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RelationResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RelationResponseType(value)
	return nil
}

// NewRelationResponseTypeFromValue returns a pointer to a valid RelationResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRelationResponseTypeFromValue(v string) (*RelationResponseType, error) {
	ev := RelationResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RelationResponseType: valid values are %v", v, allowedRelationResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RelationResponseType) IsValid() bool {
	for _, existing := range allowedRelationResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RelationResponseType value.
func (v RelationResponseType) Ptr() *RelationResponseType {
	return &v
}
