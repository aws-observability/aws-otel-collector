// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseBulkActionType The type of action to apply in a bulk update. Allowed values are `priority`, `status`, `assign`, `unassign`, `archive`, `unarchive`, `jira`, `servicenow`, `linear`, `update_project`.
type CaseBulkActionType string

// List of CaseBulkActionType.
const (
	CASEBULKACTIONTYPE_PRIORITY       CaseBulkActionType = "priority"
	CASEBULKACTIONTYPE_STATUS         CaseBulkActionType = "status"
	CASEBULKACTIONTYPE_ASSIGN         CaseBulkActionType = "assign"
	CASEBULKACTIONTYPE_UNASSIGN       CaseBulkActionType = "unassign"
	CASEBULKACTIONTYPE_ARCHIVE        CaseBulkActionType = "archive"
	CASEBULKACTIONTYPE_UNARCHIVE      CaseBulkActionType = "unarchive"
	CASEBULKACTIONTYPE_JIRA           CaseBulkActionType = "jira"
	CASEBULKACTIONTYPE_SERVICENOW     CaseBulkActionType = "servicenow"
	CASEBULKACTIONTYPE_LINEAR         CaseBulkActionType = "linear"
	CASEBULKACTIONTYPE_UPDATE_PROJECT CaseBulkActionType = "update_project"
)

var allowedCaseBulkActionTypeEnumValues = []CaseBulkActionType{
	CASEBULKACTIONTYPE_PRIORITY,
	CASEBULKACTIONTYPE_STATUS,
	CASEBULKACTIONTYPE_ASSIGN,
	CASEBULKACTIONTYPE_UNASSIGN,
	CASEBULKACTIONTYPE_ARCHIVE,
	CASEBULKACTIONTYPE_UNARCHIVE,
	CASEBULKACTIONTYPE_JIRA,
	CASEBULKACTIONTYPE_SERVICENOW,
	CASEBULKACTIONTYPE_LINEAR,
	CASEBULKACTIONTYPE_UPDATE_PROJECT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseBulkActionType) GetAllowedValues() []CaseBulkActionType {
	return allowedCaseBulkActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseBulkActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseBulkActionType(value)
	return nil
}

// NewCaseBulkActionTypeFromValue returns a pointer to a valid CaseBulkActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseBulkActionTypeFromValue(v string) (*CaseBulkActionType, error) {
	ev := CaseBulkActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseBulkActionType: valid values are %v", v, allowedCaseBulkActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseBulkActionType) IsValid() bool {
	for _, existing := range allowedCaseBulkActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseBulkActionType value.
func (v CaseBulkActionType) Ptr() *CaseBulkActionType {
	return &v
}
