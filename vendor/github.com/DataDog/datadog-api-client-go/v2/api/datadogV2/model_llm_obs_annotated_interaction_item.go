// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotatedInteractionItem - An interaction with its associated annotations.
type LLMObsAnnotatedInteractionItem struct {
	LLMObsTraceAnnotatedInteractionItem        *LLMObsTraceAnnotatedInteractionItem
	LLMObsDisplayBlockAnnotatedInteractionItem *LLMObsDisplayBlockAnnotatedInteractionItem

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LLMObsTraceAnnotatedInteractionItemAsLLMObsAnnotatedInteractionItem is a convenience function that returns LLMObsTraceAnnotatedInteractionItem wrapped in LLMObsAnnotatedInteractionItem.
func LLMObsTraceAnnotatedInteractionItemAsLLMObsAnnotatedInteractionItem(v *LLMObsTraceAnnotatedInteractionItem) LLMObsAnnotatedInteractionItem {
	return LLMObsAnnotatedInteractionItem{LLMObsTraceAnnotatedInteractionItem: v}
}

// LLMObsDisplayBlockAnnotatedInteractionItemAsLLMObsAnnotatedInteractionItem is a convenience function that returns LLMObsDisplayBlockAnnotatedInteractionItem wrapped in LLMObsAnnotatedInteractionItem.
func LLMObsDisplayBlockAnnotatedInteractionItemAsLLMObsAnnotatedInteractionItem(v *LLMObsDisplayBlockAnnotatedInteractionItem) LLMObsAnnotatedInteractionItem {
	return LLMObsAnnotatedInteractionItem{LLMObsDisplayBlockAnnotatedInteractionItem: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LLMObsAnnotatedInteractionItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LLMObsTraceAnnotatedInteractionItem
	err = datadog.Unmarshal(data, &obj.LLMObsTraceAnnotatedInteractionItem)
	if err == nil {
		if obj.LLMObsTraceAnnotatedInteractionItem != nil && obj.LLMObsTraceAnnotatedInteractionItem.UnparsedObject == nil {
			jsonLLMObsTraceAnnotatedInteractionItem, _ := datadog.Marshal(obj.LLMObsTraceAnnotatedInteractionItem)
			if string(jsonLLMObsTraceAnnotatedInteractionItem) == "{}" { // empty struct
				obj.LLMObsTraceAnnotatedInteractionItem = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsTraceAnnotatedInteractionItem = nil
		}
	} else {
		obj.LLMObsTraceAnnotatedInteractionItem = nil
	}

	// try to unmarshal data into LLMObsDisplayBlockAnnotatedInteractionItem
	err = datadog.Unmarshal(data, &obj.LLMObsDisplayBlockAnnotatedInteractionItem)
	if err == nil {
		if obj.LLMObsDisplayBlockAnnotatedInteractionItem != nil && obj.LLMObsDisplayBlockAnnotatedInteractionItem.UnparsedObject == nil {
			jsonLLMObsDisplayBlockAnnotatedInteractionItem, _ := datadog.Marshal(obj.LLMObsDisplayBlockAnnotatedInteractionItem)
			if string(jsonLLMObsDisplayBlockAnnotatedInteractionItem) == "{}" { // empty struct
				obj.LLMObsDisplayBlockAnnotatedInteractionItem = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsDisplayBlockAnnotatedInteractionItem = nil
		}
	} else {
		obj.LLMObsDisplayBlockAnnotatedInteractionItem = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LLMObsTraceAnnotatedInteractionItem = nil
		obj.LLMObsDisplayBlockAnnotatedInteractionItem = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LLMObsAnnotatedInteractionItem) MarshalJSON() ([]byte, error) {
	if obj.LLMObsTraceAnnotatedInteractionItem != nil {
		return datadog.Marshal(&obj.LLMObsTraceAnnotatedInteractionItem)
	}

	if obj.LLMObsDisplayBlockAnnotatedInteractionItem != nil {
		return datadog.Marshal(&obj.LLMObsDisplayBlockAnnotatedInteractionItem)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LLMObsAnnotatedInteractionItem) GetActualInstance() interface{} {
	if obj.LLMObsTraceAnnotatedInteractionItem != nil {
		return obj.LLMObsTraceAnnotatedInteractionItem
	}

	if obj.LLMObsDisplayBlockAnnotatedInteractionItem != nil {
		return obj.LLMObsDisplayBlockAnnotatedInteractionItem
	}

	// all schemas are nil
	return nil
}
