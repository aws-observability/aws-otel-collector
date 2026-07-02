// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetInfrastructureRequestRequestType Identifies this as an infrastructure-backed host map request.
type HostMapWidgetInfrastructureRequestRequestType string

// List of HostMapWidgetInfrastructureRequestRequestType.
const (
	HOSTMAPWIDGETINFRASTRUCTUREREQUESTREQUESTTYPE_INFRASTRUCTURE_HOSTMAP HostMapWidgetInfrastructureRequestRequestType = "infrastructure_hostmap"
)

var allowedHostMapWidgetInfrastructureRequestRequestTypeEnumValues = []HostMapWidgetInfrastructureRequestRequestType{
	HOSTMAPWIDGETINFRASTRUCTUREREQUESTREQUESTTYPE_INFRASTRUCTURE_HOSTMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostMapWidgetInfrastructureRequestRequestType) GetAllowedValues() []HostMapWidgetInfrastructureRequestRequestType {
	return allowedHostMapWidgetInfrastructureRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostMapWidgetInfrastructureRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostMapWidgetInfrastructureRequestRequestType(value)
	return nil
}

// NewHostMapWidgetInfrastructureRequestRequestTypeFromValue returns a pointer to a valid HostMapWidgetInfrastructureRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostMapWidgetInfrastructureRequestRequestTypeFromValue(v string) (*HostMapWidgetInfrastructureRequestRequestType, error) {
	ev := HostMapWidgetInfrastructureRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostMapWidgetInfrastructureRequestRequestType: valid values are %v", v, allowedHostMapWidgetInfrastructureRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostMapWidgetInfrastructureRequestRequestType) IsValid() bool {
	for _, existing := range allowedHostMapWidgetInfrastructureRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostMapWidgetInfrastructureRequestRequestType value.
func (v HostMapWidgetInfrastructureRequestRequestType) Ptr() *HostMapWidgetInfrastructureRequestRequestType {
	return &v
}
