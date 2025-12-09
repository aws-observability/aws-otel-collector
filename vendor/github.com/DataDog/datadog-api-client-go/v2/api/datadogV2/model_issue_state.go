// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueState State of the issue
type IssueState string

// List of IssueState.
const (
	ISSUESTATE_OPEN         IssueState = "OPEN"
	ISSUESTATE_ACKNOWLEDGED IssueState = "ACKNOWLEDGED"
	ISSUESTATE_RESOLVED     IssueState = "RESOLVED"
	ISSUESTATE_IGNORED      IssueState = "IGNORED"
	ISSUESTATE_EXCLUDED     IssueState = "EXCLUDED"
)

var allowedIssueStateEnumValues = []IssueState{
	ISSUESTATE_OPEN,
	ISSUESTATE_ACKNOWLEDGED,
	ISSUESTATE_RESOLVED,
	ISSUESTATE_IGNORED,
	ISSUESTATE_EXCLUDED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueState) GetAllowedValues() []IssueState {
	return allowedIssueStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueState(value)
	return nil
}

// NewIssueStateFromValue returns a pointer to a valid IssueState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueStateFromValue(v string) (*IssueState, error) {
	ev := IssueState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueState: valid values are %v", v, allowedIssueStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueState) IsValid() bool {
	for _, existing := range allowedIssueStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueState value.
func (v IssueState) Ptr() *IssueState {
	return &v
}
