// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeminiAPIKeyType The definition of the `GeminiAPIKey` object.
type GeminiAPIKeyType string

// List of GeminiAPIKeyType.
const (
	GEMINIAPIKEYTYPE_GEMINIAPIKEY GeminiAPIKeyType = "GeminiAPIKey"
)

var allowedGeminiAPIKeyTypeEnumValues = []GeminiAPIKeyType{
	GEMINIAPIKEYTYPE_GEMINIAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeminiAPIKeyType) GetAllowedValues() []GeminiAPIKeyType {
	return allowedGeminiAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeminiAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeminiAPIKeyType(value)
	return nil
}

// NewGeminiAPIKeyTypeFromValue returns a pointer to a valid GeminiAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeminiAPIKeyTypeFromValue(v string) (*GeminiAPIKeyType, error) {
	ev := GeminiAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeminiAPIKeyType: valid values are %v", v, allowedGeminiAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeminiAPIKeyType) IsValid() bool {
	for _, existing := range allowedGeminiAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeminiAPIKeyType value.
func (v GeminiAPIKeyType) Ptr() *GeminiAPIKeyType {
	return &v
}
