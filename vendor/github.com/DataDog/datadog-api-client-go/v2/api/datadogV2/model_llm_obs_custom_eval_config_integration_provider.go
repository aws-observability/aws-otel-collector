// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigIntegrationProvider Name of the LLM integration provider.
type LLMObsCustomEvalConfigIntegrationProvider string

// List of LLMObsCustomEvalConfigIntegrationProvider.
const (
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_OPENAI         LLMObsCustomEvalConfigIntegrationProvider = "openai"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_AMAZON_BEDROCK LLMObsCustomEvalConfigIntegrationProvider = "amazon-bedrock"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_ANTHROPIC      LLMObsCustomEvalConfigIntegrationProvider = "anthropic"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_AZURE_OPENAI   LLMObsCustomEvalConfigIntegrationProvider = "azure-openai"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_VERTEX_AI      LLMObsCustomEvalConfigIntegrationProvider = "vertex-ai"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_LLM_PROXY      LLMObsCustomEvalConfigIntegrationProvider = "llm-proxy"
)

var allowedLLMObsCustomEvalConfigIntegrationProviderEnumValues = []LLMObsCustomEvalConfigIntegrationProvider{
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_OPENAI,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_AMAZON_BEDROCK,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_ANTHROPIC,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_AZURE_OPENAI,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_VERTEX_AI,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_LLM_PROXY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsCustomEvalConfigIntegrationProvider) GetAllowedValues() []LLMObsCustomEvalConfigIntegrationProvider {
	return allowedLLMObsCustomEvalConfigIntegrationProviderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsCustomEvalConfigIntegrationProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsCustomEvalConfigIntegrationProvider(value)
	return nil
}

// NewLLMObsCustomEvalConfigIntegrationProviderFromValue returns a pointer to a valid LLMObsCustomEvalConfigIntegrationProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsCustomEvalConfigIntegrationProviderFromValue(v string) (*LLMObsCustomEvalConfigIntegrationProvider, error) {
	ev := LLMObsCustomEvalConfigIntegrationProvider(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsCustomEvalConfigIntegrationProvider: valid values are %v", v, allowedLLMObsCustomEvalConfigIntegrationProviderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsCustomEvalConfigIntegrationProvider) IsValid() bool {
	for _, existing := range allowedLLMObsCustomEvalConfigIntegrationProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsCustomEvalConfigIntegrationProvider value.
func (v LLMObsCustomEvalConfigIntegrationProvider) Ptr() *LLMObsCustomEvalConfigIntegrationProvider {
	return &v
}
