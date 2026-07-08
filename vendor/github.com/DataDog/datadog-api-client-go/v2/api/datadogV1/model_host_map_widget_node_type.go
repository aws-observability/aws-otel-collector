// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetNodeType Which type of infrastructure entity to visualize in the host map.
type HostMapWidgetNodeType string

// List of HostMapWidgetNodeType.
const (
	HOSTMAPWIDGETNODETYPE_HOST      HostMapWidgetNodeType = "host"
	HOSTMAPWIDGETNODETYPE_CONTAINER HostMapWidgetNodeType = "container"
	HOSTMAPWIDGETNODETYPE_POD       HostMapWidgetNodeType = "pod"
	HOSTMAPWIDGETNODETYPE_CLUSTER   HostMapWidgetNodeType = "cluster"
)

var allowedHostMapWidgetNodeTypeEnumValues = []HostMapWidgetNodeType{
	HOSTMAPWIDGETNODETYPE_HOST,
	HOSTMAPWIDGETNODETYPE_CONTAINER,
	HOSTMAPWIDGETNODETYPE_POD,
	HOSTMAPWIDGETNODETYPE_CLUSTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostMapWidgetNodeType) GetAllowedValues() []HostMapWidgetNodeType {
	return allowedHostMapWidgetNodeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostMapWidgetNodeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostMapWidgetNodeType(value)
	return nil
}

// NewHostMapWidgetNodeTypeFromValue returns a pointer to a valid HostMapWidgetNodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostMapWidgetNodeTypeFromValue(v string) (*HostMapWidgetNodeType, error) {
	ev := HostMapWidgetNodeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostMapWidgetNodeType: valid values are %v", v, allowedHostMapWidgetNodeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostMapWidgetNodeType) IsValid() bool {
	for _, existing := range allowedHostMapWidgetNodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostMapWidgetNodeType value.
func (v HostMapWidgetNodeType) Ptr() *HostMapWidgetNodeType {
	return &v
}
