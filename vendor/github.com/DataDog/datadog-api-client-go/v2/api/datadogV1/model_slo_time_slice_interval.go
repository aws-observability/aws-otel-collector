// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOTimeSliceInterval The interval used when querying data, which defines the size of a time slice.
// Two values are allowed: 60 (1 minute) and 300 (5 minutes).
// If not provided, the value defaults to 300 (5 minutes).
type SLOTimeSliceInterval int32

// List of SLOTimeSliceInterval.
const (
	SLOTIMESLICEINTERVAL_ONE_MINUTE   SLOTimeSliceInterval = 60
	SLOTIMESLICEINTERVAL_FIVE_MINUTES SLOTimeSliceInterval = 300
)

var allowedSLOTimeSliceIntervalEnumValues = []SLOTimeSliceInterval{
	SLOTIMESLICEINTERVAL_ONE_MINUTE,
	SLOTIMESLICEINTERVAL_FIVE_MINUTES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SLOTimeSliceInterval) GetAllowedValues() []SLOTimeSliceInterval {
	return allowedSLOTimeSliceIntervalEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SLOTimeSliceInterval) UnmarshalJSON(src []byte) error {
	var value int32
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SLOTimeSliceInterval(value)
	return nil
}

// NewSLOTimeSliceIntervalFromValue returns a pointer to a valid SLOTimeSliceInterval
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSLOTimeSliceIntervalFromValue(v int32) (*SLOTimeSliceInterval, error) {
	ev := SLOTimeSliceInterval(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SLOTimeSliceInterval: valid values are %v", v, allowedSLOTimeSliceIntervalEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SLOTimeSliceInterval) IsValid() bool {
	for _, existing := range allowedSLOTimeSliceIntervalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SLOTimeSliceInterval value.
func (v SLOTimeSliceInterval) Ptr() *SLOTimeSliceInterval {
	return &v
}
