// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardAuthor User who shared the dashboard.
type SharedDashboardAuthor struct {
	// Identifier of the user who shared the dashboard.
	Handle *string `json:"handle,omitempty"`
	// Name of the user who shared the dashboard.
	Name datadog.NullableString `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSharedDashboardAuthor instantiates a new SharedDashboardAuthor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSharedDashboardAuthor() *SharedDashboardAuthor {
	this := SharedDashboardAuthor{}
	return &this
}

// NewSharedDashboardAuthorWithDefaults instantiates a new SharedDashboardAuthor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSharedDashboardAuthorWithDefaults() *SharedDashboardAuthor {
	this := SharedDashboardAuthor{}
	return &this
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *SharedDashboardAuthor) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardAuthor) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *SharedDashboardAuthor) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *SharedDashboardAuthor) SetHandle(v string) {
	o.Handle = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedDashboardAuthor) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardAuthor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SharedDashboardAuthor) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *SharedDashboardAuthor) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *SharedDashboardAuthor) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *SharedDashboardAuthor) UnsetName() {
	o.Name.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SharedDashboardAuthor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SharedDashboardAuthor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle *string                `json:"handle,omitempty"`
		Name   datadog.NullableString `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "name"})
	} else {
		return err
	}
	o.Handle = all.Handle
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
