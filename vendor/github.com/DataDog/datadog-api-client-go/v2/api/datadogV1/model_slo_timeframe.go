// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOTimeframe The SLO time window options. Note that "custom" is not a valid option for creating
// or updating SLOs. It is only used when querying SLO history over custom timeframes.
type SLOTimeframe string

// List of SLOTimeframe.
const (
	SLOTIMEFRAME_SEVEN_DAYS  SLOTimeframe = "7d"
	SLOTIMEFRAME_THIRTY_DAYS SLOTimeframe = "30d"
	SLOTIMEFRAME_NINETY_DAYS SLOTimeframe = "90d"
	SLOTIMEFRAME_CUSTOM      SLOTimeframe = "custom"
)

var allowedSLOTimeframeEnumValues = []SLOTimeframe{
	SLOTIMEFRAME_SEVEN_DAYS,
	SLOTIMEFRAME_THIRTY_DAYS,
	SLOTIMEFRAME_NINETY_DAYS,
	SLOTIMEFRAME_CUSTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SLOTimeframe) GetAllowedValues() []SLOTimeframe {
	return allowedSLOTimeframeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SLOTimeframe) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SLOTimeframe(value)
	return nil
}

// NewSLOTimeframeFromValue returns a pointer to a valid SLOTimeframe
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSLOTimeframeFromValue(v string) (*SLOTimeframe, error) {
	ev := SLOTimeframe(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SLOTimeframe: valid values are %v", v, allowedSLOTimeframeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SLOTimeframe) IsValid() bool {
	for _, existing := range allowedSLOTimeframeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SLOTimeframe value.
func (v SLOTimeframe) Ptr() *SLOTimeframe {
	return &v
}
