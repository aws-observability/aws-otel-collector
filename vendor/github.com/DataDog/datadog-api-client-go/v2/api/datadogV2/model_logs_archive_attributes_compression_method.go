// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveAttributesCompressionMethod The type of compression for the archive.
type LogsArchiveAttributesCompressionMethod string

// List of LogsArchiveAttributesCompressionMethod.
const (
	LOGSARCHIVEATTRIBUTESCOMPRESSIONMETHOD_GZIP LogsArchiveAttributesCompressionMethod = "GZIP"
	LOGSARCHIVEATTRIBUTESCOMPRESSIONMETHOD_ZSTD LogsArchiveAttributesCompressionMethod = "ZSTD"
)

var allowedLogsArchiveAttributesCompressionMethodEnumValues = []LogsArchiveAttributesCompressionMethod{
	LOGSARCHIVEATTRIBUTESCOMPRESSIONMETHOD_GZIP,
	LOGSARCHIVEATTRIBUTESCOMPRESSIONMETHOD_ZSTD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsArchiveAttributesCompressionMethod) GetAllowedValues() []LogsArchiveAttributesCompressionMethod {
	return allowedLogsArchiveAttributesCompressionMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsArchiveAttributesCompressionMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsArchiveAttributesCompressionMethod(value)
	return nil
}

// NewLogsArchiveAttributesCompressionMethodFromValue returns a pointer to a valid LogsArchiveAttributesCompressionMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsArchiveAttributesCompressionMethodFromValue(v string) (*LogsArchiveAttributesCompressionMethod, error) {
	ev := LogsArchiveAttributesCompressionMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsArchiveAttributesCompressionMethod: valid values are %v", v, allowedLogsArchiveAttributesCompressionMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsArchiveAttributesCompressionMethod) IsValid() bool {
	for _, existing := range allowedLogsArchiveAttributesCompressionMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsArchiveAttributesCompressionMethod value.
func (v LogsArchiveAttributesCompressionMethod) Ptr() *LogsArchiveAttributesCompressionMethod {
	return &v
}
