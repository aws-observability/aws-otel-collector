// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultStepAssertionResult Assertion result for a browser or mobile step.
type SyntheticsTestResultStepAssertionResult struct {
	// Actual value observed during the step assertion. Its type depends on the check type.
	Actual interface{} `json:"actual,omitempty"`
	// Type of the step assertion check.
	CheckType *string `json:"check_type,omitempty"`
	// Expected value for the step assertion. Its type depends on the check type.
	Expected interface{} `json:"expected,omitempty"`
	// Whether the assertion involves secure variables.
	HasSecureVariables *bool `json:"has_secure_variables,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultStepAssertionResult instantiates a new SyntheticsTestResultStepAssertionResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultStepAssertionResult() *SyntheticsTestResultStepAssertionResult {
	this := SyntheticsTestResultStepAssertionResult{}
	return &this
}

// NewSyntheticsTestResultStepAssertionResultWithDefaults instantiates a new SyntheticsTestResultStepAssertionResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultStepAssertionResultWithDefaults() *SyntheticsTestResultStepAssertionResult {
	this := SyntheticsTestResultStepAssertionResult{}
	return &this
}

// GetActual returns the Actual field value if set, zero value otherwise.
func (o *SyntheticsTestResultStepAssertionResult) GetActual() interface{} {
	if o == nil || o.Actual == nil {
		var ret interface{}
		return ret
	}
	return o.Actual
}

// GetActualOk returns a tuple with the Actual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStepAssertionResult) GetActualOk() (*interface{}, bool) {
	if o == nil || o.Actual == nil {
		return nil, false
	}
	return &o.Actual, true
}

// HasActual returns a boolean if a field has been set.
func (o *SyntheticsTestResultStepAssertionResult) HasActual() bool {
	return o != nil && o.Actual != nil
}

// SetActual gets a reference to the given interface{} and assigns it to the Actual field.
func (o *SyntheticsTestResultStepAssertionResult) SetActual(v interface{}) {
	o.Actual = v
}

// GetCheckType returns the CheckType field value if set, zero value otherwise.
func (o *SyntheticsTestResultStepAssertionResult) GetCheckType() string {
	if o == nil || o.CheckType == nil {
		var ret string
		return ret
	}
	return *o.CheckType
}

// GetCheckTypeOk returns a tuple with the CheckType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStepAssertionResult) GetCheckTypeOk() (*string, bool) {
	if o == nil || o.CheckType == nil {
		return nil, false
	}
	return o.CheckType, true
}

// HasCheckType returns a boolean if a field has been set.
func (o *SyntheticsTestResultStepAssertionResult) HasCheckType() bool {
	return o != nil && o.CheckType != nil
}

// SetCheckType gets a reference to the given string and assigns it to the CheckType field.
func (o *SyntheticsTestResultStepAssertionResult) SetCheckType(v string) {
	o.CheckType = &v
}

// GetExpected returns the Expected field value if set, zero value otherwise.
func (o *SyntheticsTestResultStepAssertionResult) GetExpected() interface{} {
	if o == nil || o.Expected == nil {
		var ret interface{}
		return ret
	}
	return o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStepAssertionResult) GetExpectedOk() (*interface{}, bool) {
	if o == nil || o.Expected == nil {
		return nil, false
	}
	return &o.Expected, true
}

// HasExpected returns a boolean if a field has been set.
func (o *SyntheticsTestResultStepAssertionResult) HasExpected() bool {
	return o != nil && o.Expected != nil
}

// SetExpected gets a reference to the given interface{} and assigns it to the Expected field.
func (o *SyntheticsTestResultStepAssertionResult) SetExpected(v interface{}) {
	o.Expected = v
}

// GetHasSecureVariables returns the HasSecureVariables field value if set, zero value otherwise.
func (o *SyntheticsTestResultStepAssertionResult) GetHasSecureVariables() bool {
	if o == nil || o.HasSecureVariables == nil {
		var ret bool
		return ret
	}
	return *o.HasSecureVariables
}

// GetHasSecureVariablesOk returns a tuple with the HasSecureVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStepAssertionResult) GetHasSecureVariablesOk() (*bool, bool) {
	if o == nil || o.HasSecureVariables == nil {
		return nil, false
	}
	return o.HasSecureVariables, true
}

// HasHasSecureVariables returns a boolean if a field has been set.
func (o *SyntheticsTestResultStepAssertionResult) HasHasSecureVariables() bool {
	return o != nil && o.HasSecureVariables != nil
}

// SetHasSecureVariables gets a reference to the given bool and assigns it to the HasSecureVariables field.
func (o *SyntheticsTestResultStepAssertionResult) SetHasSecureVariables(v bool) {
	o.HasSecureVariables = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultStepAssertionResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Actual != nil {
		toSerialize["actual"] = o.Actual
	}
	if o.CheckType != nil {
		toSerialize["check_type"] = o.CheckType
	}
	if o.Expected != nil {
		toSerialize["expected"] = o.Expected
	}
	if o.HasSecureVariables != nil {
		toSerialize["has_secure_variables"] = o.HasSecureVariables
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultStepAssertionResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actual             interface{} `json:"actual,omitempty"`
		CheckType          *string     `json:"check_type,omitempty"`
		Expected           interface{} `json:"expected,omitempty"`
		HasSecureVariables *bool       `json:"has_secure_variables,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actual", "check_type", "expected", "has_secure_variables"})
	} else {
		return err
	}
	o.Actual = all.Actual
	o.CheckType = all.CheckType
	o.Expected = all.Expected
	o.HasSecureVariables = all.HasSecureVariables

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
