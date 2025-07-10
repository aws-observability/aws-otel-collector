// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay The weekday when the restriction period starts (Monday through Sunday).
type ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay string

// List of ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay.
const (
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_MONDAY    ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "monday"
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_TUESDAY   ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "tuesday"
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_WEDNESDAY ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "wednesday"
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_THURSDAY  ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "thursday"
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_FRIDAY    ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "friday"
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_SATURDAY  ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "saturday"
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_SUNDAY    ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay = "sunday"
)

var allowedScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDayEnumValues = []ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay{
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_MONDAY,
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_TUESDAY,
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_WEDNESDAY,
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_THURSDAY,
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_FRIDAY,
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_SATURDAY,
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_SUNDAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay) GetAllowedValues() []ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay {
	return allowedScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay(value)
	return nil
}

// NewScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDayFromValue returns a pointer to a valid ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDayFromValue(v string) (*ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay, error) {
	ev := ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay: valid values are %v", v, allowedScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay) IsValid() bool {
	for _, existing := range allowedScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay value.
func (v ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay) Ptr() *ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsStartDay {
	return &v
}
