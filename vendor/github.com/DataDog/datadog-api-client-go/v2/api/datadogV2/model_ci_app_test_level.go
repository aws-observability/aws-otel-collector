// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppTestLevel Test run level.
type CIAppTestLevel string

// List of CIAppTestLevel.
const (
	CIAPPTESTLEVEL_SESSION CIAppTestLevel = "session"
	CIAPPTESTLEVEL_MODULE  CIAppTestLevel = "module"
	CIAPPTESTLEVEL_SUITE   CIAppTestLevel = "suite"
	CIAPPTESTLEVEL_TEST    CIAppTestLevel = "test"
)

var allowedCIAppTestLevelEnumValues = []CIAppTestLevel{
	CIAPPTESTLEVEL_SESSION,
	CIAPPTESTLEVEL_MODULE,
	CIAPPTESTLEVEL_SUITE,
	CIAPPTESTLEVEL_TEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppTestLevel) GetAllowedValues() []CIAppTestLevel {
	return allowedCIAppTestLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppTestLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppTestLevel(value)
	return nil
}

// NewCIAppTestLevelFromValue returns a pointer to a valid CIAppTestLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppTestLevelFromValue(v string) (*CIAppTestLevel, error) {
	ev := CIAppTestLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppTestLevel: valid values are %v", v, allowedCIAppTestLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppTestLevel) IsValid() bool {
	for _, existing := range allowedCIAppTestLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppTestLevel value.
func (v CIAppTestLevel) Ptr() *CIAppTestLevel {
	return &v
}
