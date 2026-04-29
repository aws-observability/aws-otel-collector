// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultRunType The type of run for a Synthetic test result.
type SyntheticsTestResultRunType string

// List of SyntheticsTestResultRunType.
const (
	SYNTHETICSTESTRESULTRUNTYPE_SCHEDULED SyntheticsTestResultRunType = "scheduled"
	SYNTHETICSTESTRESULTRUNTYPE_FAST      SyntheticsTestResultRunType = "fast"
	SYNTHETICSTESTRESULTRUNTYPE_CI        SyntheticsTestResultRunType = "ci"
	SYNTHETICSTESTRESULTRUNTYPE_TRIGGERED SyntheticsTestResultRunType = "triggered"
)

var allowedSyntheticsTestResultRunTypeEnumValues = []SyntheticsTestResultRunType{
	SYNTHETICSTESTRESULTRUNTYPE_SCHEDULED,
	SYNTHETICSTESTRESULTRUNTYPE_FAST,
	SYNTHETICSTESTRESULTRUNTYPE_CI,
	SYNTHETICSTESTRESULTRUNTYPE_TRIGGERED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestResultRunType) GetAllowedValues() []SyntheticsTestResultRunType {
	return allowedSyntheticsTestResultRunTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestResultRunType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestResultRunType(value)
	return nil
}

// NewSyntheticsTestResultRunTypeFromValue returns a pointer to a valid SyntheticsTestResultRunType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestResultRunTypeFromValue(v string) (*SyntheticsTestResultRunType, error) {
	ev := SyntheticsTestResultRunType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestResultRunType: valid values are %v", v, allowedSyntheticsTestResultRunTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestResultRunType) IsValid() bool {
	for _, existing := range allowedSyntheticsTestResultRunTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestResultRunType value.
func (v SyntheticsTestResultRunType) Ptr() *SyntheticsTestResultRunType {
	return &v
}
