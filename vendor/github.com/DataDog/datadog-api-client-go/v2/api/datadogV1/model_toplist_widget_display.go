// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistWidgetDisplay - Top list widget display options.
type ToplistWidgetDisplay struct {
	ToplistWidgetStacked *ToplistWidgetStacked
	ToplistWidgetFlat    *ToplistWidgetFlat

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ToplistWidgetStackedAsToplistWidgetDisplay is a convenience function that returns ToplistWidgetStacked wrapped in ToplistWidgetDisplay.
func ToplistWidgetStackedAsToplistWidgetDisplay(v *ToplistWidgetStacked) ToplistWidgetDisplay {
	return ToplistWidgetDisplay{ToplistWidgetStacked: v}
}

// ToplistWidgetFlatAsToplistWidgetDisplay is a convenience function that returns ToplistWidgetFlat wrapped in ToplistWidgetDisplay.
func ToplistWidgetFlatAsToplistWidgetDisplay(v *ToplistWidgetFlat) ToplistWidgetDisplay {
	return ToplistWidgetDisplay{ToplistWidgetFlat: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ToplistWidgetDisplay) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ToplistWidgetStacked
	err = datadog.Unmarshal(data, &obj.ToplistWidgetStacked)
	if err == nil {
		if obj.ToplistWidgetStacked != nil && obj.ToplistWidgetStacked.UnparsedObject == nil {
			jsonToplistWidgetStacked, _ := datadog.Marshal(obj.ToplistWidgetStacked)
			if string(jsonToplistWidgetStacked) == "{}" { // empty struct
				obj.ToplistWidgetStacked = nil
			} else {
				match++
			}
		} else {
			obj.ToplistWidgetStacked = nil
		}
	} else {
		obj.ToplistWidgetStacked = nil
	}

	// try to unmarshal data into ToplistWidgetFlat
	err = datadog.Unmarshal(data, &obj.ToplistWidgetFlat)
	if err == nil {
		if obj.ToplistWidgetFlat != nil && obj.ToplistWidgetFlat.UnparsedObject == nil {
			jsonToplistWidgetFlat, _ := datadog.Marshal(obj.ToplistWidgetFlat)
			if string(jsonToplistWidgetFlat) == "{}" { // empty struct
				obj.ToplistWidgetFlat = nil
			} else {
				match++
			}
		} else {
			obj.ToplistWidgetFlat = nil
		}
	} else {
		obj.ToplistWidgetFlat = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ToplistWidgetStacked = nil
		obj.ToplistWidgetFlat = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ToplistWidgetDisplay) MarshalJSON() ([]byte, error) {
	if obj.ToplistWidgetStacked != nil {
		return datadog.Marshal(&obj.ToplistWidgetStacked)
	}

	if obj.ToplistWidgetFlat != nil {
		return datadog.Marshal(&obj.ToplistWidgetFlat)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ToplistWidgetDisplay) GetActualInstance() interface{} {
	if obj.ToplistWidgetStacked != nil {
		return obj.ToplistWidgetStacked
	}

	if obj.ToplistWidgetFlat != nil {
		return obj.ToplistWidgetFlat
	}

	// all schemas are nil
	return nil
}
