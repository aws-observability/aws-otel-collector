// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
=======
	"fmt"

	"github.com/goccy/go-json"
>>>>>>> main
)

// SpansMetricComputeAggregationType The type of aggregation to use.
type SpansMetricComputeAggregationType string

// List of SpansMetricComputeAggregationType.
const (
	SPANSMETRICCOMPUTEAGGREGATIONTYPE_COUNT        SpansMetricComputeAggregationType = "count"
	SPANSMETRICCOMPUTEAGGREGATIONTYPE_DISTRIBUTION SpansMetricComputeAggregationType = "distribution"
)

var allowedSpansMetricComputeAggregationTypeEnumValues = []SpansMetricComputeAggregationType{
	SPANSMETRICCOMPUTEAGGREGATIONTYPE_COUNT,
	SPANSMETRICCOMPUTEAGGREGATIONTYPE_DISTRIBUTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansMetricComputeAggregationType) GetAllowedValues() []SpansMetricComputeAggregationType {
	return allowedSpansMetricComputeAggregationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansMetricComputeAggregationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansMetricComputeAggregationType(value)
	return nil
}

// NewSpansMetricComputeAggregationTypeFromValue returns a pointer to a valid SpansMetricComputeAggregationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansMetricComputeAggregationTypeFromValue(v string) (*SpansMetricComputeAggregationType, error) {
	ev := SpansMetricComputeAggregationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansMetricComputeAggregationType: valid values are %v", v, allowedSpansMetricComputeAggregationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansMetricComputeAggregationType) IsValid() bool {
	for _, existing := range allowedSpansMetricComputeAggregationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansMetricComputeAggregationType value.
func (v SpansMetricComputeAggregationType) Ptr() *SpansMetricComputeAggregationType {
	return &v
}
<<<<<<< HEAD

// NullableSpansMetricComputeAggregationType handles when a null is used for SpansMetricComputeAggregationType.
type NullableSpansMetricComputeAggregationType struct {
	value *SpansMetricComputeAggregationType
	isSet bool
}

// Get returns the associated value.
func (v NullableSpansMetricComputeAggregationType) Get() *SpansMetricComputeAggregationType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableSpansMetricComputeAggregationType) Set(val *SpansMetricComputeAggregationType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableSpansMetricComputeAggregationType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableSpansMetricComputeAggregationType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSpansMetricComputeAggregationType initializes the struct as if Set has been called.
func NewNullableSpansMetricComputeAggregationType(val *SpansMetricComputeAggregationType) *NullableSpansMetricComputeAggregationType {
	return &NullableSpansMetricComputeAggregationType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableSpansMetricComputeAggregationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableSpansMetricComputeAggregationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
