// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AsanaIntegrationType The definition of the `AsanaIntegrationType` object.
type AsanaIntegrationType string

// List of AsanaIntegrationType.
const (
	ASANAINTEGRATIONTYPE_ASANA AsanaIntegrationType = "Asana"
)

var allowedAsanaIntegrationTypeEnumValues = []AsanaIntegrationType{
	ASANAINTEGRATIONTYPE_ASANA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AsanaIntegrationType) GetAllowedValues() []AsanaIntegrationType {
	return allowedAsanaIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AsanaIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AsanaIntegrationType(value)
	return nil
}

// NewAsanaIntegrationTypeFromValue returns a pointer to a valid AsanaIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAsanaIntegrationTypeFromValue(v string) (*AsanaIntegrationType, error) {
	ev := AsanaIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AsanaIntegrationType: valid values are %v", v, allowedAsanaIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AsanaIntegrationType) IsValid() bool {
	for _, existing := range allowedAsanaIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AsanaIntegrationType value.
func (v AsanaIntegrationType) Ptr() *AsanaIntegrationType {
	return &v
}
