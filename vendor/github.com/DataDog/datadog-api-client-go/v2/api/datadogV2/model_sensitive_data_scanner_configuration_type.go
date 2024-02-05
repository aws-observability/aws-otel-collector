// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerConfigurationType Sensitive Data Scanner configuration type.
type SensitiveDataScannerConfigurationType string

// List of SensitiveDataScannerConfigurationType.
const (
	SENSITIVEDATASCANNERCONFIGURATIONTYPE_SENSITIVE_DATA_SCANNER_CONFIGURATIONS SensitiveDataScannerConfigurationType = "sensitive_data_scanner_configuration"
)

var allowedSensitiveDataScannerConfigurationTypeEnumValues = []SensitiveDataScannerConfigurationType{
	SENSITIVEDATASCANNERCONFIGURATIONTYPE_SENSITIVE_DATA_SCANNER_CONFIGURATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SensitiveDataScannerConfigurationType) GetAllowedValues() []SensitiveDataScannerConfigurationType {
	return allowedSensitiveDataScannerConfigurationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SensitiveDataScannerConfigurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SensitiveDataScannerConfigurationType(value)
	return nil
}

// NewSensitiveDataScannerConfigurationTypeFromValue returns a pointer to a valid SensitiveDataScannerConfigurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSensitiveDataScannerConfigurationTypeFromValue(v string) (*SensitiveDataScannerConfigurationType, error) {
	ev := SensitiveDataScannerConfigurationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SensitiveDataScannerConfigurationType: valid values are %v", v, allowedSensitiveDataScannerConfigurationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SensitiveDataScannerConfigurationType) IsValid() bool {
	for _, existing := range allowedSensitiveDataScannerConfigurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensitiveDataScannerConfigurationType value.
func (v SensitiveDataScannerConfigurationType) Ptr() *SensitiveDataScannerConfigurationType {
	return &v
}
