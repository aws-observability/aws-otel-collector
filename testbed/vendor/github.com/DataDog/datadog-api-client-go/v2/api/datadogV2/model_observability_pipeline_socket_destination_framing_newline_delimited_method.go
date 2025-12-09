// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod The definition of `ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod` object.
type ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod string

// List of ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod.
const (
	OBSERVABILITYPIPELINESOCKETDESTINATIONFRAMINGNEWLINEDELIMITEDMETHOD_NEWLINE_DELIMITED ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod = "newline_delimited"
)

var allowedObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethodEnumValues = []ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod{
	OBSERVABILITYPIPELINESOCKETDESTINATIONFRAMINGNEWLINEDELIMITEDMETHOD_NEWLINE_DELIMITED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod) GetAllowedValues() []ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod {
	return allowedObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod(value)
	return nil
}

// NewObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethodFromValue returns a pointer to a valid ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethodFromValue(v string) (*ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod, error) {
	ev := ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod: valid values are %v", v, allowedObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod value.
func (v ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod) Ptr() *ObservabilityPipelineSocketDestinationFramingNewlineDelimitedMethod {
	return &v
}
