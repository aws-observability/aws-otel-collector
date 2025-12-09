// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppCreatePipelineEventRequestAttributesResource - Details of the CI pipeline event.
type CIAppCreatePipelineEventRequestAttributesResource struct {
	CIAppPipelineEventPipeline *CIAppPipelineEventPipeline
	CIAppPipelineEventStage    *CIAppPipelineEventStage
	CIAppPipelineEventJob      *CIAppPipelineEventJob
	CIAppPipelineEventStep     *CIAppPipelineEventStep

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CIAppPipelineEventPipelineAsCIAppCreatePipelineEventRequestAttributesResource is a convenience function that returns CIAppPipelineEventPipeline wrapped in CIAppCreatePipelineEventRequestAttributesResource.
func CIAppPipelineEventPipelineAsCIAppCreatePipelineEventRequestAttributesResource(v *CIAppPipelineEventPipeline) CIAppCreatePipelineEventRequestAttributesResource {
	return CIAppCreatePipelineEventRequestAttributesResource{CIAppPipelineEventPipeline: v}
}

// CIAppPipelineEventStageAsCIAppCreatePipelineEventRequestAttributesResource is a convenience function that returns CIAppPipelineEventStage wrapped in CIAppCreatePipelineEventRequestAttributesResource.
func CIAppPipelineEventStageAsCIAppCreatePipelineEventRequestAttributesResource(v *CIAppPipelineEventStage) CIAppCreatePipelineEventRequestAttributesResource {
	return CIAppCreatePipelineEventRequestAttributesResource{CIAppPipelineEventStage: v}
}

// CIAppPipelineEventJobAsCIAppCreatePipelineEventRequestAttributesResource is a convenience function that returns CIAppPipelineEventJob wrapped in CIAppCreatePipelineEventRequestAttributesResource.
func CIAppPipelineEventJobAsCIAppCreatePipelineEventRequestAttributesResource(v *CIAppPipelineEventJob) CIAppCreatePipelineEventRequestAttributesResource {
	return CIAppCreatePipelineEventRequestAttributesResource{CIAppPipelineEventJob: v}
}

// CIAppPipelineEventStepAsCIAppCreatePipelineEventRequestAttributesResource is a convenience function that returns CIAppPipelineEventStep wrapped in CIAppCreatePipelineEventRequestAttributesResource.
func CIAppPipelineEventStepAsCIAppCreatePipelineEventRequestAttributesResource(v *CIAppPipelineEventStep) CIAppCreatePipelineEventRequestAttributesResource {
	return CIAppCreatePipelineEventRequestAttributesResource{CIAppPipelineEventStep: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CIAppCreatePipelineEventRequestAttributesResource) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CIAppPipelineEventPipeline
	err = datadog.Unmarshal(data, &obj.CIAppPipelineEventPipeline)
	if err == nil {
		if obj.CIAppPipelineEventPipeline != nil && obj.CIAppPipelineEventPipeline.UnparsedObject == nil {
			jsonCIAppPipelineEventPipeline, _ := datadog.Marshal(obj.CIAppPipelineEventPipeline)
			if string(jsonCIAppPipelineEventPipeline) == "{}" && string(data) != "{}" { // empty struct
				obj.CIAppPipelineEventPipeline = nil
			} else {
				match++
			}
		} else {
			obj.CIAppPipelineEventPipeline = nil
		}
	} else {
		obj.CIAppPipelineEventPipeline = nil
	}

	// try to unmarshal data into CIAppPipelineEventStage
	err = datadog.Unmarshal(data, &obj.CIAppPipelineEventStage)
	if err == nil {
		if obj.CIAppPipelineEventStage != nil && obj.CIAppPipelineEventStage.UnparsedObject == nil {
			jsonCIAppPipelineEventStage, _ := datadog.Marshal(obj.CIAppPipelineEventStage)
			if string(jsonCIAppPipelineEventStage) == "{}" { // empty struct
				obj.CIAppPipelineEventStage = nil
			} else {
				match++
			}
		} else {
			obj.CIAppPipelineEventStage = nil
		}
	} else {
		obj.CIAppPipelineEventStage = nil
	}

	// try to unmarshal data into CIAppPipelineEventJob
	err = datadog.Unmarshal(data, &obj.CIAppPipelineEventJob)
	if err == nil {
		if obj.CIAppPipelineEventJob != nil && obj.CIAppPipelineEventJob.UnparsedObject == nil {
			jsonCIAppPipelineEventJob, _ := datadog.Marshal(obj.CIAppPipelineEventJob)
			if string(jsonCIAppPipelineEventJob) == "{}" { // empty struct
				obj.CIAppPipelineEventJob = nil
			} else {
				match++
			}
		} else {
			obj.CIAppPipelineEventJob = nil
		}
	} else {
		obj.CIAppPipelineEventJob = nil
	}

	// try to unmarshal data into CIAppPipelineEventStep
	err = datadog.Unmarshal(data, &obj.CIAppPipelineEventStep)
	if err == nil {
		if obj.CIAppPipelineEventStep != nil && obj.CIAppPipelineEventStep.UnparsedObject == nil {
			jsonCIAppPipelineEventStep, _ := datadog.Marshal(obj.CIAppPipelineEventStep)
			if string(jsonCIAppPipelineEventStep) == "{}" { // empty struct
				obj.CIAppPipelineEventStep = nil
			} else {
				match++
			}
		} else {
			obj.CIAppPipelineEventStep = nil
		}
	} else {
		obj.CIAppPipelineEventStep = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CIAppPipelineEventPipeline = nil
		obj.CIAppPipelineEventStage = nil
		obj.CIAppPipelineEventJob = nil
		obj.CIAppPipelineEventStep = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CIAppCreatePipelineEventRequestAttributesResource) MarshalJSON() ([]byte, error) {
	if obj.CIAppPipelineEventPipeline != nil {
		return datadog.Marshal(&obj.CIAppPipelineEventPipeline)
	}

	if obj.CIAppPipelineEventStage != nil {
		return datadog.Marshal(&obj.CIAppPipelineEventStage)
	}

	if obj.CIAppPipelineEventJob != nil {
		return datadog.Marshal(&obj.CIAppPipelineEventJob)
	}

	if obj.CIAppPipelineEventStep != nil {
		return datadog.Marshal(&obj.CIAppPipelineEventStep)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CIAppCreatePipelineEventRequestAttributesResource) GetActualInstance() interface{} {
	if obj.CIAppPipelineEventPipeline != nil {
		return obj.CIAppPipelineEventPipeline
	}

	if obj.CIAppPipelineEventStage != nil {
		return obj.CIAppPipelineEventStage
	}

	if obj.CIAppPipelineEventJob != nil {
		return obj.CIAppPipelineEventJob
	}

	if obj.CIAppPipelineEventStep != nil {
		return obj.CIAppPipelineEventStep
	}

	// all schemas are nil
	return nil
}
