// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleDataRelationshipsTeamsDataItemsType Teams resource type.
type ScheduleDataRelationshipsTeamsDataItemsType string

// List of ScheduleDataRelationshipsTeamsDataItemsType.
const (
	SCHEDULEDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS ScheduleDataRelationshipsTeamsDataItemsType = "teams"
)

var allowedScheduleDataRelationshipsTeamsDataItemsTypeEnumValues = []ScheduleDataRelationshipsTeamsDataItemsType{
	SCHEDULEDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleDataRelationshipsTeamsDataItemsType) GetAllowedValues() []ScheduleDataRelationshipsTeamsDataItemsType {
	return allowedScheduleDataRelationshipsTeamsDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleDataRelationshipsTeamsDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleDataRelationshipsTeamsDataItemsType(value)
	return nil
}

// NewScheduleDataRelationshipsTeamsDataItemsTypeFromValue returns a pointer to a valid ScheduleDataRelationshipsTeamsDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleDataRelationshipsTeamsDataItemsTypeFromValue(v string) (*ScheduleDataRelationshipsTeamsDataItemsType, error) {
	ev := ScheduleDataRelationshipsTeamsDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleDataRelationshipsTeamsDataItemsType: valid values are %v", v, allowedScheduleDataRelationshipsTeamsDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleDataRelationshipsTeamsDataItemsType) IsValid() bool {
	for _, existing := range allowedScheduleDataRelationshipsTeamsDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleDataRelationshipsTeamsDataItemsType value.
func (v ScheduleDataRelationshipsTeamsDataItemsType) Ptr() *ScheduleDataRelationshipsTeamsDataItemsType {
	return &v
}
