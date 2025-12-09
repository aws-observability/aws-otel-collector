// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListConnectionsResponseDataAttributesConnectionsItems
type ListConnectionsResponseDataAttributesConnectionsItems struct {
	//
	CreatedAt *time.Time `json:"created_at,omitempty"`
	//
	CreatedBy *string `json:"created_by,omitempty"`
	//
	Fields []CreateConnectionRequestDataAttributesFieldsItems `json:"fields,omitempty"`
	//
	Id *string `json:"id,omitempty"`
	//
	Join *ListConnectionsResponseDataAttributesConnectionsItemsJoin `json:"join,omitempty"`
	//
	Metadata map[string]string `json:"metadata,omitempty"`
	//
	Type *string `json:"type,omitempty"`
	//
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	//
	UpdatedBy *string `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListConnectionsResponseDataAttributesConnectionsItems instantiates a new ListConnectionsResponseDataAttributesConnectionsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListConnectionsResponseDataAttributesConnectionsItems() *ListConnectionsResponseDataAttributesConnectionsItems {
	this := ListConnectionsResponseDataAttributesConnectionsItems{}
	return &this
}

// NewListConnectionsResponseDataAttributesConnectionsItemsWithDefaults instantiates a new ListConnectionsResponseDataAttributesConnectionsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListConnectionsResponseDataAttributesConnectionsItemsWithDefaults() *ListConnectionsResponseDataAttributesConnectionsItems {
	this := ListConnectionsResponseDataAttributesConnectionsItems{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetFields() []CreateConnectionRequestDataAttributesFieldsItems {
	if o == nil || o.Fields == nil {
		var ret []CreateConnectionRequestDataAttributesFieldsItems
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetFieldsOk() (*[]CreateConnectionRequestDataAttributesFieldsItems, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given []CreateConnectionRequestDataAttributesFieldsItems and assigns it to the Fields field.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) SetFields(v []CreateConnectionRequestDataAttributesFieldsItems) {
	o.Fields = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) SetId(v string) {
	o.Id = &v
}

// GetJoin returns the Join field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetJoin() ListConnectionsResponseDataAttributesConnectionsItemsJoin {
	if o == nil || o.Join == nil {
		var ret ListConnectionsResponseDataAttributesConnectionsItemsJoin
		return ret
	}
	return *o.Join
}

// GetJoinOk returns a tuple with the Join field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetJoinOk() (*ListConnectionsResponseDataAttributesConnectionsItemsJoin, bool) {
	if o == nil || o.Join == nil {
		return nil, false
	}
	return o.Join, true
}

// HasJoin returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) HasJoin() bool {
	return o != nil && o.Join != nil
}

// SetJoin gets a reference to the given ListConnectionsResponseDataAttributesConnectionsItemsJoin and assigns it to the Join field.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) SetJoin(v ListConnectionsResponseDataAttributesConnectionsItemsJoin) {
	o.Join = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListConnectionsResponseDataAttributesConnectionsItems) MarshalJSON() ([]byte, error) {
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
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Join != nil {
		toSerialize["join"] = o.Join
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListConnectionsResponseDataAttributesConnectionsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *time.Time                                                 `json:"created_at,omitempty"`
		CreatedBy *string                                                    `json:"created_by,omitempty"`
		Fields    []CreateConnectionRequestDataAttributesFieldsItems         `json:"fields,omitempty"`
		Id        *string                                                    `json:"id,omitempty"`
		Join      *ListConnectionsResponseDataAttributesConnectionsItemsJoin `json:"join,omitempty"`
		Metadata  map[string]string                                          `json:"metadata,omitempty"`
		Type      *string                                                    `json:"type,omitempty"`
		UpdatedAt *time.Time                                                 `json:"updated_at,omitempty"`
		UpdatedBy *string                                                    `json:"updated_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "fields", "id", "join", "metadata", "type", "updated_at", "updated_by"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.Fields = all.Fields
	o.Id = all.Id
	if all.Join != nil && all.Join.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Join = all.Join
	o.Metadata = all.Metadata
	o.Type = all.Type
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedBy = all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
