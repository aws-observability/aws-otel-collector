// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineElasticsearchDestinationType The destination type. The value should always be `elasticsearch`.
type ObservabilityPipelineElasticsearchDestinationType string

// List of ObservabilityPipelineElasticsearchDestinationType.
const (
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONTYPE_ELASTICSEARCH ObservabilityPipelineElasticsearchDestinationType = "elasticsearch"
)

var allowedObservabilityPipelineElasticsearchDestinationTypeEnumValues = []ObservabilityPipelineElasticsearchDestinationType{
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONTYPE_ELASTICSEARCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineElasticsearchDestinationType) GetAllowedValues() []ObservabilityPipelineElasticsearchDestinationType {
	return allowedObservabilityPipelineElasticsearchDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineElasticsearchDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineElasticsearchDestinationType(value)
	return nil
}

// NewObservabilityPipelineElasticsearchDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineElasticsearchDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineElasticsearchDestinationTypeFromValue(v string) (*ObservabilityPipelineElasticsearchDestinationType, error) {
	ev := ObservabilityPipelineElasticsearchDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineElasticsearchDestinationType: valid values are %v", v, allowedObservabilityPipelineElasticsearchDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineElasticsearchDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineElasticsearchDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineElasticsearchDestinationType value.
func (v ObservabilityPipelineElasticsearchDestinationType) Ptr() *ObservabilityPipelineElasticsearchDestinationType {
	return &v
}
