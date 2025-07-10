// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineKafkaSourceType The source type. The value should always be `kafka`.
type ObservabilityPipelineKafkaSourceType string

// List of ObservabilityPipelineKafkaSourceType.
const (
	OBSERVABILITYPIPELINEKAFKASOURCETYPE_KAFKA ObservabilityPipelineKafkaSourceType = "kafka"
)

var allowedObservabilityPipelineKafkaSourceTypeEnumValues = []ObservabilityPipelineKafkaSourceType{
	OBSERVABILITYPIPELINEKAFKASOURCETYPE_KAFKA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineKafkaSourceType) GetAllowedValues() []ObservabilityPipelineKafkaSourceType {
	return allowedObservabilityPipelineKafkaSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineKafkaSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineKafkaSourceType(value)
	return nil
}

// NewObservabilityPipelineKafkaSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineKafkaSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineKafkaSourceTypeFromValue(v string) (*ObservabilityPipelineKafkaSourceType, error) {
	ev := ObservabilityPipelineKafkaSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineKafkaSourceType: valid values are %v", v, allowedObservabilityPipelineKafkaSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineKafkaSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineKafkaSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineKafkaSourceType value.
func (v ObservabilityPipelineKafkaSourceType) Ptr() *ObservabilityPipelineKafkaSourceType {
	return &v
}
