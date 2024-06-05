// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionFilterType The type of retention filter. The value should always be spans-sampling-processor.
type RetentionFilterType string

// List of RetentionFilterType.
const (
	RETENTIONFILTERTYPE_SPANS_SAMPLING_PROCESSOR RetentionFilterType = "spans-sampling-processor"
)

var allowedRetentionFilterTypeEnumValues = []RetentionFilterType{
	RETENTIONFILTERTYPE_SPANS_SAMPLING_PROCESSOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionFilterType) GetAllowedValues() []RetentionFilterType {
	return allowedRetentionFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionFilterType(value)
	return nil
}

// NewRetentionFilterTypeFromValue returns a pointer to a valid RetentionFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionFilterTypeFromValue(v string) (*RetentionFilterType, error) {
	ev := RetentionFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionFilterType: valid values are %v", v, allowedRetentionFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionFilterType) IsValid() bool {
	for _, existing := range allowedRetentionFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionFilterType value.
func (v RetentionFilterType) Ptr() *RetentionFilterType {
	return &v
}
