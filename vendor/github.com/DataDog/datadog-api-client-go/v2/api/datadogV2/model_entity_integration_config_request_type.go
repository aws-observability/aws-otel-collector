// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityIntegrationConfigRequestType JSON:API resource type for the entity integration configuration create or update request. Always `entity_integration_config_requests`.
type EntityIntegrationConfigRequestType string

// List of EntityIntegrationConfigRequestType.
const (
	ENTITYINTEGRATIONCONFIGREQUESTTYPE_ENTITY_INTEGRATION_CONFIG_REQUESTS EntityIntegrationConfigRequestType = "entity_integration_config_requests"
)

var allowedEntityIntegrationConfigRequestTypeEnumValues = []EntityIntegrationConfigRequestType{
	ENTITYINTEGRATIONCONFIGREQUESTTYPE_ENTITY_INTEGRATION_CONFIG_REQUESTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityIntegrationConfigRequestType) GetAllowedValues() []EntityIntegrationConfigRequestType {
	return allowedEntityIntegrationConfigRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityIntegrationConfigRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityIntegrationConfigRequestType(value)
	return nil
}

// NewEntityIntegrationConfigRequestTypeFromValue returns a pointer to a valid EntityIntegrationConfigRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityIntegrationConfigRequestTypeFromValue(v string) (*EntityIntegrationConfigRequestType, error) {
	ev := EntityIntegrationConfigRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityIntegrationConfigRequestType: valid values are %v", v, allowedEntityIntegrationConfigRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityIntegrationConfigRequestType) IsValid() bool {
	for _, existing := range allowedEntityIntegrationConfigRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityIntegrationConfigRequestType value.
func (v EntityIntegrationConfigRequestType) Ptr() *EntityIntegrationConfigRequestType {
	return &v
}
