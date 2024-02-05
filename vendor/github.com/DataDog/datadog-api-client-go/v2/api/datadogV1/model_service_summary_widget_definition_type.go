// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceSummaryWidgetDefinitionType Type of the service summary widget.
type ServiceSummaryWidgetDefinitionType string

// List of ServiceSummaryWidgetDefinitionType.
const (
	SERVICESUMMARYWIDGETDEFINITIONTYPE_TRACE_SERVICE ServiceSummaryWidgetDefinitionType = "trace_service"
)

var allowedServiceSummaryWidgetDefinitionTypeEnumValues = []ServiceSummaryWidgetDefinitionType{
	SERVICESUMMARYWIDGETDEFINITIONTYPE_TRACE_SERVICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceSummaryWidgetDefinitionType) GetAllowedValues() []ServiceSummaryWidgetDefinitionType {
	return allowedServiceSummaryWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceSummaryWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceSummaryWidgetDefinitionType(value)
	return nil
}

// NewServiceSummaryWidgetDefinitionTypeFromValue returns a pointer to a valid ServiceSummaryWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceSummaryWidgetDefinitionTypeFromValue(v string) (*ServiceSummaryWidgetDefinitionType, error) {
	ev := ServiceSummaryWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceSummaryWidgetDefinitionType: valid values are %v", v, allowedServiceSummaryWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceSummaryWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedServiceSummaryWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceSummaryWidgetDefinitionType value.
func (v ServiceSummaryWidgetDefinitionType) Ptr() *ServiceSummaryWidgetDefinitionType {
	return &v
}
