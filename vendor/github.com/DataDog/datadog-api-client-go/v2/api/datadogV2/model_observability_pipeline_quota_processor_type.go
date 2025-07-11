// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineQuotaProcessorType The processor type. The value should always be `quota`.
type ObservabilityPipelineQuotaProcessorType string

// List of ObservabilityPipelineQuotaProcessorType.
const (
	OBSERVABILITYPIPELINEQUOTAPROCESSORTYPE_QUOTA ObservabilityPipelineQuotaProcessorType = "quota"
)

var allowedObservabilityPipelineQuotaProcessorTypeEnumValues = []ObservabilityPipelineQuotaProcessorType{
	OBSERVABILITYPIPELINEQUOTAPROCESSORTYPE_QUOTA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineQuotaProcessorType) GetAllowedValues() []ObservabilityPipelineQuotaProcessorType {
	return allowedObservabilityPipelineQuotaProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineQuotaProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineQuotaProcessorType(value)
	return nil
}

// NewObservabilityPipelineQuotaProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineQuotaProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineQuotaProcessorTypeFromValue(v string) (*ObservabilityPipelineQuotaProcessorType, error) {
	ev := ObservabilityPipelineQuotaProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineQuotaProcessorType: valid values are %v", v, allowedObservabilityPipelineQuotaProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineQuotaProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineQuotaProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineQuotaProcessorType value.
func (v ObservabilityPipelineQuotaProcessorType) Ptr() *ObservabilityPipelineQuotaProcessorType {
	return &v
}
