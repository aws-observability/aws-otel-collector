// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestPauseStatus Define whether you want to start (`live`) or pause (`paused`) a
// Synthetic test.
type SyntheticsTestPauseStatus string

// List of SyntheticsTestPauseStatus.
const (
	SYNTHETICSTESTPAUSESTATUS_LIVE   SyntheticsTestPauseStatus = "live"
	SYNTHETICSTESTPAUSESTATUS_PAUSED SyntheticsTestPauseStatus = "paused"
)

var allowedSyntheticsTestPauseStatusEnumValues = []SyntheticsTestPauseStatus{
	SYNTHETICSTESTPAUSESTATUS_LIVE,
	SYNTHETICSTESTPAUSESTATUS_PAUSED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestPauseStatus) GetAllowedValues() []SyntheticsTestPauseStatus {
	return allowedSyntheticsTestPauseStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestPauseStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestPauseStatus(value)
	return nil
}

// NewSyntheticsTestPauseStatusFromValue returns a pointer to a valid SyntheticsTestPauseStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestPauseStatusFromValue(v string) (*SyntheticsTestPauseStatus, error) {
	ev := SyntheticsTestPauseStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestPauseStatus: valid values are %v", v, allowedSyntheticsTestPauseStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestPauseStatus) IsValid() bool {
	for _, existing := range allowedSyntheticsTestPauseStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestPauseStatus value.
func (v SyntheticsTestPauseStatus) Ptr() *SyntheticsTestPauseStatus {
	return &v
}
