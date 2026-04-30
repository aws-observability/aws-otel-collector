// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PaginationMetaPageType The pagination type used for offset-based pagination.
type PaginationMetaPageType string

// List of PaginationMetaPageType.
const (
	PAGINATIONMETAPAGETYPE_OFFSET_LIMIT PaginationMetaPageType = "offset_limit"
)

var allowedPaginationMetaPageTypeEnumValues = []PaginationMetaPageType{
	PAGINATIONMETAPAGETYPE_OFFSET_LIMIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PaginationMetaPageType) GetAllowedValues() []PaginationMetaPageType {
	return allowedPaginationMetaPageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PaginationMetaPageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PaginationMetaPageType(value)
	return nil
}

// NewPaginationMetaPageTypeFromValue returns a pointer to a valid PaginationMetaPageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPaginationMetaPageTypeFromValue(v string) (*PaginationMetaPageType, error) {
	ev := PaginationMetaPageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PaginationMetaPageType: valid values are %v", v, allowedPaginationMetaPageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PaginationMetaPageType) IsValid() bool {
	for _, existing := range allowedPaginationMetaPageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaginationMetaPageType value.
func (v PaginationMetaPageType) Ptr() *PaginationMetaPageType {
	return &v
}
