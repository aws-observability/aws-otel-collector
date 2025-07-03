// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppMeta Metadata of an app.
type AppMeta struct {
	// Timestamp of when the app was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp of when the app was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The Datadog organization ID that owns the app.
	OrgId *int64 `json:"org_id,omitempty"`
	// Timestamp of when the app was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Whether the app was updated since it was last published. Published apps are pinned to a specific version and do not automatically update when the app is updated.
	UpdatedSinceDeployment *bool `json:"updated_since_deployment,omitempty"`
	// The ID of the user who created the app.
	UserId *int64 `json:"user_id,omitempty"`
	// The name (or email address) of the user who created the app.
	UserName *string `json:"user_name,omitempty"`
	// The UUID of the user who created the app.
	UserUuid *uuid.UUID `json:"user_uuid,omitempty"`
	// The version number of the app. This starts at 1 and increments with each update.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAppMeta instantiates a new AppMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAppMeta() *AppMeta {
	this := AppMeta{}
	return &this
}

// NewAppMetaWithDefaults instantiates a new AppMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAppMetaWithDefaults() *AppMeta {
	this := AppMeta{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AppMeta) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AppMeta) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AppMeta) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AppMeta) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AppMeta) HasDeletedAt() bool {
	return o != nil && o.DeletedAt != nil
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *AppMeta) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *AppMeta) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *AppMeta) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *AppMeta) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AppMeta) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AppMeta) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AppMeta) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedSinceDeployment returns the UpdatedSinceDeployment field value if set, zero value otherwise.
func (o *AppMeta) GetUpdatedSinceDeployment() bool {
	if o == nil || o.UpdatedSinceDeployment == nil {
		var ret bool
		return ret
	}
	return *o.UpdatedSinceDeployment
}

// GetUpdatedSinceDeploymentOk returns a tuple with the UpdatedSinceDeployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetUpdatedSinceDeploymentOk() (*bool, bool) {
	if o == nil || o.UpdatedSinceDeployment == nil {
		return nil, false
	}
	return o.UpdatedSinceDeployment, true
}

// HasUpdatedSinceDeployment returns a boolean if a field has been set.
func (o *AppMeta) HasUpdatedSinceDeployment() bool {
	return o != nil && o.UpdatedSinceDeployment != nil
}

// SetUpdatedSinceDeployment gets a reference to the given bool and assigns it to the UpdatedSinceDeployment field.
func (o *AppMeta) SetUpdatedSinceDeployment(v bool) {
	o.UpdatedSinceDeployment = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AppMeta) GetUserId() int64 {
	if o == nil || o.UserId == nil {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetUserIdOk() (*int64, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AppMeta) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *AppMeta) SetUserId(v int64) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *AppMeta) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *AppMeta) HasUserName() bool {
	return o != nil && o.UserName != nil
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *AppMeta) SetUserName(v string) {
	o.UserName = &v
}

// GetUserUuid returns the UserUuid field value if set, zero value otherwise.
func (o *AppMeta) GetUserUuid() uuid.UUID {
	if o == nil || o.UserUuid == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetUserUuidOk() (*uuid.UUID, bool) {
	if o == nil || o.UserUuid == nil {
		return nil, false
	}
	return o.UserUuid, true
}

// HasUserUuid returns a boolean if a field has been set.
func (o *AppMeta) HasUserUuid() bool {
	return o != nil && o.UserUuid != nil
}

// SetUserUuid gets a reference to the given uuid.UUID and assigns it to the UserUuid field.
func (o *AppMeta) SetUserUuid(v uuid.UUID) {
	o.UserUuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AppMeta) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AppMeta) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *AppMeta) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AppMeta) MarshalJSON() ([]byte, error) {
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
	if o.DeletedAt != nil {
		if o.DeletedAt.Nanosecond() == 0 {
			toSerialize["deleted_at"] = o.DeletedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["deleted_at"] = o.DeletedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedSinceDeployment != nil {
		toSerialize["updated_since_deployment"] = o.UpdatedSinceDeployment
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
func (o *AppMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt              *time.Time `json:"created_at,omitempty"`
		DeletedAt              *time.Time `json:"deleted_at,omitempty"`
		OrgId                  *int64     `json:"org_id,omitempty"`
		UpdatedAt              *time.Time `json:"updated_at,omitempty"`
		UpdatedSinceDeployment *bool      `json:"updated_since_deployment,omitempty"`
		UserId                 *int64     `json:"user_id,omitempty"`
		UserName               *string    `json:"user_name,omitempty"`
		UserUuid               *uuid.UUID `json:"user_uuid,omitempty"`
		Version                *int64     `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "deleted_at", "org_id", "updated_at", "updated_since_deployment", "user_id", "user_name", "user_uuid", "version"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.DeletedAt = all.DeletedAt
	o.OrgId = all.OrgId
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedSinceDeployment = all.UpdatedSinceDeployment
	o.UserId = all.UserId
	o.UserName = all.UserName
	o.UserUuid = all.UserUuid
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
