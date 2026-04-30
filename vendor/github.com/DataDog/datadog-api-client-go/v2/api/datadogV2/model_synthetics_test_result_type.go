// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultType Type of the Synthetic test result resource, `result`.
type SyntheticsTestResultType string

// List of SyntheticsTestResultType.
const (
	SYNTHETICSTESTRESULTTYPE_RESULT SyntheticsTestResultType = "result"
)

var allowedSyntheticsTestResultTypeEnumValues = []SyntheticsTestResultType{
	SYNTHETICSTESTRESULTTYPE_RESULT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestResultType) GetAllowedValues() []SyntheticsTestResultType {
	return allowedSyntheticsTestResultTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestResultType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestResultType(value)
	return nil
}

// NewSyntheticsTestResultTypeFromValue returns a pointer to a valid SyntheticsTestResultType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestResultTypeFromValue(v string) (*SyntheticsTestResultType, error) {
	ev := SyntheticsTestResultType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestResultType: valid values are %v", v, allowedSyntheticsTestResultTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestResultType) IsValid() bool {
	for _, existing := range allowedSyntheticsTestResultTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestResultType value.
func (v SyntheticsTestResultType) Ptr() *SyntheticsTestResultType {
	return &v
}
