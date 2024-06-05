// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseType Case type
type CaseType string

// List of CaseType.
const (
	CASETYPE_STANDARD CaseType = "STANDARD"
)

var allowedCaseTypeEnumValues = []CaseType{
	CASETYPE_STANDARD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseType) GetAllowedValues() []CaseType {
	return allowedCaseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseType(value)
	return nil
}

// NewCaseTypeFromValue returns a pointer to a valid CaseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseTypeFromValue(v string) (*CaseType, error) {
	ev := CaseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseType: valid values are %v", v, allowedCaseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseType) IsValid() bool {
	for _, existing := range allowedCaseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseType value.
func (v CaseType) Ptr() *CaseType {
	return &v
}
