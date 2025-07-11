// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyCreateRequestDataType Indicates that the resource is of type `policies`.
type EscalationPolicyCreateRequestDataType string

// List of EscalationPolicyCreateRequestDataType.
const (
	ESCALATIONPOLICYCREATEREQUESTDATATYPE_POLICIES EscalationPolicyCreateRequestDataType = "policies"
)

var allowedEscalationPolicyCreateRequestDataTypeEnumValues = []EscalationPolicyCreateRequestDataType{
	ESCALATIONPOLICYCREATEREQUESTDATATYPE_POLICIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyCreateRequestDataType) GetAllowedValues() []EscalationPolicyCreateRequestDataType {
	return allowedEscalationPolicyCreateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyCreateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyCreateRequestDataType(value)
	return nil
}

// NewEscalationPolicyCreateRequestDataTypeFromValue returns a pointer to a valid EscalationPolicyCreateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyCreateRequestDataTypeFromValue(v string) (*EscalationPolicyCreateRequestDataType, error) {
	ev := EscalationPolicyCreateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyCreateRequestDataType: valid values are %v", v, allowedEscalationPolicyCreateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyCreateRequestDataType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyCreateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyCreateRequestDataType value.
func (v EscalationPolicyCreateRequestDataType) Ptr() *EscalationPolicyCreateRequestDataType {
	return &v
}
