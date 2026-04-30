// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultStepElementUpdates Element locator updates produced during a step.
type SyntheticsTestResultStepElementUpdates struct {
	// Updated multi-locator definition.
	MultiLocator map[string]string `json:"multi_locator,omitempty"`
	// Updated outer HTML of the targeted element.
	TargetOuterHtml *string `json:"target_outer_html,omitempty"`
	// Version of the element locator definition.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultStepElementUpdates instantiates a new SyntheticsTestResultStepElementUpdates object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultStepElementUpdates() *SyntheticsTestResultStepElementUpdates {
	this := SyntheticsTestResultStepElementUpdates{}
	return &this
}

// NewSyntheticsTestResultStepElementUpdatesWithDefaults instantiates a new SyntheticsTestResultStepElementUpdates object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultStepElementUpdatesWithDefaults() *SyntheticsTestResultStepElementUpdates {
	this := SyntheticsTestResultStepElementUpdates{}
	return &this
}

// GetMultiLocator returns the MultiLocator field value if set, zero value otherwise.
func (o *SyntheticsTestResultStepElementUpdates) GetMultiLocator() map[string]string {
	if o == nil || o.MultiLocator == nil {
		var ret map[string]string
		return ret
	}
	return o.MultiLocator
}

// GetMultiLocatorOk returns a tuple with the MultiLocator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStepElementUpdates) GetMultiLocatorOk() (*map[string]string, bool) {
	if o == nil || o.MultiLocator == nil {
		return nil, false
	}
	return &o.MultiLocator, true
}

// HasMultiLocator returns a boolean if a field has been set.
func (o *SyntheticsTestResultStepElementUpdates) HasMultiLocator() bool {
	return o != nil && o.MultiLocator != nil
}

// SetMultiLocator gets a reference to the given map[string]string and assigns it to the MultiLocator field.
func (o *SyntheticsTestResultStepElementUpdates) SetMultiLocator(v map[string]string) {
	o.MultiLocator = v
}

// GetTargetOuterHtml returns the TargetOuterHtml field value if set, zero value otherwise.
func (o *SyntheticsTestResultStepElementUpdates) GetTargetOuterHtml() string {
	if o == nil || o.TargetOuterHtml == nil {
		var ret string
		return ret
	}
	return *o.TargetOuterHtml
}

// GetTargetOuterHtmlOk returns a tuple with the TargetOuterHtml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStepElementUpdates) GetTargetOuterHtmlOk() (*string, bool) {
	if o == nil || o.TargetOuterHtml == nil {
		return nil, false
	}
	return o.TargetOuterHtml, true
}

// HasTargetOuterHtml returns a boolean if a field has been set.
func (o *SyntheticsTestResultStepElementUpdates) HasTargetOuterHtml() bool {
	return o != nil && o.TargetOuterHtml != nil
}

// SetTargetOuterHtml gets a reference to the given string and assigns it to the TargetOuterHtml field.
func (o *SyntheticsTestResultStepElementUpdates) SetTargetOuterHtml(v string) {
	o.TargetOuterHtml = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SyntheticsTestResultStepElementUpdates) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStepElementUpdates) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SyntheticsTestResultStepElementUpdates) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *SyntheticsTestResultStepElementUpdates) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultStepElementUpdates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MultiLocator != nil {
		toSerialize["multi_locator"] = o.MultiLocator
	}
	if o.TargetOuterHtml != nil {
		toSerialize["target_outer_html"] = o.TargetOuterHtml
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultStepElementUpdates) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MultiLocator    map[string]string `json:"multi_locator,omitempty"`
		TargetOuterHtml *string           `json:"target_outer_html,omitempty"`
		Version         *int64            `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"multi_locator", "target_outer_html", "version"})
	} else {
		return err
	}
	o.MultiLocator = all.MultiLocator
	o.TargetOuterHtml = all.TargetOuterHtml
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
