// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionCurveRequestType Request type for retention curve widget.
type RetentionCurveRequestType string

// List of RetentionCurveRequestType.
const (
	RETENTIONCURVEREQUESTTYPE_RETENTION_CURVE RetentionCurveRequestType = "retention_curve"
)

var allowedRetentionCurveRequestTypeEnumValues = []RetentionCurveRequestType{
	RETENTIONCURVEREQUESTTYPE_RETENTION_CURVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionCurveRequestType) GetAllowedValues() []RetentionCurveRequestType {
	return allowedRetentionCurveRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionCurveRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionCurveRequestType(value)
	return nil
}

// NewRetentionCurveRequestTypeFromValue returns a pointer to a valid RetentionCurveRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionCurveRequestTypeFromValue(v string) (*RetentionCurveRequestType, error) {
	ev := RetentionCurveRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionCurveRequestType: valid values are %v", v, allowedRetentionCurveRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionCurveRequestType) IsValid() bool {
	for _, existing := range allowedRetentionCurveRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionCurveRequestType value.
func (v RetentionCurveRequestType) Ptr() *RetentionCurveRequestType {
	return &v
}
