// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSdkConfigType The type of the resource. The value should always be `rum_sdk_config`.
type RumSdkConfigType string

// List of RumSdkConfigType.
const (
	RUMSDKCONFIGTYPE_RUM_SDK_CONFIG RumSdkConfigType = "rum_sdk_config"
)

var allowedRumSdkConfigTypeEnumValues = []RumSdkConfigType{
	RUMSDKCONFIGTYPE_RUM_SDK_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumSdkConfigType) GetAllowedValues() []RumSdkConfigType {
	return allowedRumSdkConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumSdkConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumSdkConfigType(value)
	return nil
}

// NewRumSdkConfigTypeFromValue returns a pointer to a valid RumSdkConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumSdkConfigTypeFromValue(v string) (*RumSdkConfigType, error) {
	ev := RumSdkConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumSdkConfigType: valid values are %v", v, allowedRumSdkConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumSdkConfigType) IsValid() bool {
	for _, existing := range allowedRumSdkConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumSdkConfigType value.
func (v RumSdkConfigType) Ptr() *RumSdkConfigType {
	return &v
}
