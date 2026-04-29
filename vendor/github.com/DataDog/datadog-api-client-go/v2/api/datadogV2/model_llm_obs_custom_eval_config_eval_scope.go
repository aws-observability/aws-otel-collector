// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigEvalScope Scope at which to evaluate spans.
type LLMObsCustomEvalConfigEvalScope string

// List of LLMObsCustomEvalConfigEvalScope.
const (
	LLMOBSCUSTOMEVALCONFIGEVALSCOPE_SPAN    LLMObsCustomEvalConfigEvalScope = "span"
	LLMOBSCUSTOMEVALCONFIGEVALSCOPE_TRACE   LLMObsCustomEvalConfigEvalScope = "trace"
	LLMOBSCUSTOMEVALCONFIGEVALSCOPE_SESSION LLMObsCustomEvalConfigEvalScope = "session"
)

var allowedLLMObsCustomEvalConfigEvalScopeEnumValues = []LLMObsCustomEvalConfigEvalScope{
	LLMOBSCUSTOMEVALCONFIGEVALSCOPE_SPAN,
	LLMOBSCUSTOMEVALCONFIGEVALSCOPE_TRACE,
	LLMOBSCUSTOMEVALCONFIGEVALSCOPE_SESSION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsCustomEvalConfigEvalScope) GetAllowedValues() []LLMObsCustomEvalConfigEvalScope {
	return allowedLLMObsCustomEvalConfigEvalScopeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsCustomEvalConfigEvalScope) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsCustomEvalConfigEvalScope(value)
	return nil
}

// NewLLMObsCustomEvalConfigEvalScopeFromValue returns a pointer to a valid LLMObsCustomEvalConfigEvalScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsCustomEvalConfigEvalScopeFromValue(v string) (*LLMObsCustomEvalConfigEvalScope, error) {
	ev := LLMObsCustomEvalConfigEvalScope(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsCustomEvalConfigEvalScope: valid values are %v", v, allowedLLMObsCustomEvalConfigEvalScopeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsCustomEvalConfigEvalScope) IsValid() bool {
	for _, existing := range allowedLLMObsCustomEvalConfigEvalScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsCustomEvalConfigEvalScope value.
func (v LLMObsCustomEvalConfigEvalScope) Ptr() *LLMObsCustomEvalConfigEvalScope {
	return &v
}
