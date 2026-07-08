// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumPermanentRetentionFilterType The type of the resource. The value should always be `permanent_retention_filters`.
type RumPermanentRetentionFilterType string

// List of RumPermanentRetentionFilterType.
const (
	RUMPERMANENTRETENTIONFILTERTYPE_PERMANENT_RETENTION_FILTERS RumPermanentRetentionFilterType = "permanent_retention_filters"
)

var allowedRumPermanentRetentionFilterTypeEnumValues = []RumPermanentRetentionFilterType{
	RUMPERMANENTRETENTIONFILTERTYPE_PERMANENT_RETENTION_FILTERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumPermanentRetentionFilterType) GetAllowedValues() []RumPermanentRetentionFilterType {
	return allowedRumPermanentRetentionFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumPermanentRetentionFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumPermanentRetentionFilterType(value)
	return nil
}

// NewRumPermanentRetentionFilterTypeFromValue returns a pointer to a valid RumPermanentRetentionFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumPermanentRetentionFilterTypeFromValue(v string) (*RumPermanentRetentionFilterType, error) {
	ev := RumPermanentRetentionFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumPermanentRetentionFilterType: valid values are %v", v, allowedRumPermanentRetentionFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumPermanentRetentionFilterType) IsValid() bool {
	for _, existing := range allowedRumPermanentRetentionFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumPermanentRetentionFilterType value.
func (v RumPermanentRetentionFilterType) Ptr() *RumPermanentRetentionFilterType {
	return &v
}
