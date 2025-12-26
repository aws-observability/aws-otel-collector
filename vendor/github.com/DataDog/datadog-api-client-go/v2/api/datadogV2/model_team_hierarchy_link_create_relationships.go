// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyLinkCreateRelationships The related teams that will be connected by the team hierarchy link
type TeamHierarchyLinkCreateRelationships struct {
	// Data about each team that will be connected by the team hierarchy link
	ParentTeam TeamHierarchyLinkCreateTeamRelationship `json:"parent_team"`
	// Data about each team that will be connected by the team hierarchy link
	SubTeam TeamHierarchyLinkCreateTeamRelationship `json:"sub_team"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyLinkCreateRelationships instantiates a new TeamHierarchyLinkCreateRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyLinkCreateRelationships(parentTeam TeamHierarchyLinkCreateTeamRelationship, subTeam TeamHierarchyLinkCreateTeamRelationship) *TeamHierarchyLinkCreateRelationships {
	this := TeamHierarchyLinkCreateRelationships{}
	this.ParentTeam = parentTeam
	this.SubTeam = subTeam
	return &this
}

// NewTeamHierarchyLinkCreateRelationshipsWithDefaults instantiates a new TeamHierarchyLinkCreateRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyLinkCreateRelationshipsWithDefaults() *TeamHierarchyLinkCreateRelationships {
	this := TeamHierarchyLinkCreateRelationships{}
	return &this
}

// GetParentTeam returns the ParentTeam field value.
func (o *TeamHierarchyLinkCreateRelationships) GetParentTeam() TeamHierarchyLinkCreateTeamRelationship {
	if o == nil {
		var ret TeamHierarchyLinkCreateTeamRelationship
		return ret
	}
	return o.ParentTeam
}

// GetParentTeamOk returns a tuple with the ParentTeam field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkCreateRelationships) GetParentTeamOk() (*TeamHierarchyLinkCreateTeamRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentTeam, true
}

// SetParentTeam sets field value.
func (o *TeamHierarchyLinkCreateRelationships) SetParentTeam(v TeamHierarchyLinkCreateTeamRelationship) {
	o.ParentTeam = v
}

// GetSubTeam returns the SubTeam field value.
func (o *TeamHierarchyLinkCreateRelationships) GetSubTeam() TeamHierarchyLinkCreateTeamRelationship {
	if o == nil {
		var ret TeamHierarchyLinkCreateTeamRelationship
		return ret
	}
	return o.SubTeam
}

// GetSubTeamOk returns a tuple with the SubTeam field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkCreateRelationships) GetSubTeamOk() (*TeamHierarchyLinkCreateTeamRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubTeam, true
}

// SetSubTeam sets field value.
func (o *TeamHierarchyLinkCreateRelationships) SetSubTeam(v TeamHierarchyLinkCreateTeamRelationship) {
	o.SubTeam = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyLinkCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["parent_team"] = o.ParentTeam
	toSerialize["sub_team"] = o.SubTeam

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamHierarchyLinkCreateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ParentTeam *TeamHierarchyLinkCreateTeamRelationship `json:"parent_team"`
		SubTeam    *TeamHierarchyLinkCreateTeamRelationship `json:"sub_team"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ParentTeam == nil {
		return fmt.Errorf("required field parent_team missing")
	}
	if all.SubTeam == nil {
		return fmt.Errorf("required field sub_team missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"parent_team", "sub_team"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ParentTeam.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ParentTeam = *all.ParentTeam
	if all.SubTeam.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SubTeam = *all.SubTeam

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
