// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableResultV2DataAttributesFileMetadata - Metadata specifying where and how to access the reference table's data file.
type TableResultV2DataAttributesFileMetadata struct {
	TableResultV2DataAttributesFileMetadataCloudStorage *TableResultV2DataAttributesFileMetadataCloudStorage
	TableResultV2DataAttributesFileMetadataLocalFile    *TableResultV2DataAttributesFileMetadataLocalFile

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TableResultV2DataAttributesFileMetadataCloudStorageAsTableResultV2DataAttributesFileMetadata is a convenience function that returns TableResultV2DataAttributesFileMetadataCloudStorage wrapped in TableResultV2DataAttributesFileMetadata.
func TableResultV2DataAttributesFileMetadataCloudStorageAsTableResultV2DataAttributesFileMetadata(v *TableResultV2DataAttributesFileMetadataCloudStorage) TableResultV2DataAttributesFileMetadata {
	return TableResultV2DataAttributesFileMetadata{TableResultV2DataAttributesFileMetadataCloudStorage: v}
}

// TableResultV2DataAttributesFileMetadataLocalFileAsTableResultV2DataAttributesFileMetadata is a convenience function that returns TableResultV2DataAttributesFileMetadataLocalFile wrapped in TableResultV2DataAttributesFileMetadata.
func TableResultV2DataAttributesFileMetadataLocalFileAsTableResultV2DataAttributesFileMetadata(v *TableResultV2DataAttributesFileMetadataLocalFile) TableResultV2DataAttributesFileMetadata {
	return TableResultV2DataAttributesFileMetadata{TableResultV2DataAttributesFileMetadataLocalFile: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TableResultV2DataAttributesFileMetadata) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TableResultV2DataAttributesFileMetadataCloudStorage
	err = datadog.Unmarshal(data, &obj.TableResultV2DataAttributesFileMetadataCloudStorage)
	if err == nil {
		if obj.TableResultV2DataAttributesFileMetadataCloudStorage != nil && obj.TableResultV2DataAttributesFileMetadataCloudStorage.UnparsedObject == nil {
			jsonTableResultV2DataAttributesFileMetadataCloudStorage, _ := datadog.Marshal(obj.TableResultV2DataAttributesFileMetadataCloudStorage)
			if string(jsonTableResultV2DataAttributesFileMetadataCloudStorage) == "{}" { // empty struct
				obj.TableResultV2DataAttributesFileMetadataCloudStorage = nil
			} else {
				match++
			}
		} else {
			obj.TableResultV2DataAttributesFileMetadataCloudStorage = nil
		}
	} else {
		obj.TableResultV2DataAttributesFileMetadataCloudStorage = nil
	}

	// try to unmarshal data into TableResultV2DataAttributesFileMetadataLocalFile
	err = datadog.Unmarshal(data, &obj.TableResultV2DataAttributesFileMetadataLocalFile)
	if err == nil {
		if obj.TableResultV2DataAttributesFileMetadataLocalFile != nil && obj.TableResultV2DataAttributesFileMetadataLocalFile.UnparsedObject == nil {
			jsonTableResultV2DataAttributesFileMetadataLocalFile, _ := datadog.Marshal(obj.TableResultV2DataAttributesFileMetadataLocalFile)
			if string(jsonTableResultV2DataAttributesFileMetadataLocalFile) == "{}" && string(data) != "{}" { // empty struct
				obj.TableResultV2DataAttributesFileMetadataLocalFile = nil
			} else {
				match++
			}
		} else {
			obj.TableResultV2DataAttributesFileMetadataLocalFile = nil
		}
	} else {
		obj.TableResultV2DataAttributesFileMetadataLocalFile = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TableResultV2DataAttributesFileMetadataCloudStorage = nil
		obj.TableResultV2DataAttributesFileMetadataLocalFile = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TableResultV2DataAttributesFileMetadata) MarshalJSON() ([]byte, error) {
	if obj.TableResultV2DataAttributesFileMetadataCloudStorage != nil {
		return datadog.Marshal(&obj.TableResultV2DataAttributesFileMetadataCloudStorage)
	}

	if obj.TableResultV2DataAttributesFileMetadataLocalFile != nil {
		return datadog.Marshal(&obj.TableResultV2DataAttributesFileMetadataLocalFile)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TableResultV2DataAttributesFileMetadata) GetActualInstance() interface{} {
	if obj.TableResultV2DataAttributesFileMetadataCloudStorage != nil {
		return obj.TableResultV2DataAttributesFileMetadataCloudStorage
	}

	if obj.TableResultV2DataAttributesFileMetadataLocalFile != nil {
		return obj.TableResultV2DataAttributesFileMetadataLocalFile
	}

	// all schemas are nil
	return nil
}
