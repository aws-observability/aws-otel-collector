// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsCommitmentType Type of commitment. ri for Reserved Instances, sp for Savings Plans.
type CommitmentsCommitmentType string

// List of CommitmentsCommitmentType.
const (
	COMMITMENTSCOMMITMENTTYPE_RESERVED_INSTANCES CommitmentsCommitmentType = "ri"
	COMMITMENTSCOMMITMENTTYPE_SAVINGS_PLANS      CommitmentsCommitmentType = "sp"
)

var allowedCommitmentsCommitmentTypeEnumValues = []CommitmentsCommitmentType{
	COMMITMENTSCOMMITMENTTYPE_RESERVED_INSTANCES,
	COMMITMENTSCOMMITMENTTYPE_SAVINGS_PLANS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CommitmentsCommitmentType) GetAllowedValues() []CommitmentsCommitmentType {
	return allowedCommitmentsCommitmentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CommitmentsCommitmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CommitmentsCommitmentType(value)
	return nil
}

// NewCommitmentsCommitmentTypeFromValue returns a pointer to a valid CommitmentsCommitmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCommitmentsCommitmentTypeFromValue(v string) (*CommitmentsCommitmentType, error) {
	ev := CommitmentsCommitmentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CommitmentsCommitmentType: valid values are %v", v, allowedCommitmentsCommitmentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CommitmentsCommitmentType) IsValid() bool {
	for _, existing := range allowedCommitmentsCommitmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommitmentsCommitmentType value.
func (v CommitmentsCommitmentType) Ptr() *CommitmentsCommitmentType {
	return &v
}
