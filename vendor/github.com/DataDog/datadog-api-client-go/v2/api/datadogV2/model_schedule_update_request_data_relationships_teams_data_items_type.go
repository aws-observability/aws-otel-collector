// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType Teams resource type.
type ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType string

// List of ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType.
const (
	SCHEDULEUPDATEREQUESTDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType = "teams"
)

var allowedScheduleUpdateRequestDataRelationshipsTeamsDataItemsTypeEnumValues = []ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType{
	SCHEDULEUPDATEREQUESTDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType) GetAllowedValues() []ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType {
	return allowedScheduleUpdateRequestDataRelationshipsTeamsDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType(value)
	return nil
}

// NewScheduleUpdateRequestDataRelationshipsTeamsDataItemsTypeFromValue returns a pointer to a valid ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleUpdateRequestDataRelationshipsTeamsDataItemsTypeFromValue(v string) (*ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType, error) {
	ev := ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType: valid values are %v", v, allowedScheduleUpdateRequestDataRelationshipsTeamsDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType) IsValid() bool {
	for _, existing := range allowedScheduleUpdateRequestDataRelationshipsTeamsDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType value.
func (v ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType) Ptr() *ScheduleUpdateRequestDataRelationshipsTeamsDataItemsType {
	return &v
}
