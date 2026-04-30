// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineKafkaDestinationEncoding Encoding format for log events.
type ObservabilityPipelineKafkaDestinationEncoding string

// List of ObservabilityPipelineKafkaDestinationEncoding.
const (
	OBSERVABILITYPIPELINEKAFKADESTINATIONENCODING_JSON        ObservabilityPipelineKafkaDestinationEncoding = "json"
	OBSERVABILITYPIPELINEKAFKADESTINATIONENCODING_RAW_MESSAGE ObservabilityPipelineKafkaDestinationEncoding = "raw_message"
)

var allowedObservabilityPipelineKafkaDestinationEncodingEnumValues = []ObservabilityPipelineKafkaDestinationEncoding{
	OBSERVABILITYPIPELINEKAFKADESTINATIONENCODING_JSON,
	OBSERVABILITYPIPELINEKAFKADESTINATIONENCODING_RAW_MESSAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineKafkaDestinationEncoding) GetAllowedValues() []ObservabilityPipelineKafkaDestinationEncoding {
	return allowedObservabilityPipelineKafkaDestinationEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineKafkaDestinationEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineKafkaDestinationEncoding(value)
	return nil
}

// NewObservabilityPipelineKafkaDestinationEncodingFromValue returns a pointer to a valid ObservabilityPipelineKafkaDestinationEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineKafkaDestinationEncodingFromValue(v string) (*ObservabilityPipelineKafkaDestinationEncoding, error) {
	ev := ObservabilityPipelineKafkaDestinationEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineKafkaDestinationEncoding: valid values are %v", v, allowedObservabilityPipelineKafkaDestinationEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineKafkaDestinationEncoding) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineKafkaDestinationEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineKafkaDestinationEncoding value.
func (v ObservabilityPipelineKafkaDestinationEncoding) Ptr() *ObservabilityPipelineKafkaDestinationEncoding {
	return &v
}
