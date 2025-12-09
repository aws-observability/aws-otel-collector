// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatastoreDataAttributes Detailed information about a datastore.
type DatastoreDataAttributes struct {
	// Timestamp when the datastore was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The numeric ID of the user who created the datastore.
	CreatorUserId *int64 `json:"creator_user_id,omitempty"`
	// The UUID of the user who created the datastore.
	CreatorUserUuid *string `json:"creator_user_uuid,omitempty"`
	// A human-readable description about the datastore.
	Description *string `json:"description,omitempty"`
	// Timestamp when the datastore was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The display name of the datastore.
	Name *string `json:"name,omitempty"`
	// The ID of the organization that owns this datastore.
	OrgId *int64 `json:"org_id,omitempty"`
	// The name of the primary key column for this datastore. Primary column names:
	//   - Must abide by both [PostgreSQL naming conventions](https://www.postgresql.org/docs/7.0/syntax525.htm)
	//   - Cannot exceed 63 characters
	PrimaryColumnName *string `json:"primary_column_name,omitempty"`
	// Can be set to `uuid` to automatically generate primary keys when new items are added. Default value is `none`, which requires you to supply a primary key for each new item.
	PrimaryKeyGenerationStrategy *DatastorePrimaryKeyGenerationStrategy `json:"primary_key_generation_strategy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatastoreDataAttributes instantiates a new DatastoreDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatastoreDataAttributes() *DatastoreDataAttributes {
	this := DatastoreDataAttributes{}
	return &this
}

// NewDatastoreDataAttributesWithDefaults instantiates a new DatastoreDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatastoreDataAttributesWithDefaults() *DatastoreDataAttributes {
	this := DatastoreDataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DatastoreDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DatastoreDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DatastoreDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatorUserId returns the CreatorUserId field value if set, zero value otherwise.
func (o *DatastoreDataAttributes) GetCreatorUserId() int64 {
	if o == nil || o.CreatorUserId == nil {
		var ret int64
		return ret
	}
	return *o.CreatorUserId
}

// GetCreatorUserIdOk returns a tuple with the CreatorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreDataAttributes) GetCreatorUserIdOk() (*int64, bool) {
	if o == nil || o.CreatorUserId == nil {
		return nil, false
	}
	return o.CreatorUserId, true
}

// HasCreatorUserId returns a boolean if a field has been set.
func (o *DatastoreDataAttributes) HasCreatorUserId() bool {
	return o != nil && o.CreatorUserId != nil
}

// SetCreatorUserId gets a reference to the given int64 and assigns it to the CreatorUserId field.
func (o *DatastoreDataAttributes) SetCreatorUserId(v int64) {
	o.CreatorUserId = &v
}

// GetCreatorUserUuid returns the CreatorUserUuid field value if set, zero value otherwise.
func (o *DatastoreDataAttributes) GetCreatorUserUuid() string {
	if o == nil || o.CreatorUserUuid == nil {
		var ret string
		return ret
	}
	return *o.CreatorUserUuid
}

// GetCreatorUserUuidOk returns a tuple with the CreatorUserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreDataAttributes) GetCreatorUserUuidOk() (*string, bool) {
	if o == nil || o.CreatorUserUuid == nil {
		return nil, false
	}
	return o.CreatorUserUuid, true
}

// HasCreatorUserUuid returns a boolean if a field has been set.
func (o *DatastoreDataAttributes) HasCreatorUserUuid() bool {
	return o != nil && o.CreatorUserUuid != nil
}

// SetCreatorUserUuid gets a reference to the given string and assigns it to the CreatorUserUuid field.
func (o *DatastoreDataAttributes) SetCreatorUserUuid(v string) {
	o.CreatorUserUuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DatastoreDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DatastoreDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DatastoreDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *DatastoreDataAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *DatastoreDataAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *DatastoreDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DatastoreDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DatastoreDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DatastoreDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *DatastoreDataAttributes) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreDataAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *DatastoreDataAttributes) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *DatastoreDataAttributes) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetPrimaryColumnName returns the PrimaryColumnName field value if set, zero value otherwise.
func (o *DatastoreDataAttributes) GetPrimaryColumnName() string {
	if o == nil || o.PrimaryColumnName == nil {
		var ret string
		return ret
	}
	return *o.PrimaryColumnName
}

// GetPrimaryColumnNameOk returns a tuple with the PrimaryColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreDataAttributes) GetPrimaryColumnNameOk() (*string, bool) {
	if o == nil || o.PrimaryColumnName == nil {
		return nil, false
	}
	return o.PrimaryColumnName, true
}

// HasPrimaryColumnName returns a boolean if a field has been set.
func (o *DatastoreDataAttributes) HasPrimaryColumnName() bool {
	return o != nil && o.PrimaryColumnName != nil
}

// SetPrimaryColumnName gets a reference to the given string and assigns it to the PrimaryColumnName field.
func (o *DatastoreDataAttributes) SetPrimaryColumnName(v string) {
	o.PrimaryColumnName = &v
}

// GetPrimaryKeyGenerationStrategy returns the PrimaryKeyGenerationStrategy field value if set, zero value otherwise.
func (o *DatastoreDataAttributes) GetPrimaryKeyGenerationStrategy() DatastorePrimaryKeyGenerationStrategy {
	if o == nil || o.PrimaryKeyGenerationStrategy == nil {
		var ret DatastorePrimaryKeyGenerationStrategy
		return ret
	}
	return *o.PrimaryKeyGenerationStrategy
}

// GetPrimaryKeyGenerationStrategyOk returns a tuple with the PrimaryKeyGenerationStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreDataAttributes) GetPrimaryKeyGenerationStrategyOk() (*DatastorePrimaryKeyGenerationStrategy, bool) {
	if o == nil || o.PrimaryKeyGenerationStrategy == nil {
		return nil, false
	}
	return o.PrimaryKeyGenerationStrategy, true
}

// HasPrimaryKeyGenerationStrategy returns a boolean if a field has been set.
func (o *DatastoreDataAttributes) HasPrimaryKeyGenerationStrategy() bool {
	return o != nil && o.PrimaryKeyGenerationStrategy != nil
}

// SetPrimaryKeyGenerationStrategy gets a reference to the given DatastorePrimaryKeyGenerationStrategy and assigns it to the PrimaryKeyGenerationStrategy field.
func (o *DatastoreDataAttributes) SetPrimaryKeyGenerationStrategy(v DatastorePrimaryKeyGenerationStrategy) {
	o.PrimaryKeyGenerationStrategy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatastoreDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.CreatorUserId != nil {
		toSerialize["creator_user_id"] = o.CreatorUserId
	}
	if o.CreatorUserUuid != nil {
		toSerialize["creator_user_uuid"] = o.CreatorUserUuid
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.PrimaryColumnName != nil {
		toSerialize["primary_column_name"] = o.PrimaryColumnName
	}
	if o.PrimaryKeyGenerationStrategy != nil {
		toSerialize["primary_key_generation_strategy"] = o.PrimaryKeyGenerationStrategy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatastoreDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt                    *time.Time                             `json:"created_at,omitempty"`
		CreatorUserId                *int64                                 `json:"creator_user_id,omitempty"`
		CreatorUserUuid              *string                                `json:"creator_user_uuid,omitempty"`
		Description                  *string                                `json:"description,omitempty"`
		ModifiedAt                   *time.Time                             `json:"modified_at,omitempty"`
		Name                         *string                                `json:"name,omitempty"`
		OrgId                        *int64                                 `json:"org_id,omitempty"`
		PrimaryColumnName            *string                                `json:"primary_column_name,omitempty"`
		PrimaryKeyGenerationStrategy *DatastorePrimaryKeyGenerationStrategy `json:"primary_key_generation_strategy,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "creator_user_id", "creator_user_uuid", "description", "modified_at", "name", "org_id", "primary_column_name", "primary_key_generation_strategy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.CreatorUserId = all.CreatorUserId
	o.CreatorUserUuid = all.CreatorUserUuid
	o.Description = all.Description
	o.ModifiedAt = all.ModifiedAt
	o.Name = all.Name
	o.OrgId = all.OrgId
	o.PrimaryColumnName = all.PrimaryColumnName
	if all.PrimaryKeyGenerationStrategy != nil && !all.PrimaryKeyGenerationStrategy.IsValid() {
		hasInvalidField = true
	} else {
		o.PrimaryKeyGenerationStrategy = all.PrimaryKeyGenerationStrategy
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
