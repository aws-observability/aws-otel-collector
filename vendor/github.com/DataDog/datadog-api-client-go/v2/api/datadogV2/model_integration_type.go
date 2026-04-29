// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationType Integration resource type.
type IntegrationType string

// List of IntegrationType.
const (
	INTEGRATIONTYPE_INTEGRATION IntegrationType = "integration"
)

var allowedIntegrationTypeEnumValues = []IntegrationType{
	INTEGRATIONTYPE_INTEGRATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IntegrationType) GetAllowedValues() []IntegrationType {
	return allowedIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IntegrationType(value)
	return nil
}

// NewIntegrationTypeFromValue returns a pointer to a valid IntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIntegrationTypeFromValue(v string) (*IntegrationType, error) {
	ev := IntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IntegrationType: valid values are %v", v, allowedIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IntegrationType) IsValid() bool {
	for _, existing := range allowedIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationType value.
func (v IntegrationType) Ptr() *IntegrationType {
	return &v
}
