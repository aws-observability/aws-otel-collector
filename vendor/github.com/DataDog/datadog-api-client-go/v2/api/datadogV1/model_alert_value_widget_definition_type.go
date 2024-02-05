// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertValueWidgetDefinitionType Type of the alert value widget.
type AlertValueWidgetDefinitionType string

// List of AlertValueWidgetDefinitionType.
const (
	ALERTVALUEWIDGETDEFINITIONTYPE_ALERT_VALUE AlertValueWidgetDefinitionType = "alert_value"
)

var allowedAlertValueWidgetDefinitionTypeEnumValues = []AlertValueWidgetDefinitionType{
	ALERTVALUEWIDGETDEFINITIONTYPE_ALERT_VALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AlertValueWidgetDefinitionType) GetAllowedValues() []AlertValueWidgetDefinitionType {
	return allowedAlertValueWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertValueWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertValueWidgetDefinitionType(value)
	return nil
}

// NewAlertValueWidgetDefinitionTypeFromValue returns a pointer to a valid AlertValueWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertValueWidgetDefinitionTypeFromValue(v string) (*AlertValueWidgetDefinitionType, error) {
	ev := AlertValueWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertValueWidgetDefinitionType: valid values are %v", v, allowedAlertValueWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertValueWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedAlertValueWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertValueWidgetDefinitionType value.
func (v AlertValueWidgetDefinitionType) Ptr() *AlertValueWidgetDefinitionType {
	return &v
}
