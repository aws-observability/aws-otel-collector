// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleDataRelationshipsLayersDataItemsType Layers resource type.
type ScheduleDataRelationshipsLayersDataItemsType string

// List of ScheduleDataRelationshipsLayersDataItemsType.
const (
	SCHEDULEDATARELATIONSHIPSLAYERSDATAITEMSTYPE_LAYERS ScheduleDataRelationshipsLayersDataItemsType = "layers"
)

var allowedScheduleDataRelationshipsLayersDataItemsTypeEnumValues = []ScheduleDataRelationshipsLayersDataItemsType{
	SCHEDULEDATARELATIONSHIPSLAYERSDATAITEMSTYPE_LAYERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleDataRelationshipsLayersDataItemsType) GetAllowedValues() []ScheduleDataRelationshipsLayersDataItemsType {
	return allowedScheduleDataRelationshipsLayersDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleDataRelationshipsLayersDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleDataRelationshipsLayersDataItemsType(value)
	return nil
}

// NewScheduleDataRelationshipsLayersDataItemsTypeFromValue returns a pointer to a valid ScheduleDataRelationshipsLayersDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleDataRelationshipsLayersDataItemsTypeFromValue(v string) (*ScheduleDataRelationshipsLayersDataItemsType, error) {
	ev := ScheduleDataRelationshipsLayersDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleDataRelationshipsLayersDataItemsType: valid values are %v", v, allowedScheduleDataRelationshipsLayersDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleDataRelationshipsLayersDataItemsType) IsValid() bool {
	for _, existing := range allowedScheduleDataRelationshipsLayersDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleDataRelationshipsLayersDataItemsType value.
func (v ScheduleDataRelationshipsLayersDataItemsType) Ptr() *ScheduleDataRelationshipsLayersDataItemsType {
	return &v
}
