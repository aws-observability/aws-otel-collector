// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitlabCredentialsUpdate - The definition of the `GitlabCredentialsUpdate` object.
type GitlabCredentialsUpdate struct {
	GitlabAPIKeyUpdate *GitlabAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GitlabAPIKeyUpdateAsGitlabCredentialsUpdate is a convenience function that returns GitlabAPIKeyUpdate wrapped in GitlabCredentialsUpdate.
func GitlabAPIKeyUpdateAsGitlabCredentialsUpdate(v *GitlabAPIKeyUpdate) GitlabCredentialsUpdate {
	return GitlabCredentialsUpdate{GitlabAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GitlabCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GitlabAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.GitlabAPIKeyUpdate)
	if err == nil {
		if obj.GitlabAPIKeyUpdate != nil && obj.GitlabAPIKeyUpdate.UnparsedObject == nil {
			jsonGitlabAPIKeyUpdate, _ := datadog.Marshal(obj.GitlabAPIKeyUpdate)
			if string(jsonGitlabAPIKeyUpdate) == "{}" { // empty struct
				obj.GitlabAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.GitlabAPIKeyUpdate = nil
		}
	} else {
		obj.GitlabAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GitlabAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GitlabCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.GitlabAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.GitlabAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GitlabCredentialsUpdate) GetActualInstance() interface{} {
	if obj.GitlabAPIKeyUpdate != nil {
		return obj.GitlabAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
