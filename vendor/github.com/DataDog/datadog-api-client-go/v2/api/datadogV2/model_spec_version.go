// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpecVersion The version of the CycloneDX specification a BOM conforms to.
type SpecVersion string

// List of SpecVersion.
const (
	SPECVERSION_ONE_ZERO  SpecVersion = "1.0"
	SPECVERSION_ONE_ONE   SpecVersion = "1.1"
	SPECVERSION_ONE_TWO   SpecVersion = "1.2"
	SPECVERSION_ONE_THREE SpecVersion = "1.3"
	SPECVERSION_ONE_FOUR  SpecVersion = "1.4"
	SPECVERSION_ONE_FIVE  SpecVersion = "1.5"
)

var allowedSpecVersionEnumValues = []SpecVersion{
	SPECVERSION_ONE_ZERO,
	SPECVERSION_ONE_ONE,
	SPECVERSION_ONE_TWO,
	SPECVERSION_ONE_THREE,
	SPECVERSION_ONE_FOUR,
	SPECVERSION_ONE_FIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpecVersion) GetAllowedValues() []SpecVersion {
	return allowedSpecVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpecVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpecVersion(value)
	return nil
}

// NewSpecVersionFromValue returns a pointer to a valid SpecVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpecVersionFromValue(v string) (*SpecVersion, error) {
	ev := SpecVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpecVersion: valid values are %v", v, allowedSpecVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpecVersion) IsValid() bool {
	for _, existing := range allowedSpecVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpecVersion value.
func (v SpecVersion) Ptr() *SpecVersion {
	return &v
}
