// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsEventDataSource Data source for Product Analytics event queries.
type ProductAnalyticsEventDataSource string

// List of ProductAnalyticsEventDataSource.
const (
	PRODUCTANALYTICSEVENTDATASOURCE_PRODUCT_ANALYTICS ProductAnalyticsEventDataSource = "product_analytics"
)

var allowedProductAnalyticsEventDataSourceEnumValues = []ProductAnalyticsEventDataSource{
	PRODUCTANALYTICSEVENTDATASOURCE_PRODUCT_ANALYTICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsEventDataSource) GetAllowedValues() []ProductAnalyticsEventDataSource {
	return allowedProductAnalyticsEventDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsEventDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsEventDataSource(value)
	return nil
}

// NewProductAnalyticsEventDataSourceFromValue returns a pointer to a valid ProductAnalyticsEventDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsEventDataSourceFromValue(v string) (*ProductAnalyticsEventDataSource, error) {
	ev := ProductAnalyticsEventDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsEventDataSource: valid values are %v", v, allowedProductAnalyticsEventDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsEventDataSource) IsValid() bool {
	for _, existing := range allowedProductAnalyticsEventDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsEventDataSource value.
func (v ProductAnalyticsEventDataSource) Ptr() *ProductAnalyticsEventDataSource {
	return &v
}
