// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionPointItem - List of distribution point.
type DistributionPointItem struct {
	DistributionPointTimestamp *float64
	DistributionPointData      *[]float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DistributionPointTimestampAsDistributionPointItem is a convenience function that returns float64 wrapped in DistributionPointItem.
func DistributionPointTimestampAsDistributionPointItem(v *float64) DistributionPointItem {
	return DistributionPointItem{DistributionPointTimestamp: v}
}

// DistributionPointDataAsDistributionPointItem is a convenience function that returns []float64 wrapped in DistributionPointItem.
func DistributionPointDataAsDistributionPointItem(v *[]float64) DistributionPointItem {
	return DistributionPointItem{DistributionPointData: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DistributionPointItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DistributionPointTimestamp
	err = datadog.Unmarshal(data, &obj.DistributionPointTimestamp)
	if err == nil {
		if obj.DistributionPointTimestamp != nil {
			jsonDistributionPointTimestamp, _ := datadog.Marshal(obj.DistributionPointTimestamp)
			if string(jsonDistributionPointTimestamp) == "{}" { // empty struct
				obj.DistributionPointTimestamp = nil
			} else {
				match++
			}
		} else {
			obj.DistributionPointTimestamp = nil
		}
	} else {
		obj.DistributionPointTimestamp = nil
	}

	// try to unmarshal data into DistributionPointData
	err = datadog.Unmarshal(data, &obj.DistributionPointData)
	if err == nil {
		if obj.DistributionPointData != nil {
			jsonDistributionPointData, _ := datadog.Marshal(obj.DistributionPointData)
			if string(jsonDistributionPointData) == "{}" && string(data) != "{}" { // empty struct
				obj.DistributionPointData = nil
			} else {
				match++
			}
		} else {
			obj.DistributionPointData = nil
		}
	} else {
		obj.DistributionPointData = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DistributionPointTimestamp = nil
		obj.DistributionPointData = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DistributionPointItem) MarshalJSON() ([]byte, error) {
	if obj.DistributionPointTimestamp != nil {
		return datadog.Marshal(&obj.DistributionPointTimestamp)
	}

	if obj.DistributionPointData != nil {
		return datadog.Marshal(&obj.DistributionPointData)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DistributionPointItem) GetActualInstance() interface{} {
	if obj.DistributionPointTimestamp != nil {
		return obj.DistributionPointTimestamp
	}

	if obj.DistributionPointData != nil {
		return obj.DistributionPointData
	}

	// all schemas are nil
	return nil
}
