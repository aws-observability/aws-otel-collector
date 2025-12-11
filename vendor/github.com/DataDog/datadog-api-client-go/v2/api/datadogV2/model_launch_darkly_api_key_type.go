// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LaunchDarklyAPIKeyType The definition of the `LaunchDarklyAPIKey` object.
type LaunchDarklyAPIKeyType string

// List of LaunchDarklyAPIKeyType.
const (
	LAUNCHDARKLYAPIKEYTYPE_LAUNCHDARKLYAPIKEY LaunchDarklyAPIKeyType = "LaunchDarklyAPIKey"
)

var allowedLaunchDarklyAPIKeyTypeEnumValues = []LaunchDarklyAPIKeyType{
	LAUNCHDARKLYAPIKEYTYPE_LAUNCHDARKLYAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LaunchDarklyAPIKeyType) GetAllowedValues() []LaunchDarklyAPIKeyType {
	return allowedLaunchDarklyAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LaunchDarklyAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LaunchDarklyAPIKeyType(value)
	return nil
}

// NewLaunchDarklyAPIKeyTypeFromValue returns a pointer to a valid LaunchDarklyAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLaunchDarklyAPIKeyTypeFromValue(v string) (*LaunchDarklyAPIKeyType, error) {
	ev := LaunchDarklyAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LaunchDarklyAPIKeyType: valid values are %v", v, allowedLaunchDarklyAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LaunchDarklyAPIKeyType) IsValid() bool {
	for _, existing := range allowedLaunchDarklyAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LaunchDarklyAPIKeyType value.
func (v LaunchDarklyAPIKeyType) Ptr() *LaunchDarklyAPIKeyType {
	return &v
}
