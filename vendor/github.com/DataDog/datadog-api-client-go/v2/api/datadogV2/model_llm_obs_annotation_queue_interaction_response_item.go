// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueInteractionResponseItem - A single interaction result.
type LLMObsAnnotationQueueInteractionResponseItem struct {
	LLMObsTraceInteractionResponseItem        *LLMObsTraceInteractionResponseItem
	LLMObsDisplayBlockInteractionResponseItem *LLMObsDisplayBlockInteractionResponseItem

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LLMObsTraceInteractionResponseItemAsLLMObsAnnotationQueueInteractionResponseItem is a convenience function that returns LLMObsTraceInteractionResponseItem wrapped in LLMObsAnnotationQueueInteractionResponseItem.
func LLMObsTraceInteractionResponseItemAsLLMObsAnnotationQueueInteractionResponseItem(v *LLMObsTraceInteractionResponseItem) LLMObsAnnotationQueueInteractionResponseItem {
	return LLMObsAnnotationQueueInteractionResponseItem{LLMObsTraceInteractionResponseItem: v}
}

// LLMObsDisplayBlockInteractionResponseItemAsLLMObsAnnotationQueueInteractionResponseItem is a convenience function that returns LLMObsDisplayBlockInteractionResponseItem wrapped in LLMObsAnnotationQueueInteractionResponseItem.
func LLMObsDisplayBlockInteractionResponseItemAsLLMObsAnnotationQueueInteractionResponseItem(v *LLMObsDisplayBlockInteractionResponseItem) LLMObsAnnotationQueueInteractionResponseItem {
	return LLMObsAnnotationQueueInteractionResponseItem{LLMObsDisplayBlockInteractionResponseItem: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LLMObsAnnotationQueueInteractionResponseItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LLMObsTraceInteractionResponseItem
	err = datadog.Unmarshal(data, &obj.LLMObsTraceInteractionResponseItem)
	if err == nil {
		if obj.LLMObsTraceInteractionResponseItem != nil && obj.LLMObsTraceInteractionResponseItem.UnparsedObject == nil {
			jsonLLMObsTraceInteractionResponseItem, _ := datadog.Marshal(obj.LLMObsTraceInteractionResponseItem)
			if string(jsonLLMObsTraceInteractionResponseItem) == "{}" { // empty struct
				obj.LLMObsTraceInteractionResponseItem = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsTraceInteractionResponseItem = nil
		}
	} else {
		obj.LLMObsTraceInteractionResponseItem = nil
	}

	// try to unmarshal data into LLMObsDisplayBlockInteractionResponseItem
	err = datadog.Unmarshal(data, &obj.LLMObsDisplayBlockInteractionResponseItem)
	if err == nil {
		if obj.LLMObsDisplayBlockInteractionResponseItem != nil && obj.LLMObsDisplayBlockInteractionResponseItem.UnparsedObject == nil {
			jsonLLMObsDisplayBlockInteractionResponseItem, _ := datadog.Marshal(obj.LLMObsDisplayBlockInteractionResponseItem)
			if string(jsonLLMObsDisplayBlockInteractionResponseItem) == "{}" { // empty struct
				obj.LLMObsDisplayBlockInteractionResponseItem = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsDisplayBlockInteractionResponseItem = nil
		}
	} else {
		obj.LLMObsDisplayBlockInteractionResponseItem = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LLMObsTraceInteractionResponseItem = nil
		obj.LLMObsDisplayBlockInteractionResponseItem = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LLMObsAnnotationQueueInteractionResponseItem) MarshalJSON() ([]byte, error) {
	if obj.LLMObsTraceInteractionResponseItem != nil {
		return datadog.Marshal(&obj.LLMObsTraceInteractionResponseItem)
	}

	if obj.LLMObsDisplayBlockInteractionResponseItem != nil {
		return datadog.Marshal(&obj.LLMObsDisplayBlockInteractionResponseItem)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LLMObsAnnotationQueueInteractionResponseItem) GetActualInstance() interface{} {
	if obj.LLMObsTraceInteractionResponseItem != nil {
		return obj.LLMObsTraceInteractionResponseItem
	}

	if obj.LLMObsDisplayBlockInteractionResponseItem != nil {
		return obj.LLMObsDisplayBlockInteractionResponseItem
	}

	// all schemas are nil
	return nil
}
