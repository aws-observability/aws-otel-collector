// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupMembershipRelationshipData A reference to an org group membership.
type OrgGroupMembershipRelationshipData struct {
	// The ID of the membership.
	Id uuid.UUID `json:"id"`
	// Org group memberships resource type.
	Type OrgGroupMembershipType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupMembershipRelationshipData instantiates a new OrgGroupMembershipRelationshipData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupMembershipRelationshipData(id uuid.UUID, typeVar OrgGroupMembershipType) *OrgGroupMembershipRelationshipData {
	this := OrgGroupMembershipRelationshipData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewOrgGroupMembershipRelationshipDataWithDefaults instantiates a new OrgGroupMembershipRelationshipData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupMembershipRelationshipDataWithDefaults() *OrgGroupMembershipRelationshipData {
	this := OrgGroupMembershipRelationshipData{}
	return &this
}

// GetId returns the Id field value.
func (o *OrgGroupMembershipRelationshipData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipRelationshipData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OrgGroupMembershipRelationshipData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *OrgGroupMembershipRelationshipData) GetType() OrgGroupMembershipType {
	if o == nil {
		var ret OrgGroupMembershipType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipRelationshipData) GetTypeOk() (*OrgGroupMembershipType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OrgGroupMembershipRelationshipData) SetType(v OrgGroupMembershipType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupMembershipRelationshipData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupMembershipRelationshipData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *uuid.UUID              `json:"id"`
		Type *OrgGroupMembershipType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
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
