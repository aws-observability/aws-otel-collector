// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LaunchDarklyIntegrationType The definition of the `LaunchDarklyIntegrationType` object.
type LaunchDarklyIntegrationType string

// List of LaunchDarklyIntegrationType.
const (
	LAUNCHDARKLYINTEGRATIONTYPE_LAUNCHDARKLY LaunchDarklyIntegrationType = "LaunchDarkly"
)

var allowedLaunchDarklyIntegrationTypeEnumValues = []LaunchDarklyIntegrationType{
	LAUNCHDARKLYINTEGRATIONTYPE_LAUNCHDARKLY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LaunchDarklyIntegrationType) GetAllowedValues() []LaunchDarklyIntegrationType {
	return allowedLaunchDarklyIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LaunchDarklyIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LaunchDarklyIntegrationType(value)
	return nil
}

// NewLaunchDarklyIntegrationTypeFromValue returns a pointer to a valid LaunchDarklyIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLaunchDarklyIntegrationTypeFromValue(v string) (*LaunchDarklyIntegrationType, error) {
	ev := LaunchDarklyIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LaunchDarklyIntegrationType: valid values are %v", v, allowedLaunchDarklyIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LaunchDarklyIntegrationType) IsValid() bool {
	for _, existing := range allowedLaunchDarklyIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LaunchDarklyIntegrationType value.
func (v LaunchDarklyIntegrationType) Ptr() *LaunchDarklyIntegrationType {
	return &v
}
