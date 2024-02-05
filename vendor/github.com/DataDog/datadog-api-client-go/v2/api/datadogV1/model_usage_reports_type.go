// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageReportsType The type of reports.
type UsageReportsType string

// List of UsageReportsType.
const (
	USAGEREPORTSTYPE_REPORTS UsageReportsType = "reports"
)

var allowedUsageReportsTypeEnumValues = []UsageReportsType{
	USAGEREPORTSTYPE_REPORTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UsageReportsType) GetAllowedValues() []UsageReportsType {
	return allowedUsageReportsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UsageReportsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsageReportsType(value)
	return nil
}

// NewUsageReportsTypeFromValue returns a pointer to a valid UsageReportsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUsageReportsTypeFromValue(v string) (*UsageReportsType, error) {
	ev := UsageReportsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UsageReportsType: valid values are %v", v, allowedUsageReportsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UsageReportsType) IsValid() bool {
	for _, existing := range allowedUsageReportsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageReportsType value.
func (v UsageReportsType) Ptr() *UsageReportsType {
	return &v
}
