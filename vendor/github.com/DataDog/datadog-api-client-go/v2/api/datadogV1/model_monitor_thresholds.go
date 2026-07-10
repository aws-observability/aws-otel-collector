// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorThresholds List of the different monitor threshold available.
type MonitorThresholds struct {
	// The monitor `CRITICAL` threshold.
	Critical *float64 `json:"critical,omitempty"`
	// Query evaluated as a dynamic `CRITICAL` threshold. Only supported on metric monitors with a formula query and options['variables']. Cannot be combined with static thresholds. This field is in preview.
	CriticalQuery *string `json:"critical_query,omitempty"`
	// The monitor `CRITICAL` recovery threshold.
	CriticalRecovery datadog.NullableFloat64 `json:"critical_recovery,omitempty"`
	// Query evaluated as a dynamic `CRITICAL` recovery threshold. Only supported on metric monitors with a formula query and options['variables']. Cannot be combined with static thresholds. This field is in preview.
	CriticalRecoveryQuery *string `json:"critical_recovery_query,omitempty"`
	// The monitor `OK` threshold.
	Ok datadog.NullableFloat64 `json:"ok,omitempty"`
	// The monitor UNKNOWN threshold.
	Unknown datadog.NullableFloat64 `json:"unknown,omitempty"`
	// The monitor `WARNING` threshold.
	Warning datadog.NullableFloat64 `json:"warning,omitempty"`
	// The monitor `WARNING` recovery threshold.
	WarningRecovery datadog.NullableFloat64 `json:"warning_recovery,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorThresholds instantiates a new MonitorThresholds object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorThresholds() *MonitorThresholds {
	this := MonitorThresholds{}
	return &this
}

// NewMonitorThresholdsWithDefaults instantiates a new MonitorThresholds object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorThresholdsWithDefaults() *MonitorThresholds {
	this := MonitorThresholds{}
	return &this
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *MonitorThresholds) GetCritical() float64 {
	if o == nil || o.Critical == nil {
		var ret float64
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetCriticalOk() (*float64, bool) {
	if o == nil || o.Critical == nil {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *MonitorThresholds) HasCritical() bool {
	return o != nil && o.Critical != nil
}

// SetCritical gets a reference to the given float64 and assigns it to the Critical field.
func (o *MonitorThresholds) SetCritical(v float64) {
	o.Critical = &v
}

// GetCriticalQuery returns the CriticalQuery field value if set, zero value otherwise.
func (o *MonitorThresholds) GetCriticalQuery() string {
	if o == nil || o.CriticalQuery == nil {
		var ret string
		return ret
	}
	return *o.CriticalQuery
}

// GetCriticalQueryOk returns a tuple with the CriticalQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetCriticalQueryOk() (*string, bool) {
	if o == nil || o.CriticalQuery == nil {
		return nil, false
	}
	return o.CriticalQuery, true
}

// HasCriticalQuery returns a boolean if a field has been set.
func (o *MonitorThresholds) HasCriticalQuery() bool {
	return o != nil && o.CriticalQuery != nil
}

// SetCriticalQuery gets a reference to the given string and assigns it to the CriticalQuery field.
func (o *MonitorThresholds) SetCriticalQuery(v string) {
	o.CriticalQuery = &v
}

// GetCriticalRecovery returns the CriticalRecovery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorThresholds) GetCriticalRecovery() float64 {
	if o == nil || o.CriticalRecovery.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CriticalRecovery.Get()
}

// GetCriticalRecoveryOk returns a tuple with the CriticalRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorThresholds) GetCriticalRecoveryOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CriticalRecovery.Get(), o.CriticalRecovery.IsSet()
}

// HasCriticalRecovery returns a boolean if a field has been set.
func (o *MonitorThresholds) HasCriticalRecovery() bool {
	return o != nil && o.CriticalRecovery.IsSet()
}

// SetCriticalRecovery gets a reference to the given datadog.NullableFloat64 and assigns it to the CriticalRecovery field.
func (o *MonitorThresholds) SetCriticalRecovery(v float64) {
	o.CriticalRecovery.Set(&v)
}

// SetCriticalRecoveryNil sets the value for CriticalRecovery to be an explicit nil.
func (o *MonitorThresholds) SetCriticalRecoveryNil() {
	o.CriticalRecovery.Set(nil)
}

// UnsetCriticalRecovery ensures that no value is present for CriticalRecovery, not even an explicit nil.
func (o *MonitorThresholds) UnsetCriticalRecovery() {
	o.CriticalRecovery.Unset()
}

// GetCriticalRecoveryQuery returns the CriticalRecoveryQuery field value if set, zero value otherwise.
func (o *MonitorThresholds) GetCriticalRecoveryQuery() string {
	if o == nil || o.CriticalRecoveryQuery == nil {
		var ret string
		return ret
	}
	return *o.CriticalRecoveryQuery
}

// GetCriticalRecoveryQueryOk returns a tuple with the CriticalRecoveryQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetCriticalRecoveryQueryOk() (*string, bool) {
	if o == nil || o.CriticalRecoveryQuery == nil {
		return nil, false
	}
	return o.CriticalRecoveryQuery, true
}

// HasCriticalRecoveryQuery returns a boolean if a field has been set.
func (o *MonitorThresholds) HasCriticalRecoveryQuery() bool {
	return o != nil && o.CriticalRecoveryQuery != nil
}

// SetCriticalRecoveryQuery gets a reference to the given string and assigns it to the CriticalRecoveryQuery field.
func (o *MonitorThresholds) SetCriticalRecoveryQuery(v string) {
	o.CriticalRecoveryQuery = &v
}

// GetOk returns the Ok field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorThresholds) GetOk() float64 {
	if o == nil || o.Ok.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Ok.Get()
}

// GetOkOk returns a tuple with the Ok field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorThresholds) GetOkOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ok.Get(), o.Ok.IsSet()
}

// HasOk returns a boolean if a field has been set.
func (o *MonitorThresholds) HasOk() bool {
	return o != nil && o.Ok.IsSet()
}

// SetOk gets a reference to the given datadog.NullableFloat64 and assigns it to the Ok field.
func (o *MonitorThresholds) SetOk(v float64) {
	o.Ok.Set(&v)
}

// SetOkNil sets the value for Ok to be an explicit nil.
func (o *MonitorThresholds) SetOkNil() {
	o.Ok.Set(nil)
}

// UnsetOk ensures that no value is present for Ok, not even an explicit nil.
func (o *MonitorThresholds) UnsetOk() {
	o.Ok.Unset()
}

// GetUnknown returns the Unknown field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorThresholds) GetUnknown() float64 {
	if o == nil || o.Unknown.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Unknown.Get()
}

// GetUnknownOk returns a tuple with the Unknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorThresholds) GetUnknownOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unknown.Get(), o.Unknown.IsSet()
}

// HasUnknown returns a boolean if a field has been set.
func (o *MonitorThresholds) HasUnknown() bool {
	return o != nil && o.Unknown.IsSet()
}

// SetUnknown gets a reference to the given datadog.NullableFloat64 and assigns it to the Unknown field.
func (o *MonitorThresholds) SetUnknown(v float64) {
	o.Unknown.Set(&v)
}

// SetUnknownNil sets the value for Unknown to be an explicit nil.
func (o *MonitorThresholds) SetUnknownNil() {
	o.Unknown.Set(nil)
}

// UnsetUnknown ensures that no value is present for Unknown, not even an explicit nil.
func (o *MonitorThresholds) UnsetUnknown() {
	o.Unknown.Unset()
}

// GetWarning returns the Warning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorThresholds) GetWarning() float64 {
	if o == nil || o.Warning.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Warning.Get()
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorThresholds) GetWarningOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warning.Get(), o.Warning.IsSet()
}

// HasWarning returns a boolean if a field has been set.
func (o *MonitorThresholds) HasWarning() bool {
	return o != nil && o.Warning.IsSet()
}

// SetWarning gets a reference to the given datadog.NullableFloat64 and assigns it to the Warning field.
func (o *MonitorThresholds) SetWarning(v float64) {
	o.Warning.Set(&v)
}

// SetWarningNil sets the value for Warning to be an explicit nil.
func (o *MonitorThresholds) SetWarningNil() {
	o.Warning.Set(nil)
}

// UnsetWarning ensures that no value is present for Warning, not even an explicit nil.
func (o *MonitorThresholds) UnsetWarning() {
	o.Warning.Unset()
}

// GetWarningRecovery returns the WarningRecovery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorThresholds) GetWarningRecovery() float64 {
	if o == nil || o.WarningRecovery.Get() == nil {
		var ret float64
		return ret
	}
	return *o.WarningRecovery.Get()
}

// GetWarningRecoveryOk returns a tuple with the WarningRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorThresholds) GetWarningRecoveryOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarningRecovery.Get(), o.WarningRecovery.IsSet()
}

// HasWarningRecovery returns a boolean if a field has been set.
func (o *MonitorThresholds) HasWarningRecovery() bool {
	return o != nil && o.WarningRecovery.IsSet()
}

// SetWarningRecovery gets a reference to the given datadog.NullableFloat64 and assigns it to the WarningRecovery field.
func (o *MonitorThresholds) SetWarningRecovery(v float64) {
	o.WarningRecovery.Set(&v)
}

// SetWarningRecoveryNil sets the value for WarningRecovery to be an explicit nil.
func (o *MonitorThresholds) SetWarningRecoveryNil() {
	o.WarningRecovery.Set(nil)
}

// UnsetWarningRecovery ensures that no value is present for WarningRecovery, not even an explicit nil.
func (o *MonitorThresholds) UnsetWarningRecovery() {
	o.WarningRecovery.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorThresholds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Critical != nil {
		toSerialize["critical"] = o.Critical
	}
	if o.CriticalQuery != nil {
		toSerialize["critical_query"] = o.CriticalQuery
	}
	if o.CriticalRecovery.IsSet() {
		toSerialize["critical_recovery"] = o.CriticalRecovery.Get()
	}
	if o.CriticalRecoveryQuery != nil {
		toSerialize["critical_recovery_query"] = o.CriticalRecoveryQuery
	}
	if o.Ok.IsSet() {
		toSerialize["ok"] = o.Ok.Get()
	}
	if o.Unknown.IsSet() {
		toSerialize["unknown"] = o.Unknown.Get()
	}
	if o.Warning.IsSet() {
		toSerialize["warning"] = o.Warning.Get()
	}
	if o.WarningRecovery.IsSet() {
		toSerialize["warning_recovery"] = o.WarningRecovery.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorThresholds) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Critical              *float64                `json:"critical,omitempty"`
		CriticalQuery         *string                 `json:"critical_query,omitempty"`
		CriticalRecovery      datadog.NullableFloat64 `json:"critical_recovery,omitempty"`
		CriticalRecoveryQuery *string                 `json:"critical_recovery_query,omitempty"`
		Ok                    datadog.NullableFloat64 `json:"ok,omitempty"`
		Unknown               datadog.NullableFloat64 `json:"unknown,omitempty"`
		Warning               datadog.NullableFloat64 `json:"warning,omitempty"`
		WarningRecovery       datadog.NullableFloat64 `json:"warning_recovery,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"critical", "critical_query", "critical_recovery", "critical_recovery_query", "ok", "unknown", "warning", "warning_recovery"})
	} else {
		return err
	}
	o.Critical = all.Critical
	o.CriticalQuery = all.CriticalQuery
	o.CriticalRecovery = all.CriticalRecovery
	o.CriticalRecoveryQuery = all.CriticalRecoveryQuery
	o.Ok = all.Ok
	o.Unknown = all.Unknown
	o.Warning = all.Warning
	o.WarningRecovery = all.WarningRecovery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
