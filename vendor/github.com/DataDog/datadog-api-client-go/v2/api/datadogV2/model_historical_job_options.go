// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HistoricalJobOptions Job options.
type HistoricalJobOptions struct {
	// The detection method.
	DetectionMethod *SecurityMonitoringRuleDetectionMethod `json:"detectionMethod,omitempty"`
	// A time window is specified to match when at least one of the cases matches true. This is a sliding window
	// and evaluates in real time. For third party detection method, this field is not used.
	EvaluationWindow *SecurityMonitoringRuleEvaluationWindow `json:"evaluationWindow,omitempty"`
	// Options on impossible travel detection method.
	ImpossibleTravelOptions *SecurityMonitoringRuleImpossibleTravelOptions `json:"impossibleTravelOptions,omitempty"`
	// Once a signal is generated, the signal will remain "open" if a case is matched at least once within
	// this keep alive window. For third party detection method, this field is not used.
	KeepAlive *SecurityMonitoringRuleKeepAlive `json:"keepAlive,omitempty"`
	// A signal will "close" regardless of the query being matched once the time exceeds the maximum duration.
	// This time is calculated from the first seen timestamp.
	MaxSignalDuration *SecurityMonitoringRuleMaxSignalDuration `json:"maxSignalDuration,omitempty"`
	// Options on new value detection method.
	NewValueOptions *SecurityMonitoringRuleNewValueOptions `json:"newValueOptions,omitempty"`
	// Options on third party detection method.
	ThirdPartyRuleOptions *SecurityMonitoringRuleThirdPartyOptions `json:"thirdPartyRuleOptions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHistoricalJobOptions instantiates a new HistoricalJobOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHistoricalJobOptions() *HistoricalJobOptions {
	this := HistoricalJobOptions{}
	return &this
}

// NewHistoricalJobOptionsWithDefaults instantiates a new HistoricalJobOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHistoricalJobOptionsWithDefaults() *HistoricalJobOptions {
	this := HistoricalJobOptions{}
	return &this
}

// GetDetectionMethod returns the DetectionMethod field value if set, zero value otherwise.
func (o *HistoricalJobOptions) GetDetectionMethod() SecurityMonitoringRuleDetectionMethod {
	if o == nil || o.DetectionMethod == nil {
		var ret SecurityMonitoringRuleDetectionMethod
		return ret
	}
	return *o.DetectionMethod
}

// GetDetectionMethodOk returns a tuple with the DetectionMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobOptions) GetDetectionMethodOk() (*SecurityMonitoringRuleDetectionMethod, bool) {
	if o == nil || o.DetectionMethod == nil {
		return nil, false
	}
	return o.DetectionMethod, true
}

// HasDetectionMethod returns a boolean if a field has been set.
func (o *HistoricalJobOptions) HasDetectionMethod() bool {
	return o != nil && o.DetectionMethod != nil
}

// SetDetectionMethod gets a reference to the given SecurityMonitoringRuleDetectionMethod and assigns it to the DetectionMethod field.
func (o *HistoricalJobOptions) SetDetectionMethod(v SecurityMonitoringRuleDetectionMethod) {
	o.DetectionMethod = &v
}

// GetEvaluationWindow returns the EvaluationWindow field value if set, zero value otherwise.
func (o *HistoricalJobOptions) GetEvaluationWindow() SecurityMonitoringRuleEvaluationWindow {
	if o == nil || o.EvaluationWindow == nil {
		var ret SecurityMonitoringRuleEvaluationWindow
		return ret
	}
	return *o.EvaluationWindow
}

// GetEvaluationWindowOk returns a tuple with the EvaluationWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobOptions) GetEvaluationWindowOk() (*SecurityMonitoringRuleEvaluationWindow, bool) {
	if o == nil || o.EvaluationWindow == nil {
		return nil, false
	}
	return o.EvaluationWindow, true
}

// HasEvaluationWindow returns a boolean if a field has been set.
func (o *HistoricalJobOptions) HasEvaluationWindow() bool {
	return o != nil && o.EvaluationWindow != nil
}

// SetEvaluationWindow gets a reference to the given SecurityMonitoringRuleEvaluationWindow and assigns it to the EvaluationWindow field.
func (o *HistoricalJobOptions) SetEvaluationWindow(v SecurityMonitoringRuleEvaluationWindow) {
	o.EvaluationWindow = &v
}

// GetImpossibleTravelOptions returns the ImpossibleTravelOptions field value if set, zero value otherwise.
func (o *HistoricalJobOptions) GetImpossibleTravelOptions() SecurityMonitoringRuleImpossibleTravelOptions {
	if o == nil || o.ImpossibleTravelOptions == nil {
		var ret SecurityMonitoringRuleImpossibleTravelOptions
		return ret
	}
	return *o.ImpossibleTravelOptions
}

// GetImpossibleTravelOptionsOk returns a tuple with the ImpossibleTravelOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobOptions) GetImpossibleTravelOptionsOk() (*SecurityMonitoringRuleImpossibleTravelOptions, bool) {
	if o == nil || o.ImpossibleTravelOptions == nil {
		return nil, false
	}
	return o.ImpossibleTravelOptions, true
}

// HasImpossibleTravelOptions returns a boolean if a field has been set.
func (o *HistoricalJobOptions) HasImpossibleTravelOptions() bool {
	return o != nil && o.ImpossibleTravelOptions != nil
}

// SetImpossibleTravelOptions gets a reference to the given SecurityMonitoringRuleImpossibleTravelOptions and assigns it to the ImpossibleTravelOptions field.
func (o *HistoricalJobOptions) SetImpossibleTravelOptions(v SecurityMonitoringRuleImpossibleTravelOptions) {
	o.ImpossibleTravelOptions = &v
}

// GetKeepAlive returns the KeepAlive field value if set, zero value otherwise.
func (o *HistoricalJobOptions) GetKeepAlive() SecurityMonitoringRuleKeepAlive {
	if o == nil || o.KeepAlive == nil {
		var ret SecurityMonitoringRuleKeepAlive
		return ret
	}
	return *o.KeepAlive
}

// GetKeepAliveOk returns a tuple with the KeepAlive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobOptions) GetKeepAliveOk() (*SecurityMonitoringRuleKeepAlive, bool) {
	if o == nil || o.KeepAlive == nil {
		return nil, false
	}
	return o.KeepAlive, true
}

// HasKeepAlive returns a boolean if a field has been set.
func (o *HistoricalJobOptions) HasKeepAlive() bool {
	return o != nil && o.KeepAlive != nil
}

// SetKeepAlive gets a reference to the given SecurityMonitoringRuleKeepAlive and assigns it to the KeepAlive field.
func (o *HistoricalJobOptions) SetKeepAlive(v SecurityMonitoringRuleKeepAlive) {
	o.KeepAlive = &v
}

// GetMaxSignalDuration returns the MaxSignalDuration field value if set, zero value otherwise.
func (o *HistoricalJobOptions) GetMaxSignalDuration() SecurityMonitoringRuleMaxSignalDuration {
	if o == nil || o.MaxSignalDuration == nil {
		var ret SecurityMonitoringRuleMaxSignalDuration
		return ret
	}
	return *o.MaxSignalDuration
}

// GetMaxSignalDurationOk returns a tuple with the MaxSignalDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobOptions) GetMaxSignalDurationOk() (*SecurityMonitoringRuleMaxSignalDuration, bool) {
	if o == nil || o.MaxSignalDuration == nil {
		return nil, false
	}
	return o.MaxSignalDuration, true
}

// HasMaxSignalDuration returns a boolean if a field has been set.
func (o *HistoricalJobOptions) HasMaxSignalDuration() bool {
	return o != nil && o.MaxSignalDuration != nil
}

// SetMaxSignalDuration gets a reference to the given SecurityMonitoringRuleMaxSignalDuration and assigns it to the MaxSignalDuration field.
func (o *HistoricalJobOptions) SetMaxSignalDuration(v SecurityMonitoringRuleMaxSignalDuration) {
	o.MaxSignalDuration = &v
}

// GetNewValueOptions returns the NewValueOptions field value if set, zero value otherwise.
func (o *HistoricalJobOptions) GetNewValueOptions() SecurityMonitoringRuleNewValueOptions {
	if o == nil || o.NewValueOptions == nil {
		var ret SecurityMonitoringRuleNewValueOptions
		return ret
	}
	return *o.NewValueOptions
}

// GetNewValueOptionsOk returns a tuple with the NewValueOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobOptions) GetNewValueOptionsOk() (*SecurityMonitoringRuleNewValueOptions, bool) {
	if o == nil || o.NewValueOptions == nil {
		return nil, false
	}
	return o.NewValueOptions, true
}

// HasNewValueOptions returns a boolean if a field has been set.
func (o *HistoricalJobOptions) HasNewValueOptions() bool {
	return o != nil && o.NewValueOptions != nil
}

// SetNewValueOptions gets a reference to the given SecurityMonitoringRuleNewValueOptions and assigns it to the NewValueOptions field.
func (o *HistoricalJobOptions) SetNewValueOptions(v SecurityMonitoringRuleNewValueOptions) {
	o.NewValueOptions = &v
}

// GetThirdPartyRuleOptions returns the ThirdPartyRuleOptions field value if set, zero value otherwise.
func (o *HistoricalJobOptions) GetThirdPartyRuleOptions() SecurityMonitoringRuleThirdPartyOptions {
	if o == nil || o.ThirdPartyRuleOptions == nil {
		var ret SecurityMonitoringRuleThirdPartyOptions
		return ret
	}
	return *o.ThirdPartyRuleOptions
}

// GetThirdPartyRuleOptionsOk returns a tuple with the ThirdPartyRuleOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobOptions) GetThirdPartyRuleOptionsOk() (*SecurityMonitoringRuleThirdPartyOptions, bool) {
	if o == nil || o.ThirdPartyRuleOptions == nil {
		return nil, false
	}
	return o.ThirdPartyRuleOptions, true
}

// HasThirdPartyRuleOptions returns a boolean if a field has been set.
func (o *HistoricalJobOptions) HasThirdPartyRuleOptions() bool {
	return o != nil && o.ThirdPartyRuleOptions != nil
}

// SetThirdPartyRuleOptions gets a reference to the given SecurityMonitoringRuleThirdPartyOptions and assigns it to the ThirdPartyRuleOptions field.
func (o *HistoricalJobOptions) SetThirdPartyRuleOptions(v SecurityMonitoringRuleThirdPartyOptions) {
	o.ThirdPartyRuleOptions = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HistoricalJobOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DetectionMethod != nil {
		toSerialize["detectionMethod"] = o.DetectionMethod
	}
	if o.EvaluationWindow != nil {
		toSerialize["evaluationWindow"] = o.EvaluationWindow
	}
	if o.ImpossibleTravelOptions != nil {
		toSerialize["impossibleTravelOptions"] = o.ImpossibleTravelOptions
	}
	if o.KeepAlive != nil {
		toSerialize["keepAlive"] = o.KeepAlive
	}
	if o.MaxSignalDuration != nil {
		toSerialize["maxSignalDuration"] = o.MaxSignalDuration
	}
	if o.NewValueOptions != nil {
		toSerialize["newValueOptions"] = o.NewValueOptions
	}
	if o.ThirdPartyRuleOptions != nil {
		toSerialize["thirdPartyRuleOptions"] = o.ThirdPartyRuleOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HistoricalJobOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DetectionMethod         *SecurityMonitoringRuleDetectionMethod         `json:"detectionMethod,omitempty"`
		EvaluationWindow        *SecurityMonitoringRuleEvaluationWindow        `json:"evaluationWindow,omitempty"`
		ImpossibleTravelOptions *SecurityMonitoringRuleImpossibleTravelOptions `json:"impossibleTravelOptions,omitempty"`
		KeepAlive               *SecurityMonitoringRuleKeepAlive               `json:"keepAlive,omitempty"`
		MaxSignalDuration       *SecurityMonitoringRuleMaxSignalDuration       `json:"maxSignalDuration,omitempty"`
		NewValueOptions         *SecurityMonitoringRuleNewValueOptions         `json:"newValueOptions,omitempty"`
		ThirdPartyRuleOptions   *SecurityMonitoringRuleThirdPartyOptions       `json:"thirdPartyRuleOptions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"detectionMethod", "evaluationWindow", "impossibleTravelOptions", "keepAlive", "maxSignalDuration", "newValueOptions", "thirdPartyRuleOptions"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DetectionMethod != nil && !all.DetectionMethod.IsValid() {
		hasInvalidField = true
	} else {
		o.DetectionMethod = all.DetectionMethod
	}
	if all.EvaluationWindow != nil && !all.EvaluationWindow.IsValid() {
		hasInvalidField = true
	} else {
		o.EvaluationWindow = all.EvaluationWindow
	}
	if all.ImpossibleTravelOptions != nil && all.ImpossibleTravelOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ImpossibleTravelOptions = all.ImpossibleTravelOptions
	if all.KeepAlive != nil && !all.KeepAlive.IsValid() {
		hasInvalidField = true
	} else {
		o.KeepAlive = all.KeepAlive
	}
	if all.MaxSignalDuration != nil && !all.MaxSignalDuration.IsValid() {
		hasInvalidField = true
	} else {
		o.MaxSignalDuration = all.MaxSignalDuration
	}
	if all.NewValueOptions != nil && all.NewValueOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NewValueOptions = all.NewValueOptions
	if all.ThirdPartyRuleOptions != nil && all.ThirdPartyRuleOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ThirdPartyRuleOptions = all.ThirdPartyRuleOptions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
