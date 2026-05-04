// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultRumContext RUM application context associated with a step or sub-test.
type SyntheticsTestResultRumContext struct {
	// RUM application identifier.
	ApplicationId *string `json:"application_id,omitempty"`
	// RUM session identifier.
	SessionId *string `json:"session_id,omitempty"`
	// RUM view identifier.
	ViewId *string `json:"view_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultRumContext instantiates a new SyntheticsTestResultRumContext object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultRumContext() *SyntheticsTestResultRumContext {
	this := SyntheticsTestResultRumContext{}
	return &this
}

// NewSyntheticsTestResultRumContextWithDefaults instantiates a new SyntheticsTestResultRumContext object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultRumContextWithDefaults() *SyntheticsTestResultRumContext {
	this := SyntheticsTestResultRumContext{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *SyntheticsTestResultRumContext) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRumContext) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *SyntheticsTestResultRumContext) HasApplicationId() bool {
	return o != nil && o.ApplicationId != nil
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *SyntheticsTestResultRumContext) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *SyntheticsTestResultRumContext) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRumContext) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *SyntheticsTestResultRumContext) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *SyntheticsTestResultRumContext) SetSessionId(v string) {
	o.SessionId = &v
}

// GetViewId returns the ViewId field value if set, zero value otherwise.
func (o *SyntheticsTestResultRumContext) GetViewId() string {
	if o == nil || o.ViewId == nil {
		var ret string
		return ret
	}
	return *o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRumContext) GetViewIdOk() (*string, bool) {
	if o == nil || o.ViewId == nil {
		return nil, false
	}
	return o.ViewId, true
}

// HasViewId returns a boolean if a field has been set.
func (o *SyntheticsTestResultRumContext) HasViewId() bool {
	return o != nil && o.ViewId != nil
}

// SetViewId gets a reference to the given string and assigns it to the ViewId field.
func (o *SyntheticsTestResultRumContext) SetViewId(v string) {
	o.ViewId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultRumContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}
	if o.SessionId != nil {
		toSerialize["session_id"] = o.SessionId
	}
	if o.ViewId != nil {
		toSerialize["view_id"] = o.ViewId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultRumContext) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId *string `json:"application_id,omitempty"`
		SessionId     *string `json:"session_id,omitempty"`
		ViewId        *string `json:"view_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "session_id", "view_id"})
	} else {
		return err
	}
	o.ApplicationId = all.ApplicationId
	o.SessionId = all.SessionId
	o.ViewId = all.ViewId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
