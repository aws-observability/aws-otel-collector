// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsIntegrationName The name of a supported LLM provider integration.
type LLMObsIntegrationName string

// List of LLMObsIntegrationName.
const (
	LLMOBSINTEGRATIONNAME_OPENAI         LLMObsIntegrationName = "openai"
	LLMOBSINTEGRATIONNAME_AMAZON_BEDROCK LLMObsIntegrationName = "amazon_bedrock"
	LLMOBSINTEGRATIONNAME_ANTHROPIC      LLMObsIntegrationName = "anthropic"
	LLMOBSINTEGRATIONNAME_AZURE_OPENAI   LLMObsIntegrationName = "azure_openai"
	LLMOBSINTEGRATIONNAME_VERTEX_AI      LLMObsIntegrationName = "vertex_ai"
	LLMOBSINTEGRATIONNAME_LLMPROXY       LLMObsIntegrationName = "llmproxy"
)

var allowedLLMObsIntegrationNameEnumValues = []LLMObsIntegrationName{
	LLMOBSINTEGRATIONNAME_OPENAI,
	LLMOBSINTEGRATIONNAME_AMAZON_BEDROCK,
	LLMOBSINTEGRATIONNAME_ANTHROPIC,
	LLMOBSINTEGRATIONNAME_AZURE_OPENAI,
	LLMOBSINTEGRATIONNAME_VERTEX_AI,
	LLMOBSINTEGRATIONNAME_LLMPROXY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsIntegrationName) GetAllowedValues() []LLMObsIntegrationName {
	return allowedLLMObsIntegrationNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsIntegrationName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsIntegrationName(value)
	return nil
}

// NewLLMObsIntegrationNameFromValue returns a pointer to a valid LLMObsIntegrationName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsIntegrationNameFromValue(v string) (*LLMObsIntegrationName, error) {
	ev := LLMObsIntegrationName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsIntegrationName: valid values are %v", v, allowedLLMObsIntegrationNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsIntegrationName) IsValid() bool {
	for _, existing := range allowedLLMObsIntegrationNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsIntegrationName value.
func (v LLMObsIntegrationName) Ptr() *LLMObsIntegrationName {
	return &v
}
