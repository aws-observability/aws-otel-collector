// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureScanOptionsDataType The type of the resource. The value should always be `azure_scan_options`.
type AzureScanOptionsDataType string

// List of AzureScanOptionsDataType.
const (
	AZURESCANOPTIONSDATATYPE_AZURE_SCAN_OPTIONS AzureScanOptionsDataType = "azure_scan_options"
)

var allowedAzureScanOptionsDataTypeEnumValues = []AzureScanOptionsDataType{
	AZURESCANOPTIONSDATATYPE_AZURE_SCAN_OPTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AzureScanOptionsDataType) GetAllowedValues() []AzureScanOptionsDataType {
	return allowedAzureScanOptionsDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AzureScanOptionsDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AzureScanOptionsDataType(value)
	return nil
}

// NewAzureScanOptionsDataTypeFromValue returns a pointer to a valid AzureScanOptionsDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAzureScanOptionsDataTypeFromValue(v string) (*AzureScanOptionsDataType, error) {
	ev := AzureScanOptionsDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AzureScanOptionsDataType: valid values are %v", v, allowedAzureScanOptionsDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AzureScanOptionsDataType) IsValid() bool {
	for _, existing := range allowedAzureScanOptionsDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureScanOptionsDataType value.
func (v AzureScanOptionsDataType) Ptr() *AzureScanOptionsDataType {
	return &v
}
