// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamIssuePersona Persona filter for the `issue_stream` data source.
type ListStreamIssuePersona string

// List of ListStreamIssuePersona.
const (
	LISTSTREAMISSUEPERSONA_ALL     ListStreamIssuePersona = "all"
	LISTSTREAMISSUEPERSONA_BROWSER ListStreamIssuePersona = "browser"
	LISTSTREAMISSUEPERSONA_MOBILE  ListStreamIssuePersona = "mobile"
	LISTSTREAMISSUEPERSONA_BACKEND ListStreamIssuePersona = "backend"
)

var allowedListStreamIssuePersonaEnumValues = []ListStreamIssuePersona{
	LISTSTREAMISSUEPERSONA_ALL,
	LISTSTREAMISSUEPERSONA_BROWSER,
	LISTSTREAMISSUEPERSONA_MOBILE,
	LISTSTREAMISSUEPERSONA_BACKEND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListStreamIssuePersona) GetAllowedValues() []ListStreamIssuePersona {
	return allowedListStreamIssuePersonaEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListStreamIssuePersona) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListStreamIssuePersona(value)
	return nil
}

// NewListStreamIssuePersonaFromValue returns a pointer to a valid ListStreamIssuePersona
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListStreamIssuePersonaFromValue(v string) (*ListStreamIssuePersona, error) {
	ev := ListStreamIssuePersona(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListStreamIssuePersona: valid values are %v", v, allowedListStreamIssuePersonaEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListStreamIssuePersona) IsValid() bool {
	for _, existing := range allowedListStreamIssuePersonaEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListStreamIssuePersona value.
func (v ListStreamIssuePersona) Ptr() *ListStreamIssuePersona {
	return &v
}
