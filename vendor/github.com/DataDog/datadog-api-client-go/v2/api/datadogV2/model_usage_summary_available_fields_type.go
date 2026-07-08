// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSummaryAvailableFieldsType Type of available-fields data.
type UsageSummaryAvailableFieldsType string

// List of UsageSummaryAvailableFieldsType.
const (
	USAGESUMMARYAVAILABLEFIELDSTYPE_USAGE_SUMMARY_AVAILABLE_FIELDS UsageSummaryAvailableFieldsType = "usage_summary_available_fields"
)

var allowedUsageSummaryAvailableFieldsTypeEnumValues = []UsageSummaryAvailableFieldsType{
	USAGESUMMARYAVAILABLEFIELDSTYPE_USAGE_SUMMARY_AVAILABLE_FIELDS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UsageSummaryAvailableFieldsType) GetAllowedValues() []UsageSummaryAvailableFieldsType {
	return allowedUsageSummaryAvailableFieldsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UsageSummaryAvailableFieldsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsageSummaryAvailableFieldsType(value)
	return nil
}

// NewUsageSummaryAvailableFieldsTypeFromValue returns a pointer to a valid UsageSummaryAvailableFieldsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUsageSummaryAvailableFieldsTypeFromValue(v string) (*UsageSummaryAvailableFieldsType, error) {
	ev := UsageSummaryAvailableFieldsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UsageSummaryAvailableFieldsType: valid values are %v", v, allowedUsageSummaryAvailableFieldsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UsageSummaryAvailableFieldsType) IsValid() bool {
	for _, existing := range allowedUsageSummaryAvailableFieldsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageSummaryAvailableFieldsType value.
func (v UsageSummaryAvailableFieldsType) Ptr() *UsageSummaryAvailableFieldsType {
	return &v
}
