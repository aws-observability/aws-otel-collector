// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleCreateRequestDataRelationshipsTeamsDataItemsType Teams resource type.
type ScheduleCreateRequestDataRelationshipsTeamsDataItemsType string

// List of ScheduleCreateRequestDataRelationshipsTeamsDataItemsType.
const (
	SCHEDULECREATEREQUESTDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS ScheduleCreateRequestDataRelationshipsTeamsDataItemsType = "teams"
)

var allowedScheduleCreateRequestDataRelationshipsTeamsDataItemsTypeEnumValues = []ScheduleCreateRequestDataRelationshipsTeamsDataItemsType{
	SCHEDULECREATEREQUESTDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleCreateRequestDataRelationshipsTeamsDataItemsType) GetAllowedValues() []ScheduleCreateRequestDataRelationshipsTeamsDataItemsType {
	return allowedScheduleCreateRequestDataRelationshipsTeamsDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleCreateRequestDataRelationshipsTeamsDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleCreateRequestDataRelationshipsTeamsDataItemsType(value)
	return nil
}

// NewScheduleCreateRequestDataRelationshipsTeamsDataItemsTypeFromValue returns a pointer to a valid ScheduleCreateRequestDataRelationshipsTeamsDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleCreateRequestDataRelationshipsTeamsDataItemsTypeFromValue(v string) (*ScheduleCreateRequestDataRelationshipsTeamsDataItemsType, error) {
	ev := ScheduleCreateRequestDataRelationshipsTeamsDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleCreateRequestDataRelationshipsTeamsDataItemsType: valid values are %v", v, allowedScheduleCreateRequestDataRelationshipsTeamsDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleCreateRequestDataRelationshipsTeamsDataItemsType) IsValid() bool {
	for _, existing := range allowedScheduleCreateRequestDataRelationshipsTeamsDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleCreateRequestDataRelationshipsTeamsDataItemsType value.
func (v ScheduleCreateRequestDataRelationshipsTeamsDataItemsType) Ptr() *ScheduleCreateRequestDataRelationshipsTeamsDataItemsType {
	return &v
}
