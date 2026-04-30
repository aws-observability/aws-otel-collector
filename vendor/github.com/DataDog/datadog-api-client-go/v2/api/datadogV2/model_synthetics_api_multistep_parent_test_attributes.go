// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsApiMultistepParentTestAttributes Attributes of a parent API multistep test.
type SyntheticsApiMultistepParentTestAttributes struct {
	// The name of the child subtest.
	ChildName *string `json:"child_name,omitempty"`
	// The public ID of the child subtest.
	ChildPublicId *string `json:"child_public_id,omitempty"`
	// The associated monitor ID.
	MonitorId *int64 `json:"monitor_id,omitempty"`
	// Name of the parent test.
	Name *string `json:"name,omitempty"`
	// The overall state of the parent test.
	OverallState *int64 `json:"overall_state,omitempty"`
	// Timestamp of when the overall state was last modified.
	OverallStateModified *string `json:"overall_state_modified,omitempty"`
	// The public ID of the parent test.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsApiMultistepParentTestAttributes instantiates a new SyntheticsApiMultistepParentTestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsApiMultistepParentTestAttributes() *SyntheticsApiMultistepParentTestAttributes {
	this := SyntheticsApiMultistepParentTestAttributes{}
	return &this
}

// NewSyntheticsApiMultistepParentTestAttributesWithDefaults instantiates a new SyntheticsApiMultistepParentTestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsApiMultistepParentTestAttributesWithDefaults() *SyntheticsApiMultistepParentTestAttributes {
	this := SyntheticsApiMultistepParentTestAttributes{}
	return &this
}

// GetChildName returns the ChildName field value if set, zero value otherwise.
func (o *SyntheticsApiMultistepParentTestAttributes) GetChildName() string {
	if o == nil || o.ChildName == nil {
		var ret string
		return ret
	}
	return *o.ChildName
}

// GetChildNameOk returns a tuple with the ChildName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) GetChildNameOk() (*string, bool) {
	if o == nil || o.ChildName == nil {
		return nil, false
	}
	return o.ChildName, true
}

// HasChildName returns a boolean if a field has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) HasChildName() bool {
	return o != nil && o.ChildName != nil
}

// SetChildName gets a reference to the given string and assigns it to the ChildName field.
func (o *SyntheticsApiMultistepParentTestAttributes) SetChildName(v string) {
	o.ChildName = &v
}

// GetChildPublicId returns the ChildPublicId field value if set, zero value otherwise.
func (o *SyntheticsApiMultistepParentTestAttributes) GetChildPublicId() string {
	if o == nil || o.ChildPublicId == nil {
		var ret string
		return ret
	}
	return *o.ChildPublicId
}

// GetChildPublicIdOk returns a tuple with the ChildPublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) GetChildPublicIdOk() (*string, bool) {
	if o == nil || o.ChildPublicId == nil {
		return nil, false
	}
	return o.ChildPublicId, true
}

// HasChildPublicId returns a boolean if a field has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) HasChildPublicId() bool {
	return o != nil && o.ChildPublicId != nil
}

// SetChildPublicId gets a reference to the given string and assigns it to the ChildPublicId field.
func (o *SyntheticsApiMultistepParentTestAttributes) SetChildPublicId(v string) {
	o.ChildPublicId = &v
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *SyntheticsApiMultistepParentTestAttributes) GetMonitorId() int64 {
	if o == nil || o.MonitorId == nil {
		var ret int64
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) GetMonitorIdOk() (*int64, bool) {
	if o == nil || o.MonitorId == nil {
		return nil, false
	}
	return o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) HasMonitorId() bool {
	return o != nil && o.MonitorId != nil
}

// SetMonitorId gets a reference to the given int64 and assigns it to the MonitorId field.
func (o *SyntheticsApiMultistepParentTestAttributes) SetMonitorId(v int64) {
	o.MonitorId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsApiMultistepParentTestAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsApiMultistepParentTestAttributes) SetName(v string) {
	o.Name = &v
}

// GetOverallState returns the OverallState field value if set, zero value otherwise.
func (o *SyntheticsApiMultistepParentTestAttributes) GetOverallState() int64 {
	if o == nil || o.OverallState == nil {
		var ret int64
		return ret
	}
	return *o.OverallState
}

// GetOverallStateOk returns a tuple with the OverallState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) GetOverallStateOk() (*int64, bool) {
	if o == nil || o.OverallState == nil {
		return nil, false
	}
	return o.OverallState, true
}

// HasOverallState returns a boolean if a field has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) HasOverallState() bool {
	return o != nil && o.OverallState != nil
}

// SetOverallState gets a reference to the given int64 and assigns it to the OverallState field.
func (o *SyntheticsApiMultistepParentTestAttributes) SetOverallState(v int64) {
	o.OverallState = &v
}

// GetOverallStateModified returns the OverallStateModified field value if set, zero value otherwise.
func (o *SyntheticsApiMultistepParentTestAttributes) GetOverallStateModified() string {
	if o == nil || o.OverallStateModified == nil {
		var ret string
		return ret
	}
	return *o.OverallStateModified
}

// GetOverallStateModifiedOk returns a tuple with the OverallStateModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) GetOverallStateModifiedOk() (*string, bool) {
	if o == nil || o.OverallStateModified == nil {
		return nil, false
	}
	return o.OverallStateModified, true
}

// HasOverallStateModified returns a boolean if a field has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) HasOverallStateModified() bool {
	return o != nil && o.OverallStateModified != nil
}

// SetOverallStateModified gets a reference to the given string and assigns it to the OverallStateModified field.
func (o *SyntheticsApiMultistepParentTestAttributes) SetOverallStateModified(v string) {
	o.OverallStateModified = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *SyntheticsApiMultistepParentTestAttributes) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *SyntheticsApiMultistepParentTestAttributes) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *SyntheticsApiMultistepParentTestAttributes) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsApiMultistepParentTestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChildName != nil {
		toSerialize["child_name"] = o.ChildName
	}
	if o.ChildPublicId != nil {
		toSerialize["child_public_id"] = o.ChildPublicId
	}
	if o.MonitorId != nil {
		toSerialize["monitor_id"] = o.MonitorId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OverallState != nil {
		toSerialize["overall_state"] = o.OverallState
	}
	if o.OverallStateModified != nil {
		toSerialize["overall_state_modified"] = o.OverallStateModified
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsApiMultistepParentTestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChildName            *string `json:"child_name,omitempty"`
		ChildPublicId        *string `json:"child_public_id,omitempty"`
		MonitorId            *int64  `json:"monitor_id,omitempty"`
		Name                 *string `json:"name,omitempty"`
		OverallState         *int64  `json:"overall_state,omitempty"`
		OverallStateModified *string `json:"overall_state_modified,omitempty"`
		PublicId             *string `json:"public_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"child_name", "child_public_id", "monitor_id", "name", "overall_state", "overall_state_modified", "public_id"})
	} else {
		return err
	}
	o.ChildName = all.ChildName
	o.ChildPublicId = all.ChildPublicId
	o.MonitorId = all.MonitorId
	o.Name = all.Name
	o.OverallState = all.OverallState
	o.OverallStateModified = all.OverallStateModified
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
