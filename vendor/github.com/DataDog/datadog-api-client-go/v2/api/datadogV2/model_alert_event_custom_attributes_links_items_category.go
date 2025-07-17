// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertEventCustomAttributesLinksItemsCategory The category of the link.
type AlertEventCustomAttributesLinksItemsCategory string

// List of AlertEventCustomAttributesLinksItemsCategory.
const (
	ALERTEVENTCUSTOMATTRIBUTESLINKSITEMSCATEGORY_RUNBOOK       AlertEventCustomAttributesLinksItemsCategory = "runbook"
	ALERTEVENTCUSTOMATTRIBUTESLINKSITEMSCATEGORY_DOCUMENTATION AlertEventCustomAttributesLinksItemsCategory = "documentation"
	ALERTEVENTCUSTOMATTRIBUTESLINKSITEMSCATEGORY_DASHBOARD     AlertEventCustomAttributesLinksItemsCategory = "dashboard"
)

var allowedAlertEventCustomAttributesLinksItemsCategoryEnumValues = []AlertEventCustomAttributesLinksItemsCategory{
	ALERTEVENTCUSTOMATTRIBUTESLINKSITEMSCATEGORY_RUNBOOK,
	ALERTEVENTCUSTOMATTRIBUTESLINKSITEMSCATEGORY_DOCUMENTATION,
	ALERTEVENTCUSTOMATTRIBUTESLINKSITEMSCATEGORY_DASHBOARD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AlertEventCustomAttributesLinksItemsCategory) GetAllowedValues() []AlertEventCustomAttributesLinksItemsCategory {
	return allowedAlertEventCustomAttributesLinksItemsCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertEventCustomAttributesLinksItemsCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertEventCustomAttributesLinksItemsCategory(value)
	return nil
}

// NewAlertEventCustomAttributesLinksItemsCategoryFromValue returns a pointer to a valid AlertEventCustomAttributesLinksItemsCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertEventCustomAttributesLinksItemsCategoryFromValue(v string) (*AlertEventCustomAttributesLinksItemsCategory, error) {
	ev := AlertEventCustomAttributesLinksItemsCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertEventCustomAttributesLinksItemsCategory: valid values are %v", v, allowedAlertEventCustomAttributesLinksItemsCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertEventCustomAttributesLinksItemsCategory) IsValid() bool {
	for _, existing := range allowedAlertEventCustomAttributesLinksItemsCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertEventCustomAttributesLinksItemsCategory value.
func (v AlertEventCustomAttributesLinksItemsCategory) Ptr() *AlertEventCustomAttributesLinksItemsCategory {
	return &v
}
