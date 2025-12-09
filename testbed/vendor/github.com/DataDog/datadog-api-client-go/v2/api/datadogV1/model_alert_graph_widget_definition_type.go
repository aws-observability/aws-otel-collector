// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertGraphWidgetDefinitionType Type of the alert graph widget.
type AlertGraphWidgetDefinitionType string

// List of AlertGraphWidgetDefinitionType.
const (
	ALERTGRAPHWIDGETDEFINITIONTYPE_ALERT_GRAPH AlertGraphWidgetDefinitionType = "alert_graph"
)

var allowedAlertGraphWidgetDefinitionTypeEnumValues = []AlertGraphWidgetDefinitionType{
	ALERTGRAPHWIDGETDEFINITIONTYPE_ALERT_GRAPH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AlertGraphWidgetDefinitionType) GetAllowedValues() []AlertGraphWidgetDefinitionType {
	return allowedAlertGraphWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertGraphWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertGraphWidgetDefinitionType(value)
	return nil
}

// NewAlertGraphWidgetDefinitionTypeFromValue returns a pointer to a valid AlertGraphWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertGraphWidgetDefinitionTypeFromValue(v string) (*AlertGraphWidgetDefinitionType, error) {
	ev := AlertGraphWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertGraphWidgetDefinitionType: valid values are %v", v, allowedAlertGraphWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertGraphWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedAlertGraphWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertGraphWidgetDefinitionType value.
func (v AlertGraphWidgetDefinitionType) Ptr() *AlertGraphWidgetDefinitionType {
	return &v
}
