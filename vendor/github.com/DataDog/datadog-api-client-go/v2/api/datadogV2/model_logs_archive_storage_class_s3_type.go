// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveStorageClassS3Type The storage class where the archive will be stored.
type LogsArchiveStorageClassS3Type string

// List of LogsArchiveStorageClassS3Type.
const (
	LOGSARCHIVESTORAGECLASSS3TYPE_STANDARD            LogsArchiveStorageClassS3Type = "STANDARD"
	LOGSARCHIVESTORAGECLASSS3TYPE_STANDARD_IA         LogsArchiveStorageClassS3Type = "STANDARD_IA"
	LOGSARCHIVESTORAGECLASSS3TYPE_ONEZONE_IA          LogsArchiveStorageClassS3Type = "ONEZONE_IA"
	LOGSARCHIVESTORAGECLASSS3TYPE_INTELLIGENT_TIERING LogsArchiveStorageClassS3Type = "INTELLIGENT_TIERING"
	LOGSARCHIVESTORAGECLASSS3TYPE_GLACIER_IR          LogsArchiveStorageClassS3Type = "GLACIER_IR"
)

var allowedLogsArchiveStorageClassS3TypeEnumValues = []LogsArchiveStorageClassS3Type{
	LOGSARCHIVESTORAGECLASSS3TYPE_STANDARD,
	LOGSARCHIVESTORAGECLASSS3TYPE_STANDARD_IA,
	LOGSARCHIVESTORAGECLASSS3TYPE_ONEZONE_IA,
	LOGSARCHIVESTORAGECLASSS3TYPE_INTELLIGENT_TIERING,
	LOGSARCHIVESTORAGECLASSS3TYPE_GLACIER_IR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsArchiveStorageClassS3Type) GetAllowedValues() []LogsArchiveStorageClassS3Type {
	return allowedLogsArchiveStorageClassS3TypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsArchiveStorageClassS3Type) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsArchiveStorageClassS3Type(value)
	return nil
}

// NewLogsArchiveStorageClassS3TypeFromValue returns a pointer to a valid LogsArchiveStorageClassS3Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsArchiveStorageClassS3TypeFromValue(v string) (*LogsArchiveStorageClassS3Type, error) {
	ev := LogsArchiveStorageClassS3Type(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsArchiveStorageClassS3Type: valid values are %v", v, allowedLogsArchiveStorageClassS3TypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsArchiveStorageClassS3Type) IsValid() bool {
	for _, existing := range allowedLogsArchiveStorageClassS3TypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsArchiveStorageClassS3Type value.
func (v LogsArchiveStorageClassS3Type) Ptr() *LogsArchiveStorageClassS3Type {
	return &v
}
