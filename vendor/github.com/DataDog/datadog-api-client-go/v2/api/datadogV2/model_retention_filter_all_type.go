// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionFilterAllType The type of retention filter.
type RetentionFilterAllType string

// List of RetentionFilterAllType.
const (
	RETENTIONFILTERALLTYPE_SPANS_SAMPLING_PROCESSOR        RetentionFilterAllType = "spans-sampling-processor"
	RETENTIONFILTERALLTYPE_SPANS_ERRORS_SAMPLING_PROCESSOR RetentionFilterAllType = "spans-errors-sampling-processor"
	RETENTIONFILTERALLTYPE_SPANS_APPSEC_SAMPLING_PROCESSOR RetentionFilterAllType = "spans-appsec-sampling-processor"
)

var allowedRetentionFilterAllTypeEnumValues = []RetentionFilterAllType{
	RETENTIONFILTERALLTYPE_SPANS_SAMPLING_PROCESSOR,
	RETENTIONFILTERALLTYPE_SPANS_ERRORS_SAMPLING_PROCESSOR,
	RETENTIONFILTERALLTYPE_SPANS_APPSEC_SAMPLING_PROCESSOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionFilterAllType) GetAllowedValues() []RetentionFilterAllType {
	return allowedRetentionFilterAllTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionFilterAllType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionFilterAllType(value)
	return nil
}

// NewRetentionFilterAllTypeFromValue returns a pointer to a valid RetentionFilterAllType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionFilterAllTypeFromValue(v string) (*RetentionFilterAllType, error) {
	ev := RetentionFilterAllType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionFilterAllType: valid values are %v", v, allowedRetentionFilterAllTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionFilterAllType) IsValid() bool {
	for _, existing := range allowedRetentionFilterAllTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionFilterAllType value.
func (v RetentionFilterAllType) Ptr() *RetentionFilterAllType {
	return &v
}
