// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComparisonDurationType The comparison window type.
type ComparisonDurationType string

// List of ComparisonDurationType.
const (
	COMPARISONDURATIONTYPE_PREVIOUS_TIMEFRAME ComparisonDurationType = "previous_timeframe"
	COMPARISONDURATIONTYPE_CUSTOM_TIMEFRAME   ComparisonDurationType = "custom_timeframe"
	COMPARISONDURATIONTYPE_PREVIOUS_DAY       ComparisonDurationType = "previous_day"
	COMPARISONDURATIONTYPE_PREVIOUS_WEEK      ComparisonDurationType = "previous_week"
	COMPARISONDURATIONTYPE_PREVIOUS_MONTH     ComparisonDurationType = "previous_month"
)

var allowedComparisonDurationTypeEnumValues = []ComparisonDurationType{
	COMPARISONDURATIONTYPE_PREVIOUS_TIMEFRAME,
	COMPARISONDURATIONTYPE_CUSTOM_TIMEFRAME,
	COMPARISONDURATIONTYPE_PREVIOUS_DAY,
	COMPARISONDURATIONTYPE_PREVIOUS_WEEK,
	COMPARISONDURATIONTYPE_PREVIOUS_MONTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ComparisonDurationType) GetAllowedValues() []ComparisonDurationType {
	return allowedComparisonDurationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ComparisonDurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ComparisonDurationType(value)
	return nil
}

// NewComparisonDurationTypeFromValue returns a pointer to a valid ComparisonDurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewComparisonDurationTypeFromValue(v string) (*ComparisonDurationType, error) {
	ev := ComparisonDurationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ComparisonDurationType: valid values are %v", v, allowedComparisonDurationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ComparisonDurationType) IsValid() bool {
	for _, existing := range allowedComparisonDurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ComparisonDurationType value.
func (v ComparisonDurationType) Ptr() *ComparisonDurationType {
	return &v
}
