// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsTimeseriesResponseType The resource type identifier for a timeseries analytics response.
type ProductAnalyticsTimeseriesResponseType string

// List of ProductAnalyticsTimeseriesResponseType.
const (
	PRODUCTANALYTICSTIMESERIESRESPONSETYPE_TIMESERIES_RESPONSE ProductAnalyticsTimeseriesResponseType = "timeseries_response"
)

var allowedProductAnalyticsTimeseriesResponseTypeEnumValues = []ProductAnalyticsTimeseriesResponseType{
	PRODUCTANALYTICSTIMESERIESRESPONSETYPE_TIMESERIES_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsTimeseriesResponseType) GetAllowedValues() []ProductAnalyticsTimeseriesResponseType {
	return allowedProductAnalyticsTimeseriesResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsTimeseriesResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsTimeseriesResponseType(value)
	return nil
}

// NewProductAnalyticsTimeseriesResponseTypeFromValue returns a pointer to a valid ProductAnalyticsTimeseriesResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsTimeseriesResponseTypeFromValue(v string) (*ProductAnalyticsTimeseriesResponseType, error) {
	ev := ProductAnalyticsTimeseriesResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsTimeseriesResponseType: valid values are %v", v, allowedProductAnalyticsTimeseriesResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsTimeseriesResponseType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsTimeseriesResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsTimeseriesResponseType value.
func (v ProductAnalyticsTimeseriesResponseType) Ptr() *ProductAnalyticsTimeseriesResponseType {
	return &v
}
