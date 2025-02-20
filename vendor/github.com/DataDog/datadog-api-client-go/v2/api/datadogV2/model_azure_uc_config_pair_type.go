// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureUCConfigPairType Type of Azure config pair.
type AzureUCConfigPairType string

// List of AzureUCConfigPairType.
const (
	AZUREUCCONFIGPAIRTYPE_AZURE_UC_CONFIGS AzureUCConfigPairType = "azure_uc_configs"
)

var allowedAzureUCConfigPairTypeEnumValues = []AzureUCConfigPairType{
	AZUREUCCONFIGPAIRTYPE_AZURE_UC_CONFIGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AzureUCConfigPairType) GetAllowedValues() []AzureUCConfigPairType {
	return allowedAzureUCConfigPairTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AzureUCConfigPairType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AzureUCConfigPairType(value)
	return nil
}

// NewAzureUCConfigPairTypeFromValue returns a pointer to a valid AzureUCConfigPairType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAzureUCConfigPairTypeFromValue(v string) (*AzureUCConfigPairType, error) {
	ev := AzureUCConfigPairType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AzureUCConfigPairType: valid values are %v", v, allowedAzureUCConfigPairTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AzureUCConfigPairType) IsValid() bool {
	for _, existing := range allowedAzureUCConfigPairTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureUCConfigPairType value.
func (v AzureUCConfigPairType) Ptr() *AzureUCConfigPairType {
	return &v
}
