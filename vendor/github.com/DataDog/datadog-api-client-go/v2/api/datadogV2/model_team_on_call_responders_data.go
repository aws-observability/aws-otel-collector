// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamOnCallRespondersData Defines the main on-call responder object for a team, including relationships and metadata.
type TeamOnCallRespondersData struct {
	// Unique identifier of the on-call responder configuration.
	Id *string `json:"id,omitempty"`
	// Relationship objects linked to a team's on-call responder configuration, including escalations and responders.
	Relationships *TeamOnCallRespondersDataRelationships `json:"relationships,omitempty"`
	// Represents the resource type for a group of users assigned to handle on-call duties within a team.
	Type TeamOnCallRespondersDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamOnCallRespondersData instantiates a new TeamOnCallRespondersData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamOnCallRespondersData(typeVar TeamOnCallRespondersDataType) *TeamOnCallRespondersData {
	this := TeamOnCallRespondersData{}
	this.Type = typeVar
	return &this
}

// NewTeamOnCallRespondersDataWithDefaults instantiates a new TeamOnCallRespondersData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamOnCallRespondersDataWithDefaults() *TeamOnCallRespondersData {
	this := TeamOnCallRespondersData{}
	var typeVar TeamOnCallRespondersDataType = TEAMONCALLRESPONDERSDATATYPE_TEAM_ONCALL_RESPONDERS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TeamOnCallRespondersData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamOnCallRespondersData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TeamOnCallRespondersData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TeamOnCallRespondersData) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TeamOnCallRespondersData) GetRelationships() TeamOnCallRespondersDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret TeamOnCallRespondersDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamOnCallRespondersData) GetRelationshipsOk() (*TeamOnCallRespondersDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TeamOnCallRespondersData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given TeamOnCallRespondersDataRelationships and assigns it to the Relationships field.
func (o *TeamOnCallRespondersData) SetRelationships(v TeamOnCallRespondersDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *TeamOnCallRespondersData) GetType() TeamOnCallRespondersDataType {
	if o == nil {
		var ret TeamOnCallRespondersDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamOnCallRespondersData) GetTypeOk() (*TeamOnCallRespondersDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamOnCallRespondersData) SetType(v TeamOnCallRespondersDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamOnCallRespondersData) MarshalJSON() ([]byte, error) {
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
func (o *TeamOnCallRespondersData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *string                                `json:"id,omitempty"`
		Relationships *TeamOnCallRespondersDataRelationships `json:"relationships,omitempty"`
		Type          *TeamOnCallRespondersDataType          `json:"type"`
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
