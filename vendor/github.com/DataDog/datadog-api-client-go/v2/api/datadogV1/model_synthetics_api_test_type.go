// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAPITestType Type of the Synthetic test, `api`.
type SyntheticsAPITestType string

// List of SyntheticsAPITestType.
const (
	SYNTHETICSAPITESTTYPE_API SyntheticsAPITestType = "api"
)

var allowedSyntheticsAPITestTypeEnumValues = []SyntheticsAPITestType{
	SYNTHETICSAPITESTTYPE_API,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAPITestType) GetAllowedValues() []SyntheticsAPITestType {
	return allowedSyntheticsAPITestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAPITestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAPITestType(value)
	return nil
}

// NewSyntheticsAPITestTypeFromValue returns a pointer to a valid SyntheticsAPITestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAPITestTypeFromValue(v string) (*SyntheticsAPITestType, error) {
	ev := SyntheticsAPITestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAPITestType: valid values are %v", v, allowedSyntheticsAPITestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAPITestType) IsValid() bool {
	for _, existing := range allowedSyntheticsAPITestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAPITestType value.
func (v SyntheticsAPITestType) Ptr() *SyntheticsAPITestType {
	return &v
}
