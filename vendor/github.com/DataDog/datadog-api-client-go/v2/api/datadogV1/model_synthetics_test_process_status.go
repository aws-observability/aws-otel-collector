// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestProcessStatus Status of a Synthetic test.
type SyntheticsTestProcessStatus string

// List of SyntheticsTestProcessStatus.
const (
	SYNTHETICSTESTPROCESSSTATUS_NOT_SCHEDULED       SyntheticsTestProcessStatus = "not_scheduled"
	SYNTHETICSTESTPROCESSSTATUS_SCHEDULED           SyntheticsTestProcessStatus = "scheduled"
	SYNTHETICSTESTPROCESSSTATUS_FINISHED            SyntheticsTestProcessStatus = "finished"
	SYNTHETICSTESTPROCESSSTATUS_FINISHED_WITH_ERROR SyntheticsTestProcessStatus = "finished_with_error"
)

var allowedSyntheticsTestProcessStatusEnumValues = []SyntheticsTestProcessStatus{
	SYNTHETICSTESTPROCESSSTATUS_NOT_SCHEDULED,
	SYNTHETICSTESTPROCESSSTATUS_SCHEDULED,
	SYNTHETICSTESTPROCESSSTATUS_FINISHED,
	SYNTHETICSTESTPROCESSSTATUS_FINISHED_WITH_ERROR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestProcessStatus) GetAllowedValues() []SyntheticsTestProcessStatus {
	return allowedSyntheticsTestProcessStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestProcessStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestProcessStatus(value)
	return nil
}

// NewSyntheticsTestProcessStatusFromValue returns a pointer to a valid SyntheticsTestProcessStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestProcessStatusFromValue(v string) (*SyntheticsTestProcessStatus, error) {
	ev := SyntheticsTestProcessStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestProcessStatus: valid values are %v", v, allowedSyntheticsTestProcessStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestProcessStatus) IsValid() bool {
	for _, existing := range allowedSyntheticsTestProcessStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestProcessStatus value.
func (v SyntheticsTestProcessStatus) Ptr() *SyntheticsTestProcessStatus {
	return &v
}
