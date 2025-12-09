// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatsigCredentials - The definition of the `StatsigCredentials` object.
type StatsigCredentials struct {
	StatsigAPIKey *StatsigAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StatsigAPIKeyAsStatsigCredentials is a convenience function that returns StatsigAPIKey wrapped in StatsigCredentials.
func StatsigAPIKeyAsStatsigCredentials(v *StatsigAPIKey) StatsigCredentials {
	return StatsigCredentials{StatsigAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *StatsigCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into StatsigAPIKey
	err = datadog.Unmarshal(data, &obj.StatsigAPIKey)
	if err == nil {
		if obj.StatsigAPIKey != nil && obj.StatsigAPIKey.UnparsedObject == nil {
			jsonStatsigAPIKey, _ := datadog.Marshal(obj.StatsigAPIKey)
			if string(jsonStatsigAPIKey) == "{}" { // empty struct
				obj.StatsigAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.StatsigAPIKey = nil
		}
	} else {
		obj.StatsigAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.StatsigAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj StatsigCredentials) MarshalJSON() ([]byte, error) {
	if obj.StatsigAPIKey != nil {
		return datadog.Marshal(&obj.StatsigAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *StatsigCredentials) GetActualInstance() interface{} {
	if obj.StatsigAPIKey != nil {
		return obj.StatsigAPIKey
	}

	// all schemas are nil
	return nil
}
