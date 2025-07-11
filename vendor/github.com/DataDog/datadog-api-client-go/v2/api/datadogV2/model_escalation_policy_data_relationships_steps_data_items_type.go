// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyDataRelationshipsStepsDataItemsType Indicates that the resource is of type `steps`.
type EscalationPolicyDataRelationshipsStepsDataItemsType string

// List of EscalationPolicyDataRelationshipsStepsDataItemsType.
const (
	ESCALATIONPOLICYDATARELATIONSHIPSSTEPSDATAITEMSTYPE_STEPS EscalationPolicyDataRelationshipsStepsDataItemsType = "steps"
)

var allowedEscalationPolicyDataRelationshipsStepsDataItemsTypeEnumValues = []EscalationPolicyDataRelationshipsStepsDataItemsType{
	ESCALATIONPOLICYDATARELATIONSHIPSSTEPSDATAITEMSTYPE_STEPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyDataRelationshipsStepsDataItemsType) GetAllowedValues() []EscalationPolicyDataRelationshipsStepsDataItemsType {
	return allowedEscalationPolicyDataRelationshipsStepsDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyDataRelationshipsStepsDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyDataRelationshipsStepsDataItemsType(value)
	return nil
}

// NewEscalationPolicyDataRelationshipsStepsDataItemsTypeFromValue returns a pointer to a valid EscalationPolicyDataRelationshipsStepsDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyDataRelationshipsStepsDataItemsTypeFromValue(v string) (*EscalationPolicyDataRelationshipsStepsDataItemsType, error) {
	ev := EscalationPolicyDataRelationshipsStepsDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyDataRelationshipsStepsDataItemsType: valid values are %v", v, allowedEscalationPolicyDataRelationshipsStepsDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyDataRelationshipsStepsDataItemsType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyDataRelationshipsStepsDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyDataRelationshipsStepsDataItemsType value.
func (v EscalationPolicyDataRelationshipsStepsDataItemsType) Ptr() *EscalationPolicyDataRelationshipsStepsDataItemsType {
	return &v
}
