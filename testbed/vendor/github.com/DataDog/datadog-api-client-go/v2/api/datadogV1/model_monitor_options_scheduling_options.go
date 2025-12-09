// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorOptionsSchedulingOptions Configuration options for scheduling.
type MonitorOptionsSchedulingOptions struct {
	// Configuration options for the custom schedule. **This feature is in private beta.**
	CustomSchedule *MonitorOptionsCustomSchedule `json:"custom_schedule,omitempty"`
	// Configuration options for the evaluation window. If `hour_starts` is set, no other fields may be set. Otherwise, `day_starts` and `month_starts` must be set together.
	EvaluationWindow *MonitorOptionsSchedulingOptionsEvaluationWindow `json:"evaluation_window,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorOptionsSchedulingOptions instantiates a new MonitorOptionsSchedulingOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorOptionsSchedulingOptions() *MonitorOptionsSchedulingOptions {
	this := MonitorOptionsSchedulingOptions{}
	return &this
}

// NewMonitorOptionsSchedulingOptionsWithDefaults instantiates a new MonitorOptionsSchedulingOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorOptionsSchedulingOptionsWithDefaults() *MonitorOptionsSchedulingOptions {
	this := MonitorOptionsSchedulingOptions{}
	return &this
}

// GetCustomSchedule returns the CustomSchedule field value if set, zero value otherwise.
func (o *MonitorOptionsSchedulingOptions) GetCustomSchedule() MonitorOptionsCustomSchedule {
	if o == nil || o.CustomSchedule == nil {
		var ret MonitorOptionsCustomSchedule
		return ret
	}
	return *o.CustomSchedule
}

// GetCustomScheduleOk returns a tuple with the CustomSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorOptionsSchedulingOptions) GetCustomScheduleOk() (*MonitorOptionsCustomSchedule, bool) {
	if o == nil || o.CustomSchedule == nil {
		return nil, false
	}
	return o.CustomSchedule, true
}

// HasCustomSchedule returns a boolean if a field has been set.
func (o *MonitorOptionsSchedulingOptions) HasCustomSchedule() bool {
	return o != nil && o.CustomSchedule != nil
}

// SetCustomSchedule gets a reference to the given MonitorOptionsCustomSchedule and assigns it to the CustomSchedule field.
func (o *MonitorOptionsSchedulingOptions) SetCustomSchedule(v MonitorOptionsCustomSchedule) {
	o.CustomSchedule = &v
}

// GetEvaluationWindow returns the EvaluationWindow field value if set, zero value otherwise.
func (o *MonitorOptionsSchedulingOptions) GetEvaluationWindow() MonitorOptionsSchedulingOptionsEvaluationWindow {
	if o == nil || o.EvaluationWindow == nil {
		var ret MonitorOptionsSchedulingOptionsEvaluationWindow
		return ret
	}
	return *o.EvaluationWindow
}

// GetEvaluationWindowOk returns a tuple with the EvaluationWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorOptionsSchedulingOptions) GetEvaluationWindowOk() (*MonitorOptionsSchedulingOptionsEvaluationWindow, bool) {
	if o == nil || o.EvaluationWindow == nil {
		return nil, false
	}
	return o.EvaluationWindow, true
}

// HasEvaluationWindow returns a boolean if a field has been set.
func (o *MonitorOptionsSchedulingOptions) HasEvaluationWindow() bool {
	return o != nil && o.EvaluationWindow != nil
}

// SetEvaluationWindow gets a reference to the given MonitorOptionsSchedulingOptionsEvaluationWindow and assigns it to the EvaluationWindow field.
func (o *MonitorOptionsSchedulingOptions) SetEvaluationWindow(v MonitorOptionsSchedulingOptionsEvaluationWindow) {
	o.EvaluationWindow = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorOptionsSchedulingOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomSchedule != nil {
		toSerialize["custom_schedule"] = o.CustomSchedule
	}
	if o.EvaluationWindow != nil {
		toSerialize["evaluation_window"] = o.EvaluationWindow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorOptionsSchedulingOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomSchedule   *MonitorOptionsCustomSchedule                    `json:"custom_schedule,omitempty"`
		EvaluationWindow *MonitorOptionsSchedulingOptionsEvaluationWindow `json:"evaluation_window,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_schedule", "evaluation_window"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CustomSchedule != nil && all.CustomSchedule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CustomSchedule = all.CustomSchedule
	if all.EvaluationWindow != nil && all.EvaluationWindow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EvaluationWindow = all.EvaluationWindow

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
