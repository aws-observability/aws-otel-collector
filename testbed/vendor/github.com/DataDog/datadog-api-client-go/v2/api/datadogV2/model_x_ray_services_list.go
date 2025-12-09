// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// XRayServicesList - AWS X-Ray services to collect traces from. Defaults to `include_only`.
type XRayServicesList struct {
	XRayServicesIncludeAll  *XRayServicesIncludeAll
	XRayServicesIncludeOnly *XRayServicesIncludeOnly

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// XRayServicesIncludeAllAsXRayServicesList is a convenience function that returns XRayServicesIncludeAll wrapped in XRayServicesList.
func XRayServicesIncludeAllAsXRayServicesList(v *XRayServicesIncludeAll) XRayServicesList {
	return XRayServicesList{XRayServicesIncludeAll: v}
}

// XRayServicesIncludeOnlyAsXRayServicesList is a convenience function that returns XRayServicesIncludeOnly wrapped in XRayServicesList.
func XRayServicesIncludeOnlyAsXRayServicesList(v *XRayServicesIncludeOnly) XRayServicesList {
	return XRayServicesList{XRayServicesIncludeOnly: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *XRayServicesList) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into XRayServicesIncludeAll
	err = datadog.Unmarshal(data, &obj.XRayServicesIncludeAll)
	if err == nil {
		if obj.XRayServicesIncludeAll != nil && obj.XRayServicesIncludeAll.UnparsedObject == nil {
			jsonXRayServicesIncludeAll, _ := datadog.Marshal(obj.XRayServicesIncludeAll)
			if string(jsonXRayServicesIncludeAll) == "{}" { // empty struct
				obj.XRayServicesIncludeAll = nil
			} else {
				match++
			}
		} else {
			obj.XRayServicesIncludeAll = nil
		}
	} else {
		obj.XRayServicesIncludeAll = nil
	}

	// try to unmarshal data into XRayServicesIncludeOnly
	err = datadog.Unmarshal(data, &obj.XRayServicesIncludeOnly)
	if err == nil {
		if obj.XRayServicesIncludeOnly != nil && obj.XRayServicesIncludeOnly.UnparsedObject == nil {
			jsonXRayServicesIncludeOnly, _ := datadog.Marshal(obj.XRayServicesIncludeOnly)
			if string(jsonXRayServicesIncludeOnly) == "{}" { // empty struct
				obj.XRayServicesIncludeOnly = nil
			} else {
				match++
			}
		} else {
			obj.XRayServicesIncludeOnly = nil
		}
	} else {
		obj.XRayServicesIncludeOnly = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.XRayServicesIncludeAll = nil
		obj.XRayServicesIncludeOnly = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj XRayServicesList) MarshalJSON() ([]byte, error) {
	if obj.XRayServicesIncludeAll != nil {
		return datadog.Marshal(&obj.XRayServicesIncludeAll)
	}

	if obj.XRayServicesIncludeOnly != nil {
		return datadog.Marshal(&obj.XRayServicesIncludeOnly)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *XRayServicesList) GetActualInstance() interface{} {
	if obj.XRayServicesIncludeAll != nil {
		return obj.XRayServicesIncludeAll
	}

	if obj.XRayServicesIncludeOnly != nil {
		return obj.XRayServicesIncludeOnly
	}

	// all schemas are nil
	return nil
}
