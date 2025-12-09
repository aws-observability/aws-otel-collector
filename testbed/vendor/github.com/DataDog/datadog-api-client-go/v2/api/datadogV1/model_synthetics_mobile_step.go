// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStep The steps used in a Synthetic mobile test.
type SyntheticsMobileStep struct {
	// A boolean set to allow this step to fail.
	AllowFailure *bool `json:"allowFailure,omitempty"`
	// A boolean set to determine if the step has a new step element.
	HasNewStepElement *bool `json:"hasNewStepElement,omitempty"`
	// A boolean to use in addition to `allowFailure` to determine if the test should be marked as failed when the step fails.
	IsCritical *bool `json:"isCritical,omitempty"`
	// The name of the step.
	Name string `json:"name"`
	// A boolean set to not take a screenshot for the step.
	NoScreenshot *bool `json:"noScreenshot,omitempty"`
	// The parameters of a mobile step.
	Params SyntheticsMobileStepParams `json:"params"`
	// The public ID of the step.
	PublicId *string `json:"publicId,omitempty"`
	// The time before declaring a step failed.
	Timeout *int64 `json:"timeout,omitempty"`
	// Step type used in your mobile Synthetic test.
	Type SyntheticsMobileStepType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileStep instantiates a new SyntheticsMobileStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileStep(name string, params SyntheticsMobileStepParams, typeVar SyntheticsMobileStepType) *SyntheticsMobileStep {
	this := SyntheticsMobileStep{}
	this.Name = name
	this.Params = params
	this.Type = typeVar
	return &this
}

// NewSyntheticsMobileStepWithDefaults instantiates a new SyntheticsMobileStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileStepWithDefaults() *SyntheticsMobileStep {
	this := SyntheticsMobileStep{}
	return &this
}

// GetAllowFailure returns the AllowFailure field value if set, zero value otherwise.
func (o *SyntheticsMobileStep) GetAllowFailure() bool {
	if o == nil || o.AllowFailure == nil {
		var ret bool
		return ret
	}
	return *o.AllowFailure
}

// GetAllowFailureOk returns a tuple with the AllowFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStep) GetAllowFailureOk() (*bool, bool) {
	if o == nil || o.AllowFailure == nil {
		return nil, false
	}
	return o.AllowFailure, true
}

// HasAllowFailure returns a boolean if a field has been set.
func (o *SyntheticsMobileStep) HasAllowFailure() bool {
	return o != nil && o.AllowFailure != nil
}

// SetAllowFailure gets a reference to the given bool and assigns it to the AllowFailure field.
func (o *SyntheticsMobileStep) SetAllowFailure(v bool) {
	o.AllowFailure = &v
}

// GetHasNewStepElement returns the HasNewStepElement field value if set, zero value otherwise.
func (o *SyntheticsMobileStep) GetHasNewStepElement() bool {
	if o == nil || o.HasNewStepElement == nil {
		var ret bool
		return ret
	}
	return *o.HasNewStepElement
}

// GetHasNewStepElementOk returns a tuple with the HasNewStepElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStep) GetHasNewStepElementOk() (*bool, bool) {
	if o == nil || o.HasNewStepElement == nil {
		return nil, false
	}
	return o.HasNewStepElement, true
}

// HasHasNewStepElement returns a boolean if a field has been set.
func (o *SyntheticsMobileStep) HasHasNewStepElement() bool {
	return o != nil && o.HasNewStepElement != nil
}

// SetHasNewStepElement gets a reference to the given bool and assigns it to the HasNewStepElement field.
func (o *SyntheticsMobileStep) SetHasNewStepElement(v bool) {
	o.HasNewStepElement = &v
}

// GetIsCritical returns the IsCritical field value if set, zero value otherwise.
func (o *SyntheticsMobileStep) GetIsCritical() bool {
	if o == nil || o.IsCritical == nil {
		var ret bool
		return ret
	}
	return *o.IsCritical
}

// GetIsCriticalOk returns a tuple with the IsCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStep) GetIsCriticalOk() (*bool, bool) {
	if o == nil || o.IsCritical == nil {
		return nil, false
	}
	return o.IsCritical, true
}

// HasIsCritical returns a boolean if a field has been set.
func (o *SyntheticsMobileStep) HasIsCritical() bool {
	return o != nil && o.IsCritical != nil
}

// SetIsCritical gets a reference to the given bool and assigns it to the IsCritical field.
func (o *SyntheticsMobileStep) SetIsCritical(v bool) {
	o.IsCritical = &v
}

// GetName returns the Name field value.
func (o *SyntheticsMobileStep) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStep) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsMobileStep) SetName(v string) {
	o.Name = v
}

// GetNoScreenshot returns the NoScreenshot field value if set, zero value otherwise.
func (o *SyntheticsMobileStep) GetNoScreenshot() bool {
	if o == nil || o.NoScreenshot == nil {
		var ret bool
		return ret
	}
	return *o.NoScreenshot
}

// GetNoScreenshotOk returns a tuple with the NoScreenshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStep) GetNoScreenshotOk() (*bool, bool) {
	if o == nil || o.NoScreenshot == nil {
		return nil, false
	}
	return o.NoScreenshot, true
}

// HasNoScreenshot returns a boolean if a field has been set.
func (o *SyntheticsMobileStep) HasNoScreenshot() bool {
	return o != nil && o.NoScreenshot != nil
}

// SetNoScreenshot gets a reference to the given bool and assigns it to the NoScreenshot field.
func (o *SyntheticsMobileStep) SetNoScreenshot(v bool) {
	o.NoScreenshot = &v
}

// GetParams returns the Params field value.
func (o *SyntheticsMobileStep) GetParams() SyntheticsMobileStepParams {
	if o == nil {
		var ret SyntheticsMobileStepParams
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStep) GetParamsOk() (*SyntheticsMobileStepParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value.
func (o *SyntheticsMobileStep) SetParams(v SyntheticsMobileStepParams) {
	o.Params = v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *SyntheticsMobileStep) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStep) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *SyntheticsMobileStep) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *SyntheticsMobileStep) SetPublicId(v string) {
	o.PublicId = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *SyntheticsMobileStep) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStep) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *SyntheticsMobileStep) HasTimeout() bool {
	return o != nil && o.Timeout != nil
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *SyntheticsMobileStep) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetType returns the Type field value.
func (o *SyntheticsMobileStep) GetType() SyntheticsMobileStepType {
	if o == nil {
		var ret SyntheticsMobileStepType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStep) GetTypeOk() (*SyntheticsMobileStepType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsMobileStep) SetType(v SyntheticsMobileStepType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowFailure != nil {
		toSerialize["allowFailure"] = o.AllowFailure
	}
	if o.HasNewStepElement != nil {
		toSerialize["hasNewStepElement"] = o.HasNewStepElement
	}
	if o.IsCritical != nil {
		toSerialize["isCritical"] = o.IsCritical
	}
	toSerialize["name"] = o.Name
	if o.NoScreenshot != nil {
		toSerialize["noScreenshot"] = o.NoScreenshot
	}
	toSerialize["params"] = o.Params
	if o.PublicId != nil {
		toSerialize["publicId"] = o.PublicId
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowFailure      *bool                       `json:"allowFailure,omitempty"`
		HasNewStepElement *bool                       `json:"hasNewStepElement,omitempty"`
		IsCritical        *bool                       `json:"isCritical,omitempty"`
		Name              *string                     `json:"name"`
		NoScreenshot      *bool                       `json:"noScreenshot,omitempty"`
		Params            *SyntheticsMobileStepParams `json:"params"`
		PublicId          *string                     `json:"publicId,omitempty"`
		Timeout           *int64                      `json:"timeout,omitempty"`
		Type              *SyntheticsMobileStepType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Params == nil {
		return fmt.Errorf("required field params missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowFailure", "hasNewStepElement", "isCritical", "name", "noScreenshot", "params", "publicId", "timeout", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowFailure = all.AllowFailure
	o.HasNewStepElement = all.HasNewStepElement
	o.IsCritical = all.IsCritical
	o.Name = *all.Name
	o.NoScreenshot = all.NoScreenshot
	if all.Params.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Params = *all.Params
	o.PublicId = all.PublicId
	o.Timeout = all.Timeout
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
