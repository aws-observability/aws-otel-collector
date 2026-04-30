// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueTemplateType Type identifier for Jira issue template resources
type JiraIssueTemplateType string

// List of JiraIssueTemplateType.
const (
	JIRAISSUETEMPLATETYPE_JIRA_ISSUE_TEMPLATE JiraIssueTemplateType = "jira-issue-template"
)

var allowedJiraIssueTemplateTypeEnumValues = []JiraIssueTemplateType{
	JIRAISSUETEMPLATETYPE_JIRA_ISSUE_TEMPLATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JiraIssueTemplateType) GetAllowedValues() []JiraIssueTemplateType {
	return allowedJiraIssueTemplateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JiraIssueTemplateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JiraIssueTemplateType(value)
	return nil
}

// NewJiraIssueTemplateTypeFromValue returns a pointer to a valid JiraIssueTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJiraIssueTemplateTypeFromValue(v string) (*JiraIssueTemplateType, error) {
	ev := JiraIssueTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JiraIssueTemplateType: valid values are %v", v, allowedJiraIssueTemplateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JiraIssueTemplateType) IsValid() bool {
	for _, existing := range allowedJiraIssueTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JiraIssueTemplateType value.
func (v JiraIssueTemplateType) Ptr() *JiraIssueTemplateType {
	return &v
}
