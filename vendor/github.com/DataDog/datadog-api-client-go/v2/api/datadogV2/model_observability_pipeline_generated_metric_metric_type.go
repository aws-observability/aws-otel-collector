// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGeneratedMetricMetricType Type of metric to create.
type ObservabilityPipelineGeneratedMetricMetricType string

// List of ObservabilityPipelineGeneratedMetricMetricType.
const (
	OBSERVABILITYPIPELINEGENERATEDMETRICMETRICTYPE_COUNT        ObservabilityPipelineGeneratedMetricMetricType = "count"
	OBSERVABILITYPIPELINEGENERATEDMETRICMETRICTYPE_GAUGE        ObservabilityPipelineGeneratedMetricMetricType = "gauge"
	OBSERVABILITYPIPELINEGENERATEDMETRICMETRICTYPE_DISTRIBUTION ObservabilityPipelineGeneratedMetricMetricType = "distribution"
)

var allowedObservabilityPipelineGeneratedMetricMetricTypeEnumValues = []ObservabilityPipelineGeneratedMetricMetricType{
	OBSERVABILITYPIPELINEGENERATEDMETRICMETRICTYPE_COUNT,
	OBSERVABILITYPIPELINEGENERATEDMETRICMETRICTYPE_GAUGE,
	OBSERVABILITYPIPELINEGENERATEDMETRICMETRICTYPE_DISTRIBUTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGeneratedMetricMetricType) GetAllowedValues() []ObservabilityPipelineGeneratedMetricMetricType {
	return allowedObservabilityPipelineGeneratedMetricMetricTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGeneratedMetricMetricType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGeneratedMetricMetricType(value)
	return nil
}

// NewObservabilityPipelineGeneratedMetricMetricTypeFromValue returns a pointer to a valid ObservabilityPipelineGeneratedMetricMetricType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGeneratedMetricMetricTypeFromValue(v string) (*ObservabilityPipelineGeneratedMetricMetricType, error) {
	ev := ObservabilityPipelineGeneratedMetricMetricType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGeneratedMetricMetricType: valid values are %v", v, allowedObservabilityPipelineGeneratedMetricMetricTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGeneratedMetricMetricType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGeneratedMetricMetricTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGeneratedMetricMetricType value.
func (v ObservabilityPipelineGeneratedMetricMetricType) Ptr() *ObservabilityPipelineGeneratedMetricMetricType {
	return &v
}
