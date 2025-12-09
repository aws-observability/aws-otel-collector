// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleSequenceDetectionStep Step definition for sequence detection containing the step name, condition, and evaluation window.
type SecurityMonitoringRuleSequenceDetectionStep struct {
	// Condition referencing rule queries (e.g., `a > 0`).
	Condition *string `json:"condition,omitempty"`
	// A time window is specified to match when at least one of the cases matches true. This is a sliding window
	// and evaluates in real time. For third party detection method, this field is not used.
	EvaluationWindow *SecurityMonitoringRuleEvaluationWindow `json:"evaluationWindow,omitempty"`
	// Unique name identifying the step.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleSequenceDetectionStep instantiates a new SecurityMonitoringRuleSequenceDetectionStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleSequenceDetectionStep() *SecurityMonitoringRuleSequenceDetectionStep {
	this := SecurityMonitoringRuleSequenceDetectionStep{}
	return &this
}

// NewSecurityMonitoringRuleSequenceDetectionStepWithDefaults instantiates a new SecurityMonitoringRuleSequenceDetectionStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleSequenceDetectionStepWithDefaults() *SecurityMonitoringRuleSequenceDetectionStep {
	this := SecurityMonitoringRuleSequenceDetectionStep{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleSequenceDetectionStep) GetCondition() string {
	if o == nil || o.Condition == nil {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleSequenceDetectionStep) GetConditionOk() (*string, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleSequenceDetectionStep) HasCondition() bool {
	return o != nil && o.Condition != nil
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *SecurityMonitoringRuleSequenceDetectionStep) SetCondition(v string) {
	o.Condition = &v
}

// GetEvaluationWindow returns the EvaluationWindow field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleSequenceDetectionStep) GetEvaluationWindow() SecurityMonitoringRuleEvaluationWindow {
	if o == nil || o.EvaluationWindow == nil {
		var ret SecurityMonitoringRuleEvaluationWindow
		return ret
	}
	return *o.EvaluationWindow
}

// GetEvaluationWindowOk returns a tuple with the EvaluationWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleSequenceDetectionStep) GetEvaluationWindowOk() (*SecurityMonitoringRuleEvaluationWindow, bool) {
	if o == nil || o.EvaluationWindow == nil {
		return nil, false
	}
	return o.EvaluationWindow, true
}

// HasEvaluationWindow returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleSequenceDetectionStep) HasEvaluationWindow() bool {
	return o != nil && o.EvaluationWindow != nil
}

// SetEvaluationWindow gets a reference to the given SecurityMonitoringRuleEvaluationWindow and assigns it to the EvaluationWindow field.
func (o *SecurityMonitoringRuleSequenceDetectionStep) SetEvaluationWindow(v SecurityMonitoringRuleEvaluationWindow) {
	o.EvaluationWindow = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleSequenceDetectionStep) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleSequenceDetectionStep) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleSequenceDetectionStep) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringRuleSequenceDetectionStep) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleSequenceDetectionStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.EvaluationWindow != nil {
		toSerialize["evaluationWindow"] = o.EvaluationWindow
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleSequenceDetectionStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Condition        *string                                 `json:"condition,omitempty"`
		EvaluationWindow *SecurityMonitoringRuleEvaluationWindow `json:"evaluationWindow,omitempty"`
		Name             *string                                 `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"condition", "evaluationWindow", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Condition = all.Condition
	if all.EvaluationWindow != nil && !all.EvaluationWindow.IsValid() {
		hasInvalidField = true
	} else {
		o.EvaluationWindow = all.EvaluationWindow
	}
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
