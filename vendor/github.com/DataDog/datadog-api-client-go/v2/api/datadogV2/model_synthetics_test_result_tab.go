// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultTab Information about a browser tab involved in a step.
type SyntheticsTestResultTab struct {
	// Whether the tab was focused during the step.
	Focused *bool `json:"focused,omitempty"`
	// Title of the tab.
	Title *string `json:"title,omitempty"`
	// URL loaded in the tab.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultTab instantiates a new SyntheticsTestResultTab object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultTab() *SyntheticsTestResultTab {
	this := SyntheticsTestResultTab{}
	return &this
}

// NewSyntheticsTestResultTabWithDefaults instantiates a new SyntheticsTestResultTab object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultTabWithDefaults() *SyntheticsTestResultTab {
	this := SyntheticsTestResultTab{}
	return &this
}

// GetFocused returns the Focused field value if set, zero value otherwise.
func (o *SyntheticsTestResultTab) GetFocused() bool {
	if o == nil || o.Focused == nil {
		var ret bool
		return ret
	}
	return *o.Focused
}

// GetFocusedOk returns a tuple with the Focused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTab) GetFocusedOk() (*bool, bool) {
	if o == nil || o.Focused == nil {
		return nil, false
	}
	return o.Focused, true
}

// HasFocused returns a boolean if a field has been set.
func (o *SyntheticsTestResultTab) HasFocused() bool {
	return o != nil && o.Focused != nil
}

// SetFocused gets a reference to the given bool and assigns it to the Focused field.
func (o *SyntheticsTestResultTab) SetFocused(v bool) {
	o.Focused = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SyntheticsTestResultTab) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTab) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SyntheticsTestResultTab) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SyntheticsTestResultTab) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SyntheticsTestResultTab) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTab) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SyntheticsTestResultTab) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SyntheticsTestResultTab) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultTab) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Focused != nil {
		toSerialize["focused"] = o.Focused
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultTab) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Focused *bool   `json:"focused,omitempty"`
		Title   *string `json:"title,omitempty"`
		Url     *string `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"focused", "title", "url"})
	} else {
		return err
	}
	o.Focused = all.Focused
	o.Title = all.Title
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
