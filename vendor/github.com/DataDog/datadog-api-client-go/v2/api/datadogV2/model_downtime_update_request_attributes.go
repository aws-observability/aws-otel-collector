// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeUpdateRequestAttributes Attributes of the downtime to update.
type DowntimeUpdateRequestAttributes struct {
	// The timezone in which to display the downtime's start and end times in Datadog applications. This is not used
	// as an offset for scheduling.
	DisplayTimezone datadog.NullableString `json:"display_timezone,omitempty"`
	// A message to include with notifications for this downtime. Email notifications can be sent to specific users
	// by using the same `@username` notation as events.
	Message datadog.NullableString `json:"message,omitempty"`
	// Monitor identifier for the downtime.
	MonitorIdentifier *DowntimeMonitorIdentifier `json:"monitor_identifier,omitempty"`
	// If the first recovery notification during a downtime should be muted.
	MuteFirstRecoveryNotification *bool `json:"mute_first_recovery_notification,omitempty"`
	// States that will trigger a monitor notification when the `notify_end_types` action occurs.
	NotifyEndStates []DowntimeNotifyEndStateTypes `json:"notify_end_states,omitempty"`
	// Actions that will trigger a monitor notification if the downtime is in the `notify_end_types` state.
	NotifyEndTypes []DowntimeNotifyEndStateActions `json:"notify_end_types,omitempty"`
	// Schedule for the downtime.
	Schedule *DowntimeScheduleUpdateRequest `json:"schedule,omitempty"`
	// The scope to which the downtime applies. Must follow the [common search syntax](https://docs.datadoghq.com/logs/explorer/search_syntax/).
	Scope *string `json:"scope,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDowntimeUpdateRequestAttributes instantiates a new DowntimeUpdateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeUpdateRequestAttributes() *DowntimeUpdateRequestAttributes {
	this := DowntimeUpdateRequestAttributes{}
	var displayTimezone string = "UTC"
	this.DisplayTimezone = *datadog.NewNullableString(&displayTimezone)
	return &this
}

// NewDowntimeUpdateRequestAttributesWithDefaults instantiates a new DowntimeUpdateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeUpdateRequestAttributesWithDefaults() *DowntimeUpdateRequestAttributes {
	this := DowntimeUpdateRequestAttributes{}
	var displayTimezone string = "UTC"
	this.DisplayTimezone = *datadog.NewNullableString(&displayTimezone)
	return &this
}

// GetDisplayTimezone returns the DisplayTimezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeUpdateRequestAttributes) GetDisplayTimezone() string {
	if o == nil || o.DisplayTimezone.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayTimezone.Get()
}

// GetDisplayTimezoneOk returns a tuple with the DisplayTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeUpdateRequestAttributes) GetDisplayTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayTimezone.Get(), o.DisplayTimezone.IsSet()
}

// HasDisplayTimezone returns a boolean if a field has been set.
func (o *DowntimeUpdateRequestAttributes) HasDisplayTimezone() bool {
	return o != nil && o.DisplayTimezone.IsSet()
}

// SetDisplayTimezone gets a reference to the given datadog.NullableString and assigns it to the DisplayTimezone field.
func (o *DowntimeUpdateRequestAttributes) SetDisplayTimezone(v string) {
	o.DisplayTimezone.Set(&v)
}

// SetDisplayTimezoneNil sets the value for DisplayTimezone to be an explicit nil.
func (o *DowntimeUpdateRequestAttributes) SetDisplayTimezoneNil() {
	o.DisplayTimezone.Set(nil)
}

// UnsetDisplayTimezone ensures that no value is present for DisplayTimezone, not even an explicit nil.
func (o *DowntimeUpdateRequestAttributes) UnsetDisplayTimezone() {
	o.DisplayTimezone.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeUpdateRequestAttributes) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeUpdateRequestAttributes) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *DowntimeUpdateRequestAttributes) HasMessage() bool {
	return o != nil && o.Message.IsSet()
}

// SetMessage gets a reference to the given datadog.NullableString and assigns it to the Message field.
func (o *DowntimeUpdateRequestAttributes) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil.
func (o *DowntimeUpdateRequestAttributes) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil.
func (o *DowntimeUpdateRequestAttributes) UnsetMessage() {
	o.Message.Unset()
}

// GetMonitorIdentifier returns the MonitorIdentifier field value if set, zero value otherwise.
func (o *DowntimeUpdateRequestAttributes) GetMonitorIdentifier() DowntimeMonitorIdentifier {
	if o == nil || o.MonitorIdentifier == nil {
		var ret DowntimeMonitorIdentifier
		return ret
	}
	return *o.MonitorIdentifier
}

// GetMonitorIdentifierOk returns a tuple with the MonitorIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeUpdateRequestAttributes) GetMonitorIdentifierOk() (*DowntimeMonitorIdentifier, bool) {
	if o == nil || o.MonitorIdentifier == nil {
		return nil, false
	}
	return o.MonitorIdentifier, true
}

// HasMonitorIdentifier returns a boolean if a field has been set.
func (o *DowntimeUpdateRequestAttributes) HasMonitorIdentifier() bool {
	return o != nil && o.MonitorIdentifier != nil
}

// SetMonitorIdentifier gets a reference to the given DowntimeMonitorIdentifier and assigns it to the MonitorIdentifier field.
func (o *DowntimeUpdateRequestAttributes) SetMonitorIdentifier(v DowntimeMonitorIdentifier) {
	o.MonitorIdentifier = &v
}

// GetMuteFirstRecoveryNotification returns the MuteFirstRecoveryNotification field value if set, zero value otherwise.
func (o *DowntimeUpdateRequestAttributes) GetMuteFirstRecoveryNotification() bool {
	if o == nil || o.MuteFirstRecoveryNotification == nil {
		var ret bool
		return ret
	}
	return *o.MuteFirstRecoveryNotification
}

// GetMuteFirstRecoveryNotificationOk returns a tuple with the MuteFirstRecoveryNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeUpdateRequestAttributes) GetMuteFirstRecoveryNotificationOk() (*bool, bool) {
	if o == nil || o.MuteFirstRecoveryNotification == nil {
		return nil, false
	}
	return o.MuteFirstRecoveryNotification, true
}

// HasMuteFirstRecoveryNotification returns a boolean if a field has been set.
func (o *DowntimeUpdateRequestAttributes) HasMuteFirstRecoveryNotification() bool {
	return o != nil && o.MuteFirstRecoveryNotification != nil
}

// SetMuteFirstRecoveryNotification gets a reference to the given bool and assigns it to the MuteFirstRecoveryNotification field.
func (o *DowntimeUpdateRequestAttributes) SetMuteFirstRecoveryNotification(v bool) {
	o.MuteFirstRecoveryNotification = &v
}

// GetNotifyEndStates returns the NotifyEndStates field value if set, zero value otherwise.
func (o *DowntimeUpdateRequestAttributes) GetNotifyEndStates() []DowntimeNotifyEndStateTypes {
	if o == nil || o.NotifyEndStates == nil {
		var ret []DowntimeNotifyEndStateTypes
		return ret
	}
	return o.NotifyEndStates
}

// GetNotifyEndStatesOk returns a tuple with the NotifyEndStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeUpdateRequestAttributes) GetNotifyEndStatesOk() (*[]DowntimeNotifyEndStateTypes, bool) {
	if o == nil || o.NotifyEndStates == nil {
		return nil, false
	}
	return &o.NotifyEndStates, true
}

// HasNotifyEndStates returns a boolean if a field has been set.
func (o *DowntimeUpdateRequestAttributes) HasNotifyEndStates() bool {
	return o != nil && o.NotifyEndStates != nil
}

// SetNotifyEndStates gets a reference to the given []DowntimeNotifyEndStateTypes and assigns it to the NotifyEndStates field.
func (o *DowntimeUpdateRequestAttributes) SetNotifyEndStates(v []DowntimeNotifyEndStateTypes) {
	o.NotifyEndStates = v
}

// GetNotifyEndTypes returns the NotifyEndTypes field value if set, zero value otherwise.
func (o *DowntimeUpdateRequestAttributes) GetNotifyEndTypes() []DowntimeNotifyEndStateActions {
	if o == nil || o.NotifyEndTypes == nil {
		var ret []DowntimeNotifyEndStateActions
		return ret
	}
	return o.NotifyEndTypes
}

// GetNotifyEndTypesOk returns a tuple with the NotifyEndTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeUpdateRequestAttributes) GetNotifyEndTypesOk() (*[]DowntimeNotifyEndStateActions, bool) {
	if o == nil || o.NotifyEndTypes == nil {
		return nil, false
	}
	return &o.NotifyEndTypes, true
}

// HasNotifyEndTypes returns a boolean if a field has been set.
func (o *DowntimeUpdateRequestAttributes) HasNotifyEndTypes() bool {
	return o != nil && o.NotifyEndTypes != nil
}

// SetNotifyEndTypes gets a reference to the given []DowntimeNotifyEndStateActions and assigns it to the NotifyEndTypes field.
func (o *DowntimeUpdateRequestAttributes) SetNotifyEndTypes(v []DowntimeNotifyEndStateActions) {
	o.NotifyEndTypes = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *DowntimeUpdateRequestAttributes) GetSchedule() DowntimeScheduleUpdateRequest {
	if o == nil || o.Schedule == nil {
		var ret DowntimeScheduleUpdateRequest
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeUpdateRequestAttributes) GetScheduleOk() (*DowntimeScheduleUpdateRequest, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *DowntimeUpdateRequestAttributes) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given DowntimeScheduleUpdateRequest and assigns it to the Schedule field.
func (o *DowntimeUpdateRequestAttributes) SetSchedule(v DowntimeScheduleUpdateRequest) {
	o.Schedule = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *DowntimeUpdateRequestAttributes) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeUpdateRequestAttributes) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *DowntimeUpdateRequestAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *DowntimeUpdateRequestAttributes) SetScope(v string) {
	o.Scope = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeUpdateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayTimezone.IsSet() {
		toSerialize["display_timezone"] = o.DisplayTimezone.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeUpdateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayTimezone               datadog.NullableString          `json:"display_timezone,omitempty"`
		Message                       datadog.NullableString          `json:"message,omitempty"`
		MonitorIdentifier             *DowntimeMonitorIdentifier      `json:"monitor_identifier,omitempty"`
		MuteFirstRecoveryNotification *bool                           `json:"mute_first_recovery_notification,omitempty"`
		NotifyEndStates               []DowntimeNotifyEndStateTypes   `json:"notify_end_states,omitempty"`
		NotifyEndTypes                []DowntimeNotifyEndStateActions `json:"notify_end_types,omitempty"`
		Schedule                      *DowntimeScheduleUpdateRequest  `json:"schedule,omitempty"`
		Scope                         *string                         `json:"scope,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_timezone", "message", "monitor_identifier", "mute_first_recovery_notification", "notify_end_states", "notify_end_types", "schedule", "scope"})
	} else {
		return err
	}
	o.DisplayTimezone = all.DisplayTimezone
	o.Message = all.Message
	o.MonitorIdentifier = all.MonitorIdentifier
	o.MuteFirstRecoveryNotification = all.MuteFirstRecoveryNotification
	o.NotifyEndStates = all.NotifyEndStates
	o.NotifyEndTypes = all.NotifyEndTypes
	o.Schedule = all.Schedule
	o.Scope = all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
