// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIAPIKeyType The definition of the `OpenAIAPIKey` object.
type OpenAIAPIKeyType string

// List of OpenAIAPIKeyType.
const (
	OPENAIAPIKEYTYPE_OPENAIAPIKEY OpenAIAPIKeyType = "OpenAIAPIKey"
)

var allowedOpenAIAPIKeyTypeEnumValues = []OpenAIAPIKeyType{
	OPENAIAPIKEYTYPE_OPENAIAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OpenAIAPIKeyType) GetAllowedValues() []OpenAIAPIKeyType {
	return allowedOpenAIAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpenAIAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpenAIAPIKeyType(value)
	return nil
}

// NewOpenAIAPIKeyTypeFromValue returns a pointer to a valid OpenAIAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpenAIAPIKeyTypeFromValue(v string) (*OpenAIAPIKeyType, error) {
	ev := OpenAIAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpenAIAPIKeyType: valid values are %v", v, allowedOpenAIAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpenAIAPIKeyType) IsValid() bool {
	for _, existing := range allowedOpenAIAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpenAIAPIKeyType value.
func (v OpenAIAPIKeyType) Ptr() *OpenAIAPIKeyType {
	return &v
}
