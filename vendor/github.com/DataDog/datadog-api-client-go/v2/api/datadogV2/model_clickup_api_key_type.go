// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClickupAPIKeyType The definition of the `ClickupAPIKey` object.
type ClickupAPIKeyType string

// List of ClickupAPIKeyType.
const (
	CLICKUPAPIKEYTYPE_CLICKUPAPIKEY ClickupAPIKeyType = "ClickupAPIKey"
)

var allowedClickupAPIKeyTypeEnumValues = []ClickupAPIKeyType{
	CLICKUPAPIKEYTYPE_CLICKUPAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ClickupAPIKeyType) GetAllowedValues() []ClickupAPIKeyType {
	return allowedClickupAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ClickupAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ClickupAPIKeyType(value)
	return nil
}

// NewClickupAPIKeyTypeFromValue returns a pointer to a valid ClickupAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClickupAPIKeyTypeFromValue(v string) (*ClickupAPIKeyType, error) {
	ev := ClickupAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ClickupAPIKeyType: valid values are %v", v, allowedClickupAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClickupAPIKeyType) IsValid() bool {
	for _, existing := range allowedClickupAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClickupAPIKeyType value.
func (v ClickupAPIKeyType) Ptr() *ClickupAPIKeyType {
	return &v
}
