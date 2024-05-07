// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOCorrectionType SLO correction resource type.
type SLOCorrectionType string

// List of SLOCorrectionType.
const (
	SLOCORRECTIONTYPE_CORRECTION SLOCorrectionType = "correction"
)

var allowedSLOCorrectionTypeEnumValues = []SLOCorrectionType{
	SLOCORRECTIONTYPE_CORRECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SLOCorrectionType) GetAllowedValues() []SLOCorrectionType {
	return allowedSLOCorrectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SLOCorrectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SLOCorrectionType(value)
	return nil
}

// NewSLOCorrectionTypeFromValue returns a pointer to a valid SLOCorrectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSLOCorrectionTypeFromValue(v string) (*SLOCorrectionType, error) {
	ev := SLOCorrectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SLOCorrectionType: valid values are %v", v, allowedSLOCorrectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SLOCorrectionType) IsValid() bool {
	for _, existing := range allowedSLOCorrectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SLOCorrectionType value.
func (v SLOCorrectionType) Ptr() *SLOCorrectionType {
	return &v
}
