// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelRequestType Request type for user journey funnel widget.
type ProductAnalyticsFunnelRequestType string

// List of ProductAnalyticsFunnelRequestType.
const (
	PRODUCTANALYTICSFUNNELREQUESTTYPE_USER_JOURNEY_FUNNEL ProductAnalyticsFunnelRequestType = "user_journey_funnel"
)

var allowedProductAnalyticsFunnelRequestTypeEnumValues = []ProductAnalyticsFunnelRequestType{
	PRODUCTANALYTICSFUNNELREQUESTTYPE_USER_JOURNEY_FUNNEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsFunnelRequestType) GetAllowedValues() []ProductAnalyticsFunnelRequestType {
	return allowedProductAnalyticsFunnelRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsFunnelRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsFunnelRequestType(value)
	return nil
}

// NewProductAnalyticsFunnelRequestTypeFromValue returns a pointer to a valid ProductAnalyticsFunnelRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsFunnelRequestTypeFromValue(v string) (*ProductAnalyticsFunnelRequestType, error) {
	ev := ProductAnalyticsFunnelRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsFunnelRequestType: valid values are %v", v, allowedProductAnalyticsFunnelRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsFunnelRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsFunnelRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsFunnelRequestType value.
func (v ProductAnalyticsFunnelRequestType) Ptr() *ProductAnalyticsFunnelRequestType {
	return &v
}
