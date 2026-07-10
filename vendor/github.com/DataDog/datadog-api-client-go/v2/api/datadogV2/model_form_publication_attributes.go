// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormPublicationAttributes The attributes of a form publication.
type FormPublicationAttributes struct {
	// The time at which the publication was created.
	CreatedAt time.Time `json:"created_at"`
	// The ID of the form.
	FormId uuid.UUID `json:"form_id"`
	// The version number that was published.
	FormVersion int64 `json:"form_version"`
	// The ID of the form publication.
	Id *string `json:"id,omitempty"`
	// The time at which the publication was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The ID of the organization that owns this publication.
	OrgId int64 `json:"org_id"`
	// The sequential publication number for this form.
	PublishSeq int64 `json:"publish_seq"`
	// The ID of the user who created this publication.
	UserId int64 `json:"user_id"`
	// The UUID of the user who created this publication.
	UserUuid uuid.UUID `json:"user_uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormPublicationAttributes instantiates a new FormPublicationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormPublicationAttributes(createdAt time.Time, formId uuid.UUID, formVersion int64, modifiedAt time.Time, orgId int64, publishSeq int64, userId int64, userUuid uuid.UUID) *FormPublicationAttributes {
	this := FormPublicationAttributes{}
	this.CreatedAt = createdAt
	this.FormId = formId
	this.FormVersion = formVersion
	this.ModifiedAt = modifiedAt
	this.OrgId = orgId
	this.PublishSeq = publishSeq
	this.UserId = userId
	this.UserUuid = userUuid
	return &this
}

// NewFormPublicationAttributesWithDefaults instantiates a new FormPublicationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormPublicationAttributesWithDefaults() *FormPublicationAttributes {
	this := FormPublicationAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *FormPublicationAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FormPublicationAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *FormPublicationAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetFormId returns the FormId field value.
func (o *FormPublicationAttributes) GetFormId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value
// and a boolean to check if the value has been set.
func (o *FormPublicationAttributes) GetFormIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormId, true
}

// SetFormId sets field value.
func (o *FormPublicationAttributes) SetFormId(v uuid.UUID) {
	o.FormId = v
}

// GetFormVersion returns the FormVersion field value.
func (o *FormPublicationAttributes) GetFormVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FormVersion
}

// GetFormVersionOk returns a tuple with the FormVersion field value
// and a boolean to check if the value has been set.
func (o *FormPublicationAttributes) GetFormVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormVersion, true
}

// SetFormVersion sets field value.
func (o *FormPublicationAttributes) SetFormVersion(v int64) {
	o.FormVersion = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FormPublicationAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormPublicationAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FormPublicationAttributes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FormPublicationAttributes) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *FormPublicationAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *FormPublicationAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *FormPublicationAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetOrgId returns the OrgId field value.
func (o *FormPublicationAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *FormPublicationAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *FormPublicationAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetPublishSeq returns the PublishSeq field value.
func (o *FormPublicationAttributes) GetPublishSeq() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PublishSeq
}

// GetPublishSeqOk returns a tuple with the PublishSeq field value
// and a boolean to check if the value has been set.
func (o *FormPublicationAttributes) GetPublishSeqOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishSeq, true
}

// SetPublishSeq sets field value.
func (o *FormPublicationAttributes) SetPublishSeq(v int64) {
	o.PublishSeq = v
}

// GetUserId returns the UserId field value.
func (o *FormPublicationAttributes) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FormPublicationAttributes) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *FormPublicationAttributes) SetUserId(v int64) {
	o.UserId = v
}

// GetUserUuid returns the UserUuid field value.
func (o *FormPublicationAttributes) GetUserUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value
// and a boolean to check if the value has been set.
func (o *FormPublicationAttributes) GetUserUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuid, true
}

// SetUserUuid sets field value.
func (o *FormPublicationAttributes) SetUserUuid(v uuid.UUID) {
	o.UserUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormPublicationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["form_id"] = o.FormId
	toSerialize["form_version"] = o.FormVersion
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["org_id"] = o.OrgId
	toSerialize["publish_seq"] = o.PublishSeq
	toSerialize["user_id"] = o.UserId
	toSerialize["user_uuid"] = o.UserUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormPublicationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time `json:"created_at"`
		FormId      *uuid.UUID `json:"form_id"`
		FormVersion *int64     `json:"form_version"`
		Id          *string    `json:"id,omitempty"`
		ModifiedAt  *time.Time `json:"modified_at"`
		OrgId       *int64     `json:"org_id"`
		PublishSeq  *int64     `json:"publish_seq"`
		UserId      *int64     `json:"user_id"`
		UserUuid    *uuid.UUID `json:"user_uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.FormId == nil {
		return fmt.Errorf("required field form_id missing")
	}
	if all.FormVersion == nil {
		return fmt.Errorf("required field form_version missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.PublishSeq == nil {
		return fmt.Errorf("required field publish_seq missing")
	}
	if all.UserId == nil {
		return fmt.Errorf("required field user_id missing")
	}
	if all.UserUuid == nil {
		return fmt.Errorf("required field user_uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "form_id", "form_version", "id", "modified_at", "org_id", "publish_seq", "user_id", "user_uuid"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.FormId = *all.FormId
	o.FormVersion = *all.FormVersion
	o.Id = all.Id
	o.ModifiedAt = *all.ModifiedAt
	o.OrgId = *all.OrgId
	o.PublishSeq = *all.PublishSeq
	o.UserId = *all.UserId
	o.UserUuid = *all.UserUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
