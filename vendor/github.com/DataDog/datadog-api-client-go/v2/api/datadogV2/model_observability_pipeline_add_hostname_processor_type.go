// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAddHostnameProcessorType The processor type. The value should always be `add_hostname`.
type ObservabilityPipelineAddHostnameProcessorType string

// List of ObservabilityPipelineAddHostnameProcessorType.
const (
	OBSERVABILITYPIPELINEADDHOSTNAMEPROCESSORTYPE_ADD_HOSTNAME ObservabilityPipelineAddHostnameProcessorType = "add_hostname"
)

var allowedObservabilityPipelineAddHostnameProcessorTypeEnumValues = []ObservabilityPipelineAddHostnameProcessorType{
	OBSERVABILITYPIPELINEADDHOSTNAMEPROCESSORTYPE_ADD_HOSTNAME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAddHostnameProcessorType) GetAllowedValues() []ObservabilityPipelineAddHostnameProcessorType {
	return allowedObservabilityPipelineAddHostnameProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAddHostnameProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAddHostnameProcessorType(value)
	return nil
}

// NewObservabilityPipelineAddHostnameProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineAddHostnameProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAddHostnameProcessorTypeFromValue(v string) (*ObservabilityPipelineAddHostnameProcessorType, error) {
	ev := ObservabilityPipelineAddHostnameProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAddHostnameProcessorType: valid values are %v", v, allowedObservabilityPipelineAddHostnameProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAddHostnameProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAddHostnameProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAddHostnameProcessorType value.
func (v ObservabilityPipelineAddHostnameProcessorType) Ptr() *ObservabilityPipelineAddHostnameProcessorType {
	return &v
}
