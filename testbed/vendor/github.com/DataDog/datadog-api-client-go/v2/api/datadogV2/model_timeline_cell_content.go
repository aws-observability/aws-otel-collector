// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimelineCellContent - timeline cell content
type TimelineCellContent struct {
	TimelineCellContentComment *TimelineCellContentComment

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TimelineCellContentCommentAsTimelineCellContent is a convenience function that returns TimelineCellContentComment wrapped in TimelineCellContent.
func TimelineCellContentCommentAsTimelineCellContent(v *TimelineCellContentComment) TimelineCellContent {
	return TimelineCellContent{TimelineCellContentComment: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TimelineCellContent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TimelineCellContentComment
	err = datadog.Unmarshal(data, &obj.TimelineCellContentComment)
	if err == nil {
		if obj.TimelineCellContentComment != nil && obj.TimelineCellContentComment.UnparsedObject == nil {
			jsonTimelineCellContentComment, _ := datadog.Marshal(obj.TimelineCellContentComment)
			if string(jsonTimelineCellContentComment) == "{}" && string(data) != "{}" { // empty struct
				obj.TimelineCellContentComment = nil
			} else {
				match++
			}
		} else {
			obj.TimelineCellContentComment = nil
		}
	} else {
		obj.TimelineCellContentComment = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TimelineCellContentComment = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TimelineCellContent) MarshalJSON() ([]byte, error) {
	if obj.TimelineCellContentComment != nil {
		return datadog.Marshal(&obj.TimelineCellContentComment)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TimelineCellContent) GetActualInstance() interface{} {
	if obj.TimelineCellContentComment != nil {
		return obj.TimelineCellContentComment
	}

	// all schemas are nil
	return nil
}
