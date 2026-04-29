// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowUserAttributes Attributes of a ServiceNow user
type ServiceNowUserAttributes struct {
	// The email address of the user
	Email string `json:"email"`
	// The full name of the user
	FullName *string `json:"full_name,omitempty"`
	// The ID of the ServiceNow instance
	InstanceId uuid.UUID `json:"instance_id"`
	// The username of the ServiceNow user
	UserName string `json:"user_name"`
	// The system ID of the user in ServiceNow
	UserSysId string `json:"user_sys_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceNowUserAttributes instantiates a new ServiceNowUserAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceNowUserAttributes(email string, instanceId uuid.UUID, userName string, userSysId string) *ServiceNowUserAttributes {
	this := ServiceNowUserAttributes{}
	this.Email = email
	this.InstanceId = instanceId
	this.UserName = userName
	this.UserSysId = userSysId
	return &this
}

// NewServiceNowUserAttributesWithDefaults instantiates a new ServiceNowUserAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceNowUserAttributesWithDefaults() *ServiceNowUserAttributes {
	this := ServiceNowUserAttributes{}
	return &this
}

// GetEmail returns the Email field value.
func (o *ServiceNowUserAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ServiceNowUserAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *ServiceNowUserAttributes) SetEmail(v string) {
	o.Email = v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *ServiceNowUserAttributes) GetFullName() string {
	if o == nil || o.FullName == nil {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowUserAttributes) GetFullNameOk() (*string, bool) {
	if o == nil || o.FullName == nil {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *ServiceNowUserAttributes) HasFullName() bool {
	return o != nil && o.FullName != nil
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *ServiceNowUserAttributes) SetFullName(v string) {
	o.FullName = &v
}

// GetInstanceId returns the InstanceId field value.
func (o *ServiceNowUserAttributes) GetInstanceId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *ServiceNowUserAttributes) GetInstanceIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value.
func (o *ServiceNowUserAttributes) SetInstanceId(v uuid.UUID) {
	o.InstanceId = v
}

// GetUserName returns the UserName field value.
func (o *ServiceNowUserAttributes) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *ServiceNowUserAttributes) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value.
func (o *ServiceNowUserAttributes) SetUserName(v string) {
	o.UserName = v
}

// GetUserSysId returns the UserSysId field value.
func (o *ServiceNowUserAttributes) GetUserSysId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserSysId
}

// GetUserSysIdOk returns a tuple with the UserSysId field value
// and a boolean to check if the value has been set.
func (o *ServiceNowUserAttributes) GetUserSysIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserSysId, true
}

// SetUserSysId sets field value.
func (o *ServiceNowUserAttributes) SetUserSysId(v string) {
	o.UserSysId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceNowUserAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["email"] = o.Email
	if o.FullName != nil {
		toSerialize["full_name"] = o.FullName
	}
	toSerialize["instance_id"] = o.InstanceId
	toSerialize["user_name"] = o.UserName
	toSerialize["user_sys_id"] = o.UserSysId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceNowUserAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Email      *string    `json:"email"`
		FullName   *string    `json:"full_name,omitempty"`
		InstanceId *uuid.UUID `json:"instance_id"`
		UserName   *string    `json:"user_name"`
		UserSysId  *string    `json:"user_sys_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	if all.InstanceId == nil {
		return fmt.Errorf("required field instance_id missing")
	}
	if all.UserName == nil {
		return fmt.Errorf("required field user_name missing")
	}
	if all.UserSysId == nil {
		return fmt.Errorf("required field user_sys_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"email", "full_name", "instance_id", "user_name", "user_sys_id"})
	} else {
		return err
	}
	o.Email = *all.Email
	o.FullName = all.FullName
	o.InstanceId = *all.InstanceId
	o.UserName = *all.UserName
	o.UserSysId = *all.UserSysId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
