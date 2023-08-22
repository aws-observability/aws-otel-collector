// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
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
	err = json.Unmarshal(data, &obj.CIAppPipelineEventPipeline)
	if err == nil {
		if obj.CIAppPipelineEventPipeline != nil && obj.CIAppPipelineEventPipeline.UnparsedObject == nil {
			jsonCIAppPipelineEventPipeline, _ := json.Marshal(obj.CIAppPipelineEventPipeline)
			if string(jsonCIAppPipelineEventPipeline) == "{}" { // empty struct
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
	err = json.Unmarshal(data, &obj.CIAppPipelineEventStage)
	if err == nil {
		if obj.CIAppPipelineEventStage != nil && obj.CIAppPipelineEventStage.UnparsedObject == nil {
			jsonCIAppPipelineEventStage, _ := json.Marshal(obj.CIAppPipelineEventStage)
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
	err = json.Unmarshal(data, &obj.CIAppPipelineEventJob)
	if err == nil {
		if obj.CIAppPipelineEventJob != nil && obj.CIAppPipelineEventJob.UnparsedObject == nil {
			jsonCIAppPipelineEventJob, _ := json.Marshal(obj.CIAppPipelineEventJob)
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
	err = json.Unmarshal(data, &obj.CIAppPipelineEventStep)
	if err == nil {
		if obj.CIAppPipelineEventStep != nil && obj.CIAppPipelineEventStep.UnparsedObject == nil {
			jsonCIAppPipelineEventStep, _ := json.Marshal(obj.CIAppPipelineEventStep)
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
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CIAppCreatePipelineEventRequestAttributesResource) MarshalJSON() ([]byte, error) {
	if obj.CIAppPipelineEventPipeline != nil {
		return json.Marshal(&obj.CIAppPipelineEventPipeline)
	}

	if obj.CIAppPipelineEventStage != nil {
		return json.Marshal(&obj.CIAppPipelineEventStage)
	}

	if obj.CIAppPipelineEventJob != nil {
		return json.Marshal(&obj.CIAppPipelineEventJob)
	}

	if obj.CIAppPipelineEventStep != nil {
		return json.Marshal(&obj.CIAppPipelineEventStep)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
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

// NullableCIAppCreatePipelineEventRequestAttributesResource handles when a null is used for CIAppCreatePipelineEventRequestAttributesResource.
type NullableCIAppCreatePipelineEventRequestAttributesResource struct {
	value *CIAppCreatePipelineEventRequestAttributesResource
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppCreatePipelineEventRequestAttributesResource) Get() *CIAppCreatePipelineEventRequestAttributesResource {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppCreatePipelineEventRequestAttributesResource) Set(val *CIAppCreatePipelineEventRequestAttributesResource) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppCreatePipelineEventRequestAttributesResource) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableCIAppCreatePipelineEventRequestAttributesResource) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppCreatePipelineEventRequestAttributesResource initializes the struct as if Set has been called.
func NewNullableCIAppCreatePipelineEventRequestAttributesResource(val *CIAppCreatePipelineEventRequestAttributesResource) *NullableCIAppCreatePipelineEventRequestAttributesResource {
	return &NullableCIAppCreatePipelineEventRequestAttributesResource{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppCreatePipelineEventRequestAttributesResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppCreatePipelineEventRequestAttributesResource) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
