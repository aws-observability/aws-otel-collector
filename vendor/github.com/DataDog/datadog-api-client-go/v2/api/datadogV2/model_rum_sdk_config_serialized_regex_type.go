// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSdkConfigSerializedRegexType The type identifier for a serialized regex. Always `regex`.
type RumSdkConfigSerializedRegexType string

// List of RumSdkConfigSerializedRegexType.
const (
	RUMSDKCONFIGSERIALIZEDREGEXTYPE_REGEX RumSdkConfigSerializedRegexType = "regex"
)

var allowedRumSdkConfigSerializedRegexTypeEnumValues = []RumSdkConfigSerializedRegexType{
	RUMSDKCONFIGSERIALIZEDREGEXTYPE_REGEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumSdkConfigSerializedRegexType) GetAllowedValues() []RumSdkConfigSerializedRegexType {
	return allowedRumSdkConfigSerializedRegexTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumSdkConfigSerializedRegexType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumSdkConfigSerializedRegexType(value)
	return nil
}

// NewRumSdkConfigSerializedRegexTypeFromValue returns a pointer to a valid RumSdkConfigSerializedRegexType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumSdkConfigSerializedRegexTypeFromValue(v string) (*RumSdkConfigSerializedRegexType, error) {
	ev := RumSdkConfigSerializedRegexType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumSdkConfigSerializedRegexType: valid values are %v", v, allowedRumSdkConfigSerializedRegexTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumSdkConfigSerializedRegexType) IsValid() bool {
	for _, existing := range allowedRumSdkConfigSerializedRegexTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumSdkConfigSerializedRegexType value.
func (v RumSdkConfigSerializedRegexType) Ptr() *RumSdkConfigSerializedRegexType {
	return &v
}
