// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapWidgetDefinitionType Type of the topology map widget.
type TopologyMapWidgetDefinitionType string

// List of TopologyMapWidgetDefinitionType.
const (
	TOPOLOGYMAPWIDGETDEFINITIONTYPE_TOPOLOGY_MAP TopologyMapWidgetDefinitionType = "topology_map"
)

var allowedTopologyMapWidgetDefinitionTypeEnumValues = []TopologyMapWidgetDefinitionType{
	TOPOLOGYMAPWIDGETDEFINITIONTYPE_TOPOLOGY_MAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyMapWidgetDefinitionType) GetAllowedValues() []TopologyMapWidgetDefinitionType {
	return allowedTopologyMapWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyMapWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyMapWidgetDefinitionType(value)
	return nil
}

// NewTopologyMapWidgetDefinitionTypeFromValue returns a pointer to a valid TopologyMapWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyMapWidgetDefinitionTypeFromValue(v string) (*TopologyMapWidgetDefinitionType, error) {
	ev := TopologyMapWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyMapWidgetDefinitionType: valid values are %v", v, allowedTopologyMapWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyMapWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedTopologyMapWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyMapWidgetDefinitionType value.
func (v TopologyMapWidgetDefinitionType) Ptr() *TopologyMapWidgetDefinitionType {
	return &v
}
