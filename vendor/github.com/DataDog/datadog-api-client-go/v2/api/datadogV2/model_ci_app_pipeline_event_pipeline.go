// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineEventPipeline - Details of the top level pipeline, build, or workflow of your CI.
type CIAppPipelineEventPipeline struct {
	CIAppPipelineEventFinishedPipeline   *CIAppPipelineEventFinishedPipeline
	CIAppPipelineEventInProgressPipeline *CIAppPipelineEventInProgressPipeline

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CIAppPipelineEventFinishedPipelineAsCIAppPipelineEventPipeline is a convenience function that returns CIAppPipelineEventFinishedPipeline wrapped in CIAppPipelineEventPipeline.
func CIAppPipelineEventFinishedPipelineAsCIAppPipelineEventPipeline(v *CIAppPipelineEventFinishedPipeline) CIAppPipelineEventPipeline {
	return CIAppPipelineEventPipeline{CIAppPipelineEventFinishedPipeline: v}
}

// CIAppPipelineEventInProgressPipelineAsCIAppPipelineEventPipeline is a convenience function that returns CIAppPipelineEventInProgressPipeline wrapped in CIAppPipelineEventPipeline.
func CIAppPipelineEventInProgressPipelineAsCIAppPipelineEventPipeline(v *CIAppPipelineEventInProgressPipeline) CIAppPipelineEventPipeline {
	return CIAppPipelineEventPipeline{CIAppPipelineEventInProgressPipeline: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CIAppPipelineEventPipeline) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CIAppPipelineEventFinishedPipeline
	err = datadog.Unmarshal(data, &obj.CIAppPipelineEventFinishedPipeline)
	if err == nil {
		if obj.CIAppPipelineEventFinishedPipeline != nil && obj.CIAppPipelineEventFinishedPipeline.UnparsedObject == nil {
			jsonCIAppPipelineEventFinishedPipeline, _ := datadog.Marshal(obj.CIAppPipelineEventFinishedPipeline)
			if string(jsonCIAppPipelineEventFinishedPipeline) == "{}" { // empty struct
				obj.CIAppPipelineEventFinishedPipeline = nil
			} else {
				match++
			}
		} else {
			obj.CIAppPipelineEventFinishedPipeline = nil
		}
	} else {
		obj.CIAppPipelineEventFinishedPipeline = nil
	}

	// try to unmarshal data into CIAppPipelineEventInProgressPipeline
	err = datadog.Unmarshal(data, &obj.CIAppPipelineEventInProgressPipeline)
	if err == nil {
		if obj.CIAppPipelineEventInProgressPipeline != nil && obj.CIAppPipelineEventInProgressPipeline.UnparsedObject == nil {
			jsonCIAppPipelineEventInProgressPipeline, _ := datadog.Marshal(obj.CIAppPipelineEventInProgressPipeline)
			if string(jsonCIAppPipelineEventInProgressPipeline) == "{}" { // empty struct
				obj.CIAppPipelineEventInProgressPipeline = nil
			} else {
				match++
			}
		} else {
			obj.CIAppPipelineEventInProgressPipeline = nil
		}
	} else {
		obj.CIAppPipelineEventInProgressPipeline = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CIAppPipelineEventFinishedPipeline = nil
		obj.CIAppPipelineEventInProgressPipeline = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CIAppPipelineEventPipeline) MarshalJSON() ([]byte, error) {
	if obj.CIAppPipelineEventFinishedPipeline != nil {
		return datadog.Marshal(&obj.CIAppPipelineEventFinishedPipeline)
	}

	if obj.CIAppPipelineEventInProgressPipeline != nil {
		return datadog.Marshal(&obj.CIAppPipelineEventInProgressPipeline)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CIAppPipelineEventPipeline) GetActualInstance() interface{} {
	if obj.CIAppPipelineEventFinishedPipeline != nil {
		return obj.CIAppPipelineEventFinishedPipeline
	}

	if obj.CIAppPipelineEventInProgressPipeline != nil {
		return obj.CIAppPipelineEventInProgressPipeline
	}

	// all schemas are nil
	return nil
}
