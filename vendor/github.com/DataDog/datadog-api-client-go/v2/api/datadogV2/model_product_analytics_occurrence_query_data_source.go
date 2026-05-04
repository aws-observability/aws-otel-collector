// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsOccurrenceQueryDataSource The data source identifier for occurrence queries.
type ProductAnalyticsOccurrenceQueryDataSource string

// List of ProductAnalyticsOccurrenceQueryDataSource.
const (
	PRODUCTANALYTICSOCCURRENCEQUERYDATASOURCE_PRODUCT_ANALYTICS_OCCURRENCE ProductAnalyticsOccurrenceQueryDataSource = "product_analytics_occurrence"
)

var allowedProductAnalyticsOccurrenceQueryDataSourceEnumValues = []ProductAnalyticsOccurrenceQueryDataSource{
	PRODUCTANALYTICSOCCURRENCEQUERYDATASOURCE_PRODUCT_ANALYTICS_OCCURRENCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsOccurrenceQueryDataSource) GetAllowedValues() []ProductAnalyticsOccurrenceQueryDataSource {
	return allowedProductAnalyticsOccurrenceQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsOccurrenceQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsOccurrenceQueryDataSource(value)
	return nil
}

// NewProductAnalyticsOccurrenceQueryDataSourceFromValue returns a pointer to a valid ProductAnalyticsOccurrenceQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsOccurrenceQueryDataSourceFromValue(v string) (*ProductAnalyticsOccurrenceQueryDataSource, error) {
	ev := ProductAnalyticsOccurrenceQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsOccurrenceQueryDataSource: valid values are %v", v, allowedProductAnalyticsOccurrenceQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsOccurrenceQueryDataSource) IsValid() bool {
	for _, existing := range allowedProductAnalyticsOccurrenceQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsOccurrenceQueryDataSource value.
func (v ProductAnalyticsOccurrenceQueryDataSource) Ptr() *ProductAnalyticsOccurrenceQueryDataSource {
	return &v
}
