// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCloudPremDestinationType The destination type. The value should always be `cloud_prem`.
type ObservabilityPipelineCloudPremDestinationType string

// List of ObservabilityPipelineCloudPremDestinationType.
const (
	OBSERVABILITYPIPELINECLOUDPREMDESTINATIONTYPE_CLOUD_PREM ObservabilityPipelineCloudPremDestinationType = "cloud_prem"
)

var allowedObservabilityPipelineCloudPremDestinationTypeEnumValues = []ObservabilityPipelineCloudPremDestinationType{
	OBSERVABILITYPIPELINECLOUDPREMDESTINATIONTYPE_CLOUD_PREM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineCloudPremDestinationType) GetAllowedValues() []ObservabilityPipelineCloudPremDestinationType {
	return allowedObservabilityPipelineCloudPremDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineCloudPremDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineCloudPremDestinationType(value)
	return nil
}

// NewObservabilityPipelineCloudPremDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineCloudPremDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineCloudPremDestinationTypeFromValue(v string) (*ObservabilityPipelineCloudPremDestinationType, error) {
	ev := ObservabilityPipelineCloudPremDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineCloudPremDestinationType: valid values are %v", v, allowedObservabilityPipelineCloudPremDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineCloudPremDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineCloudPremDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineCloudPremDestinationType value.
func (v ObservabilityPipelineCloudPremDestinationType) Ptr() *ObservabilityPipelineCloudPremDestinationType {
	return &v
}
