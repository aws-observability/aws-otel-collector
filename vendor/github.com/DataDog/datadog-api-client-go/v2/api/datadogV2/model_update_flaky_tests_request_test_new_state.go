// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateFlakyTestsRequestTestNewState The new state to set for the flaky test.
type UpdateFlakyTestsRequestTestNewState string

// List of UpdateFlakyTestsRequestTestNewState.
const (
	UPDATEFLAKYTESTSREQUESTTESTNEWSTATE_ACTIVE      UpdateFlakyTestsRequestTestNewState = "active"
	UPDATEFLAKYTESTSREQUESTTESTNEWSTATE_QUARANTINED UpdateFlakyTestsRequestTestNewState = "quarantined"
	UPDATEFLAKYTESTSREQUESTTESTNEWSTATE_DISABLED    UpdateFlakyTestsRequestTestNewState = "disabled"
	UPDATEFLAKYTESTSREQUESTTESTNEWSTATE_FIXED       UpdateFlakyTestsRequestTestNewState = "fixed"
)

var allowedUpdateFlakyTestsRequestTestNewStateEnumValues = []UpdateFlakyTestsRequestTestNewState{
	UPDATEFLAKYTESTSREQUESTTESTNEWSTATE_ACTIVE,
	UPDATEFLAKYTESTSREQUESTTESTNEWSTATE_QUARANTINED,
	UPDATEFLAKYTESTSREQUESTTESTNEWSTATE_DISABLED,
	UPDATEFLAKYTESTSREQUESTTESTNEWSTATE_FIXED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateFlakyTestsRequestTestNewState) GetAllowedValues() []UpdateFlakyTestsRequestTestNewState {
	return allowedUpdateFlakyTestsRequestTestNewStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateFlakyTestsRequestTestNewState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateFlakyTestsRequestTestNewState(value)
	return nil
}

// NewUpdateFlakyTestsRequestTestNewStateFromValue returns a pointer to a valid UpdateFlakyTestsRequestTestNewState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateFlakyTestsRequestTestNewStateFromValue(v string) (*UpdateFlakyTestsRequestTestNewState, error) {
	ev := UpdateFlakyTestsRequestTestNewState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateFlakyTestsRequestTestNewState: valid values are %v", v, allowedUpdateFlakyTestsRequestTestNewStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateFlakyTestsRequestTestNewState) IsValid() bool {
	for _, existing := range allowedUpdateFlakyTestsRequestTestNewStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateFlakyTestsRequestTestNewState value.
func (v UpdateFlakyTestsRequestTestNewState) Ptr() *UpdateFlakyTestsRequestTestNewState {
	return &v
}
