// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseForwardDestination - A custom destination's location to forward logs.
type CustomDestinationResponseForwardDestination struct {
	CustomDestinationResponseForwardDestinationHttp              *CustomDestinationResponseForwardDestinationHttp
	CustomDestinationResponseForwardDestinationSplunk            *CustomDestinationResponseForwardDestinationSplunk
	CustomDestinationResponseForwardDestinationElasticsearch     *CustomDestinationResponseForwardDestinationElasticsearch
	CustomDestinationResponseForwardDestinationMicrosoftSentinel *CustomDestinationResponseForwardDestinationMicrosoftSentinel

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CustomDestinationResponseForwardDestinationHttpAsCustomDestinationResponseForwardDestination is a convenience function that returns CustomDestinationResponseForwardDestinationHttp wrapped in CustomDestinationResponseForwardDestination.
func CustomDestinationResponseForwardDestinationHttpAsCustomDestinationResponseForwardDestination(v *CustomDestinationResponseForwardDestinationHttp) CustomDestinationResponseForwardDestination {
	return CustomDestinationResponseForwardDestination{CustomDestinationResponseForwardDestinationHttp: v}
}

// CustomDestinationResponseForwardDestinationSplunkAsCustomDestinationResponseForwardDestination is a convenience function that returns CustomDestinationResponseForwardDestinationSplunk wrapped in CustomDestinationResponseForwardDestination.
func CustomDestinationResponseForwardDestinationSplunkAsCustomDestinationResponseForwardDestination(v *CustomDestinationResponseForwardDestinationSplunk) CustomDestinationResponseForwardDestination {
	return CustomDestinationResponseForwardDestination{CustomDestinationResponseForwardDestinationSplunk: v}
}

// CustomDestinationResponseForwardDestinationElasticsearchAsCustomDestinationResponseForwardDestination is a convenience function that returns CustomDestinationResponseForwardDestinationElasticsearch wrapped in CustomDestinationResponseForwardDestination.
func CustomDestinationResponseForwardDestinationElasticsearchAsCustomDestinationResponseForwardDestination(v *CustomDestinationResponseForwardDestinationElasticsearch) CustomDestinationResponseForwardDestination {
	return CustomDestinationResponseForwardDestination{CustomDestinationResponseForwardDestinationElasticsearch: v}
}

// CustomDestinationResponseForwardDestinationMicrosoftSentinelAsCustomDestinationResponseForwardDestination is a convenience function that returns CustomDestinationResponseForwardDestinationMicrosoftSentinel wrapped in CustomDestinationResponseForwardDestination.
func CustomDestinationResponseForwardDestinationMicrosoftSentinelAsCustomDestinationResponseForwardDestination(v *CustomDestinationResponseForwardDestinationMicrosoftSentinel) CustomDestinationResponseForwardDestination {
	return CustomDestinationResponseForwardDestination{CustomDestinationResponseForwardDestinationMicrosoftSentinel: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CustomDestinationResponseForwardDestination) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomDestinationResponseForwardDestinationHttp
	err = datadog.Unmarshal(data, &obj.CustomDestinationResponseForwardDestinationHttp)
	if err == nil {
		if obj.CustomDestinationResponseForwardDestinationHttp != nil && obj.CustomDestinationResponseForwardDestinationHttp.UnparsedObject == nil {
			jsonCustomDestinationResponseForwardDestinationHttp, _ := datadog.Marshal(obj.CustomDestinationResponseForwardDestinationHttp)
			if string(jsonCustomDestinationResponseForwardDestinationHttp) == "{}" { // empty struct
				obj.CustomDestinationResponseForwardDestinationHttp = nil
			} else {
				match++
			}
		} else {
			obj.CustomDestinationResponseForwardDestinationHttp = nil
		}
	} else {
		obj.CustomDestinationResponseForwardDestinationHttp = nil
	}

	// try to unmarshal data into CustomDestinationResponseForwardDestinationSplunk
	err = datadog.Unmarshal(data, &obj.CustomDestinationResponseForwardDestinationSplunk)
	if err == nil {
		if obj.CustomDestinationResponseForwardDestinationSplunk != nil && obj.CustomDestinationResponseForwardDestinationSplunk.UnparsedObject == nil {
			jsonCustomDestinationResponseForwardDestinationSplunk, _ := datadog.Marshal(obj.CustomDestinationResponseForwardDestinationSplunk)
			if string(jsonCustomDestinationResponseForwardDestinationSplunk) == "{}" { // empty struct
				obj.CustomDestinationResponseForwardDestinationSplunk = nil
			} else {
				match++
			}
		} else {
			obj.CustomDestinationResponseForwardDestinationSplunk = nil
		}
	} else {
		obj.CustomDestinationResponseForwardDestinationSplunk = nil
	}

	// try to unmarshal data into CustomDestinationResponseForwardDestinationElasticsearch
	err = datadog.Unmarshal(data, &obj.CustomDestinationResponseForwardDestinationElasticsearch)
	if err == nil {
		if obj.CustomDestinationResponseForwardDestinationElasticsearch != nil && obj.CustomDestinationResponseForwardDestinationElasticsearch.UnparsedObject == nil {
			jsonCustomDestinationResponseForwardDestinationElasticsearch, _ := datadog.Marshal(obj.CustomDestinationResponseForwardDestinationElasticsearch)
			if string(jsonCustomDestinationResponseForwardDestinationElasticsearch) == "{}" { // empty struct
				obj.CustomDestinationResponseForwardDestinationElasticsearch = nil
			} else {
				match++
			}
		} else {
			obj.CustomDestinationResponseForwardDestinationElasticsearch = nil
		}
	} else {
		obj.CustomDestinationResponseForwardDestinationElasticsearch = nil
	}

	// try to unmarshal data into CustomDestinationResponseForwardDestinationMicrosoftSentinel
	err = datadog.Unmarshal(data, &obj.CustomDestinationResponseForwardDestinationMicrosoftSentinel)
	if err == nil {
		if obj.CustomDestinationResponseForwardDestinationMicrosoftSentinel != nil && obj.CustomDestinationResponseForwardDestinationMicrosoftSentinel.UnparsedObject == nil {
			jsonCustomDestinationResponseForwardDestinationMicrosoftSentinel, _ := datadog.Marshal(obj.CustomDestinationResponseForwardDestinationMicrosoftSentinel)
			if string(jsonCustomDestinationResponseForwardDestinationMicrosoftSentinel) == "{}" { // empty struct
				obj.CustomDestinationResponseForwardDestinationMicrosoftSentinel = nil
			} else {
				match++
			}
		} else {
			obj.CustomDestinationResponseForwardDestinationMicrosoftSentinel = nil
		}
	} else {
		obj.CustomDestinationResponseForwardDestinationMicrosoftSentinel = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CustomDestinationResponseForwardDestinationHttp = nil
		obj.CustomDestinationResponseForwardDestinationSplunk = nil
		obj.CustomDestinationResponseForwardDestinationElasticsearch = nil
		obj.CustomDestinationResponseForwardDestinationMicrosoftSentinel = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CustomDestinationResponseForwardDestination) MarshalJSON() ([]byte, error) {
	if obj.CustomDestinationResponseForwardDestinationHttp != nil {
		return datadog.Marshal(&obj.CustomDestinationResponseForwardDestinationHttp)
	}

	if obj.CustomDestinationResponseForwardDestinationSplunk != nil {
		return datadog.Marshal(&obj.CustomDestinationResponseForwardDestinationSplunk)
	}

	if obj.CustomDestinationResponseForwardDestinationElasticsearch != nil {
		return datadog.Marshal(&obj.CustomDestinationResponseForwardDestinationElasticsearch)
	}

	if obj.CustomDestinationResponseForwardDestinationMicrosoftSentinel != nil {
		return datadog.Marshal(&obj.CustomDestinationResponseForwardDestinationMicrosoftSentinel)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CustomDestinationResponseForwardDestination) GetActualInstance() interface{} {
	if obj.CustomDestinationResponseForwardDestinationHttp != nil {
		return obj.CustomDestinationResponseForwardDestinationHttp
	}

	if obj.CustomDestinationResponseForwardDestinationSplunk != nil {
		return obj.CustomDestinationResponseForwardDestinationSplunk
	}

	if obj.CustomDestinationResponseForwardDestinationElasticsearch != nil {
		return obj.CustomDestinationResponseForwardDestinationElasticsearch
	}

	if obj.CustomDestinationResponseForwardDestinationMicrosoftSentinel != nil {
		return obj.CustomDestinationResponseForwardDestinationMicrosoftSentinel
	}

	// all schemas are nil
	return nil
}
