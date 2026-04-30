// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SlosGroupMode How SLO results are grouped in the response.
type SlosGroupMode string

// List of SlosGroupMode.
const (
	SLOSGROUPMODE_OVERALL    SlosGroupMode = "overall"
	SLOSGROUPMODE_COMPONENTS SlosGroupMode = "components"
)

var allowedSlosGroupModeEnumValues = []SlosGroupMode{
	SLOSGROUPMODE_OVERALL,
	SLOSGROUPMODE_COMPONENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SlosGroupMode) GetAllowedValues() []SlosGroupMode {
	return allowedSlosGroupModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SlosGroupMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SlosGroupMode(value)
	return nil
}

// NewSlosGroupModeFromValue returns a pointer to a valid SlosGroupMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSlosGroupModeFromValue(v string) (*SlosGroupMode, error) {
	ev := SlosGroupMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SlosGroupMode: valid values are %v", v, allowedSlosGroupModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SlosGroupMode) IsValid() bool {
	for _, existing := range allowedSlosGroupModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SlosGroupMode value.
func (v SlosGroupMode) Ptr() *SlosGroupMode {
	return &v
}
