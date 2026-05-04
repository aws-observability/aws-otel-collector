// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelDataSource Data source for user journey funnel queries.
type ProductAnalyticsFunnelDataSource string

// List of ProductAnalyticsFunnelDataSource.
const (
	PRODUCTANALYTICSFUNNELDATASOURCE_PRODUCT_ANALYTICS_JOURNEY ProductAnalyticsFunnelDataSource = "product_analytics_journey"
)

var allowedProductAnalyticsFunnelDataSourceEnumValues = []ProductAnalyticsFunnelDataSource{
	PRODUCTANALYTICSFUNNELDATASOURCE_PRODUCT_ANALYTICS_JOURNEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsFunnelDataSource) GetAllowedValues() []ProductAnalyticsFunnelDataSource {
	return allowedProductAnalyticsFunnelDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsFunnelDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsFunnelDataSource(value)
	return nil
}

// NewProductAnalyticsFunnelDataSourceFromValue returns a pointer to a valid ProductAnalyticsFunnelDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsFunnelDataSourceFromValue(v string) (*ProductAnalyticsFunnelDataSource, error) {
	ev := ProductAnalyticsFunnelDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsFunnelDataSource: valid values are %v", v, allowedProductAnalyticsFunnelDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsFunnelDataSource) IsValid() bool {
	for _, existing := range allowedProductAnalyticsFunnelDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsFunnelDataSource value.
func (v ProductAnalyticsFunnelDataSource) Ptr() *ProductAnalyticsFunnelDataSource {
	return &v
}
