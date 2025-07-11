// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationRelationshipsRespondersDataItemsType Represents the resource type for users assigned as responders in an escalation step.
type EscalationRelationshipsRespondersDataItemsType string

// List of EscalationRelationshipsRespondersDataItemsType.
const (
	ESCALATIONRELATIONSHIPSRESPONDERSDATAITEMSTYPE_USERS EscalationRelationshipsRespondersDataItemsType = "users"
)

var allowedEscalationRelationshipsRespondersDataItemsTypeEnumValues = []EscalationRelationshipsRespondersDataItemsType{
	ESCALATIONRELATIONSHIPSRESPONDERSDATAITEMSTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationRelationshipsRespondersDataItemsType) GetAllowedValues() []EscalationRelationshipsRespondersDataItemsType {
	return allowedEscalationRelationshipsRespondersDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationRelationshipsRespondersDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationRelationshipsRespondersDataItemsType(value)
	return nil
}

// NewEscalationRelationshipsRespondersDataItemsTypeFromValue returns a pointer to a valid EscalationRelationshipsRespondersDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationRelationshipsRespondersDataItemsTypeFromValue(v string) (*EscalationRelationshipsRespondersDataItemsType, error) {
	ev := EscalationRelationshipsRespondersDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationRelationshipsRespondersDataItemsType: valid values are %v", v, allowedEscalationRelationshipsRespondersDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationRelationshipsRespondersDataItemsType) IsValid() bool {
	for _, existing := range allowedEscalationRelationshipsRespondersDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationRelationshipsRespondersDataItemsType value.
func (v EscalationRelationshipsRespondersDataItemsType) Ptr() *EscalationRelationshipsRespondersDataItemsType {
	return &v
}
