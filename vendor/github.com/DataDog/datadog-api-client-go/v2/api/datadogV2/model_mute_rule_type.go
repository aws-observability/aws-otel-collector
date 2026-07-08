// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MuteRuleType The JSON:API type for mute rules.
type MuteRuleType string

// List of MuteRuleType.
const (
	MUTERULETYPE_MUTE_RULES MuteRuleType = "mute_rules"
)

var allowedMuteRuleTypeEnumValues = []MuteRuleType{
	MUTERULETYPE_MUTE_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MuteRuleType) GetAllowedValues() []MuteRuleType {
	return allowedMuteRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MuteRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MuteRuleType(value)
	return nil
}

// NewMuteRuleTypeFromValue returns a pointer to a valid MuteRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMuteRuleTypeFromValue(v string) (*MuteRuleType, error) {
	ev := MuteRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MuteRuleType: valid values are %v", v, allowedMuteRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MuteRuleType) IsValid() bool {
	for _, existing := range allowedMuteRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MuteRuleType value.
func (v MuteRuleType) Ptr() *MuteRuleType {
	return &v
}
