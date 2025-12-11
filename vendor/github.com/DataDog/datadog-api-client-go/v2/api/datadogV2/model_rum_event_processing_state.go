// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMEventProcessingState Configures which RUM events are processed and stored for the application.
type RUMEventProcessingState string

// List of RUMEventProcessingState.
const (
	RUMEVENTPROCESSINGSTATE_ALL                RUMEventProcessingState = "ALL"
	RUMEVENTPROCESSINGSTATE_ERROR_FOCUSED_MODE RUMEventProcessingState = "ERROR_FOCUSED_MODE"
	RUMEVENTPROCESSINGSTATE_NONE               RUMEventProcessingState = "NONE"
)

var allowedRUMEventProcessingStateEnumValues = []RUMEventProcessingState{
	RUMEVENTPROCESSINGSTATE_ALL,
	RUMEVENTPROCESSINGSTATE_ERROR_FOCUSED_MODE,
	RUMEVENTPROCESSINGSTATE_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMEventProcessingState) GetAllowedValues() []RUMEventProcessingState {
	return allowedRUMEventProcessingStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMEventProcessingState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMEventProcessingState(value)
	return nil
}

// NewRUMEventProcessingStateFromValue returns a pointer to a valid RUMEventProcessingState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMEventProcessingStateFromValue(v string) (*RUMEventProcessingState, error) {
	ev := RUMEventProcessingState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMEventProcessingState: valid values are %v", v, allowedRUMEventProcessingStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMEventProcessingState) IsValid() bool {
	for _, existing := range allowedRUMEventProcessingStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMEventProcessingState value.
func (v RUMEventProcessingState) Ptr() *RUMEventProcessingState {
	return &v
}
