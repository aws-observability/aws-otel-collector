// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsBaseQuery - A query definition discriminated by the `data_source` field.
// Use `product_analytics` for standard event queries, or
// `product_analytics_occurrence` for occurrence-filtered queries.
type ProductAnalyticsBaseQuery struct {
	ProductAnalyticsEventQuery      *ProductAnalyticsEventQuery
	ProductAnalyticsOccurrenceQuery *ProductAnalyticsOccurrenceQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ProductAnalyticsEventQueryAsProductAnalyticsBaseQuery is a convenience function that returns ProductAnalyticsEventQuery wrapped in ProductAnalyticsBaseQuery.
func ProductAnalyticsEventQueryAsProductAnalyticsBaseQuery(v *ProductAnalyticsEventQuery) ProductAnalyticsBaseQuery {
	return ProductAnalyticsBaseQuery{ProductAnalyticsEventQuery: v}
}

// ProductAnalyticsOccurrenceQueryAsProductAnalyticsBaseQuery is a convenience function that returns ProductAnalyticsOccurrenceQuery wrapped in ProductAnalyticsBaseQuery.
func ProductAnalyticsOccurrenceQueryAsProductAnalyticsBaseQuery(v *ProductAnalyticsOccurrenceQuery) ProductAnalyticsBaseQuery {
	return ProductAnalyticsBaseQuery{ProductAnalyticsOccurrenceQuery: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ProductAnalyticsBaseQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ProductAnalyticsEventQuery
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsEventQuery)
	if err == nil {
		if obj.ProductAnalyticsEventQuery != nil && obj.ProductAnalyticsEventQuery.UnparsedObject == nil {
			jsonProductAnalyticsEventQuery, _ := datadog.Marshal(obj.ProductAnalyticsEventQuery)
			if string(jsonProductAnalyticsEventQuery) == "{}" { // empty struct
				obj.ProductAnalyticsEventQuery = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsEventQuery = nil
		}
	} else {
		obj.ProductAnalyticsEventQuery = nil
	}

	// try to unmarshal data into ProductAnalyticsOccurrenceQuery
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsOccurrenceQuery)
	if err == nil {
		if obj.ProductAnalyticsOccurrenceQuery != nil && obj.ProductAnalyticsOccurrenceQuery.UnparsedObject == nil {
			jsonProductAnalyticsOccurrenceQuery, _ := datadog.Marshal(obj.ProductAnalyticsOccurrenceQuery)
			if string(jsonProductAnalyticsOccurrenceQuery) == "{}" { // empty struct
				obj.ProductAnalyticsOccurrenceQuery = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsOccurrenceQuery = nil
		}
	} else {
		obj.ProductAnalyticsOccurrenceQuery = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ProductAnalyticsEventQuery = nil
		obj.ProductAnalyticsOccurrenceQuery = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ProductAnalyticsBaseQuery) MarshalJSON() ([]byte, error) {
	if obj.ProductAnalyticsEventQuery != nil {
		return datadog.Marshal(&obj.ProductAnalyticsEventQuery)
	}

	if obj.ProductAnalyticsOccurrenceQuery != nil {
		return datadog.Marshal(&obj.ProductAnalyticsOccurrenceQuery)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ProductAnalyticsBaseQuery) GetActualInstance() interface{} {
	if obj.ProductAnalyticsEventQuery != nil {
		return obj.ProductAnalyticsEventQuery
	}

	if obj.ProductAnalyticsOccurrenceQuery != nil {
		return obj.ProductAnalyticsOccurrenceQuery
	}

	// all schemas are nil
	return nil
}
