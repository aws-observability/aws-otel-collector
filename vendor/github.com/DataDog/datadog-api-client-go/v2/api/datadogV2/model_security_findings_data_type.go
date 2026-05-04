// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityFindingsDataType The type of the security finding resource.
type SecurityFindingsDataType string

// List of SecurityFindingsDataType.
const (
	SECURITYFINDINGSDATATYPE_FINDING SecurityFindingsDataType = "finding"
)

var allowedSecurityFindingsDataTypeEnumValues = []SecurityFindingsDataType{
	SECURITYFINDINGSDATATYPE_FINDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityFindingsDataType) GetAllowedValues() []SecurityFindingsDataType {
	return allowedSecurityFindingsDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityFindingsDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityFindingsDataType(value)
	return nil
}

// NewSecurityFindingsDataTypeFromValue returns a pointer to a valid SecurityFindingsDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityFindingsDataTypeFromValue(v string) (*SecurityFindingsDataType, error) {
	ev := SecurityFindingsDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityFindingsDataType: valid values are %v", v, allowedSecurityFindingsDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityFindingsDataType) IsValid() bool {
	for _, existing := range allowedSecurityFindingsDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityFindingsDataType value.
func (v SecurityFindingsDataType) Ptr() *SecurityFindingsDataType {
	return &v
}
