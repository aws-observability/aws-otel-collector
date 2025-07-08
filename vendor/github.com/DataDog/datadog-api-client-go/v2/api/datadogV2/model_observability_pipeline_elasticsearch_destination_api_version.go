// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineElasticsearchDestinationApiVersion The Elasticsearch API version to use. Set to `auto` to auto-detect.
type ObservabilityPipelineElasticsearchDestinationApiVersion string

// List of ObservabilityPipelineElasticsearchDestinationApiVersion.
const (
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONAPIVERSION_AUTO ObservabilityPipelineElasticsearchDestinationApiVersion = "auto"
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONAPIVERSION_V6   ObservabilityPipelineElasticsearchDestinationApiVersion = "v6"
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONAPIVERSION_V7   ObservabilityPipelineElasticsearchDestinationApiVersion = "v7"
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONAPIVERSION_V8   ObservabilityPipelineElasticsearchDestinationApiVersion = "v8"
)

var allowedObservabilityPipelineElasticsearchDestinationApiVersionEnumValues = []ObservabilityPipelineElasticsearchDestinationApiVersion{
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONAPIVERSION_AUTO,
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONAPIVERSION_V6,
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONAPIVERSION_V7,
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONAPIVERSION_V8,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineElasticsearchDestinationApiVersion) GetAllowedValues() []ObservabilityPipelineElasticsearchDestinationApiVersion {
	return allowedObservabilityPipelineElasticsearchDestinationApiVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineElasticsearchDestinationApiVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineElasticsearchDestinationApiVersion(value)
	return nil
}

// NewObservabilityPipelineElasticsearchDestinationApiVersionFromValue returns a pointer to a valid ObservabilityPipelineElasticsearchDestinationApiVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineElasticsearchDestinationApiVersionFromValue(v string) (*ObservabilityPipelineElasticsearchDestinationApiVersion, error) {
	ev := ObservabilityPipelineElasticsearchDestinationApiVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineElasticsearchDestinationApiVersion: valid values are %v", v, allowedObservabilityPipelineElasticsearchDestinationApiVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineElasticsearchDestinationApiVersion) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineElasticsearchDestinationApiVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineElasticsearchDestinationApiVersion value.
func (v ObservabilityPipelineElasticsearchDestinationApiVersion) Ptr() *ObservabilityPipelineElasticsearchDestinationApiVersion {
	return &v
}
