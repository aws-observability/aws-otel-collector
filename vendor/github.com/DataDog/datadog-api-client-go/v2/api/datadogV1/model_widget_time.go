// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetTime - Time setting for the widget.
type WidgetTime struct {
	WidgetLegacyLiveSpan *WidgetLegacyLiveSpan
	WidgetNewLiveSpan    *WidgetNewLiveSpan
	WidgetNewFixedSpan   *WidgetNewFixedSpan

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// WidgetLegacyLiveSpanAsWidgetTime is a convenience function that returns WidgetLegacyLiveSpan wrapped in WidgetTime.
func WidgetLegacyLiveSpanAsWidgetTime(v *WidgetLegacyLiveSpan) WidgetTime {
	return WidgetTime{WidgetLegacyLiveSpan: v}
}

// WidgetNewLiveSpanAsWidgetTime is a convenience function that returns WidgetNewLiveSpan wrapped in WidgetTime.
func WidgetNewLiveSpanAsWidgetTime(v *WidgetNewLiveSpan) WidgetTime {
	return WidgetTime{WidgetNewLiveSpan: v}
}

// WidgetNewFixedSpanAsWidgetTime is a convenience function that returns WidgetNewFixedSpan wrapped in WidgetTime.
func WidgetNewFixedSpanAsWidgetTime(v *WidgetNewFixedSpan) WidgetTime {
	return WidgetTime{WidgetNewFixedSpan: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *WidgetTime) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into WidgetLegacyLiveSpan
	err = datadog.Unmarshal(data, &obj.WidgetLegacyLiveSpan)
	if err == nil {
		if obj.WidgetLegacyLiveSpan != nil && obj.WidgetLegacyLiveSpan.UnparsedObject == nil {
			jsonWidgetLegacyLiveSpan, _ := datadog.Marshal(obj.WidgetLegacyLiveSpan)
			if string(jsonWidgetLegacyLiveSpan) == "{}" && string(data) != "{}" { // empty struct
				obj.WidgetLegacyLiveSpan = nil
			} else {
				match++
			}
		} else {
			obj.WidgetLegacyLiveSpan = nil
		}
	} else {
		obj.WidgetLegacyLiveSpan = nil
	}

	// try to unmarshal data into WidgetNewLiveSpan
	err = datadog.Unmarshal(data, &obj.WidgetNewLiveSpan)
	if err == nil {
		if obj.WidgetNewLiveSpan != nil && obj.WidgetNewLiveSpan.UnparsedObject == nil {
			jsonWidgetNewLiveSpan, _ := datadog.Marshal(obj.WidgetNewLiveSpan)
			if string(jsonWidgetNewLiveSpan) == "{}" { // empty struct
				obj.WidgetNewLiveSpan = nil
			} else {
				match++
			}
		} else {
			obj.WidgetNewLiveSpan = nil
		}
	} else {
		obj.WidgetNewLiveSpan = nil
	}

	// try to unmarshal data into WidgetNewFixedSpan
	err = datadog.Unmarshal(data, &obj.WidgetNewFixedSpan)
	if err == nil {
		if obj.WidgetNewFixedSpan != nil && obj.WidgetNewFixedSpan.UnparsedObject == nil {
			jsonWidgetNewFixedSpan, _ := datadog.Marshal(obj.WidgetNewFixedSpan)
			if string(jsonWidgetNewFixedSpan) == "{}" { // empty struct
				obj.WidgetNewFixedSpan = nil
			} else {
				match++
			}
		} else {
			obj.WidgetNewFixedSpan = nil
		}
	} else {
		obj.WidgetNewFixedSpan = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.WidgetLegacyLiveSpan = nil
		obj.WidgetNewLiveSpan = nil
		obj.WidgetNewFixedSpan = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj WidgetTime) MarshalJSON() ([]byte, error) {
	if obj.WidgetLegacyLiveSpan != nil {
		return datadog.Marshal(&obj.WidgetLegacyLiveSpan)
	}

	if obj.WidgetNewLiveSpan != nil {
		return datadog.Marshal(&obj.WidgetNewLiveSpan)
	}

	if obj.WidgetNewFixedSpan != nil {
		return datadog.Marshal(&obj.WidgetNewFixedSpan)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *WidgetTime) GetActualInstance() interface{} {
	if obj.WidgetLegacyLiveSpan != nil {
		return obj.WidgetLegacyLiveSpan
	}

	if obj.WidgetNewLiveSpan != nil {
		return obj.WidgetNewLiveSpan
	}

	if obj.WidgetNewFixedSpan != nil {
		return obj.WidgetNewFixedSpan
	}

	// all schemas are nil
	return nil
}
