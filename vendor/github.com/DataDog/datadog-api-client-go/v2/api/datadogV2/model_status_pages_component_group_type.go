// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesComponentGroupType Components resource type.
type StatusPagesComponentGroupType string

// List of StatusPagesComponentGroupType.
const (
	STATUSPAGESCOMPONENTGROUPTYPE_COMPONENTS StatusPagesComponentGroupType = "components"
)

var allowedStatusPagesComponentGroupTypeEnumValues = []StatusPagesComponentGroupType{
	STATUSPAGESCOMPONENTGROUPTYPE_COMPONENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StatusPagesComponentGroupType) GetAllowedValues() []StatusPagesComponentGroupType {
	return allowedStatusPagesComponentGroupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StatusPagesComponentGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StatusPagesComponentGroupType(value)
	return nil
}

// NewStatusPagesComponentGroupTypeFromValue returns a pointer to a valid StatusPagesComponentGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStatusPagesComponentGroupTypeFromValue(v string) (*StatusPagesComponentGroupType, error) {
	ev := StatusPagesComponentGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StatusPagesComponentGroupType: valid values are %v", v, allowedStatusPagesComponentGroupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StatusPagesComponentGroupType) IsValid() bool {
	for _, existing := range allowedStatusPagesComponentGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusPagesComponentGroupType value.
func (v StatusPagesComponentGroupType) Ptr() *StatusPagesComponentGroupType {
	return &v
}
