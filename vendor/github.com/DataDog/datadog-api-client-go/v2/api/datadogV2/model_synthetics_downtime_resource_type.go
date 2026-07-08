// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsDowntimeResourceType The resource type for a Synthetics downtime.
type SyntheticsDowntimeResourceType string

// List of SyntheticsDowntimeResourceType.
const (
	SYNTHETICSDOWNTIMERESOURCETYPE_DOWNTIME SyntheticsDowntimeResourceType = "downtime"
)

var allowedSyntheticsDowntimeResourceTypeEnumValues = []SyntheticsDowntimeResourceType{
	SYNTHETICSDOWNTIMERESOURCETYPE_DOWNTIME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsDowntimeResourceType) GetAllowedValues() []SyntheticsDowntimeResourceType {
	return allowedSyntheticsDowntimeResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsDowntimeResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsDowntimeResourceType(value)
	return nil
}

// NewSyntheticsDowntimeResourceTypeFromValue returns a pointer to a valid SyntheticsDowntimeResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsDowntimeResourceTypeFromValue(v string) (*SyntheticsDowntimeResourceType, error) {
	ev := SyntheticsDowntimeResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsDowntimeResourceType: valid values are %v", v, allowedSyntheticsDowntimeResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsDowntimeResourceType) IsValid() bool {
	for _, existing := range allowedSyntheticsDowntimeResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsDowntimeResourceType value.
func (v SyntheticsDowntimeResourceType) Ptr() *SyntheticsDowntimeResourceType {
	return &v
}
