// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserJourneyFormulaComputeMetric Metric for User Journey formula compute. `__dd.conversion` and `__dd.conversion_rate` accept `count` and `cardinality` as aggregations. `__dd.time_to_convert` accepts `avg`, `median`, `pc75`, `pc95`, `pc98`, `pc99`, `min`, and `max`.
type UserJourneyFormulaComputeMetric string

// List of UserJourneyFormulaComputeMetric.
const (
	USERJOURNEYFORMULACOMPUTEMETRIC_CONVERSION      UserJourneyFormulaComputeMetric = "__dd.conversion"
	USERJOURNEYFORMULACOMPUTEMETRIC_CONVERSION_RATE UserJourneyFormulaComputeMetric = "__dd.conversion_rate"
	USERJOURNEYFORMULACOMPUTEMETRIC_TIME_TO_CONVERT UserJourneyFormulaComputeMetric = "__dd.time_to_convert"
)

var allowedUserJourneyFormulaComputeMetricEnumValues = []UserJourneyFormulaComputeMetric{
	USERJOURNEYFORMULACOMPUTEMETRIC_CONVERSION,
	USERJOURNEYFORMULACOMPUTEMETRIC_CONVERSION_RATE,
	USERJOURNEYFORMULACOMPUTEMETRIC_TIME_TO_CONVERT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserJourneyFormulaComputeMetric) GetAllowedValues() []UserJourneyFormulaComputeMetric {
	return allowedUserJourneyFormulaComputeMetricEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserJourneyFormulaComputeMetric) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserJourneyFormulaComputeMetric(value)
	return nil
}

// NewUserJourneyFormulaComputeMetricFromValue returns a pointer to a valid UserJourneyFormulaComputeMetric
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserJourneyFormulaComputeMetricFromValue(v string) (*UserJourneyFormulaComputeMetric, error) {
	ev := UserJourneyFormulaComputeMetric(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserJourneyFormulaComputeMetric: valid values are %v", v, allowedUserJourneyFormulaComputeMetricEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserJourneyFormulaComputeMetric) IsValid() bool {
	for _, existing := range allowedUserJourneyFormulaComputeMetricEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserJourneyFormulaComputeMetric value.
func (v UserJourneyFormulaComputeMetric) Ptr() *UserJourneyFormulaComputeMetric {
	return &v
}
