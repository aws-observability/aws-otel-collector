// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardUsageUser A user referenced from a dashboard usage record (author or viewer).
type DashboardUsageUser struct {
	// Datadog handle (login) of the user.
	Handle *string `json:"handle,omitempty"`
	// The user ID.
	Id *string `json:"id,omitempty"`
	// Whether the user account is disabled.
	IsDisabled *bool `json:"is_disabled,omitempty"`
	// Display name of the user.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardUsageUser instantiates a new DashboardUsageUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardUsageUser() *DashboardUsageUser {
	this := DashboardUsageUser{}
	return &this
}

// NewDashboardUsageUserWithDefaults instantiates a new DashboardUsageUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardUsageUserWithDefaults() *DashboardUsageUser {
	this := DashboardUsageUser{}
	return &this
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *DashboardUsageUser) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUsageUser) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *DashboardUsageUser) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *DashboardUsageUser) SetHandle(v string) {
	o.Handle = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DashboardUsageUser) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUsageUser) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DashboardUsageUser) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DashboardUsageUser) SetId(v string) {
	o.Id = &v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *DashboardUsageUser) GetIsDisabled() bool {
	if o == nil || o.IsDisabled == nil {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUsageUser) GetIsDisabledOk() (*bool, bool) {
	if o == nil || o.IsDisabled == nil {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *DashboardUsageUser) HasIsDisabled() bool {
	return o != nil && o.IsDisabled != nil
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *DashboardUsageUser) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DashboardUsageUser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUsageUser) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DashboardUsageUser) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DashboardUsageUser) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardUsageUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsDisabled != nil {
		toSerialize["is_disabled"] = o.IsDisabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardUsageUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle     *string `json:"handle,omitempty"`
		Id         *string `json:"id,omitempty"`
		IsDisabled *bool   `json:"is_disabled,omitempty"`
		Name       *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "id", "is_disabled", "name"})
	} else {
		return err
	}
	o.Handle = all.Handle
	o.Id = all.Id
	o.IsDisabled = all.IsDisabled
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableDashboardUsageUser handles when a null is used for DashboardUsageUser.
type NullableDashboardUsageUser struct {
	value *DashboardUsageUser
	isSet bool
}

// Get returns the associated value.
func (v NullableDashboardUsageUser) Get() *DashboardUsageUser {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDashboardUsageUser) Set(val *DashboardUsageUser) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDashboardUsageUser) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDashboardUsageUser) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardUsageUser initializes the struct as if Set has been called.
func NewNullableDashboardUsageUser(val *DashboardUsageUser) *NullableDashboardUsageUser {
	return &NullableDashboardUsageUser{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDashboardUsageUser) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDashboardUsageUser) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
