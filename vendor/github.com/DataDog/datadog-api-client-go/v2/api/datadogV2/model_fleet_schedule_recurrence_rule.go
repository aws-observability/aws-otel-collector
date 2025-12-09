// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetScheduleRecurrenceRule Defines the recurrence pattern for the schedule. Specifies when deployments should be
// automatically triggered based on maintenance windows.
type FleetScheduleRecurrenceRule struct {
	// List of days of the week when the schedule should trigger. Valid values are:
	// "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun".
	DaysOfWeek []string `json:"days_of_week"`
	// Duration of the maintenance window in minutes.
	MaintenanceWindowDuration int64 `json:"maintenance_window_duration"`
	// Start time of the maintenance window in 24-hour clock format (HH:MM).
	// Deployments will be triggered at this time on the specified days.
	StartMaintenanceWindow string `json:"start_maintenance_window"`
	// Timezone for the schedule in IANA Time Zone Database format (e.g., "America/New_York", "UTC").
	Timezone string `json:"timezone"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetScheduleRecurrenceRule instantiates a new FleetScheduleRecurrenceRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetScheduleRecurrenceRule(daysOfWeek []string, maintenanceWindowDuration int64, startMaintenanceWindow string, timezone string) *FleetScheduleRecurrenceRule {
	this := FleetScheduleRecurrenceRule{}
	this.DaysOfWeek = daysOfWeek
	this.MaintenanceWindowDuration = maintenanceWindowDuration
	this.StartMaintenanceWindow = startMaintenanceWindow
	this.Timezone = timezone
	return &this
}

// NewFleetScheduleRecurrenceRuleWithDefaults instantiates a new FleetScheduleRecurrenceRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetScheduleRecurrenceRuleWithDefaults() *FleetScheduleRecurrenceRule {
	this := FleetScheduleRecurrenceRule{}
	return &this
}

// GetDaysOfWeek returns the DaysOfWeek field value.
func (o *FleetScheduleRecurrenceRule) GetDaysOfWeek() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DaysOfWeek
}

// GetDaysOfWeekOk returns a tuple with the DaysOfWeek field value
// and a boolean to check if the value has been set.
func (o *FleetScheduleRecurrenceRule) GetDaysOfWeekOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysOfWeek, true
}

// SetDaysOfWeek sets field value.
func (o *FleetScheduleRecurrenceRule) SetDaysOfWeek(v []string) {
	o.DaysOfWeek = v
}

// GetMaintenanceWindowDuration returns the MaintenanceWindowDuration field value.
func (o *FleetScheduleRecurrenceRule) GetMaintenanceWindowDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaintenanceWindowDuration
}

// GetMaintenanceWindowDurationOk returns a tuple with the MaintenanceWindowDuration field value
// and a boolean to check if the value has been set.
func (o *FleetScheduleRecurrenceRule) GetMaintenanceWindowDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaintenanceWindowDuration, true
}

// SetMaintenanceWindowDuration sets field value.
func (o *FleetScheduleRecurrenceRule) SetMaintenanceWindowDuration(v int64) {
	o.MaintenanceWindowDuration = v
}

// GetStartMaintenanceWindow returns the StartMaintenanceWindow field value.
func (o *FleetScheduleRecurrenceRule) GetStartMaintenanceWindow() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StartMaintenanceWindow
}

// GetStartMaintenanceWindowOk returns a tuple with the StartMaintenanceWindow field value
// and a boolean to check if the value has been set.
func (o *FleetScheduleRecurrenceRule) GetStartMaintenanceWindowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartMaintenanceWindow, true
}

// SetStartMaintenanceWindow sets field value.
func (o *FleetScheduleRecurrenceRule) SetStartMaintenanceWindow(v string) {
	o.StartMaintenanceWindow = v
}

// GetTimezone returns the Timezone field value.
func (o *FleetScheduleRecurrenceRule) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *FleetScheduleRecurrenceRule) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value.
func (o *FleetScheduleRecurrenceRule) SetTimezone(v string) {
	o.Timezone = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetScheduleRecurrenceRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["days_of_week"] = o.DaysOfWeek
	toSerialize["maintenance_window_duration"] = o.MaintenanceWindowDuration
	toSerialize["start_maintenance_window"] = o.StartMaintenanceWindow
	toSerialize["timezone"] = o.Timezone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetScheduleRecurrenceRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DaysOfWeek                *[]string `json:"days_of_week"`
		MaintenanceWindowDuration *int64    `json:"maintenance_window_duration"`
		StartMaintenanceWindow    *string   `json:"start_maintenance_window"`
		Timezone                  *string   `json:"timezone"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DaysOfWeek == nil {
		return fmt.Errorf("required field days_of_week missing")
	}
	if all.MaintenanceWindowDuration == nil {
		return fmt.Errorf("required field maintenance_window_duration missing")
	}
	if all.StartMaintenanceWindow == nil {
		return fmt.Errorf("required field start_maintenance_window missing")
	}
	if all.Timezone == nil {
		return fmt.Errorf("required field timezone missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"days_of_week", "maintenance_window_duration", "start_maintenance_window", "timezone"})
	} else {
		return err
	}
	o.DaysOfWeek = *all.DaysOfWeek
	o.MaintenanceWindowDuration = *all.MaintenanceWindowDuration
	o.StartMaintenanceWindow = *all.StartMaintenanceWindow
	o.Timezone = *all.Timezone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
