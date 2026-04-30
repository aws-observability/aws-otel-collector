// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SuiteSearchResponseType Type for the Synthetics suites search response, `suites_search`.
type SuiteSearchResponseType string

// List of SuiteSearchResponseType.
const (
	SUITESEARCHRESPONSETYPE_SUITES_SEARCH SuiteSearchResponseType = "suites_search"
)

var allowedSuiteSearchResponseTypeEnumValues = []SuiteSearchResponseType{
	SUITESEARCHRESPONSETYPE_SUITES_SEARCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SuiteSearchResponseType) GetAllowedValues() []SuiteSearchResponseType {
	return allowedSuiteSearchResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SuiteSearchResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SuiteSearchResponseType(value)
	return nil
}

// NewSuiteSearchResponseTypeFromValue returns a pointer to a valid SuiteSearchResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSuiteSearchResponseTypeFromValue(v string) (*SuiteSearchResponseType, error) {
	ev := SuiteSearchResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SuiteSearchResponseType: valid values are %v", v, allowedSuiteSearchResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SuiteSearchResponseType) IsValid() bool {
	for _, existing := range allowedSuiteSearchResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SuiteSearchResponseType value.
func (v SuiteSearchResponseType) Ptr() *SuiteSearchResponseType {
	return &v
}
