// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpClientSourceType The source type. The value should always be `http_client`.
type ObservabilityPipelineHttpClientSourceType string

// List of ObservabilityPipelineHttpClientSourceType.
const (
	OBSERVABILITYPIPELINEHTTPCLIENTSOURCETYPE_HTTP_CLIENT ObservabilityPipelineHttpClientSourceType = "http_client"
)

var allowedObservabilityPipelineHttpClientSourceTypeEnumValues = []ObservabilityPipelineHttpClientSourceType{
	OBSERVABILITYPIPELINEHTTPCLIENTSOURCETYPE_HTTP_CLIENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineHttpClientSourceType) GetAllowedValues() []ObservabilityPipelineHttpClientSourceType {
	return allowedObservabilityPipelineHttpClientSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineHttpClientSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineHttpClientSourceType(value)
	return nil
}

// NewObservabilityPipelineHttpClientSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineHttpClientSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineHttpClientSourceTypeFromValue(v string) (*ObservabilityPipelineHttpClientSourceType, error) {
	ev := ObservabilityPipelineHttpClientSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineHttpClientSourceType: valid values are %v", v, allowedObservabilityPipelineHttpClientSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineHttpClientSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineHttpClientSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineHttpClientSourceType value.
func (v ObservabilityPipelineHttpClientSourceType) Ptr() *ObservabilityPipelineHttpClientSourceType {
	return &v
}
