// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAPIStep - The steps used in a Synthetic multi-step API test.
type SyntheticsAPIStep struct {
	SyntheticsAPITestStep    *SyntheticsAPITestStep
	SyntheticsAPIWaitStep    *SyntheticsAPIWaitStep
	SyntheticsAPISubtestStep *SyntheticsAPISubtestStep

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SyntheticsAPITestStepAsSyntheticsAPIStep is a convenience function that returns SyntheticsAPITestStep wrapped in SyntheticsAPIStep.
func SyntheticsAPITestStepAsSyntheticsAPIStep(v *SyntheticsAPITestStep) SyntheticsAPIStep {
	return SyntheticsAPIStep{SyntheticsAPITestStep: v}
}

// SyntheticsAPIWaitStepAsSyntheticsAPIStep is a convenience function that returns SyntheticsAPIWaitStep wrapped in SyntheticsAPIStep.
func SyntheticsAPIWaitStepAsSyntheticsAPIStep(v *SyntheticsAPIWaitStep) SyntheticsAPIStep {
	return SyntheticsAPIStep{SyntheticsAPIWaitStep: v}
}

// SyntheticsAPISubtestStepAsSyntheticsAPIStep is a convenience function that returns SyntheticsAPISubtestStep wrapped in SyntheticsAPIStep.
func SyntheticsAPISubtestStepAsSyntheticsAPIStep(v *SyntheticsAPISubtestStep) SyntheticsAPIStep {
	return SyntheticsAPIStep{SyntheticsAPISubtestStep: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SyntheticsAPIStep) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SyntheticsAPITestStep
	err = datadog.Unmarshal(data, &obj.SyntheticsAPITestStep)
	if err == nil {
		if obj.SyntheticsAPITestStep != nil && obj.SyntheticsAPITestStep.UnparsedObject == nil {
			jsonSyntheticsAPITestStep, _ := datadog.Marshal(obj.SyntheticsAPITestStep)
			if string(jsonSyntheticsAPITestStep) == "{}" { // empty struct
				obj.SyntheticsAPITestStep = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsAPITestStep = nil
		}
	} else {
		obj.SyntheticsAPITestStep = nil
	}

	// try to unmarshal data into SyntheticsAPIWaitStep
	err = datadog.Unmarshal(data, &obj.SyntheticsAPIWaitStep)
	if err == nil {
		if obj.SyntheticsAPIWaitStep != nil && obj.SyntheticsAPIWaitStep.UnparsedObject == nil {
			jsonSyntheticsAPIWaitStep, _ := datadog.Marshal(obj.SyntheticsAPIWaitStep)
			if string(jsonSyntheticsAPIWaitStep) == "{}" { // empty struct
				obj.SyntheticsAPIWaitStep = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsAPIWaitStep = nil
		}
	} else {
		obj.SyntheticsAPIWaitStep = nil
	}

	// try to unmarshal data into SyntheticsAPISubtestStep
	err = datadog.Unmarshal(data, &obj.SyntheticsAPISubtestStep)
	if err == nil {
		if obj.SyntheticsAPISubtestStep != nil && obj.SyntheticsAPISubtestStep.UnparsedObject == nil {
			jsonSyntheticsAPISubtestStep, _ := datadog.Marshal(obj.SyntheticsAPISubtestStep)
			if string(jsonSyntheticsAPISubtestStep) == "{}" { // empty struct
				obj.SyntheticsAPISubtestStep = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsAPISubtestStep = nil
		}
	} else {
		obj.SyntheticsAPISubtestStep = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SyntheticsAPITestStep = nil
		obj.SyntheticsAPIWaitStep = nil
		obj.SyntheticsAPISubtestStep = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SyntheticsAPIStep) MarshalJSON() ([]byte, error) {
	if obj.SyntheticsAPITestStep != nil {
		return datadog.Marshal(&obj.SyntheticsAPITestStep)
	}

	if obj.SyntheticsAPIWaitStep != nil {
		return datadog.Marshal(&obj.SyntheticsAPIWaitStep)
	}

	if obj.SyntheticsAPISubtestStep != nil {
		return datadog.Marshal(&obj.SyntheticsAPISubtestStep)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SyntheticsAPIStep) GetActualInstance() interface{} {
	if obj.SyntheticsAPITestStep != nil {
		return obj.SyntheticsAPITestStep
	}

	if obj.SyntheticsAPIWaitStep != nil {
		return obj.SyntheticsAPIWaitStep
	}

	if obj.SyntheticsAPISubtestStep != nil {
		return obj.SyntheticsAPISubtestStep
	}

	// all schemas are nil
	return nil
}
