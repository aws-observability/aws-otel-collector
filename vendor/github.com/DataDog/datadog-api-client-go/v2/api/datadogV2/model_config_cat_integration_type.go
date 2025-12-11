// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfigCatIntegrationType The definition of the `ConfigCatIntegrationType` object.
type ConfigCatIntegrationType string

// List of ConfigCatIntegrationType.
const (
	CONFIGCATINTEGRATIONTYPE_CONFIGCAT ConfigCatIntegrationType = "ConfigCat"
)

var allowedConfigCatIntegrationTypeEnumValues = []ConfigCatIntegrationType{
	CONFIGCATINTEGRATIONTYPE_CONFIGCAT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ConfigCatIntegrationType) GetAllowedValues() []ConfigCatIntegrationType {
	return allowedConfigCatIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ConfigCatIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ConfigCatIntegrationType(value)
	return nil
}

// NewConfigCatIntegrationTypeFromValue returns a pointer to a valid ConfigCatIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewConfigCatIntegrationTypeFromValue(v string) (*ConfigCatIntegrationType, error) {
	ev := ConfigCatIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConfigCatIntegrationType: valid values are %v", v, allowedConfigCatIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ConfigCatIntegrationType) IsValid() bool {
	for _, existing := range allowedConfigCatIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConfigCatIntegrationType value.
func (v ConfigCatIntegrationType) Ptr() *ConfigCatIntegrationType {
	return &v
}
