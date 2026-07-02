// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpsgenieAccountType Opsgenie account resource type.
type OpsgenieAccountType string

// List of OpsgenieAccountType.
const (
	OPSGENIEACCOUNTTYPE_OPSGENIE_ACCOUNT OpsgenieAccountType = "opsgenie-account"
)

var allowedOpsgenieAccountTypeEnumValues = []OpsgenieAccountType{
	OPSGENIEACCOUNTTYPE_OPSGENIE_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OpsgenieAccountType) GetAllowedValues() []OpsgenieAccountType {
	return allowedOpsgenieAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpsgenieAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpsgenieAccountType(value)
	return nil
}

// NewOpsgenieAccountTypeFromValue returns a pointer to a valid OpsgenieAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpsgenieAccountTypeFromValue(v string) (*OpsgenieAccountType, error) {
	ev := OpsgenieAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpsgenieAccountType: valid values are %v", v, allowedOpsgenieAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpsgenieAccountType) IsValid() bool {
	for _, existing := range allowedOpsgenieAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpsgenieAccountType value.
func (v OpsgenieAccountType) Ptr() *OpsgenieAccountType {
	return &v
}
