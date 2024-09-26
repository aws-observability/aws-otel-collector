// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RelationshipToUserTeamPermission Relationship between a user team permission and a team
type RelationshipToUserTeamPermission struct {
	// Related user team permission data
	Data *RelationshipToUserTeamPermissionData `json:"data,omitempty"`
	// Links attributes.
	Links *TeamRelationshipsLinks `json:"links,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRelationshipToUserTeamPermission instantiates a new RelationshipToUserTeamPermission object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRelationshipToUserTeamPermission() *RelationshipToUserTeamPermission {
	this := RelationshipToUserTeamPermission{}
	return &this
}

// NewRelationshipToUserTeamPermissionWithDefaults instantiates a new RelationshipToUserTeamPermission object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRelationshipToUserTeamPermissionWithDefaults() *RelationshipToUserTeamPermission {
	this := RelationshipToUserTeamPermission{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RelationshipToUserTeamPermission) GetData() RelationshipToUserTeamPermissionData {
	if o == nil || o.Data == nil {
		var ret RelationshipToUserTeamPermissionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipToUserTeamPermission) GetDataOk() (*RelationshipToUserTeamPermissionData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RelationshipToUserTeamPermission) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given RelationshipToUserTeamPermissionData and assigns it to the Data field.
func (o *RelationshipToUserTeamPermission) SetData(v RelationshipToUserTeamPermissionData) {
	o.Data = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RelationshipToUserTeamPermission) GetLinks() TeamRelationshipsLinks {
	if o == nil || o.Links == nil {
		var ret TeamRelationshipsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipToUserTeamPermission) GetLinksOk() (*TeamRelationshipsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RelationshipToUserTeamPermission) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given TeamRelationshipsLinks and assigns it to the Links field.
func (o *RelationshipToUserTeamPermission) SetLinks(v TeamRelationshipsLinks) {
	o.Links = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RelationshipToUserTeamPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RelationshipToUserTeamPermission) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *RelationshipToUserTeamPermissionData `json:"data,omitempty"`
		Links *TeamRelationshipsLinks               `json:"links,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
