// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineTagCardinalityLimitProcessorAction The action to take when the cardinality limit is exceeded.
type ObservabilityPipelineTagCardinalityLimitProcessorAction string

// List of ObservabilityPipelineTagCardinalityLimitProcessorAction.
const (
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORACTION_DROP_TAG   ObservabilityPipelineTagCardinalityLimitProcessorAction = "drop_tag"
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORACTION_DROP_EVENT ObservabilityPipelineTagCardinalityLimitProcessorAction = "drop_event"
)

var allowedObservabilityPipelineTagCardinalityLimitProcessorActionEnumValues = []ObservabilityPipelineTagCardinalityLimitProcessorAction{
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORACTION_DROP_TAG,
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORACTION_DROP_EVENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineTagCardinalityLimitProcessorAction) GetAllowedValues() []ObservabilityPipelineTagCardinalityLimitProcessorAction {
	return allowedObservabilityPipelineTagCardinalityLimitProcessorActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineTagCardinalityLimitProcessorAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineTagCardinalityLimitProcessorAction(value)
	return nil
}

// NewObservabilityPipelineTagCardinalityLimitProcessorActionFromValue returns a pointer to a valid ObservabilityPipelineTagCardinalityLimitProcessorAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineTagCardinalityLimitProcessorActionFromValue(v string) (*ObservabilityPipelineTagCardinalityLimitProcessorAction, error) {
	ev := ObservabilityPipelineTagCardinalityLimitProcessorAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineTagCardinalityLimitProcessorAction: valid values are %v", v, allowedObservabilityPipelineTagCardinalityLimitProcessorActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineTagCardinalityLimitProcessorAction) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineTagCardinalityLimitProcessorActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineTagCardinalityLimitProcessorAction value.
func (v ObservabilityPipelineTagCardinalityLimitProcessorAction) Ptr() *ObservabilityPipelineTagCardinalityLimitProcessorAction {
	return &v
}
