// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultSubTest Information about a sub-test played from a parent browser test.
type SyntheticsTestResultSubTest struct {
	// Identifier of the sub-test.
	Id *string `json:"id,omitempty"`
	// Index of the browser tab playing the sub-test.
	PlayingTab *int64 `json:"playing_tab,omitempty"`
	// RUM application context associated with a step or sub-test.
	RumContext *SyntheticsTestResultRumContext `json:"rum_context,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultSubTest instantiates a new SyntheticsTestResultSubTest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultSubTest() *SyntheticsTestResultSubTest {
	this := SyntheticsTestResultSubTest{}
	return &this
}

// NewSyntheticsTestResultSubTestWithDefaults instantiates a new SyntheticsTestResultSubTest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultSubTestWithDefaults() *SyntheticsTestResultSubTest {
	this := SyntheticsTestResultSubTest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsTestResultSubTest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSubTest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsTestResultSubTest) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsTestResultSubTest) SetId(v string) {
	o.Id = &v
}

// GetPlayingTab returns the PlayingTab field value if set, zero value otherwise.
func (o *SyntheticsTestResultSubTest) GetPlayingTab() int64 {
	if o == nil || o.PlayingTab == nil {
		var ret int64
		return ret
	}
	return *o.PlayingTab
}

// GetPlayingTabOk returns a tuple with the PlayingTab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSubTest) GetPlayingTabOk() (*int64, bool) {
	if o == nil || o.PlayingTab == nil {
		return nil, false
	}
	return o.PlayingTab, true
}

// HasPlayingTab returns a boolean if a field has been set.
func (o *SyntheticsTestResultSubTest) HasPlayingTab() bool {
	return o != nil && o.PlayingTab != nil
}

// SetPlayingTab gets a reference to the given int64 and assigns it to the PlayingTab field.
func (o *SyntheticsTestResultSubTest) SetPlayingTab(v int64) {
	o.PlayingTab = &v
}

// GetRumContext returns the RumContext field value if set, zero value otherwise.
func (o *SyntheticsTestResultSubTest) GetRumContext() SyntheticsTestResultRumContext {
	if o == nil || o.RumContext == nil {
		var ret SyntheticsTestResultRumContext
		return ret
	}
	return *o.RumContext
}

// GetRumContextOk returns a tuple with the RumContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSubTest) GetRumContextOk() (*SyntheticsTestResultRumContext, bool) {
	if o == nil || o.RumContext == nil {
		return nil, false
	}
	return o.RumContext, true
}

// HasRumContext returns a boolean if a field has been set.
func (o *SyntheticsTestResultSubTest) HasRumContext() bool {
	return o != nil && o.RumContext != nil
}

// SetRumContext gets a reference to the given SyntheticsTestResultRumContext and assigns it to the RumContext field.
func (o *SyntheticsTestResultSubTest) SetRumContext(v SyntheticsTestResultRumContext) {
	o.RumContext = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultSubTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PlayingTab != nil {
		toSerialize["playing_tab"] = o.PlayingTab
	}
	if o.RumContext != nil {
		toSerialize["rum_context"] = o.RumContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultSubTest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id         *string                         `json:"id,omitempty"`
		PlayingTab *int64                          `json:"playing_tab,omitempty"`
		RumContext *SyntheticsTestResultRumContext `json:"rum_context,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "playing_tab", "rum_context"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.PlayingTab = all.PlayingTab
	if all.RumContext != nil && all.RumContext.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RumContext = all.RumContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
