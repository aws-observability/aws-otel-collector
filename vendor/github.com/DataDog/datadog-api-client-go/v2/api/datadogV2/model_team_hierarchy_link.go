// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyLink Team hierarchy link
type TeamHierarchyLink struct {
	// Team hierarchy link attributes
	Attributes TeamHierarchyLinkAttributes `json:"attributes"`
	// The team hierarchy link's identifier
	Id string `json:"id"`
	// Team hierarchy link relationships
	Relationships *TeamHierarchyLinkRelationships `json:"relationships,omitempty"`
	// Team hierarchy link type
	Type TeamHierarchyLinkType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyLink instantiates a new TeamHierarchyLink object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyLink(attributes TeamHierarchyLinkAttributes, id string, typeVar TeamHierarchyLinkType) *TeamHierarchyLink {
	this := TeamHierarchyLink{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewTeamHierarchyLinkWithDefaults instantiates a new TeamHierarchyLink object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyLinkWithDefaults() *TeamHierarchyLink {
	this := TeamHierarchyLink{}
	var typeVar TeamHierarchyLinkType = TEAMHIERARCHYLINKTYPE_TEAM_HIERARCHY_LINKS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *TeamHierarchyLink) GetAttributes() TeamHierarchyLinkAttributes {
	if o == nil {
		var ret TeamHierarchyLinkAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLink) GetAttributesOk() (*TeamHierarchyLinkAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *TeamHierarchyLink) SetAttributes(v TeamHierarchyLinkAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *TeamHierarchyLink) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLink) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TeamHierarchyLink) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TeamHierarchyLink) GetRelationships() TeamHierarchyLinkRelationships {
	if o == nil || o.Relationships == nil {
		var ret TeamHierarchyLinkRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLink) GetRelationshipsOk() (*TeamHierarchyLinkRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TeamHierarchyLink) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given TeamHierarchyLinkRelationships and assigns it to the Relationships field.
func (o *TeamHierarchyLink) SetRelationships(v TeamHierarchyLinkRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *TeamHierarchyLink) GetType() TeamHierarchyLinkType {
	if o == nil {
		var ret TeamHierarchyLinkType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLink) GetTypeOk() (*TeamHierarchyLinkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamHierarchyLink) SetType(v TeamHierarchyLinkType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamHierarchyLink) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *TeamHierarchyLinkAttributes    `json:"attributes"`
		Id            *string                         `json:"id"`
		Relationships *TeamHierarchyLinkRelationships `json:"relationships,omitempty"`
		Type          *TeamHierarchyLinkType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
