// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DegradationDataAttributesSourceType The type of the source.
type DegradationDataAttributesSourceType string

// List of DegradationDataAttributesSourceType.
const (
	DEGRADATIONDATAATTRIBUTESSOURCETYPE_INCIDENT DegradationDataAttributesSourceType = "incident"
)

var allowedDegradationDataAttributesSourceTypeEnumValues = []DegradationDataAttributesSourceType{
	DEGRADATIONDATAATTRIBUTESSOURCETYPE_INCIDENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DegradationDataAttributesSourceType) GetAllowedValues() []DegradationDataAttributesSourceType {
	return allowedDegradationDataAttributesSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DegradationDataAttributesSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DegradationDataAttributesSourceType(value)
	return nil
}

// NewDegradationDataAttributesSourceTypeFromValue returns a pointer to a valid DegradationDataAttributesSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDegradationDataAttributesSourceTypeFromValue(v string) (*DegradationDataAttributesSourceType, error) {
	ev := DegradationDataAttributesSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DegradationDataAttributesSourceType: valid values are %v", v, allowedDegradationDataAttributesSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DegradationDataAttributesSourceType) IsValid() bool {
	for _, existing := range allowedDegradationDataAttributesSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DegradationDataAttributesSourceType value.
func (v DegradationDataAttributesSourceType) Ptr() *DegradationDataAttributesSourceType {
	return &v
}
