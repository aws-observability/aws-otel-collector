// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgConnectionCreate Org connection creation data.
type OrgConnectionCreate struct {
	// Attributes for creating an org connection.
	Attributes OrgConnectionCreateAttributes `json:"attributes"`
	// Relationships for org connection creation.
	Relationships OrgConnectionCreateRelationships `json:"relationships"`
	// Org connection type.
	Type OrgConnectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgConnectionCreate instantiates a new OrgConnectionCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgConnectionCreate(attributes OrgConnectionCreateAttributes, relationships OrgConnectionCreateRelationships, typeVar OrgConnectionType) *OrgConnectionCreate {
	this := OrgConnectionCreate{}
	this.Attributes = attributes
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewOrgConnectionCreateWithDefaults instantiates a new OrgConnectionCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgConnectionCreateWithDefaults() *OrgConnectionCreate {
	this := OrgConnectionCreate{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *OrgConnectionCreate) GetAttributes() OrgConnectionCreateAttributes {
	if o == nil {
		var ret OrgConnectionCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrgConnectionCreate) GetAttributesOk() (*OrgConnectionCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *OrgConnectionCreate) SetAttributes(v OrgConnectionCreateAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value.
func (o *OrgConnectionCreate) GetRelationships() OrgConnectionCreateRelationships {
	if o == nil {
		var ret OrgConnectionCreateRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *OrgConnectionCreate) GetRelationshipsOk() (*OrgConnectionCreateRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *OrgConnectionCreate) SetRelationships(v OrgConnectionCreateRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *OrgConnectionCreate) GetType() OrgConnectionType {
	if o == nil {
		var ret OrgConnectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrgConnectionCreate) GetTypeOk() (*OrgConnectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OrgConnectionCreate) SetType(v OrgConnectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgConnectionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgConnectionCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *OrgConnectionCreateAttributes    `json:"attributes"`
		Relationships *OrgConnectionCreateRelationships `json:"relationships"`
		Type          *OrgConnectionType                `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
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
	if all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = *all.Relationships
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
