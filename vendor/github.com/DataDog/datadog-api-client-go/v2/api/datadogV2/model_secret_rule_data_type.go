// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecretRuleDataType Secret rule resource type.
type SecretRuleDataType string

// List of SecretRuleDataType.
const (
	SECRETRULEDATATYPE_SECRET_RULE SecretRuleDataType = "secret_rule"
)

var allowedSecretRuleDataTypeEnumValues = []SecretRuleDataType{
	SECRETRULEDATATYPE_SECRET_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecretRuleDataType) GetAllowedValues() []SecretRuleDataType {
	return allowedSecretRuleDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecretRuleDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecretRuleDataType(value)
	return nil
}

// NewSecretRuleDataTypeFromValue returns a pointer to a valid SecretRuleDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecretRuleDataTypeFromValue(v string) (*SecretRuleDataType, error) {
	ev := SecretRuleDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecretRuleDataType: valid values are %v", v, allowedSecretRuleDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecretRuleDataType) IsValid() bool {
	for _, existing := range allowedSecretRuleDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecretRuleDataType value.
func (v SecretRuleDataType) Ptr() *SecretRuleDataType {
	return &v
}
