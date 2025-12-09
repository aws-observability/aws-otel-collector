// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureScanOptionsInputUpdateDataType Azure scan options resource type.
type AzureScanOptionsInputUpdateDataType string

// List of AzureScanOptionsInputUpdateDataType.
const (
	AZURESCANOPTIONSINPUTUPDATEDATATYPE_AZURE_SCAN_OPTIONS AzureScanOptionsInputUpdateDataType = "azure_scan_options"
)

var allowedAzureScanOptionsInputUpdateDataTypeEnumValues = []AzureScanOptionsInputUpdateDataType{
	AZURESCANOPTIONSINPUTUPDATEDATATYPE_AZURE_SCAN_OPTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AzureScanOptionsInputUpdateDataType) GetAllowedValues() []AzureScanOptionsInputUpdateDataType {
	return allowedAzureScanOptionsInputUpdateDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AzureScanOptionsInputUpdateDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AzureScanOptionsInputUpdateDataType(value)
	return nil
}

// NewAzureScanOptionsInputUpdateDataTypeFromValue returns a pointer to a valid AzureScanOptionsInputUpdateDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAzureScanOptionsInputUpdateDataTypeFromValue(v string) (*AzureScanOptionsInputUpdateDataType, error) {
	ev := AzureScanOptionsInputUpdateDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AzureScanOptionsInputUpdateDataType: valid values are %v", v, allowedAzureScanOptionsInputUpdateDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AzureScanOptionsInputUpdateDataType) IsValid() bool {
	for _, existing := range allowedAzureScanOptionsInputUpdateDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureScanOptionsInputUpdateDataType value.
func (v AzureScanOptionsInputUpdateDataType) Ptr() *AzureScanOptionsInputUpdateDataType {
	return &v
}
