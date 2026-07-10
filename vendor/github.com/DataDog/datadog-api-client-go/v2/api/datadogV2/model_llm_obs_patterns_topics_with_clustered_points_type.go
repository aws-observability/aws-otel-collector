// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsTopicsWithClusteredPointsType Resource type of an LLM Observability patterns topics-with-clustered-points response.
type LLMObsPatternsTopicsWithClusteredPointsType string

// List of LLMObsPatternsTopicsWithClusteredPointsType.
const (
	LLMOBSPATTERNSTOPICSWITHCLUSTEREDPOINTSTYPE_GET_TOPICS_WITH_CLUSTER_POINTS_RESPONSE LLMObsPatternsTopicsWithClusteredPointsType = "get_topics_with_cluster_points_response"
)

var allowedLLMObsPatternsTopicsWithClusteredPointsTypeEnumValues = []LLMObsPatternsTopicsWithClusteredPointsType{
	LLMOBSPATTERNSTOPICSWITHCLUSTEREDPOINTSTYPE_GET_TOPICS_WITH_CLUSTER_POINTS_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPatternsTopicsWithClusteredPointsType) GetAllowedValues() []LLMObsPatternsTopicsWithClusteredPointsType {
	return allowedLLMObsPatternsTopicsWithClusteredPointsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPatternsTopicsWithClusteredPointsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPatternsTopicsWithClusteredPointsType(value)
	return nil
}

// NewLLMObsPatternsTopicsWithClusteredPointsTypeFromValue returns a pointer to a valid LLMObsPatternsTopicsWithClusteredPointsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPatternsTopicsWithClusteredPointsTypeFromValue(v string) (*LLMObsPatternsTopicsWithClusteredPointsType, error) {
	ev := LLMObsPatternsTopicsWithClusteredPointsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPatternsTopicsWithClusteredPointsType: valid values are %v", v, allowedLLMObsPatternsTopicsWithClusteredPointsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPatternsTopicsWithClusteredPointsType) IsValid() bool {
	for _, existing := range allowedLLMObsPatternsTopicsWithClusteredPointsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPatternsTopicsWithClusteredPointsType value.
func (v LLMObsPatternsTopicsWithClusteredPointsType) Ptr() *LLMObsPatternsTopicsWithClusteredPointsType {
	return &v
}
