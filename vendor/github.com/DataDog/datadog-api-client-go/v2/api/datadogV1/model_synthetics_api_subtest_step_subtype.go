// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAPISubtestStepSubtype The subtype of the Synthetic multi-step API subtest step.
type SyntheticsAPISubtestStepSubtype string

// List of SyntheticsAPISubtestStepSubtype.
const (
	SYNTHETICSAPISUBTESTSTEPSUBTYPE_PLAY_SUB_TEST SyntheticsAPISubtestStepSubtype = "playSubTest"
)

var allowedSyntheticsAPISubtestStepSubtypeEnumValues = []SyntheticsAPISubtestStepSubtype{
	SYNTHETICSAPISUBTESTSTEPSUBTYPE_PLAY_SUB_TEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAPISubtestStepSubtype) GetAllowedValues() []SyntheticsAPISubtestStepSubtype {
	return allowedSyntheticsAPISubtestStepSubtypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAPISubtestStepSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAPISubtestStepSubtype(value)
	return nil
}

// NewSyntheticsAPISubtestStepSubtypeFromValue returns a pointer to a valid SyntheticsAPISubtestStepSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAPISubtestStepSubtypeFromValue(v string) (*SyntheticsAPISubtestStepSubtype, error) {
	ev := SyntheticsAPISubtestStepSubtype(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAPISubtestStepSubtype: valid values are %v", v, allowedSyntheticsAPISubtestStepSubtypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAPISubtestStepSubtype) IsValid() bool {
	for _, existing := range allowedSyntheticsAPISubtestStepSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAPISubtestStepSubtype value.
func (v SyntheticsAPISubtestStepSubtype) Ptr() *SyntheticsAPISubtestStepSubtype {
	return &v
}
