// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldRelationships Relationships of an incident user-defined field.
type IncidentUserDefinedFieldRelationships struct {
	// Relationship to user.
	CreatedByUser RelationshipToUser `json:"created_by_user"`
	// Relationship to an incident type.
	IncidentType RelationshipToIncidentType `json:"incident_type"`
	// Relationship to user.
	LastModifiedByUser RelationshipToUser `json:"last_modified_by_user"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedFieldRelationships instantiates a new IncidentUserDefinedFieldRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedFieldRelationships(createdByUser RelationshipToUser, incidentType RelationshipToIncidentType, lastModifiedByUser RelationshipToUser) *IncidentUserDefinedFieldRelationships {
	this := IncidentUserDefinedFieldRelationships{}
	this.CreatedByUser = createdByUser
	this.IncidentType = incidentType
	this.LastModifiedByUser = lastModifiedByUser
	return &this
}

// NewIncidentUserDefinedFieldRelationshipsWithDefaults instantiates a new IncidentUserDefinedFieldRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedFieldRelationshipsWithDefaults() *IncidentUserDefinedFieldRelationships {
	this := IncidentUserDefinedFieldRelationships{}
	return &this
}

// GetCreatedByUser returns the CreatedByUser field value.
func (o *IncidentUserDefinedFieldRelationships) GetCreatedByUser() RelationshipToUser {
	if o == nil {
		var ret RelationshipToUser
		return ret
	}
	return o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldRelationships) GetCreatedByUserOk() (*RelationshipToUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByUser, true
}

// SetCreatedByUser sets field value.
func (o *IncidentUserDefinedFieldRelationships) SetCreatedByUser(v RelationshipToUser) {
	o.CreatedByUser = v
}

// GetIncidentType returns the IncidentType field value.
func (o *IncidentUserDefinedFieldRelationships) GetIncidentType() RelationshipToIncidentType {
	if o == nil {
		var ret RelationshipToIncidentType
		return ret
	}
	return o.IncidentType
}

// GetIncidentTypeOk returns a tuple with the IncidentType field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldRelationships) GetIncidentTypeOk() (*RelationshipToIncidentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentType, true
}

// SetIncidentType sets field value.
func (o *IncidentUserDefinedFieldRelationships) SetIncidentType(v RelationshipToIncidentType) {
	o.IncidentType = v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value.
func (o *IncidentUserDefinedFieldRelationships) GetLastModifiedByUser() RelationshipToUser {
	if o == nil {
		var ret RelationshipToUser
		return ret
	}
	return o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldRelationships) GetLastModifiedByUserOk() (*RelationshipToUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedByUser, true
}

// SetLastModifiedByUser sets field value.
func (o *IncidentUserDefinedFieldRelationships) SetLastModifiedByUser(v RelationshipToUser) {
	o.LastModifiedByUser = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedFieldRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["created_by_user"] = o.CreatedByUser
	toSerialize["incident_type"] = o.IncidentType
	toSerialize["last_modified_by_user"] = o.LastModifiedByUser

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedFieldRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedByUser      *RelationshipToUser         `json:"created_by_user"`
		IncidentType       *RelationshipToIncidentType `json:"incident_type"`
		LastModifiedByUser *RelationshipToUser         `json:"last_modified_by_user"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedByUser == nil {
		return fmt.Errorf("required field created_by_user missing")
	}
	if all.IncidentType == nil {
		return fmt.Errorf("required field incident_type missing")
	}
	if all.LastModifiedByUser == nil {
		return fmt.Errorf("required field last_modified_by_user missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by_user", "incident_type", "last_modified_by_user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedByUser = *all.CreatedByUser
	if all.IncidentType.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentType = *all.IncidentType
	if all.LastModifiedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastModifiedByUser = *all.LastModifiedByUser

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
