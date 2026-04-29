// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultExecutionInfo Execution details for a Synthetic test result.
type SyntheticsTestResultExecutionInfo struct {
	// Total duration of a Synthetic test execution.
	Duration *SyntheticsTestResultDuration `json:"duration,omitempty"`
	// Error message if the execution encountered an issue.
	ErrorMessage *string `json:"error_message,omitempty"`
	// Whether this result is from a fast retry.
	IsFastRetry *bool `json:"is_fast_retry,omitempty"`
	// Timing breakdown of the test execution in milliseconds.
	Timings map[string]interface{} `json:"timings,omitempty"`
	// Whether the test was executed through a tunnel.
	Tunnel *bool `json:"tunnel,omitempty"`
	// Whether the location was unhealthy during execution.
	Unhealthy *bool `json:"unhealthy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultExecutionInfo instantiates a new SyntheticsTestResultExecutionInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultExecutionInfo() *SyntheticsTestResultExecutionInfo {
	this := SyntheticsTestResultExecutionInfo{}
	return &this
}

// NewSyntheticsTestResultExecutionInfoWithDefaults instantiates a new SyntheticsTestResultExecutionInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultExecutionInfoWithDefaults() *SyntheticsTestResultExecutionInfo {
	this := SyntheticsTestResultExecutionInfo{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SyntheticsTestResultExecutionInfo) GetDuration() SyntheticsTestResultDuration {
	if o == nil || o.Duration == nil {
		var ret SyntheticsTestResultDuration
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultExecutionInfo) GetDurationOk() (*SyntheticsTestResultDuration, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SyntheticsTestResultExecutionInfo) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given SyntheticsTestResultDuration and assigns it to the Duration field.
func (o *SyntheticsTestResultExecutionInfo) SetDuration(v SyntheticsTestResultDuration) {
	o.Duration = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *SyntheticsTestResultExecutionInfo) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultExecutionInfo) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultExecutionInfo) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *SyntheticsTestResultExecutionInfo) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetIsFastRetry returns the IsFastRetry field value if set, zero value otherwise.
func (o *SyntheticsTestResultExecutionInfo) GetIsFastRetry() bool {
	if o == nil || o.IsFastRetry == nil {
		var ret bool
		return ret
	}
	return *o.IsFastRetry
}

// GetIsFastRetryOk returns a tuple with the IsFastRetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultExecutionInfo) GetIsFastRetryOk() (*bool, bool) {
	if o == nil || o.IsFastRetry == nil {
		return nil, false
	}
	return o.IsFastRetry, true
}

// HasIsFastRetry returns a boolean if a field has been set.
func (o *SyntheticsTestResultExecutionInfo) HasIsFastRetry() bool {
	return o != nil && o.IsFastRetry != nil
}

// SetIsFastRetry gets a reference to the given bool and assigns it to the IsFastRetry field.
func (o *SyntheticsTestResultExecutionInfo) SetIsFastRetry(v bool) {
	o.IsFastRetry = &v
}

// GetTimings returns the Timings field value if set, zero value otherwise.
func (o *SyntheticsTestResultExecutionInfo) GetTimings() map[string]interface{} {
	if o == nil || o.Timings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Timings
}

// GetTimingsOk returns a tuple with the Timings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultExecutionInfo) GetTimingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Timings == nil {
		return nil, false
	}
	return &o.Timings, true
}

// HasTimings returns a boolean if a field has been set.
func (o *SyntheticsTestResultExecutionInfo) HasTimings() bool {
	return o != nil && o.Timings != nil
}

// SetTimings gets a reference to the given map[string]interface{} and assigns it to the Timings field.
func (o *SyntheticsTestResultExecutionInfo) SetTimings(v map[string]interface{}) {
	o.Timings = v
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise.
func (o *SyntheticsTestResultExecutionInfo) GetTunnel() bool {
	if o == nil || o.Tunnel == nil {
		var ret bool
		return ret
	}
	return *o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultExecutionInfo) GetTunnelOk() (*bool, bool) {
	if o == nil || o.Tunnel == nil {
		return nil, false
	}
	return o.Tunnel, true
}

// HasTunnel returns a boolean if a field has been set.
func (o *SyntheticsTestResultExecutionInfo) HasTunnel() bool {
	return o != nil && o.Tunnel != nil
}

// SetTunnel gets a reference to the given bool and assigns it to the Tunnel field.
func (o *SyntheticsTestResultExecutionInfo) SetTunnel(v bool) {
	o.Tunnel = &v
}

// GetUnhealthy returns the Unhealthy field value if set, zero value otherwise.
func (o *SyntheticsTestResultExecutionInfo) GetUnhealthy() bool {
	if o == nil || o.Unhealthy == nil {
		var ret bool
		return ret
	}
	return *o.Unhealthy
}

// GetUnhealthyOk returns a tuple with the Unhealthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultExecutionInfo) GetUnhealthyOk() (*bool, bool) {
	if o == nil || o.Unhealthy == nil {
		return nil, false
	}
	return o.Unhealthy, true
}

// HasUnhealthy returns a boolean if a field has been set.
func (o *SyntheticsTestResultExecutionInfo) HasUnhealthy() bool {
	return o != nil && o.Unhealthy != nil
}

// SetUnhealthy gets a reference to the given bool and assigns it to the Unhealthy field.
func (o *SyntheticsTestResultExecutionInfo) SetUnhealthy(v bool) {
	o.Unhealthy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultExecutionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.IsFastRetry != nil {
		toSerialize["is_fast_retry"] = o.IsFastRetry
	}
	if o.Timings != nil {
		toSerialize["timings"] = o.Timings
	}
	if o.Tunnel != nil {
		toSerialize["tunnel"] = o.Tunnel
	}
	if o.Unhealthy != nil {
		toSerialize["unhealthy"] = o.Unhealthy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultExecutionInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration     *SyntheticsTestResultDuration `json:"duration,omitempty"`
		ErrorMessage *string                       `json:"error_message,omitempty"`
		IsFastRetry  *bool                         `json:"is_fast_retry,omitempty"`
		Timings      map[string]interface{}        `json:"timings,omitempty"`
		Tunnel       *bool                         `json:"tunnel,omitempty"`
		Unhealthy    *bool                         `json:"unhealthy,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration", "error_message", "is_fast_retry", "timings", "tunnel", "unhealthy"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Duration != nil && all.Duration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Duration = all.Duration
	o.ErrorMessage = all.ErrorMessage
	o.IsFastRetry = all.IsFastRetry
	o.Timings = all.Timings
	o.Tunnel = all.Tunnel
	o.Unhealthy = all.Unhealthy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
