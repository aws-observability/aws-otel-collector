// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsSuiteType Type of the Synthetic suite, `suite`.
type SyntheticsSuiteType string

// List of SyntheticsSuiteType.
const (
	SYNTHETICSSUITETYPE_SUITE SyntheticsSuiteType = "suite"
)

var allowedSyntheticsSuiteTypeEnumValues = []SyntheticsSuiteType{
	SYNTHETICSSUITETYPE_SUITE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsSuiteType) GetAllowedValues() []SyntheticsSuiteType {
	return allowedSyntheticsSuiteTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsSuiteType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsSuiteType(value)
	return nil
}

// NewSyntheticsSuiteTypeFromValue returns a pointer to a valid SyntheticsSuiteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsSuiteTypeFromValue(v string) (*SyntheticsSuiteType, error) {
	ev := SyntheticsSuiteType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsSuiteType: valid values are %v", v, allowedSyntheticsSuiteTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsSuiteType) IsValid() bool {
	for _, existing := range allowedSyntheticsSuiteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsSuiteType value.
func (v SyntheticsSuiteType) Ptr() *SyntheticsSuiteType {
	return &v
}
