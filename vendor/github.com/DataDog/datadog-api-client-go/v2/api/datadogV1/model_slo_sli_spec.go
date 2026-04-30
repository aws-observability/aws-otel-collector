// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOSliSpec - A generic SLI specification. This is used for time-slice and count-based (metric) SLOs only.
type SLOSliSpec struct {
	SLOTimeSliceSpec *SLOTimeSliceSpec
	SLOCountSpec     *SLOCountSpec

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SLOTimeSliceSpecAsSLOSliSpec is a convenience function that returns SLOTimeSliceSpec wrapped in SLOSliSpec.
func SLOTimeSliceSpecAsSLOSliSpec(v *SLOTimeSliceSpec) SLOSliSpec {
	return SLOSliSpec{SLOTimeSliceSpec: v}
}

// SLOCountSpecAsSLOSliSpec is a convenience function that returns SLOCountSpec wrapped in SLOSliSpec.
func SLOCountSpecAsSLOSliSpec(v *SLOCountSpec) SLOSliSpec {
	return SLOSliSpec{SLOCountSpec: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SLOSliSpec) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SLOTimeSliceSpec
	err = datadog.Unmarshal(data, &obj.SLOTimeSliceSpec)
	if err == nil {
		if obj.SLOTimeSliceSpec != nil && obj.SLOTimeSliceSpec.UnparsedObject == nil {
			jsonSLOTimeSliceSpec, _ := datadog.Marshal(obj.SLOTimeSliceSpec)
			if string(jsonSLOTimeSliceSpec) == "{}" { // empty struct
				obj.SLOTimeSliceSpec = nil
			} else {
				match++
			}
		} else {
			obj.SLOTimeSliceSpec = nil
		}
	} else {
		obj.SLOTimeSliceSpec = nil
	}

	// try to unmarshal data into SLOCountSpec
	err = datadog.Unmarshal(data, &obj.SLOCountSpec)
	if err == nil {
		if obj.SLOCountSpec != nil && obj.SLOCountSpec.UnparsedObject == nil {
			jsonSLOCountSpec, _ := datadog.Marshal(obj.SLOCountSpec)
			if string(jsonSLOCountSpec) == "{}" { // empty struct
				obj.SLOCountSpec = nil
			} else {
				match++
			}
		} else {
			obj.SLOCountSpec = nil
		}
	} else {
		obj.SLOCountSpec = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SLOTimeSliceSpec = nil
		obj.SLOCountSpec = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SLOSliSpec) MarshalJSON() ([]byte, error) {
	if obj.SLOTimeSliceSpec != nil {
		return datadog.Marshal(&obj.SLOTimeSliceSpec)
	}

	if obj.SLOCountSpec != nil {
		return datadog.Marshal(&obj.SLOCountSpec)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SLOSliSpec) GetActualInstance() interface{} {
	if obj.SLOTimeSliceSpec != nil {
		return obj.SLOTimeSliceSpec
	}

	if obj.SLOCountSpec != nil {
		return obj.SLOCountSpec
	}

	// all schemas are nil
	return nil
}
