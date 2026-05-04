// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldResponseData Data object for an incident user-defined field response.
type IncidentUserDefinedFieldResponseData struct {
	// Attributes of an incident user-defined field.
	Attributes IncidentUserDefinedFieldAttributesResponse `json:"attributes"`
	// The unique identifier of the user-defined field.
	Id string `json:"id"`
	// Relationships of an incident user-defined field.
	Relationships IncidentUserDefinedFieldRelationships `json:"relationships"`
	// The incident user defined fields type.
	Type IncidentUserDefinedFieldType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedFieldResponseData instantiates a new IncidentUserDefinedFieldResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedFieldResponseData(attributes IncidentUserDefinedFieldAttributesResponse, id string, relationships IncidentUserDefinedFieldRelationships, typeVar IncidentUserDefinedFieldType) *IncidentUserDefinedFieldResponseData {
	this := IncidentUserDefinedFieldResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewIncidentUserDefinedFieldResponseDataWithDefaults instantiates a new IncidentUserDefinedFieldResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedFieldResponseDataWithDefaults() *IncidentUserDefinedFieldResponseData {
	this := IncidentUserDefinedFieldResponseData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentUserDefinedFieldResponseData) GetAttributes() IncidentUserDefinedFieldAttributesResponse {
	if o == nil {
		var ret IncidentUserDefinedFieldAttributesResponse
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldResponseData) GetAttributesOk() (*IncidentUserDefinedFieldAttributesResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentUserDefinedFieldResponseData) SetAttributes(v IncidentUserDefinedFieldAttributesResponse) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *IncidentUserDefinedFieldResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentUserDefinedFieldResponseData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value.
func (o *IncidentUserDefinedFieldResponseData) GetRelationships() IncidentUserDefinedFieldRelationships {
	if o == nil {
		var ret IncidentUserDefinedFieldRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldResponseData) GetRelationshipsOk() (*IncidentUserDefinedFieldRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *IncidentUserDefinedFieldResponseData) SetRelationships(v IncidentUserDefinedFieldRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *IncidentUserDefinedFieldResponseData) GetType() IncidentUserDefinedFieldType {
	if o == nil {
		var ret IncidentUserDefinedFieldType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldResponseData) GetTypeOk() (*IncidentUserDefinedFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentUserDefinedFieldResponseData) SetType(v IncidentUserDefinedFieldType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedFieldResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedFieldResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentUserDefinedFieldAttributesResponse `json:"attributes"`
		Id            *string                                     `json:"id"`
		Relationships *IncidentUserDefinedFieldRelationships      `json:"relationships"`
		Type          *IncidentUserDefinedFieldType               `json:"type"`
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
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
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
