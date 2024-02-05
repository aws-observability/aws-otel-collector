// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerRuleType Sensitive Data Scanner rule type.
type SensitiveDataScannerRuleType string

// List of SensitiveDataScannerRuleType.
const (
	SENSITIVEDATASCANNERRULETYPE_SENSITIVE_DATA_SCANNER_RULE SensitiveDataScannerRuleType = "sensitive_data_scanner_rule"
)

var allowedSensitiveDataScannerRuleTypeEnumValues = []SensitiveDataScannerRuleType{
	SENSITIVEDATASCANNERRULETYPE_SENSITIVE_DATA_SCANNER_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SensitiveDataScannerRuleType) GetAllowedValues() []SensitiveDataScannerRuleType {
	return allowedSensitiveDataScannerRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SensitiveDataScannerRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SensitiveDataScannerRuleType(value)
	return nil
}

// NewSensitiveDataScannerRuleTypeFromValue returns a pointer to a valid SensitiveDataScannerRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSensitiveDataScannerRuleTypeFromValue(v string) (*SensitiveDataScannerRuleType, error) {
	ev := SensitiveDataScannerRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SensitiveDataScannerRuleType: valid values are %v", v, allowedSensitiveDataScannerRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SensitiveDataScannerRuleType) IsValid() bool {
	for _, existing := range allowedSensitiveDataScannerRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensitiveDataScannerRuleType value.
func (v SensitiveDataScannerRuleType) Ptr() *SensitiveDataScannerRuleType {
	return &v
}
