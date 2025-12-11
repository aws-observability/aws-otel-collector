// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClickupIntegrationType The definition of the `ClickupIntegrationType` object.
type ClickupIntegrationType string

// List of ClickupIntegrationType.
const (
	CLICKUPINTEGRATIONTYPE_CLICKUP ClickupIntegrationType = "Clickup"
)

var allowedClickupIntegrationTypeEnumValues = []ClickupIntegrationType{
	CLICKUPINTEGRATIONTYPE_CLICKUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ClickupIntegrationType) GetAllowedValues() []ClickupIntegrationType {
	return allowedClickupIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ClickupIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ClickupIntegrationType(value)
	return nil
}

// NewClickupIntegrationTypeFromValue returns a pointer to a valid ClickupIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClickupIntegrationTypeFromValue(v string) (*ClickupIntegrationType, error) {
	ev := ClickupIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ClickupIntegrationType: valid values are %v", v, allowedClickupIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClickupIntegrationType) IsValid() bool {
	for _, existing := range allowedClickupIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClickupIntegrationType value.
func (v ClickupIntegrationType) Ptr() *ClickupIntegrationType {
	return &v
}
