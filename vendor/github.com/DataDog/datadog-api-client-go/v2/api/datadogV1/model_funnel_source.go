// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelSource Source from which to query items to display in the funnel.
type FunnelSource string

// List of FunnelSource.
const (
	FUNNELSOURCE_RUM FunnelSource = "rum"
)

var allowedFunnelSourceEnumValues = []FunnelSource{
	FUNNELSOURCE_RUM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FunnelSource) GetAllowedValues() []FunnelSource {
	return allowedFunnelSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FunnelSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelSource(value)
	return nil
}

// NewFunnelSourceFromValue returns a pointer to a valid FunnelSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFunnelSourceFromValue(v string) (*FunnelSource, error) {
	ev := FunnelSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FunnelSource: valid values are %v", v, allowedFunnelSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FunnelSource) IsValid() bool {
	for _, existing := range allowedFunnelSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelSource value.
func (v FunnelSource) Ptr() *FunnelSource {
	return &v
}
