// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateFeatureFlagDataType The resource type.
type UpdateFeatureFlagDataType string

// List of UpdateFeatureFlagDataType.
const (
	UPDATEFEATUREFLAGDATATYPE_FEATURE_FLAGS UpdateFeatureFlagDataType = "feature-flags"
)

var allowedUpdateFeatureFlagDataTypeEnumValues = []UpdateFeatureFlagDataType{
	UPDATEFEATUREFLAGDATATYPE_FEATURE_FLAGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateFeatureFlagDataType) GetAllowedValues() []UpdateFeatureFlagDataType {
	return allowedUpdateFeatureFlagDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateFeatureFlagDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateFeatureFlagDataType(value)
	return nil
}

// NewUpdateFeatureFlagDataTypeFromValue returns a pointer to a valid UpdateFeatureFlagDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateFeatureFlagDataTypeFromValue(v string) (*UpdateFeatureFlagDataType, error) {
	ev := UpdateFeatureFlagDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateFeatureFlagDataType: valid values are %v", v, allowedUpdateFeatureFlagDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateFeatureFlagDataType) IsValid() bool {
	for _, existing := range allowedUpdateFeatureFlagDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateFeatureFlagDataType value.
func (v UpdateFeatureFlagDataType) Ptr() *UpdateFeatureFlagDataType {
	return &v
}
