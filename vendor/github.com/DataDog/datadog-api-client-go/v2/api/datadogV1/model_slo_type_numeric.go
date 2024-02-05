// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOTypeNumeric A numeric representation of the type of the service level objective (`0` for
// monitor, `1` for metric). Always included in service level objective responses.
// Ignored in create/update requests.
type SLOTypeNumeric int32

// List of SLOTypeNumeric.
const (
	SLOTYPENUMERIC_MONITOR    SLOTypeNumeric = 0
	SLOTYPENUMERIC_METRIC     SLOTypeNumeric = 1
	SLOTYPENUMERIC_TIME_SLICE SLOTypeNumeric = 2
)

var allowedSLOTypeNumericEnumValues = []SLOTypeNumeric{
	SLOTYPENUMERIC_MONITOR,
	SLOTYPENUMERIC_METRIC,
	SLOTYPENUMERIC_TIME_SLICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SLOTypeNumeric) GetAllowedValues() []SLOTypeNumeric {
	return allowedSLOTypeNumericEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SLOTypeNumeric) UnmarshalJSON(src []byte) error {
	var value int32
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SLOTypeNumeric(value)
	return nil
}

// NewSLOTypeNumericFromValue returns a pointer to a valid SLOTypeNumeric
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSLOTypeNumericFromValue(v int32) (*SLOTypeNumeric, error) {
	ev := SLOTypeNumeric(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SLOTypeNumeric: valid values are %v", v, allowedSLOTypeNumericEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SLOTypeNumeric) IsValid() bool {
	for _, existing := range allowedSLOTypeNumericEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SLOTypeNumeric value.
func (v SLOTypeNumeric) Ptr() *SLOTypeNumeric {
	return &v
}
