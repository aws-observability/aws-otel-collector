// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketSourceFramingBytesMethod Byte frames are passed through as-is according to the underlying I/O boundaries (for example, split between messages or stream segments).
type ObservabilityPipelineSocketSourceFramingBytesMethod string

// List of ObservabilityPipelineSocketSourceFramingBytesMethod.
const (
	OBSERVABILITYPIPELINESOCKETSOURCEFRAMINGBYTESMETHOD_BYTES ObservabilityPipelineSocketSourceFramingBytesMethod = "bytes"
)

var allowedObservabilityPipelineSocketSourceFramingBytesMethodEnumValues = []ObservabilityPipelineSocketSourceFramingBytesMethod{
	OBSERVABILITYPIPELINESOCKETSOURCEFRAMINGBYTESMETHOD_BYTES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSocketSourceFramingBytesMethod) GetAllowedValues() []ObservabilityPipelineSocketSourceFramingBytesMethod {
	return allowedObservabilityPipelineSocketSourceFramingBytesMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSocketSourceFramingBytesMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSocketSourceFramingBytesMethod(value)
	return nil
}

// NewObservabilityPipelineSocketSourceFramingBytesMethodFromValue returns a pointer to a valid ObservabilityPipelineSocketSourceFramingBytesMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSocketSourceFramingBytesMethodFromValue(v string) (*ObservabilityPipelineSocketSourceFramingBytesMethod, error) {
	ev := ObservabilityPipelineSocketSourceFramingBytesMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSocketSourceFramingBytesMethod: valid values are %v", v, allowedObservabilityPipelineSocketSourceFramingBytesMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSocketSourceFramingBytesMethod) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSocketSourceFramingBytesMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSocketSourceFramingBytesMethod value.
func (v ObservabilityPipelineSocketSourceFramingBytesMethod) Ptr() *ObservabilityPipelineSocketSourceFramingBytesMethod {
	return &v
}
