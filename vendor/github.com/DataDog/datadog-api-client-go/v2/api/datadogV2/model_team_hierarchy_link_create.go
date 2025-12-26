// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyLinkCreate Data provided when creating a team hierarchy link
type TeamHierarchyLinkCreate struct {
	// The related teams that will be connected by the team hierarchy link
	Relationships TeamHierarchyLinkCreateRelationships `json:"relationships"`
	// Team hierarchy link type
	Type TeamHierarchyLinkType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyLinkCreate instantiates a new TeamHierarchyLinkCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyLinkCreate(relationships TeamHierarchyLinkCreateRelationships, typeVar TeamHierarchyLinkType) *TeamHierarchyLinkCreate {
	this := TeamHierarchyLinkCreate{}
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewTeamHierarchyLinkCreateWithDefaults instantiates a new TeamHierarchyLinkCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyLinkCreateWithDefaults() *TeamHierarchyLinkCreate {
	this := TeamHierarchyLinkCreate{}
	var typeVar TeamHierarchyLinkType = TEAMHIERARCHYLINKTYPE_TEAM_HIERARCHY_LINKS
	this.Type = typeVar
	return &this
}

// GetRelationships returns the Relationships field value.
func (o *TeamHierarchyLinkCreate) GetRelationships() TeamHierarchyLinkCreateRelationships {
	if o == nil {
		var ret TeamHierarchyLinkCreateRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkCreate) GetRelationshipsOk() (*TeamHierarchyLinkCreateRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *TeamHierarchyLinkCreate) SetRelationships(v TeamHierarchyLinkCreateRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *TeamHierarchyLinkCreate) GetType() TeamHierarchyLinkType {
	if o == nil {
		var ret TeamHierarchyLinkType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkCreate) GetTypeOk() (*TeamHierarchyLinkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamHierarchyLinkCreate) SetType(v TeamHierarchyLinkType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyLinkCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamHierarchyLinkCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Relationships *TeamHierarchyLinkCreateRelationships `json:"relationships"`
		Type          *TeamHierarchyLinkType                `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
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
