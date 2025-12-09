// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotifyEndState A notification end state.
type NotifyEndState string

// List of NotifyEndState.
const (
	NOTIFYENDSTATE_ALERT   NotifyEndState = "alert"
	NOTIFYENDSTATE_NO_DATA NotifyEndState = "no data"
	NOTIFYENDSTATE_WARN    NotifyEndState = "warn"
)

var allowedNotifyEndStateEnumValues = []NotifyEndState{
	NOTIFYENDSTATE_ALERT,
	NOTIFYENDSTATE_NO_DATA,
	NOTIFYENDSTATE_WARN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotifyEndState) GetAllowedValues() []NotifyEndState {
	return allowedNotifyEndStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotifyEndState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotifyEndState(value)
	return nil
}

// NewNotifyEndStateFromValue returns a pointer to a valid NotifyEndState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotifyEndStateFromValue(v string) (*NotifyEndState, error) {
	ev := NotifyEndState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotifyEndState: valid values are %v", v, allowedNotifyEndStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotifyEndState) IsValid() bool {
	for _, existing := range allowedNotifyEndStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotifyEndState value.
func (v NotifyEndState) Ptr() *NotifyEndState {
	return &v
}
