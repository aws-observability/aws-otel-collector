// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestIncludedUserAttributes Attributes of an included user.
type ChangeRequestIncludedUserAttributes struct {
	// The email of the user.
	Email string `json:"email"`
	// The handle of the user.
	Handle string `json:"handle"`
	// The name of the user.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestIncludedUserAttributes instantiates a new ChangeRequestIncludedUserAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestIncludedUserAttributes(email string, handle string, name string) *ChangeRequestIncludedUserAttributes {
	this := ChangeRequestIncludedUserAttributes{}
	this.Email = email
	this.Handle = handle
	this.Name = name
	return &this
}

// NewChangeRequestIncludedUserAttributesWithDefaults instantiates a new ChangeRequestIncludedUserAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestIncludedUserAttributesWithDefaults() *ChangeRequestIncludedUserAttributes {
	this := ChangeRequestIncludedUserAttributes{}
	return &this
}

// GetEmail returns the Email field value.
func (o *ChangeRequestIncludedUserAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestIncludedUserAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *ChangeRequestIncludedUserAttributes) SetEmail(v string) {
	o.Email = v
}

// GetHandle returns the Handle field value.
func (o *ChangeRequestIncludedUserAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestIncludedUserAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *ChangeRequestIncludedUserAttributes) SetHandle(v string) {
	o.Handle = v
}

// GetName returns the Name field value.
func (o *ChangeRequestIncludedUserAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestIncludedUserAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ChangeRequestIncludedUserAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestIncludedUserAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["email"] = o.Email
	toSerialize["handle"] = o.Handle
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestIncludedUserAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Email  *string `json:"email"`
		Handle *string `json:"handle"`
		Name   *string `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"email", "handle", "name"})
	} else {
		return err
	}
	o.Email = *all.Email
	o.Handle = *all.Handle
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
