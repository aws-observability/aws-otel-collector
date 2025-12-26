// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RestrictionQueryAttributes Attributes of the restriction query.
type RestrictionQueryAttributes struct {
	// Creation time of the restriction query.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Email of the user who last modified this restriction query.
	LastModifierEmail *string `json:"last_modifier_email,omitempty"`
	// Name of the user who last modified this restriction query.
	LastModifierName *string `json:"last_modifier_name,omitempty"`
	// Time of last restriction query modification.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The query that defines the restriction. Only the content matching the query can be returned.
	RestrictionQuery *string `json:"restriction_query,omitempty"`
	// Number of roles associated with this restriction query.
	RoleCount *int64 `json:"role_count,omitempty"`
	// Number of users associated with this restriction query.
	UserCount *int64 `json:"user_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestrictionQueryAttributes instantiates a new RestrictionQueryAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestrictionQueryAttributes() *RestrictionQueryAttributes {
	this := RestrictionQueryAttributes{}
	return &this
}

// NewRestrictionQueryAttributesWithDefaults instantiates a new RestrictionQueryAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestrictionQueryAttributesWithDefaults() *RestrictionQueryAttributes {
	this := RestrictionQueryAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RestrictionQueryAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RestrictionQueryAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RestrictionQueryAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastModifierEmail returns the LastModifierEmail field value if set, zero value otherwise.
func (o *RestrictionQueryAttributes) GetLastModifierEmail() string {
	if o == nil || o.LastModifierEmail == nil {
		var ret string
		return ret
	}
	return *o.LastModifierEmail
}

// GetLastModifierEmailOk returns a tuple with the LastModifierEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryAttributes) GetLastModifierEmailOk() (*string, bool) {
	if o == nil || o.LastModifierEmail == nil {
		return nil, false
	}
	return o.LastModifierEmail, true
}

// HasLastModifierEmail returns a boolean if a field has been set.
func (o *RestrictionQueryAttributes) HasLastModifierEmail() bool {
	return o != nil && o.LastModifierEmail != nil
}

// SetLastModifierEmail gets a reference to the given string and assigns it to the LastModifierEmail field.
func (o *RestrictionQueryAttributes) SetLastModifierEmail(v string) {
	o.LastModifierEmail = &v
}

// GetLastModifierName returns the LastModifierName field value if set, zero value otherwise.
func (o *RestrictionQueryAttributes) GetLastModifierName() string {
	if o == nil || o.LastModifierName == nil {
		var ret string
		return ret
	}
	return *o.LastModifierName
}

// GetLastModifierNameOk returns a tuple with the LastModifierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryAttributes) GetLastModifierNameOk() (*string, bool) {
	if o == nil || o.LastModifierName == nil {
		return nil, false
	}
	return o.LastModifierName, true
}

// HasLastModifierName returns a boolean if a field has been set.
func (o *RestrictionQueryAttributes) HasLastModifierName() bool {
	return o != nil && o.LastModifierName != nil
}

// SetLastModifierName gets a reference to the given string and assigns it to the LastModifierName field.
func (o *RestrictionQueryAttributes) SetLastModifierName(v string) {
	o.LastModifierName = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *RestrictionQueryAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *RestrictionQueryAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *RestrictionQueryAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetRestrictionQuery returns the RestrictionQuery field value if set, zero value otherwise.
func (o *RestrictionQueryAttributes) GetRestrictionQuery() string {
	if o == nil || o.RestrictionQuery == nil {
		var ret string
		return ret
	}
	return *o.RestrictionQuery
}

// GetRestrictionQueryOk returns a tuple with the RestrictionQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryAttributes) GetRestrictionQueryOk() (*string, bool) {
	if o == nil || o.RestrictionQuery == nil {
		return nil, false
	}
	return o.RestrictionQuery, true
}

// HasRestrictionQuery returns a boolean if a field has been set.
func (o *RestrictionQueryAttributes) HasRestrictionQuery() bool {
	return o != nil && o.RestrictionQuery != nil
}

// SetRestrictionQuery gets a reference to the given string and assigns it to the RestrictionQuery field.
func (o *RestrictionQueryAttributes) SetRestrictionQuery(v string) {
	o.RestrictionQuery = &v
}

// GetRoleCount returns the RoleCount field value if set, zero value otherwise.
func (o *RestrictionQueryAttributes) GetRoleCount() int64 {
	if o == nil || o.RoleCount == nil {
		var ret int64
		return ret
	}
	return *o.RoleCount
}

// GetRoleCountOk returns a tuple with the RoleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryAttributes) GetRoleCountOk() (*int64, bool) {
	if o == nil || o.RoleCount == nil {
		return nil, false
	}
	return o.RoleCount, true
}

// HasRoleCount returns a boolean if a field has been set.
func (o *RestrictionQueryAttributes) HasRoleCount() bool {
	return o != nil && o.RoleCount != nil
}

// SetRoleCount gets a reference to the given int64 and assigns it to the RoleCount field.
func (o *RestrictionQueryAttributes) SetRoleCount(v int64) {
	o.RoleCount = &v
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *RestrictionQueryAttributes) GetUserCount() int64 {
	if o == nil || o.UserCount == nil {
		var ret int64
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryAttributes) GetUserCountOk() (*int64, bool) {
	if o == nil || o.UserCount == nil {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *RestrictionQueryAttributes) HasUserCount() bool {
	return o != nil && o.UserCount != nil
}

// SetUserCount gets a reference to the given int64 and assigns it to the UserCount field.
func (o *RestrictionQueryAttributes) SetUserCount(v int64) {
	o.UserCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestrictionQueryAttributes) MarshalJSON() ([]byte, error) {
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
	if o.LastModifierEmail != nil {
		toSerialize["last_modifier_email"] = o.LastModifierEmail
	}
	if o.LastModifierName != nil {
		toSerialize["last_modifier_name"] = o.LastModifierName
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.RestrictionQuery != nil {
		toSerialize["restriction_query"] = o.RestrictionQuery
	}
	if o.RoleCount != nil {
		toSerialize["role_count"] = o.RoleCount
	}
	if o.UserCount != nil {
		toSerialize["user_count"] = o.UserCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestrictionQueryAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt         *time.Time `json:"created_at,omitempty"`
		LastModifierEmail *string    `json:"last_modifier_email,omitempty"`
		LastModifierName  *string    `json:"last_modifier_name,omitempty"`
		ModifiedAt        *time.Time `json:"modified_at,omitempty"`
		RestrictionQuery  *string    `json:"restriction_query,omitempty"`
		RoleCount         *int64     `json:"role_count,omitempty"`
		UserCount         *int64     `json:"user_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "last_modifier_email", "last_modifier_name", "modified_at", "restriction_query", "role_count", "user_count"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.LastModifierEmail = all.LastModifierEmail
	o.LastModifierName = all.LastModifierName
	o.ModifiedAt = all.ModifiedAt
	o.RestrictionQuery = all.RestrictionQuery
	o.RoleCount = all.RoleCount
	o.UserCount = all.UserCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
