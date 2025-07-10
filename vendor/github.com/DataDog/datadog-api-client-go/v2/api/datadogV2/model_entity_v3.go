// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3 - Entity schema v3.
type EntityV3 struct {
	EntityV3Service   *EntityV3Service
	EntityV3Datastore *EntityV3Datastore
	EntityV3Queue     *EntityV3Queue
	EntityV3System    *EntityV3System
	EntityV3API       *EntityV3API

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// EntityV3ServiceAsEntityV3 is a convenience function that returns EntityV3Service wrapped in EntityV3.
func EntityV3ServiceAsEntityV3(v *EntityV3Service) EntityV3 {
	return EntityV3{EntityV3Service: v}
}

// EntityV3DatastoreAsEntityV3 is a convenience function that returns EntityV3Datastore wrapped in EntityV3.
func EntityV3DatastoreAsEntityV3(v *EntityV3Datastore) EntityV3 {
	return EntityV3{EntityV3Datastore: v}
}

// EntityV3QueueAsEntityV3 is a convenience function that returns EntityV3Queue wrapped in EntityV3.
func EntityV3QueueAsEntityV3(v *EntityV3Queue) EntityV3 {
	return EntityV3{EntityV3Queue: v}
}

// EntityV3SystemAsEntityV3 is a convenience function that returns EntityV3System wrapped in EntityV3.
func EntityV3SystemAsEntityV3(v *EntityV3System) EntityV3 {
	return EntityV3{EntityV3System: v}
}

// EntityV3APIAsEntityV3 is a convenience function that returns EntityV3API wrapped in EntityV3.
func EntityV3APIAsEntityV3(v *EntityV3API) EntityV3 {
	return EntityV3{EntityV3API: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *EntityV3) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EntityV3Service
	err = datadog.Unmarshal(data, &obj.EntityV3Service)
	if err == nil {
		if obj.EntityV3Service != nil && obj.EntityV3Service.UnparsedObject == nil {
			jsonEntityV3Service, _ := datadog.Marshal(obj.EntityV3Service)
			if string(jsonEntityV3Service) == "{}" { // empty struct
				obj.EntityV3Service = nil
			} else {
				match++
			}
		} else {
			obj.EntityV3Service = nil
		}
	} else {
		obj.EntityV3Service = nil
	}

	// try to unmarshal data into EntityV3Datastore
	err = datadog.Unmarshal(data, &obj.EntityV3Datastore)
	if err == nil {
		if obj.EntityV3Datastore != nil && obj.EntityV3Datastore.UnparsedObject == nil {
			jsonEntityV3Datastore, _ := datadog.Marshal(obj.EntityV3Datastore)
			if string(jsonEntityV3Datastore) == "{}" { // empty struct
				obj.EntityV3Datastore = nil
			} else {
				match++
			}
		} else {
			obj.EntityV3Datastore = nil
		}
	} else {
		obj.EntityV3Datastore = nil
	}

	// try to unmarshal data into EntityV3Queue
	err = datadog.Unmarshal(data, &obj.EntityV3Queue)
	if err == nil {
		if obj.EntityV3Queue != nil && obj.EntityV3Queue.UnparsedObject == nil {
			jsonEntityV3Queue, _ := datadog.Marshal(obj.EntityV3Queue)
			if string(jsonEntityV3Queue) == "{}" { // empty struct
				obj.EntityV3Queue = nil
			} else {
				match++
			}
		} else {
			obj.EntityV3Queue = nil
		}
	} else {
		obj.EntityV3Queue = nil
	}

	// try to unmarshal data into EntityV3System
	err = datadog.Unmarshal(data, &obj.EntityV3System)
	if err == nil {
		if obj.EntityV3System != nil && obj.EntityV3System.UnparsedObject == nil {
			jsonEntityV3System, _ := datadog.Marshal(obj.EntityV3System)
			if string(jsonEntityV3System) == "{}" { // empty struct
				obj.EntityV3System = nil
			} else {
				match++
			}
		} else {
			obj.EntityV3System = nil
		}
	} else {
		obj.EntityV3System = nil
	}

	// try to unmarshal data into EntityV3API
	err = datadog.Unmarshal(data, &obj.EntityV3API)
	if err == nil {
		if obj.EntityV3API != nil && obj.EntityV3API.UnparsedObject == nil {
			jsonEntityV3API, _ := datadog.Marshal(obj.EntityV3API)
			if string(jsonEntityV3API) == "{}" { // empty struct
				obj.EntityV3API = nil
			} else {
				match++
			}
		} else {
			obj.EntityV3API = nil
		}
	} else {
		obj.EntityV3API = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.EntityV3Service = nil
		obj.EntityV3Datastore = nil
		obj.EntityV3Queue = nil
		obj.EntityV3System = nil
		obj.EntityV3API = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj EntityV3) MarshalJSON() ([]byte, error) {
	if obj.EntityV3Service != nil {
		return datadog.Marshal(&obj.EntityV3Service)
	}

	if obj.EntityV3Datastore != nil {
		return datadog.Marshal(&obj.EntityV3Datastore)
	}

	if obj.EntityV3Queue != nil {
		return datadog.Marshal(&obj.EntityV3Queue)
	}

	if obj.EntityV3System != nil {
		return datadog.Marshal(&obj.EntityV3System)
	}

	if obj.EntityV3API != nil {
		return datadog.Marshal(&obj.EntityV3API)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *EntityV3) GetActualInstance() interface{} {
	if obj.EntityV3Service != nil {
		return obj.EntityV3Service
	}

	if obj.EntityV3Datastore != nil {
		return obj.EntityV3Datastore
	}

	if obj.EntityV3Queue != nil {
		return obj.EntityV3Queue
	}

	if obj.EntityV3System != nil {
		return obj.EntityV3System
	}

	if obj.EntityV3API != nil {
		return obj.EntityV3API
	}

	// all schemas are nil
	return nil
}
