// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimelineCellCreateAttributes - The timeline cell's attributes for a create request.
type IncidentTimelineCellCreateAttributes struct {
	IncidentTimelineCellMarkdownCreateAttributes *IncidentTimelineCellMarkdownCreateAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IncidentTimelineCellMarkdownCreateAttributesAsIncidentTimelineCellCreateAttributes is a convenience function that returns IncidentTimelineCellMarkdownCreateAttributes wrapped in IncidentTimelineCellCreateAttributes.
func IncidentTimelineCellMarkdownCreateAttributesAsIncidentTimelineCellCreateAttributes(v *IncidentTimelineCellMarkdownCreateAttributes) IncidentTimelineCellCreateAttributes {
	return IncidentTimelineCellCreateAttributes{IncidentTimelineCellMarkdownCreateAttributes: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IncidentTimelineCellCreateAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IncidentTimelineCellMarkdownCreateAttributes
	err = datadog.Unmarshal(data, &obj.IncidentTimelineCellMarkdownCreateAttributes)
	if err == nil {
		if obj.IncidentTimelineCellMarkdownCreateAttributes != nil && obj.IncidentTimelineCellMarkdownCreateAttributes.UnparsedObject == nil {
			jsonIncidentTimelineCellMarkdownCreateAttributes, _ := datadog.Marshal(obj.IncidentTimelineCellMarkdownCreateAttributes)
			if string(jsonIncidentTimelineCellMarkdownCreateAttributes) == "{}" { // empty struct
				obj.IncidentTimelineCellMarkdownCreateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.IncidentTimelineCellMarkdownCreateAttributes = nil
		}
	} else {
		obj.IncidentTimelineCellMarkdownCreateAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.IncidentTimelineCellMarkdownCreateAttributes = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IncidentTimelineCellCreateAttributes) MarshalJSON() ([]byte, error) {
	if obj.IncidentTimelineCellMarkdownCreateAttributes != nil {
		return datadog.Marshal(&obj.IncidentTimelineCellMarkdownCreateAttributes)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IncidentTimelineCellCreateAttributes) GetActualInstance() interface{} {
	if obj.IncidentTimelineCellMarkdownCreateAttributes != nil {
		return obj.IncidentTimelineCellMarkdownCreateAttributes
	}

	// all schemas are nil
	return nil
}
