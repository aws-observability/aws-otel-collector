// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsResponseMetaStatus The execution status of a Product Analytics query.
type ProductAnalyticsResponseMetaStatus string

// List of ProductAnalyticsResponseMetaStatus.
const (
	PRODUCTANALYTICSRESPONSEMETASTATUS_DONE    ProductAnalyticsResponseMetaStatus = "done"
	PRODUCTANALYTICSRESPONSEMETASTATUS_RUNNING ProductAnalyticsResponseMetaStatus = "running"
	PRODUCTANALYTICSRESPONSEMETASTATUS_TIMEOUT ProductAnalyticsResponseMetaStatus = "timeout"
)

var allowedProductAnalyticsResponseMetaStatusEnumValues = []ProductAnalyticsResponseMetaStatus{
	PRODUCTANALYTICSRESPONSEMETASTATUS_DONE,
	PRODUCTANALYTICSRESPONSEMETASTATUS_RUNNING,
	PRODUCTANALYTICSRESPONSEMETASTATUS_TIMEOUT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsResponseMetaStatus) GetAllowedValues() []ProductAnalyticsResponseMetaStatus {
	return allowedProductAnalyticsResponseMetaStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsResponseMetaStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsResponseMetaStatus(value)
	return nil
}

// NewProductAnalyticsResponseMetaStatusFromValue returns a pointer to a valid ProductAnalyticsResponseMetaStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsResponseMetaStatusFromValue(v string) (*ProductAnalyticsResponseMetaStatus, error) {
	ev := ProductAnalyticsResponseMetaStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsResponseMetaStatus: valid values are %v", v, allowedProductAnalyticsResponseMetaStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsResponseMetaStatus) IsValid() bool {
	for _, existing := range allowedProductAnalyticsResponseMetaStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsResponseMetaStatus value.
func (v ProductAnalyticsResponseMetaStatus) Ptr() *ProductAnalyticsResponseMetaStatus {
	return &v
}
