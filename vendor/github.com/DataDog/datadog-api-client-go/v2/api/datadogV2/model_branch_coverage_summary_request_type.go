// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BranchCoverageSummaryRequestType JSON:API type for branch coverage summary request. The value must always be `ci_app_coverage_branch_summary_request`.
type BranchCoverageSummaryRequestType string

// List of BranchCoverageSummaryRequestType.
const (
	BRANCHCOVERAGESUMMARYREQUESTTYPE_CI_APP_COVERAGE_BRANCH_SUMMARY_REQUEST BranchCoverageSummaryRequestType = "ci_app_coverage_branch_summary_request"
)

var allowedBranchCoverageSummaryRequestTypeEnumValues = []BranchCoverageSummaryRequestType{
	BRANCHCOVERAGESUMMARYREQUESTTYPE_CI_APP_COVERAGE_BRANCH_SUMMARY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BranchCoverageSummaryRequestType) GetAllowedValues() []BranchCoverageSummaryRequestType {
	return allowedBranchCoverageSummaryRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BranchCoverageSummaryRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BranchCoverageSummaryRequestType(value)
	return nil
}

// NewBranchCoverageSummaryRequestTypeFromValue returns a pointer to a valid BranchCoverageSummaryRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBranchCoverageSummaryRequestTypeFromValue(v string) (*BranchCoverageSummaryRequestType, error) {
	ev := BranchCoverageSummaryRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BranchCoverageSummaryRequestType: valid values are %v", v, allowedBranchCoverageSummaryRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BranchCoverageSummaryRequestType) IsValid() bool {
	for _, existing := range allowedBranchCoverageSummaryRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BranchCoverageSummaryRequestType value.
func (v BranchCoverageSummaryRequestType) Ptr() *BranchCoverageSummaryRequestType {
	return &v
}
