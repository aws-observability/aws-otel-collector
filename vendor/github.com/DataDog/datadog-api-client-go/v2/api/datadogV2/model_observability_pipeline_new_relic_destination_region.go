// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineNewRelicDestinationRegion The New Relic region.
type ObservabilityPipelineNewRelicDestinationRegion string

// List of ObservabilityPipelineNewRelicDestinationRegion.
const (
	OBSERVABILITYPIPELINENEWRELICDESTINATIONREGION_US ObservabilityPipelineNewRelicDestinationRegion = "us"
	OBSERVABILITYPIPELINENEWRELICDESTINATIONREGION_EU ObservabilityPipelineNewRelicDestinationRegion = "eu"
)

var allowedObservabilityPipelineNewRelicDestinationRegionEnumValues = []ObservabilityPipelineNewRelicDestinationRegion{
	OBSERVABILITYPIPELINENEWRELICDESTINATIONREGION_US,
	OBSERVABILITYPIPELINENEWRELICDESTINATIONREGION_EU,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineNewRelicDestinationRegion) GetAllowedValues() []ObservabilityPipelineNewRelicDestinationRegion {
	return allowedObservabilityPipelineNewRelicDestinationRegionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineNewRelicDestinationRegion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineNewRelicDestinationRegion(value)
	return nil
}

// NewObservabilityPipelineNewRelicDestinationRegionFromValue returns a pointer to a valid ObservabilityPipelineNewRelicDestinationRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineNewRelicDestinationRegionFromValue(v string) (*ObservabilityPipelineNewRelicDestinationRegion, error) {
	ev := ObservabilityPipelineNewRelicDestinationRegion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineNewRelicDestinationRegion: valid values are %v", v, allowedObservabilityPipelineNewRelicDestinationRegionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineNewRelicDestinationRegion) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineNewRelicDestinationRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineNewRelicDestinationRegion value.
func (v ObservabilityPipelineNewRelicDestinationRegion) Ptr() *ObservabilityPipelineNewRelicDestinationRegion {
	return &v
}
