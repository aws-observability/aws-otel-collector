// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SlosQueryType The type of SLO definition being queried.
type SlosQueryType string

// List of SlosQueryType.
const (
	SLOSQUERYTYPE_METRIC     SlosQueryType = "metric"
	SLOSQUERYTYPE_TIME_SLICE SlosQueryType = "time_slice"
	SLOSQUERYTYPE_MONITOR    SlosQueryType = "monitor"
)

var allowedSlosQueryTypeEnumValues = []SlosQueryType{
	SLOSQUERYTYPE_METRIC,
	SLOSQUERYTYPE_TIME_SLICE,
	SLOSQUERYTYPE_MONITOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SlosQueryType) GetAllowedValues() []SlosQueryType {
	return allowedSlosQueryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SlosQueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SlosQueryType(value)
	return nil
}

// NewSlosQueryTypeFromValue returns a pointer to a valid SlosQueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSlosQueryTypeFromValue(v string) (*SlosQueryType, error) {
	ev := SlosQueryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SlosQueryType: valid values are %v", v, allowedSlosQueryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SlosQueryType) IsValid() bool {
	for _, existing := range allowedSlosQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SlosQueryType value.
func (v SlosQueryType) Ptr() *SlosQueryType {
	return &v
}
