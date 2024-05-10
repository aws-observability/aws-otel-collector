// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOTimeSliceComparator The comparator used to compare the SLI value to the threshold.
type SLOTimeSliceComparator string

// List of SLOTimeSliceComparator.
const (
	SLOTIMESLICECOMPARATOR_GREATER       SLOTimeSliceComparator = ">"
	SLOTIMESLICECOMPARATOR_GREATER_EQUAL SLOTimeSliceComparator = ">="
	SLOTIMESLICECOMPARATOR_LESS          SLOTimeSliceComparator = "<"
	SLOTIMESLICECOMPARATOR_LESS_EQUAL    SLOTimeSliceComparator = "<="
)

var allowedSLOTimeSliceComparatorEnumValues = []SLOTimeSliceComparator{
	SLOTIMESLICECOMPARATOR_GREATER,
	SLOTIMESLICECOMPARATOR_GREATER_EQUAL,
	SLOTIMESLICECOMPARATOR_LESS,
	SLOTIMESLICECOMPARATOR_LESS_EQUAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SLOTimeSliceComparator) GetAllowedValues() []SLOTimeSliceComparator {
	return allowedSLOTimeSliceComparatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SLOTimeSliceComparator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SLOTimeSliceComparator(value)
	return nil
}

// NewSLOTimeSliceComparatorFromValue returns a pointer to a valid SLOTimeSliceComparator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSLOTimeSliceComparatorFromValue(v string) (*SLOTimeSliceComparator, error) {
	ev := SLOTimeSliceComparator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SLOTimeSliceComparator: valid values are %v", v, allowedSLOTimeSliceComparatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SLOTimeSliceComparator) IsValid() bool {
	for _, existing := range allowedSLOTimeSliceComparatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SLOTimeSliceComparator value.
func (v SLOTimeSliceComparator) Ptr() *SLOTimeSliceComparator {
	return &v
}
