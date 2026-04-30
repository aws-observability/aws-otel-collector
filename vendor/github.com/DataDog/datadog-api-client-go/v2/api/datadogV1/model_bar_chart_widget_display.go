// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartWidgetDisplay - Bar chart widget display options.
type BarChartWidgetDisplay struct {
	BarChartWidgetStacked *BarChartWidgetStacked
	BarChartWidgetFlat    *BarChartWidgetFlat

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// BarChartWidgetStackedAsBarChartWidgetDisplay is a convenience function that returns BarChartWidgetStacked wrapped in BarChartWidgetDisplay.
func BarChartWidgetStackedAsBarChartWidgetDisplay(v *BarChartWidgetStacked) BarChartWidgetDisplay {
	return BarChartWidgetDisplay{BarChartWidgetStacked: v}
}

// BarChartWidgetFlatAsBarChartWidgetDisplay is a convenience function that returns BarChartWidgetFlat wrapped in BarChartWidgetDisplay.
func BarChartWidgetFlatAsBarChartWidgetDisplay(v *BarChartWidgetFlat) BarChartWidgetDisplay {
	return BarChartWidgetDisplay{BarChartWidgetFlat: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *BarChartWidgetDisplay) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BarChartWidgetStacked
	err = datadog.Unmarshal(data, &obj.BarChartWidgetStacked)
	if err == nil {
		if obj.BarChartWidgetStacked != nil && obj.BarChartWidgetStacked.UnparsedObject == nil {
			jsonBarChartWidgetStacked, _ := datadog.Marshal(obj.BarChartWidgetStacked)
			if string(jsonBarChartWidgetStacked) == "{}" { // empty struct
				obj.BarChartWidgetStacked = nil
			} else {
				match++
			}
		} else {
			obj.BarChartWidgetStacked = nil
		}
	} else {
		obj.BarChartWidgetStacked = nil
	}

	// try to unmarshal data into BarChartWidgetFlat
	err = datadog.Unmarshal(data, &obj.BarChartWidgetFlat)
	if err == nil {
		if obj.BarChartWidgetFlat != nil && obj.BarChartWidgetFlat.UnparsedObject == nil {
			jsonBarChartWidgetFlat, _ := datadog.Marshal(obj.BarChartWidgetFlat)
			if string(jsonBarChartWidgetFlat) == "{}" { // empty struct
				obj.BarChartWidgetFlat = nil
			} else {
				match++
			}
		} else {
			obj.BarChartWidgetFlat = nil
		}
	} else {
		obj.BarChartWidgetFlat = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.BarChartWidgetStacked = nil
		obj.BarChartWidgetFlat = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj BarChartWidgetDisplay) MarshalJSON() ([]byte, error) {
	if obj.BarChartWidgetStacked != nil {
		return datadog.Marshal(&obj.BarChartWidgetStacked)
	}

	if obj.BarChartWidgetFlat != nil {
		return datadog.Marshal(&obj.BarChartWidgetFlat)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *BarChartWidgetDisplay) GetActualInstance() interface{} {
	if obj.BarChartWidgetStacked != nil {
		return obj.BarChartWidgetStacked
	}

	if obj.BarChartWidgetFlat != nil {
		return obj.BarChartWidgetFlat
	}

	// all schemas are nil
	return nil
}
