// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyRumQueryMode Sankey mode for RUM queries.
type SankeyRumQueryMode string

// List of SankeyRumQueryMode.
const (
	SANKEYRUMQUERYMODE_SOURCE SankeyRumQueryMode = "source"
	SANKEYRUMQUERYMODE_TARGET SankeyRumQueryMode = "target"
)

var allowedSankeyRumQueryModeEnumValues = []SankeyRumQueryMode{
	SANKEYRUMQUERYMODE_SOURCE,
	SANKEYRUMQUERYMODE_TARGET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SankeyRumQueryMode) GetAllowedValues() []SankeyRumQueryMode {
	return allowedSankeyRumQueryModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SankeyRumQueryMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SankeyRumQueryMode(value)
	return nil
}

// NewSankeyRumQueryModeFromValue returns a pointer to a valid SankeyRumQueryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSankeyRumQueryModeFromValue(v string) (*SankeyRumQueryMode, error) {
	ev := SankeyRumQueryMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SankeyRumQueryMode: valid values are %v", v, allowedSankeyRumQueryModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SankeyRumQueryMode) IsValid() bool {
	for _, existing := range allowedSankeyRumQueryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SankeyRumQueryMode value.
func (v SankeyRumQueryMode) Ptr() *SankeyRumQueryMode {
	return &v
}
