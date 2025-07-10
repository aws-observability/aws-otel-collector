// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamRoutingRulesData Represents the top-level data object for team routing rules, containing the ID, relationships, and resource type.
type TeamRoutingRulesData struct {
	// Specifies the unique identifier of this team routing rules record.
	Id *string `json:"id,omitempty"`
	// Specifies relationships for team routing rules, including rule references.
	Relationships *TeamRoutingRulesDataRelationships `json:"relationships,omitempty"`
	// Team routing rules resource type.
	Type TeamRoutingRulesDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamRoutingRulesData instantiates a new TeamRoutingRulesData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamRoutingRulesData(typeVar TeamRoutingRulesDataType) *TeamRoutingRulesData {
	this := TeamRoutingRulesData{}
	this.Type = typeVar
	return &this
}

// NewTeamRoutingRulesDataWithDefaults instantiates a new TeamRoutingRulesData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamRoutingRulesDataWithDefaults() *TeamRoutingRulesData {
	this := TeamRoutingRulesData{}
	var typeVar TeamRoutingRulesDataType = TEAMROUTINGRULESDATATYPE_TEAM_ROUTING_RULES
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TeamRoutingRulesData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRoutingRulesData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TeamRoutingRulesData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TeamRoutingRulesData) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TeamRoutingRulesData) GetRelationships() TeamRoutingRulesDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret TeamRoutingRulesDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRoutingRulesData) GetRelationshipsOk() (*TeamRoutingRulesDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TeamRoutingRulesData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given TeamRoutingRulesDataRelationships and assigns it to the Relationships field.
func (o *TeamRoutingRulesData) SetRelationships(v TeamRoutingRulesDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *TeamRoutingRulesData) GetType() TeamRoutingRulesDataType {
	if o == nil {
		var ret TeamRoutingRulesDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamRoutingRulesData) GetTypeOk() (*TeamRoutingRulesDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamRoutingRulesData) SetType(v TeamRoutingRulesDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamRoutingRulesData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
func (o *TeamRoutingRulesData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *string                            `json:"id,omitempty"`
		Relationships *TeamRoutingRulesDataRelationships `json:"relationships,omitempty"`
		Type          *TeamRoutingRulesDataType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
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
