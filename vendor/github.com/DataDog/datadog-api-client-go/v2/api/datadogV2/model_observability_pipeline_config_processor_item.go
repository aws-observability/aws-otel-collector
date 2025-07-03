// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigProcessorItem - A processor for the pipeline.
type ObservabilityPipelineConfigProcessorItem struct {
	ObservabilityPipelineFilterProcessor       *ObservabilityPipelineFilterProcessor
	ObservabilityPipelineParseJSONProcessor    *ObservabilityPipelineParseJSONProcessor
	ObservabilityPipelineQuotaProcessor        *ObservabilityPipelineQuotaProcessor
	ObservabilityPipelineAddFieldsProcessor    *ObservabilityPipelineAddFieldsProcessor
	ObservabilityPipelineRemoveFieldsProcessor *ObservabilityPipelineRemoveFieldsProcessor
	ObservabilityPipelineRenameFieldsProcessor *ObservabilityPipelineRenameFieldsProcessor

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineFilterProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineFilterProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineFilterProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineFilterProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineFilterProcessor: v}
}

// ObservabilityPipelineParseJSONProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineParseJSONProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineParseJSONProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineParseJSONProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineParseJSONProcessor: v}
}

// ObservabilityPipelineQuotaProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineQuotaProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineQuotaProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineQuotaProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineQuotaProcessor: v}
}

// ObservabilityPipelineAddFieldsProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineAddFieldsProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineAddFieldsProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineAddFieldsProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineAddFieldsProcessor: v}
}

// ObservabilityPipelineRemoveFieldsProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineRemoveFieldsProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineRemoveFieldsProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineRemoveFieldsProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineRemoveFieldsProcessor: v}
}

// ObservabilityPipelineRenameFieldsProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineRenameFieldsProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineRenameFieldsProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineRenameFieldsProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineRenameFieldsProcessor: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineConfigProcessorItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineFilterProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineFilterProcessor)
	if err == nil {
		if obj.ObservabilityPipelineFilterProcessor != nil && obj.ObservabilityPipelineFilterProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineFilterProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineFilterProcessor)
			if string(jsonObservabilityPipelineFilterProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineFilterProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineFilterProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineFilterProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineParseJSONProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineParseJSONProcessor)
	if err == nil {
		if obj.ObservabilityPipelineParseJSONProcessor != nil && obj.ObservabilityPipelineParseJSONProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineParseJSONProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineParseJSONProcessor)
			if string(jsonObservabilityPipelineParseJSONProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineParseJSONProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineParseJSONProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineParseJSONProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineQuotaProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineQuotaProcessor)
	if err == nil {
		if obj.ObservabilityPipelineQuotaProcessor != nil && obj.ObservabilityPipelineQuotaProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineQuotaProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineQuotaProcessor)
			if string(jsonObservabilityPipelineQuotaProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineQuotaProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineQuotaProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineQuotaProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineAddFieldsProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAddFieldsProcessor)
	if err == nil {
		if obj.ObservabilityPipelineAddFieldsProcessor != nil && obj.ObservabilityPipelineAddFieldsProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineAddFieldsProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineAddFieldsProcessor)
			if string(jsonObservabilityPipelineAddFieldsProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineAddFieldsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAddFieldsProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineAddFieldsProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineRemoveFieldsProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineRemoveFieldsProcessor)
	if err == nil {
		if obj.ObservabilityPipelineRemoveFieldsProcessor != nil && obj.ObservabilityPipelineRemoveFieldsProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineRemoveFieldsProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineRemoveFieldsProcessor)
			if string(jsonObservabilityPipelineRemoveFieldsProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineRemoveFieldsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineRemoveFieldsProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineRemoveFieldsProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineRenameFieldsProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineRenameFieldsProcessor)
	if err == nil {
		if obj.ObservabilityPipelineRenameFieldsProcessor != nil && obj.ObservabilityPipelineRenameFieldsProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineRenameFieldsProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineRenameFieldsProcessor)
			if string(jsonObservabilityPipelineRenameFieldsProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineRenameFieldsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineRenameFieldsProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineRenameFieldsProcessor = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineFilterProcessor = nil
		obj.ObservabilityPipelineParseJSONProcessor = nil
		obj.ObservabilityPipelineQuotaProcessor = nil
		obj.ObservabilityPipelineAddFieldsProcessor = nil
		obj.ObservabilityPipelineRemoveFieldsProcessor = nil
		obj.ObservabilityPipelineRenameFieldsProcessor = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineConfigProcessorItem) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineFilterProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineFilterProcessor)
	}

	if obj.ObservabilityPipelineParseJSONProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineParseJSONProcessor)
	}

	if obj.ObservabilityPipelineQuotaProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineQuotaProcessor)
	}

	if obj.ObservabilityPipelineAddFieldsProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAddFieldsProcessor)
	}

	if obj.ObservabilityPipelineRemoveFieldsProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineRemoveFieldsProcessor)
	}

	if obj.ObservabilityPipelineRenameFieldsProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineRenameFieldsProcessor)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineConfigProcessorItem) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineFilterProcessor != nil {
		return obj.ObservabilityPipelineFilterProcessor
	}

	if obj.ObservabilityPipelineParseJSONProcessor != nil {
		return obj.ObservabilityPipelineParseJSONProcessor
	}

	if obj.ObservabilityPipelineQuotaProcessor != nil {
		return obj.ObservabilityPipelineQuotaProcessor
	}

	if obj.ObservabilityPipelineAddFieldsProcessor != nil {
		return obj.ObservabilityPipelineAddFieldsProcessor
	}

	if obj.ObservabilityPipelineRemoveFieldsProcessor != nil {
		return obj.ObservabilityPipelineRemoveFieldsProcessor
	}

	if obj.ObservabilityPipelineRenameFieldsProcessor != nil {
		return obj.ObservabilityPipelineRenameFieldsProcessor
	}

	// all schemas are nil
	return nil
}
