// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAddEnvVarsProcessorType The processor type. The value should always be `add_env_vars`.
type ObservabilityPipelineAddEnvVarsProcessorType string

// List of ObservabilityPipelineAddEnvVarsProcessorType.
const (
	OBSERVABILITYPIPELINEADDENVVARSPROCESSORTYPE_ADD_ENV_VARS ObservabilityPipelineAddEnvVarsProcessorType = "add_env_vars"
)

var allowedObservabilityPipelineAddEnvVarsProcessorTypeEnumValues = []ObservabilityPipelineAddEnvVarsProcessorType{
	OBSERVABILITYPIPELINEADDENVVARSPROCESSORTYPE_ADD_ENV_VARS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAddEnvVarsProcessorType) GetAllowedValues() []ObservabilityPipelineAddEnvVarsProcessorType {
	return allowedObservabilityPipelineAddEnvVarsProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAddEnvVarsProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAddEnvVarsProcessorType(value)
	return nil
}

// NewObservabilityPipelineAddEnvVarsProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineAddEnvVarsProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAddEnvVarsProcessorTypeFromValue(v string) (*ObservabilityPipelineAddEnvVarsProcessorType, error) {
	ev := ObservabilityPipelineAddEnvVarsProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAddEnvVarsProcessorType: valid values are %v", v, allowedObservabilityPipelineAddEnvVarsProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAddEnvVarsProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAddEnvVarsProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAddEnvVarsProcessorType value.
func (v ObservabilityPipelineAddEnvVarsProcessorType) Ptr() *ObservabilityPipelineAddEnvVarsProcessorType {
	return &v
}
