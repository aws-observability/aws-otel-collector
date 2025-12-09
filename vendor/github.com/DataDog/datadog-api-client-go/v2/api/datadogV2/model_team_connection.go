// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamConnection A relationship between a Datadog team and a team from another external system.
type TeamConnection struct {
	// Attributes of the team connection.
	Attributes *TeamConnectionAttributes `json:"attributes,omitempty"`
	// The unique identifier of the team connection.
	Id string `json:"id"`
	// Relationships of the team connection.
	Relationships *TeamConnectionRelationships `json:"relationships,omitempty"`
	// Team connection resource type.
	Type TeamConnectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamConnection instantiates a new TeamConnection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamConnection(id string, typeVar TeamConnectionType) *TeamConnection {
	this := TeamConnection{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewTeamConnectionWithDefaults instantiates a new TeamConnection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamConnectionWithDefaults() *TeamConnection {
	this := TeamConnection{}
	var typeVar TeamConnectionType = TEAMCONNECTIONTYPE_TEAM_CONNECTION
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TeamConnection) GetAttributes() TeamConnectionAttributes {
	if o == nil || o.Attributes == nil {
		var ret TeamConnectionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamConnection) GetAttributesOk() (*TeamConnectionAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TeamConnection) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given TeamConnectionAttributes and assigns it to the Attributes field.
func (o *TeamConnection) SetAttributes(v TeamConnectionAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *TeamConnection) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamConnection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TeamConnection) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TeamConnection) GetRelationships() TeamConnectionRelationships {
	if o == nil || o.Relationships == nil {
		var ret TeamConnectionRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamConnection) GetRelationshipsOk() (*TeamConnectionRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TeamConnection) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given TeamConnectionRelationships and assigns it to the Relationships field.
func (o *TeamConnection) SetRelationships(v TeamConnectionRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *TeamConnection) GetType() TeamConnectionType {
	if o == nil {
		var ret TeamConnectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamConnection) GetTypeOk() (*TeamConnectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamConnection) SetType(v TeamConnectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
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
func (o *TeamConnection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *TeamConnectionAttributes    `json:"attributes,omitempty"`
		Id            *string                      `json:"id"`
		Relationships *TeamConnectionRelationships `json:"relationships,omitempty"`
		Type          *TeamConnectionType          `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
