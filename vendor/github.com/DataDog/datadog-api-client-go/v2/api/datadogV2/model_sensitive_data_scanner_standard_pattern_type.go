// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerStandardPatternType Sensitive Data Scanner standard pattern type.
type SensitiveDataScannerStandardPatternType string

// List of SensitiveDataScannerStandardPatternType.
const (
	SENSITIVEDATASCANNERSTANDARDPATTERNTYPE_SENSITIVE_DATA_SCANNER_STANDARD_PATTERN SensitiveDataScannerStandardPatternType = "sensitive_data_scanner_standard_pattern"
)

var allowedSensitiveDataScannerStandardPatternTypeEnumValues = []SensitiveDataScannerStandardPatternType{
	SENSITIVEDATASCANNERSTANDARDPATTERNTYPE_SENSITIVE_DATA_SCANNER_STANDARD_PATTERN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SensitiveDataScannerStandardPatternType) GetAllowedValues() []SensitiveDataScannerStandardPatternType {
	return allowedSensitiveDataScannerStandardPatternTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SensitiveDataScannerStandardPatternType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SensitiveDataScannerStandardPatternType(value)
	return nil
}

// NewSensitiveDataScannerStandardPatternTypeFromValue returns a pointer to a valid SensitiveDataScannerStandardPatternType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSensitiveDataScannerStandardPatternTypeFromValue(v string) (*SensitiveDataScannerStandardPatternType, error) {
	ev := SensitiveDataScannerStandardPatternType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SensitiveDataScannerStandardPatternType: valid values are %v", v, allowedSensitiveDataScannerStandardPatternTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SensitiveDataScannerStandardPatternType) IsValid() bool {
	for _, existing := range allowedSensitiveDataScannerStandardPatternTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensitiveDataScannerStandardPatternType value.
func (v SensitiveDataScannerStandardPatternType) Ptr() *SensitiveDataScannerStandardPatternType {
	return &v
}
