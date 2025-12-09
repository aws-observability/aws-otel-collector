// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetIssueIncludeQueryParameterItem Relationship object that should be included in the response.
type GetIssueIncludeQueryParameterItem string

// List of GetIssueIncludeQueryParameterItem.
const (
	GETISSUEINCLUDEQUERYPARAMETERITEM_ASSIGNEE    GetIssueIncludeQueryParameterItem = "assignee"
	GETISSUEINCLUDEQUERYPARAMETERITEM_CASE        GetIssueIncludeQueryParameterItem = "case"
	GETISSUEINCLUDEQUERYPARAMETERITEM_TEAM_OWNERS GetIssueIncludeQueryParameterItem = "team_owners"
)

var allowedGetIssueIncludeQueryParameterItemEnumValues = []GetIssueIncludeQueryParameterItem{
	GETISSUEINCLUDEQUERYPARAMETERITEM_ASSIGNEE,
	GETISSUEINCLUDEQUERYPARAMETERITEM_CASE,
	GETISSUEINCLUDEQUERYPARAMETERITEM_TEAM_OWNERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetIssueIncludeQueryParameterItem) GetAllowedValues() []GetIssueIncludeQueryParameterItem {
	return allowedGetIssueIncludeQueryParameterItemEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetIssueIncludeQueryParameterItem) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetIssueIncludeQueryParameterItem(value)
	return nil
}

// NewGetIssueIncludeQueryParameterItemFromValue returns a pointer to a valid GetIssueIncludeQueryParameterItem
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetIssueIncludeQueryParameterItemFromValue(v string) (*GetIssueIncludeQueryParameterItem, error) {
	ev := GetIssueIncludeQueryParameterItem(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetIssueIncludeQueryParameterItem: valid values are %v", v, allowedGetIssueIncludeQueryParameterItemEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetIssueIncludeQueryParameterItem) IsValid() bool {
	for _, existing := range allowedGetIssueIncludeQueryParameterItemEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetIssueIncludeQueryParameterItem value.
func (v GetIssueIncludeQueryParameterItem) Ptr() *GetIssueIncludeQueryParameterItem {
	return &v
}
