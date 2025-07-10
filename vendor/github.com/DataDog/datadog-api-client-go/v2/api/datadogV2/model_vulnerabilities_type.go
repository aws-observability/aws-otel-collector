// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VulnerabilitiesType The JSON:API type.
type VulnerabilitiesType string

// List of VulnerabilitiesType.
const (
	VULNERABILITIESTYPE_VULNERABILITIES VulnerabilitiesType = "vulnerabilities"
)

var allowedVulnerabilitiesTypeEnumValues = []VulnerabilitiesType{
	VULNERABILITIESTYPE_VULNERABILITIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *VulnerabilitiesType) GetAllowedValues() []VulnerabilitiesType {
	return allowedVulnerabilitiesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *VulnerabilitiesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = VulnerabilitiesType(value)
	return nil
}

// NewVulnerabilitiesTypeFromValue returns a pointer to a valid VulnerabilitiesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewVulnerabilitiesTypeFromValue(v string) (*VulnerabilitiesType, error) {
	ev := VulnerabilitiesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for VulnerabilitiesType: valid values are %v", v, allowedVulnerabilitiesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v VulnerabilitiesType) IsValid() bool {
	for _, existing := range allowedVulnerabilitiesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VulnerabilitiesType value.
func (v VulnerabilitiesType) Ptr() *VulnerabilitiesType {
	return &v
}
