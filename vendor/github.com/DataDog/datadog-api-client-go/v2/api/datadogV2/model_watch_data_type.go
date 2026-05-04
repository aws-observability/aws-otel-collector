// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WatchDataType Rum replay watch resource type.
type WatchDataType string

// List of WatchDataType.
const (
	WATCHDATATYPE_RUM_REPLAY_WATCH WatchDataType = "rum_replay_watch"
)

var allowedWatchDataTypeEnumValues = []WatchDataType{
	WATCHDATATYPE_RUM_REPLAY_WATCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WatchDataType) GetAllowedValues() []WatchDataType {
	return allowedWatchDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WatchDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WatchDataType(value)
	return nil
}

// NewWatchDataTypeFromValue returns a pointer to a valid WatchDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWatchDataTypeFromValue(v string) (*WatchDataType, error) {
	ev := WatchDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WatchDataType: valid values are %v", v, allowedWatchDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WatchDataType) IsValid() bool {
	for _, existing := range allowedWatchDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WatchDataType value.
func (v WatchDataType) Ptr() *WatchDataType {
	return &v
}
