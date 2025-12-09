// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SearchIssuesIncludeQueryParameterItem Relationship object that should be included in the search response.
type SearchIssuesIncludeQueryParameterItem string

// List of SearchIssuesIncludeQueryParameterItem.
const (
	SEARCHISSUESINCLUDEQUERYPARAMETERITEM_ISSUE             SearchIssuesIncludeQueryParameterItem = "issue"
	SEARCHISSUESINCLUDEQUERYPARAMETERITEM_ISSUE_ASSIGNEE    SearchIssuesIncludeQueryParameterItem = "issue.assignee"
	SEARCHISSUESINCLUDEQUERYPARAMETERITEM_ISSUE_CASE        SearchIssuesIncludeQueryParameterItem = "issue.case"
	SEARCHISSUESINCLUDEQUERYPARAMETERITEM_ISSUE_TEAM_OWNERS SearchIssuesIncludeQueryParameterItem = "issue.team_owners"
)

var allowedSearchIssuesIncludeQueryParameterItemEnumValues = []SearchIssuesIncludeQueryParameterItem{
	SEARCHISSUESINCLUDEQUERYPARAMETERITEM_ISSUE,
	SEARCHISSUESINCLUDEQUERYPARAMETERITEM_ISSUE_ASSIGNEE,
	SEARCHISSUESINCLUDEQUERYPARAMETERITEM_ISSUE_CASE,
	SEARCHISSUESINCLUDEQUERYPARAMETERITEM_ISSUE_TEAM_OWNERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SearchIssuesIncludeQueryParameterItem) GetAllowedValues() []SearchIssuesIncludeQueryParameterItem {
	return allowedSearchIssuesIncludeQueryParameterItemEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SearchIssuesIncludeQueryParameterItem) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SearchIssuesIncludeQueryParameterItem(value)
	return nil
}

// NewSearchIssuesIncludeQueryParameterItemFromValue returns a pointer to a valid SearchIssuesIncludeQueryParameterItem
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSearchIssuesIncludeQueryParameterItemFromValue(v string) (*SearchIssuesIncludeQueryParameterItem, error) {
	ev := SearchIssuesIncludeQueryParameterItem(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SearchIssuesIncludeQueryParameterItem: valid values are %v", v, allowedSearchIssuesIncludeQueryParameterItemEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SearchIssuesIncludeQueryParameterItem) IsValid() bool {
	for _, existing := range allowedSearchIssuesIncludeQueryParameterItemEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SearchIssuesIncludeQueryParameterItem value.
func (v SearchIssuesIncludeQueryParameterItem) Ptr() *SearchIssuesIncludeQueryParameterItem {
	return &v
}
