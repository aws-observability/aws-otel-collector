// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateFeatureFlagDataType The resource type.
type CreateFeatureFlagDataType string

// List of CreateFeatureFlagDataType.
const (
	CREATEFEATUREFLAGDATATYPE_FEATURE_FLAGS CreateFeatureFlagDataType = "feature-flags"
)

var allowedCreateFeatureFlagDataTypeEnumValues = []CreateFeatureFlagDataType{
	CREATEFEATUREFLAGDATATYPE_FEATURE_FLAGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateFeatureFlagDataType) GetAllowedValues() []CreateFeatureFlagDataType {
	return allowedCreateFeatureFlagDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateFeatureFlagDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateFeatureFlagDataType(value)
	return nil
}

// NewCreateFeatureFlagDataTypeFromValue returns a pointer to a valid CreateFeatureFlagDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateFeatureFlagDataTypeFromValue(v string) (*CreateFeatureFlagDataType, error) {
	ev := CreateFeatureFlagDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateFeatureFlagDataType: valid values are %v", v, allowedCreateFeatureFlagDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateFeatureFlagDataType) IsValid() bool {
	for _, existing := range allowedCreateFeatureFlagDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateFeatureFlagDataType value.
func (v CreateFeatureFlagDataType) Ptr() *CreateFeatureFlagDataType {
	return &v
}
