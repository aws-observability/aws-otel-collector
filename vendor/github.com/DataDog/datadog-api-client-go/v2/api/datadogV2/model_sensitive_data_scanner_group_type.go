// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerGroupType Sensitive Data Scanner group type.
type SensitiveDataScannerGroupType string

// List of SensitiveDataScannerGroupType.
const (
	SENSITIVEDATASCANNERGROUPTYPE_SENSITIVE_DATA_SCANNER_GROUP SensitiveDataScannerGroupType = "sensitive_data_scanner_group"
)

var allowedSensitiveDataScannerGroupTypeEnumValues = []SensitiveDataScannerGroupType{
	SENSITIVEDATASCANNERGROUPTYPE_SENSITIVE_DATA_SCANNER_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SensitiveDataScannerGroupType) GetAllowedValues() []SensitiveDataScannerGroupType {
	return allowedSensitiveDataScannerGroupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SensitiveDataScannerGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SensitiveDataScannerGroupType(value)
	return nil
}

// NewSensitiveDataScannerGroupTypeFromValue returns a pointer to a valid SensitiveDataScannerGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSensitiveDataScannerGroupTypeFromValue(v string) (*SensitiveDataScannerGroupType, error) {
	ev := SensitiveDataScannerGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SensitiveDataScannerGroupType: valid values are %v", v, allowedSensitiveDataScannerGroupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SensitiveDataScannerGroupType) IsValid() bool {
	for _, existing := range allowedSensitiveDataScannerGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensitiveDataScannerGroupType value.
func (v SensitiveDataScannerGroupType) Ptr() *SensitiveDataScannerGroupType {
	return &v
}
