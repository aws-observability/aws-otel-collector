// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigProcessorItem - A processor for the pipeline.
type ObservabilityPipelineConfigProcessorItem struct {
	ObservabilityPipelineFilterProcessor               *ObservabilityPipelineFilterProcessor
	ObservabilityPipelineParseJSONProcessor            *ObservabilityPipelineParseJSONProcessor
	ObservabilityPipelineQuotaProcessor                *ObservabilityPipelineQuotaProcessor
	ObservabilityPipelineAddFieldsProcessor            *ObservabilityPipelineAddFieldsProcessor
	ObservabilityPipelineRemoveFieldsProcessor         *ObservabilityPipelineRemoveFieldsProcessor
	ObservabilityPipelineRenameFieldsProcessor         *ObservabilityPipelineRenameFieldsProcessor
	ObservabilityPipelineGenerateMetricsProcessor      *ObservabilityPipelineGenerateMetricsProcessor
	ObservabilityPipelineSampleProcessor               *ObservabilityPipelineSampleProcessor
	ObservabilityPipelineParseGrokProcessor            *ObservabilityPipelineParseGrokProcessor
	ObservabilityPipelineSensitiveDataScannerProcessor *ObservabilityPipelineSensitiveDataScannerProcessor
	ObservabilityPipelineOcsfMapperProcessor           *ObservabilityPipelineOcsfMapperProcessor
	ObservabilityPipelineAddEnvVarsProcessor           *ObservabilityPipelineAddEnvVarsProcessor
	ObservabilityPipelineDedupeProcessor               *ObservabilityPipelineDedupeProcessor
	ObservabilityPipelineEnrichmentTableProcessor      *ObservabilityPipelineEnrichmentTableProcessor
	ObservabilityPipelineReduceProcessor               *ObservabilityPipelineReduceProcessor
	ObservabilityPipelineThrottleProcessor             *ObservabilityPipelineThrottleProcessor
	ObservabilityPipelineCustomProcessor               *ObservabilityPipelineCustomProcessor
	ObservabilityPipelineDatadogTagsProcessor          *ObservabilityPipelineDatadogTagsProcessor

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

// ObservabilityPipelineGenerateMetricsProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineGenerateMetricsProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineGenerateMetricsProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineGenerateMetricsProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineGenerateMetricsProcessor: v}
}

// ObservabilityPipelineSampleProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineSampleProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineSampleProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineSampleProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineSampleProcessor: v}
}

// ObservabilityPipelineParseGrokProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineParseGrokProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineParseGrokProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineParseGrokProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineParseGrokProcessor: v}
}

// ObservabilityPipelineSensitiveDataScannerProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineSensitiveDataScannerProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineSensitiveDataScannerProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineSensitiveDataScannerProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineSensitiveDataScannerProcessor: v}
}

// ObservabilityPipelineOcsfMapperProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineOcsfMapperProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineOcsfMapperProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineOcsfMapperProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineOcsfMapperProcessor: v}
}

// ObservabilityPipelineAddEnvVarsProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineAddEnvVarsProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineAddEnvVarsProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineAddEnvVarsProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineAddEnvVarsProcessor: v}
}

// ObservabilityPipelineDedupeProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineDedupeProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineDedupeProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineDedupeProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineDedupeProcessor: v}
}

// ObservabilityPipelineEnrichmentTableProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineEnrichmentTableProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineEnrichmentTableProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineEnrichmentTableProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineEnrichmentTableProcessor: v}
}

// ObservabilityPipelineReduceProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineReduceProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineReduceProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineReduceProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineReduceProcessor: v}
}

// ObservabilityPipelineThrottleProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineThrottleProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineThrottleProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineThrottleProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineThrottleProcessor: v}
}

// ObservabilityPipelineCustomProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineCustomProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineCustomProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineCustomProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineCustomProcessor: v}
}

// ObservabilityPipelineDatadogTagsProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineDatadogTagsProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineDatadogTagsProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineDatadogTagsProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineDatadogTagsProcessor: v}
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

	// try to unmarshal data into ObservabilityPipelineGenerateMetricsProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineGenerateMetricsProcessor)
	if err == nil {
		if obj.ObservabilityPipelineGenerateMetricsProcessor != nil && obj.ObservabilityPipelineGenerateMetricsProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineGenerateMetricsProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineGenerateMetricsProcessor)
			if string(jsonObservabilityPipelineGenerateMetricsProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineGenerateMetricsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineGenerateMetricsProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineGenerateMetricsProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineSampleProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSampleProcessor)
	if err == nil {
		if obj.ObservabilityPipelineSampleProcessor != nil && obj.ObservabilityPipelineSampleProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineSampleProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineSampleProcessor)
			if string(jsonObservabilityPipelineSampleProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineSampleProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSampleProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineSampleProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineParseGrokProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineParseGrokProcessor)
	if err == nil {
		if obj.ObservabilityPipelineParseGrokProcessor != nil && obj.ObservabilityPipelineParseGrokProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineParseGrokProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineParseGrokProcessor)
			if string(jsonObservabilityPipelineParseGrokProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineParseGrokProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineParseGrokProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineParseGrokProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineSensitiveDataScannerProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSensitiveDataScannerProcessor)
	if err == nil {
		if obj.ObservabilityPipelineSensitiveDataScannerProcessor != nil && obj.ObservabilityPipelineSensitiveDataScannerProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineSensitiveDataScannerProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineSensitiveDataScannerProcessor)
			if string(jsonObservabilityPipelineSensitiveDataScannerProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineSensitiveDataScannerProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSensitiveDataScannerProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineSensitiveDataScannerProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineOcsfMapperProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineOcsfMapperProcessor)
	if err == nil {
		if obj.ObservabilityPipelineOcsfMapperProcessor != nil && obj.ObservabilityPipelineOcsfMapperProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineOcsfMapperProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineOcsfMapperProcessor)
			if string(jsonObservabilityPipelineOcsfMapperProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineOcsfMapperProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineOcsfMapperProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineOcsfMapperProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineAddEnvVarsProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAddEnvVarsProcessor)
	if err == nil {
		if obj.ObservabilityPipelineAddEnvVarsProcessor != nil && obj.ObservabilityPipelineAddEnvVarsProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineAddEnvVarsProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineAddEnvVarsProcessor)
			if string(jsonObservabilityPipelineAddEnvVarsProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineAddEnvVarsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAddEnvVarsProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineAddEnvVarsProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineDedupeProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineDedupeProcessor)
	if err == nil {
		if obj.ObservabilityPipelineDedupeProcessor != nil && obj.ObservabilityPipelineDedupeProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineDedupeProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineDedupeProcessor)
			if string(jsonObservabilityPipelineDedupeProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineDedupeProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineDedupeProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineDedupeProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineEnrichmentTableProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineEnrichmentTableProcessor)
	if err == nil {
		if obj.ObservabilityPipelineEnrichmentTableProcessor != nil && obj.ObservabilityPipelineEnrichmentTableProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineEnrichmentTableProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineEnrichmentTableProcessor)
			if string(jsonObservabilityPipelineEnrichmentTableProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineEnrichmentTableProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineEnrichmentTableProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineEnrichmentTableProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineReduceProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineReduceProcessor)
	if err == nil {
		if obj.ObservabilityPipelineReduceProcessor != nil && obj.ObservabilityPipelineReduceProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineReduceProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineReduceProcessor)
			if string(jsonObservabilityPipelineReduceProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineReduceProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineReduceProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineReduceProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineThrottleProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineThrottleProcessor)
	if err == nil {
		if obj.ObservabilityPipelineThrottleProcessor != nil && obj.ObservabilityPipelineThrottleProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineThrottleProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineThrottleProcessor)
			if string(jsonObservabilityPipelineThrottleProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineThrottleProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineThrottleProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineThrottleProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineCustomProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineCustomProcessor)
	if err == nil {
		if obj.ObservabilityPipelineCustomProcessor != nil && obj.ObservabilityPipelineCustomProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineCustomProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineCustomProcessor)
			if string(jsonObservabilityPipelineCustomProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineCustomProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineCustomProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineCustomProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineDatadogTagsProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineDatadogTagsProcessor)
	if err == nil {
		if obj.ObservabilityPipelineDatadogTagsProcessor != nil && obj.ObservabilityPipelineDatadogTagsProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineDatadogTagsProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineDatadogTagsProcessor)
			if string(jsonObservabilityPipelineDatadogTagsProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineDatadogTagsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineDatadogTagsProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineDatadogTagsProcessor = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineFilterProcessor = nil
		obj.ObservabilityPipelineParseJSONProcessor = nil
		obj.ObservabilityPipelineQuotaProcessor = nil
		obj.ObservabilityPipelineAddFieldsProcessor = nil
		obj.ObservabilityPipelineRemoveFieldsProcessor = nil
		obj.ObservabilityPipelineRenameFieldsProcessor = nil
		obj.ObservabilityPipelineGenerateMetricsProcessor = nil
		obj.ObservabilityPipelineSampleProcessor = nil
		obj.ObservabilityPipelineParseGrokProcessor = nil
		obj.ObservabilityPipelineSensitiveDataScannerProcessor = nil
		obj.ObservabilityPipelineOcsfMapperProcessor = nil
		obj.ObservabilityPipelineAddEnvVarsProcessor = nil
		obj.ObservabilityPipelineDedupeProcessor = nil
		obj.ObservabilityPipelineEnrichmentTableProcessor = nil
		obj.ObservabilityPipelineReduceProcessor = nil
		obj.ObservabilityPipelineThrottleProcessor = nil
		obj.ObservabilityPipelineCustomProcessor = nil
		obj.ObservabilityPipelineDatadogTagsProcessor = nil
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

	if obj.ObservabilityPipelineGenerateMetricsProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineGenerateMetricsProcessor)
	}

	if obj.ObservabilityPipelineSampleProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSampleProcessor)
	}

	if obj.ObservabilityPipelineParseGrokProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineParseGrokProcessor)
	}

	if obj.ObservabilityPipelineSensitiveDataScannerProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSensitiveDataScannerProcessor)
	}

	if obj.ObservabilityPipelineOcsfMapperProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineOcsfMapperProcessor)
	}

	if obj.ObservabilityPipelineAddEnvVarsProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAddEnvVarsProcessor)
	}

	if obj.ObservabilityPipelineDedupeProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineDedupeProcessor)
	}

	if obj.ObservabilityPipelineEnrichmentTableProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineEnrichmentTableProcessor)
	}

	if obj.ObservabilityPipelineReduceProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineReduceProcessor)
	}

	if obj.ObservabilityPipelineThrottleProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineThrottleProcessor)
	}

	if obj.ObservabilityPipelineCustomProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineCustomProcessor)
	}

	if obj.ObservabilityPipelineDatadogTagsProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineDatadogTagsProcessor)
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

	if obj.ObservabilityPipelineGenerateMetricsProcessor != nil {
		return obj.ObservabilityPipelineGenerateMetricsProcessor
	}

	if obj.ObservabilityPipelineSampleProcessor != nil {
		return obj.ObservabilityPipelineSampleProcessor
	}

	if obj.ObservabilityPipelineParseGrokProcessor != nil {
		return obj.ObservabilityPipelineParseGrokProcessor
	}

	if obj.ObservabilityPipelineSensitiveDataScannerProcessor != nil {
		return obj.ObservabilityPipelineSensitiveDataScannerProcessor
	}

	if obj.ObservabilityPipelineOcsfMapperProcessor != nil {
		return obj.ObservabilityPipelineOcsfMapperProcessor
	}

	if obj.ObservabilityPipelineAddEnvVarsProcessor != nil {
		return obj.ObservabilityPipelineAddEnvVarsProcessor
	}

	if obj.ObservabilityPipelineDedupeProcessor != nil {
		return obj.ObservabilityPipelineDedupeProcessor
	}

	if obj.ObservabilityPipelineEnrichmentTableProcessor != nil {
		return obj.ObservabilityPipelineEnrichmentTableProcessor
	}

	if obj.ObservabilityPipelineReduceProcessor != nil {
		return obj.ObservabilityPipelineReduceProcessor
	}

	if obj.ObservabilityPipelineThrottleProcessor != nil {
		return obj.ObservabilityPipelineThrottleProcessor
	}

	if obj.ObservabilityPipelineCustomProcessor != nil {
		return obj.ObservabilityPipelineCustomProcessor
	}

	if obj.ObservabilityPipelineDatadogTagsProcessor != nil {
		return obj.ObservabilityPipelineDatadogTagsProcessor
	}

	// all schemas are nil
	return nil
}
