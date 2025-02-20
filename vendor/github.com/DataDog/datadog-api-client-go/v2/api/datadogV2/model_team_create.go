// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamCreate Team create
type TeamCreate struct {
	// Team creation attributes
	Attributes TeamCreateAttributes `json:"attributes"`
	// Relationships formed with the team on creation
	Relationships *TeamCreateRelationships `json:"relationships,omitempty"`
	// Team type
	Type TeamType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamCreate instantiates a new TeamCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamCreate(attributes TeamCreateAttributes, typeVar TeamType) *TeamCreate {
	this := TeamCreate{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewTeamCreateWithDefaults instantiates a new TeamCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamCreateWithDefaults() *TeamCreate {
	this := TeamCreate{}
	var typeVar TeamType = TEAMTYPE_TEAM
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *TeamCreate) GetAttributes() TeamCreateAttributes {
	if o == nil {
		var ret TeamCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TeamCreate) GetAttributesOk() (*TeamCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *TeamCreate) SetAttributes(v TeamCreateAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TeamCreate) GetRelationships() TeamCreateRelationships {
	if o == nil || o.Relationships == nil {
		var ret TeamCreateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamCreate) GetRelationshipsOk() (*TeamCreateRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TeamCreate) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given TeamCreateRelationships and assigns it to the Relationships field.
func (o *TeamCreate) SetRelationships(v TeamCreateRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *TeamCreate) GetType() TeamType {
	if o == nil {
		var ret TeamType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamCreate) GetTypeOk() (*TeamType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamCreate) SetType(v TeamType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
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
func (o *TeamCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *TeamCreateAttributes    `json:"attributes"`
		Relationships *TeamCreateRelationships `json:"relationships,omitempty"`
		Type          *TeamType                `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
