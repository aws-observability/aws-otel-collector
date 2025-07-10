// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SBOMType The JSON:API type.
type SBOMType string

// List of SBOMType.
const (
	SBOMTYPE_SBOMS SBOMType = "sboms"
)

var allowedSBOMTypeEnumValues = []SBOMType{
	SBOMTYPE_SBOMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SBOMType) GetAllowedValues() []SBOMType {
	return allowedSBOMTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SBOMType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SBOMType(value)
	return nil
}

// NewSBOMTypeFromValue returns a pointer to a valid SBOMType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSBOMTypeFromValue(v string) (*SBOMType, error) {
	ev := SBOMType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SBOMType: valid values are %v", v, allowedSBOMTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SBOMType) IsValid() bool {
	for _, existing := range allowedSBOMTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SBOMType value.
func (v SBOMType) Ptr() *SBOMType {
	return &v
}
