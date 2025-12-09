// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleSequenceDetectionStepTransition Transition from a parent step to a child step within a sequence detection rule.
type SecurityMonitoringRuleSequenceDetectionStepTransition struct {
	// Name of the child step.
	Child *string `json:"child,omitempty"`
	// A time window is specified to match when at least one of the cases matches true. This is a sliding window
	// and evaluates in real time. For third party detection method, this field is not used.
	EvaluationWindow *SecurityMonitoringRuleEvaluationWindow `json:"evaluationWindow,omitempty"`
	// Name of the parent step.
	Parent *string `json:"parent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleSequenceDetectionStepTransition instantiates a new SecurityMonitoringRuleSequenceDetectionStepTransition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleSequenceDetectionStepTransition() *SecurityMonitoringRuleSequenceDetectionStepTransition {
	this := SecurityMonitoringRuleSequenceDetectionStepTransition{}
	return &this
}

// NewSecurityMonitoringRuleSequenceDetectionStepTransitionWithDefaults instantiates a new SecurityMonitoringRuleSequenceDetectionStepTransition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleSequenceDetectionStepTransitionWithDefaults() *SecurityMonitoringRuleSequenceDetectionStepTransition {
	this := SecurityMonitoringRuleSequenceDetectionStepTransition{}
	return &this
}

// GetChild returns the Child field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) GetChild() string {
	if o == nil || o.Child == nil {
		var ret string
		return ret
	}
	return *o.Child
}

// GetChildOk returns a tuple with the Child field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) GetChildOk() (*string, bool) {
	if o == nil || o.Child == nil {
		return nil, false
	}
	return o.Child, true
}

// HasChild returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) HasChild() bool {
	return o != nil && o.Child != nil
}

// SetChild gets a reference to the given string and assigns it to the Child field.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) SetChild(v string) {
	o.Child = &v
}

// GetEvaluationWindow returns the EvaluationWindow field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) GetEvaluationWindow() SecurityMonitoringRuleEvaluationWindow {
	if o == nil || o.EvaluationWindow == nil {
		var ret SecurityMonitoringRuleEvaluationWindow
		return ret
	}
	return *o.EvaluationWindow
}

// GetEvaluationWindowOk returns a tuple with the EvaluationWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) GetEvaluationWindowOk() (*SecurityMonitoringRuleEvaluationWindow, bool) {
	if o == nil || o.EvaluationWindow == nil {
		return nil, false
	}
	return o.EvaluationWindow, true
}

// HasEvaluationWindow returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) HasEvaluationWindow() bool {
	return o != nil && o.EvaluationWindow != nil
}

// SetEvaluationWindow gets a reference to the given SecurityMonitoringRuleEvaluationWindow and assigns it to the EvaluationWindow field.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) SetEvaluationWindow(v SecurityMonitoringRuleEvaluationWindow) {
	o.EvaluationWindow = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) GetParent() string {
	if o == nil || o.Parent == nil {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) GetParentOk() (*string, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) HasParent() bool {
	return o != nil && o.Parent != nil
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) SetParent(v string) {
	o.Parent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleSequenceDetectionStepTransition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Child != nil {
		toSerialize["child"] = o.Child
	}
	if o.EvaluationWindow != nil {
		toSerialize["evaluationWindow"] = o.EvaluationWindow
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleSequenceDetectionStepTransition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Child            *string                                 `json:"child,omitempty"`
		EvaluationWindow *SecurityMonitoringRuleEvaluationWindow `json:"evaluationWindow,omitempty"`
		Parent           *string                                 `json:"parent,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"child", "evaluationWindow", "parent"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Child = all.Child
	if all.EvaluationWindow != nil && !all.EvaluationWindow.IsValid() {
		hasInvalidField = true
	} else {
		o.EvaluationWindow = all.EvaluationWindow
	}
	o.Parent = all.Parent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
