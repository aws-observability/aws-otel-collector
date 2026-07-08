// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineEventJob - Details of a CI job.
type CIAppPipelineEventJob struct {
	CIAppPipelineEventFinishedJob   *CIAppPipelineEventFinishedJob
	CIAppPipelineEventInProgressJob *CIAppPipelineEventInProgressJob

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CIAppPipelineEventFinishedJobAsCIAppPipelineEventJob is a convenience function that returns CIAppPipelineEventFinishedJob wrapped in CIAppPipelineEventJob.
func CIAppPipelineEventFinishedJobAsCIAppPipelineEventJob(v *CIAppPipelineEventFinishedJob) CIAppPipelineEventJob {
	return CIAppPipelineEventJob{CIAppPipelineEventFinishedJob: v}
}

// CIAppPipelineEventInProgressJobAsCIAppPipelineEventJob is a convenience function that returns CIAppPipelineEventInProgressJob wrapped in CIAppPipelineEventJob.
func CIAppPipelineEventInProgressJobAsCIAppPipelineEventJob(v *CIAppPipelineEventInProgressJob) CIAppPipelineEventJob {
	return CIAppPipelineEventJob{CIAppPipelineEventInProgressJob: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CIAppPipelineEventJob) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CIAppPipelineEventFinishedJob
	err = datadog.Unmarshal(data, &obj.CIAppPipelineEventFinishedJob)
	if err == nil {
		if obj.CIAppPipelineEventFinishedJob != nil && obj.CIAppPipelineEventFinishedJob.UnparsedObject == nil {
			jsonCIAppPipelineEventFinishedJob, _ := datadog.Marshal(obj.CIAppPipelineEventFinishedJob)
			if string(jsonCIAppPipelineEventFinishedJob) == "{}" { // empty struct
				obj.CIAppPipelineEventFinishedJob = nil
			} else {
				match++
			}
		} else {
			obj.CIAppPipelineEventFinishedJob = nil
		}
	} else {
		obj.CIAppPipelineEventFinishedJob = nil
	}

	// try to unmarshal data into CIAppPipelineEventInProgressJob
	err = datadog.Unmarshal(data, &obj.CIAppPipelineEventInProgressJob)
	if err == nil {
		if obj.CIAppPipelineEventInProgressJob != nil && obj.CIAppPipelineEventInProgressJob.UnparsedObject == nil {
			jsonCIAppPipelineEventInProgressJob, _ := datadog.Marshal(obj.CIAppPipelineEventInProgressJob)
			if string(jsonCIAppPipelineEventInProgressJob) == "{}" { // empty struct
				obj.CIAppPipelineEventInProgressJob = nil
			} else {
				match++
			}
		} else {
			obj.CIAppPipelineEventInProgressJob = nil
		}
	} else {
		obj.CIAppPipelineEventInProgressJob = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CIAppPipelineEventFinishedJob = nil
		obj.CIAppPipelineEventInProgressJob = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CIAppPipelineEventJob) MarshalJSON() ([]byte, error) {
	if obj.CIAppPipelineEventFinishedJob != nil {
		return datadog.Marshal(&obj.CIAppPipelineEventFinishedJob)
	}

	if obj.CIAppPipelineEventInProgressJob != nil {
		return datadog.Marshal(&obj.CIAppPipelineEventInProgressJob)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CIAppPipelineEventJob) GetActualInstance() interface{} {
	if obj.CIAppPipelineEventFinishedJob != nil {
		return obj.CIAppPipelineEventFinishedJob
	}

	if obj.CIAppPipelineEventInProgressJob != nil {
		return obj.CIAppPipelineEventInProgressJob
	}

	// all schemas are nil
	return nil
}
