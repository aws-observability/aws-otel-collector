// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsStorageTier Specifies storage type as indexes, online-archives or flex
type LogsStorageTier string

// List of LogsStorageTier.
const (
	LOGSSTORAGETIER_INDEXES         LogsStorageTier = "indexes"
	LOGSSTORAGETIER_ONLINE_ARCHIVES LogsStorageTier = "online-archives"
	LOGSSTORAGETIER_FLEX            LogsStorageTier = "flex"
)

var allowedLogsStorageTierEnumValues = []LogsStorageTier{
	LOGSSTORAGETIER_INDEXES,
	LOGSSTORAGETIER_ONLINE_ARCHIVES,
	LOGSSTORAGETIER_FLEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsStorageTier) GetAllowedValues() []LogsStorageTier {
	return allowedLogsStorageTierEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsStorageTier) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsStorageTier(value)
	return nil
}

// NewLogsStorageTierFromValue returns a pointer to a valid LogsStorageTier
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsStorageTierFromValue(v string) (*LogsStorageTier, error) {
	ev := LogsStorageTier(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsStorageTier: valid values are %v", v, allowedLogsStorageTierEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsStorageTier) IsValid() bool {
	for _, existing := range allowedLogsStorageTierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsStorageTier value.
func (v LogsStorageTier) Ptr() *LogsStorageTier {
	return &v
}
