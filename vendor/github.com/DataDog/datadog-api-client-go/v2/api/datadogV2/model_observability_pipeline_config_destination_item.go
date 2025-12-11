// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigDestinationItem - A destination for the pipeline.
type ObservabilityPipelineConfigDestinationItem struct {
	ObservabilityPipelineDatadogLogsDestination            *ObservabilityPipelineDatadogLogsDestination
	ObservabilityPipelineAmazonS3Destination               *ObservabilityPipelineAmazonS3Destination
	ObservabilityPipelineGoogleCloudStorageDestination     *ObservabilityPipelineGoogleCloudStorageDestination
	ObservabilityPipelineSplunkHecDestination              *ObservabilityPipelineSplunkHecDestination
	ObservabilityPipelineSumoLogicDestination              *ObservabilityPipelineSumoLogicDestination
	ObservabilityPipelineElasticsearchDestination          *ObservabilityPipelineElasticsearchDestination
	ObservabilityPipelineRsyslogDestination                *ObservabilityPipelineRsyslogDestination
	ObservabilityPipelineSyslogNgDestination               *ObservabilityPipelineSyslogNgDestination
	AzureStorageDestination                                *AzureStorageDestination
	MicrosoftSentinelDestination                           *MicrosoftSentinelDestination
	ObservabilityPipelineGoogleChronicleDestination        *ObservabilityPipelineGoogleChronicleDestination
	ObservabilityPipelineNewRelicDestination               *ObservabilityPipelineNewRelicDestination
	ObservabilityPipelineSentinelOneDestination            *ObservabilityPipelineSentinelOneDestination
	ObservabilityPipelineOpenSearchDestination             *ObservabilityPipelineOpenSearchDestination
	ObservabilityPipelineAmazonOpenSearchDestination       *ObservabilityPipelineAmazonOpenSearchDestination
	ObservabilityPipelineSocketDestination                 *ObservabilityPipelineSocketDestination
	ObservabilityPipelineAmazonSecurityLakeDestination     *ObservabilityPipelineAmazonSecurityLakeDestination
	ObservabilityPipelineCrowdStrikeNextGenSiemDestination *ObservabilityPipelineCrowdStrikeNextGenSiemDestination
	ObservabilityPipelineGooglePubSubDestination           *ObservabilityPipelineGooglePubSubDestination

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineDatadogLogsDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineDatadogLogsDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineDatadogLogsDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineDatadogLogsDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineDatadogLogsDestination: v}
}

// ObservabilityPipelineAmazonS3DestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineAmazonS3Destination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineAmazonS3DestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineAmazonS3Destination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineAmazonS3Destination: v}
}

// ObservabilityPipelineGoogleCloudStorageDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineGoogleCloudStorageDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineGoogleCloudStorageDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineGoogleCloudStorageDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineGoogleCloudStorageDestination: v}
}

// ObservabilityPipelineSplunkHecDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineSplunkHecDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineSplunkHecDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineSplunkHecDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineSplunkHecDestination: v}
}

// ObservabilityPipelineSumoLogicDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineSumoLogicDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineSumoLogicDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineSumoLogicDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineSumoLogicDestination: v}
}

// ObservabilityPipelineElasticsearchDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineElasticsearchDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineElasticsearchDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineElasticsearchDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineElasticsearchDestination: v}
}

// ObservabilityPipelineRsyslogDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineRsyslogDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineRsyslogDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineRsyslogDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineRsyslogDestination: v}
}

// ObservabilityPipelineSyslogNgDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineSyslogNgDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineSyslogNgDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineSyslogNgDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineSyslogNgDestination: v}
}

// AzureStorageDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns AzureStorageDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func AzureStorageDestinationAsObservabilityPipelineConfigDestinationItem(v *AzureStorageDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{AzureStorageDestination: v}
}

// MicrosoftSentinelDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns MicrosoftSentinelDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func MicrosoftSentinelDestinationAsObservabilityPipelineConfigDestinationItem(v *MicrosoftSentinelDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{MicrosoftSentinelDestination: v}
}

// ObservabilityPipelineGoogleChronicleDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineGoogleChronicleDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineGoogleChronicleDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineGoogleChronicleDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineGoogleChronicleDestination: v}
}

// ObservabilityPipelineNewRelicDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineNewRelicDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineNewRelicDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineNewRelicDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineNewRelicDestination: v}
}

// ObservabilityPipelineSentinelOneDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineSentinelOneDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineSentinelOneDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineSentinelOneDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineSentinelOneDestination: v}
}

// ObservabilityPipelineOpenSearchDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineOpenSearchDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineOpenSearchDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineOpenSearchDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineOpenSearchDestination: v}
}

// ObservabilityPipelineAmazonOpenSearchDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineAmazonOpenSearchDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineAmazonOpenSearchDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineAmazonOpenSearchDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineAmazonOpenSearchDestination: v}
}

// ObservabilityPipelineSocketDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineSocketDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineSocketDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineSocketDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineSocketDestination: v}
}

// ObservabilityPipelineAmazonSecurityLakeDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineAmazonSecurityLakeDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineAmazonSecurityLakeDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineAmazonSecurityLakeDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineAmazonSecurityLakeDestination: v}
}

// ObservabilityPipelineCrowdStrikeNextGenSiemDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineCrowdStrikeNextGenSiemDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineCrowdStrikeNextGenSiemDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineCrowdStrikeNextGenSiemDestination: v}
}

// ObservabilityPipelineGooglePubSubDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineGooglePubSubDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineGooglePubSubDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineGooglePubSubDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineGooglePubSubDestination: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineConfigDestinationItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineDatadogLogsDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineDatadogLogsDestination)
	if err == nil {
		if obj.ObservabilityPipelineDatadogLogsDestination != nil && obj.ObservabilityPipelineDatadogLogsDestination.UnparsedObject == nil {
			jsonObservabilityPipelineDatadogLogsDestination, _ := datadog.Marshal(obj.ObservabilityPipelineDatadogLogsDestination)
			if string(jsonObservabilityPipelineDatadogLogsDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineDatadogLogsDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineDatadogLogsDestination = nil
		}
	} else {
		obj.ObservabilityPipelineDatadogLogsDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineAmazonS3Destination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonS3Destination)
	if err == nil {
		if obj.ObservabilityPipelineAmazonS3Destination != nil && obj.ObservabilityPipelineAmazonS3Destination.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonS3Destination, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonS3Destination)
			if string(jsonObservabilityPipelineAmazonS3Destination) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonS3Destination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonS3Destination = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonS3Destination = nil
	}

	// try to unmarshal data into ObservabilityPipelineGoogleCloudStorageDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineGoogleCloudStorageDestination)
	if err == nil {
		if obj.ObservabilityPipelineGoogleCloudStorageDestination != nil && obj.ObservabilityPipelineGoogleCloudStorageDestination.UnparsedObject == nil {
			jsonObservabilityPipelineGoogleCloudStorageDestination, _ := datadog.Marshal(obj.ObservabilityPipelineGoogleCloudStorageDestination)
			if string(jsonObservabilityPipelineGoogleCloudStorageDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineGoogleCloudStorageDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineGoogleCloudStorageDestination = nil
		}
	} else {
		obj.ObservabilityPipelineGoogleCloudStorageDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineSplunkHecDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSplunkHecDestination)
	if err == nil {
		if obj.ObservabilityPipelineSplunkHecDestination != nil && obj.ObservabilityPipelineSplunkHecDestination.UnparsedObject == nil {
			jsonObservabilityPipelineSplunkHecDestination, _ := datadog.Marshal(obj.ObservabilityPipelineSplunkHecDestination)
			if string(jsonObservabilityPipelineSplunkHecDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineSplunkHecDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSplunkHecDestination = nil
		}
	} else {
		obj.ObservabilityPipelineSplunkHecDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineSumoLogicDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSumoLogicDestination)
	if err == nil {
		if obj.ObservabilityPipelineSumoLogicDestination != nil && obj.ObservabilityPipelineSumoLogicDestination.UnparsedObject == nil {
			jsonObservabilityPipelineSumoLogicDestination, _ := datadog.Marshal(obj.ObservabilityPipelineSumoLogicDestination)
			if string(jsonObservabilityPipelineSumoLogicDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineSumoLogicDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSumoLogicDestination = nil
		}
	} else {
		obj.ObservabilityPipelineSumoLogicDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineElasticsearchDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineElasticsearchDestination)
	if err == nil {
		if obj.ObservabilityPipelineElasticsearchDestination != nil && obj.ObservabilityPipelineElasticsearchDestination.UnparsedObject == nil {
			jsonObservabilityPipelineElasticsearchDestination, _ := datadog.Marshal(obj.ObservabilityPipelineElasticsearchDestination)
			if string(jsonObservabilityPipelineElasticsearchDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineElasticsearchDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineElasticsearchDestination = nil
		}
	} else {
		obj.ObservabilityPipelineElasticsearchDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineRsyslogDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineRsyslogDestination)
	if err == nil {
		if obj.ObservabilityPipelineRsyslogDestination != nil && obj.ObservabilityPipelineRsyslogDestination.UnparsedObject == nil {
			jsonObservabilityPipelineRsyslogDestination, _ := datadog.Marshal(obj.ObservabilityPipelineRsyslogDestination)
			if string(jsonObservabilityPipelineRsyslogDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineRsyslogDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineRsyslogDestination = nil
		}
	} else {
		obj.ObservabilityPipelineRsyslogDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineSyslogNgDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSyslogNgDestination)
	if err == nil {
		if obj.ObservabilityPipelineSyslogNgDestination != nil && obj.ObservabilityPipelineSyslogNgDestination.UnparsedObject == nil {
			jsonObservabilityPipelineSyslogNgDestination, _ := datadog.Marshal(obj.ObservabilityPipelineSyslogNgDestination)
			if string(jsonObservabilityPipelineSyslogNgDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineSyslogNgDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSyslogNgDestination = nil
		}
	} else {
		obj.ObservabilityPipelineSyslogNgDestination = nil
	}

	// try to unmarshal data into AzureStorageDestination
	err = datadog.Unmarshal(data, &obj.AzureStorageDestination)
	if err == nil {
		if obj.AzureStorageDestination != nil && obj.AzureStorageDestination.UnparsedObject == nil {
			jsonAzureStorageDestination, _ := datadog.Marshal(obj.AzureStorageDestination)
			if string(jsonAzureStorageDestination) == "{}" { // empty struct
				obj.AzureStorageDestination = nil
			} else {
				match++
			}
		} else {
			obj.AzureStorageDestination = nil
		}
	} else {
		obj.AzureStorageDestination = nil
	}

	// try to unmarshal data into MicrosoftSentinelDestination
	err = datadog.Unmarshal(data, &obj.MicrosoftSentinelDestination)
	if err == nil {
		if obj.MicrosoftSentinelDestination != nil && obj.MicrosoftSentinelDestination.UnparsedObject == nil {
			jsonMicrosoftSentinelDestination, _ := datadog.Marshal(obj.MicrosoftSentinelDestination)
			if string(jsonMicrosoftSentinelDestination) == "{}" { // empty struct
				obj.MicrosoftSentinelDestination = nil
			} else {
				match++
			}
		} else {
			obj.MicrosoftSentinelDestination = nil
		}
	} else {
		obj.MicrosoftSentinelDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineGoogleChronicleDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineGoogleChronicleDestination)
	if err == nil {
		if obj.ObservabilityPipelineGoogleChronicleDestination != nil && obj.ObservabilityPipelineGoogleChronicleDestination.UnparsedObject == nil {
			jsonObservabilityPipelineGoogleChronicleDestination, _ := datadog.Marshal(obj.ObservabilityPipelineGoogleChronicleDestination)
			if string(jsonObservabilityPipelineGoogleChronicleDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineGoogleChronicleDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineGoogleChronicleDestination = nil
		}
	} else {
		obj.ObservabilityPipelineGoogleChronicleDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineNewRelicDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineNewRelicDestination)
	if err == nil {
		if obj.ObservabilityPipelineNewRelicDestination != nil && obj.ObservabilityPipelineNewRelicDestination.UnparsedObject == nil {
			jsonObservabilityPipelineNewRelicDestination, _ := datadog.Marshal(obj.ObservabilityPipelineNewRelicDestination)
			if string(jsonObservabilityPipelineNewRelicDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineNewRelicDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineNewRelicDestination = nil
		}
	} else {
		obj.ObservabilityPipelineNewRelicDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineSentinelOneDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSentinelOneDestination)
	if err == nil {
		if obj.ObservabilityPipelineSentinelOneDestination != nil && obj.ObservabilityPipelineSentinelOneDestination.UnparsedObject == nil {
			jsonObservabilityPipelineSentinelOneDestination, _ := datadog.Marshal(obj.ObservabilityPipelineSentinelOneDestination)
			if string(jsonObservabilityPipelineSentinelOneDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineSentinelOneDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSentinelOneDestination = nil
		}
	} else {
		obj.ObservabilityPipelineSentinelOneDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineOpenSearchDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineOpenSearchDestination)
	if err == nil {
		if obj.ObservabilityPipelineOpenSearchDestination != nil && obj.ObservabilityPipelineOpenSearchDestination.UnparsedObject == nil {
			jsonObservabilityPipelineOpenSearchDestination, _ := datadog.Marshal(obj.ObservabilityPipelineOpenSearchDestination)
			if string(jsonObservabilityPipelineOpenSearchDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineOpenSearchDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineOpenSearchDestination = nil
		}
	} else {
		obj.ObservabilityPipelineOpenSearchDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineAmazonOpenSearchDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonOpenSearchDestination)
	if err == nil {
		if obj.ObservabilityPipelineAmazonOpenSearchDestination != nil && obj.ObservabilityPipelineAmazonOpenSearchDestination.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonOpenSearchDestination, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonOpenSearchDestination)
			if string(jsonObservabilityPipelineAmazonOpenSearchDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonOpenSearchDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonOpenSearchDestination = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonOpenSearchDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineSocketDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSocketDestination)
	if err == nil {
		if obj.ObservabilityPipelineSocketDestination != nil && obj.ObservabilityPipelineSocketDestination.UnparsedObject == nil {
			jsonObservabilityPipelineSocketDestination, _ := datadog.Marshal(obj.ObservabilityPipelineSocketDestination)
			if string(jsonObservabilityPipelineSocketDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineSocketDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSocketDestination = nil
		}
	} else {
		obj.ObservabilityPipelineSocketDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineAmazonSecurityLakeDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonSecurityLakeDestination)
	if err == nil {
		if obj.ObservabilityPipelineAmazonSecurityLakeDestination != nil && obj.ObservabilityPipelineAmazonSecurityLakeDestination.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonSecurityLakeDestination, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonSecurityLakeDestination)
			if string(jsonObservabilityPipelineAmazonSecurityLakeDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonSecurityLakeDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonSecurityLakeDestination = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonSecurityLakeDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineCrowdStrikeNextGenSiemDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineCrowdStrikeNextGenSiemDestination)
	if err == nil {
		if obj.ObservabilityPipelineCrowdStrikeNextGenSiemDestination != nil && obj.ObservabilityPipelineCrowdStrikeNextGenSiemDestination.UnparsedObject == nil {
			jsonObservabilityPipelineCrowdStrikeNextGenSiemDestination, _ := datadog.Marshal(obj.ObservabilityPipelineCrowdStrikeNextGenSiemDestination)
			if string(jsonObservabilityPipelineCrowdStrikeNextGenSiemDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineCrowdStrikeNextGenSiemDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineCrowdStrikeNextGenSiemDestination = nil
		}
	} else {
		obj.ObservabilityPipelineCrowdStrikeNextGenSiemDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineGooglePubSubDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineGooglePubSubDestination)
	if err == nil {
		if obj.ObservabilityPipelineGooglePubSubDestination != nil && obj.ObservabilityPipelineGooglePubSubDestination.UnparsedObject == nil {
			jsonObservabilityPipelineGooglePubSubDestination, _ := datadog.Marshal(obj.ObservabilityPipelineGooglePubSubDestination)
			if string(jsonObservabilityPipelineGooglePubSubDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineGooglePubSubDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineGooglePubSubDestination = nil
		}
	} else {
		obj.ObservabilityPipelineGooglePubSubDestination = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineDatadogLogsDestination = nil
		obj.ObservabilityPipelineAmazonS3Destination = nil
		obj.ObservabilityPipelineGoogleCloudStorageDestination = nil
		obj.ObservabilityPipelineSplunkHecDestination = nil
		obj.ObservabilityPipelineSumoLogicDestination = nil
		obj.ObservabilityPipelineElasticsearchDestination = nil
		obj.ObservabilityPipelineRsyslogDestination = nil
		obj.ObservabilityPipelineSyslogNgDestination = nil
		obj.AzureStorageDestination = nil
		obj.MicrosoftSentinelDestination = nil
		obj.ObservabilityPipelineGoogleChronicleDestination = nil
		obj.ObservabilityPipelineNewRelicDestination = nil
		obj.ObservabilityPipelineSentinelOneDestination = nil
		obj.ObservabilityPipelineOpenSearchDestination = nil
		obj.ObservabilityPipelineAmazonOpenSearchDestination = nil
		obj.ObservabilityPipelineSocketDestination = nil
		obj.ObservabilityPipelineAmazonSecurityLakeDestination = nil
		obj.ObservabilityPipelineCrowdStrikeNextGenSiemDestination = nil
		obj.ObservabilityPipelineGooglePubSubDestination = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineConfigDestinationItem) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineDatadogLogsDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineDatadogLogsDestination)
	}

	if obj.ObservabilityPipelineAmazonS3Destination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonS3Destination)
	}

	if obj.ObservabilityPipelineGoogleCloudStorageDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineGoogleCloudStorageDestination)
	}

	if obj.ObservabilityPipelineSplunkHecDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSplunkHecDestination)
	}

	if obj.ObservabilityPipelineSumoLogicDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSumoLogicDestination)
	}

	if obj.ObservabilityPipelineElasticsearchDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineElasticsearchDestination)
	}

	if obj.ObservabilityPipelineRsyslogDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineRsyslogDestination)
	}

	if obj.ObservabilityPipelineSyslogNgDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSyslogNgDestination)
	}

	if obj.AzureStorageDestination != nil {
		return datadog.Marshal(&obj.AzureStorageDestination)
	}

	if obj.MicrosoftSentinelDestination != nil {
		return datadog.Marshal(&obj.MicrosoftSentinelDestination)
	}

	if obj.ObservabilityPipelineGoogleChronicleDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineGoogleChronicleDestination)
	}

	if obj.ObservabilityPipelineNewRelicDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineNewRelicDestination)
	}

	if obj.ObservabilityPipelineSentinelOneDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSentinelOneDestination)
	}

	if obj.ObservabilityPipelineOpenSearchDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineOpenSearchDestination)
	}

	if obj.ObservabilityPipelineAmazonOpenSearchDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonOpenSearchDestination)
	}

	if obj.ObservabilityPipelineSocketDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSocketDestination)
	}

	if obj.ObservabilityPipelineAmazonSecurityLakeDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonSecurityLakeDestination)
	}

	if obj.ObservabilityPipelineCrowdStrikeNextGenSiemDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineCrowdStrikeNextGenSiemDestination)
	}

	if obj.ObservabilityPipelineGooglePubSubDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineGooglePubSubDestination)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineConfigDestinationItem) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineDatadogLogsDestination != nil {
		return obj.ObservabilityPipelineDatadogLogsDestination
	}

	if obj.ObservabilityPipelineAmazonS3Destination != nil {
		return obj.ObservabilityPipelineAmazonS3Destination
	}

	if obj.ObservabilityPipelineGoogleCloudStorageDestination != nil {
		return obj.ObservabilityPipelineGoogleCloudStorageDestination
	}

	if obj.ObservabilityPipelineSplunkHecDestination != nil {
		return obj.ObservabilityPipelineSplunkHecDestination
	}

	if obj.ObservabilityPipelineSumoLogicDestination != nil {
		return obj.ObservabilityPipelineSumoLogicDestination
	}

	if obj.ObservabilityPipelineElasticsearchDestination != nil {
		return obj.ObservabilityPipelineElasticsearchDestination
	}

	if obj.ObservabilityPipelineRsyslogDestination != nil {
		return obj.ObservabilityPipelineRsyslogDestination
	}

	if obj.ObservabilityPipelineSyslogNgDestination != nil {
		return obj.ObservabilityPipelineSyslogNgDestination
	}

	if obj.AzureStorageDestination != nil {
		return obj.AzureStorageDestination
	}

	if obj.MicrosoftSentinelDestination != nil {
		return obj.MicrosoftSentinelDestination
	}

	if obj.ObservabilityPipelineGoogleChronicleDestination != nil {
		return obj.ObservabilityPipelineGoogleChronicleDestination
	}

	if obj.ObservabilityPipelineNewRelicDestination != nil {
		return obj.ObservabilityPipelineNewRelicDestination
	}

	if obj.ObservabilityPipelineSentinelOneDestination != nil {
		return obj.ObservabilityPipelineSentinelOneDestination
	}

	if obj.ObservabilityPipelineOpenSearchDestination != nil {
		return obj.ObservabilityPipelineOpenSearchDestination
	}

	if obj.ObservabilityPipelineAmazonOpenSearchDestination != nil {
		return obj.ObservabilityPipelineAmazonOpenSearchDestination
	}

	if obj.ObservabilityPipelineSocketDestination != nil {
		return obj.ObservabilityPipelineSocketDestination
	}

	if obj.ObservabilityPipelineAmazonSecurityLakeDestination != nil {
		return obj.ObservabilityPipelineAmazonSecurityLakeDestination
	}

	if obj.ObservabilityPipelineCrowdStrikeNextGenSiemDestination != nil {
		return obj.ObservabilityPipelineCrowdStrikeNextGenSiemDestination
	}

	if obj.ObservabilityPipelineGooglePubSubDestination != nil {
		return obj.ObservabilityPipelineGooglePubSubDestination
	}

	// all schemas are nil
	return nil
}
