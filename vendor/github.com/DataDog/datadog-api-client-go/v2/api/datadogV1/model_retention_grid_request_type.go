// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionGridRequestType Request type for retention grid widget.
type RetentionGridRequestType string

// List of RetentionGridRequestType.
const (
	RETENTIONGRIDREQUESTTYPE_RETENTION_GRID RetentionGridRequestType = "retention_grid"
)

var allowedRetentionGridRequestTypeEnumValues = []RetentionGridRequestType{
	RETENTIONGRIDREQUESTTYPE_RETENTION_GRID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionGridRequestType) GetAllowedValues() []RetentionGridRequestType {
	return allowedRetentionGridRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionGridRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionGridRequestType(value)
	return nil
}

// NewRetentionGridRequestTypeFromValue returns a pointer to a valid RetentionGridRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionGridRequestTypeFromValue(v string) (*RetentionGridRequestType, error) {
	ev := RetentionGridRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionGridRequestType: valid values are %v", v, allowedRetentionGridRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionGridRequestType) IsValid() bool {
	for _, existing := range allowedRetentionGridRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionGridRequestType value.
func (v RetentionGridRequestType) Ptr() *RetentionGridRequestType {
	return &v
}
