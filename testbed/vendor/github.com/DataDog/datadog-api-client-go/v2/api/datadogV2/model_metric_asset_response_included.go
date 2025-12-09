// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricAssetResponseIncluded - List of included assets with full set of attributes.
type MetricAssetResponseIncluded struct {
	MetricDashboardAsset *MetricDashboardAsset
	MetricMonitorAsset   *MetricMonitorAsset
	MetricNotebookAsset  *MetricNotebookAsset
	MetricSLOAsset       *MetricSLOAsset

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MetricDashboardAssetAsMetricAssetResponseIncluded is a convenience function that returns MetricDashboardAsset wrapped in MetricAssetResponseIncluded.
func MetricDashboardAssetAsMetricAssetResponseIncluded(v *MetricDashboardAsset) MetricAssetResponseIncluded {
	return MetricAssetResponseIncluded{MetricDashboardAsset: v}
}

// MetricMonitorAssetAsMetricAssetResponseIncluded is a convenience function that returns MetricMonitorAsset wrapped in MetricAssetResponseIncluded.
func MetricMonitorAssetAsMetricAssetResponseIncluded(v *MetricMonitorAsset) MetricAssetResponseIncluded {
	return MetricAssetResponseIncluded{MetricMonitorAsset: v}
}

// MetricNotebookAssetAsMetricAssetResponseIncluded is a convenience function that returns MetricNotebookAsset wrapped in MetricAssetResponseIncluded.
func MetricNotebookAssetAsMetricAssetResponseIncluded(v *MetricNotebookAsset) MetricAssetResponseIncluded {
	return MetricAssetResponseIncluded{MetricNotebookAsset: v}
}

// MetricSLOAssetAsMetricAssetResponseIncluded is a convenience function that returns MetricSLOAsset wrapped in MetricAssetResponseIncluded.
func MetricSLOAssetAsMetricAssetResponseIncluded(v *MetricSLOAsset) MetricAssetResponseIncluded {
	return MetricAssetResponseIncluded{MetricSLOAsset: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *MetricAssetResponseIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MetricDashboardAsset
	err = datadog.Unmarshal(data, &obj.MetricDashboardAsset)
	if err == nil {
		if obj.MetricDashboardAsset != nil && obj.MetricDashboardAsset.UnparsedObject == nil {
			jsonMetricDashboardAsset, _ := datadog.Marshal(obj.MetricDashboardAsset)
			if string(jsonMetricDashboardAsset) == "{}" { // empty struct
				obj.MetricDashboardAsset = nil
			} else {
				match++
			}
		} else {
			obj.MetricDashboardAsset = nil
		}
	} else {
		obj.MetricDashboardAsset = nil
	}

	// try to unmarshal data into MetricMonitorAsset
	err = datadog.Unmarshal(data, &obj.MetricMonitorAsset)
	if err == nil {
		if obj.MetricMonitorAsset != nil && obj.MetricMonitorAsset.UnparsedObject == nil {
			jsonMetricMonitorAsset, _ := datadog.Marshal(obj.MetricMonitorAsset)
			if string(jsonMetricMonitorAsset) == "{}" { // empty struct
				obj.MetricMonitorAsset = nil
			} else {
				match++
			}
		} else {
			obj.MetricMonitorAsset = nil
		}
	} else {
		obj.MetricMonitorAsset = nil
	}

	// try to unmarshal data into MetricNotebookAsset
	err = datadog.Unmarshal(data, &obj.MetricNotebookAsset)
	if err == nil {
		if obj.MetricNotebookAsset != nil && obj.MetricNotebookAsset.UnparsedObject == nil {
			jsonMetricNotebookAsset, _ := datadog.Marshal(obj.MetricNotebookAsset)
			if string(jsonMetricNotebookAsset) == "{}" { // empty struct
				obj.MetricNotebookAsset = nil
			} else {
				match++
			}
		} else {
			obj.MetricNotebookAsset = nil
		}
	} else {
		obj.MetricNotebookAsset = nil
	}

	// try to unmarshal data into MetricSLOAsset
	err = datadog.Unmarshal(data, &obj.MetricSLOAsset)
	if err == nil {
		if obj.MetricSLOAsset != nil && obj.MetricSLOAsset.UnparsedObject == nil {
			jsonMetricSLOAsset, _ := datadog.Marshal(obj.MetricSLOAsset)
			if string(jsonMetricSLOAsset) == "{}" { // empty struct
				obj.MetricSLOAsset = nil
			} else {
				match++
			}
		} else {
			obj.MetricSLOAsset = nil
		}
	} else {
		obj.MetricSLOAsset = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MetricDashboardAsset = nil
		obj.MetricMonitorAsset = nil
		obj.MetricNotebookAsset = nil
		obj.MetricSLOAsset = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj MetricAssetResponseIncluded) MarshalJSON() ([]byte, error) {
	if obj.MetricDashboardAsset != nil {
		return datadog.Marshal(&obj.MetricDashboardAsset)
	}

	if obj.MetricMonitorAsset != nil {
		return datadog.Marshal(&obj.MetricMonitorAsset)
	}

	if obj.MetricNotebookAsset != nil {
		return datadog.Marshal(&obj.MetricNotebookAsset)
	}

	if obj.MetricSLOAsset != nil {
		return datadog.Marshal(&obj.MetricSLOAsset)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *MetricAssetResponseIncluded) GetActualInstance() interface{} {
	if obj.MetricDashboardAsset != nil {
		return obj.MetricDashboardAsset
	}

	if obj.MetricMonitorAsset != nil {
		return obj.MetricMonitorAsset
	}

	if obj.MetricNotebookAsset != nil {
		return obj.MetricNotebookAsset
	}

	if obj.MetricSLOAsset != nil {
		return obj.MetricSLOAsset
	}

	// all schemas are nil
	return nil
}
