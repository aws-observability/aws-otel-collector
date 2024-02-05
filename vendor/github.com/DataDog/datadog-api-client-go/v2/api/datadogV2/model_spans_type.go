// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansType Type of the span.
type SpansType string

// List of SpansType.
const (
	SPANSTYPE_SPANS SpansType = "spans"
)

var allowedSpansTypeEnumValues = []SpansType{
	SPANSTYPE_SPANS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansType) GetAllowedValues() []SpansType {
	return allowedSpansTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansType(value)
	return nil
}

// NewSpansTypeFromValue returns a pointer to a valid SpansType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansTypeFromValue(v string) (*SpansType, error) {
	ev := SpansType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansType: valid values are %v", v, allowedSpansTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansType) IsValid() bool {
	for _, existing := range allowedSpansTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansType value.
func (v SpansType) Ptr() *SpansType {
	return &v
}
