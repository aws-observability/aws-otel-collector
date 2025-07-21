// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardInviteesItems The allowlisted invitees for an INVITE-only shared dashboard.
type SharedDashboardInviteesItems struct {
	// Time of the invitee expiration. Null means the invite will not expire.
	AccessExpiration datadog.NullableTime `json:"access_expiration,omitempty"`
	// Time that the invitee was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Email of the invitee.
	Email string `json:"email"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSharedDashboardInviteesItems instantiates a new SharedDashboardInviteesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSharedDashboardInviteesItems(email string) *SharedDashboardInviteesItems {
	this := SharedDashboardInviteesItems{}
	this.Email = email
	return &this
}

// NewSharedDashboardInviteesItemsWithDefaults instantiates a new SharedDashboardInviteesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSharedDashboardInviteesItemsWithDefaults() *SharedDashboardInviteesItems {
	this := SharedDashboardInviteesItems{}
	return &this
}

// GetAccessExpiration returns the AccessExpiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedDashboardInviteesItems) GetAccessExpiration() time.Time {
	if o == nil || o.AccessExpiration.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AccessExpiration.Get()
}

// GetAccessExpirationOk returns a tuple with the AccessExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardInviteesItems) GetAccessExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessExpiration.Get(), o.AccessExpiration.IsSet()
}

// HasAccessExpiration returns a boolean if a field has been set.
func (o *SharedDashboardInviteesItems) HasAccessExpiration() bool {
	return o != nil && o.AccessExpiration.IsSet()
}

// SetAccessExpiration gets a reference to the given datadog.NullableTime and assigns it to the AccessExpiration field.
func (o *SharedDashboardInviteesItems) SetAccessExpiration(v time.Time) {
	o.AccessExpiration.Set(&v)
}

// SetAccessExpirationNil sets the value for AccessExpiration to be an explicit nil.
func (o *SharedDashboardInviteesItems) SetAccessExpirationNil() {
	o.AccessExpiration.Set(nil)
}

// UnsetAccessExpiration ensures that no value is present for AccessExpiration, not even an explicit nil.
func (o *SharedDashboardInviteesItems) UnsetAccessExpiration() {
	o.AccessExpiration.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SharedDashboardInviteesItems) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardInviteesItems) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SharedDashboardInviteesItems) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SharedDashboardInviteesItems) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value.
func (o *SharedDashboardInviteesItems) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardInviteesItems) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *SharedDashboardInviteesItems) SetEmail(v string) {
	o.Email = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SharedDashboardInviteesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccessExpiration.IsSet() {
		toSerialize["access_expiration"] = o.AccessExpiration.Get()
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["email"] = o.Email

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SharedDashboardInviteesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessExpiration datadog.NullableTime `json:"access_expiration,omitempty"`
		CreatedAt        *time.Time           `json:"created_at,omitempty"`
		Email            *string              `json:"email"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_expiration", "created_at", "email"})
	} else {
		return err
	}
	o.AccessExpiration = all.AccessExpiration
	o.CreatedAt = all.CreatedAt
	o.Email = *all.Email

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
