// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionFilterType The type of the resource. The value should always be retention_filters.
type RumRetentionFilterType string

// List of RumRetentionFilterType.
const (
	RUMRETENTIONFILTERTYPE_RETENTION_FILTERS RumRetentionFilterType = "retention_filters"
)

var allowedRumRetentionFilterTypeEnumValues = []RumRetentionFilterType{
	RUMRETENTIONFILTERTYPE_RETENTION_FILTERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRetentionFilterType) GetAllowedValues() []RumRetentionFilterType {
	return allowedRumRetentionFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRetentionFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRetentionFilterType(value)
	return nil
}

// NewRumRetentionFilterTypeFromValue returns a pointer to a valid RumRetentionFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRetentionFilterTypeFromValue(v string) (*RumRetentionFilterType, error) {
	ev := RumRetentionFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRetentionFilterType: valid values are %v", v, allowedRumRetentionFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRetentionFilterType) IsValid() bool {
	for _, existing := range allowedRumRetentionFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRetentionFilterType value.
func (v RumRetentionFilterType) Ptr() *RumRetentionFilterType {
	return &v
}
