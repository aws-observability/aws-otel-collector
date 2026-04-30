// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultStatus Status of a Synthetic test result.
type SyntheticsTestResultStatus string

// List of SyntheticsTestResultStatus.
const (
	SYNTHETICSTESTRESULTSTATUS_PASSED  SyntheticsTestResultStatus = "passed"
	SYNTHETICSTESTRESULTSTATUS_FAILED  SyntheticsTestResultStatus = "failed"
	SYNTHETICSTESTRESULTSTATUS_NO_DATA SyntheticsTestResultStatus = "no_data"
)

var allowedSyntheticsTestResultStatusEnumValues = []SyntheticsTestResultStatus{
	SYNTHETICSTESTRESULTSTATUS_PASSED,
	SYNTHETICSTESTRESULTSTATUS_FAILED,
	SYNTHETICSTESTRESULTSTATUS_NO_DATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestResultStatus) GetAllowedValues() []SyntheticsTestResultStatus {
	return allowedSyntheticsTestResultStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestResultStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestResultStatus(value)
	return nil
}

// NewSyntheticsTestResultStatusFromValue returns a pointer to a valid SyntheticsTestResultStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestResultStatusFromValue(v string) (*SyntheticsTestResultStatus, error) {
	ev := SyntheticsTestResultStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestResultStatus: valid values are %v", v, allowedSyntheticsTestResultStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestResultStatus) IsValid() bool {
	for _, existing := range allowedSyntheticsTestResultStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestResultStatus value.
func (v SyntheticsTestResultStatus) Ptr() *SyntheticsTestResultStatus {
	return &v
}
