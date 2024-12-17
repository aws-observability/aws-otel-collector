// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricVolumes - Possible response objects for a metric's volume.
type MetricVolumes struct {
	MetricDistinctVolume        *MetricDistinctVolume
	MetricIngestedIndexedVolume *MetricIngestedIndexedVolume

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MetricDistinctVolumeAsMetricVolumes is a convenience function that returns MetricDistinctVolume wrapped in MetricVolumes.
func MetricDistinctVolumeAsMetricVolumes(v *MetricDistinctVolume) MetricVolumes {
	return MetricVolumes{MetricDistinctVolume: v}
}

// MetricIngestedIndexedVolumeAsMetricVolumes is a convenience function that returns MetricIngestedIndexedVolume wrapped in MetricVolumes.
func MetricIngestedIndexedVolumeAsMetricVolumes(v *MetricIngestedIndexedVolume) MetricVolumes {
	return MetricVolumes{MetricIngestedIndexedVolume: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *MetricVolumes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MetricDistinctVolume
	err = datadog.Unmarshal(data, &obj.MetricDistinctVolume)
	if err == nil {
		if obj.MetricDistinctVolume != nil && obj.MetricDistinctVolume.UnparsedObject == nil {
			jsonMetricDistinctVolume, _ := datadog.Marshal(obj.MetricDistinctVolume)
			if string(jsonMetricDistinctVolume) == "{}" && string(data) != "{}" { // empty struct
				obj.MetricDistinctVolume = nil
			} else {
				match++
			}
		} else {
			obj.MetricDistinctVolume = nil
		}
	} else {
		obj.MetricDistinctVolume = nil
	}

	// try to unmarshal data into MetricIngestedIndexedVolume
	err = datadog.Unmarshal(data, &obj.MetricIngestedIndexedVolume)
	if err == nil {
		if obj.MetricIngestedIndexedVolume != nil && obj.MetricIngestedIndexedVolume.UnparsedObject == nil {
			jsonMetricIngestedIndexedVolume, _ := datadog.Marshal(obj.MetricIngestedIndexedVolume)
			if string(jsonMetricIngestedIndexedVolume) == "{}" && string(data) != "{}" { // empty struct
				obj.MetricIngestedIndexedVolume = nil
			} else {
				match++
			}
		} else {
			obj.MetricIngestedIndexedVolume = nil
		}
	} else {
		obj.MetricIngestedIndexedVolume = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MetricDistinctVolume = nil
		obj.MetricIngestedIndexedVolume = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj MetricVolumes) MarshalJSON() ([]byte, error) {
	if obj.MetricDistinctVolume != nil {
		return datadog.Marshal(&obj.MetricDistinctVolume)
	}

	if obj.MetricIngestedIndexedVolume != nil {
		return datadog.Marshal(&obj.MetricIngestedIndexedVolume)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *MetricVolumes) GetActualInstance() interface{} {
	if obj.MetricDistinctVolume != nil {
		return obj.MetricDistinctVolume
	}

	if obj.MetricIngestedIndexedVolume != nil {
		return obj.MetricIngestedIndexedVolume
	}

	// all schemas are nil
	return nil
}
