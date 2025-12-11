// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIIntegrationType The definition of the `OpenAIIntegrationType` object.
type OpenAIIntegrationType string

// List of OpenAIIntegrationType.
const (
	OPENAIINTEGRATIONTYPE_OPENAI OpenAIIntegrationType = "OpenAI"
)

var allowedOpenAIIntegrationTypeEnumValues = []OpenAIIntegrationType{
	OPENAIINTEGRATIONTYPE_OPENAI,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OpenAIIntegrationType) GetAllowedValues() []OpenAIIntegrationType {
	return allowedOpenAIIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpenAIIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpenAIIntegrationType(value)
	return nil
}

// NewOpenAIIntegrationTypeFromValue returns a pointer to a valid OpenAIIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpenAIIntegrationTypeFromValue(v string) (*OpenAIIntegrationType, error) {
	ev := OpenAIIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpenAIIntegrationType: valid values are %v", v, allowedOpenAIIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpenAIIntegrationType) IsValid() bool {
	for _, existing := range allowedOpenAIIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpenAIIntegrationType value.
func (v OpenAIIntegrationType) Ptr() *OpenAIIntegrationType {
	return &v
}
