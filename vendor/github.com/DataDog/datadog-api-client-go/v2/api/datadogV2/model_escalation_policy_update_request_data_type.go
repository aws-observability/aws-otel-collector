// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyUpdateRequestDataType Indicates that the resource is of type `policies`.
type EscalationPolicyUpdateRequestDataType string

// List of EscalationPolicyUpdateRequestDataType.
const (
	ESCALATIONPOLICYUPDATEREQUESTDATATYPE_POLICIES EscalationPolicyUpdateRequestDataType = "policies"
)

var allowedEscalationPolicyUpdateRequestDataTypeEnumValues = []EscalationPolicyUpdateRequestDataType{
	ESCALATIONPOLICYUPDATEREQUESTDATATYPE_POLICIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyUpdateRequestDataType) GetAllowedValues() []EscalationPolicyUpdateRequestDataType {
	return allowedEscalationPolicyUpdateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyUpdateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyUpdateRequestDataType(value)
	return nil
}

// NewEscalationPolicyUpdateRequestDataTypeFromValue returns a pointer to a valid EscalationPolicyUpdateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyUpdateRequestDataTypeFromValue(v string) (*EscalationPolicyUpdateRequestDataType, error) {
	ev := EscalationPolicyUpdateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyUpdateRequestDataType: valid values are %v", v, allowedEscalationPolicyUpdateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyUpdateRequestDataType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyUpdateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyUpdateRequestDataType value.
func (v EscalationPolicyUpdateRequestDataType) Ptr() *EscalationPolicyUpdateRequestDataType {
	return &v
}
