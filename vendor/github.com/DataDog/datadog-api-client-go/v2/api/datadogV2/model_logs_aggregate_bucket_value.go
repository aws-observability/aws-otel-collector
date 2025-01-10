// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsAggregateBucketValue - A bucket value, can be either a timeseries or a single value
type LogsAggregateBucketValue struct {
	LogsAggregateBucketValueSingleString *string
	LogsAggregateBucketValueSingleNumber *float64
	LogsAggregateBucketValueTimeseries   *LogsAggregateBucketValueTimeseries

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LogsAggregateBucketValueSingleStringAsLogsAggregateBucketValue is a convenience function that returns string wrapped in LogsAggregateBucketValue.
func LogsAggregateBucketValueSingleStringAsLogsAggregateBucketValue(v *string) LogsAggregateBucketValue {
	return LogsAggregateBucketValue{LogsAggregateBucketValueSingleString: v}
}

// LogsAggregateBucketValueSingleNumberAsLogsAggregateBucketValue is a convenience function that returns float64 wrapped in LogsAggregateBucketValue.
func LogsAggregateBucketValueSingleNumberAsLogsAggregateBucketValue(v *float64) LogsAggregateBucketValue {
	return LogsAggregateBucketValue{LogsAggregateBucketValueSingleNumber: v}
}

// LogsAggregateBucketValueTimeseriesAsLogsAggregateBucketValue is a convenience function that returns LogsAggregateBucketValueTimeseries wrapped in LogsAggregateBucketValue.
func LogsAggregateBucketValueTimeseriesAsLogsAggregateBucketValue(v *LogsAggregateBucketValueTimeseries) LogsAggregateBucketValue {
	return LogsAggregateBucketValue{LogsAggregateBucketValueTimeseries: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LogsAggregateBucketValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsAggregateBucketValueSingleString
	err = datadog.Unmarshal(data, &obj.LogsAggregateBucketValueSingleString)
	if err == nil {
		if obj.LogsAggregateBucketValueSingleString != nil {
			jsonLogsAggregateBucketValueSingleString, _ := datadog.Marshal(obj.LogsAggregateBucketValueSingleString)
			if string(jsonLogsAggregateBucketValueSingleString) == "{}" { // empty struct
				obj.LogsAggregateBucketValueSingleString = nil
			} else {
				match++
			}
		} else {
			obj.LogsAggregateBucketValueSingleString = nil
		}
	} else {
		obj.LogsAggregateBucketValueSingleString = nil
	}

	// try to unmarshal data into LogsAggregateBucketValueSingleNumber
	err = datadog.Unmarshal(data, &obj.LogsAggregateBucketValueSingleNumber)
	if err == nil {
		if obj.LogsAggregateBucketValueSingleNumber != nil {
			jsonLogsAggregateBucketValueSingleNumber, _ := datadog.Marshal(obj.LogsAggregateBucketValueSingleNumber)
			if string(jsonLogsAggregateBucketValueSingleNumber) == "{}" { // empty struct
				obj.LogsAggregateBucketValueSingleNumber = nil
			} else {
				match++
			}
		} else {
			obj.LogsAggregateBucketValueSingleNumber = nil
		}
	} else {
		obj.LogsAggregateBucketValueSingleNumber = nil
	}

	// try to unmarshal data into LogsAggregateBucketValueTimeseries
	err = datadog.Unmarshal(data, &obj.LogsAggregateBucketValueTimeseries)
	if err == nil {
		if obj.LogsAggregateBucketValueTimeseries != nil {
			jsonLogsAggregateBucketValueTimeseries, _ := datadog.Marshal(obj.LogsAggregateBucketValueTimeseries)
			if string(jsonLogsAggregateBucketValueTimeseries) == "{}" && string(data) != "{}" { // empty struct
				obj.LogsAggregateBucketValueTimeseries = nil
			} else {
				match++
			}
		} else {
			obj.LogsAggregateBucketValueTimeseries = nil
		}
	} else {
		obj.LogsAggregateBucketValueTimeseries = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LogsAggregateBucketValueSingleString = nil
		obj.LogsAggregateBucketValueSingleNumber = nil
		obj.LogsAggregateBucketValueTimeseries = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LogsAggregateBucketValue) MarshalJSON() ([]byte, error) {
	if obj.LogsAggregateBucketValueSingleString != nil {
		return datadog.Marshal(&obj.LogsAggregateBucketValueSingleString)
	}

	if obj.LogsAggregateBucketValueSingleNumber != nil {
		return datadog.Marshal(&obj.LogsAggregateBucketValueSingleNumber)
	}

	if obj.LogsAggregateBucketValueTimeseries != nil {
		return datadog.Marshal(&obj.LogsAggregateBucketValueTimeseries)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LogsAggregateBucketValue) GetActualInstance() interface{} {
	if obj.LogsAggregateBucketValueSingleString != nil {
		return obj.LogsAggregateBucketValueSingleString
	}

	if obj.LogsAggregateBucketValueSingleNumber != nil {
		return obj.LogsAggregateBucketValueSingleNumber
	}

	if obj.LogsAggregateBucketValueTimeseries != nil {
		return obj.LogsAggregateBucketValueTimeseries
	}

	// all schemas are nil
	return nil
}
