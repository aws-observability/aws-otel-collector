// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberFormatUnit - Number format unit.
type NumberFormatUnit struct {
	NumberFormatUnitCanonical *NumberFormatUnitCanonical
	NumberFormatUnitCustom    *NumberFormatUnitCustom

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NumberFormatUnitCanonicalAsNumberFormatUnit is a convenience function that returns NumberFormatUnitCanonical wrapped in NumberFormatUnit.
func NumberFormatUnitCanonicalAsNumberFormatUnit(v *NumberFormatUnitCanonical) NumberFormatUnit {
	return NumberFormatUnit{NumberFormatUnitCanonical: v}
}

// NumberFormatUnitCustomAsNumberFormatUnit is a convenience function that returns NumberFormatUnitCustom wrapped in NumberFormatUnit.
func NumberFormatUnitCustomAsNumberFormatUnit(v *NumberFormatUnitCustom) NumberFormatUnit {
	return NumberFormatUnit{NumberFormatUnitCustom: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *NumberFormatUnit) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NumberFormatUnitCanonical
	err = datadog.Unmarshal(data, &obj.NumberFormatUnitCanonical)
	if err == nil {
		if obj.NumberFormatUnitCanonical != nil && obj.NumberFormatUnitCanonical.UnparsedObject == nil {
			jsonNumberFormatUnitCanonical, _ := datadog.Marshal(obj.NumberFormatUnitCanonical)
			if string(jsonNumberFormatUnitCanonical) == "{}" && string(data) != "{}" { // empty struct
				obj.NumberFormatUnitCanonical = nil
			} else {
				match++
			}
		} else {
			obj.NumberFormatUnitCanonical = nil
		}
	} else {
		obj.NumberFormatUnitCanonical = nil
	}

	// try to unmarshal data into NumberFormatUnitCustom
	err = datadog.Unmarshal(data, &obj.NumberFormatUnitCustom)
	if err == nil {
		if obj.NumberFormatUnitCustom != nil && obj.NumberFormatUnitCustom.UnparsedObject == nil {
			jsonNumberFormatUnitCustom, _ := datadog.Marshal(obj.NumberFormatUnitCustom)
			if string(jsonNumberFormatUnitCustom) == "{}" && string(data) != "{}" { // empty struct
				obj.NumberFormatUnitCustom = nil
			} else {
				match++
			}
		} else {
			obj.NumberFormatUnitCustom = nil
		}
	} else {
		obj.NumberFormatUnitCustom = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.NumberFormatUnitCanonical = nil
		obj.NumberFormatUnitCustom = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj NumberFormatUnit) MarshalJSON() ([]byte, error) {
	if obj.NumberFormatUnitCanonical != nil {
		return datadog.Marshal(&obj.NumberFormatUnitCanonical)
	}

	if obj.NumberFormatUnitCustom != nil {
		return datadog.Marshal(&obj.NumberFormatUnitCustom)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *NumberFormatUnit) GetActualInstance() interface{} {
	if obj.NumberFormatUnitCanonical != nil {
		return obj.NumberFormatUnitCanonical
	}

	if obj.NumberFormatUnitCustom != nil {
		return obj.NumberFormatUnitCustom
	}

	// all schemas are nil
	return nil
}
