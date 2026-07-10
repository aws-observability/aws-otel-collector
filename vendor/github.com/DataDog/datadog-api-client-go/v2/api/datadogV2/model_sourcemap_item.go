// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SourcemapItem - A source map data object representing one of the supported map kinds.
type SourcemapItem struct {
	JSSourcemapData          *JSSourcemapData
	ReactNativeSourcemapData *ReactNativeSourcemapData
	IOSSourcemapData         *IOSSourcemapData
	JVMSourcemapData         *JVMSourcemapData
	FlutterSourcemapData     *FlutterSourcemapData
	ELFSourcemapData         *ELFSourcemapData
	NDKSourcemapData         *NDKSourcemapData
	IL2CPPSourcemapData      *IL2CPPSourcemapData

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// JSSourcemapDataAsSourcemapItem is a convenience function that returns JSSourcemapData wrapped in SourcemapItem.
func JSSourcemapDataAsSourcemapItem(v *JSSourcemapData) SourcemapItem {
	return SourcemapItem{JSSourcemapData: v}
}

// ReactNativeSourcemapDataAsSourcemapItem is a convenience function that returns ReactNativeSourcemapData wrapped in SourcemapItem.
func ReactNativeSourcemapDataAsSourcemapItem(v *ReactNativeSourcemapData) SourcemapItem {
	return SourcemapItem{ReactNativeSourcemapData: v}
}

// IOSSourcemapDataAsSourcemapItem is a convenience function that returns IOSSourcemapData wrapped in SourcemapItem.
func IOSSourcemapDataAsSourcemapItem(v *IOSSourcemapData) SourcemapItem {
	return SourcemapItem{IOSSourcemapData: v}
}

// JVMSourcemapDataAsSourcemapItem is a convenience function that returns JVMSourcemapData wrapped in SourcemapItem.
func JVMSourcemapDataAsSourcemapItem(v *JVMSourcemapData) SourcemapItem {
	return SourcemapItem{JVMSourcemapData: v}
}

// FlutterSourcemapDataAsSourcemapItem is a convenience function that returns FlutterSourcemapData wrapped in SourcemapItem.
func FlutterSourcemapDataAsSourcemapItem(v *FlutterSourcemapData) SourcemapItem {
	return SourcemapItem{FlutterSourcemapData: v}
}

// ELFSourcemapDataAsSourcemapItem is a convenience function that returns ELFSourcemapData wrapped in SourcemapItem.
func ELFSourcemapDataAsSourcemapItem(v *ELFSourcemapData) SourcemapItem {
	return SourcemapItem{ELFSourcemapData: v}
}

// NDKSourcemapDataAsSourcemapItem is a convenience function that returns NDKSourcemapData wrapped in SourcemapItem.
func NDKSourcemapDataAsSourcemapItem(v *NDKSourcemapData) SourcemapItem {
	return SourcemapItem{NDKSourcemapData: v}
}

// IL2CPPSourcemapDataAsSourcemapItem is a convenience function that returns IL2CPPSourcemapData wrapped in SourcemapItem.
func IL2CPPSourcemapDataAsSourcemapItem(v *IL2CPPSourcemapData) SourcemapItem {
	return SourcemapItem{IL2CPPSourcemapData: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SourcemapItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into JSSourcemapData
	err = datadog.Unmarshal(data, &obj.JSSourcemapData)
	if err == nil {
		if obj.JSSourcemapData != nil && obj.JSSourcemapData.UnparsedObject == nil {
			jsonJSSourcemapData, _ := datadog.Marshal(obj.JSSourcemapData)
			if string(jsonJSSourcemapData) == "{}" { // empty struct
				obj.JSSourcemapData = nil
			} else {
				match++
			}
		} else {
			obj.JSSourcemapData = nil
		}
	} else {
		obj.JSSourcemapData = nil
	}

	// try to unmarshal data into ReactNativeSourcemapData
	err = datadog.Unmarshal(data, &obj.ReactNativeSourcemapData)
	if err == nil {
		if obj.ReactNativeSourcemapData != nil && obj.ReactNativeSourcemapData.UnparsedObject == nil {
			jsonReactNativeSourcemapData, _ := datadog.Marshal(obj.ReactNativeSourcemapData)
			if string(jsonReactNativeSourcemapData) == "{}" { // empty struct
				obj.ReactNativeSourcemapData = nil
			} else {
				match++
			}
		} else {
			obj.ReactNativeSourcemapData = nil
		}
	} else {
		obj.ReactNativeSourcemapData = nil
	}

	// try to unmarshal data into IOSSourcemapData
	err = datadog.Unmarshal(data, &obj.IOSSourcemapData)
	if err == nil {
		if obj.IOSSourcemapData != nil && obj.IOSSourcemapData.UnparsedObject == nil {
			jsonIOSSourcemapData, _ := datadog.Marshal(obj.IOSSourcemapData)
			if string(jsonIOSSourcemapData) == "{}" { // empty struct
				obj.IOSSourcemapData = nil
			} else {
				match++
			}
		} else {
			obj.IOSSourcemapData = nil
		}
	} else {
		obj.IOSSourcemapData = nil
	}

	// try to unmarshal data into JVMSourcemapData
	err = datadog.Unmarshal(data, &obj.JVMSourcemapData)
	if err == nil {
		if obj.JVMSourcemapData != nil && obj.JVMSourcemapData.UnparsedObject == nil {
			jsonJVMSourcemapData, _ := datadog.Marshal(obj.JVMSourcemapData)
			if string(jsonJVMSourcemapData) == "{}" { // empty struct
				obj.JVMSourcemapData = nil
			} else {
				match++
			}
		} else {
			obj.JVMSourcemapData = nil
		}
	} else {
		obj.JVMSourcemapData = nil
	}

	// try to unmarshal data into FlutterSourcemapData
	err = datadog.Unmarshal(data, &obj.FlutterSourcemapData)
	if err == nil {
		if obj.FlutterSourcemapData != nil && obj.FlutterSourcemapData.UnparsedObject == nil {
			jsonFlutterSourcemapData, _ := datadog.Marshal(obj.FlutterSourcemapData)
			if string(jsonFlutterSourcemapData) == "{}" { // empty struct
				obj.FlutterSourcemapData = nil
			} else {
				match++
			}
		} else {
			obj.FlutterSourcemapData = nil
		}
	} else {
		obj.FlutterSourcemapData = nil
	}

	// try to unmarshal data into ELFSourcemapData
	err = datadog.Unmarshal(data, &obj.ELFSourcemapData)
	if err == nil {
		if obj.ELFSourcemapData != nil && obj.ELFSourcemapData.UnparsedObject == nil {
			jsonELFSourcemapData, _ := datadog.Marshal(obj.ELFSourcemapData)
			if string(jsonELFSourcemapData) == "{}" { // empty struct
				obj.ELFSourcemapData = nil
			} else {
				match++
			}
		} else {
			obj.ELFSourcemapData = nil
		}
	} else {
		obj.ELFSourcemapData = nil
	}

	// try to unmarshal data into NDKSourcemapData
	err = datadog.Unmarshal(data, &obj.NDKSourcemapData)
	if err == nil {
		if obj.NDKSourcemapData != nil && obj.NDKSourcemapData.UnparsedObject == nil {
			jsonNDKSourcemapData, _ := datadog.Marshal(obj.NDKSourcemapData)
			if string(jsonNDKSourcemapData) == "{}" { // empty struct
				obj.NDKSourcemapData = nil
			} else {
				match++
			}
		} else {
			obj.NDKSourcemapData = nil
		}
	} else {
		obj.NDKSourcemapData = nil
	}

	// try to unmarshal data into IL2CPPSourcemapData
	err = datadog.Unmarshal(data, &obj.IL2CPPSourcemapData)
	if err == nil {
		if obj.IL2CPPSourcemapData != nil && obj.IL2CPPSourcemapData.UnparsedObject == nil {
			jsonIL2CPPSourcemapData, _ := datadog.Marshal(obj.IL2CPPSourcemapData)
			if string(jsonIL2CPPSourcemapData) == "{}" { // empty struct
				obj.IL2CPPSourcemapData = nil
			} else {
				match++
			}
		} else {
			obj.IL2CPPSourcemapData = nil
		}
	} else {
		obj.IL2CPPSourcemapData = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.JSSourcemapData = nil
		obj.ReactNativeSourcemapData = nil
		obj.IOSSourcemapData = nil
		obj.JVMSourcemapData = nil
		obj.FlutterSourcemapData = nil
		obj.ELFSourcemapData = nil
		obj.NDKSourcemapData = nil
		obj.IL2CPPSourcemapData = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SourcemapItem) MarshalJSON() ([]byte, error) {
	if obj.JSSourcemapData != nil {
		return datadog.Marshal(&obj.JSSourcemapData)
	}

	if obj.ReactNativeSourcemapData != nil {
		return datadog.Marshal(&obj.ReactNativeSourcemapData)
	}

	if obj.IOSSourcemapData != nil {
		return datadog.Marshal(&obj.IOSSourcemapData)
	}

	if obj.JVMSourcemapData != nil {
		return datadog.Marshal(&obj.JVMSourcemapData)
	}

	if obj.FlutterSourcemapData != nil {
		return datadog.Marshal(&obj.FlutterSourcemapData)
	}

	if obj.ELFSourcemapData != nil {
		return datadog.Marshal(&obj.ELFSourcemapData)
	}

	if obj.NDKSourcemapData != nil {
		return datadog.Marshal(&obj.NDKSourcemapData)
	}

	if obj.IL2CPPSourcemapData != nil {
		return datadog.Marshal(&obj.IL2CPPSourcemapData)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SourcemapItem) GetActualInstance() interface{} {
	if obj.JSSourcemapData != nil {
		return obj.JSSourcemapData
	}

	if obj.ReactNativeSourcemapData != nil {
		return obj.ReactNativeSourcemapData
	}

	if obj.IOSSourcemapData != nil {
		return obj.IOSSourcemapData
	}

	if obj.JVMSourcemapData != nil {
		return obj.JVMSourcemapData
	}

	if obj.FlutterSourcemapData != nil {
		return obj.FlutterSourcemapData
	}

	if obj.ELFSourcemapData != nil {
		return obj.ELFSourcemapData
	}

	if obj.NDKSourcemapData != nil {
		return obj.NDKSourcemapData
	}

	if obj.IL2CPPSourcemapData != nil {
		return obj.IL2CPPSourcemapData
	}

	// all schemas are nil
	return nil
}
