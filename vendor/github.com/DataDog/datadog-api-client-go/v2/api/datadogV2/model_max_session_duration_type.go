// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaxSessionDurationType Data type of a maximum session duration update.
type MaxSessionDurationType string

// List of MaxSessionDurationType.
const (
	MAXSESSIONDURATIONTYPE_MAX_SESSION_DURATION MaxSessionDurationType = "max_session_duration"
)

var allowedMaxSessionDurationTypeEnumValues = []MaxSessionDurationType{
	MAXSESSIONDURATIONTYPE_MAX_SESSION_DURATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MaxSessionDurationType) GetAllowedValues() []MaxSessionDurationType {
	return allowedMaxSessionDurationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MaxSessionDurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MaxSessionDurationType(value)
	return nil
}

// NewMaxSessionDurationTypeFromValue returns a pointer to a valid MaxSessionDurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMaxSessionDurationTypeFromValue(v string) (*MaxSessionDurationType, error) {
	ev := MaxSessionDurationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MaxSessionDurationType: valid values are %v", v, allowedMaxSessionDurationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MaxSessionDurationType) IsValid() bool {
	for _, existing := range allowedMaxSessionDurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MaxSessionDurationType value.
func (v MaxSessionDurationType) Ptr() *MaxSessionDurationType {
	return &v
}
