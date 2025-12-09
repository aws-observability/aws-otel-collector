// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleSequenceDetectionOptions Options on sequence detection method.
type SecurityMonitoringRuleSequenceDetectionOptions struct {
	// Transitions defining the allowed order of steps and their evaluation windows.
	StepTransitions []SecurityMonitoringRuleSequenceDetectionStepTransition `json:"stepTransitions,omitempty"`
	// Steps that define the conditions to be matched in sequence.
	Steps []SecurityMonitoringRuleSequenceDetectionStep `json:"steps,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleSequenceDetectionOptions instantiates a new SecurityMonitoringRuleSequenceDetectionOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleSequenceDetectionOptions() *SecurityMonitoringRuleSequenceDetectionOptions {
	this := SecurityMonitoringRuleSequenceDetectionOptions{}
	return &this
}

// NewSecurityMonitoringRuleSequenceDetectionOptionsWithDefaults instantiates a new SecurityMonitoringRuleSequenceDetectionOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleSequenceDetectionOptionsWithDefaults() *SecurityMonitoringRuleSequenceDetectionOptions {
	this := SecurityMonitoringRuleSequenceDetectionOptions{}
	return &this
}

// GetStepTransitions returns the StepTransitions field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleSequenceDetectionOptions) GetStepTransitions() []SecurityMonitoringRuleSequenceDetectionStepTransition {
	if o == nil || o.StepTransitions == nil {
		var ret []SecurityMonitoringRuleSequenceDetectionStepTransition
		return ret
	}
	return o.StepTransitions
}

// GetStepTransitionsOk returns a tuple with the StepTransitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleSequenceDetectionOptions) GetStepTransitionsOk() (*[]SecurityMonitoringRuleSequenceDetectionStepTransition, bool) {
	if o == nil || o.StepTransitions == nil {
		return nil, false
	}
	return &o.StepTransitions, true
}

// HasStepTransitions returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleSequenceDetectionOptions) HasStepTransitions() bool {
	return o != nil && o.StepTransitions != nil
}

// SetStepTransitions gets a reference to the given []SecurityMonitoringRuleSequenceDetectionStepTransition and assigns it to the StepTransitions field.
func (o *SecurityMonitoringRuleSequenceDetectionOptions) SetStepTransitions(v []SecurityMonitoringRuleSequenceDetectionStepTransition) {
	o.StepTransitions = v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleSequenceDetectionOptions) GetSteps() []SecurityMonitoringRuleSequenceDetectionStep {
	if o == nil || o.Steps == nil {
		var ret []SecurityMonitoringRuleSequenceDetectionStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleSequenceDetectionOptions) GetStepsOk() (*[]SecurityMonitoringRuleSequenceDetectionStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleSequenceDetectionOptions) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []SecurityMonitoringRuleSequenceDetectionStep and assigns it to the Steps field.
func (o *SecurityMonitoringRuleSequenceDetectionOptions) SetSteps(v []SecurityMonitoringRuleSequenceDetectionStep) {
	o.Steps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleSequenceDetectionOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.StepTransitions != nil {
		toSerialize["stepTransitions"] = o.StepTransitions
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleSequenceDetectionOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StepTransitions []SecurityMonitoringRuleSequenceDetectionStepTransition `json:"stepTransitions,omitempty"`
		Steps           []SecurityMonitoringRuleSequenceDetectionStep           `json:"steps,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"stepTransitions", "steps"})
	} else {
		return err
	}
	o.StepTransitions = all.StepTransitions
	o.Steps = all.Steps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
