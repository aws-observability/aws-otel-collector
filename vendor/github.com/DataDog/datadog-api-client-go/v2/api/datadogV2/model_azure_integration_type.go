// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureIntegrationType The definition of the `AzureIntegrationType` object.
type AzureIntegrationType string

// List of AzureIntegrationType.
const (
	AZUREINTEGRATIONTYPE_AZURE AzureIntegrationType = "Azure"
)

var allowedAzureIntegrationTypeEnumValues = []AzureIntegrationType{
	AZUREINTEGRATIONTYPE_AZURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AzureIntegrationType) GetAllowedValues() []AzureIntegrationType {
	return allowedAzureIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AzureIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AzureIntegrationType(value)
	return nil
}

// NewAzureIntegrationTypeFromValue returns a pointer to a valid AzureIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAzureIntegrationTypeFromValue(v string) (*AzureIntegrationType, error) {
	ev := AzureIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AzureIntegrationType: valid values are %v", v, allowedAzureIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AzureIntegrationType) IsValid() bool {
	for _, existing := range allowedAzureIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureIntegrationType value.
func (v AzureIntegrationType) Ptr() *AzureIntegrationType {
	return &v
}
