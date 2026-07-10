// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayMapSubProcessor - A sub-processor used inside an array-map processor.
// Allowed types: `attribute-remapper`, `string-builder-processor`,
// `arithmetic-processor`, `category-processor`.
type LogsArrayMapSubProcessor struct {
	LogsArrayMapAttributeRemapper         *LogsArrayMapAttributeRemapper
	LogsArrayMapArithmeticSubProcessor    *LogsArrayMapArithmeticSubProcessor
	LogsArrayMapStringBuilderSubProcessor *LogsArrayMapStringBuilderSubProcessor
	LogsArrayMapCategorySubProcessor      *LogsArrayMapCategorySubProcessor

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LogsArrayMapAttributeRemapperAsLogsArrayMapSubProcessor is a convenience function that returns LogsArrayMapAttributeRemapper wrapped in LogsArrayMapSubProcessor.
func LogsArrayMapAttributeRemapperAsLogsArrayMapSubProcessor(v *LogsArrayMapAttributeRemapper) LogsArrayMapSubProcessor {
	return LogsArrayMapSubProcessor{LogsArrayMapAttributeRemapper: v}
}

// LogsArrayMapArithmeticSubProcessorAsLogsArrayMapSubProcessor is a convenience function that returns LogsArrayMapArithmeticSubProcessor wrapped in LogsArrayMapSubProcessor.
func LogsArrayMapArithmeticSubProcessorAsLogsArrayMapSubProcessor(v *LogsArrayMapArithmeticSubProcessor) LogsArrayMapSubProcessor {
	return LogsArrayMapSubProcessor{LogsArrayMapArithmeticSubProcessor: v}
}

// LogsArrayMapStringBuilderSubProcessorAsLogsArrayMapSubProcessor is a convenience function that returns LogsArrayMapStringBuilderSubProcessor wrapped in LogsArrayMapSubProcessor.
func LogsArrayMapStringBuilderSubProcessorAsLogsArrayMapSubProcessor(v *LogsArrayMapStringBuilderSubProcessor) LogsArrayMapSubProcessor {
	return LogsArrayMapSubProcessor{LogsArrayMapStringBuilderSubProcessor: v}
}

// LogsArrayMapCategorySubProcessorAsLogsArrayMapSubProcessor is a convenience function that returns LogsArrayMapCategorySubProcessor wrapped in LogsArrayMapSubProcessor.
func LogsArrayMapCategorySubProcessorAsLogsArrayMapSubProcessor(v *LogsArrayMapCategorySubProcessor) LogsArrayMapSubProcessor {
	return LogsArrayMapSubProcessor{LogsArrayMapCategorySubProcessor: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LogsArrayMapSubProcessor) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsArrayMapAttributeRemapper
	err = datadog.Unmarshal(data, &obj.LogsArrayMapAttributeRemapper)
	if err == nil {
		if obj.LogsArrayMapAttributeRemapper != nil && obj.LogsArrayMapAttributeRemapper.UnparsedObject == nil {
			jsonLogsArrayMapAttributeRemapper, _ := datadog.Marshal(obj.LogsArrayMapAttributeRemapper)
			if string(jsonLogsArrayMapAttributeRemapper) == "{}" { // empty struct
				obj.LogsArrayMapAttributeRemapper = nil
			} else {
				match++
			}
		} else {
			obj.LogsArrayMapAttributeRemapper = nil
		}
	} else {
		obj.LogsArrayMapAttributeRemapper = nil
	}

	// try to unmarshal data into LogsArrayMapArithmeticSubProcessor
	err = datadog.Unmarshal(data, &obj.LogsArrayMapArithmeticSubProcessor)
	if err == nil {
		if obj.LogsArrayMapArithmeticSubProcessor != nil && obj.LogsArrayMapArithmeticSubProcessor.UnparsedObject == nil {
			jsonLogsArrayMapArithmeticSubProcessor, _ := datadog.Marshal(obj.LogsArrayMapArithmeticSubProcessor)
			if string(jsonLogsArrayMapArithmeticSubProcessor) == "{}" { // empty struct
				obj.LogsArrayMapArithmeticSubProcessor = nil
			} else {
				match++
			}
		} else {
			obj.LogsArrayMapArithmeticSubProcessor = nil
		}
	} else {
		obj.LogsArrayMapArithmeticSubProcessor = nil
	}

	// try to unmarshal data into LogsArrayMapStringBuilderSubProcessor
	err = datadog.Unmarshal(data, &obj.LogsArrayMapStringBuilderSubProcessor)
	if err == nil {
		if obj.LogsArrayMapStringBuilderSubProcessor != nil && obj.LogsArrayMapStringBuilderSubProcessor.UnparsedObject == nil {
			jsonLogsArrayMapStringBuilderSubProcessor, _ := datadog.Marshal(obj.LogsArrayMapStringBuilderSubProcessor)
			if string(jsonLogsArrayMapStringBuilderSubProcessor) == "{}" { // empty struct
				obj.LogsArrayMapStringBuilderSubProcessor = nil
			} else {
				match++
			}
		} else {
			obj.LogsArrayMapStringBuilderSubProcessor = nil
		}
	} else {
		obj.LogsArrayMapStringBuilderSubProcessor = nil
	}

	// try to unmarshal data into LogsArrayMapCategorySubProcessor
	err = datadog.Unmarshal(data, &obj.LogsArrayMapCategorySubProcessor)
	if err == nil {
		if obj.LogsArrayMapCategorySubProcessor != nil && obj.LogsArrayMapCategorySubProcessor.UnparsedObject == nil {
			jsonLogsArrayMapCategorySubProcessor, _ := datadog.Marshal(obj.LogsArrayMapCategorySubProcessor)
			if string(jsonLogsArrayMapCategorySubProcessor) == "{}" { // empty struct
				obj.LogsArrayMapCategorySubProcessor = nil
			} else {
				match++
			}
		} else {
			obj.LogsArrayMapCategorySubProcessor = nil
		}
	} else {
		obj.LogsArrayMapCategorySubProcessor = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LogsArrayMapAttributeRemapper = nil
		obj.LogsArrayMapArithmeticSubProcessor = nil
		obj.LogsArrayMapStringBuilderSubProcessor = nil
		obj.LogsArrayMapCategorySubProcessor = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LogsArrayMapSubProcessor) MarshalJSON() ([]byte, error) {
	if obj.LogsArrayMapAttributeRemapper != nil {
		return datadog.Marshal(&obj.LogsArrayMapAttributeRemapper)
	}

	if obj.LogsArrayMapArithmeticSubProcessor != nil {
		return datadog.Marshal(&obj.LogsArrayMapArithmeticSubProcessor)
	}

	if obj.LogsArrayMapStringBuilderSubProcessor != nil {
		return datadog.Marshal(&obj.LogsArrayMapStringBuilderSubProcessor)
	}

	if obj.LogsArrayMapCategorySubProcessor != nil {
		return datadog.Marshal(&obj.LogsArrayMapCategorySubProcessor)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LogsArrayMapSubProcessor) GetActualInstance() interface{} {
	if obj.LogsArrayMapAttributeRemapper != nil {
		return obj.LogsArrayMapAttributeRemapper
	}

	if obj.LogsArrayMapArithmeticSubProcessor != nil {
		return obj.LogsArrayMapArithmeticSubProcessor
	}

	if obj.LogsArrayMapStringBuilderSubProcessor != nil {
		return obj.LogsArrayMapStringBuilderSubProcessor
	}

	if obj.LogsArrayMapCategorySubProcessor != nil {
		return obj.LogsArrayMapCategorySubProcessor
	}

	// all schemas are nil
	return nil
}
