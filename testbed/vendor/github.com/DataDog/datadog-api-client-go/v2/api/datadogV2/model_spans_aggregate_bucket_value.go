// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAggregateBucketValue - A bucket value, can be either a timeseries or a single value.
type SpansAggregateBucketValue struct {
	SpansAggregateBucketValueSingleString *string
	SpansAggregateBucketValueSingleNumber *float64
	SpansAggregateBucketValueTimeseries   *SpansAggregateBucketValueTimeseries

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SpansAggregateBucketValueSingleStringAsSpansAggregateBucketValue is a convenience function that returns string wrapped in SpansAggregateBucketValue.
func SpansAggregateBucketValueSingleStringAsSpansAggregateBucketValue(v *string) SpansAggregateBucketValue {
	return SpansAggregateBucketValue{SpansAggregateBucketValueSingleString: v}
}

// SpansAggregateBucketValueSingleNumberAsSpansAggregateBucketValue is a convenience function that returns float64 wrapped in SpansAggregateBucketValue.
func SpansAggregateBucketValueSingleNumberAsSpansAggregateBucketValue(v *float64) SpansAggregateBucketValue {
	return SpansAggregateBucketValue{SpansAggregateBucketValueSingleNumber: v}
}

// SpansAggregateBucketValueTimeseriesAsSpansAggregateBucketValue is a convenience function that returns SpansAggregateBucketValueTimeseries wrapped in SpansAggregateBucketValue.
func SpansAggregateBucketValueTimeseriesAsSpansAggregateBucketValue(v *SpansAggregateBucketValueTimeseries) SpansAggregateBucketValue {
	return SpansAggregateBucketValue{SpansAggregateBucketValueTimeseries: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SpansAggregateBucketValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SpansAggregateBucketValueSingleString
	err = datadog.Unmarshal(data, &obj.SpansAggregateBucketValueSingleString)
	if err == nil {
		if obj.SpansAggregateBucketValueSingleString != nil {
			jsonSpansAggregateBucketValueSingleString, _ := datadog.Marshal(obj.SpansAggregateBucketValueSingleString)
			if string(jsonSpansAggregateBucketValueSingleString) == "{}" { // empty struct
				obj.SpansAggregateBucketValueSingleString = nil
			} else {
				match++
			}
		} else {
			obj.SpansAggregateBucketValueSingleString = nil
		}
	} else {
		obj.SpansAggregateBucketValueSingleString = nil
	}

	// try to unmarshal data into SpansAggregateBucketValueSingleNumber
	err = datadog.Unmarshal(data, &obj.SpansAggregateBucketValueSingleNumber)
	if err == nil {
		if obj.SpansAggregateBucketValueSingleNumber != nil {
			jsonSpansAggregateBucketValueSingleNumber, _ := datadog.Marshal(obj.SpansAggregateBucketValueSingleNumber)
			if string(jsonSpansAggregateBucketValueSingleNumber) == "{}" { // empty struct
				obj.SpansAggregateBucketValueSingleNumber = nil
			} else {
				match++
			}
		} else {
			obj.SpansAggregateBucketValueSingleNumber = nil
		}
	} else {
		obj.SpansAggregateBucketValueSingleNumber = nil
	}

	// try to unmarshal data into SpansAggregateBucketValueTimeseries
	err = datadog.Unmarshal(data, &obj.SpansAggregateBucketValueTimeseries)
	if err == nil {
		if obj.SpansAggregateBucketValueTimeseries != nil {
			jsonSpansAggregateBucketValueTimeseries, _ := datadog.Marshal(obj.SpansAggregateBucketValueTimeseries)
			if string(jsonSpansAggregateBucketValueTimeseries) == "{}" && string(data) != "{}" { // empty struct
				obj.SpansAggregateBucketValueTimeseries = nil
			} else {
				match++
			}
		} else {
			obj.SpansAggregateBucketValueTimeseries = nil
		}
	} else {
		obj.SpansAggregateBucketValueTimeseries = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SpansAggregateBucketValueSingleString = nil
		obj.SpansAggregateBucketValueSingleNumber = nil
		obj.SpansAggregateBucketValueTimeseries = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SpansAggregateBucketValue) MarshalJSON() ([]byte, error) {
	if obj.SpansAggregateBucketValueSingleString != nil {
		return datadog.Marshal(&obj.SpansAggregateBucketValueSingleString)
	}

	if obj.SpansAggregateBucketValueSingleNumber != nil {
		return datadog.Marshal(&obj.SpansAggregateBucketValueSingleNumber)
	}

	if obj.SpansAggregateBucketValueTimeseries != nil {
		return datadog.Marshal(&obj.SpansAggregateBucketValueTimeseries)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SpansAggregateBucketValue) GetActualInstance() interface{} {
	if obj.SpansAggregateBucketValueSingleString != nil {
		return obj.SpansAggregateBucketValueSingleString
	}

	if obj.SpansAggregateBucketValueSingleNumber != nil {
		return obj.SpansAggregateBucketValueSingleNumber
	}

	if obj.SpansAggregateBucketValueTimeseries != nil {
		return obj.SpansAggregateBucketValueTimeseries
	}

	// all schemas are nil
	return nil
}
