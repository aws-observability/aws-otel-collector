// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DetailedFindingType The JSON:API type for findings that have the message and resource configuration.
type DetailedFindingType string

// List of DetailedFindingType.
const (
	DETAILEDFINDINGTYPE_DETAILED_FINDING DetailedFindingType = "detailed_finding"
)

var allowedDetailedFindingTypeEnumValues = []DetailedFindingType{
	DETAILEDFINDINGTYPE_DETAILED_FINDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DetailedFindingType) GetAllowedValues() []DetailedFindingType {
	return allowedDetailedFindingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DetailedFindingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DetailedFindingType(value)
	return nil
}

// NewDetailedFindingTypeFromValue returns a pointer to a valid DetailedFindingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDetailedFindingTypeFromValue(v string) (*DetailedFindingType, error) {
	ev := DetailedFindingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DetailedFindingType: valid values are %v", v, allowedDetailedFindingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DetailedFindingType) IsValid() bool {
	for _, existing := range allowedDetailedFindingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DetailedFindingType value.
func (v DetailedFindingType) Ptr() *DetailedFindingType {
	return &v
}
