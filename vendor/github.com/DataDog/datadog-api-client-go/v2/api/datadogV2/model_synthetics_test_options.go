// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestOptions Object describing the extra options for a Synthetic test.
type SyntheticsTestOptions struct {
	// Minimum amount of time in failure required to trigger an alert.
	MinFailureDuration *int64 `json:"min_failure_duration,omitempty"`
	// Minimum number of locations in failure required to trigger
	// an alert.
	MinLocationFailed *int64 `json:"min_location_failed,omitempty"`
	// The monitor name is used for the alert title as well as for all monitor dashboard widgets and SLOs.
	MonitorName *string `json:"monitor_name,omitempty"`
	// Object containing the options for a Synthetic test as a monitor
	// (for example, renotification).
	MonitorOptions *SyntheticsTestOptionsMonitorOptions `json:"monitor_options,omitempty"`
	// Integer from 1 (high) to 5 (low) indicating alert severity.
	MonitorPriority *int32 `json:"monitor_priority,omitempty"`
	// A list of role identifiers that can be pulled from the Roles API, for restricting read and write access. This field is deprecated. Use the restriction policies API to manage permissions.
	// Deprecated
	RestrictedRoles []string `json:"restricted_roles,omitempty"`
	// Object describing the retry strategy to apply to a Synthetic test.
	Retry *SyntheticsTestOptionsRetry `json:"retry,omitempty"`
	// Object containing timeframes and timezone used for advanced scheduling.
	Scheduling *SyntheticsTestOptionsScheduling `json:"scheduling,omitempty"`
	// The frequency at which to run the Synthetic test (in seconds).
	TickEvery *int64 `json:"tick_every,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestOptions instantiates a new SyntheticsTestOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestOptions() *SyntheticsTestOptions {
	this := SyntheticsTestOptions{}
	return &this
}

// NewSyntheticsTestOptionsWithDefaults instantiates a new SyntheticsTestOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestOptionsWithDefaults() *SyntheticsTestOptions {
	this := SyntheticsTestOptions{}
	return &this
}

// GetMinFailureDuration returns the MinFailureDuration field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetMinFailureDuration() int64 {
	if o == nil || o.MinFailureDuration == nil {
		var ret int64
		return ret
	}
	return *o.MinFailureDuration
}

// GetMinFailureDurationOk returns a tuple with the MinFailureDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetMinFailureDurationOk() (*int64, bool) {
	if o == nil || o.MinFailureDuration == nil {
		return nil, false
	}
	return o.MinFailureDuration, true
}

// HasMinFailureDuration returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasMinFailureDuration() bool {
	return o != nil && o.MinFailureDuration != nil
}

// SetMinFailureDuration gets a reference to the given int64 and assigns it to the MinFailureDuration field.
func (o *SyntheticsTestOptions) SetMinFailureDuration(v int64) {
	o.MinFailureDuration = &v
}

// GetMinLocationFailed returns the MinLocationFailed field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetMinLocationFailed() int64 {
	if o == nil || o.MinLocationFailed == nil {
		var ret int64
		return ret
	}
	return *o.MinLocationFailed
}

// GetMinLocationFailedOk returns a tuple with the MinLocationFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetMinLocationFailedOk() (*int64, bool) {
	if o == nil || o.MinLocationFailed == nil {
		return nil, false
	}
	return o.MinLocationFailed, true
}

// HasMinLocationFailed returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasMinLocationFailed() bool {
	return o != nil && o.MinLocationFailed != nil
}

// SetMinLocationFailed gets a reference to the given int64 and assigns it to the MinLocationFailed field.
func (o *SyntheticsTestOptions) SetMinLocationFailed(v int64) {
	o.MinLocationFailed = &v
}

// GetMonitorName returns the MonitorName field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetMonitorName() string {
	if o == nil || o.MonitorName == nil {
		var ret string
		return ret
	}
	return *o.MonitorName
}

// GetMonitorNameOk returns a tuple with the MonitorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetMonitorNameOk() (*string, bool) {
	if o == nil || o.MonitorName == nil {
		return nil, false
	}
	return o.MonitorName, true
}

// HasMonitorName returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasMonitorName() bool {
	return o != nil && o.MonitorName != nil
}

// SetMonitorName gets a reference to the given string and assigns it to the MonitorName field.
func (o *SyntheticsTestOptions) SetMonitorName(v string) {
	o.MonitorName = &v
}

// GetMonitorOptions returns the MonitorOptions field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetMonitorOptions() SyntheticsTestOptionsMonitorOptions {
	if o == nil || o.MonitorOptions == nil {
		var ret SyntheticsTestOptionsMonitorOptions
		return ret
	}
	return *o.MonitorOptions
}

// GetMonitorOptionsOk returns a tuple with the MonitorOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetMonitorOptionsOk() (*SyntheticsTestOptionsMonitorOptions, bool) {
	if o == nil || o.MonitorOptions == nil {
		return nil, false
	}
	return o.MonitorOptions, true
}

// HasMonitorOptions returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasMonitorOptions() bool {
	return o != nil && o.MonitorOptions != nil
}

// SetMonitorOptions gets a reference to the given SyntheticsTestOptionsMonitorOptions and assigns it to the MonitorOptions field.
func (o *SyntheticsTestOptions) SetMonitorOptions(v SyntheticsTestOptionsMonitorOptions) {
	o.MonitorOptions = &v
}

// GetMonitorPriority returns the MonitorPriority field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetMonitorPriority() int32 {
	if o == nil || o.MonitorPriority == nil {
		var ret int32
		return ret
	}
	return *o.MonitorPriority
}

// GetMonitorPriorityOk returns a tuple with the MonitorPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetMonitorPriorityOk() (*int32, bool) {
	if o == nil || o.MonitorPriority == nil {
		return nil, false
	}
	return o.MonitorPriority, true
}

// HasMonitorPriority returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasMonitorPriority() bool {
	return o != nil && o.MonitorPriority != nil
}

// SetMonitorPriority gets a reference to the given int32 and assigns it to the MonitorPriority field.
func (o *SyntheticsTestOptions) SetMonitorPriority(v int32) {
	o.MonitorPriority = &v
}

// GetRestrictedRoles returns the RestrictedRoles field value if set, zero value otherwise.
// Deprecated
func (o *SyntheticsTestOptions) GetRestrictedRoles() []string {
	if o == nil || o.RestrictedRoles == nil {
		var ret []string
		return ret
	}
	return o.RestrictedRoles
}

// GetRestrictedRolesOk returns a tuple with the RestrictedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SyntheticsTestOptions) GetRestrictedRolesOk() (*[]string, bool) {
	if o == nil || o.RestrictedRoles == nil {
		return nil, false
	}
	return &o.RestrictedRoles, true
}

// HasRestrictedRoles returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasRestrictedRoles() bool {
	return o != nil && o.RestrictedRoles != nil
}

// SetRestrictedRoles gets a reference to the given []string and assigns it to the RestrictedRoles field.
// Deprecated
func (o *SyntheticsTestOptions) SetRestrictedRoles(v []string) {
	o.RestrictedRoles = v
}

// GetRetry returns the Retry field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetRetry() SyntheticsTestOptionsRetry {
	if o == nil || o.Retry == nil {
		var ret SyntheticsTestOptionsRetry
		return ret
	}
	return *o.Retry
}

// GetRetryOk returns a tuple with the Retry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetRetryOk() (*SyntheticsTestOptionsRetry, bool) {
	if o == nil || o.Retry == nil {
		return nil, false
	}
	return o.Retry, true
}

// HasRetry returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasRetry() bool {
	return o != nil && o.Retry != nil
}

// SetRetry gets a reference to the given SyntheticsTestOptionsRetry and assigns it to the Retry field.
func (o *SyntheticsTestOptions) SetRetry(v SyntheticsTestOptionsRetry) {
	o.Retry = &v
}

// GetScheduling returns the Scheduling field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetScheduling() SyntheticsTestOptionsScheduling {
	if o == nil || o.Scheduling == nil {
		var ret SyntheticsTestOptionsScheduling
		return ret
	}
	return *o.Scheduling
}

// GetSchedulingOk returns a tuple with the Scheduling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetSchedulingOk() (*SyntheticsTestOptionsScheduling, bool) {
	if o == nil || o.Scheduling == nil {
		return nil, false
	}
	return o.Scheduling, true
}

// HasScheduling returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasScheduling() bool {
	return o != nil && o.Scheduling != nil
}

// SetScheduling gets a reference to the given SyntheticsTestOptionsScheduling and assigns it to the Scheduling field.
func (o *SyntheticsTestOptions) SetScheduling(v SyntheticsTestOptionsScheduling) {
	o.Scheduling = &v
}

// GetTickEvery returns the TickEvery field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetTickEvery() int64 {
	if o == nil || o.TickEvery == nil {
		var ret int64
		return ret
	}
	return *o.TickEvery
}

// GetTickEveryOk returns a tuple with the TickEvery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetTickEveryOk() (*int64, bool) {
	if o == nil || o.TickEvery == nil {
		return nil, false
	}
	return o.TickEvery, true
}

// HasTickEvery returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasTickEvery() bool {
	return o != nil && o.TickEvery != nil
}

// SetTickEvery gets a reference to the given int64 and assigns it to the TickEvery field.
func (o *SyntheticsTestOptions) SetTickEvery(v int64) {
	o.TickEvery = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MinFailureDuration != nil {
		toSerialize["min_failure_duration"] = o.MinFailureDuration
	}
	if o.MinLocationFailed != nil {
		toSerialize["min_location_failed"] = o.MinLocationFailed
	}
	if o.MonitorName != nil {
		toSerialize["monitor_name"] = o.MonitorName
	}
	if o.MonitorOptions != nil {
		toSerialize["monitor_options"] = o.MonitorOptions
	}
	if o.MonitorPriority != nil {
		toSerialize["monitor_priority"] = o.MonitorPriority
	}
	if o.RestrictedRoles != nil {
		toSerialize["restricted_roles"] = o.RestrictedRoles
	}
	if o.Retry != nil {
		toSerialize["retry"] = o.Retry
	}
	if o.Scheduling != nil {
		toSerialize["scheduling"] = o.Scheduling
	}
	if o.TickEvery != nil {
		toSerialize["tick_every"] = o.TickEvery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MinFailureDuration *int64                               `json:"min_failure_duration,omitempty"`
		MinLocationFailed  *int64                               `json:"min_location_failed,omitempty"`
		MonitorName        *string                              `json:"monitor_name,omitempty"`
		MonitorOptions     *SyntheticsTestOptionsMonitorOptions `json:"monitor_options,omitempty"`
		MonitorPriority    *int32                               `json:"monitor_priority,omitempty"`
		RestrictedRoles    []string                             `json:"restricted_roles,omitempty"`
		Retry              *SyntheticsTestOptionsRetry          `json:"retry,omitempty"`
		Scheduling         *SyntheticsTestOptionsScheduling     `json:"scheduling,omitempty"`
		TickEvery          *int64                               `json:"tick_every,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"min_failure_duration", "min_location_failed", "monitor_name", "monitor_options", "monitor_priority", "restricted_roles", "retry", "scheduling", "tick_every"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MinFailureDuration = all.MinFailureDuration
	o.MinLocationFailed = all.MinLocationFailed
	o.MonitorName = all.MonitorName
	if all.MonitorOptions != nil && all.MonitorOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MonitorOptions = all.MonitorOptions
	o.MonitorPriority = all.MonitorPriority
	o.RestrictedRoles = all.RestrictedRoles
	if all.Retry != nil && all.Retry.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Retry = all.Retry
	if all.Scheduling != nil && all.Scheduling.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scheduling = all.Scheduling
	o.TickEvery = all.TickEvery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
