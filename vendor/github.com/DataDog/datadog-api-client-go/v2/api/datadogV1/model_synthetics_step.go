// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsStep The steps used in a Synthetic browser test.
type SyntheticsStep struct {
	// A boolean set to allow this step to fail.
	AllowFailure *bool `json:"allowFailure,omitempty"`
	// A boolean set to always execute this step even if the previous step failed or was skipped.
	AlwaysExecute *bool `json:"alwaysExecute,omitempty"`
	// A boolean set to exit the test if the step succeeds.
	ExitIfSucceed *bool `json:"exitIfSucceed,omitempty"`
	// A boolean to use in addition to `allowFailure` to determine if the test should be marked as failed when the step fails.
	IsCritical *bool `json:"isCritical,omitempty"`
	// The name of the step.
	Name *string `json:"name,omitempty"`
	// A boolean set to skip taking a screenshot for the step.
	NoScreenshot *bool `json:"noScreenshot,omitempty"`
	// The parameters of the step.
	Params interface{} `json:"params,omitempty"`
	// The public ID of the step.
	PublicId *string `json:"public_id,omitempty"`
	// The time before declaring a step failed.
	Timeout *int64 `json:"timeout,omitempty"`
	// Step type used in your Synthetic test.
	Type *SyntheticsStepType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsStep instantiates a new SyntheticsStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsStep() *SyntheticsStep {
	this := SyntheticsStep{}
	return &this
}

// NewSyntheticsStepWithDefaults instantiates a new SyntheticsStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsStepWithDefaults() *SyntheticsStep {
	this := SyntheticsStep{}
	return &this
}

// GetAllowFailure returns the AllowFailure field value if set, zero value otherwise.
func (o *SyntheticsStep) GetAllowFailure() bool {
	if o == nil || o.AllowFailure == nil {
		var ret bool
		return ret
	}
	return *o.AllowFailure
}

// GetAllowFailureOk returns a tuple with the AllowFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStep) GetAllowFailureOk() (*bool, bool) {
	if o == nil || o.AllowFailure == nil {
		return nil, false
	}
	return o.AllowFailure, true
}

// HasAllowFailure returns a boolean if a field has been set.
func (o *SyntheticsStep) HasAllowFailure() bool {
	return o != nil && o.AllowFailure != nil
}

// SetAllowFailure gets a reference to the given bool and assigns it to the AllowFailure field.
func (o *SyntheticsStep) SetAllowFailure(v bool) {
	o.AllowFailure = &v
}

// GetAlwaysExecute returns the AlwaysExecute field value if set, zero value otherwise.
func (o *SyntheticsStep) GetAlwaysExecute() bool {
	if o == nil || o.AlwaysExecute == nil {
		var ret bool
		return ret
	}
	return *o.AlwaysExecute
}

// GetAlwaysExecuteOk returns a tuple with the AlwaysExecute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStep) GetAlwaysExecuteOk() (*bool, bool) {
	if o == nil || o.AlwaysExecute == nil {
		return nil, false
	}
	return o.AlwaysExecute, true
}

// HasAlwaysExecute returns a boolean if a field has been set.
func (o *SyntheticsStep) HasAlwaysExecute() bool {
	return o != nil && o.AlwaysExecute != nil
}

// SetAlwaysExecute gets a reference to the given bool and assigns it to the AlwaysExecute field.
func (o *SyntheticsStep) SetAlwaysExecute(v bool) {
	o.AlwaysExecute = &v
}

// GetExitIfSucceed returns the ExitIfSucceed field value if set, zero value otherwise.
func (o *SyntheticsStep) GetExitIfSucceed() bool {
	if o == nil || o.ExitIfSucceed == nil {
		var ret bool
		return ret
	}
	return *o.ExitIfSucceed
}

// GetExitIfSucceedOk returns a tuple with the ExitIfSucceed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStep) GetExitIfSucceedOk() (*bool, bool) {
	if o == nil || o.ExitIfSucceed == nil {
		return nil, false
	}
	return o.ExitIfSucceed, true
}

// HasExitIfSucceed returns a boolean if a field has been set.
func (o *SyntheticsStep) HasExitIfSucceed() bool {
	return o != nil && o.ExitIfSucceed != nil
}

// SetExitIfSucceed gets a reference to the given bool and assigns it to the ExitIfSucceed field.
func (o *SyntheticsStep) SetExitIfSucceed(v bool) {
	o.ExitIfSucceed = &v
}

// GetIsCritical returns the IsCritical field value if set, zero value otherwise.
func (o *SyntheticsStep) GetIsCritical() bool {
	if o == nil || o.IsCritical == nil {
		var ret bool
		return ret
	}
	return *o.IsCritical
}

// GetIsCriticalOk returns a tuple with the IsCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStep) GetIsCriticalOk() (*bool, bool) {
	if o == nil || o.IsCritical == nil {
		return nil, false
	}
	return o.IsCritical, true
}

// HasIsCritical returns a boolean if a field has been set.
func (o *SyntheticsStep) HasIsCritical() bool {
	return o != nil && o.IsCritical != nil
}

// SetIsCritical gets a reference to the given bool and assigns it to the IsCritical field.
func (o *SyntheticsStep) SetIsCritical(v bool) {
	o.IsCritical = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsStep) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStep) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsStep) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsStep) SetName(v string) {
	o.Name = &v
}

// GetNoScreenshot returns the NoScreenshot field value if set, zero value otherwise.
func (o *SyntheticsStep) GetNoScreenshot() bool {
	if o == nil || o.NoScreenshot == nil {
		var ret bool
		return ret
	}
	return *o.NoScreenshot
}

// GetNoScreenshotOk returns a tuple with the NoScreenshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStep) GetNoScreenshotOk() (*bool, bool) {
	if o == nil || o.NoScreenshot == nil {
		return nil, false
	}
	return o.NoScreenshot, true
}

// HasNoScreenshot returns a boolean if a field has been set.
func (o *SyntheticsStep) HasNoScreenshot() bool {
	return o != nil && o.NoScreenshot != nil
}

// SetNoScreenshot gets a reference to the given bool and assigns it to the NoScreenshot field.
func (o *SyntheticsStep) SetNoScreenshot(v bool) {
	o.NoScreenshot = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *SyntheticsStep) GetParams() interface{} {
	if o == nil || o.Params == nil {
		var ret interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStep) GetParamsOk() (*interface{}, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return &o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *SyntheticsStep) HasParams() bool {
	return o != nil && o.Params != nil
}

// SetParams gets a reference to the given interface{} and assigns it to the Params field.
func (o *SyntheticsStep) SetParams(v interface{}) {
	o.Params = v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *SyntheticsStep) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStep) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *SyntheticsStep) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *SyntheticsStep) SetPublicId(v string) {
	o.PublicId = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *SyntheticsStep) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStep) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *SyntheticsStep) HasTimeout() bool {
	return o != nil && o.Timeout != nil
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *SyntheticsStep) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsStep) GetType() SyntheticsStepType {
	if o == nil || o.Type == nil {
		var ret SyntheticsStepType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStep) GetTypeOk() (*SyntheticsStepType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsStep) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SyntheticsStepType and assigns it to the Type field.
func (o *SyntheticsStep) SetType(v SyntheticsStepType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsStep) MarshalJSON() ([]byte, error) {
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
	if o.IsCritical != nil {
		toSerialize["isCritical"] = o.IsCritical
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NoScreenshot != nil {
		toSerialize["noScreenshot"] = o.NoScreenshot
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowFailure  *bool               `json:"allowFailure,omitempty"`
		AlwaysExecute *bool               `json:"alwaysExecute,omitempty"`
		ExitIfSucceed *bool               `json:"exitIfSucceed,omitempty"`
		IsCritical    *bool               `json:"isCritical,omitempty"`
		Name          *string             `json:"name,omitempty"`
		NoScreenshot  *bool               `json:"noScreenshot,omitempty"`
		Params        interface{}         `json:"params,omitempty"`
		PublicId      *string             `json:"public_id,omitempty"`
		Timeout       *int64              `json:"timeout,omitempty"`
		Type          *SyntheticsStepType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowFailure", "alwaysExecute", "exitIfSucceed", "isCritical", "name", "noScreenshot", "params", "public_id", "timeout", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowFailure = all.AllowFailure
	o.AlwaysExecute = all.AlwaysExecute
	o.ExitIfSucceed = all.ExitIfSucceed
	o.IsCritical = all.IsCritical
	o.Name = all.Name
	o.NoScreenshot = all.NoScreenshot
	o.Params = all.Params
	o.PublicId = all.PublicId
	o.Timeout = all.Timeout
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
