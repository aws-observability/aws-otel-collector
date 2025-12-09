// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitlabAPIKeyType The definition of the `GitlabAPIKey` object.
type GitlabAPIKeyType string

// List of GitlabAPIKeyType.
const (
	GITLABAPIKEYTYPE_GITLABAPIKEY GitlabAPIKeyType = "GitlabAPIKey"
)

var allowedGitlabAPIKeyTypeEnumValues = []GitlabAPIKeyType{
	GITLABAPIKEYTYPE_GITLABAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GitlabAPIKeyType) GetAllowedValues() []GitlabAPIKeyType {
	return allowedGitlabAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GitlabAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GitlabAPIKeyType(value)
	return nil
}

// NewGitlabAPIKeyTypeFromValue returns a pointer to a valid GitlabAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGitlabAPIKeyTypeFromValue(v string) (*GitlabAPIKeyType, error) {
	ev := GitlabAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GitlabAPIKeyType: valid values are %v", v, allowedGitlabAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GitlabAPIKeyType) IsValid() bool {
	for _, existing := range allowedGitlabAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GitlabAPIKeyType value.
func (v GitlabAPIKeyType) Ptr() *GitlabAPIKeyType {
	return &v
}
