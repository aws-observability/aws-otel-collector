// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsEventQueryDataSource The data source identifier.
type ProductAnalyticsEventQueryDataSource string

// List of ProductAnalyticsEventQueryDataSource.
const (
	PRODUCTANALYTICSEVENTQUERYDATASOURCE_PRODUCT_ANALYTICS ProductAnalyticsEventQueryDataSource = "product_analytics"
)

var allowedProductAnalyticsEventQueryDataSourceEnumValues = []ProductAnalyticsEventQueryDataSource{
	PRODUCTANALYTICSEVENTQUERYDATASOURCE_PRODUCT_ANALYTICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsEventQueryDataSource) GetAllowedValues() []ProductAnalyticsEventQueryDataSource {
	return allowedProductAnalyticsEventQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsEventQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsEventQueryDataSource(value)
	return nil
}

// NewProductAnalyticsEventQueryDataSourceFromValue returns a pointer to a valid ProductAnalyticsEventQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsEventQueryDataSourceFromValue(v string) (*ProductAnalyticsEventQueryDataSource, error) {
	ev := ProductAnalyticsEventQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsEventQueryDataSource: valid values are %v", v, allowedProductAnalyticsEventQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsEventQueryDataSource) IsValid() bool {
	for _, existing := range allowedProductAnalyticsEventQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsEventQueryDataSource value.
func (v ProductAnalyticsEventQueryDataSource) Ptr() *ProductAnalyticsEventQueryDataSource {
	return &v
}
