// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabProjectArtifactsType The JSON:API type for a project artifacts resource.
type ModelLabProjectArtifactsType string

// List of ModelLabProjectArtifactsType.
const (
	MODELLABPROJECTARTIFACTSTYPE_PROJECT_FILES ModelLabProjectArtifactsType = "project_files"
)

var allowedModelLabProjectArtifactsTypeEnumValues = []ModelLabProjectArtifactsType{
	MODELLABPROJECTARTIFACTSTYPE_PROJECT_FILES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ModelLabProjectArtifactsType) GetAllowedValues() []ModelLabProjectArtifactsType {
	return allowedModelLabProjectArtifactsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ModelLabProjectArtifactsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ModelLabProjectArtifactsType(value)
	return nil
}

// NewModelLabProjectArtifactsTypeFromValue returns a pointer to a valid ModelLabProjectArtifactsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewModelLabProjectArtifactsTypeFromValue(v string) (*ModelLabProjectArtifactsType, error) {
	ev := ModelLabProjectArtifactsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ModelLabProjectArtifactsType: valid values are %v", v, allowedModelLabProjectArtifactsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ModelLabProjectArtifactsType) IsValid() bool {
	for _, existing := range allowedModelLabProjectArtifactsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModelLabProjectArtifactsType value.
func (v ModelLabProjectArtifactsType) Ptr() *ModelLabProjectArtifactsType {
	return &v
}
