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

// FormVersionAttributes The attributes of a form version.
type FormVersionAttributes struct {
	// The time at which the version was created.
	CreatedAt time.Time `json:"created_at"`
	// A JSON Schema definition that describes the form's data fields.
	DataDefinition FormDataDefinition `json:"data_definition"`
	// The signature of the version definition.
	DefinitionSignature string `json:"definition_signature"`
	// The ETag for optimistic concurrency control.
	Etag datadog.NullableString `json:"etag"`
	// The ID of the form version.
	Id *string `json:"id,omitempty"`
	// The time at which the version was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The state of a form version.
	State FormVersionState `json:"state"`
	// UI configuration for rendering form fields, including widget overrides, field ordering, and themes.
	UiDefinition FormUiDefinition `json:"ui_definition"`
	// The ID of the user who created this version.
	UserId int64 `json:"user_id"`
	// The UUID of the user who created this version.
	UserUuid uuid.UUID `json:"user_uuid"`
	// The sequential version number.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormVersionAttributes instantiates a new FormVersionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormVersionAttributes(createdAt time.Time, dataDefinition FormDataDefinition, definitionSignature string, etag datadog.NullableString, modifiedAt time.Time, state FormVersionState, uiDefinition FormUiDefinition, userId int64, userUuid uuid.UUID, version int64) *FormVersionAttributes {
	this := FormVersionAttributes{}
	this.CreatedAt = createdAt
	this.DataDefinition = dataDefinition
	this.DefinitionSignature = definitionSignature
	this.Etag = etag
	this.ModifiedAt = modifiedAt
	this.State = state
	this.UiDefinition = uiDefinition
	this.UserId = userId
	this.UserUuid = userUuid
	this.Version = version
	return &this
}

// NewFormVersionAttributesWithDefaults instantiates a new FormVersionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormVersionAttributesWithDefaults() *FormVersionAttributes {
	this := FormVersionAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *FormVersionAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *FormVersionAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDataDefinition returns the DataDefinition field value.
func (o *FormVersionAttributes) GetDataDefinition() FormDataDefinition {
	if o == nil {
		var ret FormDataDefinition
		return ret
	}
	return o.DataDefinition
}

// GetDataDefinitionOk returns a tuple with the DataDefinition field value
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetDataDefinitionOk() (*FormDataDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataDefinition, true
}

// SetDataDefinition sets field value.
func (o *FormVersionAttributes) SetDataDefinition(v FormDataDefinition) {
	o.DataDefinition = v
}

// GetDefinitionSignature returns the DefinitionSignature field value.
func (o *FormVersionAttributes) GetDefinitionSignature() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DefinitionSignature
}

// GetDefinitionSignatureOk returns a tuple with the DefinitionSignature field value
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetDefinitionSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefinitionSignature, true
}

// SetDefinitionSignature sets field value.
func (o *FormVersionAttributes) SetDefinitionSignature(v string) {
	o.DefinitionSignature = v
}

// GetEtag returns the Etag field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *FormVersionAttributes) GetEtag() string {
	if o == nil || o.Etag.Get() == nil {
		var ret string
		return ret
	}
	return *o.Etag.Get()
}

// GetEtagOk returns a tuple with the Etag field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FormVersionAttributes) GetEtagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Etag.Get(), o.Etag.IsSet()
}

// SetEtag sets field value.
func (o *FormVersionAttributes) SetEtag(v string) {
	o.Etag.Set(&v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FormVersionAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FormVersionAttributes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FormVersionAttributes) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *FormVersionAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *FormVersionAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetState returns the State field value.
func (o *FormVersionAttributes) GetState() FormVersionState {
	if o == nil {
		var ret FormVersionState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetStateOk() (*FormVersionState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *FormVersionAttributes) SetState(v FormVersionState) {
	o.State = v
}

// GetUiDefinition returns the UiDefinition field value.
func (o *FormVersionAttributes) GetUiDefinition() FormUiDefinition {
	if o == nil {
		var ret FormUiDefinition
		return ret
	}
	return o.UiDefinition
}

// GetUiDefinitionOk returns a tuple with the UiDefinition field value
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetUiDefinitionOk() (*FormUiDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UiDefinition, true
}

// SetUiDefinition sets field value.
func (o *FormVersionAttributes) SetUiDefinition(v FormUiDefinition) {
	o.UiDefinition = v
}

// GetUserId returns the UserId field value.
func (o *FormVersionAttributes) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *FormVersionAttributes) SetUserId(v int64) {
	o.UserId = v
}

// GetUserUuid returns the UserUuid field value.
func (o *FormVersionAttributes) GetUserUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetUserUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuid, true
}

// SetUserUuid sets field value.
func (o *FormVersionAttributes) SetUserUuid(v uuid.UUID) {
	o.UserUuid = v
}

// GetVersion returns the Version field value.
func (o *FormVersionAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *FormVersionAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormVersionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["data_definition"] = o.DataDefinition
	toSerialize["definition_signature"] = o.DefinitionSignature
	toSerialize["etag"] = o.Etag.Get()
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["state"] = o.State
	toSerialize["ui_definition"] = o.UiDefinition
	toSerialize["user_id"] = o.UserId
	toSerialize["user_uuid"] = o.UserUuid
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormVersionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt           *time.Time             `json:"created_at"`
		DataDefinition      *FormDataDefinition    `json:"data_definition"`
		DefinitionSignature *string                `json:"definition_signature"`
		Etag                datadog.NullableString `json:"etag"`
		Id                  *string                `json:"id,omitempty"`
		ModifiedAt          *time.Time             `json:"modified_at"`
		State               *FormVersionState      `json:"state"`
		UiDefinition        *FormUiDefinition      `json:"ui_definition"`
		UserId              *int64                 `json:"user_id"`
		UserUuid            *uuid.UUID             `json:"user_uuid"`
		Version             *int64                 `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.DataDefinition == nil {
		return fmt.Errorf("required field data_definition missing")
	}
	if all.DefinitionSignature == nil {
		return fmt.Errorf("required field definition_signature missing")
	}
	if !all.Etag.IsSet() {
		return fmt.Errorf("required field etag missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.UiDefinition == nil {
		return fmt.Errorf("required field ui_definition missing")
	}
	if all.UserId == nil {
		return fmt.Errorf("required field user_id missing")
	}
	if all.UserUuid == nil {
		return fmt.Errorf("required field user_uuid missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "data_definition", "definition_signature", "etag", "id", "modified_at", "state", "ui_definition", "user_id", "user_uuid", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	if all.DataDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DataDefinition = *all.DataDefinition
	o.DefinitionSignature = *all.DefinitionSignature
	o.Etag = all.Etag
	o.Id = all.Id
	o.ModifiedAt = *all.ModifiedAt
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	if all.UiDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UiDefinition = *all.UiDefinition
	o.UserId = *all.UserId
	o.UserUuid = *all.UserUuid
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
