// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsExecutionType Override the query execution strategy.
type ProductAnalyticsExecutionType string

// List of ProductAnalyticsExecutionType.
const (
	PRODUCTANALYTICSEXECUTIONTYPE_SIMPLE            ProductAnalyticsExecutionType = "simple"
	PRODUCTANALYTICSEXECUTIONTYPE_BACKGROUND        ProductAnalyticsExecutionType = "background"
	PRODUCTANALYTICSEXECUTIONTYPE_TRINO_MULTISTEP   ProductAnalyticsExecutionType = "trino-multistep"
	PRODUCTANALYTICSEXECUTIONTYPE_MATERIALIZED_VIEW ProductAnalyticsExecutionType = "materialized-view"
)

var allowedProductAnalyticsExecutionTypeEnumValues = []ProductAnalyticsExecutionType{
	PRODUCTANALYTICSEXECUTIONTYPE_SIMPLE,
	PRODUCTANALYTICSEXECUTIONTYPE_BACKGROUND,
	PRODUCTANALYTICSEXECUTIONTYPE_TRINO_MULTISTEP,
	PRODUCTANALYTICSEXECUTIONTYPE_MATERIALIZED_VIEW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsExecutionType) GetAllowedValues() []ProductAnalyticsExecutionType {
	return allowedProductAnalyticsExecutionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsExecutionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsExecutionType(value)
	return nil
}

// NewProductAnalyticsExecutionTypeFromValue returns a pointer to a valid ProductAnalyticsExecutionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsExecutionTypeFromValue(v string) (*ProductAnalyticsExecutionType, error) {
	ev := ProductAnalyticsExecutionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsExecutionType: valid values are %v", v, allowedProductAnalyticsExecutionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsExecutionType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsExecutionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsExecutionType value.
func (v ProductAnalyticsExecutionType) Ptr() *ProductAnalyticsExecutionType {
	return &v
}
