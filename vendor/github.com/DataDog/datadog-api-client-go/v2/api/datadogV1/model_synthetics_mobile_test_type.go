// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTestType Type of the Synthetic test, `mobile`.
type SyntheticsMobileTestType string

// List of SyntheticsMobileTestType.
const (
	SYNTHETICSMOBILETESTTYPE_MOBILE SyntheticsMobileTestType = "mobile"
)

var allowedSyntheticsMobileTestTypeEnumValues = []SyntheticsMobileTestType{
	SYNTHETICSMOBILETESTTYPE_MOBILE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsMobileTestType) GetAllowedValues() []SyntheticsMobileTestType {
	return allowedSyntheticsMobileTestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsMobileTestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsMobileTestType(value)
	return nil
}

// NewSyntheticsMobileTestTypeFromValue returns a pointer to a valid SyntheticsMobileTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsMobileTestTypeFromValue(v string) (*SyntheticsMobileTestType, error) {
	ev := SyntheticsMobileTestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsMobileTestType: valid values are %v", v, allowedSyntheticsMobileTestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsMobileTestType) IsValid() bool {
	for _, existing := range allowedSyntheticsMobileTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsMobileTestType value.
func (v SyntheticsMobileTestType) Ptr() *SyntheticsMobileTestType {
	return &v
}
