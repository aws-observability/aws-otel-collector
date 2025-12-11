// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfigCatSDKKeyType The definition of the `ConfigCatSDKKey` object.
type ConfigCatSDKKeyType string

// List of ConfigCatSDKKeyType.
const (
	CONFIGCATSDKKEYTYPE_CONFIGCATSDKKEY ConfigCatSDKKeyType = "ConfigCatSDKKey"
)

var allowedConfigCatSDKKeyTypeEnumValues = []ConfigCatSDKKeyType{
	CONFIGCATSDKKEYTYPE_CONFIGCATSDKKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ConfigCatSDKKeyType) GetAllowedValues() []ConfigCatSDKKeyType {
	return allowedConfigCatSDKKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ConfigCatSDKKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ConfigCatSDKKeyType(value)
	return nil
}

// NewConfigCatSDKKeyTypeFromValue returns a pointer to a valid ConfigCatSDKKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewConfigCatSDKKeyTypeFromValue(v string) (*ConfigCatSDKKeyType, error) {
	ev := ConfigCatSDKKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConfigCatSDKKeyType: valid values are %v", v, allowedConfigCatSDKKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ConfigCatSDKKeyType) IsValid() bool {
	for _, existing := range allowedConfigCatSDKKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConfigCatSDKKeyType value.
func (v ConfigCatSDKKeyType) Ptr() *ConfigCatSDKKeyType {
	return &v
}
