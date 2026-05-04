// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupMembershipUpdateData Data for updating an org group membership.
type OrgGroupMembershipUpdateData struct {
	// The ID of the membership.
	Id uuid.UUID `json:"id"`
	// Relationships for updating a membership.
	Relationships OrgGroupMembershipUpdateRelationships `json:"relationships"`
	// Org group memberships resource type.
	Type OrgGroupMembershipType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupMembershipUpdateData instantiates a new OrgGroupMembershipUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupMembershipUpdateData(id uuid.UUID, relationships OrgGroupMembershipUpdateRelationships, typeVar OrgGroupMembershipType) *OrgGroupMembershipUpdateData {
	this := OrgGroupMembershipUpdateData{}
	this.Id = id
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewOrgGroupMembershipUpdateDataWithDefaults instantiates a new OrgGroupMembershipUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupMembershipUpdateDataWithDefaults() *OrgGroupMembershipUpdateData {
	this := OrgGroupMembershipUpdateData{}
	return &this
}

// GetId returns the Id field value.
func (o *OrgGroupMembershipUpdateData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipUpdateData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OrgGroupMembershipUpdateData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetRelationships returns the Relationships field value.
func (o *OrgGroupMembershipUpdateData) GetRelationships() OrgGroupMembershipUpdateRelationships {
	if o == nil {
		var ret OrgGroupMembershipUpdateRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipUpdateData) GetRelationshipsOk() (*OrgGroupMembershipUpdateRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *OrgGroupMembershipUpdateData) SetRelationships(v OrgGroupMembershipUpdateRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *OrgGroupMembershipUpdateData) GetType() OrgGroupMembershipType {
	if o == nil {
		var ret OrgGroupMembershipType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipUpdateData) GetTypeOk() (*OrgGroupMembershipType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OrgGroupMembershipUpdateData) SetType(v OrgGroupMembershipType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupMembershipUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupMembershipUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *uuid.UUID                             `json:"id"`
		Relationships *OrgGroupMembershipUpdateRelationships `json:"relationships"`
		Type          *OrgGroupMembershipType                `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	if all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = *all.Relationships
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
