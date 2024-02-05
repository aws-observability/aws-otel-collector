// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAPIStepSubtype The subtype of the Synthetic multistep API test step, currently only supporting `http`.
type SyntheticsAPIStepSubtype string

// List of SyntheticsAPIStepSubtype.
const (
	SYNTHETICSAPISTEPSUBTYPE_HTTP SyntheticsAPIStepSubtype = "http"
)

var allowedSyntheticsAPIStepSubtypeEnumValues = []SyntheticsAPIStepSubtype{
	SYNTHETICSAPISTEPSUBTYPE_HTTP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAPIStepSubtype) GetAllowedValues() []SyntheticsAPIStepSubtype {
	return allowedSyntheticsAPIStepSubtypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAPIStepSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAPIStepSubtype(value)
	return nil
}

// NewSyntheticsAPIStepSubtypeFromValue returns a pointer to a valid SyntheticsAPIStepSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAPIStepSubtypeFromValue(v string) (*SyntheticsAPIStepSubtype, error) {
	ev := SyntheticsAPIStepSubtype(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAPIStepSubtype: valid values are %v", v, allowedSyntheticsAPIStepSubtypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAPIStepSubtype) IsValid() bool {
	for _, existing := range allowedSyntheticsAPIStepSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAPIStepSubtype value.
func (v SyntheticsAPIStepSubtype) Ptr() *SyntheticsAPIStepSubtype {
	return &v
}
