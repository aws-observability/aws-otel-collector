// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableFileEncodingType Specifies the encoding format (e.g., CSV) used for enrichment tables.
type ObservabilityPipelineEnrichmentTableFileEncodingType string

// List of ObservabilityPipelineEnrichmentTableFileEncodingType.
const (
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILEENCODINGTYPE_CSV ObservabilityPipelineEnrichmentTableFileEncodingType = "csv"
)

var allowedObservabilityPipelineEnrichmentTableFileEncodingTypeEnumValues = []ObservabilityPipelineEnrichmentTableFileEncodingType{
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILEENCODINGTYPE_CSV,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineEnrichmentTableFileEncodingType) GetAllowedValues() []ObservabilityPipelineEnrichmentTableFileEncodingType {
	return allowedObservabilityPipelineEnrichmentTableFileEncodingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineEnrichmentTableFileEncodingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineEnrichmentTableFileEncodingType(value)
	return nil
}

// NewObservabilityPipelineEnrichmentTableFileEncodingTypeFromValue returns a pointer to a valid ObservabilityPipelineEnrichmentTableFileEncodingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineEnrichmentTableFileEncodingTypeFromValue(v string) (*ObservabilityPipelineEnrichmentTableFileEncodingType, error) {
	ev := ObservabilityPipelineEnrichmentTableFileEncodingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineEnrichmentTableFileEncodingType: valid values are %v", v, allowedObservabilityPipelineEnrichmentTableFileEncodingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineEnrichmentTableFileEncodingType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineEnrichmentTableFileEncodingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineEnrichmentTableFileEncodingType value.
func (v ObservabilityPipelineEnrichmentTableFileEncodingType) Ptr() *ObservabilityPipelineEnrichmentTableFileEncodingType {
	return &v
}
