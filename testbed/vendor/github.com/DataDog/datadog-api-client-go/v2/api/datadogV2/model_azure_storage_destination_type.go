// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureStorageDestinationType The destination type. The value should always be `azure_storage`.
type AzureStorageDestinationType string

// List of AzureStorageDestinationType.
const (
	AZURESTORAGEDESTINATIONTYPE_AZURE_STORAGE AzureStorageDestinationType = "azure_storage"
)

var allowedAzureStorageDestinationTypeEnumValues = []AzureStorageDestinationType{
	AZURESTORAGEDESTINATIONTYPE_AZURE_STORAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AzureStorageDestinationType) GetAllowedValues() []AzureStorageDestinationType {
	return allowedAzureStorageDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AzureStorageDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AzureStorageDestinationType(value)
	return nil
}

// NewAzureStorageDestinationTypeFromValue returns a pointer to a valid AzureStorageDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAzureStorageDestinationTypeFromValue(v string) (*AzureStorageDestinationType, error) {
	ev := AzureStorageDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AzureStorageDestinationType: valid values are %v", v, allowedAzureStorageDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AzureStorageDestinationType) IsValid() bool {
	for _, existing := range allowedAzureStorageDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureStorageDestinationType value.
func (v AzureStorageDestinationType) Ptr() *AzureStorageDestinationType {
	return &v
}
