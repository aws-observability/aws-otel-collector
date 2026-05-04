// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionCurveWidgetDefinitionType Type of the Retention Curve widget.
type RetentionCurveWidgetDefinitionType string

// List of RetentionCurveWidgetDefinitionType.
const (
	RETENTIONCURVEWIDGETDEFINITIONTYPE_RETENTION_CURVE RetentionCurveWidgetDefinitionType = "retention_curve"
)

var allowedRetentionCurveWidgetDefinitionTypeEnumValues = []RetentionCurveWidgetDefinitionType{
	RETENTIONCURVEWIDGETDEFINITIONTYPE_RETENTION_CURVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionCurveWidgetDefinitionType) GetAllowedValues() []RetentionCurveWidgetDefinitionType {
	return allowedRetentionCurveWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionCurveWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionCurveWidgetDefinitionType(value)
	return nil
}

// NewRetentionCurveWidgetDefinitionTypeFromValue returns a pointer to a valid RetentionCurveWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionCurveWidgetDefinitionTypeFromValue(v string) (*RetentionCurveWidgetDefinitionType, error) {
	ev := RetentionCurveWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionCurveWidgetDefinitionType: valid values are %v", v, allowedRetentionCurveWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionCurveWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedRetentionCurveWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionCurveWidgetDefinitionType value.
func (v RetentionCurveWidgetDefinitionType) Ptr() *RetentionCurveWidgetDefinitionType {
	return &v
}
