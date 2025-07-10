// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineQuotaProcessorLimitEnforceType Unit for quota enforcement in bytes for data size or events for count.
type ObservabilityPipelineQuotaProcessorLimitEnforceType string

// List of ObservabilityPipelineQuotaProcessorLimitEnforceType.
const (
	OBSERVABILITYPIPELINEQUOTAPROCESSORLIMITENFORCETYPE_BYTES  ObservabilityPipelineQuotaProcessorLimitEnforceType = "bytes"
	OBSERVABILITYPIPELINEQUOTAPROCESSORLIMITENFORCETYPE_EVENTS ObservabilityPipelineQuotaProcessorLimitEnforceType = "events"
)

var allowedObservabilityPipelineQuotaProcessorLimitEnforceTypeEnumValues = []ObservabilityPipelineQuotaProcessorLimitEnforceType{
	OBSERVABILITYPIPELINEQUOTAPROCESSORLIMITENFORCETYPE_BYTES,
	OBSERVABILITYPIPELINEQUOTAPROCESSORLIMITENFORCETYPE_EVENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineQuotaProcessorLimitEnforceType) GetAllowedValues() []ObservabilityPipelineQuotaProcessorLimitEnforceType {
	return allowedObservabilityPipelineQuotaProcessorLimitEnforceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineQuotaProcessorLimitEnforceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineQuotaProcessorLimitEnforceType(value)
	return nil
}

// NewObservabilityPipelineQuotaProcessorLimitEnforceTypeFromValue returns a pointer to a valid ObservabilityPipelineQuotaProcessorLimitEnforceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineQuotaProcessorLimitEnforceTypeFromValue(v string) (*ObservabilityPipelineQuotaProcessorLimitEnforceType, error) {
	ev := ObservabilityPipelineQuotaProcessorLimitEnforceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineQuotaProcessorLimitEnforceType: valid values are %v", v, allowedObservabilityPipelineQuotaProcessorLimitEnforceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineQuotaProcessorLimitEnforceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineQuotaProcessorLimitEnforceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineQuotaProcessorLimitEnforceType value.
func (v ObservabilityPipelineQuotaProcessorLimitEnforceType) Ptr() *ObservabilityPipelineQuotaProcessorLimitEnforceType {
	return &v
}
