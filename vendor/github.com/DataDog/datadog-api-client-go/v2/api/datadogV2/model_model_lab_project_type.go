// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabProjectType The JSON:API type for a Model Lab project resource.
type ModelLabProjectType string

// List of ModelLabProjectType.
const (
	MODELLABPROJECTTYPE_PROJECTS ModelLabProjectType = "projects"
)

var allowedModelLabProjectTypeEnumValues = []ModelLabProjectType{
	MODELLABPROJECTTYPE_PROJECTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ModelLabProjectType) GetAllowedValues() []ModelLabProjectType {
	return allowedModelLabProjectTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ModelLabProjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ModelLabProjectType(value)
	return nil
}

// NewModelLabProjectTypeFromValue returns a pointer to a valid ModelLabProjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewModelLabProjectTypeFromValue(v string) (*ModelLabProjectType, error) {
	ev := ModelLabProjectType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ModelLabProjectType: valid values are %v", v, allowedModelLabProjectTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ModelLabProjectType) IsValid() bool {
	for _, existing := range allowedModelLabProjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModelLabProjectType value.
func (v ModelLabProjectType) Ptr() *ModelLabProjectType {
	return &v
}
