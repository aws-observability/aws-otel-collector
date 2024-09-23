// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeCreateRequestAttributes Downtime details.
type DowntimeCreateRequestAttributes struct {
	// The timezone in which to display the downtime's start and end times in Datadog applications. This is not used
	// as an offset for scheduling.
	DisplayTimezone datadog.NullableString `json:"display_timezone,omitempty"`
	// A message to include with notifications for this downtime. Email notifications can be sent to specific users
	// by using the same `@username` notation as events.
	Message datadog.NullableString `json:"message,omitempty"`
	// Monitor identifier for the downtime.
	MonitorIdentifier DowntimeMonitorIdentifier `json:"monitor_identifier"`
	// If the first recovery notification during a downtime should be muted.
	MuteFirstRecoveryNotification *bool `json:"mute_first_recovery_notification,omitempty"`
	// States that will trigger a monitor notification when the `notify_end_types` action occurs.
	NotifyEndStates []DowntimeNotifyEndStateTypes `json:"notify_end_states,omitempty"`
	// Actions that will trigger a monitor notification if the downtime is in the `notify_end_types` state.
	NotifyEndTypes []DowntimeNotifyEndStateActions `json:"notify_end_types,omitempty"`
	// Schedule for the downtime.
	Schedule *DowntimeScheduleCreateRequest `json:"schedule,omitempty"`
	// The scope to which the downtime applies. Must follow the [common search syntax](https://docs.datadoghq.com/logs/explorer/search_syntax/).
	Scope string `json:"scope"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDowntimeCreateRequestAttributes instantiates a new DowntimeCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeCreateRequestAttributes(monitorIdentifier DowntimeMonitorIdentifier, scope string) *DowntimeCreateRequestAttributes {
	this := DowntimeCreateRequestAttributes{}
	var displayTimezone string = "UTC"
	this.DisplayTimezone = *datadog.NewNullableString(&displayTimezone)
	this.MonitorIdentifier = monitorIdentifier
	this.Scope = scope
	return &this
}

// NewDowntimeCreateRequestAttributesWithDefaults instantiates a new DowntimeCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeCreateRequestAttributesWithDefaults() *DowntimeCreateRequestAttributes {
	this := DowntimeCreateRequestAttributes{}
	var displayTimezone string = "UTC"
	this.DisplayTimezone = *datadog.NewNullableString(&displayTimezone)
	return &this
}

// GetDisplayTimezone returns the DisplayTimezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeCreateRequestAttributes) GetDisplayTimezone() string {
	if o == nil || o.DisplayTimezone.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayTimezone.Get()
}

// GetDisplayTimezoneOk returns a tuple with the DisplayTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeCreateRequestAttributes) GetDisplayTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayTimezone.Get(), o.DisplayTimezone.IsSet()
}

// HasDisplayTimezone returns a boolean if a field has been set.
func (o *DowntimeCreateRequestAttributes) HasDisplayTimezone() bool {
	return o != nil && o.DisplayTimezone.IsSet()
}

// SetDisplayTimezone gets a reference to the given datadog.NullableString and assigns it to the DisplayTimezone field.
func (o *DowntimeCreateRequestAttributes) SetDisplayTimezone(v string) {
	o.DisplayTimezone.Set(&v)
}

// SetDisplayTimezoneNil sets the value for DisplayTimezone to be an explicit nil.
func (o *DowntimeCreateRequestAttributes) SetDisplayTimezoneNil() {
	o.DisplayTimezone.Set(nil)
}

// UnsetDisplayTimezone ensures that no value is present for DisplayTimezone, not even an explicit nil.
func (o *DowntimeCreateRequestAttributes) UnsetDisplayTimezone() {
	o.DisplayTimezone.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeCreateRequestAttributes) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeCreateRequestAttributes) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *DowntimeCreateRequestAttributes) HasMessage() bool {
	return o != nil && o.Message.IsSet()
}

// SetMessage gets a reference to the given datadog.NullableString and assigns it to the Message field.
func (o *DowntimeCreateRequestAttributes) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil.
func (o *DowntimeCreateRequestAttributes) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil.
func (o *DowntimeCreateRequestAttributes) UnsetMessage() {
	o.Message.Unset()
}

// GetMonitorIdentifier returns the MonitorIdentifier field value.
func (o *DowntimeCreateRequestAttributes) GetMonitorIdentifier() DowntimeMonitorIdentifier {
	if o == nil {
		var ret DowntimeMonitorIdentifier
		return ret
	}
	return o.MonitorIdentifier
}

// GetMonitorIdentifierOk returns a tuple with the MonitorIdentifier field value
// and a boolean to check if the value has been set.
func (o *DowntimeCreateRequestAttributes) GetMonitorIdentifierOk() (*DowntimeMonitorIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorIdentifier, true
}

// SetMonitorIdentifier sets field value.
func (o *DowntimeCreateRequestAttributes) SetMonitorIdentifier(v DowntimeMonitorIdentifier) {
	o.MonitorIdentifier = v
}

// GetMuteFirstRecoveryNotification returns the MuteFirstRecoveryNotification field value if set, zero value otherwise.
func (o *DowntimeCreateRequestAttributes) GetMuteFirstRecoveryNotification() bool {
	if o == nil || o.MuteFirstRecoveryNotification == nil {
		var ret bool
		return ret
	}
	return *o.MuteFirstRecoveryNotification
}

// GetMuteFirstRecoveryNotificationOk returns a tuple with the MuteFirstRecoveryNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeCreateRequestAttributes) GetMuteFirstRecoveryNotificationOk() (*bool, bool) {
	if o == nil || o.MuteFirstRecoveryNotification == nil {
		return nil, false
	}
	return o.MuteFirstRecoveryNotification, true
}

// HasMuteFirstRecoveryNotification returns a boolean if a field has been set.
func (o *DowntimeCreateRequestAttributes) HasMuteFirstRecoveryNotification() bool {
	return o != nil && o.MuteFirstRecoveryNotification != nil
}

// SetMuteFirstRecoveryNotification gets a reference to the given bool and assigns it to the MuteFirstRecoveryNotification field.
func (o *DowntimeCreateRequestAttributes) SetMuteFirstRecoveryNotification(v bool) {
	o.MuteFirstRecoveryNotification = &v
}

// GetNotifyEndStates returns the NotifyEndStates field value if set, zero value otherwise.
func (o *DowntimeCreateRequestAttributes) GetNotifyEndStates() []DowntimeNotifyEndStateTypes {
	if o == nil || o.NotifyEndStates == nil {
		var ret []DowntimeNotifyEndStateTypes
		return ret
	}
	return o.NotifyEndStates
}

// GetNotifyEndStatesOk returns a tuple with the NotifyEndStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeCreateRequestAttributes) GetNotifyEndStatesOk() (*[]DowntimeNotifyEndStateTypes, bool) {
	if o == nil || o.NotifyEndStates == nil {
		return nil, false
	}
	return &o.NotifyEndStates, true
}

// HasNotifyEndStates returns a boolean if a field has been set.
func (o *DowntimeCreateRequestAttributes) HasNotifyEndStates() bool {
	return o != nil && o.NotifyEndStates != nil
}

// SetNotifyEndStates gets a reference to the given []DowntimeNotifyEndStateTypes and assigns it to the NotifyEndStates field.
func (o *DowntimeCreateRequestAttributes) SetNotifyEndStates(v []DowntimeNotifyEndStateTypes) {
	o.NotifyEndStates = v
}

// GetNotifyEndTypes returns the NotifyEndTypes field value if set, zero value otherwise.
func (o *DowntimeCreateRequestAttributes) GetNotifyEndTypes() []DowntimeNotifyEndStateActions {
	if o == nil || o.NotifyEndTypes == nil {
		var ret []DowntimeNotifyEndStateActions
		return ret
	}
	return o.NotifyEndTypes
}

// GetNotifyEndTypesOk returns a tuple with the NotifyEndTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeCreateRequestAttributes) GetNotifyEndTypesOk() (*[]DowntimeNotifyEndStateActions, bool) {
	if o == nil || o.NotifyEndTypes == nil {
		return nil, false
	}
	return &o.NotifyEndTypes, true
}

// HasNotifyEndTypes returns a boolean if a field has been set.
func (o *DowntimeCreateRequestAttributes) HasNotifyEndTypes() bool {
	return o != nil && o.NotifyEndTypes != nil
}

// SetNotifyEndTypes gets a reference to the given []DowntimeNotifyEndStateActions and assigns it to the NotifyEndTypes field.
func (o *DowntimeCreateRequestAttributes) SetNotifyEndTypes(v []DowntimeNotifyEndStateActions) {
	o.NotifyEndTypes = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *DowntimeCreateRequestAttributes) GetSchedule() DowntimeScheduleCreateRequest {
	if o == nil || o.Schedule == nil {
		var ret DowntimeScheduleCreateRequest
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeCreateRequestAttributes) GetScheduleOk() (*DowntimeScheduleCreateRequest, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *DowntimeCreateRequestAttributes) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given DowntimeScheduleCreateRequest and assigns it to the Schedule field.
func (o *DowntimeCreateRequestAttributes) SetSchedule(v DowntimeScheduleCreateRequest) {
	o.Schedule = &v
}

// GetScope returns the Scope field value.
func (o *DowntimeCreateRequestAttributes) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *DowntimeCreateRequestAttributes) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *DowntimeCreateRequestAttributes) SetScope(v string) {
	o.Scope = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeCreateRequestAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["monitor_identifier"] = o.MonitorIdentifier
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
	toSerialize["scope"] = o.Scope

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayTimezone               datadog.NullableString          `json:"display_timezone,omitempty"`
		Message                       datadog.NullableString          `json:"message,omitempty"`
		MonitorIdentifier             *DowntimeMonitorIdentifier      `json:"monitor_identifier"`
		MuteFirstRecoveryNotification *bool                           `json:"mute_first_recovery_notification,omitempty"`
		NotifyEndStates               []DowntimeNotifyEndStateTypes   `json:"notify_end_states,omitempty"`
		NotifyEndTypes                []DowntimeNotifyEndStateActions `json:"notify_end_types,omitempty"`
		Schedule                      *DowntimeScheduleCreateRequest  `json:"schedule,omitempty"`
		Scope                         *string                         `json:"scope"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MonitorIdentifier == nil {
		return fmt.Errorf("required field monitor_identifier missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_timezone", "message", "monitor_identifier", "mute_first_recovery_notification", "notify_end_states", "notify_end_types", "schedule", "scope"})
	} else {
		return err
	}
	o.DisplayTimezone = all.DisplayTimezone
	o.Message = all.Message
	o.MonitorIdentifier = *all.MonitorIdentifier
	o.MuteFirstRecoveryNotification = all.MuteFirstRecoveryNotification
	o.NotifyEndStates = all.NotifyEndStates
	o.NotifyEndTypes = all.NotifyEndTypes
	o.Schedule = all.Schedule
	o.Scope = *all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
