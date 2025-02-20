// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageMetricCategory Contains the metric category.
type UsageMetricCategory string

// List of UsageMetricCategory.
const (
	USAGEMETRICCATEGORY_STANDARD UsageMetricCategory = "standard"
	USAGEMETRICCATEGORY_CUSTOM   UsageMetricCategory = "custom"
)

var allowedUsageMetricCategoryEnumValues = []UsageMetricCategory{
	USAGEMETRICCATEGORY_STANDARD,
	USAGEMETRICCATEGORY_CUSTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UsageMetricCategory) GetAllowedValues() []UsageMetricCategory {
	return allowedUsageMetricCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UsageMetricCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsageMetricCategory(value)
	return nil
}

// NewUsageMetricCategoryFromValue returns a pointer to a valid UsageMetricCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUsageMetricCategoryFromValue(v string) (*UsageMetricCategory, error) {
	ev := UsageMetricCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UsageMetricCategory: valid values are %v", v, allowedUsageMetricCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UsageMetricCategory) IsValid() bool {
	for _, existing := range allowedUsageMetricCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageMetricCategory value.
func (v UsageMetricCategory) Ptr() *UsageMetricCategory {
	return &v
}
