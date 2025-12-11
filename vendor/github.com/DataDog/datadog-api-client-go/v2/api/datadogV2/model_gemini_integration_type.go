// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeminiIntegrationType The definition of the `GeminiIntegrationType` object.
type GeminiIntegrationType string

// List of GeminiIntegrationType.
const (
	GEMINIINTEGRATIONTYPE_GEMINI GeminiIntegrationType = "Gemini"
)

var allowedGeminiIntegrationTypeEnumValues = []GeminiIntegrationType{
	GEMINIINTEGRATIONTYPE_GEMINI,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeminiIntegrationType) GetAllowedValues() []GeminiIntegrationType {
	return allowedGeminiIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeminiIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeminiIntegrationType(value)
	return nil
}

// NewGeminiIntegrationTypeFromValue returns a pointer to a valid GeminiIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeminiIntegrationTypeFromValue(v string) (*GeminiIntegrationType, error) {
	ev := GeminiIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeminiIntegrationType: valid values are %v", v, allowedGeminiIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeminiIntegrationType) IsValid() bool {
	for _, existing := range allowedGeminiIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeminiIntegrationType value.
func (v GeminiIntegrationType) Ptr() *GeminiIntegrationType {
	return &v
}
