// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOpenSearchDestinationType The destination type. The value should always be `opensearch`.
type ObservabilityPipelineOpenSearchDestinationType string

// List of ObservabilityPipelineOpenSearchDestinationType.
const (
	OBSERVABILITYPIPELINEOPENSEARCHDESTINATIONTYPE_OPENSEARCH ObservabilityPipelineOpenSearchDestinationType = "opensearch"
)

var allowedObservabilityPipelineOpenSearchDestinationTypeEnumValues = []ObservabilityPipelineOpenSearchDestinationType{
	OBSERVABILITYPIPELINEOPENSEARCHDESTINATIONTYPE_OPENSEARCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineOpenSearchDestinationType) GetAllowedValues() []ObservabilityPipelineOpenSearchDestinationType {
	return allowedObservabilityPipelineOpenSearchDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineOpenSearchDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineOpenSearchDestinationType(value)
	return nil
}

// NewObservabilityPipelineOpenSearchDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineOpenSearchDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineOpenSearchDestinationTypeFromValue(v string) (*ObservabilityPipelineOpenSearchDestinationType, error) {
	ev := ObservabilityPipelineOpenSearchDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineOpenSearchDestinationType: valid values are %v", v, allowedObservabilityPipelineOpenSearchDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineOpenSearchDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineOpenSearchDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineOpenSearchDestinationType value.
func (v ObservabilityPipelineOpenSearchDestinationType) Ptr() *ObservabilityPipelineOpenSearchDestinationType {
	return &v
}
