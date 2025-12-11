// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitlabIntegrationType The definition of the `GitlabIntegrationType` object.
type GitlabIntegrationType string

// List of GitlabIntegrationType.
const (
	GITLABINTEGRATIONTYPE_GITLAB GitlabIntegrationType = "Gitlab"
)

var allowedGitlabIntegrationTypeEnumValues = []GitlabIntegrationType{
	GITLABINTEGRATIONTYPE_GITLAB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GitlabIntegrationType) GetAllowedValues() []GitlabIntegrationType {
	return allowedGitlabIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GitlabIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GitlabIntegrationType(value)
	return nil
}

// NewGitlabIntegrationTypeFromValue returns a pointer to a valid GitlabIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGitlabIntegrationTypeFromValue(v string) (*GitlabIntegrationType, error) {
	ev := GitlabIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GitlabIntegrationType: valid values are %v", v, allowedGitlabIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GitlabIntegrationType) IsValid() bool {
	for _, existing := range allowedGitlabIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GitlabIntegrationType value.
func (v GitlabIntegrationType) Ptr() *GitlabIntegrationType {
	return &v
}
