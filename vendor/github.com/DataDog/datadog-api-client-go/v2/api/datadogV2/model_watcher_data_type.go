// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WatcherDataType Rum replay watcher resource type.
type WatcherDataType string

// List of WatcherDataType.
const (
	WATCHERDATATYPE_RUM_REPLAY_WATCHER WatcherDataType = "rum_replay_watcher"
)

var allowedWatcherDataTypeEnumValues = []WatcherDataType{
	WATCHERDATATYPE_RUM_REPLAY_WATCHER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WatcherDataType) GetAllowedValues() []WatcherDataType {
	return allowedWatcherDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WatcherDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WatcherDataType(value)
	return nil
}

// NewWatcherDataTypeFromValue returns a pointer to a valid WatcherDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWatcherDataTypeFromValue(v string) (*WatcherDataType, error) {
	ev := WatcherDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WatcherDataType: valid values are %v", v, allowedWatcherDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WatcherDataType) IsValid() bool {
	for _, existing := range allowedWatcherDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WatcherDataType value.
func (v WatcherDataType) Ptr() *WatcherDataType {
	return &v
}
