// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamRoutingRulesRequestData Holds the data necessary to create or update team routing rules, including attributes, ID, and resource type.
type TeamRoutingRulesRequestData struct {
	// Represents the attributes of a request to update or create team routing rules.
	Attributes *TeamRoutingRulesRequestDataAttributes `json:"attributes,omitempty"`
	// Specifies the unique identifier for this set of team routing rules.
	Id *string `json:"id,omitempty"`
	// Team routing rules resource type.
	Type TeamRoutingRulesRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamRoutingRulesRequestData instantiates a new TeamRoutingRulesRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamRoutingRulesRequestData(typeVar TeamRoutingRulesRequestDataType) *TeamRoutingRulesRequestData {
	this := TeamRoutingRulesRequestData{}
	this.Type = typeVar
	return &this
}

// NewTeamRoutingRulesRequestDataWithDefaults instantiates a new TeamRoutingRulesRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamRoutingRulesRequestDataWithDefaults() *TeamRoutingRulesRequestData {
	this := TeamRoutingRulesRequestData{}
	var typeVar TeamRoutingRulesRequestDataType = TEAMROUTINGRULESREQUESTDATATYPE_TEAM_ROUTING_RULES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TeamRoutingRulesRequestData) GetAttributes() TeamRoutingRulesRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TeamRoutingRulesRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRoutingRulesRequestData) GetAttributesOk() (*TeamRoutingRulesRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TeamRoutingRulesRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given TeamRoutingRulesRequestDataAttributes and assigns it to the Attributes field.
func (o *TeamRoutingRulesRequestData) SetAttributes(v TeamRoutingRulesRequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TeamRoutingRulesRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRoutingRulesRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TeamRoutingRulesRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TeamRoutingRulesRequestData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *TeamRoutingRulesRequestData) GetType() TeamRoutingRulesRequestDataType {
	if o == nil {
		var ret TeamRoutingRulesRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamRoutingRulesRequestData) GetTypeOk() (*TeamRoutingRulesRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamRoutingRulesRequestData) SetType(v TeamRoutingRulesRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamRoutingRulesRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamRoutingRulesRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TeamRoutingRulesRequestDataAttributes `json:"attributes,omitempty"`
		Id         *string                                `json:"id,omitempty"`
		Type       *TeamRoutingRulesRequestDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
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
