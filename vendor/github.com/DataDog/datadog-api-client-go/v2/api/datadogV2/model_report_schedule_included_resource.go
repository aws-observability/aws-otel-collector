// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleIncludedResource - A related resource included with a report schedule.
type ReportScheduleIncludedResource struct {
	ReportScheduleAuthor *ReportScheduleAuthor

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ReportScheduleAuthorAsReportScheduleIncludedResource is a convenience function that returns ReportScheduleAuthor wrapped in ReportScheduleIncludedResource.
func ReportScheduleAuthorAsReportScheduleIncludedResource(v *ReportScheduleAuthor) ReportScheduleIncludedResource {
	return ReportScheduleIncludedResource{ReportScheduleAuthor: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ReportScheduleIncludedResource) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ReportScheduleAuthor
	err = datadog.Unmarshal(data, &obj.ReportScheduleAuthor)
	if err == nil {
		if obj.ReportScheduleAuthor != nil && obj.ReportScheduleAuthor.UnparsedObject == nil {
			jsonReportScheduleAuthor, _ := datadog.Marshal(obj.ReportScheduleAuthor)
			if string(jsonReportScheduleAuthor) == "{}" { // empty struct
				obj.ReportScheduleAuthor = nil
			} else {
				match++
			}
		} else {
			obj.ReportScheduleAuthor = nil
		}
	} else {
		obj.ReportScheduleAuthor = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ReportScheduleAuthor = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ReportScheduleIncludedResource) MarshalJSON() ([]byte, error) {
	if obj.ReportScheduleAuthor != nil {
		return datadog.Marshal(&obj.ReportScheduleAuthor)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ReportScheduleIncludedResource) GetActualInstance() interface{} {
	if obj.ReportScheduleAuthor != nil {
		return obj.ReportScheduleAuthor
	}

	// all schemas are nil
	return nil
}
