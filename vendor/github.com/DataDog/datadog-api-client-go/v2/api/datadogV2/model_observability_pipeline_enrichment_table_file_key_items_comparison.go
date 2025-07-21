// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableFileKeyItemsComparison Defines how to compare key fields for enrichment table lookups.
type ObservabilityPipelineEnrichmentTableFileKeyItemsComparison string

// List of ObservabilityPipelineEnrichmentTableFileKeyItemsComparison.
const (
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILEKEYITEMSCOMPARISON_EQUALS ObservabilityPipelineEnrichmentTableFileKeyItemsComparison = "equals"
)

var allowedObservabilityPipelineEnrichmentTableFileKeyItemsComparisonEnumValues = []ObservabilityPipelineEnrichmentTableFileKeyItemsComparison{
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILEKEYITEMSCOMPARISON_EQUALS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineEnrichmentTableFileKeyItemsComparison) GetAllowedValues() []ObservabilityPipelineEnrichmentTableFileKeyItemsComparison {
	return allowedObservabilityPipelineEnrichmentTableFileKeyItemsComparisonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineEnrichmentTableFileKeyItemsComparison) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineEnrichmentTableFileKeyItemsComparison(value)
	return nil
}

// NewObservabilityPipelineEnrichmentTableFileKeyItemsComparisonFromValue returns a pointer to a valid ObservabilityPipelineEnrichmentTableFileKeyItemsComparison
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineEnrichmentTableFileKeyItemsComparisonFromValue(v string) (*ObservabilityPipelineEnrichmentTableFileKeyItemsComparison, error) {
	ev := ObservabilityPipelineEnrichmentTableFileKeyItemsComparison(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineEnrichmentTableFileKeyItemsComparison: valid values are %v", v, allowedObservabilityPipelineEnrichmentTableFileKeyItemsComparisonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineEnrichmentTableFileKeyItemsComparison) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineEnrichmentTableFileKeyItemsComparisonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineEnrichmentTableFileKeyItemsComparison value.
func (v ObservabilityPipelineEnrichmentTableFileKeyItemsComparison) Ptr() *ObservabilityPipelineEnrichmentTableFileKeyItemsComparison {
	return &v
}
