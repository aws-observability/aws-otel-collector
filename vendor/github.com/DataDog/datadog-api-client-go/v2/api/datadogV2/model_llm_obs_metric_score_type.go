// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsMetricScoreType Type of metric recorded for an LLM Observability experiment.
type LLMObsMetricScoreType string

// List of LLMObsMetricScoreType.
const (
	LLMOBSMETRICSCORETYPE_SCORE       LLMObsMetricScoreType = "score"
	LLMOBSMETRICSCORETYPE_CATEGORICAL LLMObsMetricScoreType = "categorical"
	LLMOBSMETRICSCORETYPE_BOOLEAN     LLMObsMetricScoreType = "boolean"
	LLMOBSMETRICSCORETYPE_JSON        LLMObsMetricScoreType = "json"
)

var allowedLLMObsMetricScoreTypeEnumValues = []LLMObsMetricScoreType{
	LLMOBSMETRICSCORETYPE_SCORE,
	LLMOBSMETRICSCORETYPE_CATEGORICAL,
	LLMOBSMETRICSCORETYPE_BOOLEAN,
	LLMOBSMETRICSCORETYPE_JSON,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsMetricScoreType) GetAllowedValues() []LLMObsMetricScoreType {
	return allowedLLMObsMetricScoreTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsMetricScoreType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsMetricScoreType(value)
	return nil
}

// NewLLMObsMetricScoreTypeFromValue returns a pointer to a valid LLMObsMetricScoreType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsMetricScoreTypeFromValue(v string) (*LLMObsMetricScoreType, error) {
	ev := LLMObsMetricScoreType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsMetricScoreType: valid values are %v", v, allowedLLMObsMetricScoreTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsMetricScoreType) IsValid() bool {
	for _, existing := range allowedLLMObsMetricScoreTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsMetricScoreType value.
func (v LLMObsMetricScoreType) Ptr() *LLMObsMetricScoreType {
	return &v
}
