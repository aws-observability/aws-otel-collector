// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueResourceType Jira issue resource type
type JiraIssueResourceType string

// List of JiraIssueResourceType.
const (
	JIRAISSUERESOURCETYPE_ISSUES JiraIssueResourceType = "issues"
)

var allowedJiraIssueResourceTypeEnumValues = []JiraIssueResourceType{
	JIRAISSUERESOURCETYPE_ISSUES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JiraIssueResourceType) GetAllowedValues() []JiraIssueResourceType {
	return allowedJiraIssueResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JiraIssueResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JiraIssueResourceType(value)
	return nil
}

// NewJiraIssueResourceTypeFromValue returns a pointer to a valid JiraIssueResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJiraIssueResourceTypeFromValue(v string) (*JiraIssueResourceType, error) {
	ev := JiraIssueResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JiraIssueResourceType: valid values are %v", v, allowedJiraIssueResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JiraIssueResourceType) IsValid() bool {
	for _, existing := range allowedJiraIssueResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JiraIssueResourceType value.
func (v JiraIssueResourceType) Ptr() *JiraIssueResourceType {
	return &v
}
