// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestRiskLevel The risk level of the change request.
type ChangeRequestRiskLevel string

// List of ChangeRequestRiskLevel.
const (
	CHANGEREQUESTRISKLEVEL_UNDEFINED ChangeRequestRiskLevel = "UNDEFINED"
	CHANGEREQUESTRISKLEVEL_LOW       ChangeRequestRiskLevel = "LOW"
	CHANGEREQUESTRISKLEVEL_MEDIUM    ChangeRequestRiskLevel = "MEDIUM"
	CHANGEREQUESTRISKLEVEL_HIGH      ChangeRequestRiskLevel = "HIGH"
)

var allowedChangeRequestRiskLevelEnumValues = []ChangeRequestRiskLevel{
	CHANGEREQUESTRISKLEVEL_UNDEFINED,
	CHANGEREQUESTRISKLEVEL_LOW,
	CHANGEREQUESTRISKLEVEL_MEDIUM,
	CHANGEREQUESTRISKLEVEL_HIGH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeRequestRiskLevel) GetAllowedValues() []ChangeRequestRiskLevel {
	return allowedChangeRequestRiskLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeRequestRiskLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeRequestRiskLevel(value)
	return nil
}

// NewChangeRequestRiskLevelFromValue returns a pointer to a valid ChangeRequestRiskLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeRequestRiskLevelFromValue(v string) (*ChangeRequestRiskLevel, error) {
	ev := ChangeRequestRiskLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeRequestRiskLevel: valid values are %v", v, allowedChangeRequestRiskLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeRequestRiskLevel) IsValid() bool {
	for _, existing := range allowedChangeRequestRiskLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeRequestRiskLevel value.
func (v ChangeRequestRiskLevel) Ptr() *ChangeRequestRiskLevel {
	return &v
}
