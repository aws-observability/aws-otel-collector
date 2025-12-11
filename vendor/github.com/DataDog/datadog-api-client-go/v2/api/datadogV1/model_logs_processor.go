// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsProcessor - Definition of a logs processor.
type LogsProcessor struct {
	LogsGrokParser                    *LogsGrokParser
	LogsDateRemapper                  *LogsDateRemapper
	LogsStatusRemapper                *LogsStatusRemapper
	LogsServiceRemapper               *LogsServiceRemapper
	LogsMessageRemapper               *LogsMessageRemapper
	LogsAttributeRemapper             *LogsAttributeRemapper
	LogsURLParser                     *LogsURLParser
	LogsUserAgentParser               *LogsUserAgentParser
	LogsCategoryProcessor             *LogsCategoryProcessor
	LogsArithmeticProcessor           *LogsArithmeticProcessor
	LogsStringBuilderProcessor        *LogsStringBuilderProcessor
	LogsPipelineProcessor             *LogsPipelineProcessor
	LogsGeoIPParser                   *LogsGeoIPParser
	LogsLookupProcessor               *LogsLookupProcessor
	ReferenceTableLogsLookupProcessor *ReferenceTableLogsLookupProcessor
	LogsTraceRemapper                 *LogsTraceRemapper
	LogsSpanRemapper                  *LogsSpanRemapper
	LogsArrayProcessor                *LogsArrayProcessor
	LogsDecoderProcessor              *LogsDecoderProcessor
	LogsSchemaProcessor               *LogsSchemaProcessor

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LogsGrokParserAsLogsProcessor is a convenience function that returns LogsGrokParser wrapped in LogsProcessor.
func LogsGrokParserAsLogsProcessor(v *LogsGrokParser) LogsProcessor {
	return LogsProcessor{LogsGrokParser: v}
}

// LogsDateRemapperAsLogsProcessor is a convenience function that returns LogsDateRemapper wrapped in LogsProcessor.
func LogsDateRemapperAsLogsProcessor(v *LogsDateRemapper) LogsProcessor {
	return LogsProcessor{LogsDateRemapper: v}
}

// LogsStatusRemapperAsLogsProcessor is a convenience function that returns LogsStatusRemapper wrapped in LogsProcessor.
func LogsStatusRemapperAsLogsProcessor(v *LogsStatusRemapper) LogsProcessor {
	return LogsProcessor{LogsStatusRemapper: v}
}

// LogsServiceRemapperAsLogsProcessor is a convenience function that returns LogsServiceRemapper wrapped in LogsProcessor.
func LogsServiceRemapperAsLogsProcessor(v *LogsServiceRemapper) LogsProcessor {
	return LogsProcessor{LogsServiceRemapper: v}
}

// LogsMessageRemapperAsLogsProcessor is a convenience function that returns LogsMessageRemapper wrapped in LogsProcessor.
func LogsMessageRemapperAsLogsProcessor(v *LogsMessageRemapper) LogsProcessor {
	return LogsProcessor{LogsMessageRemapper: v}
}

// LogsAttributeRemapperAsLogsProcessor is a convenience function that returns LogsAttributeRemapper wrapped in LogsProcessor.
func LogsAttributeRemapperAsLogsProcessor(v *LogsAttributeRemapper) LogsProcessor {
	return LogsProcessor{LogsAttributeRemapper: v}
}

// LogsURLParserAsLogsProcessor is a convenience function that returns LogsURLParser wrapped in LogsProcessor.
func LogsURLParserAsLogsProcessor(v *LogsURLParser) LogsProcessor {
	return LogsProcessor{LogsURLParser: v}
}

// LogsUserAgentParserAsLogsProcessor is a convenience function that returns LogsUserAgentParser wrapped in LogsProcessor.
func LogsUserAgentParserAsLogsProcessor(v *LogsUserAgentParser) LogsProcessor {
	return LogsProcessor{LogsUserAgentParser: v}
}

// LogsCategoryProcessorAsLogsProcessor is a convenience function that returns LogsCategoryProcessor wrapped in LogsProcessor.
func LogsCategoryProcessorAsLogsProcessor(v *LogsCategoryProcessor) LogsProcessor {
	return LogsProcessor{LogsCategoryProcessor: v}
}

// LogsArithmeticProcessorAsLogsProcessor is a convenience function that returns LogsArithmeticProcessor wrapped in LogsProcessor.
func LogsArithmeticProcessorAsLogsProcessor(v *LogsArithmeticProcessor) LogsProcessor {
	return LogsProcessor{LogsArithmeticProcessor: v}
}

// LogsStringBuilderProcessorAsLogsProcessor is a convenience function that returns LogsStringBuilderProcessor wrapped in LogsProcessor.
func LogsStringBuilderProcessorAsLogsProcessor(v *LogsStringBuilderProcessor) LogsProcessor {
	return LogsProcessor{LogsStringBuilderProcessor: v}
}

// LogsPipelineProcessorAsLogsProcessor is a convenience function that returns LogsPipelineProcessor wrapped in LogsProcessor.
func LogsPipelineProcessorAsLogsProcessor(v *LogsPipelineProcessor) LogsProcessor {
	return LogsProcessor{LogsPipelineProcessor: v}
}

// LogsGeoIPParserAsLogsProcessor is a convenience function that returns LogsGeoIPParser wrapped in LogsProcessor.
func LogsGeoIPParserAsLogsProcessor(v *LogsGeoIPParser) LogsProcessor {
	return LogsProcessor{LogsGeoIPParser: v}
}

// LogsLookupProcessorAsLogsProcessor is a convenience function that returns LogsLookupProcessor wrapped in LogsProcessor.
func LogsLookupProcessorAsLogsProcessor(v *LogsLookupProcessor) LogsProcessor {
	return LogsProcessor{LogsLookupProcessor: v}
}

// ReferenceTableLogsLookupProcessorAsLogsProcessor is a convenience function that returns ReferenceTableLogsLookupProcessor wrapped in LogsProcessor.
func ReferenceTableLogsLookupProcessorAsLogsProcessor(v *ReferenceTableLogsLookupProcessor) LogsProcessor {
	return LogsProcessor{ReferenceTableLogsLookupProcessor: v}
}

// LogsTraceRemapperAsLogsProcessor is a convenience function that returns LogsTraceRemapper wrapped in LogsProcessor.
func LogsTraceRemapperAsLogsProcessor(v *LogsTraceRemapper) LogsProcessor {
	return LogsProcessor{LogsTraceRemapper: v}
}

// LogsSpanRemapperAsLogsProcessor is a convenience function that returns LogsSpanRemapper wrapped in LogsProcessor.
func LogsSpanRemapperAsLogsProcessor(v *LogsSpanRemapper) LogsProcessor {
	return LogsProcessor{LogsSpanRemapper: v}
}

// LogsArrayProcessorAsLogsProcessor is a convenience function that returns LogsArrayProcessor wrapped in LogsProcessor.
func LogsArrayProcessorAsLogsProcessor(v *LogsArrayProcessor) LogsProcessor {
	return LogsProcessor{LogsArrayProcessor: v}
}

// LogsDecoderProcessorAsLogsProcessor is a convenience function that returns LogsDecoderProcessor wrapped in LogsProcessor.
func LogsDecoderProcessorAsLogsProcessor(v *LogsDecoderProcessor) LogsProcessor {
	return LogsProcessor{LogsDecoderProcessor: v}
}

// LogsSchemaProcessorAsLogsProcessor is a convenience function that returns LogsSchemaProcessor wrapped in LogsProcessor.
func LogsSchemaProcessorAsLogsProcessor(v *LogsSchemaProcessor) LogsProcessor {
	return LogsProcessor{LogsSchemaProcessor: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LogsProcessor) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsGrokParser
	err = datadog.Unmarshal(data, &obj.LogsGrokParser)
	if err == nil {
		if obj.LogsGrokParser != nil && obj.LogsGrokParser.UnparsedObject == nil {
			jsonLogsGrokParser, _ := datadog.Marshal(obj.LogsGrokParser)
			if string(jsonLogsGrokParser) == "{}" { // empty struct
				obj.LogsGrokParser = nil
			} else {
				match++
			}
		} else {
			obj.LogsGrokParser = nil
		}
	} else {
		obj.LogsGrokParser = nil
	}

	// try to unmarshal data into LogsDateRemapper
	err = datadog.Unmarshal(data, &obj.LogsDateRemapper)
	if err == nil {
		if obj.LogsDateRemapper != nil && obj.LogsDateRemapper.UnparsedObject == nil {
			jsonLogsDateRemapper, _ := datadog.Marshal(obj.LogsDateRemapper)
			if string(jsonLogsDateRemapper) == "{}" { // empty struct
				obj.LogsDateRemapper = nil
			} else {
				match++
			}
		} else {
			obj.LogsDateRemapper = nil
		}
	} else {
		obj.LogsDateRemapper = nil
	}

	// try to unmarshal data into LogsStatusRemapper
	err = datadog.Unmarshal(data, &obj.LogsStatusRemapper)
	if err == nil {
		if obj.LogsStatusRemapper != nil && obj.LogsStatusRemapper.UnparsedObject == nil {
			jsonLogsStatusRemapper, _ := datadog.Marshal(obj.LogsStatusRemapper)
			if string(jsonLogsStatusRemapper) == "{}" { // empty struct
				obj.LogsStatusRemapper = nil
			} else {
				match++
			}
		} else {
			obj.LogsStatusRemapper = nil
		}
	} else {
		obj.LogsStatusRemapper = nil
	}

	// try to unmarshal data into LogsServiceRemapper
	err = datadog.Unmarshal(data, &obj.LogsServiceRemapper)
	if err == nil {
		if obj.LogsServiceRemapper != nil && obj.LogsServiceRemapper.UnparsedObject == nil {
			jsonLogsServiceRemapper, _ := datadog.Marshal(obj.LogsServiceRemapper)
			if string(jsonLogsServiceRemapper) == "{}" { // empty struct
				obj.LogsServiceRemapper = nil
			} else {
				match++
			}
		} else {
			obj.LogsServiceRemapper = nil
		}
	} else {
		obj.LogsServiceRemapper = nil
	}

	// try to unmarshal data into LogsMessageRemapper
	err = datadog.Unmarshal(data, &obj.LogsMessageRemapper)
	if err == nil {
		if obj.LogsMessageRemapper != nil && obj.LogsMessageRemapper.UnparsedObject == nil {
			jsonLogsMessageRemapper, _ := datadog.Marshal(obj.LogsMessageRemapper)
			if string(jsonLogsMessageRemapper) == "{}" { // empty struct
				obj.LogsMessageRemapper = nil
			} else {
				match++
			}
		} else {
			obj.LogsMessageRemapper = nil
		}
	} else {
		obj.LogsMessageRemapper = nil
	}

	// try to unmarshal data into LogsAttributeRemapper
	err = datadog.Unmarshal(data, &obj.LogsAttributeRemapper)
	if err == nil {
		if obj.LogsAttributeRemapper != nil && obj.LogsAttributeRemapper.UnparsedObject == nil {
			jsonLogsAttributeRemapper, _ := datadog.Marshal(obj.LogsAttributeRemapper)
			if string(jsonLogsAttributeRemapper) == "{}" { // empty struct
				obj.LogsAttributeRemapper = nil
			} else {
				match++
			}
		} else {
			obj.LogsAttributeRemapper = nil
		}
	} else {
		obj.LogsAttributeRemapper = nil
	}

	// try to unmarshal data into LogsURLParser
	err = datadog.Unmarshal(data, &obj.LogsURLParser)
	if err == nil {
		if obj.LogsURLParser != nil && obj.LogsURLParser.UnparsedObject == nil {
			jsonLogsURLParser, _ := datadog.Marshal(obj.LogsURLParser)
			if string(jsonLogsURLParser) == "{}" { // empty struct
				obj.LogsURLParser = nil
			} else {
				match++
			}
		} else {
			obj.LogsURLParser = nil
		}
	} else {
		obj.LogsURLParser = nil
	}

	// try to unmarshal data into LogsUserAgentParser
	err = datadog.Unmarshal(data, &obj.LogsUserAgentParser)
	if err == nil {
		if obj.LogsUserAgentParser != nil && obj.LogsUserAgentParser.UnparsedObject == nil {
			jsonLogsUserAgentParser, _ := datadog.Marshal(obj.LogsUserAgentParser)
			if string(jsonLogsUserAgentParser) == "{}" { // empty struct
				obj.LogsUserAgentParser = nil
			} else {
				match++
			}
		} else {
			obj.LogsUserAgentParser = nil
		}
	} else {
		obj.LogsUserAgentParser = nil
	}

	// try to unmarshal data into LogsCategoryProcessor
	err = datadog.Unmarshal(data, &obj.LogsCategoryProcessor)
	if err == nil {
		if obj.LogsCategoryProcessor != nil && obj.LogsCategoryProcessor.UnparsedObject == nil {
			jsonLogsCategoryProcessor, _ := datadog.Marshal(obj.LogsCategoryProcessor)
			if string(jsonLogsCategoryProcessor) == "{}" { // empty struct
				obj.LogsCategoryProcessor = nil
			} else {
				match++
			}
		} else {
			obj.LogsCategoryProcessor = nil
		}
	} else {
		obj.LogsCategoryProcessor = nil
	}

	// try to unmarshal data into LogsArithmeticProcessor
	err = datadog.Unmarshal(data, &obj.LogsArithmeticProcessor)
	if err == nil {
		if obj.LogsArithmeticProcessor != nil && obj.LogsArithmeticProcessor.UnparsedObject == nil {
			jsonLogsArithmeticProcessor, _ := datadog.Marshal(obj.LogsArithmeticProcessor)
			if string(jsonLogsArithmeticProcessor) == "{}" { // empty struct
				obj.LogsArithmeticProcessor = nil
			} else {
				match++
			}
		} else {
			obj.LogsArithmeticProcessor = nil
		}
	} else {
		obj.LogsArithmeticProcessor = nil
	}

	// try to unmarshal data into LogsStringBuilderProcessor
	err = datadog.Unmarshal(data, &obj.LogsStringBuilderProcessor)
	if err == nil {
		if obj.LogsStringBuilderProcessor != nil && obj.LogsStringBuilderProcessor.UnparsedObject == nil {
			jsonLogsStringBuilderProcessor, _ := datadog.Marshal(obj.LogsStringBuilderProcessor)
			if string(jsonLogsStringBuilderProcessor) == "{}" { // empty struct
				obj.LogsStringBuilderProcessor = nil
			} else {
				match++
			}
		} else {
			obj.LogsStringBuilderProcessor = nil
		}
	} else {
		obj.LogsStringBuilderProcessor = nil
	}

	// try to unmarshal data into LogsPipelineProcessor
	err = datadog.Unmarshal(data, &obj.LogsPipelineProcessor)
	if err == nil {
		if obj.LogsPipelineProcessor != nil && obj.LogsPipelineProcessor.UnparsedObject == nil {
			jsonLogsPipelineProcessor, _ := datadog.Marshal(obj.LogsPipelineProcessor)
			if string(jsonLogsPipelineProcessor) == "{}" { // empty struct
				obj.LogsPipelineProcessor = nil
			} else {
				match++
			}
		} else {
			obj.LogsPipelineProcessor = nil
		}
	} else {
		obj.LogsPipelineProcessor = nil
	}

	// try to unmarshal data into LogsGeoIPParser
	err = datadog.Unmarshal(data, &obj.LogsGeoIPParser)
	if err == nil {
		if obj.LogsGeoIPParser != nil && obj.LogsGeoIPParser.UnparsedObject == nil {
			jsonLogsGeoIPParser, _ := datadog.Marshal(obj.LogsGeoIPParser)
			if string(jsonLogsGeoIPParser) == "{}" { // empty struct
				obj.LogsGeoIPParser = nil
			} else {
				match++
			}
		} else {
			obj.LogsGeoIPParser = nil
		}
	} else {
		obj.LogsGeoIPParser = nil
	}

	// try to unmarshal data into LogsLookupProcessor
	err = datadog.Unmarshal(data, &obj.LogsLookupProcessor)
	if err == nil {
		if obj.LogsLookupProcessor != nil && obj.LogsLookupProcessor.UnparsedObject == nil {
			jsonLogsLookupProcessor, _ := datadog.Marshal(obj.LogsLookupProcessor)
			if string(jsonLogsLookupProcessor) == "{}" { // empty struct
				obj.LogsLookupProcessor = nil
			} else {
				match++
			}
		} else {
			obj.LogsLookupProcessor = nil
		}
	} else {
		obj.LogsLookupProcessor = nil
	}

	// try to unmarshal data into ReferenceTableLogsLookupProcessor
	err = datadog.Unmarshal(data, &obj.ReferenceTableLogsLookupProcessor)
	if err == nil {
		if obj.ReferenceTableLogsLookupProcessor != nil && obj.ReferenceTableLogsLookupProcessor.UnparsedObject == nil {
			jsonReferenceTableLogsLookupProcessor, _ := datadog.Marshal(obj.ReferenceTableLogsLookupProcessor)
			if string(jsonReferenceTableLogsLookupProcessor) == "{}" { // empty struct
				obj.ReferenceTableLogsLookupProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ReferenceTableLogsLookupProcessor = nil
		}
	} else {
		obj.ReferenceTableLogsLookupProcessor = nil
	}

	// try to unmarshal data into LogsTraceRemapper
	err = datadog.Unmarshal(data, &obj.LogsTraceRemapper)
	if err == nil {
		if obj.LogsTraceRemapper != nil && obj.LogsTraceRemapper.UnparsedObject == nil {
			jsonLogsTraceRemapper, _ := datadog.Marshal(obj.LogsTraceRemapper)
			if string(jsonLogsTraceRemapper) == "{}" { // empty struct
				obj.LogsTraceRemapper = nil
			} else {
				match++
			}
		} else {
			obj.LogsTraceRemapper = nil
		}
	} else {
		obj.LogsTraceRemapper = nil
	}

	// try to unmarshal data into LogsSpanRemapper
	err = datadog.Unmarshal(data, &obj.LogsSpanRemapper)
	if err == nil {
		if obj.LogsSpanRemapper != nil && obj.LogsSpanRemapper.UnparsedObject == nil {
			jsonLogsSpanRemapper, _ := datadog.Marshal(obj.LogsSpanRemapper)
			if string(jsonLogsSpanRemapper) == "{}" { // empty struct
				obj.LogsSpanRemapper = nil
			} else {
				match++
			}
		} else {
			obj.LogsSpanRemapper = nil
		}
	} else {
		obj.LogsSpanRemapper = nil
	}

	// try to unmarshal data into LogsArrayProcessor
	err = datadog.Unmarshal(data, &obj.LogsArrayProcessor)
	if err == nil {
		if obj.LogsArrayProcessor != nil && obj.LogsArrayProcessor.UnparsedObject == nil {
			jsonLogsArrayProcessor, _ := datadog.Marshal(obj.LogsArrayProcessor)
			if string(jsonLogsArrayProcessor) == "{}" { // empty struct
				obj.LogsArrayProcessor = nil
			} else {
				match++
			}
		} else {
			obj.LogsArrayProcessor = nil
		}
	} else {
		obj.LogsArrayProcessor = nil
	}

	// try to unmarshal data into LogsDecoderProcessor
	err = datadog.Unmarshal(data, &obj.LogsDecoderProcessor)
	if err == nil {
		if obj.LogsDecoderProcessor != nil && obj.LogsDecoderProcessor.UnparsedObject == nil {
			jsonLogsDecoderProcessor, _ := datadog.Marshal(obj.LogsDecoderProcessor)
			if string(jsonLogsDecoderProcessor) == "{}" { // empty struct
				obj.LogsDecoderProcessor = nil
			} else {
				match++
			}
		} else {
			obj.LogsDecoderProcessor = nil
		}
	} else {
		obj.LogsDecoderProcessor = nil
	}

	// try to unmarshal data into LogsSchemaProcessor
	err = datadog.Unmarshal(data, &obj.LogsSchemaProcessor)
	if err == nil {
		if obj.LogsSchemaProcessor != nil && obj.LogsSchemaProcessor.UnparsedObject == nil {
			jsonLogsSchemaProcessor, _ := datadog.Marshal(obj.LogsSchemaProcessor)
			if string(jsonLogsSchemaProcessor) == "{}" { // empty struct
				obj.LogsSchemaProcessor = nil
			} else {
				match++
			}
		} else {
			obj.LogsSchemaProcessor = nil
		}
	} else {
		obj.LogsSchemaProcessor = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LogsGrokParser = nil
		obj.LogsDateRemapper = nil
		obj.LogsStatusRemapper = nil
		obj.LogsServiceRemapper = nil
		obj.LogsMessageRemapper = nil
		obj.LogsAttributeRemapper = nil
		obj.LogsURLParser = nil
		obj.LogsUserAgentParser = nil
		obj.LogsCategoryProcessor = nil
		obj.LogsArithmeticProcessor = nil
		obj.LogsStringBuilderProcessor = nil
		obj.LogsPipelineProcessor = nil
		obj.LogsGeoIPParser = nil
		obj.LogsLookupProcessor = nil
		obj.ReferenceTableLogsLookupProcessor = nil
		obj.LogsTraceRemapper = nil
		obj.LogsSpanRemapper = nil
		obj.LogsArrayProcessor = nil
		obj.LogsDecoderProcessor = nil
		obj.LogsSchemaProcessor = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LogsProcessor) MarshalJSON() ([]byte, error) {
	if obj.LogsGrokParser != nil {
		return datadog.Marshal(&obj.LogsGrokParser)
	}

	if obj.LogsDateRemapper != nil {
		return datadog.Marshal(&obj.LogsDateRemapper)
	}

	if obj.LogsStatusRemapper != nil {
		return datadog.Marshal(&obj.LogsStatusRemapper)
	}

	if obj.LogsServiceRemapper != nil {
		return datadog.Marshal(&obj.LogsServiceRemapper)
	}

	if obj.LogsMessageRemapper != nil {
		return datadog.Marshal(&obj.LogsMessageRemapper)
	}

	if obj.LogsAttributeRemapper != nil {
		return datadog.Marshal(&obj.LogsAttributeRemapper)
	}

	if obj.LogsURLParser != nil {
		return datadog.Marshal(&obj.LogsURLParser)
	}

	if obj.LogsUserAgentParser != nil {
		return datadog.Marshal(&obj.LogsUserAgentParser)
	}

	if obj.LogsCategoryProcessor != nil {
		return datadog.Marshal(&obj.LogsCategoryProcessor)
	}

	if obj.LogsArithmeticProcessor != nil {
		return datadog.Marshal(&obj.LogsArithmeticProcessor)
	}

	if obj.LogsStringBuilderProcessor != nil {
		return datadog.Marshal(&obj.LogsStringBuilderProcessor)
	}

	if obj.LogsPipelineProcessor != nil {
		return datadog.Marshal(&obj.LogsPipelineProcessor)
	}

	if obj.LogsGeoIPParser != nil {
		return datadog.Marshal(&obj.LogsGeoIPParser)
	}

	if obj.LogsLookupProcessor != nil {
		return datadog.Marshal(&obj.LogsLookupProcessor)
	}

	if obj.ReferenceTableLogsLookupProcessor != nil {
		return datadog.Marshal(&obj.ReferenceTableLogsLookupProcessor)
	}

	if obj.LogsTraceRemapper != nil {
		return datadog.Marshal(&obj.LogsTraceRemapper)
	}

	if obj.LogsSpanRemapper != nil {
		return datadog.Marshal(&obj.LogsSpanRemapper)
	}

	if obj.LogsArrayProcessor != nil {
		return datadog.Marshal(&obj.LogsArrayProcessor)
	}

	if obj.LogsDecoderProcessor != nil {
		return datadog.Marshal(&obj.LogsDecoderProcessor)
	}

	if obj.LogsSchemaProcessor != nil {
		return datadog.Marshal(&obj.LogsSchemaProcessor)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LogsProcessor) GetActualInstance() interface{} {
	if obj.LogsGrokParser != nil {
		return obj.LogsGrokParser
	}

	if obj.LogsDateRemapper != nil {
		return obj.LogsDateRemapper
	}

	if obj.LogsStatusRemapper != nil {
		return obj.LogsStatusRemapper
	}

	if obj.LogsServiceRemapper != nil {
		return obj.LogsServiceRemapper
	}

	if obj.LogsMessageRemapper != nil {
		return obj.LogsMessageRemapper
	}

	if obj.LogsAttributeRemapper != nil {
		return obj.LogsAttributeRemapper
	}

	if obj.LogsURLParser != nil {
		return obj.LogsURLParser
	}

	if obj.LogsUserAgentParser != nil {
		return obj.LogsUserAgentParser
	}

	if obj.LogsCategoryProcessor != nil {
		return obj.LogsCategoryProcessor
	}

	if obj.LogsArithmeticProcessor != nil {
		return obj.LogsArithmeticProcessor
	}

	if obj.LogsStringBuilderProcessor != nil {
		return obj.LogsStringBuilderProcessor
	}

	if obj.LogsPipelineProcessor != nil {
		return obj.LogsPipelineProcessor
	}

	if obj.LogsGeoIPParser != nil {
		return obj.LogsGeoIPParser
	}

	if obj.LogsLookupProcessor != nil {
		return obj.LogsLookupProcessor
	}

	if obj.ReferenceTableLogsLookupProcessor != nil {
		return obj.ReferenceTableLogsLookupProcessor
	}

	if obj.LogsTraceRemapper != nil {
		return obj.LogsTraceRemapper
	}

	if obj.LogsSpanRemapper != nil {
		return obj.LogsSpanRemapper
	}

	if obj.LogsArrayProcessor != nil {
		return obj.LogsArrayProcessor
	}

	if obj.LogsDecoderProcessor != nil {
		return obj.LogsDecoderProcessor
	}

	if obj.LogsSchemaProcessor != nil {
		return obj.LogsSchemaProcessor
	}

	// all schemas are nil
	return nil
}
