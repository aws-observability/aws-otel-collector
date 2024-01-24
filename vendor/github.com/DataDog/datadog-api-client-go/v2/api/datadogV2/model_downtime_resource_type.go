// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeResourceType Downtime resource type.
type DowntimeResourceType string

// List of DowntimeResourceType.
const (
	DOWNTIMERESOURCETYPE_DOWNTIME DowntimeResourceType = "downtime"
)

var allowedDowntimeResourceTypeEnumValues = []DowntimeResourceType{
	DOWNTIMERESOURCETYPE_DOWNTIME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DowntimeResourceType) GetAllowedValues() []DowntimeResourceType {
	return allowedDowntimeResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DowntimeResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DowntimeResourceType(value)
	return nil
}

// NewDowntimeResourceTypeFromValue returns a pointer to a valid DowntimeResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDowntimeResourceTypeFromValue(v string) (*DowntimeResourceType, error) {
	ev := DowntimeResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DowntimeResourceType: valid values are %v", v, allowedDowntimeResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DowntimeResourceType) IsValid() bool {
	for _, existing := range allowedDowntimeResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DowntimeResourceType value.
func (v DowntimeResourceType) Ptr() *DowntimeResourceType {
	return &v
}
