// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchTableRequestDataAttributesFileMetadata - Metadata specifying where and how to access the reference table's data file.
type PatchTableRequestDataAttributesFileMetadata struct {
	PatchTableRequestDataAttributesFileMetadataCloudStorage *PatchTableRequestDataAttributesFileMetadataCloudStorage
	PatchTableRequestDataAttributesFileMetadataLocalFile    *PatchTableRequestDataAttributesFileMetadataLocalFile

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// PatchTableRequestDataAttributesFileMetadataCloudStorageAsPatchTableRequestDataAttributesFileMetadata is a convenience function that returns PatchTableRequestDataAttributesFileMetadataCloudStorage wrapped in PatchTableRequestDataAttributesFileMetadata.
func PatchTableRequestDataAttributesFileMetadataCloudStorageAsPatchTableRequestDataAttributesFileMetadata(v *PatchTableRequestDataAttributesFileMetadataCloudStorage) PatchTableRequestDataAttributesFileMetadata {
	return PatchTableRequestDataAttributesFileMetadata{PatchTableRequestDataAttributesFileMetadataCloudStorage: v}
}

// PatchTableRequestDataAttributesFileMetadataLocalFileAsPatchTableRequestDataAttributesFileMetadata is a convenience function that returns PatchTableRequestDataAttributesFileMetadataLocalFile wrapped in PatchTableRequestDataAttributesFileMetadata.
func PatchTableRequestDataAttributesFileMetadataLocalFileAsPatchTableRequestDataAttributesFileMetadata(v *PatchTableRequestDataAttributesFileMetadataLocalFile) PatchTableRequestDataAttributesFileMetadata {
	return PatchTableRequestDataAttributesFileMetadata{PatchTableRequestDataAttributesFileMetadataLocalFile: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *PatchTableRequestDataAttributesFileMetadata) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PatchTableRequestDataAttributesFileMetadataCloudStorage
	err = datadog.Unmarshal(data, &obj.PatchTableRequestDataAttributesFileMetadataCloudStorage)
	if err == nil {
		if obj.PatchTableRequestDataAttributesFileMetadataCloudStorage != nil && obj.PatchTableRequestDataAttributesFileMetadataCloudStorage.UnparsedObject == nil {
			jsonPatchTableRequestDataAttributesFileMetadataCloudStorage, _ := datadog.Marshal(obj.PatchTableRequestDataAttributesFileMetadataCloudStorage)
			if string(jsonPatchTableRequestDataAttributesFileMetadataCloudStorage) == "{}" && string(data) != "{}" { // empty struct
				obj.PatchTableRequestDataAttributesFileMetadataCloudStorage = nil
			} else {
				match++
			}
		} else {
			obj.PatchTableRequestDataAttributesFileMetadataCloudStorage = nil
		}
	} else {
		obj.PatchTableRequestDataAttributesFileMetadataCloudStorage = nil
	}

	// try to unmarshal data into PatchTableRequestDataAttributesFileMetadataLocalFile
	err = datadog.Unmarshal(data, &obj.PatchTableRequestDataAttributesFileMetadataLocalFile)
	if err == nil {
		if obj.PatchTableRequestDataAttributesFileMetadataLocalFile != nil && obj.PatchTableRequestDataAttributesFileMetadataLocalFile.UnparsedObject == nil {
			jsonPatchTableRequestDataAttributesFileMetadataLocalFile, _ := datadog.Marshal(obj.PatchTableRequestDataAttributesFileMetadataLocalFile)
			if string(jsonPatchTableRequestDataAttributesFileMetadataLocalFile) == "{}" { // empty struct
				obj.PatchTableRequestDataAttributesFileMetadataLocalFile = nil
			} else {
				match++
			}
		} else {
			obj.PatchTableRequestDataAttributesFileMetadataLocalFile = nil
		}
	} else {
		obj.PatchTableRequestDataAttributesFileMetadataLocalFile = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.PatchTableRequestDataAttributesFileMetadataCloudStorage = nil
		obj.PatchTableRequestDataAttributesFileMetadataLocalFile = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj PatchTableRequestDataAttributesFileMetadata) MarshalJSON() ([]byte, error) {
	if obj.PatchTableRequestDataAttributesFileMetadataCloudStorage != nil {
		return datadog.Marshal(&obj.PatchTableRequestDataAttributesFileMetadataCloudStorage)
	}

	if obj.PatchTableRequestDataAttributesFileMetadataLocalFile != nil {
		return datadog.Marshal(&obj.PatchTableRequestDataAttributesFileMetadataLocalFile)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *PatchTableRequestDataAttributesFileMetadata) GetActualInstance() interface{} {
	if obj.PatchTableRequestDataAttributesFileMetadataCloudStorage != nil {
		return obj.PatchTableRequestDataAttributesFileMetadataCloudStorage
	}

	if obj.PatchTableRequestDataAttributesFileMetadataLocalFile != nil {
		return obj.PatchTableRequestDataAttributesFileMetadataLocalFile
	}

	// all schemas are nil
	return nil
}
