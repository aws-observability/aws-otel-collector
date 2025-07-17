// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveEncryptionS3Type Type of S3 encryption for a destination.
type LogsArchiveEncryptionS3Type string

// List of LogsArchiveEncryptionS3Type.
const (
	LOGSARCHIVEENCRYPTIONS3TYPE_NO_OVERRIDE LogsArchiveEncryptionS3Type = "NO_OVERRIDE"
	LOGSARCHIVEENCRYPTIONS3TYPE_SSE_S3      LogsArchiveEncryptionS3Type = "SSE_S3"
	LOGSARCHIVEENCRYPTIONS3TYPE_SSE_KMS     LogsArchiveEncryptionS3Type = "SSE_KMS"
)

var allowedLogsArchiveEncryptionS3TypeEnumValues = []LogsArchiveEncryptionS3Type{
	LOGSARCHIVEENCRYPTIONS3TYPE_NO_OVERRIDE,
	LOGSARCHIVEENCRYPTIONS3TYPE_SSE_S3,
	LOGSARCHIVEENCRYPTIONS3TYPE_SSE_KMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsArchiveEncryptionS3Type) GetAllowedValues() []LogsArchiveEncryptionS3Type {
	return allowedLogsArchiveEncryptionS3TypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsArchiveEncryptionS3Type) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsArchiveEncryptionS3Type(value)
	return nil
}

// NewLogsArchiveEncryptionS3TypeFromValue returns a pointer to a valid LogsArchiveEncryptionS3Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsArchiveEncryptionS3TypeFromValue(v string) (*LogsArchiveEncryptionS3Type, error) {
	ev := LogsArchiveEncryptionS3Type(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsArchiveEncryptionS3Type: valid values are %v", v, allowedLogsArchiveEncryptionS3TypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsArchiveEncryptionS3Type) IsValid() bool {
	for _, existing := range allowedLogsArchiveEncryptionS3TypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsArchiveEncryptionS3Type value.
func (v LogsArchiveEncryptionS3Type) Ptr() *LogsArchiveEncryptionS3Type {
	return &v
}
