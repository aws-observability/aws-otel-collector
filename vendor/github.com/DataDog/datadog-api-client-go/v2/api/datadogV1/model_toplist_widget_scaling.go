// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistWidgetScaling Top list widget scaling definition.
type ToplistWidgetScaling string

// List of ToplistWidgetScaling.
const (
	TOPLISTWIDGETSCALING_ABSOLUTE ToplistWidgetScaling = "absolute"
	TOPLISTWIDGETSCALING_RELATIVE ToplistWidgetScaling = "relative"
)

var allowedToplistWidgetScalingEnumValues = []ToplistWidgetScaling{
	TOPLISTWIDGETSCALING_ABSOLUTE,
	TOPLISTWIDGETSCALING_RELATIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistWidgetScaling) GetAllowedValues() []ToplistWidgetScaling {
	return allowedToplistWidgetScalingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistWidgetScaling) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistWidgetScaling(value)
	return nil
}

// NewToplistWidgetScalingFromValue returns a pointer to a valid ToplistWidgetScaling
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistWidgetScalingFromValue(v string) (*ToplistWidgetScaling, error) {
	ev := ToplistWidgetScaling(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistWidgetScaling: valid values are %v", v, allowedToplistWidgetScalingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistWidgetScaling) IsValid() bool {
	for _, existing := range allowedToplistWidgetScalingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistWidgetScaling value.
func (v ToplistWidgetScaling) Ptr() *ToplistWidgetScaling {
	return &v
}
