// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultSummaryType Type of the Synthetic test result summary resource, `result_summary`.
type SyntheticsTestResultSummaryType string

// List of SyntheticsTestResultSummaryType.
const (
	SYNTHETICSTESTRESULTSUMMARYTYPE_RESULT_SUMMARY SyntheticsTestResultSummaryType = "result_summary"
)

var allowedSyntheticsTestResultSummaryTypeEnumValues = []SyntheticsTestResultSummaryType{
	SYNTHETICSTESTRESULTSUMMARYTYPE_RESULT_SUMMARY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestResultSummaryType) GetAllowedValues() []SyntheticsTestResultSummaryType {
	return allowedSyntheticsTestResultSummaryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestResultSummaryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestResultSummaryType(value)
	return nil
}

// NewSyntheticsTestResultSummaryTypeFromValue returns a pointer to a valid SyntheticsTestResultSummaryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestResultSummaryTypeFromValue(v string) (*SyntheticsTestResultSummaryType, error) {
	ev := SyntheticsTestResultSummaryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestResultSummaryType: valid values are %v", v, allowedSyntheticsTestResultSummaryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestResultSummaryType) IsValid() bool {
	for _, existing := range allowedSyntheticsTestResultSummaryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestResultSummaryType value.
func (v SyntheticsTestResultSummaryType) Ptr() *SyntheticsTestResultSummaryType {
	return &v
}
