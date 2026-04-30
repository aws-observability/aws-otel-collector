// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsScalarResponseType The resource type identifier for a scalar analytics response.
type ProductAnalyticsScalarResponseType string

// List of ProductAnalyticsScalarResponseType.
const (
	PRODUCTANALYTICSSCALARRESPONSETYPE_SCALAR_RESPONSE ProductAnalyticsScalarResponseType = "scalar_response"
)

var allowedProductAnalyticsScalarResponseTypeEnumValues = []ProductAnalyticsScalarResponseType{
	PRODUCTANALYTICSSCALARRESPONSETYPE_SCALAR_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsScalarResponseType) GetAllowedValues() []ProductAnalyticsScalarResponseType {
	return allowedProductAnalyticsScalarResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsScalarResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsScalarResponseType(value)
	return nil
}

// NewProductAnalyticsScalarResponseTypeFromValue returns a pointer to a valid ProductAnalyticsScalarResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsScalarResponseTypeFromValue(v string) (*ProductAnalyticsScalarResponseType, error) {
	ev := ProductAnalyticsScalarResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsScalarResponseType: valid values are %v", v, allowedProductAnalyticsScalarResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsScalarResponseType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsScalarResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsScalarResponseType value.
func (v ProductAnalyticsScalarResponseType) Ptr() *ProductAnalyticsScalarResponseType {
	return &v
}
