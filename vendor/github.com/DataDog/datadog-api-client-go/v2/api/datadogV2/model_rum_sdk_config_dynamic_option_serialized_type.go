// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSdkConfigDynamicOptionSerializedType The type identifier for a dynamic option. Always `dynamic`.
type RumSdkConfigDynamicOptionSerializedType string

// List of RumSdkConfigDynamicOptionSerializedType.
const (
	RUMSDKCONFIGDYNAMICOPTIONSERIALIZEDTYPE_DYNAMIC RumSdkConfigDynamicOptionSerializedType = "dynamic"
)

var allowedRumSdkConfigDynamicOptionSerializedTypeEnumValues = []RumSdkConfigDynamicOptionSerializedType{
	RUMSDKCONFIGDYNAMICOPTIONSERIALIZEDTYPE_DYNAMIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumSdkConfigDynamicOptionSerializedType) GetAllowedValues() []RumSdkConfigDynamicOptionSerializedType {
	return allowedRumSdkConfigDynamicOptionSerializedTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumSdkConfigDynamicOptionSerializedType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumSdkConfigDynamicOptionSerializedType(value)
	return nil
}

// NewRumSdkConfigDynamicOptionSerializedTypeFromValue returns a pointer to a valid RumSdkConfigDynamicOptionSerializedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumSdkConfigDynamicOptionSerializedTypeFromValue(v string) (*RumSdkConfigDynamicOptionSerializedType, error) {
	ev := RumSdkConfigDynamicOptionSerializedType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumSdkConfigDynamicOptionSerializedType: valid values are %v", v, allowedRumSdkConfigDynamicOptionSerializedTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumSdkConfigDynamicOptionSerializedType) IsValid() bool {
	for _, existing := range allowedRumSdkConfigDynamicOptionSerializedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumSdkConfigDynamicOptionSerializedType value.
func (v RumSdkConfigDynamicOptionSerializedType) Ptr() *RumSdkConfigDynamicOptionSerializedType {
	return &v
}
