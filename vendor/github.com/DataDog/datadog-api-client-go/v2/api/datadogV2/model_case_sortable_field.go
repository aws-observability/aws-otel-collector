// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseSortableField Case field that can be sorted on
type CaseSortableField string

// List of CaseSortableField.
const (
	CASESORTABLEFIELD_CREATED_AT CaseSortableField = "created_at"
	CASESORTABLEFIELD_PRIORITY   CaseSortableField = "priority"
	CASESORTABLEFIELD_STATUS     CaseSortableField = "status"
)

var allowedCaseSortableFieldEnumValues = []CaseSortableField{
	CASESORTABLEFIELD_CREATED_AT,
	CASESORTABLEFIELD_PRIORITY,
	CASESORTABLEFIELD_STATUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseSortableField) GetAllowedValues() []CaseSortableField {
	return allowedCaseSortableFieldEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseSortableField) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseSortableField(value)
	return nil
}

// NewCaseSortableFieldFromValue returns a pointer to a valid CaseSortableField
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseSortableFieldFromValue(v string) (*CaseSortableField, error) {
	ev := CaseSortableField(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseSortableField: valid values are %v", v, allowedCaseSortableFieldEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseSortableField) IsValid() bool {
	for _, existing := range allowedCaseSortableFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseSortableField value.
func (v CaseSortableField) Ptr() *CaseSortableField {
	return &v
}
