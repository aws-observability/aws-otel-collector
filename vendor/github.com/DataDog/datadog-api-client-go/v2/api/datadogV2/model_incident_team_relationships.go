// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTeamRelationships The incident team's relationships.
type IncidentTeamRelationships struct {
	// Relationship to user.
	CreatedBy *RelationshipToUser `json:"created_by,omitempty"`
	// Relationship to user.
	LastModifiedBy *RelationshipToUser `json:"last_modified_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTeamRelationships instantiates a new IncidentTeamRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTeamRelationships() *IncidentTeamRelationships {
	this := IncidentTeamRelationships{}
	return &this
}

// NewIncidentTeamRelationshipsWithDefaults instantiates a new IncidentTeamRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTeamRelationshipsWithDefaults() *IncidentTeamRelationships {
	this := IncidentTeamRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *IncidentTeamRelationships) GetCreatedBy() RelationshipToUser {
	if o == nil || o.CreatedBy == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTeamRelationships) GetCreatedByOk() (*RelationshipToUser, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *IncidentTeamRelationships) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given RelationshipToUser and assigns it to the CreatedBy field.
func (o *IncidentTeamRelationships) SetCreatedBy(v RelationshipToUser) {
	o.CreatedBy = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *IncidentTeamRelationships) GetLastModifiedBy() RelationshipToUser {
	if o == nil || o.LastModifiedBy == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTeamRelationships) GetLastModifiedByOk() (*RelationshipToUser, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *IncidentTeamRelationships) HasLastModifiedBy() bool {
	return o != nil && o.LastModifiedBy != nil
}

// SetLastModifiedBy gets a reference to the given RelationshipToUser and assigns it to the LastModifiedBy field.
func (o *IncidentTeamRelationships) SetLastModifiedBy(v RelationshipToUser) {
	o.LastModifiedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTeamRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.LastModifiedBy != nil {
		toSerialize["last_modified_by"] = o.LastModifiedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTeamRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy      *RelationshipToUser `json:"created_by,omitempty"`
		LastModifiedBy *RelationshipToUser `json:"last_modified_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by", "last_modified_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedBy != nil && all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = all.CreatedBy
	if all.LastModifiedBy != nil && all.LastModifiedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastModifiedBy = all.LastModifiedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
