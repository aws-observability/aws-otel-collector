// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus Denotes whether mapping keys were available for this endpoint.
type BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus string

// List of BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus.
const (
	BILLINGDIMENSIONSMAPPINGBODYITEMATTRIBUTESENDPOINTSITEMSSTATUS_OK        BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus = "OK"
	BILLINGDIMENSIONSMAPPINGBODYITEMATTRIBUTESENDPOINTSITEMSSTATUS_NOT_FOUND BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus = "NOT_FOUND"
)

var allowedBillingDimensionsMappingBodyItemAttributesEndpointsItemsStatusEnumValues = []BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus{
	BILLINGDIMENSIONSMAPPINGBODYITEMATTRIBUTESENDPOINTSITEMSSTATUS_OK,
	BILLINGDIMENSIONSMAPPINGBODYITEMATTRIBUTESENDPOINTSITEMSSTATUS_NOT_FOUND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus) GetAllowedValues() []BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus {
	return allowedBillingDimensionsMappingBodyItemAttributesEndpointsItemsStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus(value)
	return nil
}

// NewBillingDimensionsMappingBodyItemAttributesEndpointsItemsStatusFromValue returns a pointer to a valid BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBillingDimensionsMappingBodyItemAttributesEndpointsItemsStatusFromValue(v string) (*BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus, error) {
	ev := BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus: valid values are %v", v, allowedBillingDimensionsMappingBodyItemAttributesEndpointsItemsStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus) IsValid() bool {
	for _, existing := range allowedBillingDimensionsMappingBodyItemAttributesEndpointsItemsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus value.
func (v BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus) Ptr() *BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus {
	return &v
}
