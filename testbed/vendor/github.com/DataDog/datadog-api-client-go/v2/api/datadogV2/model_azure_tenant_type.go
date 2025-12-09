// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureTenantType The definition of the `AzureTenant` object.
type AzureTenantType string

// List of AzureTenantType.
const (
	AZURETENANTTYPE_AZURETENANT AzureTenantType = "AzureTenant"
)

var allowedAzureTenantTypeEnumValues = []AzureTenantType{
	AZURETENANTTYPE_AZURETENANT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AzureTenantType) GetAllowedValues() []AzureTenantType {
	return allowedAzureTenantTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AzureTenantType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AzureTenantType(value)
	return nil
}

// NewAzureTenantTypeFromValue returns a pointer to a valid AzureTenantType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAzureTenantTypeFromValue(v string) (*AzureTenantType, error) {
	ev := AzureTenantType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AzureTenantType: valid values are %v", v, allowedAzureTenantTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AzureTenantType) IsValid() bool {
	for _, existing := range allowedAzureTenantTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureTenantType value.
func (v AzureTenantType) Ptr() *AzureTenantType {
	return &v
}
