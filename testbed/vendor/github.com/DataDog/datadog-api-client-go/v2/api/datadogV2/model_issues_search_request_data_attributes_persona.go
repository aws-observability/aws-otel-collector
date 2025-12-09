// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssuesSearchRequestDataAttributesPersona Persona for the search. Either track(s) or persona(s) must be specified.
type IssuesSearchRequestDataAttributesPersona string

// List of IssuesSearchRequestDataAttributesPersona.
const (
	ISSUESSEARCHREQUESTDATAATTRIBUTESPERSONA_ALL     IssuesSearchRequestDataAttributesPersona = "ALL"
	ISSUESSEARCHREQUESTDATAATTRIBUTESPERSONA_BROWSER IssuesSearchRequestDataAttributesPersona = "BROWSER"
	ISSUESSEARCHREQUESTDATAATTRIBUTESPERSONA_MOBILE  IssuesSearchRequestDataAttributesPersona = "MOBILE"
	ISSUESSEARCHREQUESTDATAATTRIBUTESPERSONA_BACKEND IssuesSearchRequestDataAttributesPersona = "BACKEND"
)

var allowedIssuesSearchRequestDataAttributesPersonaEnumValues = []IssuesSearchRequestDataAttributesPersona{
	ISSUESSEARCHREQUESTDATAATTRIBUTESPERSONA_ALL,
	ISSUESSEARCHREQUESTDATAATTRIBUTESPERSONA_BROWSER,
	ISSUESSEARCHREQUESTDATAATTRIBUTESPERSONA_MOBILE,
	ISSUESSEARCHREQUESTDATAATTRIBUTESPERSONA_BACKEND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssuesSearchRequestDataAttributesPersona) GetAllowedValues() []IssuesSearchRequestDataAttributesPersona {
	return allowedIssuesSearchRequestDataAttributesPersonaEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssuesSearchRequestDataAttributesPersona) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssuesSearchRequestDataAttributesPersona(value)
	return nil
}

// NewIssuesSearchRequestDataAttributesPersonaFromValue returns a pointer to a valid IssuesSearchRequestDataAttributesPersona
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssuesSearchRequestDataAttributesPersonaFromValue(v string) (*IssuesSearchRequestDataAttributesPersona, error) {
	ev := IssuesSearchRequestDataAttributesPersona(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssuesSearchRequestDataAttributesPersona: valid values are %v", v, allowedIssuesSearchRequestDataAttributesPersonaEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssuesSearchRequestDataAttributesPersona) IsValid() bool {
	for _, existing := range allowedIssuesSearchRequestDataAttributesPersonaEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssuesSearchRequestDataAttributesPersona value.
func (v IssuesSearchRequestDataAttributesPersona) Ptr() *IssuesSearchRequestDataAttributesPersona {
	return &v
}
