// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansMetricType The type of resource. The value should always be spans_metrics.
type SpansMetricType string

// List of SpansMetricType.
const (
	SPANSMETRICTYPE_SPANS_METRICS SpansMetricType = "spans_metrics"
)

var allowedSpansMetricTypeEnumValues = []SpansMetricType{
	SPANSMETRICTYPE_SPANS_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansMetricType) GetAllowedValues() []SpansMetricType {
	return allowedSpansMetricTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansMetricType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansMetricType(value)
	return nil
}

// NewSpansMetricTypeFromValue returns a pointer to a valid SpansMetricType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansMetricTypeFromValue(v string) (*SpansMetricType, error) {
	ev := SpansMetricType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansMetricType: valid values are %v", v, allowedSpansMetricTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansMetricType) IsValid() bool {
	for _, existing := range allowedSpansMetricTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansMetricType value.
func (v SpansMetricType) Ptr() *SpansMetricType {
	return &v
}
