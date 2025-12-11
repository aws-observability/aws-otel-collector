// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GreyNoiseAPIKeyType The definition of the `GreyNoiseAPIKey` object.
type GreyNoiseAPIKeyType string

// List of GreyNoiseAPIKeyType.
const (
	GREYNOISEAPIKEYTYPE_GREYNOISEAPIKEY GreyNoiseAPIKeyType = "GreyNoiseAPIKey"
)

var allowedGreyNoiseAPIKeyTypeEnumValues = []GreyNoiseAPIKeyType{
	GREYNOISEAPIKEYTYPE_GREYNOISEAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GreyNoiseAPIKeyType) GetAllowedValues() []GreyNoiseAPIKeyType {
	return allowedGreyNoiseAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GreyNoiseAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GreyNoiseAPIKeyType(value)
	return nil
}

// NewGreyNoiseAPIKeyTypeFromValue returns a pointer to a valid GreyNoiseAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGreyNoiseAPIKeyTypeFromValue(v string) (*GreyNoiseAPIKeyType, error) {
	ev := GreyNoiseAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GreyNoiseAPIKeyType: valid values are %v", v, allowedGreyNoiseAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GreyNoiseAPIKeyType) IsValid() bool {
	for _, existing := range allowedGreyNoiseAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GreyNoiseAPIKeyType value.
func (v GreyNoiseAPIKeyType) Ptr() *GreyNoiseAPIKeyType {
	return &v
}
