// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateJiraIssueRequestArrayIncluded - Attributes and relationships of the case linked to the Jira issue. Should contain all of the following: case, project, and security findings.
type CreateJiraIssueRequestArrayIncluded struct {
	CreateCaseRequestData     *CreateCaseRequestData
	CaseManagementProjectData *CaseManagementProjectData
	FindingData               *FindingData

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CreateCaseRequestDataAsCreateJiraIssueRequestArrayIncluded is a convenience function that returns CreateCaseRequestData wrapped in CreateJiraIssueRequestArrayIncluded.
func CreateCaseRequestDataAsCreateJiraIssueRequestArrayIncluded(v *CreateCaseRequestData) CreateJiraIssueRequestArrayIncluded {
	return CreateJiraIssueRequestArrayIncluded{CreateCaseRequestData: v}
}

// CaseManagementProjectDataAsCreateJiraIssueRequestArrayIncluded is a convenience function that returns CaseManagementProjectData wrapped in CreateJiraIssueRequestArrayIncluded.
func CaseManagementProjectDataAsCreateJiraIssueRequestArrayIncluded(v *CaseManagementProjectData) CreateJiraIssueRequestArrayIncluded {
	return CreateJiraIssueRequestArrayIncluded{CaseManagementProjectData: v}
}

// FindingDataAsCreateJiraIssueRequestArrayIncluded is a convenience function that returns FindingData wrapped in CreateJiraIssueRequestArrayIncluded.
func FindingDataAsCreateJiraIssueRequestArrayIncluded(v *FindingData) CreateJiraIssueRequestArrayIncluded {
	return CreateJiraIssueRequestArrayIncluded{FindingData: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CreateJiraIssueRequestArrayIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateCaseRequestData
	err = datadog.Unmarshal(data, &obj.CreateCaseRequestData)
	if err == nil {
		if obj.CreateCaseRequestData != nil && obj.CreateCaseRequestData.UnparsedObject == nil {
			jsonCreateCaseRequestData, _ := datadog.Marshal(obj.CreateCaseRequestData)
			if string(jsonCreateCaseRequestData) == "{}" { // empty struct
				obj.CreateCaseRequestData = nil
			} else {
				match++
			}
		} else {
			obj.CreateCaseRequestData = nil
		}
	} else {
		obj.CreateCaseRequestData = nil
	}

	// try to unmarshal data into CaseManagementProjectData
	err = datadog.Unmarshal(data, &obj.CaseManagementProjectData)
	if err == nil {
		if obj.CaseManagementProjectData != nil && obj.CaseManagementProjectData.UnparsedObject == nil {
			jsonCaseManagementProjectData, _ := datadog.Marshal(obj.CaseManagementProjectData)
			if string(jsonCaseManagementProjectData) == "{}" { // empty struct
				obj.CaseManagementProjectData = nil
			} else {
				match++
			}
		} else {
			obj.CaseManagementProjectData = nil
		}
	} else {
		obj.CaseManagementProjectData = nil
	}

	// try to unmarshal data into FindingData
	err = datadog.Unmarshal(data, &obj.FindingData)
	if err == nil {
		if obj.FindingData != nil && obj.FindingData.UnparsedObject == nil {
			jsonFindingData, _ := datadog.Marshal(obj.FindingData)
			if string(jsonFindingData) == "{}" { // empty struct
				obj.FindingData = nil
			} else {
				match++
			}
		} else {
			obj.FindingData = nil
		}
	} else {
		obj.FindingData = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CreateCaseRequestData = nil
		obj.CaseManagementProjectData = nil
		obj.FindingData = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CreateJiraIssueRequestArrayIncluded) MarshalJSON() ([]byte, error) {
	if obj.CreateCaseRequestData != nil {
		return datadog.Marshal(&obj.CreateCaseRequestData)
	}

	if obj.CaseManagementProjectData != nil {
		return datadog.Marshal(&obj.CaseManagementProjectData)
	}

	if obj.FindingData != nil {
		return datadog.Marshal(&obj.FindingData)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CreateJiraIssueRequestArrayIncluded) GetActualInstance() interface{} {
	if obj.CreateCaseRequestData != nil {
		return obj.CreateCaseRequestData
	}

	if obj.CaseManagementProjectData != nil {
		return obj.CaseManagementProjectData
	}

	if obj.FindingData != nil {
		return obj.FindingData
	}

	// all schemas are nil
	return nil
}
