// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityIntegrationConfigType JSON:API resource type for an entity integration configuration. Always `entity_integration_configs`.
type EntityIntegrationConfigType string

// List of EntityIntegrationConfigType.
const (
	ENTITYINTEGRATIONCONFIGTYPE_ENTITY_INTEGRATION_CONFIGS EntityIntegrationConfigType = "entity_integration_configs"
)

var allowedEntityIntegrationConfigTypeEnumValues = []EntityIntegrationConfigType{
	ENTITYINTEGRATIONCONFIGTYPE_ENTITY_INTEGRATION_CONFIGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityIntegrationConfigType) GetAllowedValues() []EntityIntegrationConfigType {
	return allowedEntityIntegrationConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityIntegrationConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityIntegrationConfigType(value)
	return nil
}

// NewEntityIntegrationConfigTypeFromValue returns a pointer to a valid EntityIntegrationConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityIntegrationConfigTypeFromValue(v string) (*EntityIntegrationConfigType, error) {
	ev := EntityIntegrationConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityIntegrationConfigType: valid values are %v", v, allowedEntityIntegrationConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityIntegrationConfigType) IsValid() bool {
	for _, existing := range allowedEntityIntegrationConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityIntegrationConfigType value.
func (v EntityIntegrationConfigType) Ptr() *EntityIntegrationConfigType {
	return &v
}
