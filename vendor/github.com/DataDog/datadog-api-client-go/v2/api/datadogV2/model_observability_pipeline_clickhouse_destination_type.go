// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineClickhouseDestinationType The destination type. The value must be `clickhouse`.
type ObservabilityPipelineClickhouseDestinationType string

// List of ObservabilityPipelineClickhouseDestinationType.
const (
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONTYPE_CLICKHOUSE ObservabilityPipelineClickhouseDestinationType = "clickhouse"
)

var allowedObservabilityPipelineClickhouseDestinationTypeEnumValues = []ObservabilityPipelineClickhouseDestinationType{
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONTYPE_CLICKHOUSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineClickhouseDestinationType) GetAllowedValues() []ObservabilityPipelineClickhouseDestinationType {
	return allowedObservabilityPipelineClickhouseDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineClickhouseDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineClickhouseDestinationType(value)
	return nil
}

// NewObservabilityPipelineClickhouseDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineClickhouseDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineClickhouseDestinationTypeFromValue(v string) (*ObservabilityPipelineClickhouseDestinationType, error) {
	ev := ObservabilityPipelineClickhouseDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineClickhouseDestinationType: valid values are %v", v, allowedObservabilityPipelineClickhouseDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineClickhouseDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineClickhouseDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineClickhouseDestinationType value.
func (v ObservabilityPipelineClickhouseDestinationType) Ptr() *ObservabilityPipelineClickhouseDestinationType {
	return &v
}
