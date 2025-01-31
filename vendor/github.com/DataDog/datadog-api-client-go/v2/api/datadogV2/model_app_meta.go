// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppMeta The definition of `AppMeta` object.
type AppMeta struct {
	// The `AppMeta` `created_at`.
	CreatedAt *string `json:"created_at,omitempty"`
	// The `AppMeta` `deleted_at`.
	DeletedAt *string `json:"deleted_at,omitempty"`
	// The `AppMeta` `org_id`.
	OrgId *int64 `json:"org_id,omitempty"`
	// The `AppMeta` `run_as_user`.
	RunAsUser *string `json:"run_as_user,omitempty"`
	// The `AppMeta` `updated_at`.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// The `AppMeta` `updated_since_deployment`.
	UpdatedSinceDeployment *bool `json:"updated_since_deployment,omitempty"`
	// The `AppMeta` `user_id`.
	UserId *int64 `json:"user_id,omitempty"`
	// The `AppMeta` `user_name`.
	UserName *string `json:"user_name,omitempty"`
	// The `AppMeta` `user_uuid`.
	UserUuid *string `json:"user_uuid,omitempty"`
	// The `AppMeta` `version`.
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
func (o *AppMeta) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AppMeta) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AppMeta) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AppMeta) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AppMeta) HasDeletedAt() bool {
	return o != nil && o.DeletedAt != nil
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AppMeta) SetDeletedAt(v string) {
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

// GetRunAsUser returns the RunAsUser field value if set, zero value otherwise.
func (o *AppMeta) GetRunAsUser() string {
	if o == nil || o.RunAsUser == nil {
		var ret string
		return ret
	}
	return *o.RunAsUser
}

// GetRunAsUserOk returns a tuple with the RunAsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetRunAsUserOk() (*string, bool) {
	if o == nil || o.RunAsUser == nil {
		return nil, false
	}
	return o.RunAsUser, true
}

// HasRunAsUser returns a boolean if a field has been set.
func (o *AppMeta) HasRunAsUser() bool {
	return o != nil && o.RunAsUser != nil
}

// SetRunAsUser gets a reference to the given string and assigns it to the RunAsUser field.
func (o *AppMeta) SetRunAsUser(v string) {
	o.RunAsUser = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AppMeta) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AppMeta) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AppMeta) SetUpdatedAt(v string) {
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
func (o *AppMeta) GetUserUuid() string {
	if o == nil || o.UserUuid == nil {
		var ret string
		return ret
	}
	return *o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMeta) GetUserUuidOk() (*string, bool) {
	if o == nil || o.UserUuid == nil {
		return nil, false
	}
	return o.UserUuid, true
}

// HasUserUuid returns a boolean if a field has been set.
func (o *AppMeta) HasUserUuid() bool {
	return o != nil && o.UserUuid != nil
}

// SetUserUuid gets a reference to the given string and assigns it to the UserUuid field.
func (o *AppMeta) SetUserUuid(v string) {
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
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.RunAsUser != nil {
		toSerialize["run_as_user"] = o.RunAsUser
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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
		CreatedAt              *string `json:"created_at,omitempty"`
		DeletedAt              *string `json:"deleted_at,omitempty"`
		OrgId                  *int64  `json:"org_id,omitempty"`
		RunAsUser              *string `json:"run_as_user,omitempty"`
		UpdatedAt              *string `json:"updated_at,omitempty"`
		UpdatedSinceDeployment *bool   `json:"updated_since_deployment,omitempty"`
		UserId                 *int64  `json:"user_id,omitempty"`
		UserName               *string `json:"user_name,omitempty"`
		UserUuid               *string `json:"user_uuid,omitempty"`
		Version                *int64  `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "deleted_at", "org_id", "run_as_user", "updated_at", "updated_since_deployment", "user_id", "user_name", "user_uuid", "version"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.DeletedAt = all.DeletedAt
	o.OrgId = all.OrgId
	o.RunAsUser = all.RunAsUser
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
