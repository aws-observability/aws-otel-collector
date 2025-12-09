// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestAttributesFlakyState The current state of the flaky test.
type FlakyTestAttributesFlakyState string

// List of FlakyTestAttributesFlakyState.
const (
	FLAKYTESTATTRIBUTESFLAKYSTATE_ACTIVE      FlakyTestAttributesFlakyState = "active"
	FLAKYTESTATTRIBUTESFLAKYSTATE_FIXED       FlakyTestAttributesFlakyState = "fixed"
	FLAKYTESTATTRIBUTESFLAKYSTATE_QUARANTINED FlakyTestAttributesFlakyState = "quarantined"
	FLAKYTESTATTRIBUTESFLAKYSTATE_DISABLED    FlakyTestAttributesFlakyState = "disabled"
)

var allowedFlakyTestAttributesFlakyStateEnumValues = []FlakyTestAttributesFlakyState{
	FLAKYTESTATTRIBUTESFLAKYSTATE_ACTIVE,
	FLAKYTESTATTRIBUTESFLAKYSTATE_FIXED,
	FLAKYTESTATTRIBUTESFLAKYSTATE_QUARANTINED,
	FLAKYTESTATTRIBUTESFLAKYSTATE_DISABLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlakyTestAttributesFlakyState) GetAllowedValues() []FlakyTestAttributesFlakyState {
	return allowedFlakyTestAttributesFlakyStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlakyTestAttributesFlakyState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlakyTestAttributesFlakyState(value)
	return nil
}

// NewFlakyTestAttributesFlakyStateFromValue returns a pointer to a valid FlakyTestAttributesFlakyState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlakyTestAttributesFlakyStateFromValue(v string) (*FlakyTestAttributesFlakyState, error) {
	ev := FlakyTestAttributesFlakyState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlakyTestAttributesFlakyState: valid values are %v", v, allowedFlakyTestAttributesFlakyStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlakyTestAttributesFlakyState) IsValid() bool {
	for _, existing := range allowedFlakyTestAttributesFlakyStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlakyTestAttributesFlakyState value.
func (v FlakyTestAttributesFlakyState) Ptr() *FlakyTestAttributesFlakyState {
	return &v
}
