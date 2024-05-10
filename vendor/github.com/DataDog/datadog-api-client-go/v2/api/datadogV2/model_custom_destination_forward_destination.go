// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationForwardDestination - A custom destination's location to forward logs.
type CustomDestinationForwardDestination struct {
	CustomDestinationForwardDestinationHttp          *CustomDestinationForwardDestinationHttp
	CustomDestinationForwardDestinationSplunk        *CustomDestinationForwardDestinationSplunk
	CustomDestinationForwardDestinationElasticsearch *CustomDestinationForwardDestinationElasticsearch

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CustomDestinationForwardDestinationHttpAsCustomDestinationForwardDestination is a convenience function that returns CustomDestinationForwardDestinationHttp wrapped in CustomDestinationForwardDestination.
func CustomDestinationForwardDestinationHttpAsCustomDestinationForwardDestination(v *CustomDestinationForwardDestinationHttp) CustomDestinationForwardDestination {
	return CustomDestinationForwardDestination{CustomDestinationForwardDestinationHttp: v}
}

// CustomDestinationForwardDestinationSplunkAsCustomDestinationForwardDestination is a convenience function that returns CustomDestinationForwardDestinationSplunk wrapped in CustomDestinationForwardDestination.
func CustomDestinationForwardDestinationSplunkAsCustomDestinationForwardDestination(v *CustomDestinationForwardDestinationSplunk) CustomDestinationForwardDestination {
	return CustomDestinationForwardDestination{CustomDestinationForwardDestinationSplunk: v}
}

// CustomDestinationForwardDestinationElasticsearchAsCustomDestinationForwardDestination is a convenience function that returns CustomDestinationForwardDestinationElasticsearch wrapped in CustomDestinationForwardDestination.
func CustomDestinationForwardDestinationElasticsearchAsCustomDestinationForwardDestination(v *CustomDestinationForwardDestinationElasticsearch) CustomDestinationForwardDestination {
	return CustomDestinationForwardDestination{CustomDestinationForwardDestinationElasticsearch: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CustomDestinationForwardDestination) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomDestinationForwardDestinationHttp
	err = datadog.Unmarshal(data, &obj.CustomDestinationForwardDestinationHttp)
	if err == nil {
		if obj.CustomDestinationForwardDestinationHttp != nil && obj.CustomDestinationForwardDestinationHttp.UnparsedObject == nil {
			jsonCustomDestinationForwardDestinationHttp, _ := datadog.Marshal(obj.CustomDestinationForwardDestinationHttp)
			if string(jsonCustomDestinationForwardDestinationHttp) == "{}" { // empty struct
				obj.CustomDestinationForwardDestinationHttp = nil
			} else {
				match++
			}
		} else {
			obj.CustomDestinationForwardDestinationHttp = nil
		}
	} else {
		obj.CustomDestinationForwardDestinationHttp = nil
	}

	// try to unmarshal data into CustomDestinationForwardDestinationSplunk
	err = datadog.Unmarshal(data, &obj.CustomDestinationForwardDestinationSplunk)
	if err == nil {
		if obj.CustomDestinationForwardDestinationSplunk != nil && obj.CustomDestinationForwardDestinationSplunk.UnparsedObject == nil {
			jsonCustomDestinationForwardDestinationSplunk, _ := datadog.Marshal(obj.CustomDestinationForwardDestinationSplunk)
			if string(jsonCustomDestinationForwardDestinationSplunk) == "{}" { // empty struct
				obj.CustomDestinationForwardDestinationSplunk = nil
			} else {
				match++
			}
		} else {
			obj.CustomDestinationForwardDestinationSplunk = nil
		}
	} else {
		obj.CustomDestinationForwardDestinationSplunk = nil
	}

	// try to unmarshal data into CustomDestinationForwardDestinationElasticsearch
	err = datadog.Unmarshal(data, &obj.CustomDestinationForwardDestinationElasticsearch)
	if err == nil {
		if obj.CustomDestinationForwardDestinationElasticsearch != nil && obj.CustomDestinationForwardDestinationElasticsearch.UnparsedObject == nil {
			jsonCustomDestinationForwardDestinationElasticsearch, _ := datadog.Marshal(obj.CustomDestinationForwardDestinationElasticsearch)
			if string(jsonCustomDestinationForwardDestinationElasticsearch) == "{}" { // empty struct
				obj.CustomDestinationForwardDestinationElasticsearch = nil
			} else {
				match++
			}
		} else {
			obj.CustomDestinationForwardDestinationElasticsearch = nil
		}
	} else {
		obj.CustomDestinationForwardDestinationElasticsearch = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CustomDestinationForwardDestinationHttp = nil
		obj.CustomDestinationForwardDestinationSplunk = nil
		obj.CustomDestinationForwardDestinationElasticsearch = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CustomDestinationForwardDestination) MarshalJSON() ([]byte, error) {
	if obj.CustomDestinationForwardDestinationHttp != nil {
		return datadog.Marshal(&obj.CustomDestinationForwardDestinationHttp)
	}

	if obj.CustomDestinationForwardDestinationSplunk != nil {
		return datadog.Marshal(&obj.CustomDestinationForwardDestinationSplunk)
	}

	if obj.CustomDestinationForwardDestinationElasticsearch != nil {
		return datadog.Marshal(&obj.CustomDestinationForwardDestinationElasticsearch)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CustomDestinationForwardDestination) GetActualInstance() interface{} {
	if obj.CustomDestinationForwardDestinationHttp != nil {
		return obj.CustomDestinationForwardDestinationHttp
	}

	if obj.CustomDestinationForwardDestinationSplunk != nil {
		return obj.CustomDestinationForwardDestinationSplunk
	}

	if obj.CustomDestinationForwardDestinationElasticsearch != nil {
		return obj.CustomDestinationForwardDestinationElasticsearch
	}

	// all schemas are nil
	return nil
}
