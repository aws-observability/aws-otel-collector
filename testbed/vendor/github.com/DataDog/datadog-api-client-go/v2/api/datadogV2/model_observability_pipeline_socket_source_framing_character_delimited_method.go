// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod Byte frames which are delimited by a chosen character.
type ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod string

// List of ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod.
const (
	OBSERVABILITYPIPELINESOCKETSOURCEFRAMINGCHARACTERDELIMITEDMETHOD_CHARACTER_DELIMITED ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod = "character_delimited"
)

var allowedObservabilityPipelineSocketSourceFramingCharacterDelimitedMethodEnumValues = []ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod{
	OBSERVABILITYPIPELINESOCKETSOURCEFRAMINGCHARACTERDELIMITEDMETHOD_CHARACTER_DELIMITED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod) GetAllowedValues() []ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod {
	return allowedObservabilityPipelineSocketSourceFramingCharacterDelimitedMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod(value)
	return nil
}

// NewObservabilityPipelineSocketSourceFramingCharacterDelimitedMethodFromValue returns a pointer to a valid ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSocketSourceFramingCharacterDelimitedMethodFromValue(v string) (*ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod, error) {
	ev := ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod: valid values are %v", v, allowedObservabilityPipelineSocketSourceFramingCharacterDelimitedMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSocketSourceFramingCharacterDelimitedMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod value.
func (v ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod) Ptr() *ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod {
	return &v
}
