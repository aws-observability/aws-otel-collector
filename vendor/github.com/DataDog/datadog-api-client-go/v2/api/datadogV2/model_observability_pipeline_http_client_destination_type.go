// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpClientDestinationType The destination type. The value should always be `http_client`.
type ObservabilityPipelineHttpClientDestinationType string

// List of ObservabilityPipelineHttpClientDestinationType.
const (
	OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONTYPE_HTTP_CLIENT ObservabilityPipelineHttpClientDestinationType = "http_client"
)

var allowedObservabilityPipelineHttpClientDestinationTypeEnumValues = []ObservabilityPipelineHttpClientDestinationType{
	OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONTYPE_HTTP_CLIENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineHttpClientDestinationType) GetAllowedValues() []ObservabilityPipelineHttpClientDestinationType {
	return allowedObservabilityPipelineHttpClientDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineHttpClientDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineHttpClientDestinationType(value)
	return nil
}

// NewObservabilityPipelineHttpClientDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineHttpClientDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineHttpClientDestinationTypeFromValue(v string) (*ObservabilityPipelineHttpClientDestinationType, error) {
	ev := ObservabilityPipelineHttpClientDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineHttpClientDestinationType: valid values are %v", v, allowedObservabilityPipelineHttpClientDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineHttpClientDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineHttpClientDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineHttpClientDestinationType value.
func (v ObservabilityPipelineHttpClientDestinationType) Ptr() *ObservabilityPipelineHttpClientDestinationType {
	return &v
}
