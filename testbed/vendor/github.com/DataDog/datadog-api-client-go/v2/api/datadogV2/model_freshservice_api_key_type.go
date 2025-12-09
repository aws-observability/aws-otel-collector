// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FreshserviceAPIKeyType The definition of the `FreshserviceAPIKey` object.
type FreshserviceAPIKeyType string

// List of FreshserviceAPIKeyType.
const (
	FRESHSERVICEAPIKEYTYPE_FRESHSERVICEAPIKEY FreshserviceAPIKeyType = "FreshserviceAPIKey"
)

var allowedFreshserviceAPIKeyTypeEnumValues = []FreshserviceAPIKeyType{
	FRESHSERVICEAPIKEYTYPE_FRESHSERVICEAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FreshserviceAPIKeyType) GetAllowedValues() []FreshserviceAPIKeyType {
	return allowedFreshserviceAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FreshserviceAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FreshserviceAPIKeyType(value)
	return nil
}

// NewFreshserviceAPIKeyTypeFromValue returns a pointer to a valid FreshserviceAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFreshserviceAPIKeyTypeFromValue(v string) (*FreshserviceAPIKeyType, error) {
	ev := FreshserviceAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FreshserviceAPIKeyType: valid values are %v", v, allowedFreshserviceAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FreshserviceAPIKeyType) IsValid() bool {
	for _, existing := range allowedFreshserviceAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FreshserviceAPIKeyType value.
func (v FreshserviceAPIKeyType) Ptr() *FreshserviceAPIKeyType {
	return &v
}
