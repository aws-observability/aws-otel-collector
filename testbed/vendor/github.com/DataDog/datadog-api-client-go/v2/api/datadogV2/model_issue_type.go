// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueType Type of the object.
type IssueType string

// List of IssueType.
const (
	ISSUETYPE_ISSUE IssueType = "issue"
)

var allowedIssueTypeEnumValues = []IssueType{
	ISSUETYPE_ISSUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueType) GetAllowedValues() []IssueType {
	return allowedIssueTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueType(value)
	return nil
}

// NewIssueTypeFromValue returns a pointer to a valid IssueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueTypeFromValue(v string) (*IssueType, error) {
	ev := IssueType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueType: valid values are %v", v, allowedIssueTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueType) IsValid() bool {
	for _, existing := range allowedIssueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueType value.
func (v IssueType) Ptr() *IssueType {
	return &v
}
