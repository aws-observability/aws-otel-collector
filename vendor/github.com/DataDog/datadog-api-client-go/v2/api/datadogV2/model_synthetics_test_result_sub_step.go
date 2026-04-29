// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultSubStep Information about a sub-step in a nested test execution.
type SyntheticsTestResultSubStep struct {
	// Depth of the sub-step in the execution tree.
	Level *int64 `json:"level,omitempty"`
	// Reference to the parent step of a sub-step.
	ParentStep *SyntheticsTestResultParentStep `json:"parent_step,omitempty"`
	// Reference to the parent test of a sub-step.
	ParentTest *SyntheticsTestResultParentTest `json:"parent_test,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultSubStep instantiates a new SyntheticsTestResultSubStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultSubStep() *SyntheticsTestResultSubStep {
	this := SyntheticsTestResultSubStep{}
	return &this
}

// NewSyntheticsTestResultSubStepWithDefaults instantiates a new SyntheticsTestResultSubStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultSubStepWithDefaults() *SyntheticsTestResultSubStep {
	this := SyntheticsTestResultSubStep{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *SyntheticsTestResultSubStep) GetLevel() int64 {
	if o == nil || o.Level == nil {
		var ret int64
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSubStep) GetLevelOk() (*int64, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *SyntheticsTestResultSubStep) HasLevel() bool {
	return o != nil && o.Level != nil
}

// SetLevel gets a reference to the given int64 and assigns it to the Level field.
func (o *SyntheticsTestResultSubStep) SetLevel(v int64) {
	o.Level = &v
}

// GetParentStep returns the ParentStep field value if set, zero value otherwise.
func (o *SyntheticsTestResultSubStep) GetParentStep() SyntheticsTestResultParentStep {
	if o == nil || o.ParentStep == nil {
		var ret SyntheticsTestResultParentStep
		return ret
	}
	return *o.ParentStep
}

// GetParentStepOk returns a tuple with the ParentStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSubStep) GetParentStepOk() (*SyntheticsTestResultParentStep, bool) {
	if o == nil || o.ParentStep == nil {
		return nil, false
	}
	return o.ParentStep, true
}

// HasParentStep returns a boolean if a field has been set.
func (o *SyntheticsTestResultSubStep) HasParentStep() bool {
	return o != nil && o.ParentStep != nil
}

// SetParentStep gets a reference to the given SyntheticsTestResultParentStep and assigns it to the ParentStep field.
func (o *SyntheticsTestResultSubStep) SetParentStep(v SyntheticsTestResultParentStep) {
	o.ParentStep = &v
}

// GetParentTest returns the ParentTest field value if set, zero value otherwise.
func (o *SyntheticsTestResultSubStep) GetParentTest() SyntheticsTestResultParentTest {
	if o == nil || o.ParentTest == nil {
		var ret SyntheticsTestResultParentTest
		return ret
	}
	return *o.ParentTest
}

// GetParentTestOk returns a tuple with the ParentTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSubStep) GetParentTestOk() (*SyntheticsTestResultParentTest, bool) {
	if o == nil || o.ParentTest == nil {
		return nil, false
	}
	return o.ParentTest, true
}

// HasParentTest returns a boolean if a field has been set.
func (o *SyntheticsTestResultSubStep) HasParentTest() bool {
	return o != nil && o.ParentTest != nil
}

// SetParentTest gets a reference to the given SyntheticsTestResultParentTest and assigns it to the ParentTest field.
func (o *SyntheticsTestResultSubStep) SetParentTest(v SyntheticsTestResultParentTest) {
	o.ParentTest = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultSubStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.ParentStep != nil {
		toSerialize["parent_step"] = o.ParentStep
	}
	if o.ParentTest != nil {
		toSerialize["parent_test"] = o.ParentTest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultSubStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Level      *int64                          `json:"level,omitempty"`
		ParentStep *SyntheticsTestResultParentStep `json:"parent_step,omitempty"`
		ParentTest *SyntheticsTestResultParentTest `json:"parent_test,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"level", "parent_step", "parent_test"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Level = all.Level
	if all.ParentStep != nil && all.ParentStep.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ParentStep = all.ParentStep
	if all.ParentTest != nil && all.ParentTest.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ParentTest = all.ParentTest

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
