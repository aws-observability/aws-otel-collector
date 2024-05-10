// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CasePriority Case priority
type CasePriority string

// List of CasePriority.
const (
	CASEPRIORITY_NOT_DEFINED CasePriority = "NOT_DEFINED"
	CASEPRIORITY_P1          CasePriority = "P1"
	CASEPRIORITY_P2          CasePriority = "P2"
	CASEPRIORITY_P3          CasePriority = "P3"
	CASEPRIORITY_P4          CasePriority = "P4"
	CASEPRIORITY_P5          CasePriority = "P5"
)

var allowedCasePriorityEnumValues = []CasePriority{
	CASEPRIORITY_NOT_DEFINED,
	CASEPRIORITY_P1,
	CASEPRIORITY_P2,
	CASEPRIORITY_P3,
	CASEPRIORITY_P4,
	CASEPRIORITY_P5,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CasePriority) GetAllowedValues() []CasePriority {
	return allowedCasePriorityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CasePriority) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CasePriority(value)
	return nil
}

// NewCasePriorityFromValue returns a pointer to a valid CasePriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCasePriorityFromValue(v string) (*CasePriority, error) {
	ev := CasePriority(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CasePriority: valid values are %v", v, allowedCasePriorityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CasePriority) IsValid() bool {
	for _, existing := range allowedCasePriorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CasePriority value.
func (v CasePriority) Ptr() *CasePriority {
	return &v
}
