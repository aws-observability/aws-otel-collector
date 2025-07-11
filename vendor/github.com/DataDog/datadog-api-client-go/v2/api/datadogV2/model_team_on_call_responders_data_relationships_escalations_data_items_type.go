// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType Identifies the resource type for escalation policy steps linked to a team's on-call configuration.
type TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType string

// List of TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType.
const (
	TEAMONCALLRESPONDERSDATARELATIONSHIPSESCALATIONSDATAITEMSTYPE_ESCALATION_POLICY_STEPS TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType = "escalation_policy_steps"
)

var allowedTeamOnCallRespondersDataRelationshipsEscalationsDataItemsTypeEnumValues = []TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType{
	TEAMONCALLRESPONDERSDATARELATIONSHIPSESCALATIONSDATAITEMSTYPE_ESCALATION_POLICY_STEPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType) GetAllowedValues() []TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType {
	return allowedTeamOnCallRespondersDataRelationshipsEscalationsDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType(value)
	return nil
}

// NewTeamOnCallRespondersDataRelationshipsEscalationsDataItemsTypeFromValue returns a pointer to a valid TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamOnCallRespondersDataRelationshipsEscalationsDataItemsTypeFromValue(v string) (*TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType, error) {
	ev := TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType: valid values are %v", v, allowedTeamOnCallRespondersDataRelationshipsEscalationsDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType) IsValid() bool {
	for _, existing := range allowedTeamOnCallRespondersDataRelationshipsEscalationsDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType value.
func (v TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType) Ptr() *TeamOnCallRespondersDataRelationshipsEscalationsDataItemsType {
	return &v
}
