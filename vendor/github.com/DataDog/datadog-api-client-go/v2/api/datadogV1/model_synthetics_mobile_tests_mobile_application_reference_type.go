// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTestsMobileApplicationReferenceType Reference type for the mobile application for a mobile synthetics test.
type SyntheticsMobileTestsMobileApplicationReferenceType string

// List of SyntheticsMobileTestsMobileApplicationReferenceType.
const (
	SYNTHETICSMOBILETESTSMOBILEAPPLICATIONREFERENCETYPE_LATEST  SyntheticsMobileTestsMobileApplicationReferenceType = "latest"
	SYNTHETICSMOBILETESTSMOBILEAPPLICATIONREFERENCETYPE_VERSION SyntheticsMobileTestsMobileApplicationReferenceType = "version"
)

var allowedSyntheticsMobileTestsMobileApplicationReferenceTypeEnumValues = []SyntheticsMobileTestsMobileApplicationReferenceType{
	SYNTHETICSMOBILETESTSMOBILEAPPLICATIONREFERENCETYPE_LATEST,
	SYNTHETICSMOBILETESTSMOBILEAPPLICATIONREFERENCETYPE_VERSION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsMobileTestsMobileApplicationReferenceType) GetAllowedValues() []SyntheticsMobileTestsMobileApplicationReferenceType {
	return allowedSyntheticsMobileTestsMobileApplicationReferenceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsMobileTestsMobileApplicationReferenceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsMobileTestsMobileApplicationReferenceType(value)
	return nil
}

// NewSyntheticsMobileTestsMobileApplicationReferenceTypeFromValue returns a pointer to a valid SyntheticsMobileTestsMobileApplicationReferenceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsMobileTestsMobileApplicationReferenceTypeFromValue(v string) (*SyntheticsMobileTestsMobileApplicationReferenceType, error) {
	ev := SyntheticsMobileTestsMobileApplicationReferenceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsMobileTestsMobileApplicationReferenceType: valid values are %v", v, allowedSyntheticsMobileTestsMobileApplicationReferenceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsMobileTestsMobileApplicationReferenceType) IsValid() bool {
	for _, existing := range allowedSyntheticsMobileTestsMobileApplicationReferenceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsMobileTestsMobileApplicationReferenceType value.
func (v SyntheticsMobileTestsMobileApplicationReferenceType) Ptr() *SyntheticsMobileTestsMobileApplicationReferenceType {
	return &v
}
