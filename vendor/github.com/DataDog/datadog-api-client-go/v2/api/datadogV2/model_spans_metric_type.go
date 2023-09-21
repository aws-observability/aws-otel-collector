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
	err := json.Unmarshal(src, &value)
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
<<<<<<< HEAD

// NullableSpansMetricType handles when a null is used for SpansMetricType.
type NullableSpansMetricType struct {
	value *SpansMetricType
	isSet bool
}

// Get returns the associated value.
func (v NullableSpansMetricType) Get() *SpansMetricType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableSpansMetricType) Set(val *SpansMetricType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableSpansMetricType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableSpansMetricType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSpansMetricType initializes the struct as if Set has been called.
func NewNullableSpansMetricType(val *SpansMetricType) *NullableSpansMetricType {
	return &NullableSpansMetricType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableSpansMetricType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableSpansMetricType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
