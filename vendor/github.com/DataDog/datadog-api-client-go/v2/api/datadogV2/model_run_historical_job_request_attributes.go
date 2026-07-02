// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunHistoricalJobRequestAttributes Run a historical job request.
type RunHistoricalJobRequestAttributes struct {
	// Definition of a historical job based on a security monitoring rule.
	FromRule *JobDefinitionFromRule `json:"fromRule,omitempty"`
	// Definition of a historical job.
	JobDefinition *JobDefinition `json:"jobDefinition,omitempty"`
	// Whether the job outputs signals when results are converted.
	SignalOutput *bool `json:"signalOutput,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRunHistoricalJobRequestAttributes instantiates a new RunHistoricalJobRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRunHistoricalJobRequestAttributes() *RunHistoricalJobRequestAttributes {
	this := RunHistoricalJobRequestAttributes{}
	return &this
}

// NewRunHistoricalJobRequestAttributesWithDefaults instantiates a new RunHistoricalJobRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRunHistoricalJobRequestAttributesWithDefaults() *RunHistoricalJobRequestAttributes {
	this := RunHistoricalJobRequestAttributes{}
	return &this
}

// GetFromRule returns the FromRule field value if set, zero value otherwise.
func (o *RunHistoricalJobRequestAttributes) GetFromRule() JobDefinitionFromRule {
	if o == nil || o.FromRule == nil {
		var ret JobDefinitionFromRule
		return ret
	}
	return *o.FromRule
}

// GetFromRuleOk returns a tuple with the FromRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunHistoricalJobRequestAttributes) GetFromRuleOk() (*JobDefinitionFromRule, bool) {
	if o == nil || o.FromRule == nil {
		return nil, false
	}
	return o.FromRule, true
}

// HasFromRule returns a boolean if a field has been set.
func (o *RunHistoricalJobRequestAttributes) HasFromRule() bool {
	return o != nil && o.FromRule != nil
}

// SetFromRule gets a reference to the given JobDefinitionFromRule and assigns it to the FromRule field.
func (o *RunHistoricalJobRequestAttributes) SetFromRule(v JobDefinitionFromRule) {
	o.FromRule = &v
}

// GetJobDefinition returns the JobDefinition field value if set, zero value otherwise.
func (o *RunHistoricalJobRequestAttributes) GetJobDefinition() JobDefinition {
	if o == nil || o.JobDefinition == nil {
		var ret JobDefinition
		return ret
	}
	return *o.JobDefinition
}

// GetJobDefinitionOk returns a tuple with the JobDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunHistoricalJobRequestAttributes) GetJobDefinitionOk() (*JobDefinition, bool) {
	if o == nil || o.JobDefinition == nil {
		return nil, false
	}
	return o.JobDefinition, true
}

// HasJobDefinition returns a boolean if a field has been set.
func (o *RunHistoricalJobRequestAttributes) HasJobDefinition() bool {
	return o != nil && o.JobDefinition != nil
}

// SetJobDefinition gets a reference to the given JobDefinition and assigns it to the JobDefinition field.
func (o *RunHistoricalJobRequestAttributes) SetJobDefinition(v JobDefinition) {
	o.JobDefinition = &v
}

// GetSignalOutput returns the SignalOutput field value if set, zero value otherwise.
func (o *RunHistoricalJobRequestAttributes) GetSignalOutput() bool {
	if o == nil || o.SignalOutput == nil {
		var ret bool
		return ret
	}
	return *o.SignalOutput
}

// GetSignalOutputOk returns a tuple with the SignalOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunHistoricalJobRequestAttributes) GetSignalOutputOk() (*bool, bool) {
	if o == nil || o.SignalOutput == nil {
		return nil, false
	}
	return o.SignalOutput, true
}

// HasSignalOutput returns a boolean if a field has been set.
func (o *RunHistoricalJobRequestAttributes) HasSignalOutput() bool {
	return o != nil && o.SignalOutput != nil
}

// SetSignalOutput gets a reference to the given bool and assigns it to the SignalOutput field.
func (o *RunHistoricalJobRequestAttributes) SetSignalOutput(v bool) {
	o.SignalOutput = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RunHistoricalJobRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FromRule != nil {
		toSerialize["fromRule"] = o.FromRule
	}
	if o.JobDefinition != nil {
		toSerialize["jobDefinition"] = o.JobDefinition
	}
	if o.SignalOutput != nil {
		toSerialize["signalOutput"] = o.SignalOutput
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RunHistoricalJobRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FromRule      *JobDefinitionFromRule `json:"fromRule,omitempty"`
		JobDefinition *JobDefinition         `json:"jobDefinition,omitempty"`
		SignalOutput  *bool                  `json:"signalOutput,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fromRule", "jobDefinition", "signalOutput"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.FromRule != nil && all.FromRule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FromRule = all.FromRule
	if all.JobDefinition != nil && all.JobDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JobDefinition = all.JobDefinition
	o.SignalOutput = all.SignalOutput

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
