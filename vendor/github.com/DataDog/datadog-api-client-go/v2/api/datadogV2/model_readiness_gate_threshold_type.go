// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReadinessGateThresholdType The definition of `ReadinessGateThresholdType` object.
type ReadinessGateThresholdType string

// List of ReadinessGateThresholdType.
const (
	READINESSGATETHRESHOLDTYPE_ANY ReadinessGateThresholdType = "ANY"
	READINESSGATETHRESHOLDTYPE_ALL ReadinessGateThresholdType = "ALL"
)

var allowedReadinessGateThresholdTypeEnumValues = []ReadinessGateThresholdType{
	READINESSGATETHRESHOLDTYPE_ANY,
	READINESSGATETHRESHOLDTYPE_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReadinessGateThresholdType) GetAllowedValues() []ReadinessGateThresholdType {
	return allowedReadinessGateThresholdTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReadinessGateThresholdType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReadinessGateThresholdType(value)
	return nil
}

// NewReadinessGateThresholdTypeFromValue returns a pointer to a valid ReadinessGateThresholdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReadinessGateThresholdTypeFromValue(v string) (*ReadinessGateThresholdType, error) {
	ev := ReadinessGateThresholdType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReadinessGateThresholdType: valid values are %v", v, allowedReadinessGateThresholdTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReadinessGateThresholdType) IsValid() bool {
	for _, existing := range allowedReadinessGateThresholdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReadinessGateThresholdType value.
func (v ReadinessGateThresholdType) Ptr() *ReadinessGateThresholdType {
	return &v
}
