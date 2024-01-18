// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeNotifyEndStateTypes State that will trigger a monitor notification when the `notify_end_types` action occurs.
type DowntimeNotifyEndStateTypes string

// List of DowntimeNotifyEndStateTypes.
const (
	DOWNTIMENOTIFYENDSTATETYPES_ALERT   DowntimeNotifyEndStateTypes = "alert"
	DOWNTIMENOTIFYENDSTATETYPES_NO_DATA DowntimeNotifyEndStateTypes = "no data"
	DOWNTIMENOTIFYENDSTATETYPES_WARN    DowntimeNotifyEndStateTypes = "warn"
)

var allowedDowntimeNotifyEndStateTypesEnumValues = []DowntimeNotifyEndStateTypes{
	DOWNTIMENOTIFYENDSTATETYPES_ALERT,
	DOWNTIMENOTIFYENDSTATETYPES_NO_DATA,
	DOWNTIMENOTIFYENDSTATETYPES_WARN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DowntimeNotifyEndStateTypes) GetAllowedValues() []DowntimeNotifyEndStateTypes {
	return allowedDowntimeNotifyEndStateTypesEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DowntimeNotifyEndStateTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DowntimeNotifyEndStateTypes(value)
	return nil
}

// NewDowntimeNotifyEndStateTypesFromValue returns a pointer to a valid DowntimeNotifyEndStateTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDowntimeNotifyEndStateTypesFromValue(v string) (*DowntimeNotifyEndStateTypes, error) {
	ev := DowntimeNotifyEndStateTypes(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DowntimeNotifyEndStateTypes: valid values are %v", v, allowedDowntimeNotifyEndStateTypesEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DowntimeNotifyEndStateTypes) IsValid() bool {
	for _, existing := range allowedDowntimeNotifyEndStateTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DowntimeNotifyEndStateTypes value.
func (v DowntimeNotifyEndStateTypes) Ptr() *DowntimeNotifyEndStateTypes {
	return &v
}
