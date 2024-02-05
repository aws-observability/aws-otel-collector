// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceMapWidgetDefinitionType Type of the service map widget.
type ServiceMapWidgetDefinitionType string

// List of ServiceMapWidgetDefinitionType.
const (
	SERVICEMAPWIDGETDEFINITIONTYPE_SERVICEMAP ServiceMapWidgetDefinitionType = "servicemap"
)

var allowedServiceMapWidgetDefinitionTypeEnumValues = []ServiceMapWidgetDefinitionType{
	SERVICEMAPWIDGETDEFINITIONTYPE_SERVICEMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceMapWidgetDefinitionType) GetAllowedValues() []ServiceMapWidgetDefinitionType {
	return allowedServiceMapWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceMapWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceMapWidgetDefinitionType(value)
	return nil
}

// NewServiceMapWidgetDefinitionTypeFromValue returns a pointer to a valid ServiceMapWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceMapWidgetDefinitionTypeFromValue(v string) (*ServiceMapWidgetDefinitionType, error) {
	ev := ServiceMapWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceMapWidgetDefinitionType: valid values are %v", v, allowedServiceMapWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceMapWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedServiceMapWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceMapWidgetDefinitionType value.
func (v ServiceMapWidgetDefinitionType) Ptr() *ServiceMapWidgetDefinitionType {
	return &v
}
