// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestVersionChangeType Type of the version metadata resource.
type SyntheticsTestVersionChangeType string

// List of SyntheticsTestVersionChangeType.
const (
	SYNTHETICSTESTVERSIONCHANGETYPE_VERSION_METADATA SyntheticsTestVersionChangeType = "version_metadata"
)

var allowedSyntheticsTestVersionChangeTypeEnumValues = []SyntheticsTestVersionChangeType{
	SYNTHETICSTESTVERSIONCHANGETYPE_VERSION_METADATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestVersionChangeType) GetAllowedValues() []SyntheticsTestVersionChangeType {
	return allowedSyntheticsTestVersionChangeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestVersionChangeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestVersionChangeType(value)
	return nil
}

// NewSyntheticsTestVersionChangeTypeFromValue returns a pointer to a valid SyntheticsTestVersionChangeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestVersionChangeTypeFromValue(v string) (*SyntheticsTestVersionChangeType, error) {
	ev := SyntheticsTestVersionChangeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestVersionChangeType: valid values are %v", v, allowedSyntheticsTestVersionChangeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestVersionChangeType) IsValid() bool {
	for _, existing := range allowedSyntheticsTestVersionChangeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestVersionChangeType value.
func (v SyntheticsTestVersionChangeType) Ptr() *SyntheticsTestVersionChangeType {
	return &v
}
