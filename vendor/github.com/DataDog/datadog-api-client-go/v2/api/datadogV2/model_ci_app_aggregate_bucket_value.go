// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// CIAppAggregateBucketValue - A bucket value, can either be a timeseries or a single value.
type CIAppAggregateBucketValue struct {
	CIAppAggregateBucketValueSingleString *string
	CIAppAggregateBucketValueSingleNumber *float64
	CIAppAggregateBucketValueTimeseries   *CIAppAggregateBucketValueTimeseries

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CIAppAggregateBucketValueSingleStringAsCIAppAggregateBucketValue is a convenience function that returns string wrapped in CIAppAggregateBucketValue.
func CIAppAggregateBucketValueSingleStringAsCIAppAggregateBucketValue(v *string) CIAppAggregateBucketValue {
	return CIAppAggregateBucketValue{CIAppAggregateBucketValueSingleString: v}
}

// CIAppAggregateBucketValueSingleNumberAsCIAppAggregateBucketValue is a convenience function that returns float64 wrapped in CIAppAggregateBucketValue.
func CIAppAggregateBucketValueSingleNumberAsCIAppAggregateBucketValue(v *float64) CIAppAggregateBucketValue {
	return CIAppAggregateBucketValue{CIAppAggregateBucketValueSingleNumber: v}
}

// CIAppAggregateBucketValueTimeseriesAsCIAppAggregateBucketValue is a convenience function that returns CIAppAggregateBucketValueTimeseries wrapped in CIAppAggregateBucketValue.
func CIAppAggregateBucketValueTimeseriesAsCIAppAggregateBucketValue(v *CIAppAggregateBucketValueTimeseries) CIAppAggregateBucketValue {
	return CIAppAggregateBucketValue{CIAppAggregateBucketValueTimeseries: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CIAppAggregateBucketValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CIAppAggregateBucketValueSingleString
	err = json.Unmarshal(data, &obj.CIAppAggregateBucketValueSingleString)
	if err == nil {
		if obj.CIAppAggregateBucketValueSingleString != nil {
			jsonCIAppAggregateBucketValueSingleString, _ := json.Marshal(obj.CIAppAggregateBucketValueSingleString)
			if string(jsonCIAppAggregateBucketValueSingleString) == "{}" { // empty struct
				obj.CIAppAggregateBucketValueSingleString = nil
			} else {
				match++
			}
		} else {
			obj.CIAppAggregateBucketValueSingleString = nil
		}
	} else {
		obj.CIAppAggregateBucketValueSingleString = nil
	}

	// try to unmarshal data into CIAppAggregateBucketValueSingleNumber
	err = json.Unmarshal(data, &obj.CIAppAggregateBucketValueSingleNumber)
	if err == nil {
		if obj.CIAppAggregateBucketValueSingleNumber != nil {
			jsonCIAppAggregateBucketValueSingleNumber, _ := json.Marshal(obj.CIAppAggregateBucketValueSingleNumber)
			if string(jsonCIAppAggregateBucketValueSingleNumber) == "{}" { // empty struct
				obj.CIAppAggregateBucketValueSingleNumber = nil
			} else {
				match++
			}
		} else {
			obj.CIAppAggregateBucketValueSingleNumber = nil
		}
	} else {
		obj.CIAppAggregateBucketValueSingleNumber = nil
	}

	// try to unmarshal data into CIAppAggregateBucketValueTimeseries
	err = json.Unmarshal(data, &obj.CIAppAggregateBucketValueTimeseries)
	if err == nil {
		if obj.CIAppAggregateBucketValueTimeseries != nil {
			jsonCIAppAggregateBucketValueTimeseries, _ := json.Marshal(obj.CIAppAggregateBucketValueTimeseries)
			if string(jsonCIAppAggregateBucketValueTimeseries) == "{}" { // empty struct
				obj.CIAppAggregateBucketValueTimeseries = nil
			} else {
				match++
			}
		} else {
			obj.CIAppAggregateBucketValueTimeseries = nil
		}
	} else {
		obj.CIAppAggregateBucketValueTimeseries = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CIAppAggregateBucketValueSingleString = nil
		obj.CIAppAggregateBucketValueSingleNumber = nil
		obj.CIAppAggregateBucketValueTimeseries = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CIAppAggregateBucketValue) MarshalJSON() ([]byte, error) {
	if obj.CIAppAggregateBucketValueSingleString != nil {
		return json.Marshal(&obj.CIAppAggregateBucketValueSingleString)
	}

	if obj.CIAppAggregateBucketValueSingleNumber != nil {
		return json.Marshal(&obj.CIAppAggregateBucketValueSingleNumber)
	}

	if obj.CIAppAggregateBucketValueTimeseries != nil {
		return json.Marshal(&obj.CIAppAggregateBucketValueTimeseries)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CIAppAggregateBucketValue) GetActualInstance() interface{} {
	if obj.CIAppAggregateBucketValueSingleString != nil {
		return obj.CIAppAggregateBucketValueSingleString
	}

	if obj.CIAppAggregateBucketValueSingleNumber != nil {
		return obj.CIAppAggregateBucketValueSingleNumber
	}

	if obj.CIAppAggregateBucketValueTimeseries != nil {
		return obj.CIAppAggregateBucketValueTimeseries
	}

	// all schemas are nil
	return nil
}
