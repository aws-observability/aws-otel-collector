// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpServerSourceType The source type. The value should always be `http_server`.
type ObservabilityPipelineHttpServerSourceType string

// List of ObservabilityPipelineHttpServerSourceType.
const (
	OBSERVABILITYPIPELINEHTTPSERVERSOURCETYPE_HTTP_SERVER ObservabilityPipelineHttpServerSourceType = "http_server"
)

var allowedObservabilityPipelineHttpServerSourceTypeEnumValues = []ObservabilityPipelineHttpServerSourceType{
	OBSERVABILITYPIPELINEHTTPSERVERSOURCETYPE_HTTP_SERVER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineHttpServerSourceType) GetAllowedValues() []ObservabilityPipelineHttpServerSourceType {
	return allowedObservabilityPipelineHttpServerSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineHttpServerSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineHttpServerSourceType(value)
	return nil
}

// NewObservabilityPipelineHttpServerSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineHttpServerSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineHttpServerSourceTypeFromValue(v string) (*ObservabilityPipelineHttpServerSourceType, error) {
	ev := ObservabilityPipelineHttpServerSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineHttpServerSourceType: valid values are %v", v, allowedObservabilityPipelineHttpServerSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineHttpServerSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineHttpServerSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineHttpServerSourceType value.
func (v ObservabilityPipelineHttpServerSourceType) Ptr() *ObservabilityPipelineHttpServerSourceType {
	return &v
}
