// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssuesSearchResultIncluded - An array of related resources, returned when the `include` query parameter is used.
type IssuesSearchResultIncluded struct {
	Issue     *Issue
	Case      *Case
	IssueUser *IssueUser
	IssueTeam *IssueTeam

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IssueAsIssuesSearchResultIncluded is a convenience function that returns Issue wrapped in IssuesSearchResultIncluded.
func IssueAsIssuesSearchResultIncluded(v *Issue) IssuesSearchResultIncluded {
	return IssuesSearchResultIncluded{Issue: v}
}

// CaseAsIssuesSearchResultIncluded is a convenience function that returns Case wrapped in IssuesSearchResultIncluded.
func CaseAsIssuesSearchResultIncluded(v *Case) IssuesSearchResultIncluded {
	return IssuesSearchResultIncluded{Case: v}
}

// IssueUserAsIssuesSearchResultIncluded is a convenience function that returns IssueUser wrapped in IssuesSearchResultIncluded.
func IssueUserAsIssuesSearchResultIncluded(v *IssueUser) IssuesSearchResultIncluded {
	return IssuesSearchResultIncluded{IssueUser: v}
}

// IssueTeamAsIssuesSearchResultIncluded is a convenience function that returns IssueTeam wrapped in IssuesSearchResultIncluded.
func IssueTeamAsIssuesSearchResultIncluded(v *IssueTeam) IssuesSearchResultIncluded {
	return IssuesSearchResultIncluded{IssueTeam: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IssuesSearchResultIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Issue
	err = datadog.Unmarshal(data, &obj.Issue)
	if err == nil {
		if obj.Issue != nil && obj.Issue.UnparsedObject == nil {
			jsonIssue, _ := datadog.Marshal(obj.Issue)
			if string(jsonIssue) == "{}" { // empty struct
				obj.Issue = nil
			} else {
				match++
			}
		} else {
			obj.Issue = nil
		}
	} else {
		obj.Issue = nil
	}

	// try to unmarshal data into Case
	err = datadog.Unmarshal(data, &obj.Case)
	if err == nil {
		if obj.Case != nil && obj.Case.UnparsedObject == nil {
			jsonCase, _ := datadog.Marshal(obj.Case)
			if string(jsonCase) == "{}" { // empty struct
				obj.Case = nil
			} else {
				match++
			}
		} else {
			obj.Case = nil
		}
	} else {
		obj.Case = nil
	}

	// try to unmarshal data into IssueUser
	err = datadog.Unmarshal(data, &obj.IssueUser)
	if err == nil {
		if obj.IssueUser != nil && obj.IssueUser.UnparsedObject == nil {
			jsonIssueUser, _ := datadog.Marshal(obj.IssueUser)
			if string(jsonIssueUser) == "{}" { // empty struct
				obj.IssueUser = nil
			} else {
				match++
			}
		} else {
			obj.IssueUser = nil
		}
	} else {
		obj.IssueUser = nil
	}

	// try to unmarshal data into IssueTeam
	err = datadog.Unmarshal(data, &obj.IssueTeam)
	if err == nil {
		if obj.IssueTeam != nil && obj.IssueTeam.UnparsedObject == nil {
			jsonIssueTeam, _ := datadog.Marshal(obj.IssueTeam)
			if string(jsonIssueTeam) == "{}" { // empty struct
				obj.IssueTeam = nil
			} else {
				match++
			}
		} else {
			obj.IssueTeam = nil
		}
	} else {
		obj.IssueTeam = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.Issue = nil
		obj.Case = nil
		obj.IssueUser = nil
		obj.IssueTeam = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IssuesSearchResultIncluded) MarshalJSON() ([]byte, error) {
	if obj.Issue != nil {
		return datadog.Marshal(&obj.Issue)
	}

	if obj.Case != nil {
		return datadog.Marshal(&obj.Case)
	}

	if obj.IssueUser != nil {
		return datadog.Marshal(&obj.IssueUser)
	}

	if obj.IssueTeam != nil {
		return datadog.Marshal(&obj.IssueTeam)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IssuesSearchResultIncluded) GetActualInstance() interface{} {
	if obj.Issue != nil {
		return obj.Issue
	}

	if obj.Case != nil {
		return obj.Case
	}

	if obj.IssueUser != nil {
		return obj.IssueUser
	}

	if obj.IssueTeam != nil {
		return obj.IssueTeam
	}

	// all schemas are nil
	return nil
}
