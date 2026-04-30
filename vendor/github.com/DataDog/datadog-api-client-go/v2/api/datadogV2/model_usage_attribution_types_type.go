// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageAttributionTypesType Type of usage attribution types data.
type UsageAttributionTypesType string

// List of UsageAttributionTypesType.
const (
	USAGEATTRIBUTIONTYPESTYPE_USAGE_ATTRIBUTION_TYPES UsageAttributionTypesType = "usage_attribution_types"
)

var allowedUsageAttributionTypesTypeEnumValues = []UsageAttributionTypesType{
	USAGEATTRIBUTIONTYPESTYPE_USAGE_ATTRIBUTION_TYPES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UsageAttributionTypesType) GetAllowedValues() []UsageAttributionTypesType {
	return allowedUsageAttributionTypesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UsageAttributionTypesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsageAttributionTypesType(value)
	return nil
}

// NewUsageAttributionTypesTypeFromValue returns a pointer to a valid UsageAttributionTypesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUsageAttributionTypesTypeFromValue(v string) (*UsageAttributionTypesType, error) {
	ev := UsageAttributionTypesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UsageAttributionTypesType: valid values are %v", v, allowedUsageAttributionTypesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UsageAttributionTypesType) IsValid() bool {
	for _, existing := range allowedUsageAttributionTypesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageAttributionTypesType value.
func (v UsageAttributionTypesType) Ptr() *UsageAttributionTypesType {
	return &v
}
