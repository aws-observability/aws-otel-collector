// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringTriageUser Object representing a given user entity.
type SecurityMonitoringTriageUser struct {
	// The handle for this user account.
	Handle *string `json:"handle,omitempty"`
	// Gravatar icon associated to the user.
	Icon *string `json:"icon,omitempty"`
	// Numerical ID assigned by Datadog to this user account.
	Id *int64 `json:"id,omitempty"`
	// The name for this user account.
	Name datadog.NullableString `json:"name,omitempty"`
	// UUID assigned by Datadog to this user account.
	Uuid string `json:"uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringTriageUser instantiates a new SecurityMonitoringTriageUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringTriageUser(uuid string) *SecurityMonitoringTriageUser {
	this := SecurityMonitoringTriageUser{}
	this.Uuid = uuid
	return &this
}

// NewSecurityMonitoringTriageUserWithDefaults instantiates a new SecurityMonitoringTriageUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringTriageUserWithDefaults() *SecurityMonitoringTriageUser {
	this := SecurityMonitoringTriageUser{}
	return &this
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *SecurityMonitoringTriageUser) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringTriageUser) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *SecurityMonitoringTriageUser) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *SecurityMonitoringTriageUser) SetHandle(v string) {
	o.Handle = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *SecurityMonitoringTriageUser) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringTriageUser) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *SecurityMonitoringTriageUser) HasIcon() bool {
	return o != nil && o.Icon != nil
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *SecurityMonitoringTriageUser) SetIcon(v string) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityMonitoringTriageUser) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringTriageUser) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityMonitoringTriageUser) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SecurityMonitoringTriageUser) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityMonitoringTriageUser) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityMonitoringTriageUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringTriageUser) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *SecurityMonitoringTriageUser) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *SecurityMonitoringTriageUser) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *SecurityMonitoringTriageUser) UnsetName() {
	o.Name.Unset()
}

// GetUuid returns the Uuid field value.
func (o *SecurityMonitoringTriageUser) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringTriageUser) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value.
func (o *SecurityMonitoringTriageUser) SetUuid(v string) {
	o.Uuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringTriageUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringTriageUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle *string                `json:"handle,omitempty"`
		Icon   *string                `json:"icon,omitempty"`
		Id     *int64                 `json:"id,omitempty"`
		Name   datadog.NullableString `json:"name,omitempty"`
		Uuid   *string                `json:"uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Uuid == nil {
		return fmt.Errorf("required field uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "icon", "id", "name", "uuid"})
	} else {
		return err
	}
	o.Handle = all.Handle
	o.Icon = all.Icon
	o.Id = all.Id
	o.Name = all.Name
	o.Uuid = *all.Uuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
