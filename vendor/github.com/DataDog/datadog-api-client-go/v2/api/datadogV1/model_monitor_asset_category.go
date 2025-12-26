// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorAssetCategory Indicates the type of asset this entity represents on a monitor.
type MonitorAssetCategory string

// List of MonitorAssetCategory.
const (
	MONITORASSETCATEGORY_RUNBOOK MonitorAssetCategory = "runbook"
)

var allowedMonitorAssetCategoryEnumValues = []MonitorAssetCategory{
	MONITORASSETCATEGORY_RUNBOOK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorAssetCategory) GetAllowedValues() []MonitorAssetCategory {
	return allowedMonitorAssetCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorAssetCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorAssetCategory(value)
	return nil
}

// NewMonitorAssetCategoryFromValue returns a pointer to a valid MonitorAssetCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorAssetCategoryFromValue(v string) (*MonitorAssetCategory, error) {
	ev := MonitorAssetCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorAssetCategory: valid values are %v", v, allowedMonitorAssetCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorAssetCategory) IsValid() bool {
	for _, existing := range allowedMonitorAssetCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorAssetCategory value.
func (v MonitorAssetCategory) Ptr() *MonitorAssetCategory {
	return &v
}
