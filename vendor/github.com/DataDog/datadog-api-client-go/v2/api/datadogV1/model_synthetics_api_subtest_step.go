// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAPISubtestStep The subtest step used in a Synthetics multi-step API test.
type SyntheticsAPISubtestStep struct {
	// Determines whether or not to continue with test if this step fails.
	AllowFailure *bool `json:"allowFailure,omitempty"`
	// A boolean set to always execute this step even if the previous step failed or was skipped.
	AlwaysExecute *bool `json:"alwaysExecute,omitempty"`
	// Determines whether or not to exit the test if the step succeeds.
	ExitIfSucceed *bool `json:"exitIfSucceed,omitempty"`
	// Generate variables using JavaScript.
	ExtractedValuesFromScript *string `json:"extractedValuesFromScript,omitempty"`
	// ID of the step.
	Id *string `json:"id,omitempty"`
	// Determines whether or not to consider the entire test as failed if this step fails.
	// Can be used only if `allowFailure` is `true`.
	IsCritical *bool `json:"isCritical,omitempty"`
	// The name of the step.
	Name string `json:"name"`
	// Object describing the retry strategy to apply to a Synthetic test.
	Retry *SyntheticsTestOptionsRetry `json:"retry,omitempty"`
	// Public ID of the test to be played as part of a `playSubTest` step type.
	SubtestPublicId string `json:"subtestPublicId"`
	// The subtype of the Synthetic multi-step API subtest step.
	Subtype SyntheticsAPISubtestStepSubtype `json:"subtype"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsAPISubtestStep instantiates a new SyntheticsAPISubtestStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsAPISubtestStep(name string, subtestPublicId string, subtype SyntheticsAPISubtestStepSubtype) *SyntheticsAPISubtestStep {
	this := SyntheticsAPISubtestStep{}
	this.Name = name
	this.SubtestPublicId = subtestPublicId
	this.Subtype = subtype
	return &this
}

// NewSyntheticsAPISubtestStepWithDefaults instantiates a new SyntheticsAPISubtestStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsAPISubtestStepWithDefaults() *SyntheticsAPISubtestStep {
	this := SyntheticsAPISubtestStep{}
	return &this
}

// GetAllowFailure returns the AllowFailure field value if set, zero value otherwise.
func (o *SyntheticsAPISubtestStep) GetAllowFailure() bool {
	if o == nil || o.AllowFailure == nil {
		var ret bool
		return ret
	}
	return *o.AllowFailure
}

// GetAllowFailureOk returns a tuple with the AllowFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPISubtestStep) GetAllowFailureOk() (*bool, bool) {
	if o == nil || o.AllowFailure == nil {
		return nil, false
	}
	return o.AllowFailure, true
}

// HasAllowFailure returns a boolean if a field has been set.
func (o *SyntheticsAPISubtestStep) HasAllowFailure() bool {
	return o != nil && o.AllowFailure != nil
}

// SetAllowFailure gets a reference to the given bool and assigns it to the AllowFailure field.
func (o *SyntheticsAPISubtestStep) SetAllowFailure(v bool) {
	o.AllowFailure = &v
}

// GetAlwaysExecute returns the AlwaysExecute field value if set, zero value otherwise.
func (o *SyntheticsAPISubtestStep) GetAlwaysExecute() bool {
	if o == nil || o.AlwaysExecute == nil {
		var ret bool
		return ret
	}
	return *o.AlwaysExecute
}

// GetAlwaysExecuteOk returns a tuple with the AlwaysExecute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPISubtestStep) GetAlwaysExecuteOk() (*bool, bool) {
	if o == nil || o.AlwaysExecute == nil {
		return nil, false
	}
	return o.AlwaysExecute, true
}

// HasAlwaysExecute returns a boolean if a field has been set.
func (o *SyntheticsAPISubtestStep) HasAlwaysExecute() bool {
	return o != nil && o.AlwaysExecute != nil
}

// SetAlwaysExecute gets a reference to the given bool and assigns it to the AlwaysExecute field.
func (o *SyntheticsAPISubtestStep) SetAlwaysExecute(v bool) {
	o.AlwaysExecute = &v
}

// GetExitIfSucceed returns the ExitIfSucceed field value if set, zero value otherwise.
func (o *SyntheticsAPISubtestStep) GetExitIfSucceed() bool {
	if o == nil || o.ExitIfSucceed == nil {
		var ret bool
		return ret
	}
	return *o.ExitIfSucceed
}

// GetExitIfSucceedOk returns a tuple with the ExitIfSucceed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPISubtestStep) GetExitIfSucceedOk() (*bool, bool) {
	if o == nil || o.ExitIfSucceed == nil {
		return nil, false
	}
	return o.ExitIfSucceed, true
}

// HasExitIfSucceed returns a boolean if a field has been set.
func (o *SyntheticsAPISubtestStep) HasExitIfSucceed() bool {
	return o != nil && o.ExitIfSucceed != nil
}

// SetExitIfSucceed gets a reference to the given bool and assigns it to the ExitIfSucceed field.
func (o *SyntheticsAPISubtestStep) SetExitIfSucceed(v bool) {
	o.ExitIfSucceed = &v
}

// GetExtractedValuesFromScript returns the ExtractedValuesFromScript field value if set, zero value otherwise.
func (o *SyntheticsAPISubtestStep) GetExtractedValuesFromScript() string {
	if o == nil || o.ExtractedValuesFromScript == nil {
		var ret string
		return ret
	}
	return *o.ExtractedValuesFromScript
}

// GetExtractedValuesFromScriptOk returns a tuple with the ExtractedValuesFromScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPISubtestStep) GetExtractedValuesFromScriptOk() (*string, bool) {
	if o == nil || o.ExtractedValuesFromScript == nil {
		return nil, false
	}
	return o.ExtractedValuesFromScript, true
}

// HasExtractedValuesFromScript returns a boolean if a field has been set.
func (o *SyntheticsAPISubtestStep) HasExtractedValuesFromScript() bool {
	return o != nil && o.ExtractedValuesFromScript != nil
}

// SetExtractedValuesFromScript gets a reference to the given string and assigns it to the ExtractedValuesFromScript field.
func (o *SyntheticsAPISubtestStep) SetExtractedValuesFromScript(v string) {
	o.ExtractedValuesFromScript = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsAPISubtestStep) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPISubtestStep) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsAPISubtestStep) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsAPISubtestStep) SetId(v string) {
	o.Id = &v
}

// GetIsCritical returns the IsCritical field value if set, zero value otherwise.
func (o *SyntheticsAPISubtestStep) GetIsCritical() bool {
	if o == nil || o.IsCritical == nil {
		var ret bool
		return ret
	}
	return *o.IsCritical
}

// GetIsCriticalOk returns a tuple with the IsCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPISubtestStep) GetIsCriticalOk() (*bool, bool) {
	if o == nil || o.IsCritical == nil {
		return nil, false
	}
	return o.IsCritical, true
}

// HasIsCritical returns a boolean if a field has been set.
func (o *SyntheticsAPISubtestStep) HasIsCritical() bool {
	return o != nil && o.IsCritical != nil
}

// SetIsCritical gets a reference to the given bool and assigns it to the IsCritical field.
func (o *SyntheticsAPISubtestStep) SetIsCritical(v bool) {
	o.IsCritical = &v
}

// GetName returns the Name field value.
func (o *SyntheticsAPISubtestStep) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAPISubtestStep) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsAPISubtestStep) SetName(v string) {
	o.Name = v
}

// GetRetry returns the Retry field value if set, zero value otherwise.
func (o *SyntheticsAPISubtestStep) GetRetry() SyntheticsTestOptionsRetry {
	if o == nil || o.Retry == nil {
		var ret SyntheticsTestOptionsRetry
		return ret
	}
	return *o.Retry
}

// GetRetryOk returns a tuple with the Retry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPISubtestStep) GetRetryOk() (*SyntheticsTestOptionsRetry, bool) {
	if o == nil || o.Retry == nil {
		return nil, false
	}
	return o.Retry, true
}

// HasRetry returns a boolean if a field has been set.
func (o *SyntheticsAPISubtestStep) HasRetry() bool {
	return o != nil && o.Retry != nil
}

// SetRetry gets a reference to the given SyntheticsTestOptionsRetry and assigns it to the Retry field.
func (o *SyntheticsAPISubtestStep) SetRetry(v SyntheticsTestOptionsRetry) {
	o.Retry = &v
}

// GetSubtestPublicId returns the SubtestPublicId field value.
func (o *SyntheticsAPISubtestStep) GetSubtestPublicId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SubtestPublicId
}

// GetSubtestPublicIdOk returns a tuple with the SubtestPublicId field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAPISubtestStep) GetSubtestPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubtestPublicId, true
}

// SetSubtestPublicId sets field value.
func (o *SyntheticsAPISubtestStep) SetSubtestPublicId(v string) {
	o.SubtestPublicId = v
}

// GetSubtype returns the Subtype field value.
func (o *SyntheticsAPISubtestStep) GetSubtype() SyntheticsAPISubtestStepSubtype {
	if o == nil {
		var ret SyntheticsAPISubtestStepSubtype
		return ret
	}
	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAPISubtestStep) GetSubtypeOk() (*SyntheticsAPISubtestStepSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value.
func (o *SyntheticsAPISubtestStep) SetSubtype(v SyntheticsAPISubtestStepSubtype) {
	o.Subtype = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsAPISubtestStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowFailure != nil {
		toSerialize["allowFailure"] = o.AllowFailure
	}
	if o.AlwaysExecute != nil {
		toSerialize["alwaysExecute"] = o.AlwaysExecute
	}
	if o.ExitIfSucceed != nil {
		toSerialize["exitIfSucceed"] = o.ExitIfSucceed
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
	if o.Retry != nil {
		toSerialize["retry"] = o.Retry
	}
	toSerialize["subtestPublicId"] = o.SubtestPublicId
	toSerialize["subtype"] = o.Subtype

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsAPISubtestStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowFailure              *bool                            `json:"allowFailure,omitempty"`
		AlwaysExecute             *bool                            `json:"alwaysExecute,omitempty"`
		ExitIfSucceed             *bool                            `json:"exitIfSucceed,omitempty"`
		ExtractedValuesFromScript *string                          `json:"extractedValuesFromScript,omitempty"`
		Id                        *string                          `json:"id,omitempty"`
		IsCritical                *bool                            `json:"isCritical,omitempty"`
		Name                      *string                          `json:"name"`
		Retry                     *SyntheticsTestOptionsRetry      `json:"retry,omitempty"`
		SubtestPublicId           *string                          `json:"subtestPublicId"`
		Subtype                   *SyntheticsAPISubtestStepSubtype `json:"subtype"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SubtestPublicId == nil {
		return fmt.Errorf("required field subtestPublicId missing")
	}
	if all.Subtype == nil {
		return fmt.Errorf("required field subtype missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowFailure", "alwaysExecute", "exitIfSucceed", "extractedValuesFromScript", "id", "isCritical", "name", "retry", "subtestPublicId", "subtype"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowFailure = all.AllowFailure
	o.AlwaysExecute = all.AlwaysExecute
	o.ExitIfSucceed = all.ExitIfSucceed
	o.ExtractedValuesFromScript = all.ExtractedValuesFromScript
	o.Id = all.Id
	o.IsCritical = all.IsCritical
	o.Name = *all.Name
	if all.Retry != nil && all.Retry.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Retry = all.Retry
	o.SubtestPublicId = *all.SubtestPublicId
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
