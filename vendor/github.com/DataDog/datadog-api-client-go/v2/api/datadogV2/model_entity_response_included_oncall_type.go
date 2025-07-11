// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedOncallType Oncall type.
type EntityResponseIncludedOncallType string

// List of EntityResponseIncludedOncallType.
const (
	ENTITYRESPONSEINCLUDEDONCALLTYPE_ONCALL EntityResponseIncludedOncallType = "oncall"
)

var allowedEntityResponseIncludedOncallTypeEnumValues = []EntityResponseIncludedOncallType{
	ENTITYRESPONSEINCLUDEDONCALLTYPE_ONCALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityResponseIncludedOncallType) GetAllowedValues() []EntityResponseIncludedOncallType {
	return allowedEntityResponseIncludedOncallTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityResponseIncludedOncallType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityResponseIncludedOncallType(value)
	return nil
}

// NewEntityResponseIncludedOncallTypeFromValue returns a pointer to a valid EntityResponseIncludedOncallType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityResponseIncludedOncallTypeFromValue(v string) (*EntityResponseIncludedOncallType, error) {
	ev := EntityResponseIncludedOncallType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityResponseIncludedOncallType: valid values are %v", v, allowedEntityResponseIncludedOncallTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityResponseIncludedOncallType) IsValid() bool {
	for _, existing := range allowedEntityResponseIncludedOncallTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityResponseIncludedOncallType value.
func (v EntityResponseIncludedOncallType) Ptr() *EntityResponseIncludedOncallType {
	return &v
}
