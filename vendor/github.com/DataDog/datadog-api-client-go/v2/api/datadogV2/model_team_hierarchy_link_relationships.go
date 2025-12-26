// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyLinkRelationships Team hierarchy link relationships
type TeamHierarchyLinkRelationships struct {
	// Team hierarchy link team relationship
	ParentTeam TeamHierarchyLinkTeamRelationship `json:"parent_team"`
	// Team hierarchy link team relationship
	SubTeam TeamHierarchyLinkTeamRelationship `json:"sub_team"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyLinkRelationships instantiates a new TeamHierarchyLinkRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyLinkRelationships(parentTeam TeamHierarchyLinkTeamRelationship, subTeam TeamHierarchyLinkTeamRelationship) *TeamHierarchyLinkRelationships {
	this := TeamHierarchyLinkRelationships{}
	this.ParentTeam = parentTeam
	this.SubTeam = subTeam
	return &this
}

// NewTeamHierarchyLinkRelationshipsWithDefaults instantiates a new TeamHierarchyLinkRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyLinkRelationshipsWithDefaults() *TeamHierarchyLinkRelationships {
	this := TeamHierarchyLinkRelationships{}
	return &this
}

// GetParentTeam returns the ParentTeam field value.
func (o *TeamHierarchyLinkRelationships) GetParentTeam() TeamHierarchyLinkTeamRelationship {
	if o == nil {
		var ret TeamHierarchyLinkTeamRelationship
		return ret
	}
	return o.ParentTeam
}

// GetParentTeamOk returns a tuple with the ParentTeam field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkRelationships) GetParentTeamOk() (*TeamHierarchyLinkTeamRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentTeam, true
}

// SetParentTeam sets field value.
func (o *TeamHierarchyLinkRelationships) SetParentTeam(v TeamHierarchyLinkTeamRelationship) {
	o.ParentTeam = v
}

// GetSubTeam returns the SubTeam field value.
func (o *TeamHierarchyLinkRelationships) GetSubTeam() TeamHierarchyLinkTeamRelationship {
	if o == nil {
		var ret TeamHierarchyLinkTeamRelationship
		return ret
	}
	return o.SubTeam
}

// GetSubTeamOk returns a tuple with the SubTeam field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkRelationships) GetSubTeamOk() (*TeamHierarchyLinkTeamRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubTeam, true
}

// SetSubTeam sets field value.
func (o *TeamHierarchyLinkRelationships) SetSubTeam(v TeamHierarchyLinkTeamRelationship) {
	o.SubTeam = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyLinkRelationships) MarshalJSON() ([]byte, error) {
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
func (o *TeamHierarchyLinkRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ParentTeam *TeamHierarchyLinkTeamRelationship `json:"parent_team"`
		SubTeam    *TeamHierarchyLinkTeamRelationship `json:"sub_team"`
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
