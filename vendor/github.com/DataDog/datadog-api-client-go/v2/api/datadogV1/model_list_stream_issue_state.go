// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamIssueState Issue state filter for the `issue_stream` data source.
type ListStreamIssueState string

// List of ListStreamIssueState.
const (
	LISTSTREAMISSUESTATE_OPEN         ListStreamIssueState = "OPEN"
	LISTSTREAMISSUESTATE_IGNORED      ListStreamIssueState = "IGNORED"
	LISTSTREAMISSUESTATE_ACKNOWLEDGED ListStreamIssueState = "ACKNOWLEDGED"
	LISTSTREAMISSUESTATE_RESOLVED     ListStreamIssueState = "RESOLVED"
)

var allowedListStreamIssueStateEnumValues = []ListStreamIssueState{
	LISTSTREAMISSUESTATE_OPEN,
	LISTSTREAMISSUESTATE_IGNORED,
	LISTSTREAMISSUESTATE_ACKNOWLEDGED,
	LISTSTREAMISSUESTATE_RESOLVED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListStreamIssueState) GetAllowedValues() []ListStreamIssueState {
	return allowedListStreamIssueStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListStreamIssueState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListStreamIssueState(value)
	return nil
}

// NewListStreamIssueStateFromValue returns a pointer to a valid ListStreamIssueState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListStreamIssueStateFromValue(v string) (*ListStreamIssueState, error) {
	ev := ListStreamIssueState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListStreamIssueState: valid values are %v", v, allowedListStreamIssueStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListStreamIssueState) IsValid() bool {
	for _, existing := range allowedListStreamIssueStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListStreamIssueState value.
func (v ListStreamIssueState) Ptr() *ListStreamIssueState {
	return &v
}
