// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssuesDataType Jira issues resource type.
type JiraIssuesDataType string

// List of JiraIssuesDataType.
const (
	JIRAISSUESDATATYPE_JIRA_ISSUES JiraIssuesDataType = "jira_issues"
)

var allowedJiraIssuesDataTypeEnumValues = []JiraIssuesDataType{
	JIRAISSUESDATATYPE_JIRA_ISSUES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JiraIssuesDataType) GetAllowedValues() []JiraIssuesDataType {
	return allowedJiraIssuesDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JiraIssuesDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JiraIssuesDataType(value)
	return nil
}

// NewJiraIssuesDataTypeFromValue returns a pointer to a valid JiraIssuesDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJiraIssuesDataTypeFromValue(v string) (*JiraIssuesDataType, error) {
	ev := JiraIssuesDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JiraIssuesDataType: valid values are %v", v, allowedJiraIssuesDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JiraIssuesDataType) IsValid() bool {
	for _, existing := range allowedJiraIssuesDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JiraIssuesDataType value.
func (v JiraIssuesDataType) Ptr() *JiraIssuesDataType {
	return &v
}
