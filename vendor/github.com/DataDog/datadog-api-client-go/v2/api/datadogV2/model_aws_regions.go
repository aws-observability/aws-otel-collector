// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSRegions - AWS Regions to collect data from. Defaults to `include_all`.
type AWSRegions struct {
	AWSRegionsIncludeAll  *AWSRegionsIncludeAll
	AWSRegionsIncludeOnly *AWSRegionsIncludeOnly

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AWSRegionsIncludeAllAsAWSRegions is a convenience function that returns AWSRegionsIncludeAll wrapped in AWSRegions.
func AWSRegionsIncludeAllAsAWSRegions(v *AWSRegionsIncludeAll) AWSRegions {
	return AWSRegions{AWSRegionsIncludeAll: v}
}

// AWSRegionsIncludeOnlyAsAWSRegions is a convenience function that returns AWSRegionsIncludeOnly wrapped in AWSRegions.
func AWSRegionsIncludeOnlyAsAWSRegions(v *AWSRegionsIncludeOnly) AWSRegions {
	return AWSRegions{AWSRegionsIncludeOnly: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AWSRegions) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSRegionsIncludeAll
	err = datadog.Unmarshal(data, &obj.AWSRegionsIncludeAll)
	if err == nil {
		if obj.AWSRegionsIncludeAll != nil && obj.AWSRegionsIncludeAll.UnparsedObject == nil {
			jsonAWSRegionsIncludeAll, _ := datadog.Marshal(obj.AWSRegionsIncludeAll)
			if string(jsonAWSRegionsIncludeAll) == "{}" { // empty struct
				obj.AWSRegionsIncludeAll = nil
			} else {
				match++
			}
		} else {
			obj.AWSRegionsIncludeAll = nil
		}
	} else {
		obj.AWSRegionsIncludeAll = nil
	}

	// try to unmarshal data into AWSRegionsIncludeOnly
	err = datadog.Unmarshal(data, &obj.AWSRegionsIncludeOnly)
	if err == nil {
		if obj.AWSRegionsIncludeOnly != nil && obj.AWSRegionsIncludeOnly.UnparsedObject == nil {
			jsonAWSRegionsIncludeOnly, _ := datadog.Marshal(obj.AWSRegionsIncludeOnly)
			if string(jsonAWSRegionsIncludeOnly) == "{}" { // empty struct
				obj.AWSRegionsIncludeOnly = nil
			} else {
				match++
			}
		} else {
			obj.AWSRegionsIncludeOnly = nil
		}
	} else {
		obj.AWSRegionsIncludeOnly = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AWSRegionsIncludeAll = nil
		obj.AWSRegionsIncludeOnly = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AWSRegions) MarshalJSON() ([]byte, error) {
	if obj.AWSRegionsIncludeAll != nil {
		return datadog.Marshal(&obj.AWSRegionsIncludeAll)
	}

	if obj.AWSRegionsIncludeOnly != nil {
		return datadog.Marshal(&obj.AWSRegionsIncludeOnly)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AWSRegions) GetActualInstance() interface{} {
	if obj.AWSRegionsIncludeAll != nil {
		return obj.AWSRegionsIncludeAll
	}

	if obj.AWSRegionsIncludeOnly != nil {
		return obj.AWSRegionsIncludeOnly
	}

	// all schemas are nil
	return nil
}
