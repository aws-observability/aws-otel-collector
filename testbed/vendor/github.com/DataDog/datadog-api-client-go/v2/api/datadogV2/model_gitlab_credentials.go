// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitlabCredentials - The definition of the `GitlabCredentials` object.
type GitlabCredentials struct {
	GitlabAPIKey *GitlabAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GitlabAPIKeyAsGitlabCredentials is a convenience function that returns GitlabAPIKey wrapped in GitlabCredentials.
func GitlabAPIKeyAsGitlabCredentials(v *GitlabAPIKey) GitlabCredentials {
	return GitlabCredentials{GitlabAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GitlabCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GitlabAPIKey
	err = datadog.Unmarshal(data, &obj.GitlabAPIKey)
	if err == nil {
		if obj.GitlabAPIKey != nil && obj.GitlabAPIKey.UnparsedObject == nil {
			jsonGitlabAPIKey, _ := datadog.Marshal(obj.GitlabAPIKey)
			if string(jsonGitlabAPIKey) == "{}" { // empty struct
				obj.GitlabAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.GitlabAPIKey = nil
		}
	} else {
		obj.GitlabAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GitlabAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GitlabCredentials) MarshalJSON() ([]byte, error) {
	if obj.GitlabAPIKey != nil {
		return datadog.Marshal(&obj.GitlabAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GitlabCredentials) GetActualInstance() interface{} {
	if obj.GitlabAPIKey != nil {
		return obj.GitlabAPIKey
	}

	// all schemas are nil
	return nil
}
