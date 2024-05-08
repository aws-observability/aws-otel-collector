// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstWidgetLegend - Configuration of the legend.
type SunburstWidgetLegend struct {
	SunburstWidgetLegendTable           *SunburstWidgetLegendTable
	SunburstWidgetLegendInlineAutomatic *SunburstWidgetLegendInlineAutomatic

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SunburstWidgetLegendTableAsSunburstWidgetLegend is a convenience function that returns SunburstWidgetLegendTable wrapped in SunburstWidgetLegend.
func SunburstWidgetLegendTableAsSunburstWidgetLegend(v *SunburstWidgetLegendTable) SunburstWidgetLegend {
	return SunburstWidgetLegend{SunburstWidgetLegendTable: v}
}

// SunburstWidgetLegendInlineAutomaticAsSunburstWidgetLegend is a convenience function that returns SunburstWidgetLegendInlineAutomatic wrapped in SunburstWidgetLegend.
func SunburstWidgetLegendInlineAutomaticAsSunburstWidgetLegend(v *SunburstWidgetLegendInlineAutomatic) SunburstWidgetLegend {
	return SunburstWidgetLegend{SunburstWidgetLegendInlineAutomatic: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SunburstWidgetLegend) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SunburstWidgetLegendTable
	err = datadog.Unmarshal(data, &obj.SunburstWidgetLegendTable)
	if err == nil {
		if obj.SunburstWidgetLegendTable != nil && obj.SunburstWidgetLegendTable.UnparsedObject == nil {
			jsonSunburstWidgetLegendTable, _ := datadog.Marshal(obj.SunburstWidgetLegendTable)
			if string(jsonSunburstWidgetLegendTable) == "{}" { // empty struct
				obj.SunburstWidgetLegendTable = nil
			} else {
				match++
			}
		} else {
			obj.SunburstWidgetLegendTable = nil
		}
	} else {
		obj.SunburstWidgetLegendTable = nil
	}

	// try to unmarshal data into SunburstWidgetLegendInlineAutomatic
	err = datadog.Unmarshal(data, &obj.SunburstWidgetLegendInlineAutomatic)
	if err == nil {
		if obj.SunburstWidgetLegendInlineAutomatic != nil && obj.SunburstWidgetLegendInlineAutomatic.UnparsedObject == nil {
			jsonSunburstWidgetLegendInlineAutomatic, _ := datadog.Marshal(obj.SunburstWidgetLegendInlineAutomatic)
			if string(jsonSunburstWidgetLegendInlineAutomatic) == "{}" { // empty struct
				obj.SunburstWidgetLegendInlineAutomatic = nil
			} else {
				match++
			}
		} else {
			obj.SunburstWidgetLegendInlineAutomatic = nil
		}
	} else {
		obj.SunburstWidgetLegendInlineAutomatic = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SunburstWidgetLegendTable = nil
		obj.SunburstWidgetLegendInlineAutomatic = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SunburstWidgetLegend) MarshalJSON() ([]byte, error) {
	if obj.SunburstWidgetLegendTable != nil {
		return datadog.Marshal(&obj.SunburstWidgetLegendTable)
	}

	if obj.SunburstWidgetLegendInlineAutomatic != nil {
		return datadog.Marshal(&obj.SunburstWidgetLegendInlineAutomatic)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SunburstWidgetLegend) GetActualInstance() interface{} {
	if obj.SunburstWidgetLegendTable != nil {
		return obj.SunburstWidgetLegendTable
	}

	if obj.SunburstWidgetLegendInlineAutomatic != nil {
		return obj.SunburstWidgetLegendInlineAutomatic
	}

	// all schemas are nil
	return nil
}
