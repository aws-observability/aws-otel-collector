// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigSourceItem - A data source for the pipeline.
type ObservabilityPipelineConfigSourceItem struct {
	ObservabilityPipelineKafkaSource              *ObservabilityPipelineKafkaSource
	ObservabilityPipelineDatadogAgentSource       *ObservabilityPipelineDatadogAgentSource
	ObservabilityPipelineSplunkTcpSource          *ObservabilityPipelineSplunkTcpSource
	ObservabilityPipelineSplunkHecSource          *ObservabilityPipelineSplunkHecSource
	ObservabilityPipelineAmazonS3Source           *ObservabilityPipelineAmazonS3Source
	ObservabilityPipelineFluentdSource            *ObservabilityPipelineFluentdSource
	ObservabilityPipelineFluentBitSource          *ObservabilityPipelineFluentBitSource
	ObservabilityPipelineHttpServerSource         *ObservabilityPipelineHttpServerSource
	ObservabilityPipelineSumoLogicSource          *ObservabilityPipelineSumoLogicSource
	ObservabilityPipelineRsyslogSource            *ObservabilityPipelineRsyslogSource
	ObservabilityPipelineSyslogNgSource           *ObservabilityPipelineSyslogNgSource
	ObservabilityPipelineAmazonDataFirehoseSource *ObservabilityPipelineAmazonDataFirehoseSource
	ObservabilityPipelineGooglePubSubSource       *ObservabilityPipelineGooglePubSubSource
	ObservabilityPipelineHttpClientSource         *ObservabilityPipelineHttpClientSource
	ObservabilityPipelineLogstashSource           *ObservabilityPipelineLogstashSource
	ObservabilityPipelineSocketSource             *ObservabilityPipelineSocketSource

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineKafkaSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineKafkaSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineKafkaSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineKafkaSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineKafkaSource: v}
}

// ObservabilityPipelineDatadogAgentSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineDatadogAgentSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineDatadogAgentSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineDatadogAgentSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineDatadogAgentSource: v}
}

// ObservabilityPipelineSplunkTcpSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineSplunkTcpSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineSplunkTcpSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineSplunkTcpSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineSplunkTcpSource: v}
}

// ObservabilityPipelineSplunkHecSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineSplunkHecSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineSplunkHecSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineSplunkHecSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineSplunkHecSource: v}
}

// ObservabilityPipelineAmazonS3SourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineAmazonS3Source wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineAmazonS3SourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineAmazonS3Source) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineAmazonS3Source: v}
}

// ObservabilityPipelineFluentdSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineFluentdSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineFluentdSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineFluentdSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineFluentdSource: v}
}

// ObservabilityPipelineFluentBitSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineFluentBitSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineFluentBitSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineFluentBitSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineFluentBitSource: v}
}

// ObservabilityPipelineHttpServerSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineHttpServerSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineHttpServerSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineHttpServerSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineHttpServerSource: v}
}

// ObservabilityPipelineSumoLogicSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineSumoLogicSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineSumoLogicSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineSumoLogicSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineSumoLogicSource: v}
}

// ObservabilityPipelineRsyslogSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineRsyslogSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineRsyslogSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineRsyslogSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineRsyslogSource: v}
}

// ObservabilityPipelineSyslogNgSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineSyslogNgSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineSyslogNgSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineSyslogNgSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineSyslogNgSource: v}
}

// ObservabilityPipelineAmazonDataFirehoseSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineAmazonDataFirehoseSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineAmazonDataFirehoseSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineAmazonDataFirehoseSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineAmazonDataFirehoseSource: v}
}

// ObservabilityPipelineGooglePubSubSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineGooglePubSubSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineGooglePubSubSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineGooglePubSubSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineGooglePubSubSource: v}
}

// ObservabilityPipelineHttpClientSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineHttpClientSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineHttpClientSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineHttpClientSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineHttpClientSource: v}
}

// ObservabilityPipelineLogstashSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineLogstashSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineLogstashSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineLogstashSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineLogstashSource: v}
}

// ObservabilityPipelineSocketSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineSocketSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineSocketSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineSocketSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineSocketSource: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineConfigSourceItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineKafkaSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineKafkaSource)
	if err == nil {
		if obj.ObservabilityPipelineKafkaSource != nil && obj.ObservabilityPipelineKafkaSource.UnparsedObject == nil {
			jsonObservabilityPipelineKafkaSource, _ := datadog.Marshal(obj.ObservabilityPipelineKafkaSource)
			if string(jsonObservabilityPipelineKafkaSource) == "{}" { // empty struct
				obj.ObservabilityPipelineKafkaSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineKafkaSource = nil
		}
	} else {
		obj.ObservabilityPipelineKafkaSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineDatadogAgentSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineDatadogAgentSource)
	if err == nil {
		if obj.ObservabilityPipelineDatadogAgentSource != nil && obj.ObservabilityPipelineDatadogAgentSource.UnparsedObject == nil {
			jsonObservabilityPipelineDatadogAgentSource, _ := datadog.Marshal(obj.ObservabilityPipelineDatadogAgentSource)
			if string(jsonObservabilityPipelineDatadogAgentSource) == "{}" { // empty struct
				obj.ObservabilityPipelineDatadogAgentSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineDatadogAgentSource = nil
		}
	} else {
		obj.ObservabilityPipelineDatadogAgentSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineSplunkTcpSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSplunkTcpSource)
	if err == nil {
		if obj.ObservabilityPipelineSplunkTcpSource != nil && obj.ObservabilityPipelineSplunkTcpSource.UnparsedObject == nil {
			jsonObservabilityPipelineSplunkTcpSource, _ := datadog.Marshal(obj.ObservabilityPipelineSplunkTcpSource)
			if string(jsonObservabilityPipelineSplunkTcpSource) == "{}" { // empty struct
				obj.ObservabilityPipelineSplunkTcpSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSplunkTcpSource = nil
		}
	} else {
		obj.ObservabilityPipelineSplunkTcpSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineSplunkHecSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSplunkHecSource)
	if err == nil {
		if obj.ObservabilityPipelineSplunkHecSource != nil && obj.ObservabilityPipelineSplunkHecSource.UnparsedObject == nil {
			jsonObservabilityPipelineSplunkHecSource, _ := datadog.Marshal(obj.ObservabilityPipelineSplunkHecSource)
			if string(jsonObservabilityPipelineSplunkHecSource) == "{}" { // empty struct
				obj.ObservabilityPipelineSplunkHecSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSplunkHecSource = nil
		}
	} else {
		obj.ObservabilityPipelineSplunkHecSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineAmazonS3Source
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonS3Source)
	if err == nil {
		if obj.ObservabilityPipelineAmazonS3Source != nil && obj.ObservabilityPipelineAmazonS3Source.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonS3Source, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonS3Source)
			if string(jsonObservabilityPipelineAmazonS3Source) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonS3Source = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonS3Source = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonS3Source = nil
	}

	// try to unmarshal data into ObservabilityPipelineFluentdSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineFluentdSource)
	if err == nil {
		if obj.ObservabilityPipelineFluentdSource != nil && obj.ObservabilityPipelineFluentdSource.UnparsedObject == nil {
			jsonObservabilityPipelineFluentdSource, _ := datadog.Marshal(obj.ObservabilityPipelineFluentdSource)
			if string(jsonObservabilityPipelineFluentdSource) == "{}" { // empty struct
				obj.ObservabilityPipelineFluentdSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineFluentdSource = nil
		}
	} else {
		obj.ObservabilityPipelineFluentdSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineFluentBitSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineFluentBitSource)
	if err == nil {
		if obj.ObservabilityPipelineFluentBitSource != nil && obj.ObservabilityPipelineFluentBitSource.UnparsedObject == nil {
			jsonObservabilityPipelineFluentBitSource, _ := datadog.Marshal(obj.ObservabilityPipelineFluentBitSource)
			if string(jsonObservabilityPipelineFluentBitSource) == "{}" { // empty struct
				obj.ObservabilityPipelineFluentBitSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineFluentBitSource = nil
		}
	} else {
		obj.ObservabilityPipelineFluentBitSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineHttpServerSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineHttpServerSource)
	if err == nil {
		if obj.ObservabilityPipelineHttpServerSource != nil && obj.ObservabilityPipelineHttpServerSource.UnparsedObject == nil {
			jsonObservabilityPipelineHttpServerSource, _ := datadog.Marshal(obj.ObservabilityPipelineHttpServerSource)
			if string(jsonObservabilityPipelineHttpServerSource) == "{}" { // empty struct
				obj.ObservabilityPipelineHttpServerSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineHttpServerSource = nil
		}
	} else {
		obj.ObservabilityPipelineHttpServerSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineSumoLogicSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSumoLogicSource)
	if err == nil {
		if obj.ObservabilityPipelineSumoLogicSource != nil && obj.ObservabilityPipelineSumoLogicSource.UnparsedObject == nil {
			jsonObservabilityPipelineSumoLogicSource, _ := datadog.Marshal(obj.ObservabilityPipelineSumoLogicSource)
			if string(jsonObservabilityPipelineSumoLogicSource) == "{}" { // empty struct
				obj.ObservabilityPipelineSumoLogicSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSumoLogicSource = nil
		}
	} else {
		obj.ObservabilityPipelineSumoLogicSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineRsyslogSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineRsyslogSource)
	if err == nil {
		if obj.ObservabilityPipelineRsyslogSource != nil && obj.ObservabilityPipelineRsyslogSource.UnparsedObject == nil {
			jsonObservabilityPipelineRsyslogSource, _ := datadog.Marshal(obj.ObservabilityPipelineRsyslogSource)
			if string(jsonObservabilityPipelineRsyslogSource) == "{}" { // empty struct
				obj.ObservabilityPipelineRsyslogSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineRsyslogSource = nil
		}
	} else {
		obj.ObservabilityPipelineRsyslogSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineSyslogNgSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSyslogNgSource)
	if err == nil {
		if obj.ObservabilityPipelineSyslogNgSource != nil && obj.ObservabilityPipelineSyslogNgSource.UnparsedObject == nil {
			jsonObservabilityPipelineSyslogNgSource, _ := datadog.Marshal(obj.ObservabilityPipelineSyslogNgSource)
			if string(jsonObservabilityPipelineSyslogNgSource) == "{}" { // empty struct
				obj.ObservabilityPipelineSyslogNgSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSyslogNgSource = nil
		}
	} else {
		obj.ObservabilityPipelineSyslogNgSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineAmazonDataFirehoseSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonDataFirehoseSource)
	if err == nil {
		if obj.ObservabilityPipelineAmazonDataFirehoseSource != nil && obj.ObservabilityPipelineAmazonDataFirehoseSource.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonDataFirehoseSource, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonDataFirehoseSource)
			if string(jsonObservabilityPipelineAmazonDataFirehoseSource) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonDataFirehoseSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonDataFirehoseSource = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonDataFirehoseSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineGooglePubSubSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineGooglePubSubSource)
	if err == nil {
		if obj.ObservabilityPipelineGooglePubSubSource != nil && obj.ObservabilityPipelineGooglePubSubSource.UnparsedObject == nil {
			jsonObservabilityPipelineGooglePubSubSource, _ := datadog.Marshal(obj.ObservabilityPipelineGooglePubSubSource)
			if string(jsonObservabilityPipelineGooglePubSubSource) == "{}" { // empty struct
				obj.ObservabilityPipelineGooglePubSubSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineGooglePubSubSource = nil
		}
	} else {
		obj.ObservabilityPipelineGooglePubSubSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineHttpClientSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineHttpClientSource)
	if err == nil {
		if obj.ObservabilityPipelineHttpClientSource != nil && obj.ObservabilityPipelineHttpClientSource.UnparsedObject == nil {
			jsonObservabilityPipelineHttpClientSource, _ := datadog.Marshal(obj.ObservabilityPipelineHttpClientSource)
			if string(jsonObservabilityPipelineHttpClientSource) == "{}" { // empty struct
				obj.ObservabilityPipelineHttpClientSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineHttpClientSource = nil
		}
	} else {
		obj.ObservabilityPipelineHttpClientSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineLogstashSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineLogstashSource)
	if err == nil {
		if obj.ObservabilityPipelineLogstashSource != nil && obj.ObservabilityPipelineLogstashSource.UnparsedObject == nil {
			jsonObservabilityPipelineLogstashSource, _ := datadog.Marshal(obj.ObservabilityPipelineLogstashSource)
			if string(jsonObservabilityPipelineLogstashSource) == "{}" { // empty struct
				obj.ObservabilityPipelineLogstashSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineLogstashSource = nil
		}
	} else {
		obj.ObservabilityPipelineLogstashSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineSocketSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSocketSource)
	if err == nil {
		if obj.ObservabilityPipelineSocketSource != nil && obj.ObservabilityPipelineSocketSource.UnparsedObject == nil {
			jsonObservabilityPipelineSocketSource, _ := datadog.Marshal(obj.ObservabilityPipelineSocketSource)
			if string(jsonObservabilityPipelineSocketSource) == "{}" { // empty struct
				obj.ObservabilityPipelineSocketSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSocketSource = nil
		}
	} else {
		obj.ObservabilityPipelineSocketSource = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineKafkaSource = nil
		obj.ObservabilityPipelineDatadogAgentSource = nil
		obj.ObservabilityPipelineSplunkTcpSource = nil
		obj.ObservabilityPipelineSplunkHecSource = nil
		obj.ObservabilityPipelineAmazonS3Source = nil
		obj.ObservabilityPipelineFluentdSource = nil
		obj.ObservabilityPipelineFluentBitSource = nil
		obj.ObservabilityPipelineHttpServerSource = nil
		obj.ObservabilityPipelineSumoLogicSource = nil
		obj.ObservabilityPipelineRsyslogSource = nil
		obj.ObservabilityPipelineSyslogNgSource = nil
		obj.ObservabilityPipelineAmazonDataFirehoseSource = nil
		obj.ObservabilityPipelineGooglePubSubSource = nil
		obj.ObservabilityPipelineHttpClientSource = nil
		obj.ObservabilityPipelineLogstashSource = nil
		obj.ObservabilityPipelineSocketSource = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineConfigSourceItem) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineKafkaSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineKafkaSource)
	}

	if obj.ObservabilityPipelineDatadogAgentSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineDatadogAgentSource)
	}

	if obj.ObservabilityPipelineSplunkTcpSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSplunkTcpSource)
	}

	if obj.ObservabilityPipelineSplunkHecSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSplunkHecSource)
	}

	if obj.ObservabilityPipelineAmazonS3Source != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonS3Source)
	}

	if obj.ObservabilityPipelineFluentdSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineFluentdSource)
	}

	if obj.ObservabilityPipelineFluentBitSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineFluentBitSource)
	}

	if obj.ObservabilityPipelineHttpServerSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineHttpServerSource)
	}

	if obj.ObservabilityPipelineSumoLogicSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSumoLogicSource)
	}

	if obj.ObservabilityPipelineRsyslogSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineRsyslogSource)
	}

	if obj.ObservabilityPipelineSyslogNgSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSyslogNgSource)
	}

	if obj.ObservabilityPipelineAmazonDataFirehoseSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonDataFirehoseSource)
	}

	if obj.ObservabilityPipelineGooglePubSubSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineGooglePubSubSource)
	}

	if obj.ObservabilityPipelineHttpClientSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineHttpClientSource)
	}

	if obj.ObservabilityPipelineLogstashSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineLogstashSource)
	}

	if obj.ObservabilityPipelineSocketSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSocketSource)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineConfigSourceItem) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineKafkaSource != nil {
		return obj.ObservabilityPipelineKafkaSource
	}

	if obj.ObservabilityPipelineDatadogAgentSource != nil {
		return obj.ObservabilityPipelineDatadogAgentSource
	}

	if obj.ObservabilityPipelineSplunkTcpSource != nil {
		return obj.ObservabilityPipelineSplunkTcpSource
	}

	if obj.ObservabilityPipelineSplunkHecSource != nil {
		return obj.ObservabilityPipelineSplunkHecSource
	}

	if obj.ObservabilityPipelineAmazonS3Source != nil {
		return obj.ObservabilityPipelineAmazonS3Source
	}

	if obj.ObservabilityPipelineFluentdSource != nil {
		return obj.ObservabilityPipelineFluentdSource
	}

	if obj.ObservabilityPipelineFluentBitSource != nil {
		return obj.ObservabilityPipelineFluentBitSource
	}

	if obj.ObservabilityPipelineHttpServerSource != nil {
		return obj.ObservabilityPipelineHttpServerSource
	}

	if obj.ObservabilityPipelineSumoLogicSource != nil {
		return obj.ObservabilityPipelineSumoLogicSource
	}

	if obj.ObservabilityPipelineRsyslogSource != nil {
		return obj.ObservabilityPipelineRsyslogSource
	}

	if obj.ObservabilityPipelineSyslogNgSource != nil {
		return obj.ObservabilityPipelineSyslogNgSource
	}

	if obj.ObservabilityPipelineAmazonDataFirehoseSource != nil {
		return obj.ObservabilityPipelineAmazonDataFirehoseSource
	}

	if obj.ObservabilityPipelineGooglePubSubSource != nil {
		return obj.ObservabilityPipelineGooglePubSubSource
	}

	if obj.ObservabilityPipelineHttpClientSource != nil {
		return obj.ObservabilityPipelineHttpClientSource
	}

	if obj.ObservabilityPipelineLogstashSource != nil {
		return obj.ObservabilityPipelineLogstashSource
	}

	if obj.ObservabilityPipelineSocketSource != nil {
		return obj.ObservabilityPipelineSocketSource
	}

	// all schemas are nil
	return nil
}
