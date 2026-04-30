// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineKafkaSaslMechanism SASL mechanism used for Kafka authentication.
type ObservabilityPipelineKafkaSaslMechanism string

// List of ObservabilityPipelineKafkaSaslMechanism.
const (
	OBSERVABILITYPIPELINEKAFKASASLMECHANISM_PLAIN               ObservabilityPipelineKafkaSaslMechanism = "PLAIN"
	OBSERVABILITYPIPELINEKAFKASASLMECHANISM_SCRAMNOT_SHANOT_256 ObservabilityPipelineKafkaSaslMechanism = "SCRAM-SHA-256"
	OBSERVABILITYPIPELINEKAFKASASLMECHANISM_SCRAMNOT_SHANOT_512 ObservabilityPipelineKafkaSaslMechanism = "SCRAM-SHA-512"
)

var allowedObservabilityPipelineKafkaSaslMechanismEnumValues = []ObservabilityPipelineKafkaSaslMechanism{
	OBSERVABILITYPIPELINEKAFKASASLMECHANISM_PLAIN,
	OBSERVABILITYPIPELINEKAFKASASLMECHANISM_SCRAMNOT_SHANOT_256,
	OBSERVABILITYPIPELINEKAFKASASLMECHANISM_SCRAMNOT_SHANOT_512,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineKafkaSaslMechanism) GetAllowedValues() []ObservabilityPipelineKafkaSaslMechanism {
	return allowedObservabilityPipelineKafkaSaslMechanismEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineKafkaSaslMechanism) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineKafkaSaslMechanism(value)
	return nil
}

// NewObservabilityPipelineKafkaSaslMechanismFromValue returns a pointer to a valid ObservabilityPipelineKafkaSaslMechanism
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineKafkaSaslMechanismFromValue(v string) (*ObservabilityPipelineKafkaSaslMechanism, error) {
	ev := ObservabilityPipelineKafkaSaslMechanism(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineKafkaSaslMechanism: valid values are %v", v, allowedObservabilityPipelineKafkaSaslMechanismEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineKafkaSaslMechanism) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineKafkaSaslMechanismEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineKafkaSaslMechanism value.
func (v ObservabilityPipelineKafkaSaslMechanism) Ptr() *ObservabilityPipelineKafkaSaslMechanism {
	return &v
}
