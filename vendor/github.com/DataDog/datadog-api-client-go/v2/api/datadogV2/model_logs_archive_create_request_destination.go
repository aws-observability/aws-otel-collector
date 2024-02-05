// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveCreateRequestDestination - An archive's destination.
type LogsArchiveCreateRequestDestination struct {
	LogsArchiveDestinationAzure *LogsArchiveDestinationAzure
	LogsArchiveDestinationGCS   *LogsArchiveDestinationGCS
	LogsArchiveDestinationS3    *LogsArchiveDestinationS3

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LogsArchiveDestinationAzureAsLogsArchiveCreateRequestDestination is a convenience function that returns LogsArchiveDestinationAzure wrapped in LogsArchiveCreateRequestDestination.
func LogsArchiveDestinationAzureAsLogsArchiveCreateRequestDestination(v *LogsArchiveDestinationAzure) LogsArchiveCreateRequestDestination {
	return LogsArchiveCreateRequestDestination{LogsArchiveDestinationAzure: v}
}

// LogsArchiveDestinationGCSAsLogsArchiveCreateRequestDestination is a convenience function that returns LogsArchiveDestinationGCS wrapped in LogsArchiveCreateRequestDestination.
func LogsArchiveDestinationGCSAsLogsArchiveCreateRequestDestination(v *LogsArchiveDestinationGCS) LogsArchiveCreateRequestDestination {
	return LogsArchiveCreateRequestDestination{LogsArchiveDestinationGCS: v}
}

// LogsArchiveDestinationS3AsLogsArchiveCreateRequestDestination is a convenience function that returns LogsArchiveDestinationS3 wrapped in LogsArchiveCreateRequestDestination.
func LogsArchiveDestinationS3AsLogsArchiveCreateRequestDestination(v *LogsArchiveDestinationS3) LogsArchiveCreateRequestDestination {
	return LogsArchiveCreateRequestDestination{LogsArchiveDestinationS3: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LogsArchiveCreateRequestDestination) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsArchiveDestinationAzure
	err = datadog.Unmarshal(data, &obj.LogsArchiveDestinationAzure)
	if err == nil {
		if obj.LogsArchiveDestinationAzure != nil && obj.LogsArchiveDestinationAzure.UnparsedObject == nil {
			jsonLogsArchiveDestinationAzure, _ := datadog.Marshal(obj.LogsArchiveDestinationAzure)
			if string(jsonLogsArchiveDestinationAzure) == "{}" { // empty struct
				obj.LogsArchiveDestinationAzure = nil
			} else {
				match++
			}
		} else {
			obj.LogsArchiveDestinationAzure = nil
		}
	} else {
		obj.LogsArchiveDestinationAzure = nil
	}

	// try to unmarshal data into LogsArchiveDestinationGCS
	err = datadog.Unmarshal(data, &obj.LogsArchiveDestinationGCS)
	if err == nil {
		if obj.LogsArchiveDestinationGCS != nil && obj.LogsArchiveDestinationGCS.UnparsedObject == nil {
			jsonLogsArchiveDestinationGCS, _ := datadog.Marshal(obj.LogsArchiveDestinationGCS)
			if string(jsonLogsArchiveDestinationGCS) == "{}" { // empty struct
				obj.LogsArchiveDestinationGCS = nil
			} else {
				match++
			}
		} else {
			obj.LogsArchiveDestinationGCS = nil
		}
	} else {
		obj.LogsArchiveDestinationGCS = nil
	}

	// try to unmarshal data into LogsArchiveDestinationS3
	err = datadog.Unmarshal(data, &obj.LogsArchiveDestinationS3)
	if err == nil {
		if obj.LogsArchiveDestinationS3 != nil && obj.LogsArchiveDestinationS3.UnparsedObject == nil {
			jsonLogsArchiveDestinationS3, _ := datadog.Marshal(obj.LogsArchiveDestinationS3)
			if string(jsonLogsArchiveDestinationS3) == "{}" { // empty struct
				obj.LogsArchiveDestinationS3 = nil
			} else {
				match++
			}
		} else {
			obj.LogsArchiveDestinationS3 = nil
		}
	} else {
		obj.LogsArchiveDestinationS3 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LogsArchiveDestinationAzure = nil
		obj.LogsArchiveDestinationGCS = nil
		obj.LogsArchiveDestinationS3 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LogsArchiveCreateRequestDestination) MarshalJSON() ([]byte, error) {
	if obj.LogsArchiveDestinationAzure != nil {
		return datadog.Marshal(&obj.LogsArchiveDestinationAzure)
	}

	if obj.LogsArchiveDestinationGCS != nil {
		return datadog.Marshal(&obj.LogsArchiveDestinationGCS)
	}

	if obj.LogsArchiveDestinationS3 != nil {
		return datadog.Marshal(&obj.LogsArchiveDestinationS3)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LogsArchiveCreateRequestDestination) GetActualInstance() interface{} {
	if obj.LogsArchiveDestinationAzure != nil {
		return obj.LogsArchiveDestinationAzure
	}

	if obj.LogsArchiveDestinationGCS != nil {
		return obj.LogsArchiveDestinationGCS
	}

	if obj.LogsArchiveDestinationS3 != nil {
		return obj.LogsArchiveDestinationS3
	}

	// all schemas are nil
	return nil
}
