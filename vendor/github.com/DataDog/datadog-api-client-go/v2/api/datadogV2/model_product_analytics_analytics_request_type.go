// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsAnalyticsRequestType The resource type for analytics requests.
type ProductAnalyticsAnalyticsRequestType string

// List of ProductAnalyticsAnalyticsRequestType.
const (
	PRODUCTANALYTICSANALYTICSREQUESTTYPE_FORMULA_ANALYTICS_EXTENDED_REQUEST ProductAnalyticsAnalyticsRequestType = "formula_analytics_extended_request"
)

var allowedProductAnalyticsAnalyticsRequestTypeEnumValues = []ProductAnalyticsAnalyticsRequestType{
	PRODUCTANALYTICSANALYTICSREQUESTTYPE_FORMULA_ANALYTICS_EXTENDED_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsAnalyticsRequestType) GetAllowedValues() []ProductAnalyticsAnalyticsRequestType {
	return allowedProductAnalyticsAnalyticsRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsAnalyticsRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsAnalyticsRequestType(value)
	return nil
}

// NewProductAnalyticsAnalyticsRequestTypeFromValue returns a pointer to a valid ProductAnalyticsAnalyticsRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsAnalyticsRequestTypeFromValue(v string) (*ProductAnalyticsAnalyticsRequestType, error) {
	ev := ProductAnalyticsAnalyticsRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsAnalyticsRequestType: valid values are %v", v, allowedProductAnalyticsAnalyticsRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsAnalyticsRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsAnalyticsRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsAnalyticsRequestType value.
func (v ProductAnalyticsAnalyticsRequestType) Ptr() *ProductAnalyticsAnalyticsRequestType {
	return &v
}
