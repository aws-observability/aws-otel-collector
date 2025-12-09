// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamConnectionCreateData Data for creating a team connection.
type TeamConnectionCreateData struct {
	// Attributes of the team connection.
	Attributes *TeamConnectionAttributes `json:"attributes,omitempty"`
	// Relationships of the team connection.
	Relationships *TeamConnectionRelationships `json:"relationships,omitempty"`
	// Team connection resource type.
	Type TeamConnectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamConnectionCreateData instantiates a new TeamConnectionCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamConnectionCreateData(typeVar TeamConnectionType) *TeamConnectionCreateData {
	this := TeamConnectionCreateData{}
	this.Type = typeVar
	return &this
}

// NewTeamConnectionCreateDataWithDefaults instantiates a new TeamConnectionCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamConnectionCreateDataWithDefaults() *TeamConnectionCreateData {
	this := TeamConnectionCreateData{}
	var typeVar TeamConnectionType = TEAMCONNECTIONTYPE_TEAM_CONNECTION
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TeamConnectionCreateData) GetAttributes() TeamConnectionAttributes {
	if o == nil || o.Attributes == nil {
		var ret TeamConnectionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamConnectionCreateData) GetAttributesOk() (*TeamConnectionAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TeamConnectionCreateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given TeamConnectionAttributes and assigns it to the Attributes field.
func (o *TeamConnectionCreateData) SetAttributes(v TeamConnectionAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TeamConnectionCreateData) GetRelationships() TeamConnectionRelationships {
	if o == nil || o.Relationships == nil {
		var ret TeamConnectionRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamConnectionCreateData) GetRelationshipsOk() (*TeamConnectionRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TeamConnectionCreateData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given TeamConnectionRelationships and assigns it to the Relationships field.
func (o *TeamConnectionCreateData) SetRelationships(v TeamConnectionRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *TeamConnectionCreateData) GetType() TeamConnectionType {
	if o == nil {
		var ret TeamConnectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamConnectionCreateData) GetTypeOk() (*TeamConnectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamConnectionCreateData) SetType(v TeamConnectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamConnectionCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
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
func (o *TeamConnectionCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *TeamConnectionAttributes    `json:"attributes,omitempty"`
		Relationships *TeamConnectionRelationships `json:"relationships,omitempty"`
		Type          *TeamConnectionType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
