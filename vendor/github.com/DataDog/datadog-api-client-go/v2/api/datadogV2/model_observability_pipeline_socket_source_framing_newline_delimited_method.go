// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod Byte frames which are delimited by a newline character.
type ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod string

// List of ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod.
const (
	OBSERVABILITYPIPELINESOCKETSOURCEFRAMINGNEWLINEDELIMITEDMETHOD_NEWLINE_DELIMITED ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod = "newline_delimited"
)

var allowedObservabilityPipelineSocketSourceFramingNewlineDelimitedMethodEnumValues = []ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod{
	OBSERVABILITYPIPELINESOCKETSOURCEFRAMINGNEWLINEDELIMITEDMETHOD_NEWLINE_DELIMITED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod) GetAllowedValues() []ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod {
	return allowedObservabilityPipelineSocketSourceFramingNewlineDelimitedMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod(value)
	return nil
}

// NewObservabilityPipelineSocketSourceFramingNewlineDelimitedMethodFromValue returns a pointer to a valid ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSocketSourceFramingNewlineDelimitedMethodFromValue(v string) (*ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod, error) {
	ev := ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod: valid values are %v", v, allowedObservabilityPipelineSocketSourceFramingNewlineDelimitedMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSocketSourceFramingNewlineDelimitedMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod value.
func (v ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod) Ptr() *ObservabilityPipelineSocketSourceFramingNewlineDelimitedMethod {
	return &v
}
