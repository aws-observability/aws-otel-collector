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

// FormDataAttributes The attributes of a form.
type FormDataAttributes struct {
	// Whether the form is currently active.
	Active bool `json:"active"`
	// Whether the form accepts anonymous submissions.
	Anonymous bool `json:"anonymous"`
	// The time at which the form was created.
	CreatedAt time.Time `json:"created_at"`
	// The datastore configuration for a form.
	DatastoreConfig FormDatastoreConfigAttributes `json:"datastore_config"`
	// The description of the form.
	Description string `json:"description"`
	// The date and time at which the form stops accepting responses.
	EndDate datadog.NullableTime `json:"end_date,omitempty"`
	// Whether the current user has already submitted this form. Only present for forms with `single_response` set to `true`.
	HasSubmitted datadog.NullableBool `json:"has_submitted,omitempty"`
	// Whether the form is an IDP survey.
	IdpSurvey bool `json:"idp_survey"`
	// The time at which the form was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The name of the form.
	Name string `json:"name"`
	// The ID of the organization that owns this form.
	OrgId int64 `json:"org_id"`
	// The attributes of a form publication.
	Publication *FormPublicationAttributes `json:"publication,omitempty"`
	// Whether the form is available in the self-service catalog.
	SelfService bool `json:"self_service"`
	// Whether each user can only submit one response.
	SingleResponse bool `json:"single_response"`
	// The ID of the user who created this form.
	UserId int64 `json:"user_id"`
	// The UUID of the user who created this form.
	UserUuid uuid.UUID `json:"user_uuid"`
	// The attributes of a form version.
	Version *FormVersionAttributes `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormDataAttributes instantiates a new FormDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormDataAttributes(active bool, anonymous bool, createdAt time.Time, datastoreConfig FormDatastoreConfigAttributes, description string, idpSurvey bool, modifiedAt time.Time, name string, orgId int64, selfService bool, singleResponse bool, userId int64, userUuid uuid.UUID) *FormDataAttributes {
	this := FormDataAttributes{}
	this.Active = active
	this.Anonymous = anonymous
	this.CreatedAt = createdAt
	this.DatastoreConfig = datastoreConfig
	this.Description = description
	this.IdpSurvey = idpSurvey
	this.ModifiedAt = modifiedAt
	this.Name = name
	this.OrgId = orgId
	this.SelfService = selfService
	this.SingleResponse = singleResponse
	this.UserId = userId
	this.UserUuid = userUuid
	return &this
}

// NewFormDataAttributesWithDefaults instantiates a new FormDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormDataAttributesWithDefaults() *FormDataAttributes {
	this := FormDataAttributes{}
	return &this
}

// GetActive returns the Active field value.
func (o *FormDataAttributes) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value.
func (o *FormDataAttributes) SetActive(v bool) {
	o.Active = v
}

// GetAnonymous returns the Anonymous field value.
func (o *FormDataAttributes) GetAnonymous() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Anonymous
}

// GetAnonymousOk returns a tuple with the Anonymous field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetAnonymousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Anonymous, true
}

// SetAnonymous sets field value.
func (o *FormDataAttributes) SetAnonymous(v bool) {
	o.Anonymous = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *FormDataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *FormDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDatastoreConfig returns the DatastoreConfig field value.
func (o *FormDataAttributes) GetDatastoreConfig() FormDatastoreConfigAttributes {
	if o == nil {
		var ret FormDatastoreConfigAttributes
		return ret
	}
	return o.DatastoreConfig
}

// GetDatastoreConfigOk returns a tuple with the DatastoreConfig field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetDatastoreConfigOk() (*FormDatastoreConfigAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatastoreConfig, true
}

// SetDatastoreConfig sets field value.
func (o *FormDataAttributes) SetDatastoreConfig(v FormDatastoreConfigAttributes) {
	o.DatastoreConfig = v
}

// GetDescription returns the Description field value.
func (o *FormDataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *FormDataAttributes) SetDescription(v string) {
	o.Description = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormDataAttributes) GetEndDate() time.Time {
	if o == nil || o.EndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FormDataAttributes) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *FormDataAttributes) HasEndDate() bool {
	return o != nil && o.EndDate.IsSet()
}

// SetEndDate gets a reference to the given datadog.NullableTime and assigns it to the EndDate field.
func (o *FormDataAttributes) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}

// SetEndDateNil sets the value for EndDate to be an explicit nil.
func (o *FormDataAttributes) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil.
func (o *FormDataAttributes) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetHasSubmitted returns the HasSubmitted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormDataAttributes) GetHasSubmitted() bool {
	if o == nil || o.HasSubmitted.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasSubmitted.Get()
}

// GetHasSubmittedOk returns a tuple with the HasSubmitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FormDataAttributes) GetHasSubmittedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasSubmitted.Get(), o.HasSubmitted.IsSet()
}

// HasHasSubmitted returns a boolean if a field has been set.
func (o *FormDataAttributes) HasHasSubmitted() bool {
	return o != nil && o.HasSubmitted.IsSet()
}

// SetHasSubmitted gets a reference to the given datadog.NullableBool and assigns it to the HasSubmitted field.
func (o *FormDataAttributes) SetHasSubmitted(v bool) {
	o.HasSubmitted.Set(&v)
}

// SetHasSubmittedNil sets the value for HasSubmitted to be an explicit nil.
func (o *FormDataAttributes) SetHasSubmittedNil() {
	o.HasSubmitted.Set(nil)
}

// UnsetHasSubmitted ensures that no value is present for HasSubmitted, not even an explicit nil.
func (o *FormDataAttributes) UnsetHasSubmitted() {
	o.HasSubmitted.Unset()
}

// GetIdpSurvey returns the IdpSurvey field value.
func (o *FormDataAttributes) GetIdpSurvey() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IdpSurvey
}

// GetIdpSurveyOk returns a tuple with the IdpSurvey field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetIdpSurveyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpSurvey, true
}

// SetIdpSurvey sets field value.
func (o *FormDataAttributes) SetIdpSurvey(v bool) {
	o.IdpSurvey = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *FormDataAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *FormDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetName returns the Name field value.
func (o *FormDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormDataAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value.
func (o *FormDataAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *FormDataAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetPublication returns the Publication field value if set, zero value otherwise.
func (o *FormDataAttributes) GetPublication() FormPublicationAttributes {
	if o == nil || o.Publication == nil {
		var ret FormPublicationAttributes
		return ret
	}
	return *o.Publication
}

// GetPublicationOk returns a tuple with the Publication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetPublicationOk() (*FormPublicationAttributes, bool) {
	if o == nil || o.Publication == nil {
		return nil, false
	}
	return o.Publication, true
}

// HasPublication returns a boolean if a field has been set.
func (o *FormDataAttributes) HasPublication() bool {
	return o != nil && o.Publication != nil
}

// SetPublication gets a reference to the given FormPublicationAttributes and assigns it to the Publication field.
func (o *FormDataAttributes) SetPublication(v FormPublicationAttributes) {
	o.Publication = &v
}

// GetSelfService returns the SelfService field value.
func (o *FormDataAttributes) GetSelfService() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SelfService
}

// GetSelfServiceOk returns a tuple with the SelfService field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetSelfServiceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfService, true
}

// SetSelfService sets field value.
func (o *FormDataAttributes) SetSelfService(v bool) {
	o.SelfService = v
}

// GetSingleResponse returns the SingleResponse field value.
func (o *FormDataAttributes) GetSingleResponse() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SingleResponse
}

// GetSingleResponseOk returns a tuple with the SingleResponse field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetSingleResponseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SingleResponse, true
}

// SetSingleResponse sets field value.
func (o *FormDataAttributes) SetSingleResponse(v bool) {
	o.SingleResponse = v
}

// GetUserId returns the UserId field value.
func (o *FormDataAttributes) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *FormDataAttributes) SetUserId(v int64) {
	o.UserId = v
}

// GetUserUuid returns the UserUuid field value.
func (o *FormDataAttributes) GetUserUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetUserUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuid, true
}

// SetUserUuid sets field value.
func (o *FormDataAttributes) SetUserUuid(v uuid.UUID) {
	o.UserUuid = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FormDataAttributes) GetVersion() FormVersionAttributes {
	if o == nil || o.Version == nil {
		var ret FormVersionAttributes
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataAttributes) GetVersionOk() (*FormVersionAttributes, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FormDataAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given FormVersionAttributes and assigns it to the Version field.
func (o *FormDataAttributes) SetVersion(v FormVersionAttributes) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["active"] = o.Active
	toSerialize["anonymous"] = o.Anonymous
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["datastore_config"] = o.DatastoreConfig
	toSerialize["description"] = o.Description
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.HasSubmitted.IsSet() {
		toSerialize["has_submitted"] = o.HasSubmitted.Get()
	}
	toSerialize["idp_survey"] = o.IdpSurvey
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["org_id"] = o.OrgId
	if o.Publication != nil {
		toSerialize["publication"] = o.Publication
	}
	toSerialize["self_service"] = o.SelfService
	toSerialize["single_response"] = o.SingleResponse
	toSerialize["user_id"] = o.UserId
	toSerialize["user_uuid"] = o.UserUuid
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Active          *bool                          `json:"active"`
		Anonymous       *bool                          `json:"anonymous"`
		CreatedAt       *time.Time                     `json:"created_at"`
		DatastoreConfig *FormDatastoreConfigAttributes `json:"datastore_config"`
		Description     *string                        `json:"description"`
		EndDate         datadog.NullableTime           `json:"end_date,omitempty"`
		HasSubmitted    datadog.NullableBool           `json:"has_submitted,omitempty"`
		IdpSurvey       *bool                          `json:"idp_survey"`
		ModifiedAt      *time.Time                     `json:"modified_at"`
		Name            *string                        `json:"name"`
		OrgId           *int64                         `json:"org_id"`
		Publication     *FormPublicationAttributes     `json:"publication,omitempty"`
		SelfService     *bool                          `json:"self_service"`
		SingleResponse  *bool                          `json:"single_response"`
		UserId          *int64                         `json:"user_id"`
		UserUuid        *uuid.UUID                     `json:"user_uuid"`
		Version         *FormVersionAttributes         `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Active == nil {
		return fmt.Errorf("required field active missing")
	}
	if all.Anonymous == nil {
		return fmt.Errorf("required field anonymous missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.DatastoreConfig == nil {
		return fmt.Errorf("required field datastore_config missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.IdpSurvey == nil {
		return fmt.Errorf("required field idp_survey missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.SelfService == nil {
		return fmt.Errorf("required field self_service missing")
	}
	if all.SingleResponse == nil {
		return fmt.Errorf("required field single_response missing")
	}
	if all.UserId == nil {
		return fmt.Errorf("required field user_id missing")
	}
	if all.UserUuid == nil {
		return fmt.Errorf("required field user_uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"active", "anonymous", "created_at", "datastore_config", "description", "end_date", "has_submitted", "idp_survey", "modified_at", "name", "org_id", "publication", "self_service", "single_response", "user_id", "user_uuid", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Active = *all.Active
	o.Anonymous = *all.Anonymous
	o.CreatedAt = *all.CreatedAt
	if all.DatastoreConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatastoreConfig = *all.DatastoreConfig
	o.Description = *all.Description
	o.EndDate = all.EndDate
	o.HasSubmitted = all.HasSubmitted
	o.IdpSurvey = *all.IdpSurvey
	o.ModifiedAt = *all.ModifiedAt
	o.Name = *all.Name
	o.OrgId = *all.OrgId
	if all.Publication != nil && all.Publication.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Publication = all.Publication
	o.SelfService = *all.SelfService
	o.SingleResponse = *all.SingleResponse
	o.UserId = *all.UserId
	o.UserUuid = *all.UserUuid
	if all.Version != nil && all.Version.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
