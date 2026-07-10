// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TicketCreationTarget The ticketing system to create tickets in.
type TicketCreationTarget string

// List of TicketCreationTarget.
const (
	TICKETCREATIONTARGET_JIRA            TicketCreationTarget = "jira"
	TICKETCREATIONTARGET_CASE_MANAGEMENT TicketCreationTarget = "case_management"
)

var allowedTicketCreationTargetEnumValues = []TicketCreationTarget{
	TICKETCREATIONTARGET_JIRA,
	TICKETCREATIONTARGET_CASE_MANAGEMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TicketCreationTarget) GetAllowedValues() []TicketCreationTarget {
	return allowedTicketCreationTargetEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TicketCreationTarget) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TicketCreationTarget(value)
	return nil
}

// NewTicketCreationTargetFromValue returns a pointer to a valid TicketCreationTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTicketCreationTargetFromValue(v string) (*TicketCreationTarget, error) {
	ev := TicketCreationTarget(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TicketCreationTarget: valid values are %v", v, allowedTicketCreationTargetEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TicketCreationTarget) IsValid() bool {
	for _, existing := range allowedTicketCreationTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TicketCreationTarget value.
func (v TicketCreationTarget) Ptr() *TicketCreationTarget {
	return &v
}
