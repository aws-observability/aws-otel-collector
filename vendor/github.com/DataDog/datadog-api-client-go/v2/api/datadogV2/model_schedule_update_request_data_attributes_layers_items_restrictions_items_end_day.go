// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay Defines the day of the week on which the time restriction ends.
type ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay string

// List of ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay.
const (
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_MONDAY    ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "monday"
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_TUESDAY   ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "tuesday"
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_WEDNESDAY ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "wednesday"
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_THURSDAY  ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "thursday"
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_FRIDAY    ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "friday"
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_SATURDAY  ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "saturday"
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_SUNDAY    ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "sunday"
)

var allowedScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDayEnumValues = []ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay{
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_MONDAY,
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_TUESDAY,
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_WEDNESDAY,
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_THURSDAY,
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_FRIDAY,
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_SATURDAY,
	SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_SUNDAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay) GetAllowedValues() []ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay {
	return allowedScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay(value)
	return nil
}

// NewScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDayFromValue returns a pointer to a valid ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDayFromValue(v string) (*ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay, error) {
	ev := ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay: valid values are %v", v, allowedScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay) IsValid() bool {
	for _, existing := range allowedScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay value.
func (v ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay) Ptr() *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay {
	return &v
}
