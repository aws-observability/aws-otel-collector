// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppVersionAttributes Attributes describing an app version.
type AppVersionAttributes struct {
	// The ID of the app this version belongs to.
	AppId *uuid.UUID `json:"app_id,omitempty"`
	// Timestamp of when the version was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Whether this version has ever been published.
	HasEverBeenPublished *bool `json:"has_ever_been_published,omitempty"`
	// The optional human-readable name of the version.
	Name *string `json:"name,omitempty"`
	// Timestamp of when the version was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The ID of the user who created the version.
	UserId *int64 `json:"user_id,omitempty"`
	// The name (or email) of the user who created the version.
	UserName *string `json:"user_name,omitempty"`
	// The UUID of the user who created the version.
	UserUuid *uuid.UUID `json:"user_uuid,omitempty"`
	// The version number of the app, starting at 1.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAppVersionAttributes instantiates a new AppVersionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAppVersionAttributes() *AppVersionAttributes {
	this := AppVersionAttributes{}
	return &this
}

// NewAppVersionAttributesWithDefaults instantiates a new AppVersionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAppVersionAttributesWithDefaults() *AppVersionAttributes {
	this := AppVersionAttributes{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *AppVersionAttributes) GetAppId() uuid.UUID {
	if o == nil || o.AppId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionAttributes) GetAppIdOk() (*uuid.UUID, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *AppVersionAttributes) HasAppId() bool {
	return o != nil && o.AppId != nil
}

// SetAppId gets a reference to the given uuid.UUID and assigns it to the AppId field.
func (o *AppVersionAttributes) SetAppId(v uuid.UUID) {
	o.AppId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AppVersionAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AppVersionAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AppVersionAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHasEverBeenPublished returns the HasEverBeenPublished field value if set, zero value otherwise.
func (o *AppVersionAttributes) GetHasEverBeenPublished() bool {
	if o == nil || o.HasEverBeenPublished == nil {
		var ret bool
		return ret
	}
	return *o.HasEverBeenPublished
}

// GetHasEverBeenPublishedOk returns a tuple with the HasEverBeenPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionAttributes) GetHasEverBeenPublishedOk() (*bool, bool) {
	if o == nil || o.HasEverBeenPublished == nil {
		return nil, false
	}
	return o.HasEverBeenPublished, true
}

// HasHasEverBeenPublished returns a boolean if a field has been set.
func (o *AppVersionAttributes) HasHasEverBeenPublished() bool {
	return o != nil && o.HasEverBeenPublished != nil
}

// SetHasEverBeenPublished gets a reference to the given bool and assigns it to the HasEverBeenPublished field.
func (o *AppVersionAttributes) SetHasEverBeenPublished(v bool) {
	o.HasEverBeenPublished = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppVersionAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppVersionAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppVersionAttributes) SetName(v string) {
	o.Name = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AppVersionAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AppVersionAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AppVersionAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AppVersionAttributes) GetUserId() int64 {
	if o == nil || o.UserId == nil {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionAttributes) GetUserIdOk() (*int64, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AppVersionAttributes) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *AppVersionAttributes) SetUserId(v int64) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *AppVersionAttributes) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionAttributes) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *AppVersionAttributes) HasUserName() bool {
	return o != nil && o.UserName != nil
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *AppVersionAttributes) SetUserName(v string) {
	o.UserName = &v
}

// GetUserUuid returns the UserUuid field value if set, zero value otherwise.
func (o *AppVersionAttributes) GetUserUuid() uuid.UUID {
	if o == nil || o.UserUuid == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionAttributes) GetUserUuidOk() (*uuid.UUID, bool) {
	if o == nil || o.UserUuid == nil {
		return nil, false
	}
	return o.UserUuid, true
}

// HasUserUuid returns a boolean if a field has been set.
func (o *AppVersionAttributes) HasUserUuid() bool {
	return o != nil && o.UserUuid != nil
}

// SetUserUuid gets a reference to the given uuid.UUID and assigns it to the UserUuid field.
func (o *AppVersionAttributes) SetUserUuid(v uuid.UUID) {
	o.UserUuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AppVersionAttributes) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionAttributes) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AppVersionAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *AppVersionAttributes) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AppVersionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.HasEverBeenPublished != nil {
		toSerialize["has_ever_been_published"] = o.HasEverBeenPublished
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.UserName != nil {
		toSerialize["user_name"] = o.UserName
	}
	if o.UserUuid != nil {
		toSerialize["user_uuid"] = o.UserUuid
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AppVersionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppId                *uuid.UUID `json:"app_id,omitempty"`
		CreatedAt            *time.Time `json:"created_at,omitempty"`
		HasEverBeenPublished *bool      `json:"has_ever_been_published,omitempty"`
		Name                 *string    `json:"name,omitempty"`
		UpdatedAt            *time.Time `json:"updated_at,omitempty"`
		UserId               *int64     `json:"user_id,omitempty"`
		UserName             *string    `json:"user_name,omitempty"`
		UserUuid             *uuid.UUID `json:"user_uuid,omitempty"`
		Version              *int64     `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"app_id", "created_at", "has_ever_been_published", "name", "updated_at", "user_id", "user_name", "user_uuid", "version"})
	} else {
		return err
	}
	o.AppId = all.AppId
	o.CreatedAt = all.CreatedAt
	o.HasEverBeenPublished = all.HasEverBeenPublished
	o.Name = all.Name
	o.UpdatedAt = all.UpdatedAt
	o.UserId = all.UserId
	o.UserName = all.UserName
	o.UserUuid = all.UserUuid
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
