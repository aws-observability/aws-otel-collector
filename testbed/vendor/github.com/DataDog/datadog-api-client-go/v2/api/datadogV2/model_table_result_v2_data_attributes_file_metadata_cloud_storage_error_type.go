// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableResultV2DataAttributesFileMetadataCloudStorageErrorType The type of error that occurred during file processing. This field provides high-level error categories for easier troubleshooting and is only present when there are errors.
type TableResultV2DataAttributesFileMetadataCloudStorageErrorType string

// List of TableResultV2DataAttributesFileMetadataCloudStorageErrorType.
const (
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_TABLE_SCHEMA_ERROR  TableResultV2DataAttributesFileMetadataCloudStorageErrorType = "TABLE_SCHEMA_ERROR"
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_FILE_FORMAT_ERROR   TableResultV2DataAttributesFileMetadataCloudStorageErrorType = "FILE_FORMAT_ERROR"
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_CONFIGURATION_ERROR TableResultV2DataAttributesFileMetadataCloudStorageErrorType = "CONFIGURATION_ERROR"
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_QUOTA_EXCEEDED      TableResultV2DataAttributesFileMetadataCloudStorageErrorType = "QUOTA_EXCEEDED"
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_CONFLICT_ERROR      TableResultV2DataAttributesFileMetadataCloudStorageErrorType = "CONFLICT_ERROR"
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_VALIDATION_ERROR    TableResultV2DataAttributesFileMetadataCloudStorageErrorType = "VALIDATION_ERROR"
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_STATE_ERROR         TableResultV2DataAttributesFileMetadataCloudStorageErrorType = "STATE_ERROR"
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_OPERATION_ERROR     TableResultV2DataAttributesFileMetadataCloudStorageErrorType = "OPERATION_ERROR"
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_SYSTEM_ERROR        TableResultV2DataAttributesFileMetadataCloudStorageErrorType = "SYSTEM_ERROR"
)

var allowedTableResultV2DataAttributesFileMetadataCloudStorageErrorTypeEnumValues = []TableResultV2DataAttributesFileMetadataCloudStorageErrorType{
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_TABLE_SCHEMA_ERROR,
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_FILE_FORMAT_ERROR,
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_CONFIGURATION_ERROR,
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_QUOTA_EXCEEDED,
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_CONFLICT_ERROR,
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_VALIDATION_ERROR,
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_STATE_ERROR,
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_OPERATION_ERROR,
	TABLERESULTV2DATAATTRIBUTESFILEMETADATACLOUDSTORAGEERRORTYPE_SYSTEM_ERROR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableResultV2DataAttributesFileMetadataCloudStorageErrorType) GetAllowedValues() []TableResultV2DataAttributesFileMetadataCloudStorageErrorType {
	return allowedTableResultV2DataAttributesFileMetadataCloudStorageErrorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableResultV2DataAttributesFileMetadataCloudStorageErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableResultV2DataAttributesFileMetadataCloudStorageErrorType(value)
	return nil
}

// NewTableResultV2DataAttributesFileMetadataCloudStorageErrorTypeFromValue returns a pointer to a valid TableResultV2DataAttributesFileMetadataCloudStorageErrorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableResultV2DataAttributesFileMetadataCloudStorageErrorTypeFromValue(v string) (*TableResultV2DataAttributesFileMetadataCloudStorageErrorType, error) {
	ev := TableResultV2DataAttributesFileMetadataCloudStorageErrorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableResultV2DataAttributesFileMetadataCloudStorageErrorType: valid values are %v", v, allowedTableResultV2DataAttributesFileMetadataCloudStorageErrorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableResultV2DataAttributesFileMetadataCloudStorageErrorType) IsValid() bool {
	for _, existing := range allowedTableResultV2DataAttributesFileMetadataCloudStorageErrorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableResultV2DataAttributesFileMetadataCloudStorageErrorType value.
func (v TableResultV2DataAttributesFileMetadataCloudStorageErrorType) Ptr() *TableResultV2DataAttributesFileMetadataCloudStorageErrorType {
	return &v
}
