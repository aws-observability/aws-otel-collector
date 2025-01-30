// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSNamespaceFilters - AWS Metrics namespace filters. Defaults to `exclude_only`.
type AWSNamespaceFilters struct {
	AWSNamespaceFiltersExcludeOnly *AWSNamespaceFiltersExcludeOnly
	AWSNamespaceFiltersIncludeOnly *AWSNamespaceFiltersIncludeOnly

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AWSNamespaceFiltersExcludeOnlyAsAWSNamespaceFilters is a convenience function that returns AWSNamespaceFiltersExcludeOnly wrapped in AWSNamespaceFilters.
func AWSNamespaceFiltersExcludeOnlyAsAWSNamespaceFilters(v *AWSNamespaceFiltersExcludeOnly) AWSNamespaceFilters {
	return AWSNamespaceFilters{AWSNamespaceFiltersExcludeOnly: v}
}

// AWSNamespaceFiltersIncludeOnlyAsAWSNamespaceFilters is a convenience function that returns AWSNamespaceFiltersIncludeOnly wrapped in AWSNamespaceFilters.
func AWSNamespaceFiltersIncludeOnlyAsAWSNamespaceFilters(v *AWSNamespaceFiltersIncludeOnly) AWSNamespaceFilters {
	return AWSNamespaceFilters{AWSNamespaceFiltersIncludeOnly: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AWSNamespaceFilters) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSNamespaceFiltersExcludeOnly
	err = datadog.Unmarshal(data, &obj.AWSNamespaceFiltersExcludeOnly)
	if err == nil {
		if obj.AWSNamespaceFiltersExcludeOnly != nil && obj.AWSNamespaceFiltersExcludeOnly.UnparsedObject == nil {
			jsonAWSNamespaceFiltersExcludeOnly, _ := datadog.Marshal(obj.AWSNamespaceFiltersExcludeOnly)
			if string(jsonAWSNamespaceFiltersExcludeOnly) == "{}" { // empty struct
				obj.AWSNamespaceFiltersExcludeOnly = nil
			} else {
				match++
			}
		} else {
			obj.AWSNamespaceFiltersExcludeOnly = nil
		}
	} else {
		obj.AWSNamespaceFiltersExcludeOnly = nil
	}

	// try to unmarshal data into AWSNamespaceFiltersIncludeOnly
	err = datadog.Unmarshal(data, &obj.AWSNamespaceFiltersIncludeOnly)
	if err == nil {
		if obj.AWSNamespaceFiltersIncludeOnly != nil && obj.AWSNamespaceFiltersIncludeOnly.UnparsedObject == nil {
			jsonAWSNamespaceFiltersIncludeOnly, _ := datadog.Marshal(obj.AWSNamespaceFiltersIncludeOnly)
			if string(jsonAWSNamespaceFiltersIncludeOnly) == "{}" { // empty struct
				obj.AWSNamespaceFiltersIncludeOnly = nil
			} else {
				match++
			}
		} else {
			obj.AWSNamespaceFiltersIncludeOnly = nil
		}
	} else {
		obj.AWSNamespaceFiltersIncludeOnly = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AWSNamespaceFiltersExcludeOnly = nil
		obj.AWSNamespaceFiltersIncludeOnly = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AWSNamespaceFilters) MarshalJSON() ([]byte, error) {
	if obj.AWSNamespaceFiltersExcludeOnly != nil {
		return datadog.Marshal(&obj.AWSNamespaceFiltersExcludeOnly)
	}

	if obj.AWSNamespaceFiltersIncludeOnly != nil {
		return datadog.Marshal(&obj.AWSNamespaceFiltersIncludeOnly)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AWSNamespaceFilters) GetActualInstance() interface{} {
	if obj.AWSNamespaceFiltersExcludeOnly != nil {
		return obj.AWSNamespaceFiltersExcludeOnly
	}

	if obj.AWSNamespaceFiltersIncludeOnly != nil {
		return obj.AWSNamespaceFiltersIncludeOnly
	}

	// all schemas are nil
	return nil
}
