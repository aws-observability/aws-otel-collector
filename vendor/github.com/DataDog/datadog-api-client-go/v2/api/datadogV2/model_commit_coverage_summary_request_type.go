// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitCoverageSummaryRequestType JSON:API type for commit coverage summary request. The value must always be `ci_app_coverage_commit_summary_request`.
type CommitCoverageSummaryRequestType string

// List of CommitCoverageSummaryRequestType.
const (
	COMMITCOVERAGESUMMARYREQUESTTYPE_CI_APP_COVERAGE_COMMIT_SUMMARY_REQUEST CommitCoverageSummaryRequestType = "ci_app_coverage_commit_summary_request"
)

var allowedCommitCoverageSummaryRequestTypeEnumValues = []CommitCoverageSummaryRequestType{
	COMMITCOVERAGESUMMARYREQUESTTYPE_CI_APP_COVERAGE_COMMIT_SUMMARY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CommitCoverageSummaryRequestType) GetAllowedValues() []CommitCoverageSummaryRequestType {
	return allowedCommitCoverageSummaryRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CommitCoverageSummaryRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CommitCoverageSummaryRequestType(value)
	return nil
}

// NewCommitCoverageSummaryRequestTypeFromValue returns a pointer to a valid CommitCoverageSummaryRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCommitCoverageSummaryRequestTypeFromValue(v string) (*CommitCoverageSummaryRequestType, error) {
	ev := CommitCoverageSummaryRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CommitCoverageSummaryRequestType: valid values are %v", v, allowedCommitCoverageSummaryRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CommitCoverageSummaryRequestType) IsValid() bool {
	for _, existing := range allowedCommitCoverageSummaryRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommitCoverageSummaryRequestType value.
func (v CommitCoverageSummaryRequestType) Ptr() *CommitCoverageSummaryRequestType {
	return &v
}
