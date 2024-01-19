// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansComputeType The type of compute.
type SpansComputeType string

// List of SpansComputeType.
const (
	SPANSCOMPUTETYPE_TIMESERIES SpansComputeType = "timeseries"
	SPANSCOMPUTETYPE_TOTAL      SpansComputeType = "total"
)

var allowedSpansComputeTypeEnumValues = []SpansComputeType{
	SPANSCOMPUTETYPE_TIMESERIES,
	SPANSCOMPUTETYPE_TOTAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansComputeType) GetAllowedValues() []SpansComputeType {
	return allowedSpansComputeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansComputeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansComputeType(value)
	return nil
}

// NewSpansComputeTypeFromValue returns a pointer to a valid SpansComputeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansComputeTypeFromValue(v string) (*SpansComputeType, error) {
	ev := SpansComputeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansComputeType: valid values are %v", v, allowedSpansComputeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansComputeType) IsValid() bool {
	for _, existing := range allowedSpansComputeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansComputeType value.
func (v SpansComputeType) Ptr() *SpansComputeType {
	return &v
}
