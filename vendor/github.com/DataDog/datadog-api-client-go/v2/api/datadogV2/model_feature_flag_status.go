// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagStatus The status of a feature flag in an environment.
type FeatureFlagStatus string

// List of FeatureFlagStatus.
const (
	FEATUREFLAGSTATUS_ENABLED  FeatureFlagStatus = "ENABLED"
	FEATUREFLAGSTATUS_DISABLED FeatureFlagStatus = "DISABLED"
)

var allowedFeatureFlagStatusEnumValues = []FeatureFlagStatus{
	FEATUREFLAGSTATUS_ENABLED,
	FEATUREFLAGSTATUS_DISABLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagStatus) GetAllowedValues() []FeatureFlagStatus {
	return allowedFeatureFlagStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagStatus(value)
	return nil
}

// NewFeatureFlagStatusFromValue returns a pointer to a valid FeatureFlagStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagStatusFromValue(v string) (*FeatureFlagStatus, error) {
	ev := FeatureFlagStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagStatus: valid values are %v", v, allowedFeatureFlagStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagStatus) IsValid() bool {
	for _, existing := range allowedFeatureFlagStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagStatus value.
func (v FeatureFlagStatus) Ptr() *FeatureFlagStatus {
	return &v
}
