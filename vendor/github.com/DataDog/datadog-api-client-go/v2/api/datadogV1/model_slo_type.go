// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOType The type of the service level objective.
type SLOType string

// List of SLOType.
const (
	SLOTYPE_METRIC     SLOType = "metric"
	SLOTYPE_MONITOR    SLOType = "monitor"
	SLOTYPE_TIME_SLICE SLOType = "time_slice"
)

var allowedSLOTypeEnumValues = []SLOType{
	SLOTYPE_METRIC,
	SLOTYPE_MONITOR,
	SLOTYPE_TIME_SLICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SLOType) GetAllowedValues() []SLOType {
	return allowedSLOTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SLOType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SLOType(value)
	return nil
}

// NewSLOTypeFromValue returns a pointer to a valid SLOType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSLOTypeFromValue(v string) (*SLOType, error) {
	ev := SLOType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SLOType: valid values are %v", v, allowedSLOTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SLOType) IsValid() bool {
	for _, existing := range allowedSLOTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SLOType value.
func (v SLOType) Ptr() *SLOType {
	return &v
}
