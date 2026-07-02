// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabRunArtifactsType The JSON:API type for a run artifacts resource.
type ModelLabRunArtifactsType string

// List of ModelLabRunArtifactsType.
const (
	MODELLABRUNARTIFACTSTYPE_ARTIFACTS ModelLabRunArtifactsType = "artifacts"
)

var allowedModelLabRunArtifactsTypeEnumValues = []ModelLabRunArtifactsType{
	MODELLABRUNARTIFACTSTYPE_ARTIFACTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ModelLabRunArtifactsType) GetAllowedValues() []ModelLabRunArtifactsType {
	return allowedModelLabRunArtifactsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ModelLabRunArtifactsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ModelLabRunArtifactsType(value)
	return nil
}

// NewModelLabRunArtifactsTypeFromValue returns a pointer to a valid ModelLabRunArtifactsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewModelLabRunArtifactsTypeFromValue(v string) (*ModelLabRunArtifactsType, error) {
	ev := ModelLabRunArtifactsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ModelLabRunArtifactsType: valid values are %v", v, allowedModelLabRunArtifactsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ModelLabRunArtifactsType) IsValid() bool {
	for _, existing := range allowedModelLabRunArtifactsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModelLabRunArtifactsType value.
func (v ModelLabRunArtifactsType) Ptr() *ModelLabRunArtifactsType {
	return &v
}
