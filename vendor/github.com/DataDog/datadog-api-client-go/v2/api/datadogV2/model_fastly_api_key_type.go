// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FastlyAPIKeyType The definition of the `FastlyAPIKey` object.
type FastlyAPIKeyType string

// List of FastlyAPIKeyType.
const (
	FASTLYAPIKEYTYPE_FASTLYAPIKEY FastlyAPIKeyType = "FastlyAPIKey"
)

var allowedFastlyAPIKeyTypeEnumValues = []FastlyAPIKeyType{
	FASTLYAPIKEYTYPE_FASTLYAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FastlyAPIKeyType) GetAllowedValues() []FastlyAPIKeyType {
	return allowedFastlyAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FastlyAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FastlyAPIKeyType(value)
	return nil
}

// NewFastlyAPIKeyTypeFromValue returns a pointer to a valid FastlyAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFastlyAPIKeyTypeFromValue(v string) (*FastlyAPIKeyType, error) {
	ev := FastlyAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FastlyAPIKeyType: valid values are %v", v, allowedFastlyAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FastlyAPIKeyType) IsValid() bool {
	for _, existing := range allowedFastlyAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FastlyAPIKeyType value.
func (v FastlyAPIKeyType) Ptr() *FastlyAPIKeyType {
	return &v
}
