// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSentinelOneDestinationRegion The SentinelOne region to send logs to.
type ObservabilityPipelineSentinelOneDestinationRegion string

// List of ObservabilityPipelineSentinelOneDestinationRegion.
const (
	OBSERVABILITYPIPELINESENTINELONEDESTINATIONREGION_US          ObservabilityPipelineSentinelOneDestinationRegion = "us"
	OBSERVABILITYPIPELINESENTINELONEDESTINATIONREGION_EU          ObservabilityPipelineSentinelOneDestinationRegion = "eu"
	OBSERVABILITYPIPELINESENTINELONEDESTINATIONREGION_CA          ObservabilityPipelineSentinelOneDestinationRegion = "ca"
	OBSERVABILITYPIPELINESENTINELONEDESTINATIONREGION_DATA_SET_US ObservabilityPipelineSentinelOneDestinationRegion = "data_set_us"
)

var allowedObservabilityPipelineSentinelOneDestinationRegionEnumValues = []ObservabilityPipelineSentinelOneDestinationRegion{
	OBSERVABILITYPIPELINESENTINELONEDESTINATIONREGION_US,
	OBSERVABILITYPIPELINESENTINELONEDESTINATIONREGION_EU,
	OBSERVABILITYPIPELINESENTINELONEDESTINATIONREGION_CA,
	OBSERVABILITYPIPELINESENTINELONEDESTINATIONREGION_DATA_SET_US,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSentinelOneDestinationRegion) GetAllowedValues() []ObservabilityPipelineSentinelOneDestinationRegion {
	return allowedObservabilityPipelineSentinelOneDestinationRegionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSentinelOneDestinationRegion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSentinelOneDestinationRegion(value)
	return nil
}

// NewObservabilityPipelineSentinelOneDestinationRegionFromValue returns a pointer to a valid ObservabilityPipelineSentinelOneDestinationRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSentinelOneDestinationRegionFromValue(v string) (*ObservabilityPipelineSentinelOneDestinationRegion, error) {
	ev := ObservabilityPipelineSentinelOneDestinationRegion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSentinelOneDestinationRegion: valid values are %v", v, allowedObservabilityPipelineSentinelOneDestinationRegionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSentinelOneDestinationRegion) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSentinelOneDestinationRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSentinelOneDestinationRegion value.
func (v ObservabilityPipelineSentinelOneDestinationRegion) Ptr() *ObservabilityPipelineSentinelOneDestinationRegion {
	return &v
}
