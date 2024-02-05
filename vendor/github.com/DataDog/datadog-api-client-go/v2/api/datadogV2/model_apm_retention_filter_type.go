// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmRetentionFilterType The type of the resource.
type ApmRetentionFilterType string

// List of ApmRetentionFilterType.
const (
	APMRETENTIONFILTERTYPE_apm_retention_filter ApmRetentionFilterType = "apm_retention_filter"
)

var allowedApmRetentionFilterTypeEnumValues = []ApmRetentionFilterType{
	APMRETENTIONFILTERTYPE_apm_retention_filter,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmRetentionFilterType) GetAllowedValues() []ApmRetentionFilterType {
	return allowedApmRetentionFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmRetentionFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmRetentionFilterType(value)
	return nil
}

// NewApmRetentionFilterTypeFromValue returns a pointer to a valid ApmRetentionFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmRetentionFilterTypeFromValue(v string) (*ApmRetentionFilterType, error) {
	ev := ApmRetentionFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmRetentionFilterType: valid values are %v", v, allowedApmRetentionFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmRetentionFilterType) IsValid() bool {
	for _, existing := range allowedApmRetentionFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmRetentionFilterType value.
func (v ApmRetentionFilterType) Ptr() *ApmRetentionFilterType {
	return &v
}
