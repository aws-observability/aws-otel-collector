// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketDestinationFramingBytesMethod The definition of `ObservabilityPipelineSocketDestinationFramingBytesMethod` object.
type ObservabilityPipelineSocketDestinationFramingBytesMethod string

// List of ObservabilityPipelineSocketDestinationFramingBytesMethod.
const (
	OBSERVABILITYPIPELINESOCKETDESTINATIONFRAMINGBYTESMETHOD_BYTES ObservabilityPipelineSocketDestinationFramingBytesMethod = "bytes"
)

var allowedObservabilityPipelineSocketDestinationFramingBytesMethodEnumValues = []ObservabilityPipelineSocketDestinationFramingBytesMethod{
	OBSERVABILITYPIPELINESOCKETDESTINATIONFRAMINGBYTESMETHOD_BYTES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSocketDestinationFramingBytesMethod) GetAllowedValues() []ObservabilityPipelineSocketDestinationFramingBytesMethod {
	return allowedObservabilityPipelineSocketDestinationFramingBytesMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSocketDestinationFramingBytesMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSocketDestinationFramingBytesMethod(value)
	return nil
}

// NewObservabilityPipelineSocketDestinationFramingBytesMethodFromValue returns a pointer to a valid ObservabilityPipelineSocketDestinationFramingBytesMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSocketDestinationFramingBytesMethodFromValue(v string) (*ObservabilityPipelineSocketDestinationFramingBytesMethod, error) {
	ev := ObservabilityPipelineSocketDestinationFramingBytesMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSocketDestinationFramingBytesMethod: valid values are %v", v, allowedObservabilityPipelineSocketDestinationFramingBytesMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSocketDestinationFramingBytesMethod) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSocketDestinationFramingBytesMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSocketDestinationFramingBytesMethod value.
func (v ObservabilityPipelineSocketDestinationFramingBytesMethod) Ptr() *ObservabilityPipelineSocketDestinationFramingBytesMethod {
	return &v
}
