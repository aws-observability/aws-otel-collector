// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppCreatePipelineEventRequestDataSingleOrArray - Data of the pipeline events to create.
type CIAppCreatePipelineEventRequestDataSingleOrArray struct {
	CIAppCreatePipelineEventRequestData      *CIAppCreatePipelineEventRequestData
	CIAppCreatePipelineEventRequestDataArray *[]CIAppCreatePipelineEventRequestData

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CIAppCreatePipelineEventRequestDataAsCIAppCreatePipelineEventRequestDataSingleOrArray is a convenience function that returns CIAppCreatePipelineEventRequestData wrapped in CIAppCreatePipelineEventRequestDataSingleOrArray.
func CIAppCreatePipelineEventRequestDataAsCIAppCreatePipelineEventRequestDataSingleOrArray(v *CIAppCreatePipelineEventRequestData) CIAppCreatePipelineEventRequestDataSingleOrArray {
	return CIAppCreatePipelineEventRequestDataSingleOrArray{CIAppCreatePipelineEventRequestData: v}
}

// CIAppCreatePipelineEventRequestDataArrayAsCIAppCreatePipelineEventRequestDataSingleOrArray is a convenience function that returns []CIAppCreatePipelineEventRequestData wrapped in CIAppCreatePipelineEventRequestDataSingleOrArray.
func CIAppCreatePipelineEventRequestDataArrayAsCIAppCreatePipelineEventRequestDataSingleOrArray(v *[]CIAppCreatePipelineEventRequestData) CIAppCreatePipelineEventRequestDataSingleOrArray {
	return CIAppCreatePipelineEventRequestDataSingleOrArray{CIAppCreatePipelineEventRequestDataArray: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CIAppCreatePipelineEventRequestDataSingleOrArray) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CIAppCreatePipelineEventRequestData
	err = datadog.Unmarshal(data, &obj.CIAppCreatePipelineEventRequestData)
	if err == nil {
		if obj.CIAppCreatePipelineEventRequestData != nil && obj.CIAppCreatePipelineEventRequestData.UnparsedObject == nil {
			jsonCIAppCreatePipelineEventRequestData, _ := datadog.Marshal(obj.CIAppCreatePipelineEventRequestData)
			if string(jsonCIAppCreatePipelineEventRequestData) == "{}" && string(data) != "{}" { // empty struct
				obj.CIAppCreatePipelineEventRequestData = nil
			} else {
				match++
			}
		} else {
			obj.CIAppCreatePipelineEventRequestData = nil
		}
	} else {
		obj.CIAppCreatePipelineEventRequestData = nil
	}

	// try to unmarshal data into CIAppCreatePipelineEventRequestDataArray
	err = datadog.Unmarshal(data, &obj.CIAppCreatePipelineEventRequestDataArray)
	if err == nil {
		if obj.CIAppCreatePipelineEventRequestDataArray != nil {
			jsonCIAppCreatePipelineEventRequestDataArray, _ := datadog.Marshal(obj.CIAppCreatePipelineEventRequestDataArray)
			if string(jsonCIAppCreatePipelineEventRequestDataArray) == "{}" && string(data) != "{}" { // empty struct
				obj.CIAppCreatePipelineEventRequestDataArray = nil
			} else {
				match++
			}
		} else {
			obj.CIAppCreatePipelineEventRequestDataArray = nil
		}
	} else {
		obj.CIAppCreatePipelineEventRequestDataArray = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CIAppCreatePipelineEventRequestData = nil
		obj.CIAppCreatePipelineEventRequestDataArray = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CIAppCreatePipelineEventRequestDataSingleOrArray) MarshalJSON() ([]byte, error) {
	if obj.CIAppCreatePipelineEventRequestData != nil {
		return datadog.Marshal(&obj.CIAppCreatePipelineEventRequestData)
	}

	if obj.CIAppCreatePipelineEventRequestDataArray != nil {
		return datadog.Marshal(&obj.CIAppCreatePipelineEventRequestDataArray)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CIAppCreatePipelineEventRequestDataSingleOrArray) GetActualInstance() interface{} {
	if obj.CIAppCreatePipelineEventRequestData != nil {
		return obj.CIAppCreatePipelineEventRequestData
	}

	if obj.CIAppCreatePipelineEventRequestDataArray != nil {
		return obj.CIAppCreatePipelineEventRequestDataArray
	}

	// all schemas are nil
	return nil
}
