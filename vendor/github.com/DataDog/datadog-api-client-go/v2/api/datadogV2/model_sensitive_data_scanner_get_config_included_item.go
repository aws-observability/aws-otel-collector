// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerGetConfigIncludedItem - An object related to the configuration.
type SensitiveDataScannerGetConfigIncludedItem struct {
	SensitiveDataScannerRuleIncludedItem  *SensitiveDataScannerRuleIncludedItem
	SensitiveDataScannerGroupIncludedItem *SensitiveDataScannerGroupIncludedItem

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SensitiveDataScannerRuleIncludedItemAsSensitiveDataScannerGetConfigIncludedItem is a convenience function that returns SensitiveDataScannerRuleIncludedItem wrapped in SensitiveDataScannerGetConfigIncludedItem.
func SensitiveDataScannerRuleIncludedItemAsSensitiveDataScannerGetConfigIncludedItem(v *SensitiveDataScannerRuleIncludedItem) SensitiveDataScannerGetConfigIncludedItem {
	return SensitiveDataScannerGetConfigIncludedItem{SensitiveDataScannerRuleIncludedItem: v}
}

// SensitiveDataScannerGroupIncludedItemAsSensitiveDataScannerGetConfigIncludedItem is a convenience function that returns SensitiveDataScannerGroupIncludedItem wrapped in SensitiveDataScannerGetConfigIncludedItem.
func SensitiveDataScannerGroupIncludedItemAsSensitiveDataScannerGetConfigIncludedItem(v *SensitiveDataScannerGroupIncludedItem) SensitiveDataScannerGetConfigIncludedItem {
	return SensitiveDataScannerGetConfigIncludedItem{SensitiveDataScannerGroupIncludedItem: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SensitiveDataScannerGetConfigIncludedItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SensitiveDataScannerRuleIncludedItem
	err = datadog.Unmarshal(data, &obj.SensitiveDataScannerRuleIncludedItem)
	if err == nil {
		if obj.SensitiveDataScannerRuleIncludedItem != nil && obj.SensitiveDataScannerRuleIncludedItem.UnparsedObject == nil {
			jsonSensitiveDataScannerRuleIncludedItem, _ := datadog.Marshal(obj.SensitiveDataScannerRuleIncludedItem)
			if string(jsonSensitiveDataScannerRuleIncludedItem) == "{}" && string(data) != "{}" { // empty struct
				obj.SensitiveDataScannerRuleIncludedItem = nil
			} else {
				match++
			}
		} else {
			obj.SensitiveDataScannerRuleIncludedItem = nil
		}
	} else {
		obj.SensitiveDataScannerRuleIncludedItem = nil
	}

	// try to unmarshal data into SensitiveDataScannerGroupIncludedItem
	err = datadog.Unmarshal(data, &obj.SensitiveDataScannerGroupIncludedItem)
	if err == nil {
		if obj.SensitiveDataScannerGroupIncludedItem != nil && obj.SensitiveDataScannerGroupIncludedItem.UnparsedObject == nil {
			jsonSensitiveDataScannerGroupIncludedItem, _ := datadog.Marshal(obj.SensitiveDataScannerGroupIncludedItem)
			if string(jsonSensitiveDataScannerGroupIncludedItem) == "{}" && string(data) != "{}" { // empty struct
				obj.SensitiveDataScannerGroupIncludedItem = nil
			} else {
				match++
			}
		} else {
			obj.SensitiveDataScannerGroupIncludedItem = nil
		}
	} else {
		obj.SensitiveDataScannerGroupIncludedItem = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SensitiveDataScannerRuleIncludedItem = nil
		obj.SensitiveDataScannerGroupIncludedItem = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SensitiveDataScannerGetConfigIncludedItem) MarshalJSON() ([]byte, error) {
	if obj.SensitiveDataScannerRuleIncludedItem != nil {
		return datadog.Marshal(&obj.SensitiveDataScannerRuleIncludedItem)
	}

	if obj.SensitiveDataScannerGroupIncludedItem != nil {
		return datadog.Marshal(&obj.SensitiveDataScannerGroupIncludedItem)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SensitiveDataScannerGetConfigIncludedItem) GetActualInstance() interface{} {
	if obj.SensitiveDataScannerRuleIncludedItem != nil {
		return obj.SensitiveDataScannerRuleIncludedItem
	}

	if obj.SensitiveDataScannerGroupIncludedItem != nil {
		return obj.SensitiveDataScannerGroupIncludedItem
	}

	// all schemas are nil
	return nil
}
