// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSdkConfigMatchOptionSerializedType The type of match pattern, either a literal string or a regex.
type RumSdkConfigMatchOptionSerializedType string

// List of RumSdkConfigMatchOptionSerializedType.
const (
	RUMSDKCONFIGMATCHOPTIONSERIALIZEDTYPE_STRING RumSdkConfigMatchOptionSerializedType = "string"
	RUMSDKCONFIGMATCHOPTIONSERIALIZEDTYPE_REGEX  RumSdkConfigMatchOptionSerializedType = "regex"
)

var allowedRumSdkConfigMatchOptionSerializedTypeEnumValues = []RumSdkConfigMatchOptionSerializedType{
	RUMSDKCONFIGMATCHOPTIONSERIALIZEDTYPE_STRING,
	RUMSDKCONFIGMATCHOPTIONSERIALIZEDTYPE_REGEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumSdkConfigMatchOptionSerializedType) GetAllowedValues() []RumSdkConfigMatchOptionSerializedType {
	return allowedRumSdkConfigMatchOptionSerializedTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumSdkConfigMatchOptionSerializedType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumSdkConfigMatchOptionSerializedType(value)
	return nil
}

// NewRumSdkConfigMatchOptionSerializedTypeFromValue returns a pointer to a valid RumSdkConfigMatchOptionSerializedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumSdkConfigMatchOptionSerializedTypeFromValue(v string) (*RumSdkConfigMatchOptionSerializedType, error) {
	ev := RumSdkConfigMatchOptionSerializedType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumSdkConfigMatchOptionSerializedType: valid values are %v", v, allowedRumSdkConfigMatchOptionSerializedTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumSdkConfigMatchOptionSerializedType) IsValid() bool {
	for _, existing := range allowedRumSdkConfigMatchOptionSerializedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumSdkConfigMatchOptionSerializedType value.
func (v RumSdkConfigMatchOptionSerializedType) Ptr() *RumSdkConfigMatchOptionSerializedType {
	return &v
}
