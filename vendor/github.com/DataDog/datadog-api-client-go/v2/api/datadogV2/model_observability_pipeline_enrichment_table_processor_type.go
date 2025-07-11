// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableProcessorType The processor type. The value should always be `enrichment_table`.
type ObservabilityPipelineEnrichmentTableProcessorType string

// List of ObservabilityPipelineEnrichmentTableProcessorType.
const (
	OBSERVABILITYPIPELINEENRICHMENTTABLEPROCESSORTYPE_ENRICHMENT_TABLE ObservabilityPipelineEnrichmentTableProcessorType = "enrichment_table"
)

var allowedObservabilityPipelineEnrichmentTableProcessorTypeEnumValues = []ObservabilityPipelineEnrichmentTableProcessorType{
	OBSERVABILITYPIPELINEENRICHMENTTABLEPROCESSORTYPE_ENRICHMENT_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineEnrichmentTableProcessorType) GetAllowedValues() []ObservabilityPipelineEnrichmentTableProcessorType {
	return allowedObservabilityPipelineEnrichmentTableProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineEnrichmentTableProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineEnrichmentTableProcessorType(value)
	return nil
}

// NewObservabilityPipelineEnrichmentTableProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineEnrichmentTableProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineEnrichmentTableProcessorTypeFromValue(v string) (*ObservabilityPipelineEnrichmentTableProcessorType, error) {
	ev := ObservabilityPipelineEnrichmentTableProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineEnrichmentTableProcessorType: valid values are %v", v, allowedObservabilityPipelineEnrichmentTableProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineEnrichmentTableProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineEnrichmentTableProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineEnrichmentTableProcessorType value.
func (v ObservabilityPipelineEnrichmentTableProcessorType) Ptr() *ObservabilityPipelineEnrichmentTableProcessorType {
	return &v
}
