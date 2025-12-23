// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorAssetResourceType Type of internal Datadog resource associated with a monitor asset.
type MonitorAssetResourceType string

// List of MonitorAssetResourceType.
const (
	MONITORASSETRESOURCETYPE_NOTEBOOK MonitorAssetResourceType = "notebook"
)

var allowedMonitorAssetResourceTypeEnumValues = []MonitorAssetResourceType{
	MONITORASSETRESOURCETYPE_NOTEBOOK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorAssetResourceType) GetAllowedValues() []MonitorAssetResourceType {
	return allowedMonitorAssetResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorAssetResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorAssetResourceType(value)
	return nil
}

// NewMonitorAssetResourceTypeFromValue returns a pointer to a valid MonitorAssetResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorAssetResourceTypeFromValue(v string) (*MonitorAssetResourceType, error) {
	ev := MonitorAssetResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorAssetResourceType: valid values are %v", v, allowedMonitorAssetResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorAssetResourceType) IsValid() bool {
	for _, existing := range allowedMonitorAssetResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorAssetResourceType value.
func (v MonitorAssetResourceType) Ptr() *MonitorAssetResourceType {
	return &v
}
