// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyWidgetRequest - Request definition for Sankey widget.
type SankeyWidgetRequest struct {
	SankeyRumRequest     *SankeyRumRequest
	SankeyNetworkRequest *SankeyNetworkRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SankeyRumRequestAsSankeyWidgetRequest is a convenience function that returns SankeyRumRequest wrapped in SankeyWidgetRequest.
func SankeyRumRequestAsSankeyWidgetRequest(v *SankeyRumRequest) SankeyWidgetRequest {
	return SankeyWidgetRequest{SankeyRumRequest: v}
}

// SankeyNetworkRequestAsSankeyWidgetRequest is a convenience function that returns SankeyNetworkRequest wrapped in SankeyWidgetRequest.
func SankeyNetworkRequestAsSankeyWidgetRequest(v *SankeyNetworkRequest) SankeyWidgetRequest {
	return SankeyWidgetRequest{SankeyNetworkRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SankeyWidgetRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SankeyRumRequest
	err = datadog.Unmarshal(data, &obj.SankeyRumRequest)
	if err == nil {
		if obj.SankeyRumRequest != nil && obj.SankeyRumRequest.UnparsedObject == nil {
			jsonSankeyRumRequest, _ := datadog.Marshal(obj.SankeyRumRequest)
			if string(jsonSankeyRumRequest) == "{}" { // empty struct
				obj.SankeyRumRequest = nil
			} else {
				match++
			}
		} else {
			obj.SankeyRumRequest = nil
		}
	} else {
		obj.SankeyRumRequest = nil
	}

	// try to unmarshal data into SankeyNetworkRequest
	err = datadog.Unmarshal(data, &obj.SankeyNetworkRequest)
	if err == nil {
		if obj.SankeyNetworkRequest != nil && obj.SankeyNetworkRequest.UnparsedObject == nil {
			jsonSankeyNetworkRequest, _ := datadog.Marshal(obj.SankeyNetworkRequest)
			if string(jsonSankeyNetworkRequest) == "{}" { // empty struct
				obj.SankeyNetworkRequest = nil
			} else {
				match++
			}
		} else {
			obj.SankeyNetworkRequest = nil
		}
	} else {
		obj.SankeyNetworkRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SankeyRumRequest = nil
		obj.SankeyNetworkRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SankeyWidgetRequest) MarshalJSON() ([]byte, error) {
	if obj.SankeyRumRequest != nil {
		return datadog.Marshal(&obj.SankeyRumRequest)
	}

	if obj.SankeyNetworkRequest != nil {
		return datadog.Marshal(&obj.SankeyNetworkRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SankeyWidgetRequest) GetActualInstance() interface{} {
	if obj.SankeyRumRequest != nil {
		return obj.SankeyRumRequest
	}

	if obj.SankeyNetworkRequest != nil {
		return obj.SankeyNetworkRequest
	}

	// all schemas are nil
	return nil
}
