// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WildcardWidgetRequest - Request object for the wildcard widget. Each variant represents a distinct data-fetching pattern: scalar formulas, timeseries formulas, list streams, and histograms.
type WildcardWidgetRequest struct {
	TreeMapWidgetRequest      *TreeMapWidgetRequest
	TimeseriesWidgetRequest   *TimeseriesWidgetRequest
	ListStreamWidgetRequest   *ListStreamWidgetRequest
	DistributionWidgetRequest *DistributionWidgetRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TreeMapWidgetRequestAsWildcardWidgetRequest is a convenience function that returns TreeMapWidgetRequest wrapped in WildcardWidgetRequest.
func TreeMapWidgetRequestAsWildcardWidgetRequest(v *TreeMapWidgetRequest) WildcardWidgetRequest {
	return WildcardWidgetRequest{TreeMapWidgetRequest: v}
}

// TimeseriesWidgetRequestAsWildcardWidgetRequest is a convenience function that returns TimeseriesWidgetRequest wrapped in WildcardWidgetRequest.
func TimeseriesWidgetRequestAsWildcardWidgetRequest(v *TimeseriesWidgetRequest) WildcardWidgetRequest {
	return WildcardWidgetRequest{TimeseriesWidgetRequest: v}
}

// ListStreamWidgetRequestAsWildcardWidgetRequest is a convenience function that returns ListStreamWidgetRequest wrapped in WildcardWidgetRequest.
func ListStreamWidgetRequestAsWildcardWidgetRequest(v *ListStreamWidgetRequest) WildcardWidgetRequest {
	return WildcardWidgetRequest{ListStreamWidgetRequest: v}
}

// DistributionWidgetRequestAsWildcardWidgetRequest is a convenience function that returns DistributionWidgetRequest wrapped in WildcardWidgetRequest.
func DistributionWidgetRequestAsWildcardWidgetRequest(v *DistributionWidgetRequest) WildcardWidgetRequest {
	return WildcardWidgetRequest{DistributionWidgetRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *WildcardWidgetRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TreeMapWidgetRequest
	err = datadog.Unmarshal(data, &obj.TreeMapWidgetRequest)
	if err == nil {
		if obj.TreeMapWidgetRequest != nil && obj.TreeMapWidgetRequest.UnparsedObject == nil {
			jsonTreeMapWidgetRequest, _ := datadog.Marshal(obj.TreeMapWidgetRequest)
			if string(jsonTreeMapWidgetRequest) == "{}" && string(data) != "{}" { // empty struct
				obj.TreeMapWidgetRequest = nil
			} else {
				match++
			}
		} else {
			obj.TreeMapWidgetRequest = nil
		}
	} else {
		obj.TreeMapWidgetRequest = nil
	}

	// try to unmarshal data into TimeseriesWidgetRequest
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetRequest)
	if err == nil {
		if obj.TimeseriesWidgetRequest != nil && obj.TimeseriesWidgetRequest.UnparsedObject == nil {
			jsonTimeseriesWidgetRequest, _ := datadog.Marshal(obj.TimeseriesWidgetRequest)
			if string(jsonTimeseriesWidgetRequest) == "{}" && string(data) != "{}" { // empty struct
				obj.TimeseriesWidgetRequest = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetRequest = nil
		}
	} else {
		obj.TimeseriesWidgetRequest = nil
	}

	// try to unmarshal data into ListStreamWidgetRequest
	err = datadog.Unmarshal(data, &obj.ListStreamWidgetRequest)
	if err == nil {
		if obj.ListStreamWidgetRequest != nil && obj.ListStreamWidgetRequest.UnparsedObject == nil {
			jsonListStreamWidgetRequest, _ := datadog.Marshal(obj.ListStreamWidgetRequest)
			if string(jsonListStreamWidgetRequest) == "{}" { // empty struct
				obj.ListStreamWidgetRequest = nil
			} else {
				match++
			}
		} else {
			obj.ListStreamWidgetRequest = nil
		}
	} else {
		obj.ListStreamWidgetRequest = nil
	}

	// try to unmarshal data into DistributionWidgetRequest
	err = datadog.Unmarshal(data, &obj.DistributionWidgetRequest)
	if err == nil {
		if obj.DistributionWidgetRequest != nil && obj.DistributionWidgetRequest.UnparsedObject == nil {
			jsonDistributionWidgetRequest, _ := datadog.Marshal(obj.DistributionWidgetRequest)
			if string(jsonDistributionWidgetRequest) == "{}" && string(data) != "{}" { // empty struct
				obj.DistributionWidgetRequest = nil
			} else {
				match++
			}
		} else {
			obj.DistributionWidgetRequest = nil
		}
	} else {
		obj.DistributionWidgetRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TreeMapWidgetRequest = nil
		obj.TimeseriesWidgetRequest = nil
		obj.ListStreamWidgetRequest = nil
		obj.DistributionWidgetRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj WildcardWidgetRequest) MarshalJSON() ([]byte, error) {
	if obj.TreeMapWidgetRequest != nil {
		return datadog.Marshal(&obj.TreeMapWidgetRequest)
	}

	if obj.TimeseriesWidgetRequest != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetRequest)
	}

	if obj.ListStreamWidgetRequest != nil {
		return datadog.Marshal(&obj.ListStreamWidgetRequest)
	}

	if obj.DistributionWidgetRequest != nil {
		return datadog.Marshal(&obj.DistributionWidgetRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *WildcardWidgetRequest) GetActualInstance() interface{} {
	if obj.TreeMapWidgetRequest != nil {
		return obj.TreeMapWidgetRequest
	}

	if obj.TimeseriesWidgetRequest != nil {
		return obj.TimeseriesWidgetRequest
	}

	if obj.ListStreamWidgetRequest != nil {
		return obj.ListStreamWidgetRequest
	}

	if obj.DistributionWidgetRequest != nil {
		return obj.DistributionWidgetRequest
	}

	// all schemas are nil
	return nil
}
