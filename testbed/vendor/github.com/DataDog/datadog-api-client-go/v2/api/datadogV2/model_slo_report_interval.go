// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOReportInterval The frequency at which report data is to be generated.
type SLOReportInterval string

// List of SLOReportInterval.
const (
	SLOREPORTINTERVAL_DAILY   SLOReportInterval = "daily"
	SLOREPORTINTERVAL_WEEKLY  SLOReportInterval = "weekly"
	SLOREPORTINTERVAL_MONTHLY SLOReportInterval = "monthly"
)

var allowedSLOReportIntervalEnumValues = []SLOReportInterval{
	SLOREPORTINTERVAL_DAILY,
	SLOREPORTINTERVAL_WEEKLY,
	SLOREPORTINTERVAL_MONTHLY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SLOReportInterval) GetAllowedValues() []SLOReportInterval {
	return allowedSLOReportIntervalEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SLOReportInterval) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SLOReportInterval(value)
	return nil
}

// NewSLOReportIntervalFromValue returns a pointer to a valid SLOReportInterval
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSLOReportIntervalFromValue(v string) (*SLOReportInterval, error) {
	ev := SLOReportInterval(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SLOReportInterval: valid values are %v", v, allowedSLOReportIntervalEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SLOReportInterval) IsValid() bool {
	for _, existing := range allowedSLOReportIntervalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SLOReportInterval value.
func (v SLOReportInterval) Ptr() *SLOReportInterval {
	return &v
}
