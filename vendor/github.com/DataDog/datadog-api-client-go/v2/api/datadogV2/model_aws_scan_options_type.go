// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsScanOptionsType The type of the resource. The value should always be `aws_scan_options`.
type AwsScanOptionsType string

// List of AwsScanOptionsType.
const (
	AWSSCANOPTIONSTYPE_AWS_SCAN_OPTIONS AwsScanOptionsType = "aws_scan_options"
)

var allowedAwsScanOptionsTypeEnumValues = []AwsScanOptionsType{
	AWSSCANOPTIONSTYPE_AWS_SCAN_OPTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AwsScanOptionsType) GetAllowedValues() []AwsScanOptionsType {
	return allowedAwsScanOptionsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AwsScanOptionsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AwsScanOptionsType(value)
	return nil
}

// NewAwsScanOptionsTypeFromValue returns a pointer to a valid AwsScanOptionsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAwsScanOptionsTypeFromValue(v string) (*AwsScanOptionsType, error) {
	ev := AwsScanOptionsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AwsScanOptionsType: valid values are %v", v, allowedAwsScanOptionsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AwsScanOptionsType) IsValid() bool {
	for _, existing := range allowedAwsScanOptionsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsScanOptionsType value.
func (v AwsScanOptionsType) Ptr() *AwsScanOptionsType {
	return &v
}
