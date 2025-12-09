// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTableRequestDataAttributesFileMetadata - Metadata specifying where and how to access the reference table's data file.
type CreateTableRequestDataAttributesFileMetadata struct {
	CreateTableRequestDataAttributesFileMetadataCloudStorage *CreateTableRequestDataAttributesFileMetadataCloudStorage
	CreateTableRequestDataAttributesFileMetadataLocalFile    *CreateTableRequestDataAttributesFileMetadataLocalFile

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CreateTableRequestDataAttributesFileMetadataCloudStorageAsCreateTableRequestDataAttributesFileMetadata is a convenience function that returns CreateTableRequestDataAttributesFileMetadataCloudStorage wrapped in CreateTableRequestDataAttributesFileMetadata.
func CreateTableRequestDataAttributesFileMetadataCloudStorageAsCreateTableRequestDataAttributesFileMetadata(v *CreateTableRequestDataAttributesFileMetadataCloudStorage) CreateTableRequestDataAttributesFileMetadata {
	return CreateTableRequestDataAttributesFileMetadata{CreateTableRequestDataAttributesFileMetadataCloudStorage: v}
}

// CreateTableRequestDataAttributesFileMetadataLocalFileAsCreateTableRequestDataAttributesFileMetadata is a convenience function that returns CreateTableRequestDataAttributesFileMetadataLocalFile wrapped in CreateTableRequestDataAttributesFileMetadata.
func CreateTableRequestDataAttributesFileMetadataLocalFileAsCreateTableRequestDataAttributesFileMetadata(v *CreateTableRequestDataAttributesFileMetadataLocalFile) CreateTableRequestDataAttributesFileMetadata {
	return CreateTableRequestDataAttributesFileMetadata{CreateTableRequestDataAttributesFileMetadataLocalFile: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CreateTableRequestDataAttributesFileMetadata) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateTableRequestDataAttributesFileMetadataCloudStorage
	err = datadog.Unmarshal(data, &obj.CreateTableRequestDataAttributesFileMetadataCloudStorage)
	if err == nil {
		if obj.CreateTableRequestDataAttributesFileMetadataCloudStorage != nil && obj.CreateTableRequestDataAttributesFileMetadataCloudStorage.UnparsedObject == nil {
			jsonCreateTableRequestDataAttributesFileMetadataCloudStorage, _ := datadog.Marshal(obj.CreateTableRequestDataAttributesFileMetadataCloudStorage)
			if string(jsonCreateTableRequestDataAttributesFileMetadataCloudStorage) == "{}" { // empty struct
				obj.CreateTableRequestDataAttributesFileMetadataCloudStorage = nil
			} else {
				match++
			}
		} else {
			obj.CreateTableRequestDataAttributesFileMetadataCloudStorage = nil
		}
	} else {
		obj.CreateTableRequestDataAttributesFileMetadataCloudStorage = nil
	}

	// try to unmarshal data into CreateTableRequestDataAttributesFileMetadataLocalFile
	err = datadog.Unmarshal(data, &obj.CreateTableRequestDataAttributesFileMetadataLocalFile)
	if err == nil {
		if obj.CreateTableRequestDataAttributesFileMetadataLocalFile != nil && obj.CreateTableRequestDataAttributesFileMetadataLocalFile.UnparsedObject == nil {
			jsonCreateTableRequestDataAttributesFileMetadataLocalFile, _ := datadog.Marshal(obj.CreateTableRequestDataAttributesFileMetadataLocalFile)
			if string(jsonCreateTableRequestDataAttributesFileMetadataLocalFile) == "{}" { // empty struct
				obj.CreateTableRequestDataAttributesFileMetadataLocalFile = nil
			} else {
				match++
			}
		} else {
			obj.CreateTableRequestDataAttributesFileMetadataLocalFile = nil
		}
	} else {
		obj.CreateTableRequestDataAttributesFileMetadataLocalFile = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CreateTableRequestDataAttributesFileMetadataCloudStorage = nil
		obj.CreateTableRequestDataAttributesFileMetadataLocalFile = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CreateTableRequestDataAttributesFileMetadata) MarshalJSON() ([]byte, error) {
	if obj.CreateTableRequestDataAttributesFileMetadataCloudStorage != nil {
		return datadog.Marshal(&obj.CreateTableRequestDataAttributesFileMetadataCloudStorage)
	}

	if obj.CreateTableRequestDataAttributesFileMetadataLocalFile != nil {
		return datadog.Marshal(&obj.CreateTableRequestDataAttributesFileMetadataLocalFile)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CreateTableRequestDataAttributesFileMetadata) GetActualInstance() interface{} {
	if obj.CreateTableRequestDataAttributesFileMetadataCloudStorage != nil {
		return obj.CreateTableRequestDataAttributesFileMetadataCloudStorage
	}

	if obj.CreateTableRequestDataAttributesFileMetadataLocalFile != nil {
		return obj.CreateTableRequestDataAttributesFileMetadataLocalFile
	}

	// all schemas are nil
	return nil
}
