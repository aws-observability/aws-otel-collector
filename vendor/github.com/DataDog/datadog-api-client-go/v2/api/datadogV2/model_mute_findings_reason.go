// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MuteFindingsReason The reason why the findings are muted or unmuted.
type MuteFindingsReason string

// List of MuteFindingsReason.
const (
	MUTEFINDINGSREASON_PENDING_FIX             MuteFindingsReason = "PENDING_FIX"
	MUTEFINDINGSREASON_FALSE_POSITIVE          MuteFindingsReason = "FALSE_POSITIVE"
	MUTEFINDINGSREASON_OTHER                   MuteFindingsReason = "OTHER"
	MUTEFINDINGSREASON_NO_FIX                  MuteFindingsReason = "NO_FIX"
	MUTEFINDINGSREASON_DUPLICATE               MuteFindingsReason = "DUPLICATE"
	MUTEFINDINGSREASON_RISK_ACCEPTED           MuteFindingsReason = "RISK_ACCEPTED"
	MUTEFINDINGSREASON_NO_PENDING_FIX          MuteFindingsReason = "NO_PENDING_FIX"
	MUTEFINDINGSREASON_HUMAN_ERROR             MuteFindingsReason = "HUMAN_ERROR"
	MUTEFINDINGSREASON_NO_LONGER_ACCEPTED_RISK MuteFindingsReason = "NO_LONGER_ACCEPTED_RISK"
)

var allowedMuteFindingsReasonEnumValues = []MuteFindingsReason{
	MUTEFINDINGSREASON_PENDING_FIX,
	MUTEFINDINGSREASON_FALSE_POSITIVE,
	MUTEFINDINGSREASON_OTHER,
	MUTEFINDINGSREASON_NO_FIX,
	MUTEFINDINGSREASON_DUPLICATE,
	MUTEFINDINGSREASON_RISK_ACCEPTED,
	MUTEFINDINGSREASON_NO_PENDING_FIX,
	MUTEFINDINGSREASON_HUMAN_ERROR,
	MUTEFINDINGSREASON_NO_LONGER_ACCEPTED_RISK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MuteFindingsReason) GetAllowedValues() []MuteFindingsReason {
	return allowedMuteFindingsReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MuteFindingsReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MuteFindingsReason(value)
	return nil
}

// NewMuteFindingsReasonFromValue returns a pointer to a valid MuteFindingsReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMuteFindingsReasonFromValue(v string) (*MuteFindingsReason, error) {
	ev := MuteFindingsReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MuteFindingsReason: valid values are %v", v, allowedMuteFindingsReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MuteFindingsReason) IsValid() bool {
	for _, existing := range allowedMuteFindingsReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MuteFindingsReason value.
func (v MuteFindingsReason) Ptr() *MuteFindingsReason {
	return &v
}
