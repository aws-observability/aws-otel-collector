// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamUpdateRelationships Team update relationships
type TeamUpdateRelationships struct {
	// Relationship between a team and a team link
	TeamLinks *RelationshipToTeamLinks `json:"team_links,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamUpdateRelationships instantiates a new TeamUpdateRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamUpdateRelationships() *TeamUpdateRelationships {
	this := TeamUpdateRelationships{}
	return &this
}

// NewTeamUpdateRelationshipsWithDefaults instantiates a new TeamUpdateRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamUpdateRelationshipsWithDefaults() *TeamUpdateRelationships {
	this := TeamUpdateRelationships{}
	return &this
}

// GetTeamLinks returns the TeamLinks field value if set, zero value otherwise.
func (o *TeamUpdateRelationships) GetTeamLinks() RelationshipToTeamLinks {
	if o == nil || o.TeamLinks == nil {
		var ret RelationshipToTeamLinks
		return ret
	}
	return *o.TeamLinks
}

// GetTeamLinksOk returns a tuple with the TeamLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdateRelationships) GetTeamLinksOk() (*RelationshipToTeamLinks, bool) {
	if o == nil || o.TeamLinks == nil {
		return nil, false
	}
	return o.TeamLinks, true
}

// HasTeamLinks returns a boolean if a field has been set.
func (o *TeamUpdateRelationships) HasTeamLinks() bool {
	return o != nil && o.TeamLinks != nil
}

// SetTeamLinks gets a reference to the given RelationshipToTeamLinks and assigns it to the TeamLinks field.
func (o *TeamUpdateRelationships) SetTeamLinks(v RelationshipToTeamLinks) {
	o.TeamLinks = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamUpdateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TeamLinks != nil {
		toSerialize["team_links"] = o.TeamLinks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamUpdateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TeamLinks *RelationshipToTeamLinks `json:"team_links,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"team_links"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.TeamLinks != nil && all.TeamLinks.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TeamLinks = all.TeamLinks

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
