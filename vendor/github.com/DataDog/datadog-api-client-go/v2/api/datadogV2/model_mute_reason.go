// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MuteReason The reason for muting a security finding.
type MuteReason string

// List of MuteReason.
const (
	MUTEREASON_DUPLICATE      MuteReason = "duplicate"
	MUTEREASON_FALSE_POSITIVE MuteReason = "false_positive"
	MUTEREASON_NO_FIX         MuteReason = "no_fix"
	MUTEREASON_OTHER          MuteReason = "other"
	MUTEREASON_PENDING_FIX    MuteReason = "pending_fix"
	MUTEREASON_RISK_ACCEPTED  MuteReason = "risk_accepted"
)

var allowedMuteReasonEnumValues = []MuteReason{
	MUTEREASON_DUPLICATE,
	MUTEREASON_FALSE_POSITIVE,
	MUTEREASON_NO_FIX,
	MUTEREASON_OTHER,
	MUTEREASON_PENDING_FIX,
	MUTEREASON_RISK_ACCEPTED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MuteReason) GetAllowedValues() []MuteReason {
	return allowedMuteReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MuteReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MuteReason(value)
	return nil
}

// NewMuteReasonFromValue returns a pointer to a valid MuteReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMuteReasonFromValue(v string) (*MuteReason, error) {
	ev := MuteReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MuteReason: valid values are %v", v, allowedMuteReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MuteReason) IsValid() bool {
	for _, existing := range allowedMuteReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MuteReason value.
func (v MuteReason) Ptr() *MuteReason {
	return &v
}
