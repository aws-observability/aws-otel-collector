// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestDecisionStatusType The status of a change request decision.
type ChangeRequestDecisionStatusType string

// List of ChangeRequestDecisionStatusType.
const (
	CHANGEREQUESTDECISIONSTATUSTYPE_REQUESTED ChangeRequestDecisionStatusType = "REQUESTED"
	CHANGEREQUESTDECISIONSTATUSTYPE_APPROVED  ChangeRequestDecisionStatusType = "APPROVED"
	CHANGEREQUESTDECISIONSTATUSTYPE_DECLINED  ChangeRequestDecisionStatusType = "DECLINED"
)

var allowedChangeRequestDecisionStatusTypeEnumValues = []ChangeRequestDecisionStatusType{
	CHANGEREQUESTDECISIONSTATUSTYPE_REQUESTED,
	CHANGEREQUESTDECISIONSTATUSTYPE_APPROVED,
	CHANGEREQUESTDECISIONSTATUSTYPE_DECLINED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeRequestDecisionStatusType) GetAllowedValues() []ChangeRequestDecisionStatusType {
	return allowedChangeRequestDecisionStatusTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeRequestDecisionStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeRequestDecisionStatusType(value)
	return nil
}

// NewChangeRequestDecisionStatusTypeFromValue returns a pointer to a valid ChangeRequestDecisionStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeRequestDecisionStatusTypeFromValue(v string) (*ChangeRequestDecisionStatusType, error) {
	ev := ChangeRequestDecisionStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeRequestDecisionStatusType: valid values are %v", v, allowedChangeRequestDecisionStatusTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeRequestDecisionStatusType) IsValid() bool {
	for _, existing := range allowedChangeRequestDecisionStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeRequestDecisionStatusType value.
func (v ChangeRequestDecisionStatusType) Ptr() *ChangeRequestDecisionStatusType {
	return &v
}
