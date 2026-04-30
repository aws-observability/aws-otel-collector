// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CoverageSummaryType JSON:API type for coverage summary response. The value must always be `ci_app_coverage_summary`.
type CoverageSummaryType string

// List of CoverageSummaryType.
const (
	COVERAGESUMMARYTYPE_CI_APP_COVERAGE_SUMMARY CoverageSummaryType = "ci_app_coverage_summary"
)

var allowedCoverageSummaryTypeEnumValues = []CoverageSummaryType{
	COVERAGESUMMARYTYPE_CI_APP_COVERAGE_SUMMARY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CoverageSummaryType) GetAllowedValues() []CoverageSummaryType {
	return allowedCoverageSummaryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CoverageSummaryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CoverageSummaryType(value)
	return nil
}

// NewCoverageSummaryTypeFromValue returns a pointer to a valid CoverageSummaryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCoverageSummaryTypeFromValue(v string) (*CoverageSummaryType, error) {
	ev := CoverageSummaryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CoverageSummaryType: valid values are %v", v, allowedCoverageSummaryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CoverageSummaryType) IsValid() bool {
	for _, existing := range allowedCoverageSummaryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CoverageSummaryType value.
func (v CoverageSummaryType) Ptr() *CoverageSummaryType {
	return &v
}
