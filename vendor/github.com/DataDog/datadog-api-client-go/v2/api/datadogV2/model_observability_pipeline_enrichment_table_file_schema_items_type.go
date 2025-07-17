// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableFileSchemaItemsType Declares allowed data types for enrichment table columns.
type ObservabilityPipelineEnrichmentTableFileSchemaItemsType string

// List of ObservabilityPipelineEnrichmentTableFileSchemaItemsType.
const (
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_STRING    ObservabilityPipelineEnrichmentTableFileSchemaItemsType = "string"
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_BOOLEAN   ObservabilityPipelineEnrichmentTableFileSchemaItemsType = "boolean"
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_INTEGER   ObservabilityPipelineEnrichmentTableFileSchemaItemsType = "integer"
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_FLOAT     ObservabilityPipelineEnrichmentTableFileSchemaItemsType = "float"
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_DATE      ObservabilityPipelineEnrichmentTableFileSchemaItemsType = "date"
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_TIMESTAMP ObservabilityPipelineEnrichmentTableFileSchemaItemsType = "timestamp"
)

var allowedObservabilityPipelineEnrichmentTableFileSchemaItemsTypeEnumValues = []ObservabilityPipelineEnrichmentTableFileSchemaItemsType{
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_STRING,
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_BOOLEAN,
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_INTEGER,
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_FLOAT,
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_DATE,
	OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_TIMESTAMP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineEnrichmentTableFileSchemaItemsType) GetAllowedValues() []ObservabilityPipelineEnrichmentTableFileSchemaItemsType {
	return allowedObservabilityPipelineEnrichmentTableFileSchemaItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineEnrichmentTableFileSchemaItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineEnrichmentTableFileSchemaItemsType(value)
	return nil
}

// NewObservabilityPipelineEnrichmentTableFileSchemaItemsTypeFromValue returns a pointer to a valid ObservabilityPipelineEnrichmentTableFileSchemaItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineEnrichmentTableFileSchemaItemsTypeFromValue(v string) (*ObservabilityPipelineEnrichmentTableFileSchemaItemsType, error) {
	ev := ObservabilityPipelineEnrichmentTableFileSchemaItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineEnrichmentTableFileSchemaItemsType: valid values are %v", v, allowedObservabilityPipelineEnrichmentTableFileSchemaItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineEnrichmentTableFileSchemaItemsType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineEnrichmentTableFileSchemaItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineEnrichmentTableFileSchemaItemsType value.
func (v ObservabilityPipelineEnrichmentTableFileSchemaItemsType) Ptr() *ObservabilityPipelineEnrichmentTableFileSchemaItemsType {
	return &v
}
