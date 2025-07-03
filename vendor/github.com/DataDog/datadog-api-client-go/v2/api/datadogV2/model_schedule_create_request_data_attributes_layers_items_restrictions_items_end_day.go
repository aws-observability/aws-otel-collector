// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay The weekday when the restriction period ends (Monday through Sunday).
type ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay string

// List of ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay.
const (
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_MONDAY    ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "monday"
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_TUESDAY   ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "tuesday"
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_WEDNESDAY ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "wednesday"
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_THURSDAY  ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "thursday"
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_FRIDAY    ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "friday"
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_SATURDAY  ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "saturday"
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_SUNDAY    ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay = "sunday"
)

var allowedScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDayEnumValues = []ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay{
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_MONDAY,
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_TUESDAY,
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_WEDNESDAY,
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_THURSDAY,
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_FRIDAY,
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_SATURDAY,
	SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_SUNDAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay) GetAllowedValues() []ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay {
	return allowedScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay(value)
	return nil
}

// NewScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDayFromValue returns a pointer to a valid ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDayFromValue(v string) (*ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay, error) {
	ev := ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay: valid values are %v", v, allowedScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay) IsValid() bool {
	for _, existing := range allowedScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay value.
func (v ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay) Ptr() *ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItemsEndDay {
	return &v
}
