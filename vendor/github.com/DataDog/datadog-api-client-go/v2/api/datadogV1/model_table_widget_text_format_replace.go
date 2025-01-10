// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableWidgetTextFormatReplace - Replace rule for the table widget text format.
type TableWidgetTextFormatReplace struct {
	TableWidgetTextFormatReplaceAll       *TableWidgetTextFormatReplaceAll
	TableWidgetTextFormatReplaceSubstring *TableWidgetTextFormatReplaceSubstring

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TableWidgetTextFormatReplaceAllAsTableWidgetTextFormatReplace is a convenience function that returns TableWidgetTextFormatReplaceAll wrapped in TableWidgetTextFormatReplace.
func TableWidgetTextFormatReplaceAllAsTableWidgetTextFormatReplace(v *TableWidgetTextFormatReplaceAll) TableWidgetTextFormatReplace {
	return TableWidgetTextFormatReplace{TableWidgetTextFormatReplaceAll: v}
}

// TableWidgetTextFormatReplaceSubstringAsTableWidgetTextFormatReplace is a convenience function that returns TableWidgetTextFormatReplaceSubstring wrapped in TableWidgetTextFormatReplace.
func TableWidgetTextFormatReplaceSubstringAsTableWidgetTextFormatReplace(v *TableWidgetTextFormatReplaceSubstring) TableWidgetTextFormatReplace {
	return TableWidgetTextFormatReplace{TableWidgetTextFormatReplaceSubstring: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TableWidgetTextFormatReplace) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TableWidgetTextFormatReplaceAll
	err = datadog.Unmarshal(data, &obj.TableWidgetTextFormatReplaceAll)
	if err == nil {
		if obj.TableWidgetTextFormatReplaceAll != nil && obj.TableWidgetTextFormatReplaceAll.UnparsedObject == nil {
			jsonTableWidgetTextFormatReplaceAll, _ := datadog.Marshal(obj.TableWidgetTextFormatReplaceAll)
			if string(jsonTableWidgetTextFormatReplaceAll) == "{}" { // empty struct
				obj.TableWidgetTextFormatReplaceAll = nil
			} else {
				match++
			}
		} else {
			obj.TableWidgetTextFormatReplaceAll = nil
		}
	} else {
		obj.TableWidgetTextFormatReplaceAll = nil
	}

	// try to unmarshal data into TableWidgetTextFormatReplaceSubstring
	err = datadog.Unmarshal(data, &obj.TableWidgetTextFormatReplaceSubstring)
	if err == nil {
		if obj.TableWidgetTextFormatReplaceSubstring != nil && obj.TableWidgetTextFormatReplaceSubstring.UnparsedObject == nil {
			jsonTableWidgetTextFormatReplaceSubstring, _ := datadog.Marshal(obj.TableWidgetTextFormatReplaceSubstring)
			if string(jsonTableWidgetTextFormatReplaceSubstring) == "{}" { // empty struct
				obj.TableWidgetTextFormatReplaceSubstring = nil
			} else {
				match++
			}
		} else {
			obj.TableWidgetTextFormatReplaceSubstring = nil
		}
	} else {
		obj.TableWidgetTextFormatReplaceSubstring = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TableWidgetTextFormatReplaceAll = nil
		obj.TableWidgetTextFormatReplaceSubstring = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TableWidgetTextFormatReplace) MarshalJSON() ([]byte, error) {
	if obj.TableWidgetTextFormatReplaceAll != nil {
		return datadog.Marshal(&obj.TableWidgetTextFormatReplaceAll)
	}

	if obj.TableWidgetTextFormatReplaceSubstring != nil {
		return datadog.Marshal(&obj.TableWidgetTextFormatReplaceSubstring)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TableWidgetTextFormatReplace) GetActualInstance() interface{} {
	if obj.TableWidgetTextFormatReplaceAll != nil {
		return obj.TableWidgetTextFormatReplaceAll
	}

	if obj.TableWidgetTextFormatReplaceSubstring != nil {
		return obj.TableWidgetTextFormatReplaceSubstring
	}

	// all schemas are nil
	return nil
}
