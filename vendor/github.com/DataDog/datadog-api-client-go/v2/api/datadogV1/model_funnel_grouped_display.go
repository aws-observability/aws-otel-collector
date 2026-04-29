// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelGroupedDisplay Display mode for grouped funnel results.
type FunnelGroupedDisplay string

// List of FunnelGroupedDisplay.
const (
	FUNNELGROUPEDDISPLAY_STACKED      FunnelGroupedDisplay = "stacked"
	FUNNELGROUPEDDISPLAY_SIDE_BY_SIDE FunnelGroupedDisplay = "side_by_side"
)

var allowedFunnelGroupedDisplayEnumValues = []FunnelGroupedDisplay{
	FUNNELGROUPEDDISPLAY_STACKED,
	FUNNELGROUPEDDISPLAY_SIDE_BY_SIDE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FunnelGroupedDisplay) GetAllowedValues() []FunnelGroupedDisplay {
	return allowedFunnelGroupedDisplayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FunnelGroupedDisplay) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelGroupedDisplay(value)
	return nil
}

// NewFunnelGroupedDisplayFromValue returns a pointer to a valid FunnelGroupedDisplay
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFunnelGroupedDisplayFromValue(v string) (*FunnelGroupedDisplay, error) {
	ev := FunnelGroupedDisplay(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FunnelGroupedDisplay: valid values are %v", v, allowedFunnelGroupedDisplayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FunnelGroupedDisplay) IsValid() bool {
	for _, existing := range allowedFunnelGroupedDisplayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelGroupedDisplay value.
func (v FunnelGroupedDisplay) Ptr() *FunnelGroupedDisplay {
	return &v
}
