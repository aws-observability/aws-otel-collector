// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseLinkAttributes Attributes describing a directional relationship between two entities (cases, incidents, or pages).
type CaseLinkAttributes struct {
	// The UUID of the child (target) entity in the relationship.
	ChildEntityId string `json:"child_entity_id"`
	// The type of the child entity. Allowed values: `CASE`, `INCIDENT`, `PAGE`, `AGENT_CONVERSATION`.
	ChildEntityType string `json:"child_entity_type"`
	// The UUID of the parent (source) entity in the relationship.
	ParentEntityId string `json:"parent_entity_id"`
	// The type of the parent entity. Allowed values: `CASE`, `INCIDENT`, `PAGE`, `AGENT_CONVERSATION`.
	ParentEntityType string `json:"parent_entity_type"`
	// The type of directional relationship. Allowed values: `RELATES_TO` (bidirectional association), `CAUSES` (parent causes child), `BLOCKS` (parent blocks child), `DUPLICATES` (parent duplicates child), `PARENT_OF` (hierarchical), `SUCCESSOR_OF` (sequence), `ESCALATES_TO` (priority escalation).
	Relationship string `json:"relationship"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseLinkAttributes instantiates a new CaseLinkAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseLinkAttributes(childEntityId string, childEntityType string, parentEntityId string, parentEntityType string, relationship string) *CaseLinkAttributes {
	this := CaseLinkAttributes{}
	this.ChildEntityId = childEntityId
	this.ChildEntityType = childEntityType
	this.ParentEntityId = parentEntityId
	this.ParentEntityType = parentEntityType
	this.Relationship = relationship
	return &this
}

// NewCaseLinkAttributesWithDefaults instantiates a new CaseLinkAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseLinkAttributesWithDefaults() *CaseLinkAttributes {
	this := CaseLinkAttributes{}
	return &this
}

// GetChildEntityId returns the ChildEntityId field value.
func (o *CaseLinkAttributes) GetChildEntityId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChildEntityId
}

// GetChildEntityIdOk returns a tuple with the ChildEntityId field value
// and a boolean to check if the value has been set.
func (o *CaseLinkAttributes) GetChildEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChildEntityId, true
}

// SetChildEntityId sets field value.
func (o *CaseLinkAttributes) SetChildEntityId(v string) {
	o.ChildEntityId = v
}

// GetChildEntityType returns the ChildEntityType field value.
func (o *CaseLinkAttributes) GetChildEntityType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChildEntityType
}

// GetChildEntityTypeOk returns a tuple with the ChildEntityType field value
// and a boolean to check if the value has been set.
func (o *CaseLinkAttributes) GetChildEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChildEntityType, true
}

// SetChildEntityType sets field value.
func (o *CaseLinkAttributes) SetChildEntityType(v string) {
	o.ChildEntityType = v
}

// GetParentEntityId returns the ParentEntityId field value.
func (o *CaseLinkAttributes) GetParentEntityId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ParentEntityId
}

// GetParentEntityIdOk returns a tuple with the ParentEntityId field value
// and a boolean to check if the value has been set.
func (o *CaseLinkAttributes) GetParentEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentEntityId, true
}

// SetParentEntityId sets field value.
func (o *CaseLinkAttributes) SetParentEntityId(v string) {
	o.ParentEntityId = v
}

// GetParentEntityType returns the ParentEntityType field value.
func (o *CaseLinkAttributes) GetParentEntityType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ParentEntityType
}

// GetParentEntityTypeOk returns a tuple with the ParentEntityType field value
// and a boolean to check if the value has been set.
func (o *CaseLinkAttributes) GetParentEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentEntityType, true
}

// SetParentEntityType sets field value.
func (o *CaseLinkAttributes) SetParentEntityType(v string) {
	o.ParentEntityType = v
}

// GetRelationship returns the Relationship field value.
func (o *CaseLinkAttributes) GetRelationship() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value
// and a boolean to check if the value has been set.
func (o *CaseLinkAttributes) GetRelationshipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationship, true
}

// SetRelationship sets field value.
func (o *CaseLinkAttributes) SetRelationship(v string) {
	o.Relationship = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseLinkAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["child_entity_id"] = o.ChildEntityId
	toSerialize["child_entity_type"] = o.ChildEntityType
	toSerialize["parent_entity_id"] = o.ParentEntityId
	toSerialize["parent_entity_type"] = o.ParentEntityType
	toSerialize["relationship"] = o.Relationship

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseLinkAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChildEntityId    *string `json:"child_entity_id"`
		ChildEntityType  *string `json:"child_entity_type"`
		ParentEntityId   *string `json:"parent_entity_id"`
		ParentEntityType *string `json:"parent_entity_type"`
		Relationship     *string `json:"relationship"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChildEntityId == nil {
		return fmt.Errorf("required field child_entity_id missing")
	}
	if all.ChildEntityType == nil {
		return fmt.Errorf("required field child_entity_type missing")
	}
	if all.ParentEntityId == nil {
		return fmt.Errorf("required field parent_entity_id missing")
	}
	if all.ParentEntityType == nil {
		return fmt.Errorf("required field parent_entity_type missing")
	}
	if all.Relationship == nil {
		return fmt.Errorf("required field relationship missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"child_entity_id", "child_entity_type", "parent_entity_id", "parent_entity_type", "relationship"})
	} else {
		return err
	}
	o.ChildEntityId = *all.ChildEntityId
	o.ChildEntityType = *all.ChildEntityType
	o.ParentEntityId = *all.ParentEntityId
	o.ParentEntityType = *all.ParentEntityType
	o.Relationship = *all.Relationship

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
