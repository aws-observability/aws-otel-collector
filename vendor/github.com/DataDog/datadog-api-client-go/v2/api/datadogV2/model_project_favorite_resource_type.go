// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectFavoriteResourceType JSON:API resource type for project favorites.
type ProjectFavoriteResourceType string

// List of ProjectFavoriteResourceType.
const (
	PROJECTFAVORITERESOURCETYPE_PROJECT_FAVORITE ProjectFavoriteResourceType = "project_favorite"
)

var allowedProjectFavoriteResourceTypeEnumValues = []ProjectFavoriteResourceType{
	PROJECTFAVORITERESOURCETYPE_PROJECT_FAVORITE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProjectFavoriteResourceType) GetAllowedValues() []ProjectFavoriteResourceType {
	return allowedProjectFavoriteResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProjectFavoriteResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProjectFavoriteResourceType(value)
	return nil
}

// NewProjectFavoriteResourceTypeFromValue returns a pointer to a valid ProjectFavoriteResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProjectFavoriteResourceTypeFromValue(v string) (*ProjectFavoriteResourceType, error) {
	ev := ProjectFavoriteResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProjectFavoriteResourceType: valid values are %v", v, allowedProjectFavoriteResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProjectFavoriteResourceType) IsValid() bool {
	for _, existing := range allowedProjectFavoriteResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProjectFavoriteResourceType value.
func (v ProjectFavoriteResourceType) Ptr() *ProjectFavoriteResourceType {
	return &v
}
