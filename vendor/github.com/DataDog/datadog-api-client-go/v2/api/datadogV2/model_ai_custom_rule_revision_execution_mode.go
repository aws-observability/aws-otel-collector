// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiCustomRuleRevisionExecutionMode The execution mode for an AI rule revision.
type AiCustomRuleRevisionExecutionMode string

// List of AiCustomRuleRevisionExecutionMode.
const (
	AICUSTOMRULEREVISIONEXECUTIONMODE_AUTO   AiCustomRuleRevisionExecutionMode = "auto"
	AICUSTOMRULEREVISIONEXECUTIONMODE_MANUAL AiCustomRuleRevisionExecutionMode = "manual"
	AICUSTOMRULEREVISIONEXECUTIONMODE_ALWAYS AiCustomRuleRevisionExecutionMode = "always"
)

var allowedAiCustomRuleRevisionExecutionModeEnumValues = []AiCustomRuleRevisionExecutionMode{
	AICUSTOMRULEREVISIONEXECUTIONMODE_AUTO,
	AICUSTOMRULEREVISIONEXECUTIONMODE_MANUAL,
	AICUSTOMRULEREVISIONEXECUTIONMODE_ALWAYS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiCustomRuleRevisionExecutionMode) GetAllowedValues() []AiCustomRuleRevisionExecutionMode {
	return allowedAiCustomRuleRevisionExecutionModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiCustomRuleRevisionExecutionMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiCustomRuleRevisionExecutionMode(value)
	return nil
}

// NewAiCustomRuleRevisionExecutionModeFromValue returns a pointer to a valid AiCustomRuleRevisionExecutionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiCustomRuleRevisionExecutionModeFromValue(v string) (*AiCustomRuleRevisionExecutionMode, error) {
	ev := AiCustomRuleRevisionExecutionMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiCustomRuleRevisionExecutionMode: valid values are %v", v, allowedAiCustomRuleRevisionExecutionModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiCustomRuleRevisionExecutionMode) IsValid() bool {
	for _, existing := range allowedAiCustomRuleRevisionExecutionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiCustomRuleRevisionExecutionMode value.
func (v AiCustomRuleRevisionExecutionMode) Ptr() *AiCustomRuleRevisionExecutionMode {
	return &v
}
