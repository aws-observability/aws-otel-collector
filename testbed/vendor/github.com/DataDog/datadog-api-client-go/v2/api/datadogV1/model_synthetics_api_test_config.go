// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAPITestConfig Configuration object for a Synthetic API test.
type SyntheticsAPITestConfig struct {
	// Array of assertions used for the test. Required for single API tests.
	Assertions []SyntheticsAssertion `json:"assertions,omitempty"`
	// Array of variables used for the test.
	ConfigVariables []SyntheticsConfigVariable `json:"configVariables,omitempty"`
	// Object describing the Synthetic test request.
	Request *SyntheticsTestRequest `json:"request,omitempty"`
	// When the test subtype is `multi`, the steps of the test.
	Steps []SyntheticsAPIStep `json:"steps,omitempty"`
	// Variables defined from JavaScript code.
	VariablesFromScript *string `json:"variablesFromScript,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsAPITestConfig instantiates a new SyntheticsAPITestConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsAPITestConfig() *SyntheticsAPITestConfig {
	this := SyntheticsAPITestConfig{}
	return &this
}

// NewSyntheticsAPITestConfigWithDefaults instantiates a new SyntheticsAPITestConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsAPITestConfigWithDefaults() *SyntheticsAPITestConfig {
	this := SyntheticsAPITestConfig{}
	return &this
}

// GetAssertions returns the Assertions field value if set, zero value otherwise.
func (o *SyntheticsAPITestConfig) GetAssertions() []SyntheticsAssertion {
	if o == nil || o.Assertions == nil {
		var ret []SyntheticsAssertion
		return ret
	}
	return o.Assertions
}

// GetAssertionsOk returns a tuple with the Assertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestConfig) GetAssertionsOk() (*[]SyntheticsAssertion, bool) {
	if o == nil || o.Assertions == nil {
		return nil, false
	}
	return &o.Assertions, true
}

// HasAssertions returns a boolean if a field has been set.
func (o *SyntheticsAPITestConfig) HasAssertions() bool {
	return o != nil && o.Assertions != nil
}

// SetAssertions gets a reference to the given []SyntheticsAssertion and assigns it to the Assertions field.
func (o *SyntheticsAPITestConfig) SetAssertions(v []SyntheticsAssertion) {
	o.Assertions = v
}

// GetConfigVariables returns the ConfigVariables field value if set, zero value otherwise.
func (o *SyntheticsAPITestConfig) GetConfigVariables() []SyntheticsConfigVariable {
	if o == nil || o.ConfigVariables == nil {
		var ret []SyntheticsConfigVariable
		return ret
	}
	return o.ConfigVariables
}

// GetConfigVariablesOk returns a tuple with the ConfigVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestConfig) GetConfigVariablesOk() (*[]SyntheticsConfigVariable, bool) {
	if o == nil || o.ConfigVariables == nil {
		return nil, false
	}
	return &o.ConfigVariables, true
}

// HasConfigVariables returns a boolean if a field has been set.
func (o *SyntheticsAPITestConfig) HasConfigVariables() bool {
	return o != nil && o.ConfigVariables != nil
}

// SetConfigVariables gets a reference to the given []SyntheticsConfigVariable and assigns it to the ConfigVariables field.
func (o *SyntheticsAPITestConfig) SetConfigVariables(v []SyntheticsConfigVariable) {
	o.ConfigVariables = v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SyntheticsAPITestConfig) GetRequest() SyntheticsTestRequest {
	if o == nil || o.Request == nil {
		var ret SyntheticsTestRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestConfig) GetRequestOk() (*SyntheticsTestRequest, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SyntheticsAPITestConfig) HasRequest() bool {
	return o != nil && o.Request != nil
}

// SetRequest gets a reference to the given SyntheticsTestRequest and assigns it to the Request field.
func (o *SyntheticsAPITestConfig) SetRequest(v SyntheticsTestRequest) {
	o.Request = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *SyntheticsAPITestConfig) GetSteps() []SyntheticsAPIStep {
	if o == nil || o.Steps == nil {
		var ret []SyntheticsAPIStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestConfig) GetStepsOk() (*[]SyntheticsAPIStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *SyntheticsAPITestConfig) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []SyntheticsAPIStep and assigns it to the Steps field.
func (o *SyntheticsAPITestConfig) SetSteps(v []SyntheticsAPIStep) {
	o.Steps = v
}

// GetVariablesFromScript returns the VariablesFromScript field value if set, zero value otherwise.
func (o *SyntheticsAPITestConfig) GetVariablesFromScript() string {
	if o == nil || o.VariablesFromScript == nil {
		var ret string
		return ret
	}
	return *o.VariablesFromScript
}

// GetVariablesFromScriptOk returns a tuple with the VariablesFromScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestConfig) GetVariablesFromScriptOk() (*string, bool) {
	if o == nil || o.VariablesFromScript == nil {
		return nil, false
	}
	return o.VariablesFromScript, true
}

// HasVariablesFromScript returns a boolean if a field has been set.
func (o *SyntheticsAPITestConfig) HasVariablesFromScript() bool {
	return o != nil && o.VariablesFromScript != nil
}

// SetVariablesFromScript gets a reference to the given string and assigns it to the VariablesFromScript field.
func (o *SyntheticsAPITestConfig) SetVariablesFromScript(v string) {
	o.VariablesFromScript = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsAPITestConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assertions != nil {
		toSerialize["assertions"] = o.Assertions
	}
	if o.ConfigVariables != nil {
		toSerialize["configVariables"] = o.ConfigVariables
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.VariablesFromScript != nil {
		toSerialize["variablesFromScript"] = o.VariablesFromScript
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsAPITestConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assertions          []SyntheticsAssertion      `json:"assertions,omitempty"`
		ConfigVariables     []SyntheticsConfigVariable `json:"configVariables,omitempty"`
		Request             *SyntheticsTestRequest     `json:"request,omitempty"`
		Steps               []SyntheticsAPIStep        `json:"steps,omitempty"`
		VariablesFromScript *string                    `json:"variablesFromScript,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assertions", "configVariables", "request", "steps", "variablesFromScript"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Assertions = all.Assertions
	o.ConfigVariables = all.ConfigVariables
	if all.Request != nil && all.Request.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Request = all.Request
	o.Steps = all.Steps
	o.VariablesFromScript = all.VariablesFromScript

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
