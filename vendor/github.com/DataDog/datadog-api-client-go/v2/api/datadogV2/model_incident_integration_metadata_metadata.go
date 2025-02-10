// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentIntegrationMetadataMetadata - Incident integration metadata's metadata attribute.
type IncidentIntegrationMetadataMetadata struct {
	SlackIntegrationMetadata   *SlackIntegrationMetadata
	JiraIntegrationMetadata    *JiraIntegrationMetadata
	MSTeamsIntegrationMetadata *MSTeamsIntegrationMetadata

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SlackIntegrationMetadataAsIncidentIntegrationMetadataMetadata is a convenience function that returns SlackIntegrationMetadata wrapped in IncidentIntegrationMetadataMetadata.
func SlackIntegrationMetadataAsIncidentIntegrationMetadataMetadata(v *SlackIntegrationMetadata) IncidentIntegrationMetadataMetadata {
	return IncidentIntegrationMetadataMetadata{SlackIntegrationMetadata: v}
}

// JiraIntegrationMetadataAsIncidentIntegrationMetadataMetadata is a convenience function that returns JiraIntegrationMetadata wrapped in IncidentIntegrationMetadataMetadata.
func JiraIntegrationMetadataAsIncidentIntegrationMetadataMetadata(v *JiraIntegrationMetadata) IncidentIntegrationMetadataMetadata {
	return IncidentIntegrationMetadataMetadata{JiraIntegrationMetadata: v}
}

// MSTeamsIntegrationMetadataAsIncidentIntegrationMetadataMetadata is a convenience function that returns MSTeamsIntegrationMetadata wrapped in IncidentIntegrationMetadataMetadata.
func MSTeamsIntegrationMetadataAsIncidentIntegrationMetadataMetadata(v *MSTeamsIntegrationMetadata) IncidentIntegrationMetadataMetadata {
	return IncidentIntegrationMetadataMetadata{MSTeamsIntegrationMetadata: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IncidentIntegrationMetadataMetadata) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SlackIntegrationMetadata
	err = datadog.Unmarshal(data, &obj.SlackIntegrationMetadata)
	if err == nil {
		if obj.SlackIntegrationMetadata != nil && obj.SlackIntegrationMetadata.UnparsedObject == nil {
			jsonSlackIntegrationMetadata, _ := datadog.Marshal(obj.SlackIntegrationMetadata)
			if string(jsonSlackIntegrationMetadata) == "{}" { // empty struct
				obj.SlackIntegrationMetadata = nil
			} else {
				match++
			}
		} else {
			obj.SlackIntegrationMetadata = nil
		}
	} else {
		obj.SlackIntegrationMetadata = nil
	}

	// try to unmarshal data into JiraIntegrationMetadata
	err = datadog.Unmarshal(data, &obj.JiraIntegrationMetadata)
	if err == nil {
		if obj.JiraIntegrationMetadata != nil && obj.JiraIntegrationMetadata.UnparsedObject == nil {
			jsonJiraIntegrationMetadata, _ := datadog.Marshal(obj.JiraIntegrationMetadata)
			if string(jsonJiraIntegrationMetadata) == "{}" { // empty struct
				obj.JiraIntegrationMetadata = nil
			} else {
				match++
			}
		} else {
			obj.JiraIntegrationMetadata = nil
		}
	} else {
		obj.JiraIntegrationMetadata = nil
	}

	// try to unmarshal data into MSTeamsIntegrationMetadata
	err = datadog.Unmarshal(data, &obj.MSTeamsIntegrationMetadata)
	if err == nil {
		if obj.MSTeamsIntegrationMetadata != nil && obj.MSTeamsIntegrationMetadata.UnparsedObject == nil {
			jsonMSTeamsIntegrationMetadata, _ := datadog.Marshal(obj.MSTeamsIntegrationMetadata)
			if string(jsonMSTeamsIntegrationMetadata) == "{}" { // empty struct
				obj.MSTeamsIntegrationMetadata = nil
			} else {
				match++
			}
		} else {
			obj.MSTeamsIntegrationMetadata = nil
		}
	} else {
		obj.MSTeamsIntegrationMetadata = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SlackIntegrationMetadata = nil
		obj.JiraIntegrationMetadata = nil
		obj.MSTeamsIntegrationMetadata = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IncidentIntegrationMetadataMetadata) MarshalJSON() ([]byte, error) {
	if obj.SlackIntegrationMetadata != nil {
		return datadog.Marshal(&obj.SlackIntegrationMetadata)
	}

	if obj.JiraIntegrationMetadata != nil {
		return datadog.Marshal(&obj.JiraIntegrationMetadata)
	}

	if obj.MSTeamsIntegrationMetadata != nil {
		return datadog.Marshal(&obj.MSTeamsIntegrationMetadata)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IncidentIntegrationMetadataMetadata) GetActualInstance() interface{} {
	if obj.SlackIntegrationMetadata != nil {
		return obj.SlackIntegrationMetadata
	}

	if obj.JiraIntegrationMetadata != nil {
		return obj.JiraIntegrationMetadata
	}

	if obj.MSTeamsIntegrationMetadata != nil {
		return obj.MSTeamsIntegrationMetadata
	}

	// all schemas are nil
	return nil
}
