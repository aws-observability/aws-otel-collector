// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardInvitee Invitee that can access an invite-only shared dashboard.
type SharedDashboardInvitee struct {
	// Time when the invitee's access expires.
	AccessExpiration datadog.NullableTime `json:"access_expiration"`
	// Time when the invitee was added.
	CreatedAt time.Time `json:"created_at"`
	// Email address of the invitee.
	Email string `json:"email"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSharedDashboardInvitee instantiates a new SharedDashboardInvitee object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSharedDashboardInvitee(accessExpiration datadog.NullableTime, createdAt time.Time, email string) *SharedDashboardInvitee {
	this := SharedDashboardInvitee{}
	this.AccessExpiration = accessExpiration
	this.CreatedAt = createdAt
	this.Email = email
	return &this
}

// NewSharedDashboardInviteeWithDefaults instantiates a new SharedDashboardInvitee object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSharedDashboardInviteeWithDefaults() *SharedDashboardInvitee {
	this := SharedDashboardInvitee{}
	return &this
}

// GetAccessExpiration returns the AccessExpiration field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *SharedDashboardInvitee) GetAccessExpiration() time.Time {
	if o == nil || o.AccessExpiration.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AccessExpiration.Get()
}

// GetAccessExpirationOk returns a tuple with the AccessExpiration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardInvitee) GetAccessExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessExpiration.Get(), o.AccessExpiration.IsSet()
}

// SetAccessExpiration sets field value.
func (o *SharedDashboardInvitee) SetAccessExpiration(v time.Time) {
	o.AccessExpiration.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value.
func (o *SharedDashboardInvitee) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardInvitee) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *SharedDashboardInvitee) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEmail returns the Email field value.
func (o *SharedDashboardInvitee) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardInvitee) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *SharedDashboardInvitee) SetEmail(v string) {
	o.Email = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SharedDashboardInvitee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["access_expiration"] = o.AccessExpiration.Get()
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["email"] = o.Email

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SharedDashboardInvitee) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessExpiration datadog.NullableTime `json:"access_expiration"`
		CreatedAt        *time.Time           `json:"created_at"`
		Email            *string              `json:"email"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.AccessExpiration.IsSet() {
		return fmt.Errorf("required field access_expiration missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_expiration", "created_at", "email"})
	} else {
		return err
	}
	o.AccessExpiration = all.AccessExpiration
	o.CreatedAt = *all.CreatedAt
	o.Email = *all.Email

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
