// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineQuotaProcessorOverflowAction The action to take when the quota is exceeded. Options:
// - `drop`: Drop the event.
// - `no_action`: Let the event pass through.
// - `overflow_routing`: Route to an overflow destination.
type ObservabilityPipelineQuotaProcessorOverflowAction string

// List of ObservabilityPipelineQuotaProcessorOverflowAction.
const (
	OBSERVABILITYPIPELINEQUOTAPROCESSOROVERFLOWACTION_DROP             ObservabilityPipelineQuotaProcessorOverflowAction = "drop"
	OBSERVABILITYPIPELINEQUOTAPROCESSOROVERFLOWACTION_NO_ACTION        ObservabilityPipelineQuotaProcessorOverflowAction = "no_action"
	OBSERVABILITYPIPELINEQUOTAPROCESSOROVERFLOWACTION_OVERFLOW_ROUTING ObservabilityPipelineQuotaProcessorOverflowAction = "overflow_routing"
)

var allowedObservabilityPipelineQuotaProcessorOverflowActionEnumValues = []ObservabilityPipelineQuotaProcessorOverflowAction{
	OBSERVABILITYPIPELINEQUOTAPROCESSOROVERFLOWACTION_DROP,
	OBSERVABILITYPIPELINEQUOTAPROCESSOROVERFLOWACTION_NO_ACTION,
	OBSERVABILITYPIPELINEQUOTAPROCESSOROVERFLOWACTION_OVERFLOW_ROUTING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineQuotaProcessorOverflowAction) GetAllowedValues() []ObservabilityPipelineQuotaProcessorOverflowAction {
	return allowedObservabilityPipelineQuotaProcessorOverflowActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineQuotaProcessorOverflowAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineQuotaProcessorOverflowAction(value)
	return nil
}

// NewObservabilityPipelineQuotaProcessorOverflowActionFromValue returns a pointer to a valid ObservabilityPipelineQuotaProcessorOverflowAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineQuotaProcessorOverflowActionFromValue(v string) (*ObservabilityPipelineQuotaProcessorOverflowAction, error) {
	ev := ObservabilityPipelineQuotaProcessorOverflowAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineQuotaProcessorOverflowAction: valid values are %v", v, allowedObservabilityPipelineQuotaProcessorOverflowActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineQuotaProcessorOverflowAction) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineQuotaProcessorOverflowActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineQuotaProcessorOverflowAction value.
func (v ObservabilityPipelineQuotaProcessorOverflowAction) Ptr() *ObservabilityPipelineQuotaProcessorOverflowAction {
	return &v
}
