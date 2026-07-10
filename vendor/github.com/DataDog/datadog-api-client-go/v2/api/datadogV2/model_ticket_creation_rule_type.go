// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TicketCreationRuleType The JSON:API type for ticket creation rules.
type TicketCreationRuleType string

// List of TicketCreationRuleType.
const (
	TICKETCREATIONRULETYPE_TICKET_CREATION_RULES TicketCreationRuleType = "ticket_creation_rules"
)

var allowedTicketCreationRuleTypeEnumValues = []TicketCreationRuleType{
	TICKETCREATIONRULETYPE_TICKET_CREATION_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TicketCreationRuleType) GetAllowedValues() []TicketCreationRuleType {
	return allowedTicketCreationRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TicketCreationRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TicketCreationRuleType(value)
	return nil
}

// NewTicketCreationRuleTypeFromValue returns a pointer to a valid TicketCreationRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTicketCreationRuleTypeFromValue(v string) (*TicketCreationRuleType, error) {
	ev := TicketCreationRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TicketCreationRuleType: valid values are %v", v, allowedTicketCreationRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TicketCreationRuleType) IsValid() bool {
	for _, existing := range allowedTicketCreationRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TicketCreationRuleType value.
func (v TicketCreationRuleType) Ptr() *TicketCreationRuleType {
	return &v
}
