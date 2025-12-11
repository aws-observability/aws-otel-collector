// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSchemaMapper - Configuration of the schema processor mapper to use.
type LogsSchemaMapper struct {
	LogsSchemaRemapper       *LogsSchemaRemapper
	LogsSchemaCategoryMapper *LogsSchemaCategoryMapper

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LogsSchemaRemapperAsLogsSchemaMapper is a convenience function that returns LogsSchemaRemapper wrapped in LogsSchemaMapper.
func LogsSchemaRemapperAsLogsSchemaMapper(v *LogsSchemaRemapper) LogsSchemaMapper {
	return LogsSchemaMapper{LogsSchemaRemapper: v}
}

// LogsSchemaCategoryMapperAsLogsSchemaMapper is a convenience function that returns LogsSchemaCategoryMapper wrapped in LogsSchemaMapper.
func LogsSchemaCategoryMapperAsLogsSchemaMapper(v *LogsSchemaCategoryMapper) LogsSchemaMapper {
	return LogsSchemaMapper{LogsSchemaCategoryMapper: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LogsSchemaMapper) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsSchemaRemapper
	err = datadog.Unmarshal(data, &obj.LogsSchemaRemapper)
	if err == nil {
		if obj.LogsSchemaRemapper != nil && obj.LogsSchemaRemapper.UnparsedObject == nil {
			jsonLogsSchemaRemapper, _ := datadog.Marshal(obj.LogsSchemaRemapper)
			if string(jsonLogsSchemaRemapper) == "{}" { // empty struct
				obj.LogsSchemaRemapper = nil
			} else {
				match++
			}
		} else {
			obj.LogsSchemaRemapper = nil
		}
	} else {
		obj.LogsSchemaRemapper = nil
	}

	// try to unmarshal data into LogsSchemaCategoryMapper
	err = datadog.Unmarshal(data, &obj.LogsSchemaCategoryMapper)
	if err == nil {
		if obj.LogsSchemaCategoryMapper != nil && obj.LogsSchemaCategoryMapper.UnparsedObject == nil {
			jsonLogsSchemaCategoryMapper, _ := datadog.Marshal(obj.LogsSchemaCategoryMapper)
			if string(jsonLogsSchemaCategoryMapper) == "{}" { // empty struct
				obj.LogsSchemaCategoryMapper = nil
			} else {
				match++
			}
		} else {
			obj.LogsSchemaCategoryMapper = nil
		}
	} else {
		obj.LogsSchemaCategoryMapper = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LogsSchemaRemapper = nil
		obj.LogsSchemaCategoryMapper = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LogsSchemaMapper) MarshalJSON() ([]byte, error) {
	if obj.LogsSchemaRemapper != nil {
		return datadog.Marshal(&obj.LogsSchemaRemapper)
	}

	if obj.LogsSchemaCategoryMapper != nil {
		return datadog.Marshal(&obj.LogsSchemaCategoryMapper)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LogsSchemaMapper) GetActualInstance() interface{} {
	if obj.LogsSchemaRemapper != nil {
		return obj.LogsSchemaRemapper
	}

	if obj.LogsSchemaCategoryMapper != nil {
		return obj.LogsSchemaCategoryMapper
	}

	// all schemas are nil
	return nil
}
