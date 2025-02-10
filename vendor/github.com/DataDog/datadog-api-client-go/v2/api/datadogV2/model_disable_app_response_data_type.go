// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DisableAppResponseDataType The definition of `DisableAppResponseDataType` object.
type DisableAppResponseDataType string

// List of DisableAppResponseDataType.
const (
	DISABLEAPPRESPONSEDATATYPE_DEPLOYMENT DisableAppResponseDataType = "deployment"
)

var allowedDisableAppResponseDataTypeEnumValues = []DisableAppResponseDataType{
	DISABLEAPPRESPONSEDATATYPE_DEPLOYMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DisableAppResponseDataType) GetAllowedValues() []DisableAppResponseDataType {
	return allowedDisableAppResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DisableAppResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DisableAppResponseDataType(value)
	return nil
}

// NewDisableAppResponseDataTypeFromValue returns a pointer to a valid DisableAppResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDisableAppResponseDataTypeFromValue(v string) (*DisableAppResponseDataType, error) {
	ev := DisableAppResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DisableAppResponseDataType: valid values are %v", v, allowedDisableAppResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DisableAppResponseDataType) IsValid() bool {
	for _, existing := range allowedDisableAppResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DisableAppResponseDataType value.
func (v DisableAppResponseDataType) Ptr() *DisableAppResponseDataType {
	return &v
}
