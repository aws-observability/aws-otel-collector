// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricTagConfigurationMetricTypeCategory The metric's type category.
type MetricTagConfigurationMetricTypeCategory string

// List of MetricTagConfigurationMetricTypeCategory.
const (
	METRICTAGCONFIGURATIONMETRICTYPECATEGORY_NON_DISTRIBUTION MetricTagConfigurationMetricTypeCategory = "non_distribution"
	METRICTAGCONFIGURATIONMETRICTYPECATEGORY_DISTRIBUTION     MetricTagConfigurationMetricTypeCategory = "distribution"
)

var allowedMetricTagConfigurationMetricTypeCategoryEnumValues = []MetricTagConfigurationMetricTypeCategory{
	METRICTAGCONFIGURATIONMETRICTYPECATEGORY_NON_DISTRIBUTION,
	METRICTAGCONFIGURATIONMETRICTYPECATEGORY_DISTRIBUTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricTagConfigurationMetricTypeCategory) GetAllowedValues() []MetricTagConfigurationMetricTypeCategory {
	return allowedMetricTagConfigurationMetricTypeCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricTagConfigurationMetricTypeCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricTagConfigurationMetricTypeCategory(value)
	return nil
}

// NewMetricTagConfigurationMetricTypeCategoryFromValue returns a pointer to a valid MetricTagConfigurationMetricTypeCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricTagConfigurationMetricTypeCategoryFromValue(v string) (*MetricTagConfigurationMetricTypeCategory, error) {
	ev := MetricTagConfigurationMetricTypeCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricTagConfigurationMetricTypeCategory: valid values are %v", v, allowedMetricTagConfigurationMetricTypeCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricTagConfigurationMetricTypeCategory) IsValid() bool {
	for _, existing := range allowedMetricTagConfigurationMetricTypeCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricTagConfigurationMetricTypeCategory value.
func (v MetricTagConfigurationMetricTypeCategory) Ptr() *MetricTagConfigurationMetricTypeCategory {
	return &v
}
