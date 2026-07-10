// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CycloneDXComponentType The type of the scanned component.
type CycloneDXComponentType string

// List of CycloneDXComponentType.
const (
	CYCLONEDXCOMPONENTTYPE_LIBRARY          CycloneDXComponentType = "library"
	CYCLONEDXCOMPONENTTYPE_APPLICATION      CycloneDXComponentType = "application"
	CYCLONEDXCOMPONENTTYPE_OPERATING_SYSTEM CycloneDXComponentType = "operating-system"
)

var allowedCycloneDXComponentTypeEnumValues = []CycloneDXComponentType{
	CYCLONEDXCOMPONENTTYPE_LIBRARY,
	CYCLONEDXCOMPONENTTYPE_APPLICATION,
	CYCLONEDXCOMPONENTTYPE_OPERATING_SYSTEM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CycloneDXComponentType) GetAllowedValues() []CycloneDXComponentType {
	return allowedCycloneDXComponentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CycloneDXComponentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CycloneDXComponentType(value)
	return nil
}

// NewCycloneDXComponentTypeFromValue returns a pointer to a valid CycloneDXComponentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCycloneDXComponentTypeFromValue(v string) (*CycloneDXComponentType, error) {
	ev := CycloneDXComponentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CycloneDXComponentType: valid values are %v", v, allowedCycloneDXComponentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CycloneDXComponentType) IsValid() bool {
	for _, existing := range allowedCycloneDXComponentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CycloneDXComponentType value.
func (v CycloneDXComponentType) Ptr() *CycloneDXComponentType {
	return &v
}
