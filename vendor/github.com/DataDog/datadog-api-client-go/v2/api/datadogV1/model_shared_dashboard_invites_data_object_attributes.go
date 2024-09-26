// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardInvitesDataObjectAttributes Attributes of the shared dashboard invitation
type SharedDashboardInvitesDataObjectAttributes struct {
	// When the invitation was sent.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// An email address that an invitation has been (or if used in invitation request, will be) sent to.
	Email *string `json:"email,omitempty"`
	// Indicates whether an active session exists for the invitation (produced when a user clicks the link in the email).
	HasSession *bool `json:"has_session,omitempty"`
	// When the invitation expires.
	InvitationExpiry *time.Time `json:"invitation_expiry,omitempty"`
	// When the invited user's session expires. null if the invitation has no associated session.
	SessionExpiry datadog.NullableTime `json:"session_expiry,omitempty"`
	// The unique token of the shared dashboard that was (or is to be) shared.
	ShareToken *string `json:"share_token,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSharedDashboardInvitesDataObjectAttributes instantiates a new SharedDashboardInvitesDataObjectAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSharedDashboardInvitesDataObjectAttributes() *SharedDashboardInvitesDataObjectAttributes {
	this := SharedDashboardInvitesDataObjectAttributes{}
	return &this
}

// NewSharedDashboardInvitesDataObjectAttributesWithDefaults instantiates a new SharedDashboardInvitesDataObjectAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSharedDashboardInvitesDataObjectAttributesWithDefaults() *SharedDashboardInvitesDataObjectAttributes {
	this := SharedDashboardInvitesDataObjectAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SharedDashboardInvitesDataObjectAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardInvitesDataObjectAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SharedDashboardInvitesDataObjectAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SharedDashboardInvitesDataObjectAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SharedDashboardInvitesDataObjectAttributes) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardInvitesDataObjectAttributes) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SharedDashboardInvitesDataObjectAttributes) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SharedDashboardInvitesDataObjectAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetHasSession returns the HasSession field value if set, zero value otherwise.
func (o *SharedDashboardInvitesDataObjectAttributes) GetHasSession() bool {
	if o == nil || o.HasSession == nil {
		var ret bool
		return ret
	}
	return *o.HasSession
}

// GetHasSessionOk returns a tuple with the HasSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardInvitesDataObjectAttributes) GetHasSessionOk() (*bool, bool) {
	if o == nil || o.HasSession == nil {
		return nil, false
	}
	return o.HasSession, true
}

// HasHasSession returns a boolean if a field has been set.
func (o *SharedDashboardInvitesDataObjectAttributes) HasHasSession() bool {
	return o != nil && o.HasSession != nil
}

// SetHasSession gets a reference to the given bool and assigns it to the HasSession field.
func (o *SharedDashboardInvitesDataObjectAttributes) SetHasSession(v bool) {
	o.HasSession = &v
}

// GetInvitationExpiry returns the InvitationExpiry field value if set, zero value otherwise.
func (o *SharedDashboardInvitesDataObjectAttributes) GetInvitationExpiry() time.Time {
	if o == nil || o.InvitationExpiry == nil {
		var ret time.Time
		return ret
	}
	return *o.InvitationExpiry
}

// GetInvitationExpiryOk returns a tuple with the InvitationExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardInvitesDataObjectAttributes) GetInvitationExpiryOk() (*time.Time, bool) {
	if o == nil || o.InvitationExpiry == nil {
		return nil, false
	}
	return o.InvitationExpiry, true
}

// HasInvitationExpiry returns a boolean if a field has been set.
func (o *SharedDashboardInvitesDataObjectAttributes) HasInvitationExpiry() bool {
	return o != nil && o.InvitationExpiry != nil
}

// SetInvitationExpiry gets a reference to the given time.Time and assigns it to the InvitationExpiry field.
func (o *SharedDashboardInvitesDataObjectAttributes) SetInvitationExpiry(v time.Time) {
	o.InvitationExpiry = &v
}

// GetSessionExpiry returns the SessionExpiry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedDashboardInvitesDataObjectAttributes) GetSessionExpiry() time.Time {
	if o == nil || o.SessionExpiry.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SessionExpiry.Get()
}

// GetSessionExpiryOk returns a tuple with the SessionExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardInvitesDataObjectAttributes) GetSessionExpiryOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionExpiry.Get(), o.SessionExpiry.IsSet()
}

// HasSessionExpiry returns a boolean if a field has been set.
func (o *SharedDashboardInvitesDataObjectAttributes) HasSessionExpiry() bool {
	return o != nil && o.SessionExpiry.IsSet()
}

// SetSessionExpiry gets a reference to the given datadog.NullableTime and assigns it to the SessionExpiry field.
func (o *SharedDashboardInvitesDataObjectAttributes) SetSessionExpiry(v time.Time) {
	o.SessionExpiry.Set(&v)
}

// SetSessionExpiryNil sets the value for SessionExpiry to be an explicit nil.
func (o *SharedDashboardInvitesDataObjectAttributes) SetSessionExpiryNil() {
	o.SessionExpiry.Set(nil)
}

// UnsetSessionExpiry ensures that no value is present for SessionExpiry, not even an explicit nil.
func (o *SharedDashboardInvitesDataObjectAttributes) UnsetSessionExpiry() {
	o.SessionExpiry.Unset()
}

// GetShareToken returns the ShareToken field value if set, zero value otherwise.
func (o *SharedDashboardInvitesDataObjectAttributes) GetShareToken() string {
	if o == nil || o.ShareToken == nil {
		var ret string
		return ret
	}
	return *o.ShareToken
}

// GetShareTokenOk returns a tuple with the ShareToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardInvitesDataObjectAttributes) GetShareTokenOk() (*string, bool) {
	if o == nil || o.ShareToken == nil {
		return nil, false
	}
	return o.ShareToken, true
}

// HasShareToken returns a boolean if a field has been set.
func (o *SharedDashboardInvitesDataObjectAttributes) HasShareToken() bool {
	return o != nil && o.ShareToken != nil
}

// SetShareToken gets a reference to the given string and assigns it to the ShareToken field.
func (o *SharedDashboardInvitesDataObjectAttributes) SetShareToken(v string) {
	o.ShareToken = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SharedDashboardInvitesDataObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.HasSession != nil {
		toSerialize["has_session"] = o.HasSession
	}
	if o.InvitationExpiry != nil {
		if o.InvitationExpiry.Nanosecond() == 0 {
			toSerialize["invitation_expiry"] = o.InvitationExpiry.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["invitation_expiry"] = o.InvitationExpiry.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.SessionExpiry.IsSet() {
		toSerialize["session_expiry"] = o.SessionExpiry.Get()
	}
	if o.ShareToken != nil {
		toSerialize["share_token"] = o.ShareToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SharedDashboardInvitesDataObjectAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt        *time.Time           `json:"created_at,omitempty"`
		Email            *string              `json:"email,omitempty"`
		HasSession       *bool                `json:"has_session,omitempty"`
		InvitationExpiry *time.Time           `json:"invitation_expiry,omitempty"`
		SessionExpiry    datadog.NullableTime `json:"session_expiry,omitempty"`
		ShareToken       *string              `json:"share_token,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "email", "has_session", "invitation_expiry", "session_expiry", "share_token"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.Email = all.Email
	o.HasSession = all.HasSession
	o.InvitationExpiry = all.InvitationExpiry
	o.SessionExpiry = all.SessionExpiry
	o.ShareToken = all.ShareToken

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
