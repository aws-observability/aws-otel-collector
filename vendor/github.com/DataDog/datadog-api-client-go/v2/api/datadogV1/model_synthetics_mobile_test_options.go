// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTestOptions Object describing the extra options for a Synthetic test.
type SyntheticsMobileTestOptions struct {
	// A boolean to set if an application crash would mark the test as failed.
	AllowApplicationCrash *bool `json:"allowApplicationCrash,omitempty"`
	// Array of bindings used for the mobile test.
	Bindings []SyntheticsTestRestrictionPolicyBinding `json:"bindings,omitempty"`
	// CI/CD options for a Synthetic test.
	Ci *SyntheticsTestCiOptions `json:"ci,omitempty"`
	// The default timeout for steps in the test (in seconds).
	DefaultStepTimeout *int32 `json:"defaultStepTimeout,omitempty"`
	// For mobile test, array with the different device IDs used to run the test.
	DeviceIds []string `json:"device_ids"`
	// A boolean to disable auto accepting alerts.
	DisableAutoAcceptAlert *bool `json:"disableAutoAcceptAlert,omitempty"`
	// Minimum amount of time in failure required to trigger an alert.
	MinFailureDuration *int64 `json:"min_failure_duration,omitempty"`
	// Mobile application for mobile synthetics test.
	MobileApplication SyntheticsMobileTestsMobileApplication `json:"mobileApplication"`
	// The monitor name is used for the alert title as well as for all monitor dashboard widgets and SLOs.
	MonitorName *string `json:"monitor_name,omitempty"`
	// Object containing the options for a Synthetic test as a monitor
	// (for example, renotification).
	MonitorOptions *SyntheticsTestOptionsMonitorOptions `json:"monitor_options,omitempty"`
	// Integer from 1 (high) to 5 (low) indicating alert severity.
	MonitorPriority *int32 `json:"monitor_priority,omitempty"`
	// A boolean set to not take a screenshot for the step.
	NoScreenshot *bool `json:"noScreenshot,omitempty"`
	// A list of role identifiers that can be pulled from the Roles API, for restricting read and write access.
	RestrictedRoles []string `json:"restricted_roles,omitempty"`
	// Object describing the retry strategy to apply to a Synthetic test.
	Retry *SyntheticsTestOptionsRetry `json:"retry,omitempty"`
	// Object containing timeframes and timezone used for advanced scheduling.
	Scheduling *SyntheticsTestOptionsScheduling `json:"scheduling,omitempty"`
	// The frequency at which to run the Synthetic test (in seconds).
	TickEvery int64 `json:"tick_every"`
	// The level of verbosity for the mobile test. This field can not be set by a user.
	Verbosity *int32 `json:"verbosity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileTestOptions instantiates a new SyntheticsMobileTestOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileTestOptions(deviceIds []string, mobileApplication SyntheticsMobileTestsMobileApplication, tickEvery int64) *SyntheticsMobileTestOptions {
	this := SyntheticsMobileTestOptions{}
	this.DeviceIds = deviceIds
	this.MobileApplication = mobileApplication
	this.TickEvery = tickEvery
	return &this
}

// NewSyntheticsMobileTestOptionsWithDefaults instantiates a new SyntheticsMobileTestOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileTestOptionsWithDefaults() *SyntheticsMobileTestOptions {
	this := SyntheticsMobileTestOptions{}
	return &this
}

// GetAllowApplicationCrash returns the AllowApplicationCrash field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetAllowApplicationCrash() bool {
	if o == nil || o.AllowApplicationCrash == nil {
		var ret bool
		return ret
	}
	return *o.AllowApplicationCrash
}

// GetAllowApplicationCrashOk returns a tuple with the AllowApplicationCrash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetAllowApplicationCrashOk() (*bool, bool) {
	if o == nil || o.AllowApplicationCrash == nil {
		return nil, false
	}
	return o.AllowApplicationCrash, true
}

// HasAllowApplicationCrash returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasAllowApplicationCrash() bool {
	return o != nil && o.AllowApplicationCrash != nil
}

// SetAllowApplicationCrash gets a reference to the given bool and assigns it to the AllowApplicationCrash field.
func (o *SyntheticsMobileTestOptions) SetAllowApplicationCrash(v bool) {
	o.AllowApplicationCrash = &v
}

// GetBindings returns the Bindings field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetBindings() []SyntheticsTestRestrictionPolicyBinding {
	if o == nil || o.Bindings == nil {
		var ret []SyntheticsTestRestrictionPolicyBinding
		return ret
	}
	return o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetBindingsOk() (*[]SyntheticsTestRestrictionPolicyBinding, bool) {
	if o == nil || o.Bindings == nil {
		return nil, false
	}
	return &o.Bindings, true
}

// HasBindings returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasBindings() bool {
	return o != nil && o.Bindings != nil
}

// SetBindings gets a reference to the given []SyntheticsTestRestrictionPolicyBinding and assigns it to the Bindings field.
func (o *SyntheticsMobileTestOptions) SetBindings(v []SyntheticsTestRestrictionPolicyBinding) {
	o.Bindings = v
}

// GetCi returns the Ci field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetCi() SyntheticsTestCiOptions {
	if o == nil || o.Ci == nil {
		var ret SyntheticsTestCiOptions
		return ret
	}
	return *o.Ci
}

// GetCiOk returns a tuple with the Ci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetCiOk() (*SyntheticsTestCiOptions, bool) {
	if o == nil || o.Ci == nil {
		return nil, false
	}
	return o.Ci, true
}

// HasCi returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasCi() bool {
	return o != nil && o.Ci != nil
}

// SetCi gets a reference to the given SyntheticsTestCiOptions and assigns it to the Ci field.
func (o *SyntheticsMobileTestOptions) SetCi(v SyntheticsTestCiOptions) {
	o.Ci = &v
}

// GetDefaultStepTimeout returns the DefaultStepTimeout field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetDefaultStepTimeout() int32 {
	if o == nil || o.DefaultStepTimeout == nil {
		var ret int32
		return ret
	}
	return *o.DefaultStepTimeout
}

// GetDefaultStepTimeoutOk returns a tuple with the DefaultStepTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetDefaultStepTimeoutOk() (*int32, bool) {
	if o == nil || o.DefaultStepTimeout == nil {
		return nil, false
	}
	return o.DefaultStepTimeout, true
}

// HasDefaultStepTimeout returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasDefaultStepTimeout() bool {
	return o != nil && o.DefaultStepTimeout != nil
}

// SetDefaultStepTimeout gets a reference to the given int32 and assigns it to the DefaultStepTimeout field.
func (o *SyntheticsMobileTestOptions) SetDefaultStepTimeout(v int32) {
	o.DefaultStepTimeout = &v
}

// GetDeviceIds returns the DeviceIds field value.
func (o *SyntheticsMobileTestOptions) GetDeviceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetDeviceIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceIds, true
}

// SetDeviceIds sets field value.
func (o *SyntheticsMobileTestOptions) SetDeviceIds(v []string) {
	o.DeviceIds = v
}

// GetDisableAutoAcceptAlert returns the DisableAutoAcceptAlert field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetDisableAutoAcceptAlert() bool {
	if o == nil || o.DisableAutoAcceptAlert == nil {
		var ret bool
		return ret
	}
	return *o.DisableAutoAcceptAlert
}

// GetDisableAutoAcceptAlertOk returns a tuple with the DisableAutoAcceptAlert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetDisableAutoAcceptAlertOk() (*bool, bool) {
	if o == nil || o.DisableAutoAcceptAlert == nil {
		return nil, false
	}
	return o.DisableAutoAcceptAlert, true
}

// HasDisableAutoAcceptAlert returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasDisableAutoAcceptAlert() bool {
	return o != nil && o.DisableAutoAcceptAlert != nil
}

// SetDisableAutoAcceptAlert gets a reference to the given bool and assigns it to the DisableAutoAcceptAlert field.
func (o *SyntheticsMobileTestOptions) SetDisableAutoAcceptAlert(v bool) {
	o.DisableAutoAcceptAlert = &v
}

// GetMinFailureDuration returns the MinFailureDuration field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetMinFailureDuration() int64 {
	if o == nil || o.MinFailureDuration == nil {
		var ret int64
		return ret
	}
	return *o.MinFailureDuration
}

// GetMinFailureDurationOk returns a tuple with the MinFailureDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetMinFailureDurationOk() (*int64, bool) {
	if o == nil || o.MinFailureDuration == nil {
		return nil, false
	}
	return o.MinFailureDuration, true
}

// HasMinFailureDuration returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasMinFailureDuration() bool {
	return o != nil && o.MinFailureDuration != nil
}

// SetMinFailureDuration gets a reference to the given int64 and assigns it to the MinFailureDuration field.
func (o *SyntheticsMobileTestOptions) SetMinFailureDuration(v int64) {
	o.MinFailureDuration = &v
}

// GetMobileApplication returns the MobileApplication field value.
func (o *SyntheticsMobileTestOptions) GetMobileApplication() SyntheticsMobileTestsMobileApplication {
	if o == nil {
		var ret SyntheticsMobileTestsMobileApplication
		return ret
	}
	return o.MobileApplication
}

// GetMobileApplicationOk returns a tuple with the MobileApplication field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetMobileApplicationOk() (*SyntheticsMobileTestsMobileApplication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MobileApplication, true
}

// SetMobileApplication sets field value.
func (o *SyntheticsMobileTestOptions) SetMobileApplication(v SyntheticsMobileTestsMobileApplication) {
	o.MobileApplication = v
}

// GetMonitorName returns the MonitorName field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetMonitorName() string {
	if o == nil || o.MonitorName == nil {
		var ret string
		return ret
	}
	return *o.MonitorName
}

// GetMonitorNameOk returns a tuple with the MonitorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetMonitorNameOk() (*string, bool) {
	if o == nil || o.MonitorName == nil {
		return nil, false
	}
	return o.MonitorName, true
}

// HasMonitorName returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasMonitorName() bool {
	return o != nil && o.MonitorName != nil
}

// SetMonitorName gets a reference to the given string and assigns it to the MonitorName field.
func (o *SyntheticsMobileTestOptions) SetMonitorName(v string) {
	o.MonitorName = &v
}

// GetMonitorOptions returns the MonitorOptions field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetMonitorOptions() SyntheticsTestOptionsMonitorOptions {
	if o == nil || o.MonitorOptions == nil {
		var ret SyntheticsTestOptionsMonitorOptions
		return ret
	}
	return *o.MonitorOptions
}

// GetMonitorOptionsOk returns a tuple with the MonitorOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetMonitorOptionsOk() (*SyntheticsTestOptionsMonitorOptions, bool) {
	if o == nil || o.MonitorOptions == nil {
		return nil, false
	}
	return o.MonitorOptions, true
}

// HasMonitorOptions returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasMonitorOptions() bool {
	return o != nil && o.MonitorOptions != nil
}

// SetMonitorOptions gets a reference to the given SyntheticsTestOptionsMonitorOptions and assigns it to the MonitorOptions field.
func (o *SyntheticsMobileTestOptions) SetMonitorOptions(v SyntheticsTestOptionsMonitorOptions) {
	o.MonitorOptions = &v
}

// GetMonitorPriority returns the MonitorPriority field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetMonitorPriority() int32 {
	if o == nil || o.MonitorPriority == nil {
		var ret int32
		return ret
	}
	return *o.MonitorPriority
}

// GetMonitorPriorityOk returns a tuple with the MonitorPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetMonitorPriorityOk() (*int32, bool) {
	if o == nil || o.MonitorPriority == nil {
		return nil, false
	}
	return o.MonitorPriority, true
}

// HasMonitorPriority returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasMonitorPriority() bool {
	return o != nil && o.MonitorPriority != nil
}

// SetMonitorPriority gets a reference to the given int32 and assigns it to the MonitorPriority field.
func (o *SyntheticsMobileTestOptions) SetMonitorPriority(v int32) {
	o.MonitorPriority = &v
}

// GetNoScreenshot returns the NoScreenshot field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetNoScreenshot() bool {
	if o == nil || o.NoScreenshot == nil {
		var ret bool
		return ret
	}
	return *o.NoScreenshot
}

// GetNoScreenshotOk returns a tuple with the NoScreenshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetNoScreenshotOk() (*bool, bool) {
	if o == nil || o.NoScreenshot == nil {
		return nil, false
	}
	return o.NoScreenshot, true
}

// HasNoScreenshot returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasNoScreenshot() bool {
	return o != nil && o.NoScreenshot != nil
}

// SetNoScreenshot gets a reference to the given bool and assigns it to the NoScreenshot field.
func (o *SyntheticsMobileTestOptions) SetNoScreenshot(v bool) {
	o.NoScreenshot = &v
}

// GetRestrictedRoles returns the RestrictedRoles field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetRestrictedRoles() []string {
	if o == nil || o.RestrictedRoles == nil {
		var ret []string
		return ret
	}
	return o.RestrictedRoles
}

// GetRestrictedRolesOk returns a tuple with the RestrictedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetRestrictedRolesOk() (*[]string, bool) {
	if o == nil || o.RestrictedRoles == nil {
		return nil, false
	}
	return &o.RestrictedRoles, true
}

// HasRestrictedRoles returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasRestrictedRoles() bool {
	return o != nil && o.RestrictedRoles != nil
}

// SetRestrictedRoles gets a reference to the given []string and assigns it to the RestrictedRoles field.
func (o *SyntheticsMobileTestOptions) SetRestrictedRoles(v []string) {
	o.RestrictedRoles = v
}

// GetRetry returns the Retry field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetRetry() SyntheticsTestOptionsRetry {
	if o == nil || o.Retry == nil {
		var ret SyntheticsTestOptionsRetry
		return ret
	}
	return *o.Retry
}

// GetRetryOk returns a tuple with the Retry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetRetryOk() (*SyntheticsTestOptionsRetry, bool) {
	if o == nil || o.Retry == nil {
		return nil, false
	}
	return o.Retry, true
}

// HasRetry returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasRetry() bool {
	return o != nil && o.Retry != nil
}

// SetRetry gets a reference to the given SyntheticsTestOptionsRetry and assigns it to the Retry field.
func (o *SyntheticsMobileTestOptions) SetRetry(v SyntheticsTestOptionsRetry) {
	o.Retry = &v
}

// GetScheduling returns the Scheduling field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetScheduling() SyntheticsTestOptionsScheduling {
	if o == nil || o.Scheduling == nil {
		var ret SyntheticsTestOptionsScheduling
		return ret
	}
	return *o.Scheduling
}

// GetSchedulingOk returns a tuple with the Scheduling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetSchedulingOk() (*SyntheticsTestOptionsScheduling, bool) {
	if o == nil || o.Scheduling == nil {
		return nil, false
	}
	return o.Scheduling, true
}

// HasScheduling returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasScheduling() bool {
	return o != nil && o.Scheduling != nil
}

// SetScheduling gets a reference to the given SyntheticsTestOptionsScheduling and assigns it to the Scheduling field.
func (o *SyntheticsMobileTestOptions) SetScheduling(v SyntheticsTestOptionsScheduling) {
	o.Scheduling = &v
}

// GetTickEvery returns the TickEvery field value.
func (o *SyntheticsMobileTestOptions) GetTickEvery() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TickEvery
}

// GetTickEveryOk returns a tuple with the TickEvery field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetTickEveryOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TickEvery, true
}

// SetTickEvery sets field value.
func (o *SyntheticsMobileTestOptions) SetTickEvery(v int64) {
	o.TickEvery = v
}

// GetVerbosity returns the Verbosity field value if set, zero value otherwise.
func (o *SyntheticsMobileTestOptions) GetVerbosity() int32 {
	if o == nil || o.Verbosity == nil {
		var ret int32
		return ret
	}
	return *o.Verbosity
}

// GetVerbosityOk returns a tuple with the Verbosity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestOptions) GetVerbosityOk() (*int32, bool) {
	if o == nil || o.Verbosity == nil {
		return nil, false
	}
	return o.Verbosity, true
}

// HasVerbosity returns a boolean if a field has been set.
func (o *SyntheticsMobileTestOptions) HasVerbosity() bool {
	return o != nil && o.Verbosity != nil
}

// SetVerbosity gets a reference to the given int32 and assigns it to the Verbosity field.
func (o *SyntheticsMobileTestOptions) SetVerbosity(v int32) {
	o.Verbosity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileTestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowApplicationCrash != nil {
		toSerialize["allowApplicationCrash"] = o.AllowApplicationCrash
	}
	if o.Bindings != nil {
		toSerialize["bindings"] = o.Bindings
	}
	if o.Ci != nil {
		toSerialize["ci"] = o.Ci
	}
	if o.DefaultStepTimeout != nil {
		toSerialize["defaultStepTimeout"] = o.DefaultStepTimeout
	}
	toSerialize["device_ids"] = o.DeviceIds
	if o.DisableAutoAcceptAlert != nil {
		toSerialize["disableAutoAcceptAlert"] = o.DisableAutoAcceptAlert
	}
	if o.MinFailureDuration != nil {
		toSerialize["min_failure_duration"] = o.MinFailureDuration
	}
	toSerialize["mobileApplication"] = o.MobileApplication
	if o.MonitorName != nil {
		toSerialize["monitor_name"] = o.MonitorName
	}
	if o.MonitorOptions != nil {
		toSerialize["monitor_options"] = o.MonitorOptions
	}
	if o.MonitorPriority != nil {
		toSerialize["monitor_priority"] = o.MonitorPriority
	}
	if o.NoScreenshot != nil {
		toSerialize["noScreenshot"] = o.NoScreenshot
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
	toSerialize["tick_every"] = o.TickEvery
	if o.Verbosity != nil {
		toSerialize["verbosity"] = o.Verbosity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileTestOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowApplicationCrash  *bool                                    `json:"allowApplicationCrash,omitempty"`
		Bindings               []SyntheticsTestRestrictionPolicyBinding `json:"bindings,omitempty"`
		Ci                     *SyntheticsTestCiOptions                 `json:"ci,omitempty"`
		DefaultStepTimeout     *int32                                   `json:"defaultStepTimeout,omitempty"`
		DeviceIds              *[]string                                `json:"device_ids"`
		DisableAutoAcceptAlert *bool                                    `json:"disableAutoAcceptAlert,omitempty"`
		MinFailureDuration     *int64                                   `json:"min_failure_duration,omitempty"`
		MobileApplication      *SyntheticsMobileTestsMobileApplication  `json:"mobileApplication"`
		MonitorName            *string                                  `json:"monitor_name,omitempty"`
		MonitorOptions         *SyntheticsTestOptionsMonitorOptions     `json:"monitor_options,omitempty"`
		MonitorPriority        *int32                                   `json:"monitor_priority,omitempty"`
		NoScreenshot           *bool                                    `json:"noScreenshot,omitempty"`
		RestrictedRoles        []string                                 `json:"restricted_roles,omitempty"`
		Retry                  *SyntheticsTestOptionsRetry              `json:"retry,omitempty"`
		Scheduling             *SyntheticsTestOptionsScheduling         `json:"scheduling,omitempty"`
		TickEvery              *int64                                   `json:"tick_every"`
		Verbosity              *int32                                   `json:"verbosity,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DeviceIds == nil {
		return fmt.Errorf("required field device_ids missing")
	}
	if all.MobileApplication == nil {
		return fmt.Errorf("required field mobileApplication missing")
	}
	if all.TickEvery == nil {
		return fmt.Errorf("required field tick_every missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowApplicationCrash", "bindings", "ci", "defaultStepTimeout", "device_ids", "disableAutoAcceptAlert", "min_failure_duration", "mobileApplication", "monitor_name", "monitor_options", "monitor_priority", "noScreenshot", "restricted_roles", "retry", "scheduling", "tick_every", "verbosity"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowApplicationCrash = all.AllowApplicationCrash
	o.Bindings = all.Bindings
	if all.Ci != nil && all.Ci.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Ci = all.Ci
	o.DefaultStepTimeout = all.DefaultStepTimeout
	o.DeviceIds = *all.DeviceIds
	o.DisableAutoAcceptAlert = all.DisableAutoAcceptAlert
	o.MinFailureDuration = all.MinFailureDuration
	if all.MobileApplication.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MobileApplication = *all.MobileApplication
	o.MonitorName = all.MonitorName
	if all.MonitorOptions != nil && all.MonitorOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MonitorOptions = all.MonitorOptions
	o.MonitorPriority = all.MonitorPriority
	o.NoScreenshot = all.NoScreenshot
	o.RestrictedRoles = all.RestrictedRoles
	if all.Retry != nil && all.Retry.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Retry = all.Retry
	if all.Scheduling != nil && all.Scheduling.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scheduling = all.Scheduling
	o.TickEvery = *all.TickEvery
	o.Verbosity = all.Verbosity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
