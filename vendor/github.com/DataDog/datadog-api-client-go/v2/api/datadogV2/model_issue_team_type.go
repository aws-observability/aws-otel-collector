// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueTeamType Type of the object.
type IssueTeamType string

// List of IssueTeamType.
const (
	ISSUETEAMTYPE_TEAM IssueTeamType = "team"
)

var allowedIssueTeamTypeEnumValues = []IssueTeamType{
	ISSUETEAMTYPE_TEAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueTeamType) GetAllowedValues() []IssueTeamType {
	return allowedIssueTeamTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueTeamType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueTeamType(value)
	return nil
}

// NewIssueTeamTypeFromValue returns a pointer to a valid IssueTeamType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueTeamTypeFromValue(v string) (*IssueTeamType, error) {
	ev := IssueTeamType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueTeamType: valid values are %v", v, allowedIssueTeamTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueTeamType) IsValid() bool {
	for _, existing := range allowedIssueTeamTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueTeamType value.
func (v IssueTeamType) Ptr() *IssueTeamType {
	return &v
}
