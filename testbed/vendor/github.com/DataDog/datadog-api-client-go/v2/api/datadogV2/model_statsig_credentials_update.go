// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatsigCredentialsUpdate - The definition of the `StatsigCredentialsUpdate` object.
type StatsigCredentialsUpdate struct {
	StatsigAPIKeyUpdate *StatsigAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StatsigAPIKeyUpdateAsStatsigCredentialsUpdate is a convenience function that returns StatsigAPIKeyUpdate wrapped in StatsigCredentialsUpdate.
func StatsigAPIKeyUpdateAsStatsigCredentialsUpdate(v *StatsigAPIKeyUpdate) StatsigCredentialsUpdate {
	return StatsigCredentialsUpdate{StatsigAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *StatsigCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into StatsigAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.StatsigAPIKeyUpdate)
	if err == nil {
		if obj.StatsigAPIKeyUpdate != nil && obj.StatsigAPIKeyUpdate.UnparsedObject == nil {
			jsonStatsigAPIKeyUpdate, _ := datadog.Marshal(obj.StatsigAPIKeyUpdate)
			if string(jsonStatsigAPIKeyUpdate) == "{}" { // empty struct
				obj.StatsigAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.StatsigAPIKeyUpdate = nil
		}
	} else {
		obj.StatsigAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.StatsigAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj StatsigCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.StatsigAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.StatsigAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *StatsigCredentialsUpdate) GetActualInstance() interface{} {
	if obj.StatsigAPIKeyUpdate != nil {
		return obj.StatsigAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
