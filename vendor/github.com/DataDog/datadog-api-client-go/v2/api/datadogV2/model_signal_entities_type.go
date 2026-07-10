// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SignalEntitiesType The type of the resource. The value should always be `entities`.
type SignalEntitiesType string

// List of SignalEntitiesType.
const (
	SIGNALENTITIESTYPE_ENTITIES SignalEntitiesType = "entities"
)

var allowedSignalEntitiesTypeEnumValues = []SignalEntitiesType{
	SIGNALENTITIESTYPE_ENTITIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SignalEntitiesType) GetAllowedValues() []SignalEntitiesType {
	return allowedSignalEntitiesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SignalEntitiesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SignalEntitiesType(value)
	return nil
}

// NewSignalEntitiesTypeFromValue returns a pointer to a valid SignalEntitiesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSignalEntitiesTypeFromValue(v string) (*SignalEntitiesType, error) {
	ev := SignalEntitiesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SignalEntitiesType: valid values are %v", v, allowedSignalEntitiesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SignalEntitiesType) IsValid() bool {
	for _, existing := range allowedSignalEntitiesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SignalEntitiesType value.
func (v SignalEntitiesType) Ptr() *SignalEntitiesType {
	return &v
}
