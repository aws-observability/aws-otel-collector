// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsScalarColumnType Column type.
type ProductAnalyticsScalarColumnType string

// List of ProductAnalyticsScalarColumnType.
const (
	PRODUCTANALYTICSSCALARCOLUMNTYPE_NUMBER ProductAnalyticsScalarColumnType = "number"
	PRODUCTANALYTICSSCALARCOLUMNTYPE_GROUP  ProductAnalyticsScalarColumnType = "group"
)

var allowedProductAnalyticsScalarColumnTypeEnumValues = []ProductAnalyticsScalarColumnType{
	PRODUCTANALYTICSSCALARCOLUMNTYPE_NUMBER,
	PRODUCTANALYTICSSCALARCOLUMNTYPE_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsScalarColumnType) GetAllowedValues() []ProductAnalyticsScalarColumnType {
	return allowedProductAnalyticsScalarColumnTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsScalarColumnType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsScalarColumnType(value)
	return nil
}

// NewProductAnalyticsScalarColumnTypeFromValue returns a pointer to a valid ProductAnalyticsScalarColumnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsScalarColumnTypeFromValue(v string) (*ProductAnalyticsScalarColumnType, error) {
	ev := ProductAnalyticsScalarColumnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsScalarColumnType: valid values are %v", v, allowedProductAnalyticsScalarColumnTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsScalarColumnType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsScalarColumnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsScalarColumnType value.
func (v ProductAnalyticsScalarColumnType) Ptr() *ProductAnalyticsScalarColumnType {
	return &v
}
