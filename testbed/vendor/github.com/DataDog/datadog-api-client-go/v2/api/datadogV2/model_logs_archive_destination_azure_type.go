// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveDestinationAzureType Type of the Azure archive destination.
type LogsArchiveDestinationAzureType string

// List of LogsArchiveDestinationAzureType.
const (
	LOGSARCHIVEDESTINATIONAZURETYPE_AZURE LogsArchiveDestinationAzureType = "azure"
)

var allowedLogsArchiveDestinationAzureTypeEnumValues = []LogsArchiveDestinationAzureType{
	LOGSARCHIVEDESTINATIONAZURETYPE_AZURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsArchiveDestinationAzureType) GetAllowedValues() []LogsArchiveDestinationAzureType {
	return allowedLogsArchiveDestinationAzureTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsArchiveDestinationAzureType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsArchiveDestinationAzureType(value)
	return nil
}

// NewLogsArchiveDestinationAzureTypeFromValue returns a pointer to a valid LogsArchiveDestinationAzureType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsArchiveDestinationAzureTypeFromValue(v string) (*LogsArchiveDestinationAzureType, error) {
	ev := LogsArchiveDestinationAzureType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsArchiveDestinationAzureType: valid values are %v", v, allowedLogsArchiveDestinationAzureTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsArchiveDestinationAzureType) IsValid() bool {
	for _, existing := range allowedLogsArchiveDestinationAzureTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsArchiveDestinationAzureType value.
func (v LogsArchiveDestinationAzureType) Ptr() *LogsArchiveDestinationAzureType {
	return &v
}
