// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayProcessorOperation - Configuration of the array processor operation to perform.
type LogsArrayProcessorOperation struct {
	LogsArrayProcessorOperationAppend *LogsArrayProcessorOperationAppend
	LogsArrayProcessorOperationLength *LogsArrayProcessorOperationLength
	LogsArrayProcessorOperationSelect *LogsArrayProcessorOperationSelect

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LogsArrayProcessorOperationAppendAsLogsArrayProcessorOperation is a convenience function that returns LogsArrayProcessorOperationAppend wrapped in LogsArrayProcessorOperation.
func LogsArrayProcessorOperationAppendAsLogsArrayProcessorOperation(v *LogsArrayProcessorOperationAppend) LogsArrayProcessorOperation {
	return LogsArrayProcessorOperation{LogsArrayProcessorOperationAppend: v}
}

// LogsArrayProcessorOperationLengthAsLogsArrayProcessorOperation is a convenience function that returns LogsArrayProcessorOperationLength wrapped in LogsArrayProcessorOperation.
func LogsArrayProcessorOperationLengthAsLogsArrayProcessorOperation(v *LogsArrayProcessorOperationLength) LogsArrayProcessorOperation {
	return LogsArrayProcessorOperation{LogsArrayProcessorOperationLength: v}
}

// LogsArrayProcessorOperationSelectAsLogsArrayProcessorOperation is a convenience function that returns LogsArrayProcessorOperationSelect wrapped in LogsArrayProcessorOperation.
func LogsArrayProcessorOperationSelectAsLogsArrayProcessorOperation(v *LogsArrayProcessorOperationSelect) LogsArrayProcessorOperation {
	return LogsArrayProcessorOperation{LogsArrayProcessorOperationSelect: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LogsArrayProcessorOperation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsArrayProcessorOperationAppend
	err = datadog.Unmarshal(data, &obj.LogsArrayProcessorOperationAppend)
	if err == nil {
		if obj.LogsArrayProcessorOperationAppend != nil && obj.LogsArrayProcessorOperationAppend.UnparsedObject == nil {
			jsonLogsArrayProcessorOperationAppend, _ := datadog.Marshal(obj.LogsArrayProcessorOperationAppend)
			if string(jsonLogsArrayProcessorOperationAppend) == "{}" { // empty struct
				obj.LogsArrayProcessorOperationAppend = nil
			} else {
				match++
			}
		} else {
			obj.LogsArrayProcessorOperationAppend = nil
		}
	} else {
		obj.LogsArrayProcessorOperationAppend = nil
	}

	// try to unmarshal data into LogsArrayProcessorOperationLength
	err = datadog.Unmarshal(data, &obj.LogsArrayProcessorOperationLength)
	if err == nil {
		if obj.LogsArrayProcessorOperationLength != nil && obj.LogsArrayProcessorOperationLength.UnparsedObject == nil {
			jsonLogsArrayProcessorOperationLength, _ := datadog.Marshal(obj.LogsArrayProcessorOperationLength)
			if string(jsonLogsArrayProcessorOperationLength) == "{}" { // empty struct
				obj.LogsArrayProcessorOperationLength = nil
			} else {
				match++
			}
		} else {
			obj.LogsArrayProcessorOperationLength = nil
		}
	} else {
		obj.LogsArrayProcessorOperationLength = nil
	}

	// try to unmarshal data into LogsArrayProcessorOperationSelect
	err = datadog.Unmarshal(data, &obj.LogsArrayProcessorOperationSelect)
	if err == nil {
		if obj.LogsArrayProcessorOperationSelect != nil && obj.LogsArrayProcessorOperationSelect.UnparsedObject == nil {
			jsonLogsArrayProcessorOperationSelect, _ := datadog.Marshal(obj.LogsArrayProcessorOperationSelect)
			if string(jsonLogsArrayProcessorOperationSelect) == "{}" { // empty struct
				obj.LogsArrayProcessorOperationSelect = nil
			} else {
				match++
			}
		} else {
			obj.LogsArrayProcessorOperationSelect = nil
		}
	} else {
		obj.LogsArrayProcessorOperationSelect = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LogsArrayProcessorOperationAppend = nil
		obj.LogsArrayProcessorOperationLength = nil
		obj.LogsArrayProcessorOperationSelect = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LogsArrayProcessorOperation) MarshalJSON() ([]byte, error) {
	if obj.LogsArrayProcessorOperationAppend != nil {
		return datadog.Marshal(&obj.LogsArrayProcessorOperationAppend)
	}

	if obj.LogsArrayProcessorOperationLength != nil {
		return datadog.Marshal(&obj.LogsArrayProcessorOperationLength)
	}

	if obj.LogsArrayProcessorOperationSelect != nil {
		return datadog.Marshal(&obj.LogsArrayProcessorOperationSelect)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LogsArrayProcessorOperation) GetActualInstance() interface{} {
	if obj.LogsArrayProcessorOperationAppend != nil {
		return obj.LogsArrayProcessorOperationAppend
	}

	if obj.LogsArrayProcessorOperationLength != nil {
		return obj.LogsArrayProcessorOperationLength
	}

	if obj.LogsArrayProcessorOperationSelect != nil {
		return obj.LogsArrayProcessorOperationSelect
	}

	// all schemas are nil
	return nil
}
