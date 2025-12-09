// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssuesSearchResultType Type of the object.
type IssuesSearchResultType string

// List of IssuesSearchResultType.
const (
	ISSUESSEARCHRESULTTYPE_ERROR_TRACKING_SEARCH_RESULT IssuesSearchResultType = "error_tracking_search_result"
)

var allowedIssuesSearchResultTypeEnumValues = []IssuesSearchResultType{
	ISSUESSEARCHRESULTTYPE_ERROR_TRACKING_SEARCH_RESULT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssuesSearchResultType) GetAllowedValues() []IssuesSearchResultType {
	return allowedIssuesSearchResultTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssuesSearchResultType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssuesSearchResultType(value)
	return nil
}

// NewIssuesSearchResultTypeFromValue returns a pointer to a valid IssuesSearchResultType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssuesSearchResultTypeFromValue(v string) (*IssuesSearchResultType, error) {
	ev := IssuesSearchResultType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssuesSearchResultType: valid values are %v", v, allowedIssuesSearchResultTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssuesSearchResultType) IsValid() bool {
	for _, existing := range allowedIssuesSearchResultTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssuesSearchResultType value.
func (v IssuesSearchResultType) Ptr() *IssuesSearchResultType {
	return &v
}
