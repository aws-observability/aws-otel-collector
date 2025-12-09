// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAPITestStep The Test step used in a Synthetic multi-step API test.
type SyntheticsAPITestStep struct {
	// Determines whether or not to continue with test if this step fails.
	AllowFailure *bool `json:"allowFailure,omitempty"`
	// Array of assertions used for the test.
	Assertions []SyntheticsAssertion `json:"assertions"`
	// Determines whether or not to exit the test if the step succeeds.
	ExitIfSucceed *bool `json:"exitIfSucceed,omitempty"`
	// Array of values to parse and save as variables from the response.
	ExtractedValues []SyntheticsParsingOptions `json:"extractedValues,omitempty"`
	// Generate variables using JavaScript.
	ExtractedValuesFromScript *string `json:"extractedValuesFromScript,omitempty"`
	// ID of the step.
	Id *string `json:"id,omitempty"`
	// Determines whether or not to consider the entire test as failed if this step fails.
	// Can be used only if `allowFailure` is `true`.
	IsCritical *bool `json:"isCritical,omitempty"`
	// The name of the step.
	Name string `json:"name"`
	// Object describing the Synthetic test request.
	Request SyntheticsTestRequest `json:"request"`
	// Object describing the retry strategy to apply to a Synthetic test.
	Retry *SyntheticsTestOptionsRetry `json:"retry,omitempty"`
	// The subtype of the Synthetic multi-step API test step.
	Subtype SyntheticsAPITestStepSubtype `json:"subtype"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsAPITestStep instantiates a new SyntheticsAPITestStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsAPITestStep(assertions []SyntheticsAssertion, name string, request SyntheticsTestRequest, subtype SyntheticsAPITestStepSubtype) *SyntheticsAPITestStep {
	this := SyntheticsAPITestStep{}
	this.Assertions = assertions
	this.Name = name
	this.Request = request
	this.Subtype = subtype
	return &this
}

// NewSyntheticsAPITestStepWithDefaults instantiates a new SyntheticsAPITestStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsAPITestStepWithDefaults() *SyntheticsAPITestStep {
	this := SyntheticsAPITestStep{}
	return &this
}

// GetAllowFailure returns the AllowFailure field value if set, zero value otherwise.
func (o *SyntheticsAPITestStep) GetAllowFailure() bool {
	if o == nil || o.AllowFailure == nil {
		var ret bool
		return ret
	}
	return *o.AllowFailure
}

// GetAllowFailureOk returns a tuple with the AllowFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestStep) GetAllowFailureOk() (*bool, bool) {
	if o == nil || o.AllowFailure == nil {
		return nil, false
	}
	return o.AllowFailure, true
}

// HasAllowFailure returns a boolean if a field has been set.
func (o *SyntheticsAPITestStep) HasAllowFailure() bool {
	return o != nil && o.AllowFailure != nil
}

// SetAllowFailure gets a reference to the given bool and assigns it to the AllowFailure field.
func (o *SyntheticsAPITestStep) SetAllowFailure(v bool) {
	o.AllowFailure = &v
}

// GetAssertions returns the Assertions field value.
func (o *SyntheticsAPITestStep) GetAssertions() []SyntheticsAssertion {
	if o == nil {
		var ret []SyntheticsAssertion
		return ret
	}
	return o.Assertions
}

// GetAssertionsOk returns a tuple with the Assertions field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestStep) GetAssertionsOk() (*[]SyntheticsAssertion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assertions, true
}

// SetAssertions sets field value.
func (o *SyntheticsAPITestStep) SetAssertions(v []SyntheticsAssertion) {
	o.Assertions = v
}

// GetExitIfSucceed returns the ExitIfSucceed field value if set, zero value otherwise.
func (o *SyntheticsAPITestStep) GetExitIfSucceed() bool {
	if o == nil || o.ExitIfSucceed == nil {
		var ret bool
		return ret
	}
	return *o.ExitIfSucceed
}

// GetExitIfSucceedOk returns a tuple with the ExitIfSucceed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestStep) GetExitIfSucceedOk() (*bool, bool) {
	if o == nil || o.ExitIfSucceed == nil {
		return nil, false
	}
	return o.ExitIfSucceed, true
}

// HasExitIfSucceed returns a boolean if a field has been set.
func (o *SyntheticsAPITestStep) HasExitIfSucceed() bool {
	return o != nil && o.ExitIfSucceed != nil
}

// SetExitIfSucceed gets a reference to the given bool and assigns it to the ExitIfSucceed field.
func (o *SyntheticsAPITestStep) SetExitIfSucceed(v bool) {
	o.ExitIfSucceed = &v
}

// GetExtractedValues returns the ExtractedValues field value if set, zero value otherwise.
func (o *SyntheticsAPITestStep) GetExtractedValues() []SyntheticsParsingOptions {
	if o == nil || o.ExtractedValues == nil {
		var ret []SyntheticsParsingOptions
		return ret
	}
	return o.ExtractedValues
}

// GetExtractedValuesOk returns a tuple with the ExtractedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestStep) GetExtractedValuesOk() (*[]SyntheticsParsingOptions, bool) {
	if o == nil || o.ExtractedValues == nil {
		return nil, false
	}
	return &o.ExtractedValues, true
}

// HasExtractedValues returns a boolean if a field has been set.
func (o *SyntheticsAPITestStep) HasExtractedValues() bool {
	return o != nil && o.ExtractedValues != nil
}

// SetExtractedValues gets a reference to the given []SyntheticsParsingOptions and assigns it to the ExtractedValues field.
func (o *SyntheticsAPITestStep) SetExtractedValues(v []SyntheticsParsingOptions) {
	o.ExtractedValues = v
}

// GetExtractedValuesFromScript returns the ExtractedValuesFromScript field value if set, zero value otherwise.
func (o *SyntheticsAPITestStep) GetExtractedValuesFromScript() string {
	if o == nil || o.ExtractedValuesFromScript == nil {
		var ret string
		return ret
	}
	return *o.ExtractedValuesFromScript
}

// GetExtractedValuesFromScriptOk returns a tuple with the ExtractedValuesFromScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestStep) GetExtractedValuesFromScriptOk() (*string, bool) {
	if o == nil || o.ExtractedValuesFromScript == nil {
		return nil, false
	}
	return o.ExtractedValuesFromScript, true
}

// HasExtractedValuesFromScript returns a boolean if a field has been set.
func (o *SyntheticsAPITestStep) HasExtractedValuesFromScript() bool {
	return o != nil && o.ExtractedValuesFromScript != nil
}

// SetExtractedValuesFromScript gets a reference to the given string and assigns it to the ExtractedValuesFromScript field.
func (o *SyntheticsAPITestStep) SetExtractedValuesFromScript(v string) {
	o.ExtractedValuesFromScript = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsAPITestStep) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestStep) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsAPITestStep) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsAPITestStep) SetId(v string) {
	o.Id = &v
}

// GetIsCritical returns the IsCritical field value if set, zero value otherwise.
func (o *SyntheticsAPITestStep) GetIsCritical() bool {
	if o == nil || o.IsCritical == nil {
		var ret bool
		return ret
	}
	return *o.IsCritical
}

// GetIsCriticalOk returns a tuple with the IsCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestStep) GetIsCriticalOk() (*bool, bool) {
	if o == nil || o.IsCritical == nil {
		return nil, false
	}
	return o.IsCritical, true
}

// HasIsCritical returns a boolean if a field has been set.
func (o *SyntheticsAPITestStep) HasIsCritical() bool {
	return o != nil && o.IsCritical != nil
}

// SetIsCritical gets a reference to the given bool and assigns it to the IsCritical field.
func (o *SyntheticsAPITestStep) SetIsCritical(v bool) {
	o.IsCritical = &v
}

// GetName returns the Name field value.
func (o *SyntheticsAPITestStep) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestStep) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsAPITestStep) SetName(v string) {
	o.Name = v
}

// GetRequest returns the Request field value.
func (o *SyntheticsAPITestStep) GetRequest() SyntheticsTestRequest {
	if o == nil {
		var ret SyntheticsTestRequest
		return ret
	}
	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestStep) GetRequestOk() (*SyntheticsTestRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value.
func (o *SyntheticsAPITestStep) SetRequest(v SyntheticsTestRequest) {
	o.Request = v
}

// GetRetry returns the Retry field value if set, zero value otherwise.
func (o *SyntheticsAPITestStep) GetRetry() SyntheticsTestOptionsRetry {
	if o == nil || o.Retry == nil {
		var ret SyntheticsTestOptionsRetry
		return ret
	}
	return *o.Retry
}

// GetRetryOk returns a tuple with the Retry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestStep) GetRetryOk() (*SyntheticsTestOptionsRetry, bool) {
	if o == nil || o.Retry == nil {
		return nil, false
	}
	return o.Retry, true
}

// HasRetry returns a boolean if a field has been set.
func (o *SyntheticsAPITestStep) HasRetry() bool {
	return o != nil && o.Retry != nil
}

// SetRetry gets a reference to the given SyntheticsTestOptionsRetry and assigns it to the Retry field.
func (o *SyntheticsAPITestStep) SetRetry(v SyntheticsTestOptionsRetry) {
	o.Retry = &v
}

// GetSubtype returns the Subtype field value.
func (o *SyntheticsAPITestStep) GetSubtype() SyntheticsAPITestStepSubtype {
	if o == nil {
		var ret SyntheticsAPITestStepSubtype
		return ret
	}
	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestStep) GetSubtypeOk() (*SyntheticsAPITestStepSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value.
func (o *SyntheticsAPITestStep) SetSubtype(v SyntheticsAPITestStepSubtype) {
	o.Subtype = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsAPITestStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowFailure != nil {
		toSerialize["allowFailure"] = o.AllowFailure
	}
	toSerialize["assertions"] = o.Assertions
	if o.ExitIfSucceed != nil {
		toSerialize["exitIfSucceed"] = o.ExitIfSucceed
	}
	if o.ExtractedValues != nil {
		toSerialize["extractedValues"] = o.ExtractedValues
	}
	if o.ExtractedValuesFromScript != nil {
		toSerialize["extractedValuesFromScript"] = o.ExtractedValuesFromScript
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsCritical != nil {
		toSerialize["isCritical"] = o.IsCritical
	}
	toSerialize["name"] = o.Name
	toSerialize["request"] = o.Request
	if o.Retry != nil {
		toSerialize["retry"] = o.Retry
	}
	toSerialize["subtype"] = o.Subtype

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsAPITestStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowFailure              *bool                         `json:"allowFailure,omitempty"`
		Assertions                *[]SyntheticsAssertion        `json:"assertions"`
		ExitIfSucceed             *bool                         `json:"exitIfSucceed,omitempty"`
		ExtractedValues           []SyntheticsParsingOptions    `json:"extractedValues,omitempty"`
		ExtractedValuesFromScript *string                       `json:"extractedValuesFromScript,omitempty"`
		Id                        *string                       `json:"id,omitempty"`
		IsCritical                *bool                         `json:"isCritical,omitempty"`
		Name                      *string                       `json:"name"`
		Request                   *SyntheticsTestRequest        `json:"request"`
		Retry                     *SyntheticsTestOptionsRetry   `json:"retry,omitempty"`
		Subtype                   *SyntheticsAPITestStepSubtype `json:"subtype"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Assertions == nil {
		return fmt.Errorf("required field assertions missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Request == nil {
		return fmt.Errorf("required field request missing")
	}
	if all.Subtype == nil {
		return fmt.Errorf("required field subtype missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowFailure", "assertions", "exitIfSucceed", "extractedValues", "extractedValuesFromScript", "id", "isCritical", "name", "request", "retry", "subtype"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowFailure = all.AllowFailure
	o.Assertions = *all.Assertions
	o.ExitIfSucceed = all.ExitIfSucceed
	o.ExtractedValues = all.ExtractedValues
	o.ExtractedValuesFromScript = all.ExtractedValuesFromScript
	o.Id = all.Id
	o.IsCritical = all.IsCritical
	o.Name = *all.Name
	if all.Request.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Request = *all.Request
	if all.Retry != nil && all.Retry.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Retry = all.Retry
	if !all.Subtype.IsValid() {
		hasInvalidField = true
	} else {
		o.Subtype = *all.Subtype
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
