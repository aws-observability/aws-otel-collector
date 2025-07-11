// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyDataType Indicates that the resource is of type `policies`.
type EscalationPolicyDataType string

// List of EscalationPolicyDataType.
const (
	ESCALATIONPOLICYDATATYPE_POLICIES EscalationPolicyDataType = "policies"
)

var allowedEscalationPolicyDataTypeEnumValues = []EscalationPolicyDataType{
	ESCALATIONPOLICYDATATYPE_POLICIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyDataType) GetAllowedValues() []EscalationPolicyDataType {
	return allowedEscalationPolicyDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyDataType(value)
	return nil
}

// NewEscalationPolicyDataTypeFromValue returns a pointer to a valid EscalationPolicyDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyDataTypeFromValue(v string) (*EscalationPolicyDataType, error) {
	ev := EscalationPolicyDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyDataType: valid values are %v", v, allowedEscalationPolicyDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyDataType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyDataType value.
func (v EscalationPolicyDataType) Ptr() *EscalationPolicyDataType {
	return &v
}
