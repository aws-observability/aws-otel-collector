// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReferenceTableSortType Sort field and direction for reference tables. Use field name for ascending, prefix with "-" for descending.
type ReferenceTableSortType string

// List of ReferenceTableSortType.
const (
	REFERENCETABLESORTTYPE_UPDATED_AT       ReferenceTableSortType = "updated_at"
	REFERENCETABLESORTTYPE_TABLE_NAME       ReferenceTableSortType = "table_name"
	REFERENCETABLESORTTYPE_STATUS           ReferenceTableSortType = "status"
	REFERENCETABLESORTTYPE_MINUS_UPDATED_AT ReferenceTableSortType = "-updated_at"
	REFERENCETABLESORTTYPE_MINUS_TABLE_NAME ReferenceTableSortType = "-table_name"
	REFERENCETABLESORTTYPE_MINUS_STATUS     ReferenceTableSortType = "-status"
)

var allowedReferenceTableSortTypeEnumValues = []ReferenceTableSortType{
	REFERENCETABLESORTTYPE_UPDATED_AT,
	REFERENCETABLESORTTYPE_TABLE_NAME,
	REFERENCETABLESORTTYPE_STATUS,
	REFERENCETABLESORTTYPE_MINUS_UPDATED_AT,
	REFERENCETABLESORTTYPE_MINUS_TABLE_NAME,
	REFERENCETABLESORTTYPE_MINUS_STATUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReferenceTableSortType) GetAllowedValues() []ReferenceTableSortType {
	return allowedReferenceTableSortTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReferenceTableSortType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReferenceTableSortType(value)
	return nil
}

// NewReferenceTableSortTypeFromValue returns a pointer to a valid ReferenceTableSortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReferenceTableSortTypeFromValue(v string) (*ReferenceTableSortType, error) {
	ev := ReferenceTableSortType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReferenceTableSortType: valid values are %v", v, allowedReferenceTableSortTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReferenceTableSortType) IsValid() bool {
	for _, existing := range allowedReferenceTableSortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReferenceTableSortType value.
func (v ReferenceTableSortType) Ptr() *ReferenceTableSortType {
	return &v
}
