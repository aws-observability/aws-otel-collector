// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PersonalAccessTokensSort Sorting options
type PersonalAccessTokensSort string

// List of PersonalAccessTokensSort.
const (
	PERSONALACCESSTOKENSSORT_NAME_ASCENDING        PersonalAccessTokensSort = "name"
	PERSONALACCESSTOKENSSORT_NAME_DESCENDING       PersonalAccessTokensSort = "-name"
	PERSONALACCESSTOKENSSORT_CREATED_AT_ASCENDING  PersonalAccessTokensSort = "created_at"
	PERSONALACCESSTOKENSSORT_CREATED_AT_DESCENDING PersonalAccessTokensSort = "-created_at"
	PERSONALACCESSTOKENSSORT_EXPIRES_AT_ASCENDING  PersonalAccessTokensSort = "expires_at"
	PERSONALACCESSTOKENSSORT_EXPIRES_AT_DESCENDING PersonalAccessTokensSort = "-expires_at"
)

var allowedPersonalAccessTokensSortEnumValues = []PersonalAccessTokensSort{
	PERSONALACCESSTOKENSSORT_NAME_ASCENDING,
	PERSONALACCESSTOKENSSORT_NAME_DESCENDING,
	PERSONALACCESSTOKENSSORT_CREATED_AT_ASCENDING,
	PERSONALACCESSTOKENSSORT_CREATED_AT_DESCENDING,
	PERSONALACCESSTOKENSSORT_EXPIRES_AT_ASCENDING,
	PERSONALACCESSTOKENSSORT_EXPIRES_AT_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PersonalAccessTokensSort) GetAllowedValues() []PersonalAccessTokensSort {
	return allowedPersonalAccessTokensSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PersonalAccessTokensSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PersonalAccessTokensSort(value)
	return nil
}

// NewPersonalAccessTokensSortFromValue returns a pointer to a valid PersonalAccessTokensSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPersonalAccessTokensSortFromValue(v string) (*PersonalAccessTokensSort, error) {
	ev := PersonalAccessTokensSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PersonalAccessTokensSort: valid values are %v", v, allowedPersonalAccessTokensSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PersonalAccessTokensSort) IsValid() bool {
	for _, existing := range allowedPersonalAccessTokensSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PersonalAccessTokensSort value.
func (v PersonalAccessTokensSort) Ptr() *PersonalAccessTokensSort {
	return &v
}
