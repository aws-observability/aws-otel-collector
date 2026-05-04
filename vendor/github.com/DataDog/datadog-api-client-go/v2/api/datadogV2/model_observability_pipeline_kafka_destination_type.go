// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineKafkaDestinationType The destination type. The value should always be `kafka`.
type ObservabilityPipelineKafkaDestinationType string

// List of ObservabilityPipelineKafkaDestinationType.
const (
	OBSERVABILITYPIPELINEKAFKADESTINATIONTYPE_KAFKA ObservabilityPipelineKafkaDestinationType = "kafka"
)

var allowedObservabilityPipelineKafkaDestinationTypeEnumValues = []ObservabilityPipelineKafkaDestinationType{
	OBSERVABILITYPIPELINEKAFKADESTINATIONTYPE_KAFKA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineKafkaDestinationType) GetAllowedValues() []ObservabilityPipelineKafkaDestinationType {
	return allowedObservabilityPipelineKafkaDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineKafkaDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineKafkaDestinationType(value)
	return nil
}

// NewObservabilityPipelineKafkaDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineKafkaDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineKafkaDestinationTypeFromValue(v string) (*ObservabilityPipelineKafkaDestinationType, error) {
	ev := ObservabilityPipelineKafkaDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineKafkaDestinationType: valid values are %v", v, allowedObservabilityPipelineKafkaDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineKafkaDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineKafkaDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineKafkaDestinationType value.
func (v ObservabilityPipelineKafkaDestinationType) Ptr() *ObservabilityPipelineKafkaDestinationType {
	return &v
}
