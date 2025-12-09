// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitAPIKeyType The definition of the `SplitAPIKey` object.
type SplitAPIKeyType string

// List of SplitAPIKeyType.
const (
	SPLITAPIKEYTYPE_SPLITAPIKEY SplitAPIKeyType = "SplitAPIKey"
)

var allowedSplitAPIKeyTypeEnumValues = []SplitAPIKeyType{
	SPLITAPIKEYTYPE_SPLITAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SplitAPIKeyType) GetAllowedValues() []SplitAPIKeyType {
	return allowedSplitAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SplitAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SplitAPIKeyType(value)
	return nil
}

// NewSplitAPIKeyTypeFromValue returns a pointer to a valid SplitAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSplitAPIKeyTypeFromValue(v string) (*SplitAPIKeyType, error) {
	ev := SplitAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SplitAPIKeyType: valid values are %v", v, allowedSplitAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SplitAPIKeyType) IsValid() bool {
	for _, existing := range allowedSplitAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SplitAPIKeyType value.
func (v SplitAPIKeyType) Ptr() *SplitAPIKeyType {
	return &v
}
