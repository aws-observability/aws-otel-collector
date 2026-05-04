// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionDataSource Data source for retention queries.
type RetentionDataSource string

// List of RetentionDataSource.
const (
	RETENTIONDATASOURCE_PRODUCT_ANALYTICS_RETENTION RetentionDataSource = "product_analytics_retention"
)

var allowedRetentionDataSourceEnumValues = []RetentionDataSource{
	RETENTIONDATASOURCE_PRODUCT_ANALYTICS_RETENTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionDataSource) GetAllowedValues() []RetentionDataSource {
	return allowedRetentionDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionDataSource(value)
	return nil
}

// NewRetentionDataSourceFromValue returns a pointer to a valid RetentionDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionDataSourceFromValue(v string) (*RetentionDataSource, error) {
	ev := RetentionDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionDataSource: valid values are %v", v, allowedRetentionDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionDataSource) IsValid() bool {
	for _, existing := range allowedRetentionDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionDataSource value.
func (v RetentionDataSource) Ptr() *RetentionDataSource {
	return &v
}
