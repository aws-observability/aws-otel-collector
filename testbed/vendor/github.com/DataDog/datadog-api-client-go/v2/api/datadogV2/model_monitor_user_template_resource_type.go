// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorUserTemplateResourceType Monitor user template resource type.
type MonitorUserTemplateResourceType string

// List of MonitorUserTemplateResourceType.
const (
	MONITORUSERTEMPLATERESOURCETYPE_MONITOR_USER_TEMPLATE MonitorUserTemplateResourceType = "monitor-user-template"
)

var allowedMonitorUserTemplateResourceTypeEnumValues = []MonitorUserTemplateResourceType{
	MONITORUSERTEMPLATERESOURCETYPE_MONITOR_USER_TEMPLATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorUserTemplateResourceType) GetAllowedValues() []MonitorUserTemplateResourceType {
	return allowedMonitorUserTemplateResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorUserTemplateResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorUserTemplateResourceType(value)
	return nil
}

// NewMonitorUserTemplateResourceTypeFromValue returns a pointer to a valid MonitorUserTemplateResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorUserTemplateResourceTypeFromValue(v string) (*MonitorUserTemplateResourceType, error) {
	ev := MonitorUserTemplateResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorUserTemplateResourceType: valid values are %v", v, allowedMonitorUserTemplateResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorUserTemplateResourceType) IsValid() bool {
	for _, existing := range allowedMonitorUserTemplateResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorUserTemplateResourceType value.
func (v MonitorUserTemplateResourceType) Ptr() *MonitorUserTemplateResourceType {
	return &v
}
