// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicAPIKeyType The definition of the `AnthropicAPIKey` object.
type AnthropicAPIKeyType string

// List of AnthropicAPIKeyType.
const (
	ANTHROPICAPIKEYTYPE_ANTHROPICAPIKEY AnthropicAPIKeyType = "AnthropicAPIKey"
)

var allowedAnthropicAPIKeyTypeEnumValues = []AnthropicAPIKeyType{
	ANTHROPICAPIKEYTYPE_ANTHROPICAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnthropicAPIKeyType) GetAllowedValues() []AnthropicAPIKeyType {
	return allowedAnthropicAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnthropicAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnthropicAPIKeyType(value)
	return nil
}

// NewAnthropicAPIKeyTypeFromValue returns a pointer to a valid AnthropicAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnthropicAPIKeyTypeFromValue(v string) (*AnthropicAPIKeyType, error) {
	ev := AnthropicAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnthropicAPIKeyType: valid values are %v", v, allowedAnthropicAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnthropicAPIKeyType) IsValid() bool {
	for _, existing := range allowedAnthropicAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnthropicAPIKeyType value.
func (v AnthropicAPIKeyType) Ptr() *AnthropicAPIKeyType {
	return &v
}
