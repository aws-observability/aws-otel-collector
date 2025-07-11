// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelinePipelineKafkaSourceSaslMechanism SASL mechanism used for Kafka authentication.
type ObservabilityPipelinePipelineKafkaSourceSaslMechanism string

// List of ObservabilityPipelinePipelineKafkaSourceSaslMechanism.
const (
	OBSERVABILITYPIPELINEPIPELINEKAFKASOURCESASLMECHANISM_PLAIN               ObservabilityPipelinePipelineKafkaSourceSaslMechanism = "PLAIN"
	OBSERVABILITYPIPELINEPIPELINEKAFKASOURCESASLMECHANISM_SCRAMNOT_SHANOT_256 ObservabilityPipelinePipelineKafkaSourceSaslMechanism = "SCRAM-SHA-256"
	OBSERVABILITYPIPELINEPIPELINEKAFKASOURCESASLMECHANISM_SCRAMNOT_SHANOT_512 ObservabilityPipelinePipelineKafkaSourceSaslMechanism = "SCRAM-SHA-512"
)

var allowedObservabilityPipelinePipelineKafkaSourceSaslMechanismEnumValues = []ObservabilityPipelinePipelineKafkaSourceSaslMechanism{
	OBSERVABILITYPIPELINEPIPELINEKAFKASOURCESASLMECHANISM_PLAIN,
	OBSERVABILITYPIPELINEPIPELINEKAFKASOURCESASLMECHANISM_SCRAMNOT_SHANOT_256,
	OBSERVABILITYPIPELINEPIPELINEKAFKASOURCESASLMECHANISM_SCRAMNOT_SHANOT_512,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelinePipelineKafkaSourceSaslMechanism) GetAllowedValues() []ObservabilityPipelinePipelineKafkaSourceSaslMechanism {
	return allowedObservabilityPipelinePipelineKafkaSourceSaslMechanismEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelinePipelineKafkaSourceSaslMechanism) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelinePipelineKafkaSourceSaslMechanism(value)
	return nil
}

// NewObservabilityPipelinePipelineKafkaSourceSaslMechanismFromValue returns a pointer to a valid ObservabilityPipelinePipelineKafkaSourceSaslMechanism
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelinePipelineKafkaSourceSaslMechanismFromValue(v string) (*ObservabilityPipelinePipelineKafkaSourceSaslMechanism, error) {
	ev := ObservabilityPipelinePipelineKafkaSourceSaslMechanism(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelinePipelineKafkaSourceSaslMechanism: valid values are %v", v, allowedObservabilityPipelinePipelineKafkaSourceSaslMechanismEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelinePipelineKafkaSourceSaslMechanism) IsValid() bool {
	for _, existing := range allowedObservabilityPipelinePipelineKafkaSourceSaslMechanismEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelinePipelineKafkaSourceSaslMechanism value.
func (v ObservabilityPipelinePipelineKafkaSourceSaslMechanism) Ptr() *ObservabilityPipelinePipelineKafkaSourceSaslMechanism {
	return &v
}
