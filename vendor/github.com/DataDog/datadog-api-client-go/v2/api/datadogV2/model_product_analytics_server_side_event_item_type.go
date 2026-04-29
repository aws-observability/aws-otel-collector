// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsServerSideEventItemType The type of Product Analytics event. Must be `server` for server-side events.
type ProductAnalyticsServerSideEventItemType string

// List of ProductAnalyticsServerSideEventItemType.
const (
	PRODUCTANALYTICSSERVERSIDEEVENTITEMTYPE_SERVER ProductAnalyticsServerSideEventItemType = "server"
)

var allowedProductAnalyticsServerSideEventItemTypeEnumValues = []ProductAnalyticsServerSideEventItemType{
	PRODUCTANALYTICSSERVERSIDEEVENTITEMTYPE_SERVER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsServerSideEventItemType) GetAllowedValues() []ProductAnalyticsServerSideEventItemType {
	return allowedProductAnalyticsServerSideEventItemTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsServerSideEventItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsServerSideEventItemType(value)
	return nil
}

// NewProductAnalyticsServerSideEventItemTypeFromValue returns a pointer to a valid ProductAnalyticsServerSideEventItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsServerSideEventItemTypeFromValue(v string) (*ProductAnalyticsServerSideEventItemType, error) {
	ev := ProductAnalyticsServerSideEventItemType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsServerSideEventItemType: valid values are %v", v, allowedProductAnalyticsServerSideEventItemTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsServerSideEventItemType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsServerSideEventItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsServerSideEventItemType value.
func (v ProductAnalyticsServerSideEventItemType) Ptr() *ProductAnalyticsServerSideEventItemType {
	return &v
}
