// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsSuiteTypes Type for the Synthetics suites responses, `suites`.
type SyntheticsSuiteTypes string

// List of SyntheticsSuiteTypes.
const (
	SYNTHETICSSUITETYPES_SUITES SyntheticsSuiteTypes = "suites"
)

var allowedSyntheticsSuiteTypesEnumValues = []SyntheticsSuiteTypes{
	SYNTHETICSSUITETYPES_SUITES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsSuiteTypes) GetAllowedValues() []SyntheticsSuiteTypes {
	return allowedSyntheticsSuiteTypesEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsSuiteTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsSuiteTypes(value)
	return nil
}

// NewSyntheticsSuiteTypesFromValue returns a pointer to a valid SyntheticsSuiteTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsSuiteTypesFromValue(v string) (*SyntheticsSuiteTypes, error) {
	ev := SyntheticsSuiteTypes(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsSuiteTypes: valid values are %v", v, allowedSyntheticsSuiteTypesEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsSuiteTypes) IsValid() bool {
	for _, existing := range allowedSyntheticsSuiteTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsSuiteTypes value.
func (v SyntheticsSuiteTypes) Ptr() *SyntheticsSuiteTypes {
	return &v
}
