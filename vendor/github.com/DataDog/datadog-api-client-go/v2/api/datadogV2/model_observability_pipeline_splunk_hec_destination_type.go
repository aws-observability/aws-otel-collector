// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplunkHecDestinationType The destination type. Always `splunk_hec`.
type ObservabilityPipelineSplunkHecDestinationType string

// List of ObservabilityPipelineSplunkHecDestinationType.
const (
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONTYPE_SPLUNK_HEC ObservabilityPipelineSplunkHecDestinationType = "splunk_hec"
)

var allowedObservabilityPipelineSplunkHecDestinationTypeEnumValues = []ObservabilityPipelineSplunkHecDestinationType{
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONTYPE_SPLUNK_HEC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSplunkHecDestinationType) GetAllowedValues() []ObservabilityPipelineSplunkHecDestinationType {
	return allowedObservabilityPipelineSplunkHecDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSplunkHecDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSplunkHecDestinationType(value)
	return nil
}

// NewObservabilityPipelineSplunkHecDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineSplunkHecDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSplunkHecDestinationTypeFromValue(v string) (*ObservabilityPipelineSplunkHecDestinationType, error) {
	ev := ObservabilityPipelineSplunkHecDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSplunkHecDestinationType: valid values are %v", v, allowedObservabilityPipelineSplunkHecDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSplunkHecDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSplunkHecDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSplunkHecDestinationType value.
func (v ObservabilityPipelineSplunkHecDestinationType) Ptr() *ObservabilityPipelineSplunkHecDestinationType {
	return &v
}
