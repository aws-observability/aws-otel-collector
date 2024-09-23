// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeResponseAttributes Downtime details.
type DowntimeResponseAttributes struct {
	// Time that the downtime was canceled.
	Canceled datadog.NullableTime `json:"canceled,omitempty"`
	// Creation time of the downtime.
	Created *time.Time `json:"created,omitempty"`
	// The timezone in which to display the downtime's start and end times in Datadog applications. This is not used
	// as an offset for scheduling.
	DisplayTimezone datadog.NullableString `json:"display_timezone,omitempty"`
	// A message to include with notifications for this downtime. Email notifications can be sent to specific users
	// by using the same `@username` notation as events.
	Message datadog.NullableString `json:"message,omitempty"`
	// Time that the downtime was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// Monitor identifier for the downtime.
	MonitorIdentifier *DowntimeMonitorIdentifier `json:"monitor_identifier,omitempty"`
	// If the first recovery notification during a downtime should be muted.
	MuteFirstRecoveryNotification *bool `json:"mute_first_recovery_notification,omitempty"`
	// States that will trigger a monitor notification when the `notify_end_types` action occurs.
	NotifyEndStates []DowntimeNotifyEndStateTypes `json:"notify_end_states,omitempty"`
	// Actions that will trigger a monitor notification if the downtime is in the `notify_end_types` state.
	NotifyEndTypes []DowntimeNotifyEndStateActions `json:"notify_end_types,omitempty"`
	// The schedule that defines when the monitor starts, stops, and recurs. There are two types of schedules:
	// one-time and recurring. Recurring schedules may have up to five RRULE-based recurrences. If no schedules are
	// provided, the downtime will begin immediately and never end.
	Schedule *DowntimeScheduleResponse `json:"schedule,omitempty"`
	// The scope to which the downtime applies. Must follow the [common search syntax](https://docs.datadoghq.com/logs/explorer/search_syntax/).
	Scope *string `json:"scope,omitempty"`
	// The current status of the downtime.
	Status *DowntimeStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDowntimeResponseAttributes instantiates a new DowntimeResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeResponseAttributes() *DowntimeResponseAttributes {
	this := DowntimeResponseAttributes{}
	var displayTimezone string = "UTC"
	this.DisplayTimezone = *datadog.NewNullableString(&displayTimezone)
	return &this
}

// NewDowntimeResponseAttributesWithDefaults instantiates a new DowntimeResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeResponseAttributesWithDefaults() *DowntimeResponseAttributes {
	this := DowntimeResponseAttributes{}
	var displayTimezone string = "UTC"
	this.DisplayTimezone = *datadog.NewNullableString(&displayTimezone)
	return &this
}

// GetCanceled returns the Canceled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeResponseAttributes) GetCanceled() time.Time {
	if o == nil || o.Canceled.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Canceled.Get()
}

// GetCanceledOk returns a tuple with the Canceled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeResponseAttributes) GetCanceledOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Canceled.Get(), o.Canceled.IsSet()
}

// HasCanceled returns a boolean if a field has been set.
func (o *DowntimeResponseAttributes) HasCanceled() bool {
	return o != nil && o.Canceled.IsSet()
}

// SetCanceled gets a reference to the given datadog.NullableTime and assigns it to the Canceled field.
func (o *DowntimeResponseAttributes) SetCanceled(v time.Time) {
	o.Canceled.Set(&v)
}

// SetCanceledNil sets the value for Canceled to be an explicit nil.
func (o *DowntimeResponseAttributes) SetCanceledNil() {
	o.Canceled.Set(nil)
}

// UnsetCanceled ensures that no value is present for Canceled, not even an explicit nil.
func (o *DowntimeResponseAttributes) UnsetCanceled() {
	o.Canceled.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DowntimeResponseAttributes) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeResponseAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DowntimeResponseAttributes) HasCreated() bool {
	return o != nil && o.Created != nil
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *DowntimeResponseAttributes) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDisplayTimezone returns the DisplayTimezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeResponseAttributes) GetDisplayTimezone() string {
	if o == nil || o.DisplayTimezone.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayTimezone.Get()
}

// GetDisplayTimezoneOk returns a tuple with the DisplayTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeResponseAttributes) GetDisplayTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayTimezone.Get(), o.DisplayTimezone.IsSet()
}

// HasDisplayTimezone returns a boolean if a field has been set.
func (o *DowntimeResponseAttributes) HasDisplayTimezone() bool {
	return o != nil && o.DisplayTimezone.IsSet()
}

// SetDisplayTimezone gets a reference to the given datadog.NullableString and assigns it to the DisplayTimezone field.
func (o *DowntimeResponseAttributes) SetDisplayTimezone(v string) {
	o.DisplayTimezone.Set(&v)
}

// SetDisplayTimezoneNil sets the value for DisplayTimezone to be an explicit nil.
func (o *DowntimeResponseAttributes) SetDisplayTimezoneNil() {
	o.DisplayTimezone.Set(nil)
}

// UnsetDisplayTimezone ensures that no value is present for DisplayTimezone, not even an explicit nil.
func (o *DowntimeResponseAttributes) UnsetDisplayTimezone() {
	o.DisplayTimezone.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeResponseAttributes) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeResponseAttributes) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *DowntimeResponseAttributes) HasMessage() bool {
	return o != nil && o.Message.IsSet()
}

// SetMessage gets a reference to the given datadog.NullableString and assigns it to the Message field.
func (o *DowntimeResponseAttributes) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil.
func (o *DowntimeResponseAttributes) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil.
func (o *DowntimeResponseAttributes) UnsetMessage() {
	o.Message.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *DowntimeResponseAttributes) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeResponseAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *DowntimeResponseAttributes) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *DowntimeResponseAttributes) SetModified(v time.Time) {
	o.Modified = &v
}

// GetMonitorIdentifier returns the MonitorIdentifier field value if set, zero value otherwise.
func (o *DowntimeResponseAttributes) GetMonitorIdentifier() DowntimeMonitorIdentifier {
	if o == nil || o.MonitorIdentifier == nil {
		var ret DowntimeMonitorIdentifier
		return ret
	}
	return *o.MonitorIdentifier
}

// GetMonitorIdentifierOk returns a tuple with the MonitorIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeResponseAttributes) GetMonitorIdentifierOk() (*DowntimeMonitorIdentifier, bool) {
	if o == nil || o.MonitorIdentifier == nil {
		return nil, false
	}
	return o.MonitorIdentifier, true
}

// HasMonitorIdentifier returns a boolean if a field has been set.
func (o *DowntimeResponseAttributes) HasMonitorIdentifier() bool {
	return o != nil && o.MonitorIdentifier != nil
}

// SetMonitorIdentifier gets a reference to the given DowntimeMonitorIdentifier and assigns it to the MonitorIdentifier field.
func (o *DowntimeResponseAttributes) SetMonitorIdentifier(v DowntimeMonitorIdentifier) {
	o.MonitorIdentifier = &v
}

// GetMuteFirstRecoveryNotification returns the MuteFirstRecoveryNotification field value if set, zero value otherwise.
func (o *DowntimeResponseAttributes) GetMuteFirstRecoveryNotification() bool {
	if o == nil || o.MuteFirstRecoveryNotification == nil {
		var ret bool
		return ret
	}
	return *o.MuteFirstRecoveryNotification
}

// GetMuteFirstRecoveryNotificationOk returns a tuple with the MuteFirstRecoveryNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeResponseAttributes) GetMuteFirstRecoveryNotificationOk() (*bool, bool) {
	if o == nil || o.MuteFirstRecoveryNotification == nil {
		return nil, false
	}
	return o.MuteFirstRecoveryNotification, true
}

// HasMuteFirstRecoveryNotification returns a boolean if a field has been set.
func (o *DowntimeResponseAttributes) HasMuteFirstRecoveryNotification() bool {
	return o != nil && o.MuteFirstRecoveryNotification != nil
}

// SetMuteFirstRecoveryNotification gets a reference to the given bool and assigns it to the MuteFirstRecoveryNotification field.
func (o *DowntimeResponseAttributes) SetMuteFirstRecoveryNotification(v bool) {
	o.MuteFirstRecoveryNotification = &v
}

// GetNotifyEndStates returns the NotifyEndStates field value if set, zero value otherwise.
func (o *DowntimeResponseAttributes) GetNotifyEndStates() []DowntimeNotifyEndStateTypes {
	if o == nil || o.NotifyEndStates == nil {
		var ret []DowntimeNotifyEndStateTypes
		return ret
	}
	return o.NotifyEndStates
}

// GetNotifyEndStatesOk returns a tuple with the NotifyEndStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeResponseAttributes) GetNotifyEndStatesOk() (*[]DowntimeNotifyEndStateTypes, bool) {
	if o == nil || o.NotifyEndStates == nil {
		return nil, false
	}
	return &o.NotifyEndStates, true
}

// HasNotifyEndStates returns a boolean if a field has been set.
func (o *DowntimeResponseAttributes) HasNotifyEndStates() bool {
	return o != nil && o.NotifyEndStates != nil
}

// SetNotifyEndStates gets a reference to the given []DowntimeNotifyEndStateTypes and assigns it to the NotifyEndStates field.
func (o *DowntimeResponseAttributes) SetNotifyEndStates(v []DowntimeNotifyEndStateTypes) {
	o.NotifyEndStates = v
}

// GetNotifyEndTypes returns the NotifyEndTypes field value if set, zero value otherwise.
func (o *DowntimeResponseAttributes) GetNotifyEndTypes() []DowntimeNotifyEndStateActions {
	if o == nil || o.NotifyEndTypes == nil {
		var ret []DowntimeNotifyEndStateActions
		return ret
	}
	return o.NotifyEndTypes
}

// GetNotifyEndTypesOk returns a tuple with the NotifyEndTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeResponseAttributes) GetNotifyEndTypesOk() (*[]DowntimeNotifyEndStateActions, bool) {
	if o == nil || o.NotifyEndTypes == nil {
		return nil, false
	}
	return &o.NotifyEndTypes, true
}

// HasNotifyEndTypes returns a boolean if a field has been set.
func (o *DowntimeResponseAttributes) HasNotifyEndTypes() bool {
	return o != nil && o.NotifyEndTypes != nil
}

// SetNotifyEndTypes gets a reference to the given []DowntimeNotifyEndStateActions and assigns it to the NotifyEndTypes field.
func (o *DowntimeResponseAttributes) SetNotifyEndTypes(v []DowntimeNotifyEndStateActions) {
	o.NotifyEndTypes = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *DowntimeResponseAttributes) GetSchedule() DowntimeScheduleResponse {
	if o == nil || o.Schedule == nil {
		var ret DowntimeScheduleResponse
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeResponseAttributes) GetScheduleOk() (*DowntimeScheduleResponse, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *DowntimeResponseAttributes) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given DowntimeScheduleResponse and assigns it to the Schedule field.
func (o *DowntimeResponseAttributes) SetSchedule(v DowntimeScheduleResponse) {
	o.Schedule = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *DowntimeResponseAttributes) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeResponseAttributes) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *DowntimeResponseAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *DowntimeResponseAttributes) SetScope(v string) {
	o.Scope = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DowntimeResponseAttributes) GetStatus() DowntimeStatus {
	if o == nil || o.Status == nil {
		var ret DowntimeStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeResponseAttributes) GetStatusOk() (*DowntimeStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DowntimeResponseAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given DowntimeStatus and assigns it to the Status field.
func (o *DowntimeResponseAttributes) SetStatus(v DowntimeStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Canceled.IsSet() {
		toSerialize["canceled"] = o.Canceled.Get()
	}
	if o.Created != nil {
		if o.Created.Nanosecond() == 0 {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DisplayTimezone.IsSet() {
		toSerialize["display_timezone"] = o.DisplayTimezone.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Modified != nil {
		if o.Modified.Nanosecond() == 0 {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.MonitorIdentifier != nil {
		toSerialize["monitor_identifier"] = o.MonitorIdentifier
	}
	if o.MuteFirstRecoveryNotification != nil {
		toSerialize["mute_first_recovery_notification"] = o.MuteFirstRecoveryNotification
	}
	if o.NotifyEndStates != nil {
		toSerialize["notify_end_states"] = o.NotifyEndStates
	}
	if o.NotifyEndTypes != nil {
		toSerialize["notify_end_types"] = o.NotifyEndTypes
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Canceled                      datadog.NullableTime            `json:"canceled,omitempty"`
		Created                       *time.Time                      `json:"created,omitempty"`
		DisplayTimezone               datadog.NullableString          `json:"display_timezone,omitempty"`
		Message                       datadog.NullableString          `json:"message,omitempty"`
		Modified                      *time.Time                      `json:"modified,omitempty"`
		MonitorIdentifier             *DowntimeMonitorIdentifier      `json:"monitor_identifier,omitempty"`
		MuteFirstRecoveryNotification *bool                           `json:"mute_first_recovery_notification,omitempty"`
		NotifyEndStates               []DowntimeNotifyEndStateTypes   `json:"notify_end_states,omitempty"`
		NotifyEndTypes                []DowntimeNotifyEndStateActions `json:"notify_end_types,omitempty"`
		Schedule                      *DowntimeScheduleResponse       `json:"schedule,omitempty"`
		Scope                         *string                         `json:"scope,omitempty"`
		Status                        *DowntimeStatus                 `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"canceled", "created", "display_timezone", "message", "modified", "monitor_identifier", "mute_first_recovery_notification", "notify_end_states", "notify_end_types", "schedule", "scope", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Canceled = all.Canceled
	o.Created = all.Created
	o.DisplayTimezone = all.DisplayTimezone
	o.Message = all.Message
	o.Modified = all.Modified
	o.MonitorIdentifier = all.MonitorIdentifier
	o.MuteFirstRecoveryNotification = all.MuteFirstRecoveryNotification
	o.NotifyEndStates = all.NotifyEndStates
	o.NotifyEndTypes = all.NotifyEndTypes
	o.Schedule = all.Schedule
	o.Scope = all.Scope
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
