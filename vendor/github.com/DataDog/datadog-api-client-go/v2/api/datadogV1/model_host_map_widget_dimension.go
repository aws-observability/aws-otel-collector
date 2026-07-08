// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetDimension Visual dimension driven by a formula in the infrastructure host map widget.
type HostMapWidgetDimension string

// List of HostMapWidgetDimension.
const (
	HOSTMAPWIDGETDIMENSION_NODE HostMapWidgetDimension = "node"
	HOSTMAPWIDGETDIMENSION_FILL HostMapWidgetDimension = "fill"
	HOSTMAPWIDGETDIMENSION_SIZE HostMapWidgetDimension = "size"
)

var allowedHostMapWidgetDimensionEnumValues = []HostMapWidgetDimension{
	HOSTMAPWIDGETDIMENSION_NODE,
	HOSTMAPWIDGETDIMENSION_FILL,
	HOSTMAPWIDGETDIMENSION_SIZE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostMapWidgetDimension) GetAllowedValues() []HostMapWidgetDimension {
	return allowedHostMapWidgetDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostMapWidgetDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostMapWidgetDimension(value)
	return nil
}

// NewHostMapWidgetDimensionFromValue returns a pointer to a valid HostMapWidgetDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostMapWidgetDimensionFromValue(v string) (*HostMapWidgetDimension, error) {
	ev := HostMapWidgetDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostMapWidgetDimension: valid values are %v", v, allowedHostMapWidgetDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostMapWidgetDimension) IsValid() bool {
	for _, existing := range allowedHostMapWidgetDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostMapWidgetDimension value.
func (v HostMapWidgetDimension) Ptr() *HostMapWidgetDimension {
	return &v
}
