// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatsigIntegrationType The definition of the `StatsigIntegrationType` object.
type StatsigIntegrationType string

// List of StatsigIntegrationType.
const (
	STATSIGINTEGRATIONTYPE_STATSIG StatsigIntegrationType = "Statsig"
)

var allowedStatsigIntegrationTypeEnumValues = []StatsigIntegrationType{
	STATSIGINTEGRATIONTYPE_STATSIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StatsigIntegrationType) GetAllowedValues() []StatsigIntegrationType {
	return allowedStatsigIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StatsigIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StatsigIntegrationType(value)
	return nil
}

// NewStatsigIntegrationTypeFromValue returns a pointer to a valid StatsigIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStatsigIntegrationTypeFromValue(v string) (*StatsigIntegrationType, error) {
	ev := StatsigIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StatsigIntegrationType: valid values are %v", v, allowedStatsigIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StatsigIntegrationType) IsValid() bool {
	for _, existing := range allowedStatsigIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatsigIntegrationType value.
func (v StatsigIntegrationType) Ptr() *StatsigIntegrationType {
	return &v
}
