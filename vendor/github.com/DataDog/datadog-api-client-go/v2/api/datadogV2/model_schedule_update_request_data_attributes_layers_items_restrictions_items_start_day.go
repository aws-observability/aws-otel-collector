// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay Defines the day of the week on which the time restriction starts.
type ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay string

// List of ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay.
const (
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_MONDAY    ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "monday"
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_TUESDAY   ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "tuesday"
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_WEDNESDAY ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "wednesday"
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_THURSDAY  ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "thursday"
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_FRIDAY    ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "friday"
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_SATURDAY  ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "saturday"
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_SUNDAY    ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "sunday"
)

var allowedScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDayEnumValues = []ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay{
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_MONDAY,
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_TUESDAY,
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_WEDNESDAY,
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_THURSDAY,
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_FRIDAY,
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_SATURDAY,
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_SUNDAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay) GetAllowedValues() []ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay {
	return allowedScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay(value)
	return nil
}

// NewScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDayFromValue returns a pointer to a valid ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDayFromValue(v string) (*ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay, error) {
	ev := ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay: valid values are %v", v, allowedScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay) IsValid() bool {
	for _, existing := range allowedScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay value.
func (v ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay) Ptr() *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay {
	return &v
}
