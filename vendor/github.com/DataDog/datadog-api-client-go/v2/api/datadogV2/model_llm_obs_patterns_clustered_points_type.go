// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsClusteredPointsType Resource type of an LLM Observability patterns clustered points response.
type LLMObsPatternsClusteredPointsType string

// List of LLMObsPatternsClusteredPointsType.
const (
	LLMOBSPATTERNSCLUSTEREDPOINTSTYPE_CLUSTERED_POINTS_RESPONSE LLMObsPatternsClusteredPointsType = "clustered_points_response"
)

var allowedLLMObsPatternsClusteredPointsTypeEnumValues = []LLMObsPatternsClusteredPointsType{
	LLMOBSPATTERNSCLUSTEREDPOINTSTYPE_CLUSTERED_POINTS_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPatternsClusteredPointsType) GetAllowedValues() []LLMObsPatternsClusteredPointsType {
	return allowedLLMObsPatternsClusteredPointsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPatternsClusteredPointsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPatternsClusteredPointsType(value)
	return nil
}

// NewLLMObsPatternsClusteredPointsTypeFromValue returns a pointer to a valid LLMObsPatternsClusteredPointsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPatternsClusteredPointsTypeFromValue(v string) (*LLMObsPatternsClusteredPointsType, error) {
	ev := LLMObsPatternsClusteredPointsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPatternsClusteredPointsType: valid values are %v", v, allowedLLMObsPatternsClusteredPointsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPatternsClusteredPointsType) IsValid() bool {
	for _, existing := range allowedLLMObsPatternsClusteredPointsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPatternsClusteredPointsType value.
func (v LLMObsPatternsClusteredPointsType) Ptr() *LLMObsPatternsClusteredPointsType {
	return &v
}
